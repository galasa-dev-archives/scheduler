/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/util"
)
var (
    loginCmd = &cobra.Command{
		Use:   "login [hostname]",
		Short: "login to a scheduler api server",
		Long:  "login to a scheduler api server",
		Args:  cobra.ExactArgs(1),
		Run:   loginExecute,
	}

)

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginExecute(cmd *cobra.Command, args []string) {
	hostname := args[0]

	util.Login(hostname)
}
