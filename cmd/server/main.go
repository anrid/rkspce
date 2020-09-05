package main

import (
	"github.com/anrid/rkspce/internal/config"
	"github.com/anrid/rkspce/internal/controller"
	"github.com/anrid/rkspce/internal/pkg/httpserver"
	"go.uber.org/zap"
)

func main() {
	// Flags.

	// Setup config.
	c := config.New()

	// Setup zap logger.
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer log.Sync()
	zap.ReplaceGlobals(log) // Allows us to optionally use the zap.L() global instance.

	// Setup controller.
	ctrl := controller.New()

	// Setup HTTP server.
	serv := httpserver.New()

	// Add routes.
	serv.AddRoutes(ctrl.SetupRoutes)

	// Start HTTP server.
	serv.Start(c.Host)
}
