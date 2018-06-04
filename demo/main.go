package main

import (
	"fmt"
	"github.com/ryannedolan/go-dark/predef"
	"strconv"
)

func main() {
	a := []int{1, 2, 3}
	b := predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) }).From(func() chan interface{} {
		ch := make(chan interface{})
		go func() {
			for _, e := range a {
				ch <- e
			}
			close(ch)
		}()
		return ch
	}).
		Fmap(func(a int) interface{} { return a + 1 }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) })).
		Fmap(func(a int) interface{} { return strconv.Itoa(a) + "!" }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x string) interface{})(x.(string)) })).
		Build(func(ch chan interface{}) interface{} {
			arr := []string{}
			for x := range ch {
				arr = append(arr, x.(string))
			}
			return arr
		}).([]string)
	fmt.Println(b)
}
