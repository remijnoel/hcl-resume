package resume

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func Load(filename string, resume *Resume) error{
	err := hclsimple.DecodeFile(filename, nil, resume)
    if err != nil {
        return fmt.Errorf("failed to decode resume file %s: %w", filename, err)
    }
    return nil
}

// LoadFromDir loads all .hcl files in a directory, merges them, and parses into Resume.
func LoadFromDir(dir string) (*Resume, error) {
	var allHCL []string

	// Find all .hcl files (ignoring non-files)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".hcl") {
			b, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("reading %s: %w", path, err)
			}
			allHCL = append(allHCL, string(b))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(allHCL) == 0 {
		return nil, fmt.Errorf("no .hcl files found in %s", dir)
	}

	// Combine all HCL content into one string
	combined := strings.Join(allHCL, "\n")

	var resume Resume
	err = hclsimple.Decode("resume-all.hcl", []byte(combined), nil, &resume)
	if err != nil {
		return nil, fmt.Errorf("HCL parse error: %w", err)
	}
	return &resume, nil
}