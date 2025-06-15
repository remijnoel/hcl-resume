package cmd

import (
	"os"

	"github.com/remijnoel/hcl-resume/template"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new resume",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Creating a new resume...")
		includeScaffolding, _ := cmd.Flags().GetBool("include-scaffolding")

		if includeScaffolding {
			log.Info("Including scaffolding files for a new resume...")
			// Write the basic scaffolding files
			scaffolding := template.HCLScaffolding
			fileName := "resume.hcl"
			err := os.WriteFile(fileName, []byte(scaffolding), 0644)
			if err != nil {
				log.Fatalf("Failed to write scaffolding file: %v", err)
			}
			log.Infof("Scaffolding written to %s", fileName)
		}

		// Write the template files to the current directory
		latexTemplate := template.LateXResumeTemplate
		fileName := "resume.cls"
		err := os.WriteFile(fileName, []byte(latexTemplate), 0644)
		if err != nil {
			log.Fatalf("Failed to write template file: %v", err)
		}
		log.Infof("Template written to %s", fileName)
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolP("include-scaffolding", "s", false, "Include scaffolding files for a new resume (default true)")
}
