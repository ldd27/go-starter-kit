package cmd

import (
	"context"
	"sync"

	"github.com/ldd27/go-starter-kit/internal/conf"
	"github.com/ldd27/go-starter-kit/internal/setup"
	"github.com/ldd27/go-starter-kit/internal/setup/g"
	"github.com/spf13/cobra"
)

func init() {
	RootCMD.AddCommand(serveCMD())
}

func serveCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "run server",
		Long:  "run server",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks := []func(conf.Conf) error{
				setup.SetupZAP,
				setup.SetupDB,
			}

			for _, task := range tasks {
				if err := task(g.C); err != nil {
					panic(err)
				}
			}

			wg := sync.WaitGroup{}
			_, cancel := context.WithCancel(context.Background())
			defer cancel()

			if err := setup.SetupEcho(g.C, &wg, cancel); err != nil {
				return err
			}

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup.SetupViper(cmd.Flags().Lookup("config").Value.String())
		},
	}

	cmd.PersistentFlags().StringP("config", "c", "", "config file path")

	return cmd
}
