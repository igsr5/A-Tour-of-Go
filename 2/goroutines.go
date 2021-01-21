package main

import (
  "fmt"
  "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func sum (s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum
}

func fibonacci(c, quit chan int){
  x, y := 0, 1
  for {
    select {
    case c <- x:
      x, y = y, x+y
    case <-quit:
      fmt.Println("quit")
      return
    }
  }
}

func main() {
  //go say("world")
  //say("hello")

  //s := []int{7, 2, 8, -9, 4, 0}

  //c := make(chan int)
  //go sum(s[:len(s) / 2], c)
  //go sum(s[len(s) / 2:], c)

  //x, y := <-c, <-c

  //fmt.Println(x, y, x+y)

  c := make(chan int)
  quit := make(chan int)

  go func(){
    for i := 0; i < 10; i++ {
      a, ok := <-c
      fmt.Println(a)
      fmt.Println(ok)
    }
    quit<-0
  }()
  fibonacci(c, quit)
}

