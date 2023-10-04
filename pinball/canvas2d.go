package main

import "syscall/js"

type Canvas2d struct {
	domNode js.Value
}

func Init(domNodeId string, width float64, height float64) (*Canvas2d, error) {
	var canvas Canvas2d
	canvas.domNode = js.Global().Get("document").Call("getElementById", domNodeId)
	canvas.domNode.Set("height", js.Global().Get("innerHeight").Float()*0.9)
	canvas.domNode.Set("width", js.Global().Get("innerWidth").Float()*0.4)

	return &canvas, nil
}
