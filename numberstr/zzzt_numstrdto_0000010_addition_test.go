package numberstr

import (
	"testing"
)

func TestNumStrDto_AddNumStrs_10(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_10() "
	nStr1 := "-9589.21"
	nStr2 := "9211.40"
	nStr3 := "-377.81"

	var err error
	var n1, n2, nExpected, nResult NumStrDto

	nDto := NumStrDto{}.New()
	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_20(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_20() "
	nStr1 := "9589.21"
	nStr2 := "-9211.40"
	nStr3 := "377.81"

	nDto := NumStrDto{}.New()

	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_AddNumStrs_30(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_30() "
	nStr1 := "9589.21"
	nStr2 := "9211.40"
	nStr3 := "18800.61"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_AddNumStrs_40(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_40() "
	nStr1 := "-9589.21"
	nStr2 := "-9211.40"
	nStr3 := "-18800.61"

	nDto := NumStrDto{}.New()

	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_AddNumStrs_50(t *testing.T) {
	ePrefix := " TestNumStrDto_AddNumStrs_50() "
	nStr1 := "2"
	nStr2 := "3"
	nStr3 := "5"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_AddNumStrs_60(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_60() "
	nStr1 := "2"
	nStr2 := "0.0"
	nStr3 := "2.0"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestNumStrDto_AddNumStrs_70(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_70() "
	nStr1 := "0"
	nStr2 := "0.0"
	nStr3 := "0.0"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_AddNumStrs_80(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_80() "
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr("nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

}

func TestNumStrDto_AddNumStrs_90(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_90() "
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

}

func TestNumStrDto_AddNumStrs_100(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_100() "
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_AddNumStrs_110(t *testing.T) {
	ePrefix := "TestNumStrDto_AddNumStrs_110() "
	nStr1 := "67.521"
	nStr2 := "-67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()
	var n1, n2, nExpected, nResult NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
		return
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
		return
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return
	}

	err = nResult.IsValidInstanceError(ePrefix + "nResult ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}
