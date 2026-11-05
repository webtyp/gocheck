package main
import "webtyp.com/fmt"
var n, s = 7, "42"
func main() { i, _ := fmt.Convert(s).Int(); println(fmt.Convert(n).String(), i) }
