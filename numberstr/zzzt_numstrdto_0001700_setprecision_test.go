package numberstr

import (
	"fmt"
	"testing"
)

func TestNumStrDto_SetPrecision_10(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_10() "

	nStr := "-123456"
	precision := uint(3)
	expectedNumStr := "-123456.000"
	iStr := "123456"
	fracStr := "000"
	signVal := -1

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

func TestNumStrDto_SetPrecision_20(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_20() "

	nStr := "123456"
	precision := uint(9)
	expectedNumStr := "123456.000000000"
	iStr := "123456"
	fracStr := "000000000"
	signVal := 1

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

func TestNumStrDto_SetPrecision_30(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_30() "

	nStr := "123456.7890"
	precision := uint(2)
	expectedNumStr := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}
}

func TestNumStrDto_SetPrecision_40(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_40() "

	nStr := "123456.789"
	precision := uint(5)
	expectedNumStr := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_50(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_50() "

	nStr := "-123456.789"
	precision := uint(5)
	expectedNumStr := "-123456.78900"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_60(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_60() "

	nStr := "123456.789"
	precision := uint(3)
	expectedNumStr := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_70(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_70() "

	nStr := "123456789"
	precision := uint(7)
	expectedPrecision := uint(7)
	expectedNumStr := "123456789.0000000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_80(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_80() "

	nStr := "123456789"
	precision := uint(7)
	expectedPrecision := uint(7)
	expectedNumStr := "123456789.0000000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_90(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_90() "

	nStr := "-123456789"
	precision := uint(7)
	expectedPrecision := uint(7)
	expectedNumStr := "-123456789.0000000"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_100(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_100() "

	nStr := "-123456789"
	precision := uint(7)
	expectedPrecision := uint(7)
	expectedNumStr := "-123456789.0000000"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_110(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_110() "

	nStr := "123456.789"
	precision := uint(2)
	expectedPrecision := uint(2)
	expectedNumStr := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_120(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_120() "

	nStr := "123456.789"
	precision := uint(2)
	expectedPrecision := uint(2)
	expectedNumStr := "123456.78"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_130(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_130() "

	nStr := "123456.789"
	precision := uint(5)
	expectedPrecision := uint(5)
	expectedNumStr := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_140(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_140() "

	nStr := "123.456789"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "123.4"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "4"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_150(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_150() "

	nStr := "123.456789"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "123.5"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "5"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_160(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_160() "

	nStr := "-123.456789"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "-123.5"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "5"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_170(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_170() "

	nStr := "123456.789"
	precision := uint(0)
	expectedPrecision := uint(0)
	expectedNumStr := "123457"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_180(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_180() "

	nStr := "-123456.789"
	precision := uint(0)
	expectedPrecision := uint(0)
	expectedNumStr := "-123457"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_190(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_190() "

	nStr := "123456.789"
	precision := uint(0)
	expectedPrecision := uint(0)
	expectedNumStr := "123456"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_200(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_200() "

	nStr := "-123456.789"
	precision := uint(0)
	expectedPrecision := uint(0)
	expectedNumStr := "-123456"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_210(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_210() "

	nStr := "123457"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "123457.0"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_220(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_220() "

	nStr := "123457"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "123457.0"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_230(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_230() "

	nStr := "-123457"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "-123457.0"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_240(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_240() "

	nStr := "-123457"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "-123457.0"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		true,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetPrecision_250(t *testing.T) {

	ePrefix := "TestNumStrDto_SetPrecision_250() "

	nStr := "-123.456789"
	precision := uint(1)
	expectedPrecision := uint(1)
	expectedNumStr := "-123.4"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "4"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.SetPrecision(
		nStr,
		precision,
		false,
		ePrefix+"nStr ")

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetThisPrecision_10(t *testing.T) {

	ePrefix := "TestNumStrDto_SetThisPrecision_10() "

	nStr := "123456789"
	initialPrecision := uint(0)
	precision := uint(3)
	expectedPrecision := uint(3)
	expectedNumStr := "123456789.000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "000"

	var nsDto NumStrDto
	var err error

	nsDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if initialPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Error: Expected Initial precision='%v'.\n"+
			"Actual Initial precision='%v'.\n",
			initialPrecision, nsDto.GetPrecisionUint())
		return
	}

	err =
		nsDto.SetThisPrecision(
			precision,
			true,
			ePrefix+
				fmt.Sprintf("precision='%v' ", precision))

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_SetThisPrecision_20(t *testing.T) {

	ePrefix := "TestNumStrDto_SetThisPrecision_20() "

	nStr := "123456.789"
	initialPrecision := uint(3)
	precision := uint(2)
	expectedPrecision := uint(2)
	expectedNumStr := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	var nsDto NumStrDto
	var err error

	nsDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if initialPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Error: Expected Initial precision='%v'.\n"+
			"Actual Initial precision='%v'.\n",
			initialPrecision, nsDto.GetPrecisionUint())
		return
	}

	err =
		nsDto.SetThisPrecision(
			precision,
			true,
			ePrefix+
				fmt.Sprintf("precision='%v' ", precision))

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

	if expectedPrecision != nsDto.GetPrecisionUint() {
		t.Errorf("Expected precision='%v'.\n"+
			"Instead, got %v.\n",
			expectedPrecision, nsDto.GetPrecisionUint())
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

	if actualNumStr != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'.\n"+
			"Instead, got  AbsIntRunes='%v'.", absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if actualNumStr != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got AbsFracRunes='%v'.",
			absFracRuneStr, actualNumStr)
	}

}
