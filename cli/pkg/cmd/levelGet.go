/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/level"
)
var (
    levelGetCmd = &cobra.Command{
		Use:   "get",
		Short: "get level from the scheduler",
		Long:  "get level from the scheduler",
		Args:  cobra.ExactArgs(1),
		Run:   levelGetExecute,
	}

	levelGetOutputType string

)

func init() {
	levelGetCmd.Flags().StringVarP(&levelGetOutputType, "output", "o", "", "output type,  yaml/json")
	levelGetCmd.MarkFlagRequired("output")

	levelCmd.AddCommand(levelGetCmd)
}

func levelGetExecute(cmd *cobra.Command, args []string) {
	currentName := args[0]

	level.LevelGet(levelApp, currentName, levelGetOutputType)
}
