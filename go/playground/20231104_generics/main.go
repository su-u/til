package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))
	fmt.Println(f([]MyInt2{1, 2, 3, 4}))
	// fmt.Println(f([]MyInt3{1, 2, 3, 4})) // インターフェースを実装していないため負荷
}

func f[T Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}

	return result
}

type Stringer interface {
	String() string
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

type MyInt2 int

func (i MyInt2) String() string {
	return strconv.Itoa(int(i + 1))
}

type MyInt3 int

func (i MyInt3) String2() string {
	return strconv.Itoa(int(i + 1))
}
