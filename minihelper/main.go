package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func ERR(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, tmp := os.ReadFile(os.Args[1])
	ERR(tmp)
	log, err := os.Create("/tmp/minihelper.log")
	ERR(err)
	var url string
	var remote_host string
	var dash string
	var method string
	var local_host string
	var localport string

	for {
		fmt.Scanf("%s %s %s %s %s %s", &url, &remote_host, &dash, &method, &local_host, &localport)
		new_url := strings.Split(url, "//")
		new_url2 := strings.ReplaceAll(new_url[1], "/", "")
		log.Write([]byte(new_url2))

		if bytes.Contains(data, []byte(new_url2)) {
			fmt.Printf("OK rewrite-url=\"%s\"\n", os.Args[2])
			log.Write([]byte("OK"))
		} else {
			log.Write([]byte("ERR\n"))
			fmt.Printf("ERR\n")
		}
	}
}
