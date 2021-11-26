/*
  もし関数fが、可変数引数pに何も渡さない状態で呼び出された場合、pに渡される値はnilとなります。
*/

package main

import "fmt"

func main() {
  f()
}

func f(p ...string) {
  // pにはnilが格納される
  fmt.Println(p == nil)
}