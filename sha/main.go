package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 固定为20个字节
	s1 := sha1.New()
	s1.Write([]byte("1"))
	r1 := s1.Sum(nil)
	fmt.Println(len(r1))

	// 固定为32个字节
	s256 := sha256.New()
	s256.Write([]byte("1"))
	r256 := s256.Sum(nil)
	fmt.Println(len(r256))
}
