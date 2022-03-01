/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/level"
)
var (
    levelDescribeCmd = &cobra.Command{
		Use:   "describe [levelname]",
		Short: "describe a new level",
		Long:  "describe a new level",
		Args:  cobra.ExactArgs(1),
		Run:   levelDescribeExecute,
	}

)

func init() {
	levelCmd.AddCommand(levelDescribeCmd)
}

func levelDescribeExecute(cmd *cobra.Command, args []string) {
	oldName := args[0]

	level.LevelDescribe(levelApp, oldName)
}
