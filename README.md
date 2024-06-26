# Go Command Reference

This README provides an overview of the principal Go commands you'll need when working on Go projects.

## Table of Contents

- [Installing Go](#installing-go)
- [Setting Up a Go Project](#setting-up-a-go-project)
- [Running Go Code](#running-go-code)
- [Managing Modules](#managing-modules)
- [Testing](#testing)
- [Building Executables](#building-executables)
- [Commonly Used Commands](#commonly-used-commands)
- [Help and Documentation](#help-and-documentation)

## Installing Go

To install Go, download the latest version from the [official Go website](https://golang.org/dl/) and follow the installation instructions for your operating system.

## Setting Up a Go Project

### Initializing a New Module

To create a new Go module, navigate to your project directory and run:

```sh
go mod init <module-name>
```

Example:

```sh
go mod init github.com/yourusername/yourproject
```

## Running Go Code

### Running a Go File

To run a Go file, use the `go run` command followed by the file name:

```sh
go run main.go
```

### Running a Go Package

To run a Go package, use:

```sh
go run .
```

This command runs the `main` package in the current directory.

## Managing Modules

### Adding a Dependency

To add a dependency to your project, use:

```sh
go get <module-path>
```

Example:

```sh
go get github.com/gorilla/mux
```

### Tidying Up Dependencies

To clean up your `go.mod` file and remove unnecessary dependencies, use:

```sh
go mod tidy
```

### Downloading All Dependencies

To download all dependencies specified in your `go.mod` file, use:

```sh
go mod download
```

## Testing

### Running Tests

To run tests in your project, use:

```sh
go test ./...
```

This command runs all tests in the current module.

### Running Tests with Verbose Output

To run tests with verbose output, use:

```sh
go test -v ./...
```

### Running a Specific Test

To run a specific test function, use:

```sh
go test -run <test-name>
```

Example:

```sh
go test -run TestFunctionName
```

## Building Executables

### Building for the Current System

To build an executable for the current system, use:

```sh
go build
```

This command creates an executable in the current directory.

### Building for a Different Operating System

To build an executable for a different operating system, set the `GOOS` and `GOARCH` environment variables:

```sh
GOOS=linux GOARCH=amd64 go build
```

## Commonly Used Commands

### Formatting Code

To format your code according to Go standards, use:

```sh
go fmt ./...
```

### Vetting Code

To examine your code for potential issues, use:

```sh
go vet ./...
```

### Generating Documentation

To generate HTML documentation for your project, use:

```sh
godoc -http=:6060
```

Then navigate to `http://localhost:6060` in your web browser.

## Help and Documentation

### Getting Help

To get help with Go commands, use:

```sh
go help
```

### Getting Help for a Specific Command

To get help for a specific command, use:

```sh
go help <command>
```

Example:

```sh
go help build
```

### Viewing Documentation

To view the documentation for a package, use:

```sh
go doc <package>
```

Example:

```sh
go doc fmt
```

---

This README provides a quick reference to the most commonly used Go commands. For more detailed information, visit the [official Go documentation](https://golang.org/doc/).
