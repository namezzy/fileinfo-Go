# fileinfo: A Go CLI Tool for File and Directory Information

This README provides a detailed guide to developing a command-line tool using Go.  Go is well-suited for CLI tools due to its ability to compile to a single executable, fast startup time, and rich standard and third-party library support.  This guide walks you through creating a simple yet practical tool.

## Functionality

We'll build a command-line tool named `fileinfo` that can:

* Display basic information about a specified file or directory.
* Recursively count the number of files and total size within a directory.
* Support filtering files by specific types.

## Prerequisites

* **Go:** Ensure you have Go installed.  Download and install it from the [official Go website](https://go.dev/dl/) if needed.  Verify your installation with:

```bash
go version
```

## Project Setup

1. **Create Project Directory:**

```bash
mkdir fileinfo
cd fileinfo
```

2. **Initialize Go Module:** Replace `yourusername` with your actual GitHub username.

```bash
go mod init github.com/yourusername/fileinfo
```

3. **Install Cobra:** We'll use the Cobra library for command-line argument parsing.

```bash
go get -u github.com/spf13/cobra@latest
```

4. **Install color library:**

```bash
go get -u github.com/fatih/color
```

5. **Build and Install:**

```bash
# linux platform
go build -o fileinfo

# Windows platform
go build -o fileinfo.exe
```

This creates an executable `fileinfo` in the current directory. To install it system-wide (requires appropriate permissions):

```bash
go install
```

This installs the executable to your `$GOPATH/bin` directory (or equivalent).

## Usage

```bash
# Get information about the current directory
./fileinfo

# Get information about a specific file
./fileinfo main.go

# Recursively count files in a directory
./fileinfo -r /path/to/directory

# Count only files of a specific type
./fileinfo -r -t .go /path/to/directory
```

## Extending Functionality

You can extend this tool by:

* Adding more commands (e.g., `fileinfo search` to search files).
* Adding more filtering options (e.g., by size, modification time).
* Implementing sorting functionality (e.g., by size or time).
* Improving error handling and user feedback.


This example demonstrates the basic workflow for developing CLI tools in Go.  Leverage Go's strengths—single executable compilation, cross-platform support, rich libraries, and performance—to build powerful and efficient command-line applications.
