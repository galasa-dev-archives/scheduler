/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	resource "galasa.dev/scheduler/pkg/resource"
)
var (
    createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create resoources in scheduler from a yaml file",
		Long:  "Create App/Level/Levelset in the scheduler",
		Run:   createExecute,
	}

	createFilename string

)

func init() {
	createCmd.Flags().StringVarP(&createFilename, "file", "f", "", "yaml to create")
	createCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(createCmd)
}

func createExecute(cmd *cobra.Command, args []string) {
	resource.Create(createFilename)
}
