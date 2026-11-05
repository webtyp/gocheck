// Baseline for mall_*: webtyp/fmt as a real app uses it, without string helpers.
package main

import "webtyp.com/fmt"

var n, s = 7, "a,b,c"

func main() {
	println(fmt.Sprintf("%d %s", n, s), fmt.Err("x", s).Error())
}
