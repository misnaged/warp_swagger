package main

import (
	"github.com/gateway-fm/warp_swagger/cmd/dummy"
	"github.com/gateway-fm/warp_swagger/cmd/root"
	"github.com/gateway-fm/warp_swagger/cmd/summon"
	"github.com/gateway-fm/warp_swagger/cmd/swagger"
	"github.com/gateway-fm/warp_swagger/internal"
	"github.com/misnaged/annales/logger"
	"os"
)

func main() {
	app, err := internal.NewApplication()

	if err != nil {
		logger.Log().Errorf("An error occurred %v", err)
		os.Exit(1)
	}

	rootCmd := root.Cmd(app)
	rootCmd.AddCommand(summon.Cmd(app))
	rootCmd.AddCommand(swagger.Cmd(app))
	rootCmd.AddCommand(dummy.Cmd(app))

	if err = rootCmd.Execute(); err != nil {
		logger.Log().Errorf("An error occurred %v", err)
		os.Exit(1)
	}
}
