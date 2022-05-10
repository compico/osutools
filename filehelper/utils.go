package filehelper

import (
	"fmt"
	"os"
	"path/filepath"
)

func lsdir(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open path: %w", err)
	}
	defer f.Close()
	finfo, err := f.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("cannot read directory: %w", err)
	}
	var files []string
	for i := 0; i < len(finfo); i++ {
		files = append(files, filepath.Join(path, finfo[i].Name()))
	}
	return files, nil
}
