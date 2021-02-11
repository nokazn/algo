package main

import (
  "fmt"
)

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func GreatestCommonDivisor(a, b int) int {
  n := 1
  for i:= min(a, b); i > 0; i-- {
    if b % i == 0 && a & i == 0 {
      return i
    }
  }
  return n
}

func Euclid(a, b int) int {
  r := a % b
  if r == 0 {
    return b
  }
  return Euclid(b, r)
}

func main() {
  fmt.Println(GreatestCommonDivisor(214, 45))
  fmt.Println(Euclid(214, 45))
}
