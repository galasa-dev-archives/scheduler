/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"
)
var (
    getCmd = &cobra.Command{
		Use:   "get",
		Short: "Display one or more scheduler resources",
		Long:  "Display one or more scheduler resources",
	}

)

func init() {
	rootCmd.AddCommand(getCmd)
}
