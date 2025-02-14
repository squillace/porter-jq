package main

import (
	"github.com/squillace/porter-jq/pkg/jq"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *jq.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
