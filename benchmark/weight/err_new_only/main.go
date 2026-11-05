package main
import "errors"
var got = errors.New("x")
func main() { println(got.Error()) }
