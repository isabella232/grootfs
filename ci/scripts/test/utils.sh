function mount_btrfs {
  # Configure cgroup
  mount -tcgroup -odevices cgroup:devices /sys/fs/cgroup
  devices_mount_info=$(cat /proc/self/cgroup | grep devices)
  devices_subdir=$(echo $devices_mount_info | cut -d: -f3)
  echo 'b 7:* rwm' > /sys/fs/cgroup/devices.allow
  echo 'b 7:* rwm' > /sys/fs/cgroup${devices_subdir}/devices.allow

  # Setup loop devices
  for i in {0..256}
  do
    mknod -m777 /dev/loop$i b 7 $i
  done

  # Make BTRFS volume
  truncate -s 1G /btrfs_volume
  mkfs.btrfs /btrfs_volume

  # Mount BTRFS
  mkdir /mnt/btrfs
  mount -t btrfs -o user_subvol_rm_allowed,rw /btrfs_volume /mnt/btrfs
  chmod 777 -R /mnt/btrfs
  btrfs quota enable /mnt/btrfs
}

function compile_drax {
  tmp_dir=$(mktemp -d)
  go build -o $tmp_dir/drax code.cloudfoundry.org/grootfs/store/volume_driver/drax
  echo $tmp_dir/drax
}

function cleanup_drax {
  drax_path=$1
  rm -Rf $(dirname $drax_path)
}

function setup_drax {
  drax_path=$1
  cp $drax_path /usr/local/bin/drax
  chmod u+s /usr/local/bin/drax
}

function sudo_mount_btrfs {
  local MOUNT_BTRFS_FUNC=$(declare -f mount_btrfs)
  sudo bash -c "$MOUNT_BTRFS_FUNC; mount_btrfs"
}

function sudo_setup_drax {
  drax_path=$1

  local SETUP_DRAX_FUNC=$(declare -f setup_drax)
  sudo bash -c "$SETUP_DRAX_FUNC; setup_drax $drax_path"
}

function move_to_gopath {
  grootfsPath=/go/src/code.cloudfoundry.org/grootfs
  rmdir $grootfsPath
  ln -s $PWD/src/code.cloudfoundry.org/grootfs $grootfsPath
  cd $grootfsPath
}