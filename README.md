# golang-wasm-example

Example app using Go's wasm

## Build

```
$ GOOS=js GOARCH=wasm go build -o test.wasm test.go
```

## Run

```
$ go get github.com/mattn/serve
$ serve
```

## License

MIT

## Author

Yasuhrio Matsumoto (a.k.a. mattn)
