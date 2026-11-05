package main
import "webtyp.com/fmt"
var n = 7
var a, b error = fmt.Err("a"), fmt.Err("b")
func main() { println(fmt.Sprintf("%d", n), a == b) }
