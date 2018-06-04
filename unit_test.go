package main

import (
	"github.com/ryannedolan/go-dark/predef"
	"testing"
)

func TestFmap(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	res1 := predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) }).From(func() chan interface{} {
		ch := make(chan interface{})
		go func() {
			for _, e := range ints {
				ch <- e
			}
			close(ch)
		}()
		return ch
	}).Fmap(func(x int) interface{} { return x + 1 }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) })).Build(func(ch chan interface{}) interface{} {
		arr := []int{}
		for x := range ch {
			arr = append(arr, x.(int))
		}
		return arr
	}).([]int)
	t.Log(res1)
	if len(res1) != 5 {
		t.Fatalf("len wrong")
	}
	if res1[0] != 2 || res1[1] != 3 || res1[2] != 4 || res1[3] != 5 || res1[4] != 6 {
		t.Fatalf("Build(func(ch chan interface{}) interface{} { arr := []{}; for x := range ch { arr = append(arr, x.()) }; return arr } ).([]) content wrong")
	}
	res2 := predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) }).From(func() chan interface{} {
		ch := make(chan interface{})
		go func() {
			for _, e := range res1 {
				ch <- e
			}
			close(ch)
		}()
		return ch
	}).Fmap(func(x int) interface{} { return x - 1 }, predef.BuildIterator(func(x interface{}, f interface{}) interface{} { return f.(func(x int) interface{})(x.(int)) })).Build(func(ch chan interface{}) interface{} {
		arr := []int{}
		for x := range ch {
			arr = append(arr, x.(int))
		}
		return arr
	}).([]int)
	t.Log(res2)
	if len(res2) != 5 {
		t.Fatalf("len wrong")
	}
	if res2[0] != 1 || res2[1] != 2 || res2[2] != 3 || res2[3] != 4 || res2[4] != 5 {
		t.Fatalf("Build(func(ch chan interface{}) interface{} { arr := []{}; for x := range ch { arr = append(arr, x.()) }; return arr } ).([]) content wrong")
	}
}
