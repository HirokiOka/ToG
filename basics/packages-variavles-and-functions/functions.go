package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func lazyAdd(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(lazyAdd(42, 13))
}
