package main

import (
	"github.com/squillace/porter-jq/pkg/jq"
	"github.com/spf13/cobra"
)

func buildBuildCommand(m *jq.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
