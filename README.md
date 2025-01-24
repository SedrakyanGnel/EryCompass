# QueryCompass

A command-line tool written in Go that converts a list of URLs into structured JSON.

## Installation

```sh
GO111MODULE=on go install github.com/CyberCompass/QueryCompass/querycompass@latest
```

Ensure your Go `bin` directory is in your system `PATH`:

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage

```sh
querycompass input.txt output.json
```

If you do not provide an output file, the JSON will be printed to the console.
