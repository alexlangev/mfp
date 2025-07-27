package main

import (
	"fmt"

	"github.com/alexlangev/mfp/internal/episodes"
)

func main() {
	fmt.Println("Hello there!")
	fmt.Println()
	x, _ := episodes.GetEpisodes()
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
}
