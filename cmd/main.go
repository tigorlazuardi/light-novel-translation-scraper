package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use: "lnt-scraper",
	RunE: func(cmd *cobra.Command, _args []string) error {
		return errors.New("unimplemented")
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Exec(ctx context.Context) error {
	return mainCmd.ExecuteContext(ctx)
}
