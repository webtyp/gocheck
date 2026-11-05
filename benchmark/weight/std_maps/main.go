package main
import "maps"
var m = map[string]int{"a": 1}
func main() { c := maps.Clone(m); println(len(c)) }
