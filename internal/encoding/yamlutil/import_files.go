package yamlutil

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"sort"
)

func GetYmlImportFiles(e embed.FS, f string) ([]string, error) {
	var files []string
	for _, pat := range []string{
		path.Join(f, "*.yml"),
		path.Join(f, "*.yaml"),
	} {
		matches, err := fs.Glob(e, pat)
		if err != nil {
			return nil, fmt.Errorf("glob dir: %s, pattern: %s, error: %w", f, pat, err)
		}
		files = append(files, matches...)
	}
	sort.Strings(files)
	return files, nil
}
