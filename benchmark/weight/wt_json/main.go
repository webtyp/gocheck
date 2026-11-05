package main
import ("webtyp.com/json"; "webtyp.com/model")
type U struct{ Name string; Age int64 }
func (u *U) EncodeFields(w model.FieldWriter) { w.String("name", u.Name); w.Int("age", u.Age) }
func (u *U) IsNil() bool { return u == nil }
var u = &U{"a", 3}
func main() { var out string; _ = json.Encode(u, &out); println(out) }
