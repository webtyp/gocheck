package main

import "encoding/json"

type U struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var u = U{"a", 3}

func main() { b, _ := json.Marshal(u); println(string(b)) }
