package numberstr

import "testing"

func TestNumStrDto_Multiply_00100(t *testing.T) {

	nStr1 := "35.123456"
	nStr2 := "47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_Multiply_00100() "

	var n1, n2, nExpected NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err = n1.Multiply(
		n2,
		ePrefix+"n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr,
		err = n1.GetNumStr(
		ePrefix + "n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	expectedNumStr, err = nExpected.GetNumStr(
		ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Actual Number String = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_MultiplyNumStrs_00100(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00100() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00200(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "-1685.4921624912384"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00200() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00300(t *testing.T) {
	nStr1 := "-35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00300() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00400(t *testing.T) {
	nStr1 := "57"
	nStr2 := "123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00400() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00500(t *testing.T) {
	nStr1 := "57"
	nStr2 := "-123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00500() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00600(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00600() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00700(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "-123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00700() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00800(t *testing.T) {
	nStr1 := "0"
	nStr2 := "123"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00800() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_00900(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_00900() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_01000(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_01000() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_01100(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0.123"
	nStr3 := "7.011"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_01100() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_01200(t *testing.T) {
	nStr1 := "62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_01200() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_01300(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "-194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_01300() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}

func TestNumStrDto_MultiplyNumStrs_01400(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "-3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	ePrefix := "TestNumStrDto_MultiplyNumStrs_01400() "

	var n1, n2, nExpected, nProduct NumStrDto
	var actualNumStr, expectedNumStr string
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2->n1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nExpected, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3->nExpected ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nProduct, err = nDto.MultiplyNumStrs(
		n1,
		n2,
		ePrefix+"n1 x n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	actualNumStr, err = nProduct.GetNumStr(
		ePrefix + "nProduct ")

	expectedNumStr, err =
		nExpected.GetNumStr(
			ePrefix + "nExpected ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Product Number String= '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsIntRunes())
	expectedNumStr = string(nExpected.GetAbsIntRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	actualNumStr = string(nProduct.GetAbsFracRunes())
	expectedNumStr = string(nExpected.GetAbsFracRunes())

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
		return
	}

	if nExpected.GetSign() != nProduct.GetSign() {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetSign(), nProduct.GetSign())
		return
	}

	if nExpected.HasNumericDigits() != nProduct.HasNumericDigits() {
		t.Errorf("Expected  nProduct.HasNumericDigist= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.HasNumericDigits(), nProduct.HasNumericDigits())
		return
	}

	if nExpected.IsFractionalValue() != nProduct.IsFractionalValue() {
		t.Errorf("Expected nProduct.IsFractionalValue= '%v'.\n"+
			"Instead, got '%v'\n",
			nExpected.IsFractionalValue(), nProduct.IsFractionalValue())
		return
	}

	if nExpected.GetPrecisionInt() != nProduct.GetPrecisionInt() {
		t.Errorf("Expected nProduct.precision= '%v'.\n"+
			"Instead, got %v\n",
			nExpected.GetPrecisionInt(), nProduct.GetPrecisionInt())
		return
	}

	err = nProduct.IsValidInstanceError(ePrefix + "nProduct ")

	if err != nil {
		t.Errorf("'nProduct' NumStr is INVALD! Error= %v\n", err.Error())
	}

}
