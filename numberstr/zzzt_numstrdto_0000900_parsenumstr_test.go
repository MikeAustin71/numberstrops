package numberstr

import (
	"fmt"
	"strings"
	"testing"
)

func TestNumStrDto_ParseNumStr_10(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_10() "

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n",
			nStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_20(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_20() "

	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n",
			nStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, actualNumStr)

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_30(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_30() "

	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n",
			nStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, actualNumStr)

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_40(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_40() "

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n",
			nStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, actualNumStr)

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_50(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_50() "

	nStr := "-000.000"
	expectedStr := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != expectedStr {
		t.Errorf("Expected Number String = '%v'.\n"+
			"Instead, got %v\n",
			expectedStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, actualNumStr)

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestNumStrDto_ParseNumStr_60(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_60() "

	nStr := "-123.4567#"
	expectedStr := "-123.4567"
	iStr := "123"
	fracStr := "4567"
	signVal := -1
	precision := uint(4)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != expectedStr {
		t.Errorf("Expected Number String = '%v'.\n"+
			"Instead, got %v\n",
			expectedStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, actualNumStr)

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}
}

func TestNumStrDto_ParseNumStr_70(t *testing.T) {

	ePrefix := "TestNumStrDto_ParseNumStr_70() "

	nStr := "-123.4#-567#"
	expectedStr := "-123.4567"
	iStr := "123"
	fracStr := "4567"
	signVal := -1
	precision := uint(4)

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nDto, err = nXDto.ParseNumStr(
		nStr,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != expectedStr {
		t.Errorf("Expected Number String = '%v'.\n"+
			"Instead, got %v\n",
			expectedStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, actualNumStr)

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if signVal != nDto.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue='%v'.\n"+
			"Instead, got IsFractionalValue='%v'.\n",
			isFractionalValue, nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}
}
