package main

import (
	"fmt"
	"github.com/ryannedolan/go-dark/predef"
	"strconv"
)

func main() {
	a := []int{1, 2, 3}
	b := predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) }).From(func() []interface{} {
		arr := make([]interface{}, 0)
		go func() {
			for _, e := range a {
				arr = append(arr, e)
			}
		}()
		return arr
	}).
		Fmap(func(a int) interface{} { return a + 1 }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) })).
		Fmap(func(a int) interface{} { return strconv.Itoa(a) + "!" }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x string) interface{})(x.(string)) })).
		Build(func(elems []interface{}) interface{} {
			arr := []string{}
			for _, x := range elems {
				arr = append(arr, x.(string))
			}
			return arr
		}).([]string)
	fmt.Println(b)
}
