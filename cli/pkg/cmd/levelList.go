/*
 * Copyright contributors to the Galasa project
 */
package cmd

import (
	"github.com/spf13/cobra"

	"galasa.dev/scheduler/pkg/level"
)
var (
    levelListCmd = &cobra.Command{
		Use:   "list",
		Short: "list levels defined to the scheduler",
		Long:  "list levels defined to the scheduler",
		Run:   levelListExecute,
	}

)

func init() {
	levelCmd.AddCommand(levelListCmd)
}

func levelListExecute(cmd *cobra.Command, args []string) {
	level.LevelListAll(levelApp)
}
