/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var (
    statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Fetch the current status of the scheduler",
		Long:  "Report on the current configuration and what test runs are active",
		Run:   statusExecute,
	}

)

func init() {
	RootCmd.AddCommand(statusCmd)

}

func statusExecute(cmd *cobra.Command, args []string) {
	fmt.Println("No idea what you are talking about :-)")
}
