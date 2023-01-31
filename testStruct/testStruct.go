package testStruct

import "strconv"

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

	tS.Str = strconv.Itoa(tS.A) + strconv.Itoa(tS.B)
	return tS.Str
}
