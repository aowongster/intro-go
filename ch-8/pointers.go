package main

import (
  "fmt"
)

func square (x *float64) {
  *x = *x * *x
}

func swap(x, y *int) {
  *x, *y = *y, *x
}

func main() {
  x := 1.5
  square(&x)
  fmt.Println(x)

  y, z := 2, 4
  fmt.Println(y, z)
  swap(&y, &z)
  fmt.Println(y, z)

}
