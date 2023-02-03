package main

import (
	"fmt"
	"newTest/pack2"
	"newTest/testStruct"
)

// Main ep
func main() {
	fmt.Println("Syka")
	var (
		a = testStruct.TestStruct{}
		b = pack2.PackA{}
	)

	a.CrStr()
	b.CreatePoints()

	fmt.Println("start   ", a, "     ", b)

	a.A, a.B = 12, 21
	b.X, b.Y = 4, 6

	fmt.Println("changed ", a, "     ", b)

	fmt.Println("        ", a.GetSumStr(), "   -   ", b.SumPoint())

	fmt.Println("sq1= ", b.Square(), "  b = ", b)
	fmt.Println("sq2= ", b.Square2(), "  b = ", b)
}
