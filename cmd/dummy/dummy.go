package dummy

import (
	"github.com/misnaged/warp_swagger/internal"
	"github.com/spf13/cobra"
)

func Cmd(app *internal.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dummy",
		Short: "dummy",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.CallDummy(); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
