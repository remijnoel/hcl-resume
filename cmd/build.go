package cmd

import (
	_ "embed"

	"github.com/remijnoel/hcl-resume/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the resume from source files",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Building the resume from source files...")
		sourceFile, _ := cmd.Flags().GetString("source")
		outputDir, _ := cmd.Flags().GetString("output-dir")
		err := internal.RunPDFLatex(sourceFile, outputDir)
		if err != nil {
			log.Fatalf("Failed to run pdflatex: %v", err)
		}
		log.Info("Resume built successfully")
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringP("source", "s", "resume.tex", "Source file to build the resume from")
	buildCmd.Flags().StringP("output-dir", "o", ".", "Output directory for the built resume")
	buildCmd.Flags().BoolP("pdflatex", "p", true, "Use pdflatex to build the resume (default true)")
}
