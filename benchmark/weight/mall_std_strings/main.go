// Every string helper webtyp/fmt duplicates, used together (mall_std_strings).
package main

import (
	"strings"
	"webtyp.com/fmt"
)

var n, s = 7, "a,b,c"

func main() {
	println(fmt.Sprintf("%d %s", n, s), fmt.Err("x", s).Error())
	println(strings.Contains(s, "b"), strings.HasPrefix(s, "a"), strings.HasSuffix(s, "c"), strings.Index(s, "b"), strings.LastIndex(s, ","), len(strings.Split(s, ",")), strings.TrimSpace(" "+s), strings.TrimPrefix(s, "a"), strings.TrimSuffix(s, "c"), strings.ToLower(s), strings.ToUpper(s), strings.Repeat(s, 2), strings.Count(s, ","), strings.ReplaceAll(s, ",", ";"), strings.Join([]string{s, s}, ","))
}
