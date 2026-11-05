package main
type e struct{ s string }
func (x *e) Error() string { return x.s }
var EOF = &e{"eof"}
var got error = &e{"x"}
func main() { x, ok := got.(*e); println(ok && x == EOF) }
