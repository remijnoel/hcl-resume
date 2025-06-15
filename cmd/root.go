package cmd

import (
	"github.com/openai/openai-go"
	"github.com/remijnoel/hcl-resume/internal"
	"github.com/spf13/cobra"
)

var ai *openai.Client = internal.NewOpenAIClient()

var RootCmd = &cobra.Command{
	Use:   "hcl-resume",
	Short: "Manage your resume using HCL",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config()
		return nil
	},

}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	// RootCmd.PersistentFlags().Bool("debug", false, "Enable verbose logging")
}
