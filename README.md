# golang-wasm-example

![wasm](https://raw.githubusercontent.com/mattn/golang-wasm-example/master/screenshot.gif)

Example app using Go's wasm support.

## Build


```sh
GOOS=js GOARCH=wasm go generate
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```


## Run

```sh
go get github.com/mattn/serve
serve
```

## Build and run

```sh
make
```

## License

MIT

## Author

Yasuhrio Matsumoto (a.k.a. mattn)
