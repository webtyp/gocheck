package main

import (
	"strings"
	"webtyp.com/fmt"
)

var n, s = 7, "a,b,c"

func main() {
	println(fmt.Sprintf("%d %s", n, s), strings.Contains(s, "b"), len(strings.Split(s, ",")), strings.ToUpper(s))
}
