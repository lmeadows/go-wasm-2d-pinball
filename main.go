package main

import (
	"syscall/js"

	"github.com/markfarnan/go-canvas/canvas"
)

var runForever chan struct{}

var cvs *canvas.Canvas2d
var width float64
var height float64

func main() {
	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(int(js.Global().Get("innerWidth").Float()*0.5), int(js.Global().Get("innerWidth").Float()*0.8))

	<-runForever
}
