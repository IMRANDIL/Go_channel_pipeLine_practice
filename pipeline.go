package main

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int) //making a non buffered channel to make it synchronous.

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(numChannel <-chan int) <-chan int {
	out := make(chan int) //making a non buffered channel to make it synchronous

	go func() {
		for n := range numChannel {
			out <- n * n
		}
		close(out)
	}()

	return out
}
