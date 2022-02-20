/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/app"
)
var (
    appCreateCmd = &cobra.Command{
		Use:   "create [appname]",
		Short: "create a new app",
		Long:  "create a new app",
		Args:  cobra.ExactArgs(1),
		Run:   appCreateExecute,
	}

)

func init() {
	appCmd.AddCommand(appCreateCmd)
}

func appCreateExecute(cmd *cobra.Command, args []string) {
	newName := args[0]

	app.AppCreate(newName)
}
