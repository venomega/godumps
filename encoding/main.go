package main

import "encoding/json"

type asd struct {
	Id       uint64
	Name     string
	Category string
}

func main() {
	foo := asd{1, "Guillermo Plasencia Semanat", "Alumni"}
	data, err := json.Marshal(foo)
	if err != nil {
		panic(err)
	}
	println("the json is", string(data))
	json.Unmarshal(data, &foo)

	println("the struct is", foo.Id)
}
