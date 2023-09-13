package summon

import (
	"github.com/gateway-fm/warp_swagger/internal"
	"github.com/spf13/cobra"
)

func Cmd(app *internal.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "summon",
		Short: "summon",
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO: Rebuild with cobra flags
			if err := app.Summon(); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
