package main
import "strings"
var s = "a,b,c"
func main() { println(strings.Contains(s, "b"), len(strings.Split(s, ",")), strings.ToUpper(s)) }
