package main

import (
	"fmt"

	"github.com/imrandil/go_channel/pipeline" // Replace with your full module path
)

func main() {
	//input
	nums := []int{3, 4, 9, 2, 1}

	//stage 1
	dataChannel := pipeline.SliceToChannel(nums)

	//stage 2
	finalChannel := pipeline.Sq(dataChannel)

	//output it now in the main channel
	for n := range finalChannel {
		fmt.Println(n)
	}

}
