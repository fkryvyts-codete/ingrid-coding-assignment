// Package cmd contains commands that can be run from the command line
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web service",
	Long:  `You can change values in config.yaml to customize the behavior of the app`,
	Run: func(_ *cobra.Command, _ []string) {
		httpAddr := viper.GetString("server.listen")

		var logger log.Logger
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)

		mux := http.NewServeMux()

		http.Handle("/", accessControl(mux))

		errs := make(chan error, 2)
		go func() {
			logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
			errs <- http.ListenAndServe(httpAddr, nil)
		}()
		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT)
			errs <- fmt.Errorf("%s", <-c)
		}()

		logger.Log("terminated", <-errs)
	},
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
