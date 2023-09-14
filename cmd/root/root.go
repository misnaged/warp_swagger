package root

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/misnaged/warp_swagger/internal"
)

func Cmd(app *internal.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "warp_swagger",
		Short:            "warp_swagger",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgs := []any{app.WarpCfg(), app.SwagCfg()}
			return initializeConfig(cmd, cfgs)
		},
	}
	return cmd
}

func initializeConfig(cmd *cobra.Command, cfgs []any) error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("read config file: %w", err)
		}
	}

	//   set config via env vars
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	bindFlags(cmd)
	var unmarshall error
	for i := range cfgs {
		unmarshall = viper.Unmarshal(&cfgs[i])
	}
	return unmarshall
}

func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
