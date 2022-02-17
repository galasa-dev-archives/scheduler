/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/app"
)
var (
    appGetCmd = &cobra.Command{
		Use:   "app",
		Short: "get an app",
		Long:  "get an app",
		Run:   appGetExecute,
	}

)

func init() {
	getCmd.AddCommand(appGetCmd)
}

func appGetExecute(cmd *cobra.Command, args []string) {
	app.GetAllApps()
}
