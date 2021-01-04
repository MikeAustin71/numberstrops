package numberstr

import "testing"

func TestNumStrDto_CopyIn_10(t *testing.T) {
	ePrefix := "TestNumStrDto_CopyIn_10() "

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDtoX := NumStrDto{}.New()

	var n1 NumStrDto
	var err error

	n1, err = nDtoX.ParseNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	var s string

	s, err = nDto.GetNumStr(ePrefix + "nDto ")

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", nStr, s)
		return
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, s)
		return
	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, s)
		return
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'.\n"+
			"Instead, got %v\n", nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return

	}

	err = nDto.IsValidInstanceError(ePrefix + "Testing validity of 'nDto' ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_CopyIn_20(t *testing.T) {
	ePrefix := "TestNumStrDto_CopyIn_20() "

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDtoX := NumStrDto{}.New()

	var n1 NumStrDto
	var err error

	n1, err = nDtoX.ParseNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	var s string

	s, err = nDto.GetNumStr(ePrefix + "nDto ")

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", nStr, s)
		return
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, s)
		return
	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, s)
		return
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'.\n"+
			"Instead, got %v\n", nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return

	}

	err = nDto.IsValidInstanceError(ePrefix + "Testing validity of 'nDto' ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_CopyOut_10(t *testing.T) {

	ePrefix := "TestNumStrDto_CopyOut_10() "

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDtoX := NumStrDto{}.New()

	var n1 NumStrDto
	var err error

	n1, err = nDtoX.ParseNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto := n1.CopyOut()

	var s string

	s, err = nDto.GetNumStr(ePrefix + "nDto ")

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", nStr, s)
		return
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, s)
		return
	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, s)
		return
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'.\n"+
			"Instead, got %v\n", nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return

	}

	err = nDto.IsValidInstanceError(ePrefix + "Testing validity of 'nDto' ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}

func TestNumStrDto_CopyOut_20(t *testing.T) {

	ePrefix := "TestNumStrDto_CopyOut_20() "

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDtoX := NumStrDto{}.New()

	var n1 NumStrDto
	var err error

	n1, err = nDtoX.ParseNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto := n1.CopyOut()

	var s string

	s, err = nDto.GetNumStr(ePrefix + "nDto ")

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n", nStr, s)
		return
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			iStr, s)
		return
	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, s)
		return
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'.\n"+
			"Instead, got %v\n", nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return

	}

	err = nDto.IsValidInstanceError(ePrefix + "Testing validity of 'nDto' ")

	if err != nil {
		t.Errorf("%v", err.Error())
	}

}
