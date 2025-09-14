# grep-go

A concurrent file searching tool written in Go, similar to the classic Unix `grep` utility but with parallel processing capabilities.

## Overview

`grep-go` searches for a specified text pattern recursively through files in a directory. Uses Go's concurrency features to perform searches in parallel.

## Installation

Clone the repository and build:

```bash
git clone <repository-url>
cd grep-go
go build -o grepgo ./grepgo
```

## Usage

```bash
./grepgo SEARCH_TERM [SEARCH_DIRECTORY]
```

Arguments:

- `SEARCH_TERM`: The text pattern to search for (required)
- `SEARCH_DIRECTORY`: The directory to search in (if omitted, uses current directory)

## Example

Search for the term "TODO" in the current directory:

```bash
./grepgo TODO .
```

This will display all occurrences of "TODO" in files under the current directory, formatted as:

```
path/to/file[line_number]:matching line content
```

## Implementation Details

The project consists of three main packages:

1. **worklist**: Manages the queue of files to be processed
2. **worker**: Handles file searching operations and result collection
3. **grepgo**: Main package that coordinates directory traversal and worker management

The implementation uses a concurrent approach with multiple worker goroutines to process files in parallel.

## Dependencies

- [github.com/alexflint/go-arg](https://github.com/alexflint/go-arg) - Command-line argument parsing
