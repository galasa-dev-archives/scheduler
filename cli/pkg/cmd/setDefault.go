/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/util"
)
var (
    setDefaultCmd = &cobra.Command{
		Use:   "set-default [app]",
		Short: "Set the default app",
		Long:  "Set the default app",
		Args:  cobra.ExactArgs(1),
		Run:   setDefaultExecute,
	}

)

func init() {
	appCmd.AddCommand(setDefaultCmd)
}

func setDefaultExecute(cmd *cobra.Command, args []string) {
	app := args[0]

	util.SetDefaultApp(app)
}
