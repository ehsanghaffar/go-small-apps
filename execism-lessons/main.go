package main

import (
	"fmt"

	lasagna "golanglearn/execism/lasagna"
	"golanglearn/execism/myday"
	"golanglearn/sololearn/saynumber"
)

func main() {
	layers := 3
	actual := 30
	fmt.Println("Lasagna has", layers, "layers and has been in the oven for", actual, "minutes.")
	elapsed := lasagna.ElapsedTime(layers, actual)
	fmt.Println("Total elapsed time:", elapsed, "minutes.")
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)
	nums := [3]int{num1, num2, num3}
	for i := range nums {
		saynumber.Saynumber(nums[i])
	}
	myday.MyDay()
}
