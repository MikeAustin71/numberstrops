package numberstr

import (
	"fmt"
	"testing"
)

func TestNumStrDto_ScaleNumStr_10(t *testing.T) {

	ePrefix := "TestNumStrDto_ScaleNumStr_10"

	// "123456.789"  3   "123.456789"
	nStr := "123456.789"
	scaleMode := SCALEPRECISIONLEFT
	scale := uint(3)
	precision := uint(6)
	expectedNumStr := "123.456789"
	iStr := "123"
	fracStr := "456789"
	signVal := 1

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err =
		nXDto.ScaleNumStr(
			nStr,
			scale,
			scaleMode,
			ePrefix+
				fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = nsDto.GetNumStr(
		ePrefix + "nsDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != expectedNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, got Number String='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

	if precision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			precision, nsDto.GetPrecisionUint())
		return
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'.\n"+
			"Instead, got signVal='%v'\n",
			signVal, nsDto.GetSign())
		return
	}

	err = nsDto.IsValidInstanceError(
		ePrefix + "Testing 'nsDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nsDto.HasNumericDigits() {
		t.Errorf("Expected nsDto.HasNumericDigits='true'.\n"+
			"Instead, got %v.\n", nsDto.HasNumericDigits())
		return
	}

	actualNumStr = string(nsDto.GetAbsIntRunes())

	if actualNumStr != iStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", iStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			fracStr, actualNumStr)
	}
}
