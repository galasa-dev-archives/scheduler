/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"
)
var (
    levelCmd = &cobra.Command{
		Use:   "level",
		Short: "Control an level",
		Long:  "Control an level",
	}

	levelApp string

)

func init() {
	levelCmd.Flags().StringVarP(&levelApp, "app", "a", "", "the app this level applies to")

	rootCmd.AddCommand(levelCmd)
}
