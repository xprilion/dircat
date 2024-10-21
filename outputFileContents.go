package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// isBinaryFile checks whether a file is likely a binary file by inspecting the first few bytes
func isBinaryFile(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// Check the first 512 bytes of the file
	buffer := make([]byte, 512)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		return false, err
	}

	for i := 0; i < n; i++ {
		// If any byte is not printable and not a control character (like newline), assume binary
		if !unicode.IsPrint(rune(buffer[i])) && !unicode.IsControl(rune(buffer[i])) {
			return true, nil
		}
	}
	return false, nil
}

// outputFileContents appends the contents of the files with relative paths, skipping binary files
func outputFileContents(folderPath string, out io.Writer) error {
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

		// Only append file contents (skip directories)
		if !info.IsDir() {
			relativePath, err := filepath.Rel(folderPath, path)
			if err != nil {
				return err
			}

			// Check if the file is binary
			isBinary, err := isBinaryFile(path)
			if err != nil {
				return err
			}
			if isBinary {
				return nil // Skip binary files silently
			}

			// Append file header and content
			_, err = fmt.Fprintf(out, "\n---\n\n>> Filename: %s\n\n", relativePath)
			if err != nil {
				return err
			}

			// Append file content
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(out) // Add a blank line between files
			return err
		}

		return nil
	})
}
