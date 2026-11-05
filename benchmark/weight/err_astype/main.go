package main
import "errors"
type NotFound struct{ s string }
func (x *NotFound) Error() string { return x.s }
var got error = &NotFound{"x"}
func main() { _, ok := errors.AsType[*NotFound](got); println(ok) }
