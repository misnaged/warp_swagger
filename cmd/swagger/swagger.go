package swagger

import (
	"github.com/gateway-fm/warp_swagger/internal"
	"github.com/spf13/cobra"
)

func Cmd(app *internal.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.SwaggerGenerate(args); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
