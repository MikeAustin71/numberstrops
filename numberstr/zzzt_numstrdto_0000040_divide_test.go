package numberstr

import "testing"

func TestNumStrDto_Divide_10(t *testing.T) {

	ePrefix := "TestNumStrDto_Divide_10() "
	nStr1 := "12"
	nStr2 := "3"
	expectedStr := "4.0"
	var dividendDto, divisorDto NumStrDto
	var err error

	dividendDto, err = NumStrDto{}.NewNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	divisorDto, err = NumStrDto{}.NewNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nXDto := NumStrDto{}

	var quotient NumStrDto

	quotient,
		err = nXDto.DivideFractionNumStrs(
		dividendDto,
		divisorDto,
		1,
		ePrefix+"dividendDto/divisorDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = quotient.GetNumStr(ePrefix + "quotient ")

	if expectedStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'.",
			expectedStr, actualNumStr)
	}

}
