package main

import "fmt"
//import "math"

func add(x, y int) int {
  return x + y
}

func swap(x, y string) (string, string) {
  return y, x
}

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func sqrt(x float64) float64 {
  z := 1.0
  for z*z != x {
    z -= (z*z - x) / (2 * z)
  }
  return z
}

func main() {
  //var hello string = "hello world!"
  //fmt.Println(hello)

  //fmt.Println(add(42, 13))

  //a, b := swap("hello", "world")
  //fmt.Println(a, b)

  //x, y := split(10)
  //fmt.Println(x,y)

  //var x1, y1 int = 3, 4
  //var f float64 = math.Sqrt(float64(x1*x1 + y1*y1))
  //var z uint = uint(f)
  //fmt.Println(x1, y1, z)

  //const PI = 1 << 1
  //fmt.Println(PI)

  //for i := 0; i < 10; i++ {
  //  fmt.Println(i)
  //  if i % 2 == 0 {
  //    fmt.Println("even")
  //  }
  //}
  //sum := 1
  //for sum < 100 {
  //  sum +=2
  //  fmt.Println(sum)
  //}

  sqrt := sqrt(81)
  fmt.Println(sqrt)
}
