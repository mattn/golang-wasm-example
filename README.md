# golang-wasm-example

Example app using Go's wasm support.

## Build


    GOOS=js GOARCH=wasm go build -o main.wasm main.go


## Run

    go get github.com/mattn/serve
    serve

## Build and run

    make

## License

MIT

## Author

Yasuhrio Matsumoto (a.k.a. mattn)
