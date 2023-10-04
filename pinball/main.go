package main

import (
	"fmt"
	"math"
	"syscall/js"
)

var runForever chan struct{}

func main() {
	canvas, _ := Init("gamecanvas", js.Global().Get("innerWidth").Float()*0.4, js.Global().Get("innerHeight").Float()*0.9)

	js.Global().Get("document").Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		key, code := event.Get("key").String(), event.Get("code").String()
		println(fmt.Printf("Pressed %v with code %v", key, code))
		return nil
	}))

	ctx := canvas.domNode.Call("getContext", "2d")

	drawBall(ctx)

	<-runForever
}

func drawBall(ctx js.Value) {
	// Need this code to get this to work:
	// https://github.com/markfarnan/go-canvas/blob/6971ccd00770ac5ee39846076d21935a933364ee/canvas/canvas2d.go#L97C2-L97C66
	// https://github.com/markfarnan/go-canvas/blob/6971ccd00770ac5ee39846076d21935a933364ee/canvas/canvas2d.go#L181-L183
	// Must figure out copybuff AND create a canvas abstraction
	//image := image.NewRGBA(image.Rect(0, 0, 100, 100))
	//c.copybuff = js.Global().Get("Uint8Array").New(len(c.image.Pix))
	imageData := ctx.Call("createImageData", 10, 10)
	ctx.Call("beginPath")
	ctx.Call("arc", 200, 175, 30, 0, 2*math.Pi)
	ctx.Call("stroke")
	ctx.Call("putImageData", imageData, 300, 100)
}

func log(args ...interface{}) {
	js.Global().Get("console").Call("log", args...)
}
