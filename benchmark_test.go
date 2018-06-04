package main

import (
	"github.com/ryannedolan/go-dark/predef"
	"testing"
)

func BenchmarkFmap(b *testing.B) {
	ints := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		ints[i] = i
	}
	b.ResetTimer()
	_ = predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) }).From(func() []interface{} {
		arr := make([]interface{}, 0)
		go func() {
			for _, e := range ints {
				arr = append(arr, e)
			}
		}()
		return arr
	}).Fmap(func(x int) interface{} { return x + 1 }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) }))
}
