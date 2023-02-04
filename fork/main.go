package main

import (
	"syscall"
	"time"
)

func loop() {
	for true {
		println("asdasd")
	}
}

func main() {
	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if id != 0 {
		println("We are the children", id)
		time.Sleep(1e9)
		loop()
	} else {
		println("We are the parent", id)
		syscall.Exit(0)
	}

}
