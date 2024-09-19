package main

const size = 1

type myStruct struct {
	arr [size]int64 // 8 * size bytes
}

func NewMyStructVal() myStruct {
	var ms myStruct
	for i := 0; i < 1; i++ {
	}
	return ms
}

func NewMyStructPtr() *myStruct {
	var ms myStruct
	for i := 0; i < 1; i++ {
	}
	return &ms
}
