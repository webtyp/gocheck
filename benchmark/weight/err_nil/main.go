package main

type e struct{ s string }

func (x *e) Error() string { return x.s }

var got error = &e{"x"}

func main() { println(got != nil) }
