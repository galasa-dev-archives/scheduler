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
		Use:   "get",
		Short: "get app from the scheduler",
		Long:  "get app from the scheduler",
		Args:  cobra.ExactArgs(1),
		Run:   appGetExecute,
	}

	appGetOutputType string

)

func init() {
	appGetCmd.Flags().StringVarP(&appGetOutputType, "output", "o", "", "output type,  yaml/json")
	appGetCmd.MarkFlagRequired("output")

	appCmd.AddCommand(appGetCmd)
}

func appGetExecute(cmd *cobra.Command, args []string) {
	currentName := args[0]

	app.AppGet(currentName, appGetOutputType)
}
