package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ldd27/go-starter-kit/internal/conf"
	"github.com/ldd27/go-starter-kit/internal/setup"
	"github.com/ldd27/go-starter-kit/internal/setup/g"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	RootCMD.AddCommand(initDbCMD())
}

func initDbCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init_db",
		Short: "init_db",
		Long:  "init_db",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := time.Now()
			tasks := []func(conf.Conf) error{
				setup.SetupZAP,
				setup.SetupDB,
				setup.AutoMigrate,
				setup.SetupHealth,
			}

			for _, task := range tasks {
				if err := task(g.C); err != nil {
					panic(err)
				}
			}

			logger := zap.L()
			logger.Info("init success", zap.String("duration", time.Since(t).String()))
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			logger.Info(fmt.Sprintf("Received signal %s", <-ch))

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup.SetupViper(cmd.Flags().Lookup("config").Value.String())
		},
	}

	cmd.PersistentFlags().StringP("config", "c", "", "config file path")

	return cmd
}
