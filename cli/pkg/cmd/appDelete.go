/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/app"
)
var (
    appDeleteCmd = &cobra.Command{
		Use:   "delete [appname]",
		Short: "delete a new app",
		Long:  "delete a new app",
		Args:  cobra.ExactArgs(1),
		Run:   appDeleteExecute,
	}

)

func init() {
	appCmd.AddCommand(appDeleteCmd)
}

func appDeleteExecute(cmd *cobra.Command, args []string) {
	oldName := args[0]

	app.AppDelete(oldName)
}
