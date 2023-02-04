package main

func the_func(b string, a chan int) {
	a <- len(b)
}

func main() {
	channel := make(chan int)
	go the_func("Theda", channel)
	go the_func("new_string", channel)

	a, b := <-channel, <-channel
	println(a, b)
}
