package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func write(bridge js.Value, format string, args ...any) {
	bridge.Call("write", fmt.Sprintf(format, args...))
}

func writePrompt(bridge js.Value) {
	write(bridge, "> ")
}

func writeSize(bridge js.Value, cols, rows int) {
	write(bridge, "terminal size: %d rows, %d columns\r\n", rows, cols)
}

func main() {
	bridge := js.Global().Get("goPty")
	if bridge.IsUndefined() || bridge.IsNull() {
		panic("goPty is not defined")
	}

	inputCh := make(chan string, 16)
	resizeCh := make(chan [2]int, 8)

	onData := js.FuncOf(func(_ js.Value, args []js.Value) any {
		inputCh <- args[0].String()
		return nil
	})

	onResize := js.FuncOf(func(_ js.Value, args []js.Value) any {
		resizeCh <- [2]int{args[0].Int(), args[1].Int()}
		return nil
	})

	dataSubscription := bridge.Call("onData", onData)
	resizeSubscription := bridge.Call("onResize", onResize)
	defer dataSubscription.Call("dispose")
	defer resizeSubscription.Call("dispose")
	defer onData.Release()
	defer onResize.Release()

	size := bridge.Call("getSize")
	write(bridge, "Go/Wasm + xterm-pty\r\n")
	write(bridge, "Type anything and press Enter. Type `size` or `exit`.\r\n")
	writeSize(bridge, size.Get("cols").Int(), size.Get("rows").Int())
	writePrompt(bridge)

	var pending string

	for {
		select {
		case chunk := <-inputCh:
			pending += chunk
			for {
				index := strings.IndexAny(pending, "\r\n")
				if index < 0 {
					break
				}

				line := strings.TrimSpace(pending[:index])
				pending = strings.TrimLeft(pending[index+1:], "\r\n")

				switch line {
				case "":
				case "size":
					size := bridge.Call("getSize")
					writeSize(bridge, size.Get("cols").Int(), size.Get("rows").Int())
				case "exit":
					write(bridge, "bye!\r\n")
					return
				default:
					write(bridge, "echo: %s\r\n", line)
				}

				writePrompt(bridge)
			}
		case size := <-resizeCh:
			write(bridge, "\r\n")
			writeSize(bridge, size[0], size[1])
			writePrompt(bridge)
		}
	}
}
