# gobuggery

Lower level wrapper for the basic remote debugger connection to the
github.com/bnagy/rBuggery json stub ( json_debug_server )

This is only used by github.com/bnagy/alpcbuggery at the moment.

## Documentation

Use `import "github.com/bnagy/gobuggery"` in your own Go code.

Get godoc at: http://godoc.org/github.com/bnagy/gobuggery

## Installation

You should follow the [instructions](https://golang.org/doc/install) to
install Go, if you haven't already done so. Then:

```bash
$ go get github.com/bnagy/gobuggery
```

## TODO:

Very few of the rBuggery methods are mapped over to idiomatic Go. Pretty much
all that's working so far is Execute ( to run any windbg command ) and local
kernel connection.

## BUGS

## Contributing

Fork & pullreq

## License

BSD Style, See LICENSE file for details



