# goserve

Simple command-line static http-server.

## Installation

You need `go` installed and `GOBIN` in your `PATH`. Once that is done, run the
command:

```shell
$ go get -u github.com/thiamsantos/goserve
```

## Usage

```
Usage
  $ goserve [options]

Options
  -path string
        Path to serve (default ".")
  -port int
        Port to use (default 8080)

Examples
  $ goserve
  $ goserve -port 3000
  $ goserve -path /tmp/static
  $ goserve -port 8888 -path /tmp/static
```
