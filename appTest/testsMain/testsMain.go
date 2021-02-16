package testsMain

import (
	"fmt"
	numstr "github.com/MikeAustin71/numberstrops/numberstr"
	"strconv"
)

type TestMain struct {
	input  string
	output string
}

func (tMain *TestMain) Test000010AddNumStrs(
	ePrefix string) (
	err error) {

	errPrefixDto := numstr.ErrPrefixDto{}.NewEPrefOld(ePrefix)

	errPrefixDto.SetEPref("TestMain.TestAddNumStrs000010()")

	nStr1 := "-9589.21"
	nStr2 := "9211.40"
	nStr3 := "-377.81"
	var n1, n2, nExpected, nResult numstr.NumStrDto

	nDto := numstr.NumStrDto{}.New()

	n1, err = nDto.ParseNumStr(
		nStr1,
		errPrefixDto.XCtx("nStr1"))

	if err != nil {
		return err
	}

	n2, err = nDto.ParseNumStr(nStr2,
		errPrefixDto.XCtx("nStr2"))

	if err != nil {
		return err
	}

	nExpected, err = nDto.ParseNumStr(nStr3,
		errPrefixDto.XCtx("nStr3"))

	if err != nil {
		return err
	}

	nResult, err = nDto.AddNumStrs(n1, n2,
		errPrefixDto.XCtx("n1+n2"))

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
		return err
	}

	fmt.Printf(ePrefix + "- Successful Completion!\n")

	return err
}

func (tMain *TestMain) Test001000NewInt(
	ePrefix string) (
	err error) {

	// From Test TestNumStrDto_NewInt_0020()

	ePrefix += "TestNumStrDto_NewInt_0010() "

	intNum := 7
	precision := uint(0)

	expectedNumStr := "7"

	var nDto numstr.NumStrDto

	nXDto := numstr.NumStrDto{}

	nDto, err =
		nXDto.NewInt(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum= '%v' ",
					strconv.Itoa(intNum)))

	if err != nil {
		return err
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		return err
	}

	if expectedNumStr != actualNumStr {
		err = fmt.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

	return err
}
