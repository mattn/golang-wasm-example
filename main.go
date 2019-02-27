// +build js,wasm

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"syscall/js"
)

func main() {
	href := js.Global().Get("location").Get("href")
	u, err := url.Parse(href.String())
	if err != nil {
		log.Fatal(err)
	}
	u.Path = "/logo.png"

	log.Println("loading image file: " + u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	enc := base64.StdEncoding.EncodeToString(b)

	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")
	image := js.Global().Call("eval", "new Image()")
	image.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx.Call("drawImage", image, 0, 0)
		return nil
	}))
	image.Set("src", "data:image/png;base64,"+enc)

	canvas.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("window").Call("alert", "Don't click me!")
		return nil
	}))

	select {}
}
