package internal

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/remijnoel/hcl-resume/resume"
)

func escapeLaTeX(s string) string {
	return strings.ReplaceAll(s, "&", `\&`)
}

// RenderToTeX renders the Resume struct to a LaTeX source file.
//   tmplPath: path to resume.tex.tmpl
//   outPath: path to output .tex file
func RenderToTeX(resume *resume.Resume, tmplContent string) ([]byte, error) {
	// Parse the template
	tmpl, err := template.
	New("resume").
	Delims("<<", ">>").
	Funcs(template.FuncMap{"escape": escapeLaTeX}).
	Parse(tmplContent)

	if err != nil {
		return nil, fmt.Errorf("cannot parse template: %w", err)
	}

	// Render the template with the Resume struct
	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, resume); err != nil {
		return nil, fmt.Errorf("template execution failed: %w", err)
	}

	return rendered.Bytes(), nil
}
