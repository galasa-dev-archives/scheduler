/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	resource "galasa.dev/scheduler/pkg/resource"
)
var (
    applyCmd = &cobra.Command{
		Use:   "apply",
		Short: "apply resoources in scheduler from a yaml file",
		Long:  "Apply App/Level/Levelset in the scheduler",
		Run:   applyExecute,
	}

	applyFilename string

)

func init() {
	applyCmd.Flags().StringVarP(&applyFilename, "file", "f", "", "yaml to apply")
	applyCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(applyCmd)
}

func applyExecute(cmd *cobra.Command, args []string) {
	resource.Apply(createFilename)
}
