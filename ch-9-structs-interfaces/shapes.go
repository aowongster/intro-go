package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  x, y, r float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

// aka circumference
func (c *Circle) perimeter() float64 {
  return math.Pi * 2 * c.r
}

type Rectangle struct {
  x1, x2, y1, y2 float64
}

// easier to think of h x w, h=> y, w => x
func (r *Rectangle) area() float64 {
  h := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)

  return h * w
}

func (r *Rectangle) perimeter() float64 {
  h := distance(r.x1, r.y1) // calc.. individually... or?
}

// pythags theorem a^2 + b^2 = c^2
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func main() {

}
