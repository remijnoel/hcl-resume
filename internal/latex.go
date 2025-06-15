package internal

import "os/exec"

func RunPDFLatex(path string, outputDir string) error {
	cmd := exec.Command("pdflatex", "-output-directory", outputDir, path)
	return cmd.Run()
}

