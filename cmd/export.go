package cmd

import (
	"os"

	_ "embed"

	"github.com/remijnoel/hcl-resume/internal"
	"github.com/remijnoel/hcl-resume/resume"
	"github.com/remijnoel/hcl-resume/template"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)


var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the resume to LaTeX",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Exporting the resume to LaTeX...")
		r, err := resume.LoadFromDir(".")
		if err != nil {
			log.Fatalf("Failed to load resume: %v", err)
		}

		log.Infof("Rendering in LaTeX format")
		latexContent, err := internal.RenderToTeX(r, template.ResumeTemplate)
		if err != nil {
			log.Fatalf("Failed to render resume: %v", err)
		}
		outPath := "resume.tex"
		err = os.WriteFile(outPath, []byte(latexContent), 0644)
		if err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}
		log.Infof("Resume exported to %s", outPath)


	},
}



func init() {
	RootCmd.AddCommand(exportCmd)
}
