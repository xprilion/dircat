# dircat

`dircat` is a lightweight command-line tool that recursively concatenates files from a directory into a single text output, while also providing an option to generate a file tree. By default, it outputs the result to the console, but users can redirect the output to a file if needed.

## Features

- Recursively generates a file tree of a given directory.
- Outputs the contents of all non-binary files in a directory.
- Skips hidden files, `.pyc` files, `__pycache__` directories, and binary files.
- Supports clean, simple console output by default.
- Optional file output via a command-line flag.
- Modular with subcommands for tree generation and file content aggregation.

## Installation

### Install via Binary

1. Build the Go binary:
   ```bash
   go build -o dircat
   ```

2. Move the binary to your desired location:
   ```bash
   mv dircat ~/xpritools/
   ```

3. Add `~/xpritools` to your `$PATH`:
   For **Zsh** (default shell on macOS since Catalina):
   ```bash
   echo 'export PATH="$HOME/xpritools:$PATH"' >> ~/.zshrc
   source ~/.zshrc
   ```

   For **Bash**:
   ```bash
   echo 'export PATH="$HOME/xpritools:$PATH"' >> ~/.bash_profile
   source ~/.bash_profile
   ```

4. Now you can run `dircat` from anywhere in your terminal!

### Install via Homebrew (coming soon)

You will soon be able to install `dircat` via Homebrew with:

```bash
brew tap xprilion/tools
brew install dircat
```

## Usage

### Basic Commands

- **File Tree**: Display a tree of the directory structure.
  
  ```bash
  dircat tree /path/to/folder
  ```

- **File Contents**: Display the contents of all non-binary files in the directory.
  
  ```bash
  dircat content /path/to/folder
  ```

- **All**: Display both the file tree and contents of the files.
  
  ```bash
  dircat all /path/to/folder
  ```

### Redirecting Output to a File

To save the output to a file, simply redirect the output using `>>`:

```bash
dircat all /path/to/folder >> output.txt
```

Or use the `-o` flag to specify an output file:

```bash
dircat all /path/to/folder -o output.txt
```

## Skipped Files

The following files are skipped by default:
- Binary files (those that cannot be displayed as text)
- Hidden files (starting with `.`, like `.git`)
- `.pyc` files and `__pycache__` directories

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request. Please ensure your code follows the existing style and passes all tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
