package main

import "webtyp.com/fmt"

var n, s = 7, "a,b,c"

func main() {
	println(fmt.Sprintf("%d %s", n, s), fmt.Contains(s, "b"), len(fmt.Split(s, ",")), fmt.Convert(s).ToUpper().String())
}
