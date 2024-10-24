package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const version = "dircat v0.1.2"

func main() {
	// Define flags
	var outputFile string
	var showVersion bool

	flag.StringVar(&outputFile, "o", "", "Specify output file (optional)")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&showVersion, "v", false, "Show version information (shorthand)")
	flag.Parse()

	// Handle version flag
	if showVersion {
		fmt.Println(version)
		return
	}

	if flag.NArg() < 2 {
		fmt.Println("Usage: dircat <subcommand> <folder_path> [-o output_file]")
		fmt.Println("Subcommands:")
		fmt.Println("  tree    - Output file tree structure")
		fmt.Println("  content - Output file contents")
		fmt.Println("  all     - Output both file tree and contents")
		return
	}

	subcommand := flag.Arg(0)
	folderPath := flag.Arg(1)

	var out io.Writer = os.Stdout
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer file.Close()
		out = file
	}

	switch subcommand {
	case "tree":
		err := outputFileTree(folderPath, out)
		if err != nil {
			fmt.Printf("Error generating file tree: %v\n", err)
		}
	case "content":
		err := outputFileContents(folderPath, out)
		if err != nil {
			fmt.Printf("Error appending file contents: %v\n", err)
		}
	case "all":
		err := outputFileTree(folderPath, out)
		if err != nil {
			fmt.Printf("Error generating file tree: %v\n", err)
		}
		err = outputFileContents(folderPath, out)
		if err != nil {
			fmt.Printf("Error appending file contents: %v\n", err)
		}
	default:
		fmt.Println("Invalid subcommand. Use 'tree', 'content', or 'all'.")
	}
}
