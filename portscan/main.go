package main

import (
	"fmt"
	"net"
	"time"
)

var count int = 0

func monitor() {
	for true {
		print("\b\b\b\b\b\b\b\b")
		print("        ")
		print("\b\b\b\b\b\b\b\b")
		print(count)
		time.Sleep(1.0e9)
	}
}

func exec(a string) {
	_, err := net.Dial("tcp", a)
	if err == nil {
		println(a, "open")
	}
	count -= 1
}

func main() {
	for i := 0; i < 1081; i++ {
		count += 1
		string := fmt.Sprint("scanme.nmap.org:", i)
		go exec(string)
	}
	go monitor()
	for true {
		if count == 0 {
			break
		}
	}

}
