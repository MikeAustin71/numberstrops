package numberstr

import (
	"testing"
)

func TestNumStrDto_CompareAbsoluteVals_10(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_10() "
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompareResult := 1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_20(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_20() "
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompareResult := 1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(nStr1, ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(nStr2, ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err =
		nDto.CompareAbsoluteValues(&n1,
			&n2,
			ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_30(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_30 "
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompareResult := 0
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult, err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_40(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_40 "
	nStr1 := "567.21"
	nStr2 := "12567.218956"
	expectedCompareResult := -1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult, err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'.\n"+
			"Instead got '%v'",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_50(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_50() "
	nStr1 := "567.21"
	nStr2 := "-12567.218956"
	expectedCompareResult := -1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult, err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'.\n"+
			"Instead got '%v'",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_60(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_60() "
	nStr1 := "567.21"
	nStr2 := "-567.21"
	expectedCompareResult := 0
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult, err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'.\n"+
			"Instead got '%v'",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_70(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_70() "
	nStr1 := "567.21"
	nStr2 := "567.21"
	expectedCompareResult := 0
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult, err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' vs n2='%v'.\n"+
			"Expected Comparison Result= '%v'.\n"+
			"Instead got '%v'",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareAbsoluteVals_80(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareAbsoluteVals_80() "
	nStr1 := "567.21"
	nStr2 := "1567.21"
	expectedCompareResult := -1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult, err = nDto.CompareAbsoluteValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Absolute Values n1= '%v' vs n2='%v'.\n"+
			"Expected Comparison Result= '%v'.\n"+
			"Instead got '%v'",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_10(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_10() "
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompareResult := -1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_20(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_20() "
	nStr1 := "12567.218956"
	nStr2 := "9211.40"
	expectedCompareResult := 1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_30(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_30() "
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompareResult := -1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_40(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_40() "
	nStr1 := "12567.218956"
	nStr2 := "-9211.40"
	expectedCompareResult := 1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_50(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_50() "
	nStr1 := "12567.218956"
	nStr2 := "-12567.218956"
	expectedCompareResult := 1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_60(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_60() "
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompareResult := -1
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_70(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_70() "
	nStr1 := "-12567.218956"
	nStr2 := "-12567.218956"
	expectedCompareResult := 0
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}

func TestNumStrDto_CompareSignedVals_80(t *testing.T) {
	ePrefix := "TestNumStrDto_CompareSignedVals_80() "
	nStr1 := "12567.218956"
	nStr2 := "12567.218956"
	expectedCompareResult := 0
	nDto := NumStrDto{}.New()
	var n1, n2 NumStrDto
	var err error

	n1, err = nDto.ParseNumStr(
		nStr1,
		ePrefix+"nStr1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2, err = nDto.ParseNumStr(
		nStr2,
		ePrefix+"nStr2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualCompareResult int

	actualCompareResult,
		err = nDto.CompareSignedValues(
		&n1,
		&n2,
		ePrefix+"Compare n1 vs n2 ")

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'.\n"+
			"Expected Comparison= '%v'. Instead got '%v'\n",
			nStr1, nStr2, expectedCompareResult, actualCompareResult)
	}

}
