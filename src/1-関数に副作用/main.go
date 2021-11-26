/*
  関数の戻り値は、return文が呼ばれ復帰するときに呼び出し元に値渡しされる。
*/
package main

import (
  "fmt"
)

var s = 1

func f() int {
  s++
  return s
}

func main() {
  v := f()
  // vは値渡しなので、sは影響を受けない
  v++
  fmt.Println(s)
  fmt.Println(v)
}
