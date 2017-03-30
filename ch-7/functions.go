package main

import (
  "fmt"
)

// write a function for sum
func sum(args ...int) int {
  total :=0
  for _, v:= range args {
    total += v
  }
  return total
}

// write a halving functions
func half(x int) (int, bool) {
  return x / 2 , x % 2 == 0
}

// variadic max
func vmax(args ...int) int {
  max :=0
  for _, v:= range args {
    if v > max {
      max = v
    }
  }
  return max
}

// odd generator
func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

// recursive fibonacci
func fibonacci(x int) int {
  if x == 0 {
    return 0
  } else if x == 1 {
    return 1
  } else {
    return fibonacci(x-1) + fibonacci(x-2)
  }
}

// main test runnner
func main() {
  // test sum
  fmt.Println(sum(1,2,3))

  // test half
  fmt.Println(half(4))

  // test vmax
  fmt.Println(vmax(1,2,3,10,9))

  // odd generator
  nextOdd := makeOddGenerator()
  for i:=0; i<10; i++ {
    fmt.Print(nextOdd(), " ")
  }
  fmt.Println()

  // fibonacci
  fmt.Println(fibonacci(10))

  // defer recover panic pattern
  defer func() {
    str := recover()
    fmt.Println("CAUGHT: ", str)
  }()
  panic("PANIC")
}
