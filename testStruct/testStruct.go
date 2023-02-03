package testStruct

import (
	"fmt"
	"math"
	"strconv"
)

type TestStruct struct {
	A, B int
	Str  string
}

func (tS *TestStruct) CrStr() {
	tS.A = 0
	tS.B = 0
	tS.Str = ""
}

func (tS *TestStruct) GetSumStr() string {

	tS.Str = strconv.Itoa(tS.A + tS.B)
	return tS.Str
}

func (tS *TestStruct) GetSubStr() string {

	tS.Str = fmt.Sprint(math.Abs(float64(tS.A) - float64(tS.B)))
	return tS.Str
}
