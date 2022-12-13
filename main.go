package golanglearn

import (
	"fmt"

	lasagna "golanglearn/execism/lasagna"
)

func main() {
	layers := 3
	actual := 30
	fmt.Println("Lasagna has", layers, "layers and has been in the oven for", actual, "minutes.")
	elapsed := lasagna.ElapsedTime(layers, actual)
	fmt.Println("Total elapsed time:", elapsed, "minutes.")
}
