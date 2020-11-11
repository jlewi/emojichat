// main provides a simple http server for serving the front end and javascript.
package main

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/jlewi/emojichat/go/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
)

func initLogger() {
	// TODO(jlewi): Make the verbosity level configurable.

	// Start with a production logger config.
	config := zap.NewProductionConfig()

	// TODO(jlewi): In development mode we should use the console encoder as opposed to json formatted logs.

	// Increment the logging level.
	// TODO(jlewi): Make this a flag.
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	zapLog, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("Could not create zap instance (%v)?", err))
	}
	log = zapr.NewLogger(zapLog)

	zap.ReplaceGlobals(zapLog)
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&options.port, "port", "", "80", "The port to serve on.")
	serveCmd.Flags().StringVarP(&options.staticFiles, "static-files", "", "", "Optional directory of files to serve statically")
}

type cliOptions struct {
	port          string
	staticFiles string
}

var (
	Verbose  int
	options  = cliOptions{}

	rootCmd = &cobra.Command{
		Short: "Serve",
	}

	log logr.Logger

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start webserver.",
		Long:  `starts the chatbot server`,
		Run: func(cmd *cobra.Command, args []string) {
			initLogger()
			s, err := server.New(options.staticFiles, log)

			if err != nil {
				log.Error(err, "Failed to initialize server")
				return
			}

			address := ":" + options.port
			log.Info("Serving", "address", address)
			if err := http.ListenAndServe(address, s.Router); err != nil {
				log.Error(err, "Failed to listen and serve")
				return
			}
		},
	}
)

func main() {
	rootCmd.Execute()
}
