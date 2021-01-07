package numberstr

import (
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
