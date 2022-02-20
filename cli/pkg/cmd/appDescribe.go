/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/app"
)
var (
    appDescribeCmd = &cobra.Command{
		Use:   "describe [appname]",
		Short: "describe a new app",
		Long:  "describe a new app",
		Args:  cobra.ExactArgs(1),
		Run:   appDescribeExecute,
	}

)

func init() {
	appCmd.AddCommand(appDescribeCmd)
}

func appDescribeExecute(cmd *cobra.Command, args []string) {
	oldName := args[0]

	app.AppDescribe(oldName)
}
