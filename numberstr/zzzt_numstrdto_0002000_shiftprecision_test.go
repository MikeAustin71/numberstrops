package numberstr

import "testing"

func TestNumStrDto_ShiftPrecisionLeft_000100(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000100() "

	nStr := "123456.789"
	precision := uint(3)
	expectedPrecision := uint(6)
	expectedNumStr := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000200(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000200() "

	nStr := "123456.789"
	precision := uint(2)
	expectedPrecision := uint(5)
	expectedNumStr := "1234.56789"
	signVal := 1
	absIntRuneStr := "1234"
	absFracRuneStr := "56789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000300(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000300() "

	nStr := "123456.789"
	precision := uint(6)
	expectedPrecision := uint(9)
	expectedNumStr := "0.123456789"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "123456789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000400(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000400() "

	nStr := "123456789"
	precision := uint(6)
	expectedPrecision := uint(6)
	expectedNumStr := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000500(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000500() "

	nStr := "123"
	precision := uint(5)
	expectedPrecision := uint(5)
	expectedNumStr := "0.00123"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00123"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000600(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000600() "

	nStr := "0"
	precision := uint(3)
	expectedPrecision := uint(3)
	expectedNumStr := "0.000"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "000"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000700(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000700() "

	nStr := "0.000"
	precision := uint(2)
	expectedPrecision := uint(5)
	expectedNumStr := "0.00000"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00000"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000800(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000800() "

	nStr := "123456.789"
	precision := uint(0)
	expectedPrecision := uint(3)
	expectedNumStr := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_000900(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_000900() "

	nStr := "-123456.789"
	precision := uint(0)
	expectedPrecision := uint(3)
	expectedNumStr := "-123456.789"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_001000(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_001000() "

	nStr := "-123456.789"
	precision := uint(3)
	expectedPrecision := uint(6)
	expectedNumStr := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_001100(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionLeft_001100() "

	nStr := "-123456789"
	precision := uint(6)
	expectedPrecision := uint(6)
	expectedNumStr := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionLeft(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_010(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_010() "

	nStr := "123456.789"
	precision := uint(3)
	expectedPrecision := uint(0)
	expectedNumStr := "123456789"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_020(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_020() "

	nStr := "123456.789"
	precision := uint(2)
	expectedPrecision := uint(1)
	expectedNumStr := "12345678.9"
	signVal := 1
	absIntRuneStr := "12345678"
	absFracRuneStr := "9"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_030(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_030() "

	nStr := "123456.789"
	precision := uint(6)
	expectedPrecision := uint(0)
	expectedNumStr := "123456789000"
	signVal := 1
	absIntRuneStr := "123456789000"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_040(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_040() "

	nStr := "123456789"
	precision := uint(6)
	expectedPrecision := uint(0)
	expectedNumStr := "123456789000000"
	signVal := 1
	absIntRuneStr := "123456789000000"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_050(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_050() "

	nStr := "123"
	precision := uint(5)
	expectedPrecision := uint(0)
	expectedNumStr := "12300000"
	signVal := 1
	absIntRuneStr := "12300000"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_060(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_060() "

	nStr := "0"
	precision := uint(3)
	expectedPrecision := uint(0)
	expectedNumStr := "0"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_070(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_070() "

	nStr := "123456.789"
	precision := uint(0)
	expectedPrecision := uint(3)
	expectedNumStr := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_080(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_080() "

	nStr := "-123456.789"
	precision := uint(0)
	expectedPrecision := uint(3)
	expectedNumStr := "-123456.789"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_090(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_090() "

	nStr := "-123456.789"
	precision := uint(3)
	expectedPrecision := uint(0)
	expectedNumStr := "-123456789"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}

func TestNumStrDto_ShiftPrecisionRight_100(t *testing.T) {

	ePrefix := "TestNumStrDto_ShiftPrecisionRight_100() "

	nStr := "-123456789"
	precision := uint(6)
	expectedPrecision := uint(0)
	expectedNumStr := "-123456789000000"
	signVal := -1
	absIntRuneStr := "123456789000000"
	absFracRuneStr := ""

	var nsDto, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}

	nsDto, err = nXDto.ShiftPrecisionRight(
		nStr,
		precision,
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
			"Instead, got %v.\n",
			absIntRuneStr, actualNumStr)
		return
	}

	actualNumStr = string(nsDto.GetAbsFracRunes())

	if absFracRuneStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes='%v'.\n"+
			"Instead, got %v\n", absFracRuneStr, actualNumStr)
	}

}
