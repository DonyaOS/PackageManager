# Contributing

To print your Go environment variables run:

```
go env
```

From the root of this project install the Golang libraries with:

```
go get -u ./...
```

Example run:

```
go run src/donya.go
Usage:
  donya [command]

Available Commands:
  help        Help about any command
  install     Installing package(s) in DonyaOS

Flags:
  -h, --help   help for donya

Use "donya [command] --help" for more information about a command.
```

- [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
- [revive](https://github.com/mgechev/revive) - 🔥 ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for golint.
- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
