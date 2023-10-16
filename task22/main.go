package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	_, ok := a.SetString("100000000000000000", 10)
	if !ok {
		fmt.Println("Wrong data")
	}
	_, ok = b.SetString("99999999999999999", 10)
	if !ok {
		fmt.Println("Wrong data")
	}

	result := new(big.Int)
	fmt.Print("new ")
  fmt.Println(result)

	result.Add(a,b)
	fmt.Print("add ")
	fmt.Println(result)

	result.Sub(result, b)
	fmt.Print("sub ")
	fmt.Println(result)

	result.Mul(a,b)
	fmt.Print("mul ")
	fmt.Println(result)

	if b.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Divison by zero, select another operation")
	} else {
		result.Div(a,b)
	  fmt.Print("div ")
		fmt.Println(result)
	}
}
