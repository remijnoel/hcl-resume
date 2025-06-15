package template

import (
	_ "embed"
)

//go:embed resume.tex
var ResumeTemplate string

//go:embed resume.cls
var LateXResumeTemplate string

//go:embed hcl_scaffolding.hcl
var HCLScaffolding string