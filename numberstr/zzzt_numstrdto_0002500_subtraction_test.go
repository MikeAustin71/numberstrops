package numberstr

import (
	"fmt"
	"testing"
)

func TestNumStrDto_SubtractNumStrs_10(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_10() "

	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "73.521"

	var nDto NumStrDto
	var err error

	nDto = NumStrDto{}.New()

	var n1, n2, nResult NumStrDto

	n1, err =
		nDto.ParseNumStr(
			nStr1,
			ePrefix+
				fmt.Sprintf("nStr1=%v ", nStr1))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err =
		nDto.ParseNumStr(
			nStr2,
			ePrefix+
				fmt.Sprintf("nStr2=%v ", nStr2))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err =
		nDto.ParseNumStr(
			nStr3,
			ePrefix+
				fmt.Sprintf("nStr3=%v ", nStr3))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto, err =
		nDto.SubtractNumStrs(
			n1,
			n2,
			ePrefix+"n1 - n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr, expectedNumStr string

	actualNumStr,
		err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr,
		err = nResult.GetNumStr(
		ePrefix + "nResult ")

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

	actualNumStr = string(nDto.GetAbsIntRunes())
	expectedNumStr = string(nResult.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())
	expectedNumStr = string(nResult.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expectedNumStr, actualNumStr)
		return
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'\n"+
			"Instead, got %v\n",
			nResult.GetSign(), nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v", nDto.HasNumericDigits())
		return
	}

	if nDto.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got '%v'\n",
			nResult.IsFractionalValue(), nDto.IsFractionalValue())
		return
	}

	if nDto.GetPrecisionUint() != nResult.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nResult.GetPrecisionUint(), nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError("Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_SubtractNumStrs_20(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_20() "

	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-73.521"

	var nDto NumStrDto
	var err error

	nDto = NumStrDto{}.New()

	var n1, n2, nResult NumStrDto

	n1, err =
		nDto.ParseNumStr(
			nStr1,
			ePrefix+
				fmt.Sprintf("nStr1=%v ", nStr1))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err =
		nDto.ParseNumStr(
			nStr2,
			ePrefix+
				fmt.Sprintf("nStr2=%v ", nStr2))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err =
		nDto.ParseNumStr(
			nStr3,
			ePrefix+
				fmt.Sprintf("nStr3=%v ", nStr3))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto, err =
		nDto.SubtractNumStrs(
			n1,
			n2,
			ePrefix+"n1 - n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr, expectedNumStr string

	actualNumStr,
		err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr,
		err = nResult.GetNumStr(
		ePrefix + "nResult ")

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

	actualNumStr = string(nDto.GetAbsIntRunes())
	expectedNumStr = string(nResult.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())
	expectedNumStr = string(nResult.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expectedNumStr, actualNumStr)
		return
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'\n"+
			"Instead, got %v\n",
			nResult.GetSign(), nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v", nDto.HasNumericDigits())
		return
	}

	if nDto.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got '%v'\n",
			nResult.IsFractionalValue(), nDto.IsFractionalValue())
		return
	}

	if nDto.GetPrecisionUint() != nResult.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nResult.GetPrecisionUint(), nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError("Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_SubtractNumStrs_30(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_30() "

	nStr1 := "67.521"
	nStr2 := "691.1"
	nStr3 := "-623.579"

	var nDto NumStrDto
	var err error

	nDto = NumStrDto{}.New()

	var n1, n2, nResult NumStrDto

	n1, err =
		nDto.ParseNumStr(
			nStr1,
			ePrefix+
				fmt.Sprintf("nStr1=%v ", nStr1))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err =
		nDto.ParseNumStr(
			nStr2,
			ePrefix+
				fmt.Sprintf("nStr2=%v ", nStr2))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err =
		nDto.ParseNumStr(
			nStr3,
			ePrefix+
				fmt.Sprintf("nStr3=%v ", nStr3))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto, err =
		nDto.SubtractNumStrs(
			n1,
			n2,
			ePrefix+"n1 - n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr, expectedNumStr string

	actualNumStr,
		err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr,
		err = nResult.GetNumStr(
		ePrefix + "nResult ")

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

	actualNumStr = string(nDto.GetAbsIntRunes())
	expectedNumStr = string(nResult.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())
	expectedNumStr = string(nResult.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expectedNumStr, actualNumStr)
		return
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'\n"+
			"Instead, got %v\n",
			nResult.GetSign(), nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v", nDto.HasNumericDigits())
		return
	}

	if nDto.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got '%v'\n",
			nResult.IsFractionalValue(), nDto.IsFractionalValue())
		return
	}

	if nDto.GetPrecisionUint() != nResult.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nResult.GetPrecisionUint(), nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError("Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_SubtractNumStrs_40(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_40() "

	nStr1 := "691.1"
	nStr2 := "67.521"
	nStr3 := "623.579"

	var nDto NumStrDto
	var err error

	nDto = NumStrDto{}.New()

	var n1, n2, nResult NumStrDto

	n1, err =
		nDto.ParseNumStr(
			nStr1,
			ePrefix+
				fmt.Sprintf("nStr1=%v ", nStr1))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err =
		nDto.ParseNumStr(
			nStr2,
			ePrefix+
				fmt.Sprintf("nStr2=%v ", nStr2))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err =
		nDto.ParseNumStr(
			nStr3,
			ePrefix+
				fmt.Sprintf("nStr3=%v ", nStr3))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto, err =
		nDto.SubtractNumStrs(
			n1,
			n2,
			ePrefix+"n1 - n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr, expectedNumStr string

	actualNumStr,
		err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr,
		err = nResult.GetNumStr(
		ePrefix + "nResult ")

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

	actualNumStr = string(nDto.GetAbsIntRunes())
	expectedNumStr = string(nResult.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())
	expectedNumStr = string(nResult.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expectedNumStr, actualNumStr)
		return
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'\n"+
			"Instead, got %v\n",
			nResult.GetSign(), nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v", nDto.HasNumericDigits())
		return
	}

	if nDto.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got '%v'\n",
			nResult.IsFractionalValue(), nDto.IsFractionalValue())
		return
	}

	if nDto.GetPrecisionUint() != nResult.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nResult.GetPrecisionUint(), nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError("Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_SubtractNumStrs_50(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_50() "

	nStr1 := "691.1"
	nStr2 := "0"
	nStr3 := "691.1"

	var nDto NumStrDto
	var err error

	nDto = NumStrDto{}.New()

	var n1, n2, nResult NumStrDto

	n1, err =
		nDto.ParseNumStr(
			nStr1,
			ePrefix+
				fmt.Sprintf("nStr1=%v ", nStr1))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err =
		nDto.ParseNumStr(
			nStr2,
			ePrefix+
				fmt.Sprintf("nStr2=%v ", nStr2))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err =
		nDto.ParseNumStr(
			nStr3,
			ePrefix+
				fmt.Sprintf("nStr3=%v ", nStr3))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto, err =
		nDto.SubtractNumStrs(
			n1,
			n2,
			ePrefix+"n1 - n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr, expectedNumStr string

	actualNumStr,
		err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr,
		err = nResult.GetNumStr(
		ePrefix + "nResult ")

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

	actualNumStr = string(nDto.GetAbsIntRunes())
	expectedNumStr = string(nResult.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())
	expectedNumStr = string(nResult.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expectedNumStr, actualNumStr)
		return
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'\n"+
			"Instead, got %v\n",
			nResult.GetSign(), nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v", nDto.HasNumericDigits())
		return
	}

	if nDto.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got '%v'\n",
			nResult.IsFractionalValue(), nDto.IsFractionalValue())
		return
	}

	if nDto.GetPrecisionUint() != nResult.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nResult.GetPrecisionUint(), nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError("Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_SubtractNumStrs_60(t *testing.T) {

	ePrefix := "TestNumStrDto_SubtractNumStrs_60() "

	nStr1 := "0"
	nStr2 := "691.1"
	nStr3 := "-691.1"

	var nDto NumStrDto
	var err error

	nDto = NumStrDto{}.New()

	var n1, n2, nResult NumStrDto

	n1, err =
		nDto.ParseNumStr(
			nStr1,
			ePrefix+
				fmt.Sprintf("nStr1=%v ", nStr1))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err =
		nDto.ParseNumStr(
			nStr2,
			ePrefix+
				fmt.Sprintf("nStr2=%v ", nStr2))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nResult, err =
		nDto.ParseNumStr(
			nStr3,
			ePrefix+
				fmt.Sprintf("nStr3=%v ", nStr3))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto, err =
		nDto.SubtractNumStrs(
			n1,
			n2,
			ePrefix+"n1 - n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr, expectedNumStr string

	actualNumStr,
		err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr,
		err = nResult.GetNumStr(
		ePrefix + "nResult ")

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

	actualNumStr = string(nDto.GetAbsIntRunes())
	expectedNumStr = string(nResult.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())
	expectedNumStr = string(nResult.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", expectedNumStr, actualNumStr)
		return
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'\n"+
			"Instead, got %v\n",
			nResult.GetSign(), nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v", nDto.HasNumericDigits())
		return
	}

	if nDto.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got '%v'\n",
			nResult.IsFractionalValue(), nDto.IsFractionalValue())
		return
	}

	if nDto.GetPrecisionUint() != nResult.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			nResult.GetPrecisionUint(), nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError("Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}
