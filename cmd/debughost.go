package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the resume to LaTeX",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exporting resume to LaTeX...")
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new resume",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Creating a new resume...")
		// includeScaffold, _ := cmd.Flags().GetBool("include-scaffold")
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)

	RootCmd.AddCommand(newCmd)
	RootCmd.Flags().BoolP("include-scaffold", "s", false, "Include scaffold in the new resume")
}
