package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("greet", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return "Hello, World!"
		}
		return fmt.Sprintf("Hello, %s!", args[0].String())
	}))

	select {}
}
