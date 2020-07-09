package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 字符串中确实不会出现0
	a := [11]byte{}
	copy(a[:], "helloworld")
	fmt.Println(string(a[:]))
	for _, i := range a {
		fmt.Println(string(i), i)
	}
	// 数组len和cap是一样的
	fmt.Println(len(a))
	fmt.Println(cap(a))
	// 可以计算指定数据类型占用内存大小
	fmt.Println(unsafe.Sizeof(*(*Hello)(unsafe.Pointer(nil))))
}

type Hello struct {
	hello [10]byte
	world [10]byte
}
