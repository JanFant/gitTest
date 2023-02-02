package main

import (
	"fmt"
	"newTest/pack2"
	"newTest/testStruct"
)

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

}
