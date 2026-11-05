// Every string helper webtyp/fmt duplicates, used together (mall_wt_strings).
package main

import (
	"webtyp.com/fmt"
)

var n, s = 7, "a,b,c"

func main() {
	println(fmt.Sprintf("%d %s", n, s), fmt.Err("x", s).Error())
	println(fmt.Contains(s, "b"), fmt.HasPrefix(s, "a"), fmt.HasSuffix(s, "c"), fmt.Index(s, "b"), fmt.LastIndex(s, ","), len(fmt.Split(s, ",")), fmt.TrimSpace(" "+s), fmt.TrimPrefix(s, "a"), fmt.TrimSuffix(s, "c"), fmt.ToLower(s), fmt.ToUpper(s), fmt.Repeat(s, 2), fmt.Count(s, ","), fmt.ReplaceAll(s, ",", ";"), fmt.JoinSlice([]string{s, s}, ","))
}
