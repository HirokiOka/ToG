package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("seed", rand.Seed)
  fmt.Println("My favorite number is", rand.Intn(10))
}
