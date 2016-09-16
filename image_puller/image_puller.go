package image_puller

import (
	"fmt"
	"io"
	"net/url"

	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/lager"
	specsv1 "github.com/opencontainers/image-spec/specs-go/v1"
)

//go:generate counterfeiter . VolumeDriver
//go:generate counterfeiter . Fetcher
//go:generate counterfeiter . Unpacker

type UnpackSpec struct {
	Stream      io.ReadCloser
	TargetPath  string
	UIDMappings []groot.IDMappingSpec
	GIDMappings []groot.IDMappingSpec
}

type LayerDigest struct {
	BlobID        string
	ChainID       string
	ParentChainID string
}

type ImageInfo struct {
	LayersDigest []LayerDigest
	Config       specsv1.Image
}

type VolumeDriver interface {
	Path(logger lager.Logger, id string) (string, error)
	Create(logger lager.Logger, parentID, id string) (string, error)
	DestroyVolume(logger lager.Logger, id string) error
}

type Fetcher interface {
	ImageInfo(logger lager.Logger, imageURL *url.URL) (ImageInfo, error)
	StreamBlob(logger lager.Logger, imageURL *url.URL, source string) (io.ReadCloser, int64, error)
}

type Unpacker interface {
	Unpack(logger lager.Logger, spec UnpackSpec) error
}

type ImagePuller struct {
	fetcher      Fetcher
	unpacker     Unpacker
	volumeDriver VolumeDriver
}

func NewImagePuller(fetcher Fetcher, unpacker Unpacker, volumeDriver VolumeDriver) *ImagePuller {
	return &ImagePuller{
		fetcher:      fetcher,
		unpacker:     unpacker,
		volumeDriver: volumeDriver,
	}
}

func (p *ImagePuller) Pull(logger lager.Logger, spec groot.ImageSpec) (groot.BundleSpec, error) {
	logger = logger.Session("image-pulling", lager.Data{"spec": spec})
	logger.Info("start")
	defer logger.Info("end")
	var err error

	imageInfo, err := p.fetcher.ImageInfo(logger, spec.ImageSrc)
	if err != nil {
		return groot.BundleSpec{}, fmt.Errorf("fetching list of digests: %s", err)
	}
	logger.Debug("fetched-layers-digests", lager.Data{"digests": imageInfo.LayersDigest})

	var (
		volumePath string
		totalSize  int64
	)
	for _, digest := range imageInfo.LayersDigest {
		volumePath, err = p.volumeDriver.Path(logger, wrapVolumeID(spec, digest.ChainID))
		if err == nil {
			logger.Debug("volume-exists", lager.Data{
				"volumePath":    volumePath,
				"blobID":        digest.BlobID,
				"chainID":       digest.ChainID,
				"parentChainID": digest.ParentChainID,
			})
			continue
		}

		stream, size, err := p.fetcher.StreamBlob(logger, spec.ImageSrc, digest.BlobID)
		if err != nil {
			return groot.BundleSpec{}, fmt.Errorf("streaming blob `%s`: %s", digest.BlobID, err)
		}
		totalSize += size
		if !spec.ExcludeImageFromQuota && spec.DiskLimit != 0 && spec.DiskLimit < totalSize {
			err := fmt.Errorf("exceeded disk quota, using %d out of %d bytes", totalSize, spec.DiskLimit)
			logger.Error("blob-streaming-failed", err, lager.Data{
				"totalSize":             totalSize,
				"diskLimit":             spec.DiskLimit,
				"excludeImageFromQuota": spec.ExcludeImageFromQuota,
				"blobID":                digest.BlobID,
			})
			return groot.BundleSpec{}, err
		}

		logger.Debug("got-stream-for-blob", lager.Data{
			"size":                  size,
			"diskLimit":             spec.DiskLimit,
			"excludeImageFromQuota": spec.ExcludeImageFromQuota,
			"blobID":                digest.BlobID,
			"chainID":               digest.ChainID,
			"parentChainID":         digest.ParentChainID,
		})

		volumePath, err = p.volumeDriver.Create(logger,
			wrapVolumeID(spec, digest.ParentChainID),
			wrapVolumeID(spec, digest.ChainID),
		)
		if err != nil {
			return groot.BundleSpec{}, fmt.Errorf("creating volume for layer `%s`: %s", digest.BlobID, err)
		}
		logger.Debug("volume-created", lager.Data{
			"volumePath":    volumePath,
			"blobID":        digest.BlobID,
			"chainID":       digest.ChainID,
			"parentChainID": digest.ParentChainID,
		})

		unpackSpec := UnpackSpec{
			TargetPath:  volumePath,
			Stream:      stream,
			UIDMappings: spec.UIDMappings,
			GIDMappings: spec.GIDMappings,
		}
		if err := p.unpacker.Unpack(logger, unpackSpec); err != nil {
			if err := p.volumeDriver.DestroyVolume(logger, digest.ChainID); err != nil {
				logger.Error("volume-cleanup-failed", err, lager.Data{
					"blobID":        digest.BlobID,
					"chainID":       digest.ChainID,
					"parentChainID": digest.ParentChainID,
				})
			}
			return groot.BundleSpec{}, fmt.Errorf("unpacking layer `%s`: %s", digest.BlobID, err)
		}
		logger.Debug("layer-unpacked", lager.Data{
			"blobID":        digest.BlobID,
			"chainID":       digest.ChainID,
			"parentChainID": digest.ParentChainID,
		})
	}

	bundleSpec := groot.BundleSpec{
		Image:      imageInfo.Config,
		VolumePath: volumePath,
	}
	return bundleSpec, nil
}

func wrapVolumeID(spec groot.ImageSpec, volumeID string) string {
	if volumeID == "" {
		return ""
	}

	if len(spec.UIDMappings) > 0 || len(spec.GIDMappings) > 0 {
		return fmt.Sprintf("%s-namespaced", volumeID)
	}

	return volumeID
}
