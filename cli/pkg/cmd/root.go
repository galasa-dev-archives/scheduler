/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "galasasch",
	Short: "Creating and maintaining Galasa test schdules",
	Long:  "",
	Version: "0.0.1",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
