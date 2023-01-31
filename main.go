package main

import (
	"fmt"
	"newTest/testStruct"
)

func main() {
	fmt.Println("Syka")
	a := testStruct.TestStruct{}
	a.CrStr()
	fmt.Println(a)
	a.A = 12
	a.B = 21
	fmt.Println(a.GetSumStr())
}
