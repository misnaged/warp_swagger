package main

import (
	"github.com/misnaged/annales/logger"
	"github.com/misnaged/warp_swagger/cmd/dummy"
	"github.com/misnaged/warp_swagger/cmd/root"
	"github.com/misnaged/warp_swagger/cmd/summon"
	"github.com/misnaged/warp_swagger/cmd/swagger"
	"github.com/misnaged/warp_swagger/internal"
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
