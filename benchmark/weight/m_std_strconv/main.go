package main

import (
	"strconv"
	"webtyp.com/fmt"
)

var n, s = 7, "42"

func main() { i, _ := strconv.Atoi(s); println(fmt.Sprintf("%d", n), strconv.Itoa(n), i) }
