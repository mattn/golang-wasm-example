.PHONY: all clean serve

all: main.wasm serve

%.wasm: %.go
	GOOS=js GOARCH=wasm go build -o "$@" "$<"

serve:
	xdg-open 'http://localhost:5000'
	serve || (go get -v github.com/mattn/serve && serve)

clean:
	rm -f *.wasm
