// +build js,wasm

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

func main() {
	href := js.Global().Get("location").Get("href")
	u, err := url.Parse(href.String())
	if err != nil {
		log.Fatal(err)
	}
	u.Path = "/logo.png"
	u.RawQuery = fmt.Sprint(time.Now().UnixNano())

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
		canvas.Set("width", image.Get("naturalWidth"))
		canvas.Set("height", image.Get("naturalHeight"))
		ctx.Call("drawImage", image, 0, 0)
		js.Global().Call("setInterval", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			style := canvas.Get("style")
			left := style.Get("left")
			if left == js.Undefined() {
				left = js.ValueOf("0px")
			} else {
				n, _ := strconv.Atoi(strings.TrimRight(left.String(), "px"))
				left = js.ValueOf(fmt.Sprintf("%dpx", n+10))
			}
			style.Set("left", left)
			return nil
		}), js.ValueOf(200))
		return nil
	}))
	image.Set("src", "data:image/png;base64,"+enc)

	canvas.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("window").Call("alert", "Don't click me!")
		return nil
	}))

	select {}
}
