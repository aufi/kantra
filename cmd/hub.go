package cmd

import (
	"context"

	"github.com/konveyor-ecosystem/kantra/cmd/hub"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

type hubCommand struct {
	log     logr.Logger
	cleanup bool
}

func NewHubCommand(log logr.Logger) *cobra.Command {
	hubCmd := &hubCommand{
		log:     log,
		cleanup: true,
	}

	hubCommand := &cobra.Command{
		Use: "hub",
		Short: "Interact with Konveyor Hub component",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := hubCmd.RunPing(cmd.Context())
			if err != nil {
				log.Error(err, "failed to execute hub check")
				return err
			}
			return nil
		},
	}

	exportCommand := &cobra.Command{
		Use: "export",

		Short: "Export Hub resources to local dump",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := hub.Connect()
			if err != nil {
				log.Error(err, "hub connection failed")
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := hub.Export()
			if err != nil {
				log.Error(err, "failed to execute hub export")
				return err
			}
			return nil
		},
	}
	hubCommand.AddCommand(exportCommand)

	importCommand := &cobra.Command{
		Use: "import",	// add arg for clean before import

		Short: "Import local dump to Hub",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := hub.Connect()
			if err != nil {
				log.Error(err, "hub connection failed")
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := hub.Import()
			if err != nil {
				log.Error(err, "failed to execute hub import")
				return err
			}
			return nil
		},
	}
	hubCommand.AddCommand(importCommand)

	return hubCommand
}

func (w *hubCommand) RunPing(ctx context.Context) error {
	err := hub.Connect()
	if err != nil {
		w.log.V(1).Error(err, "Cannot reach Hub.")
	} else {
		w.log.V(1).Info("Hub is alive.")
	}

	return nil
}
