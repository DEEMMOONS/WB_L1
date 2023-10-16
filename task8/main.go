package main

import "fmt"

func main() {
  var num int64
  var i, changedBit uint
  fmt.Scan(&num)
  fmt.Scan(&i)
  fmt.Scan(&changedBit)

  mask := int64(1 << i)
  if changedBit == 1 {
    num = num | mask
  } else {
    num = num & ^(mask)
  }
  fmt.Println(num)
}
