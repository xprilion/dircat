package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// outputFileTree generates the file tree structure
func outputFileTree(folderPath string, out io.Writer) error {
	_, err := fmt.Fprintln(out, "\n\n>> File Tree:")
	if err != nil {
		return err
	}

	return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == folderPath {
			return nil
		}

		// Skip __pycache__, .pyc files, and hidden files
		if strings.Contains(path, "__pycache__") || strings.HasPrefix(info.Name(), ".") || strings.HasSuffix(info.Name(), ".pyc") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Print folder or file with proper indentation
		relativePath, err := filepath.Rel(folderPath, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			_, err = fmt.Fprintf(out, "%s\n", relativePath)
		} else {
			_, err = fmt.Fprintf(out, "    ├── %s\n", relativePath)
		}

		return err
	})
}
