package main

import (
	// "fmt"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func check(data, id string) bool {
	data2 := string(data)
	if !strings.Contains(data2, id) {
		return false
	}
	return true
}

func get_until_char(data string, char rune) string {
	buff := make([]rune, 199)
	var position int
	for pos, v := range data {
		if v == char {
			position = pos
			break
		}
		buff[pos] = v
		position += 1
	}
	return string(buff[:position][1 : len(buff[:position])-1])

}

func get_id(data []byte, id string) string {
	data2 := string(data)
	if !strings.Contains(data2, id) {
		return ""
	}
	arr := strings.Split(data2, id)
	return get_until_char(arr[1], '>')

}

func get_href(data string) string {
	if !check(data, "href=") {
		return ""
	}
	arr := strings.Split(data, "href=")
	return get_until_char(arr[1], ' ')
}

func query_body(address string) []byte {
	res, err := http.Get(address)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return data
}

func GetUrl(address string) string {
	data := query_body(address)
	element := get_id(data, "preload")
	url := get_href(element)
	return address + url
}

func GetImage(address string) []byte {
	return query_body(address)

}

func main() {
	url := GetUrl("https://www.bing.com")
	fmt.Printf("%s", url)
	img := GetImage(url)

	path := os.Getenv("HOME")
	file, err := os.Create(path + "/.wallpaper.jpeg")
	if err != nil {
		panic(err)
	}
	file.Write(img)

	cc := exec.Command("nitrogen", "--set-auto", path+"/.wallpaper.jpeg")
	_, err2 := cc.Output()
	if err2 != nil {
		panic(err2)
	}

}
