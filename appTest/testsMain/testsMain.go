package testsMain

import (
	"fmt"
	numstr "github.com/MikeAustin71/numberstrops/numberstr"
)

type TestMain struct {
	input  string
	output string
}

func (tMain *TestMain) Test000010AddNumStrs(
	ePrefix string) (
	err error) {

	ePrefix += "TestMain.TestAddNumStrs000010() "

	nStr1 := "-9589.21"
	nStr2 := "9211.40"
	nStr3 := "-377.81"
	var n1, n2, nExpected, nResult numstr.NumStrDto

	nDto := numstr.NumStrDto{}.New()
	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		return err
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		return err
	}

	nExpected, err = nDto.ParseNumStr(nStr3, ePrefix+"nStr3 ")

	if err != nil {
		return err
	}

	nResult, err = nDto.AddNumStrs(n1, n2, ePrefix+"n1+n2 ")

	if err != nil {
		return err
	}

	var s, expected string

	s, err = nResult.GetNumStr(ePrefix + "nResult ")

	if err != nil {
		return err
	}

	expected, err = nExpected.GetNumStr(ePrefix + "nExpected ")

	if err != nil {
		return err
	}

	if s != expected {
		err = fmt.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
		return err
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		err = fmt.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
		return err
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		err = fmt.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
		return err
	}

	if nExpected.GetSign() != nResult.GetSign() {
		err = fmt.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
		return err
	}

	if nExpected.HasNumericDigits() != nResult.HasNumericDigits() {
		err = fmt.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits(), nResult.HasNumericDigits())
		return err
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		err = fmt.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
		return err
	}

	if nExpected.GetPrecisionInt() != nResult.GetPrecisionInt() {
		err = fmt.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecisionInt(), nResult.GetPrecisionInt())
		return err
	}

	err = nResult.IsValidInstanceError(ePrefix)

	if err != nil {
		err = fmt.Errorf("Error returned from nDto.IsValidInstanceError(&nResult). Error= %v", err)
	}

	return err
}
