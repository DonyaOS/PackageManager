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
```

To build the Go file into an executable for Linux run:

```
go build src/donya.go
```

And generally for Windows:

```
env GOOS=windows GOARCH=amd64 go build src/donya.go
```

To see all the platforms run:

```
go tool dist list -json
```

You can install `revive` with:

```
go get -u github.com/mgechev/revive
```

- [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
- [revive](https://github.com/mgechev/revive) - 🔥 ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for golint.
- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
