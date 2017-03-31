package main

import (
  "fmt"
  "time"
)

func sleep() {
  time.Sleep(time.Second * 1)
}

func ping(c chan string) {
  for i:=0; ; i++ {
    c <- "."
  }
}

func sleepAfter(c chan string) {
  for {
    select {
    case msg := <- c:
      fmt.Print(msg)
      time.Sleep(time.Second)
    case <- time.After(time.Second * 10):
      fmt.Println("Sleep over")
    }
  }
}

func main() {
  c := make(chan string)
  fmt.Println("Going to sleep")

  go ping(c)
  go sleepAfter(c)

  input := ""
  fmt.Scanln(&input)
}
