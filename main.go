package main

import "fmt"

func main() {
	//input
	nums := []int{3, 4, 9, 2, 1}

	//stage 1
	dataChannel := sliceToChannel(nums)

	//stage 2
	finalChannel := sq(dataChannel)

	//output it now in the main channel
	for n := range finalChannel {
		fmt.Println(n)
	}

}
