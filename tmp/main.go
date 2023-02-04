package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type new_int int

func (value new_int) String() string {
	return fmt.Sprintf("%d", value)
}

func main() {
	println(time.Now().Second())
	//time.Sleep(1.0e9)
	println(time.Now().Second())

	var asd new_int = 12

	println(asd.String())

	dict := map[string]int{}

	dict["00"] = 12
	dict["10"] = 12
	println(dict["00"])
	v, _ := json.Marshal(dict)
	fmt.Printf("%v", string(v))

}
