package main
import "errors"
var EOF = errors.New("eof")
var got = errors.New("x")
func main() { println(errors.Is(got, EOF)) }
