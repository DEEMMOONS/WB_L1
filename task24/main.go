package main

import (
  "fmt"
  "math"
)

type Point struct {
  x, y float64
}

func NewPoint (x, y float64) Point {
  return Point{x, y}
}

func (p *Point) Distance(other Point) float64 {
  dx := other.x - p.x
  dy := other.y - p.y
  return math.Sqrt(dx * dx + dy * dy)
}

func main() {
  point1 := NewPoint(1, 2)
  point2 := NewPoint(4, 6)
  fmt.Println(point1.Distance(point2))
}
