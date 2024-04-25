package main

import (
	"log/slog"

	"github.com/tarannum22/olive-tree-pgp/jobs"
	"github.com/tarannum22/olive-tree-pgp/logger"
	"github.com/tarannum22/olive-tree-pgp/server"
	"github.com/tarannum22/olive-tree-pgp/utils"
)

var (
	globalLogger *slog.Logger
	config       utils.AppConfig
)

func init() {
	globalLogger = logger.InitGlobalLogger()
	slog.SetDefault(globalLogger)
	utils.LoadAppConfig("config.yaml", &config)
}

func main() {
	slog.Info("WatermelonPGP has started")

	jobs.InitJobs()
	//App()
	server.StartHTTPServer(8080)

	slog.Info("Stopping WatermelonPGP")
}
