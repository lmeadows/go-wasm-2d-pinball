package main

import (
	"fmt"
	"syscall/js"
)

var runForever chan struct{}

func main() {
	js.Global().Get("document").Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		key, code := event.Get("key").String(), event.Get("code").String()
		println(fmt.Printf("Pressed %v with code %v", key, code))
		return nil
	}))

	<-runForever
}
