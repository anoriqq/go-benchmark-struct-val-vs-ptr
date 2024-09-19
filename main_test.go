package main

import (
	"testing"
)

func BenchmarkVal_ShortLife(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		v := NewMyStructVal()
		sum += v.arr[0]
	}
}

func BenchmarkPtr_ShortLife(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		v := NewMyStructPtr()
		sum += v.arr[0]
	}
}

func BenchmarkVal_LongLife(b *testing.B) {
	var list []myStruct
	for i := 0; i < b.N; i++ {
		v := NewMyStructVal()
		list = append(list, v)
	}
}

func BenchmarkPtr_LongLife(b *testing.B) {
	var list []*myStruct
	for i := 0; i < b.N; i++ {
		v := NewMyStructPtr()
		list = append(list, v)
	}
}

func BenchmarkVal_LongLife_NoGrow(b *testing.B) {
	list := make([]myStruct, 0, b.N)
	for i := 0; i < b.N; i++ {
		v := NewMyStructVal()
		list = append(list, v)
	}
}

func BenchmarkPtr_LongLife_NoGrow(b *testing.B) {
	list := make([]*myStruct, 0, b.N)
	for i := 0; i < b.N; i++ {
		v := NewMyStructPtr()
		list = append(list, v)
	}
}
