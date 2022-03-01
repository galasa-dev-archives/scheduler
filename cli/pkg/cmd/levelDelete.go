/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/level"
)
var (
    levelDeleteCmd = &cobra.Command{
		Use:   "delete [levelname]",
		Short: "delete a new level",
		Long:  "delete a new level",
		Args:  cobra.ExactArgs(1),
		Run:   levelDeleteExecute,
	}

)

func init() {
	levelCmd.AddCommand(levelDeleteCmd)
}

func levelDeleteExecute(cmd *cobra.Command, args []string) {
	oldName := args[0]

	level.LevelDelete(levelApp, oldName)
}
