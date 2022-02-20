/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"
)
var (
    appCmd = &cobra.Command{
		Use:   "app",
		Short: "Control an app or application",
		Long:  "Control an app or application",
	}

)

func init() {
	rootCmd.AddCommand(appCmd)
}
