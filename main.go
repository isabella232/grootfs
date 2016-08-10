package main

import (
	"os"
	"strings"

	"code.cloudfoundry.org/grootfs/commands"
	"code.cloudfoundry.org/grootfs/graph"
	"code.cloudfoundry.org/lager"

	"github.com/urfave/cli"
)

const GraphPath = "/tmp/grootfs"

func main() {
	grootfs := cli.NewApp()
	grootfs.Name = "grootfs"
	grootfs.Usage = "I am Groot!"
	grootfs.Version = "0.1.0"

	grootfs.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "graph",
			Usage: "Path to the graph directory",
			Value: GraphPath,
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "Set logging level. <info|debug|error|fatal>",
			Value: "fatal",
		},
		cli.StringFlag{
			Name:  "log-file",
			Usage: "Forward logs to file",
		},
	}

	grootfs.Commands = []cli.Command{
		commands.CreateCommand,
		commands.DeleteCommand,
		commands.UntarCommand,
	}

	grootfs.Before = func(ctx *cli.Context) error {
		graphPath := ctx.String("graph")

		cli.ErrWriter = os.Stdout

		logger := configureLog(ctx)

		configurer := graph.NewConfigurer()
		return configurer.Ensure(logger, graphPath)
	}

	grootfs.Run(os.Args)
}

func configureLog(ctx *cli.Context) lager.Logger {
	logFile := ctx.GlobalString("log-file")
	logLevel := ctx.String("log-level")
	logWriter := os.Stderr

	if logFile != "" {
		logWriter, _ = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	}

	logger := lager.NewLogger("grootfs")
	lagerLogLevel := translateLogLevel(logLevel)

	logger.RegisterSink(lager.NewWriterSink(logWriter, lagerLogLevel))
	ctx.App.Metadata["logger"] = logger

	return logger
}

func translateLogLevel(logLevel string) lager.LogLevel {
	switch strings.ToUpper(logLevel) {
	case "FATAL":
		return lager.FATAL
	case "DEBUG":
		return lager.DEBUG
	case "INFO":
		return lager.INFO
	case "ERROR":
		return lager.ERROR
	}

	return lager.FATAL
}
