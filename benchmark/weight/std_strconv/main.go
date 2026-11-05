package main

import "strconv"

var n, s = 7, "42"

func main() { i, _ := strconv.Atoi(s); println(strconv.Itoa(n), i) }
