/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/app"
)
var (
    appListCmd = &cobra.Command{
		Use:   "list",
		Short: "list apps defined to the scheduler",
		Long:  "list apps defined to the scheduler",
		Run:   appListExecute,
	}

)

func init() {
	appCmd.AddCommand(appListCmd)
}

func appListExecute(cmd *cobra.Command, args []string) {
	app.AppListAll()
}
