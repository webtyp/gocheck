package main

import "webtyp.com/fmt"

var s = "a,b,c"

func main() { println(fmt.Contains(s, "b"), len(fmt.Split(s, ",")), fmt.Convert(s).ToUpper().String()) }
