package numberstr

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNumStrDto_GetAbsoluteBigInt_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetAbsoluteBigInt_10() "
	nStr := "-123.456"
	absNumStr := "123456"
	expected, isOk := big.NewInt(0).SetString(absNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(absNumStr,10) Failed!. absNumStr= '%v'", absNumStr)
	}

	nxDto := NumStrDto{}

	n1, err := nxDto.ParseNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var absBigInt *big.Int
	absBigInt, err = n1.GetAbsoluteBigInt(ePrefix)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if absBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected absBigInt= %v .\n"+
			"Instead got, %v",
			expected.String(), absBigInt.String())
		return
	}

}

func TestNumStrDto_GetAbsFracRunes_10(t *testing.T) {
	ePrefix := "TestNumStrDto_GetAbsFracRunes_10() "
	nStr := "-123.456"
	expectedRunes := []rune{'4', '5', '6'}

	nDto, err := NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"Creating nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	fracRunes := nDto.GetAbsFracRunes()

	if len(expectedRunes) != len(fracRunes) {
		t.Errorf("Expected a rune array length of '%v'.\n"+
			"Instead length='%v' .\n",
			len(expectedRunes), len(fracRunes))
		return
	}

	for i := 0; i < len(fracRunes); i++ {
		if expectedRunes[i] != fracRunes[i] {
			t.Errorf("Error! Rune Idx='%v'.\nExpected rune='%v'.\n"+
				"Instead rune='%v'.\n",
				i, expectedRunes[i], fracRunes[i])
			return
		}

	}

}

func TestNumStrDto_GetAbsFracRunes_20(t *testing.T) {

	ePrefix := "TestNumStrDto_GetAbsFracRunes_20() "

	nStr := "123"

	nDto, err := NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"Creating nDto ")

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	fracRunes := nDto.GetAbsFracRunes()

	if len(fracRunes) != 0 {
		t.Errorf("Expected rune array length=0.\n"+
			"Instead length='%v'\n"+
			"Array String='%v'\n",
			len(fracRunes), string(fracRunes))
	}

}

func TestNumStrDto_GetAbsIntRunes_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetAbsIntRunes_10() "

	nStr := "-123.456"
	expectedRunes := []rune{'1', '2', '3'}

	nDto, err := NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	intRunes := nDto.GetAbsIntRunes()

	if len(expectedRunes) != len(intRunes) {
		t.Errorf("Error: Expected Rune Array Length='%v'.\n"+
			"Instead, Array Length='%v'\n",
			len(expectedRunes), len(intRunes))
		return
	}

	for i := 0; i < len(intRunes); i++ {
		if expectedRunes[i] != intRunes[i] {
			t.Errorf("Error: Index='%v' Expected Rune='%v'. Actual Rune='%v'",
				i, expectedRunes[i], intRunes[i])
			return
		}
	}

}

func TestNumStrDto_GetAbsIntRunes_20(t *testing.T) {

	ePrefix := "TestNumStrDto_GetAbsIntRunes_10() "
	nStr := "-123"
	expectedRunes := []rune{'1', '2', '3'}

	nDto, err := NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	intRunes := nDto.GetAbsIntRunes()

	if len(expectedRunes) != len(intRunes) {
		t.Errorf("Error: Expected Rune Array Length='%v'.\n"+
			"Instead, Array Length='%v'\n",
			len(expectedRunes), len(intRunes))
		return
	}

	for i := 0; i < len(intRunes); i++ {
		if expectedRunes[i] != intRunes[i] {
			t.Errorf("Error: Index='%v'\n"+
				"Expected Rune='%v'.\n"+
				"Actual Rune='%v'\n",
				i, expectedRunes[i], intRunes[i])
			return
		}
	}

}

func TestNumStrDto_GetAbsIntRunes_30(t *testing.T) {

	ePrefix := "TestNumStrDto_GetAbsIntRunes_30() "

	nStr := "123456789.7865"
	expectedRunes := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	nDto, err := NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	intRunes := nDto.GetAbsIntRunes()

	if len(expectedRunes) != len(intRunes) {
		t.Errorf("Error: Expected Rune Array Length='%v'.  Instead, Array Length='%v'",
			len(expectedRunes), len(intRunes))
	}

	for i := 0; i < len(intRunes); i++ {
		if expectedRunes[i] != intRunes[i] {
			t.Errorf("Error: Index='%v' Expected Rune='%v'. Actual Rune='%v'",
				i, expectedRunes[i], intRunes[i])
		}
	}

}
