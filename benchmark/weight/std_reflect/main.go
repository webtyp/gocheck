package main

import "reflect"

var v any = 7

func main() { println(reflect.TypeOf(v).Kind() == reflect.Int) }
