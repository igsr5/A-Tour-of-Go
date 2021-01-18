package main

import "fmt"
//import "time"
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

  //sqrt := sqrt(81)
  //defer fmt.Println(sqrt)

  //fmt.Println("When's Saturday?")
	//today := time.Now().Weekday()
	//switch time.Saturday {
	//case today + 0:
	//	fmt.Println("Today.")
	//case today + 1:
	//	fmt.Println("Tomorrow.")
	//case today + 2:
	//	fmt.Println("In two days.")
	//default:
	//	fmt.Println("Too far away.")
	//}
  //i, j := 42, 2701
  //p := &i
  //fmt.Println(*p)
  //*p = 21
  //fmt.Println(i)

  //p = &j
  //*p = *p / 37
  //fmt.Println(j)

  //type Vertex struct {
  //  X int
  //  Y int
  //}

  //v := &Vertex{1, 2}
  //v.X = 4
  //fmt.Println(v.X, v.Y)
   names := [4]string{
     "John",
     "Paul",
     "George",
     "Ringo",
   }
   fmt.Println(names)
  
   a := names[0:2]
   b := names[1:3]
   fmt.Println(a, b)

   b[0] = "XXX"
   fmt.Println(a, b)
   fmt.Println(names)
}
