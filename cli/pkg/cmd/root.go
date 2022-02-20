/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "galasasch",
	Short: "Creating and maintaining Galasa test schdules",
	Long:  "",
	Version: "unknowncliversion-unknowngithash",
}

func Execute() {
	reportVersion()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func reportVersion() {
	fmt.Printf("Galasa Scheduler - version %s\n\n", rootCmd.Version)
}
