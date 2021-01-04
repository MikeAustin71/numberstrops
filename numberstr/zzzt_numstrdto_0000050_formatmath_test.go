package numberstr

import "testing"

func TestNumStrDto_FormatForMathOps_10(t *testing.T) {

	ePrefix := "TestNumStrDto_FormatForMathOps_10() "
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompareResult := 1

	nDto := NumStrDto{}.New()
	var n1, n2, nOut1, nOut2 NumStrDto
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

	nOut1, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nOut2, err = nDto.ParseNumStr(
		nStr4,
		ePrefix+"nStr4 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var n1OutDto, n2OutDto NumStrDto
	var actualCompareResult int
	var actualIsReversed bool

	n1OutDto,
		n2OutDto,
		actualCompareResult,
		actualIsReversed,
		err = nDto.FormatForMathOps(
		n1,
		n2,
		ePrefix+"n1,n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if false != actualIsReversed {
		t.Errorf("Expected isReverse = '%v'.\n"+
			"Instead got '%v'\n",
			false, actualIsReversed)
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Expected actualCompareResult result = '%v'.\n"+
			"Instead got '%v'\n",
			expectedCompareResult, actualCompareResult)
		return
	}

	var nOut1NumStr, n1OutDtoNumStr string

	nOut1NumStr,
		err = nOut1.GetNumStr(ePrefix + "nOut1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n1OutDtoNumStr,
		err = n1OutDto.GetNumStr(ePrefix + "n1OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut1NumStr != n1OutDtoNumStr {
		t.Errorf("Expected nOut1NumStr= '%v'\n"+
			" Instead got '%v'", nOut1NumStr, n1OutDtoNumStr)
		return
	}

	var nOut2NumStr, n2OutDtoNumStr string

	nOut2NumStr,
		err = nOut2.GetNumStr(ePrefix + "nOut2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2OutDtoNumStr,
		err = n2OutDto.GetNumStr(ePrefix + "n2OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut2NumStr != n2OutDtoNumStr {
		t.Errorf("Expected n2OutDtoNumStr= '%v'.\n"+
			"Instead got '%v'\n",
			nOut2NumStr, n2OutDtoNumStr)
		return
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'.\n"+
			"Instead got '%v'\n",
			nOut1.GetSign(), n1OutDto.GetSign())
		return
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()='%v'.\n"+
			"Instead got '%v'\n",
			nOut2.GetSign(), n2OutDto.GetSign())
		return
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'\n"+
			" Instead got '%v'\n",
			nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
		return
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()='%v'.\n"+
			"Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_20(t *testing.T) {
	ePrefix := "TestNumStrDto_FormatForMathOps_20() "
	nStr1 := "-9211.40"
	nStr2 := "-12567.218956"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompareResult := 1

	nDto := NumStrDto{}.New()
	var n1, n2, nOut1, nOut2 NumStrDto
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

	nOut1, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nOut2, err = nDto.ParseNumStr(
		nStr4,
		ePrefix+"nStr4 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var n1OutDto, n2OutDto NumStrDto
	var actualCompareResult int
	var actualIsReversed bool

	n1OutDto,
		n2OutDto,
		actualCompareResult,
		actualIsReversed,
		err =
		nDto.FormatForMathOps(n1,
			n2,
			ePrefix+"n1,n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if true != actualIsReversed {
		t.Errorf("Expected isReverse = '%v'.\n"+
			"Instead got '%v'\n", true, actualIsReversed)
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Expected actualCompareResult result = '%v'.\n"+
			"Instead got '%v'\n",
			expectedCompareResult, actualCompareResult)
		return
	}

	var nOut1NumStr, n1OutDtoNumStr string

	nOut1NumStr,
		err = nOut1.GetNumStr(ePrefix + "nOut1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n1OutDtoNumStr,
		err = n1OutDto.GetNumStr(ePrefix + "n1OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut1NumStr != n1OutDtoNumStr {
		t.Errorf("Expected nOut1NumStr= '%v'\n"+
			" Instead got '%v'", nOut1NumStr, n1OutDtoNumStr)
		return
	}

	var nOut2NumStr, n2OutDtoNumStr string

	nOut2NumStr,
		err = nOut2.GetNumStr(ePrefix + "nOut2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2OutDtoNumStr,
		err = n2OutDto.GetNumStr(ePrefix + "n2OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut2NumStr != n2OutDtoNumStr {
		t.Errorf("Expected n2OutDtoNumStr= '%v'.\n"+
			"Instead got '%v'\n",
			nOut2NumStr, n2OutDtoNumStr)
		return
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'.\n"+
			"Instead got '%v'\n",
			nOut1.GetSign(), n1OutDto.GetSign())
		return
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()='%v'.\n"+
			"Instead got '%v'\n",
			nOut2.GetSign(), n2OutDto.GetSign())
		return
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'\n"+
			" Instead got '%v'\n",
			nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
		return
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()='%v'.\n"+
			"Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_30(t *testing.T) {
	ePrefix := "TestNumStrDto_FormatForMathOps_30() "
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompareResult := 1

	nDto := NumStrDto{}.New()
	var n1, n2, nOut1, nOut2 NumStrDto
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

	nOut1, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nOut2, err = nDto.ParseNumStr(
		nStr4,
		ePrefix+"nStr4 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var n1OutDto, n2OutDto NumStrDto
	var actualCompareResult int
	var actualIsReversed bool

	n1OutDto,
		n2OutDto,
		actualCompareResult,
		actualIsReversed,
		err =
		nDto.FormatForMathOps(n1,
			n2,
			ePrefix+"n1,n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if true != actualIsReversed {
		t.Errorf("Expected isReverse = '%v'.\n"+
			"Instead got '%v'\n",
			true, actualIsReversed)
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Expected actualCompareResult result = '%v'.\n"+
			"Instead got '%v'\n",
			expectedCompareResult, actualCompareResult)
		return
	}

	var nOut1NumStr, n1OutDtoNumStr string

	nOut1NumStr,
		err = nOut1.GetNumStr(ePrefix + "nOut1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n1OutDtoNumStr,
		err = n1OutDto.GetNumStr(ePrefix + "n1OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut1NumStr != n1OutDtoNumStr {
		t.Errorf("Expected nOut1NumStr= '%v'\n"+
			" Instead got '%v'", nOut1NumStr, n1OutDtoNumStr)
		return
	}

	var nOut2NumStr, n2OutDtoNumStr string

	nOut2NumStr,
		err = nOut2.GetNumStr(ePrefix + "nOut2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2OutDtoNumStr,
		err = n2OutDto.GetNumStr(ePrefix + "n2OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut2NumStr != n2OutDtoNumStr {
		t.Errorf("Expected n2OutDtoNumStr= '%v'.\n"+
			"Instead got '%v'\n",
			nOut2NumStr, n2OutDtoNumStr)
		return
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'.\n"+
			"Instead got '%v'\n",
			nOut1.GetSign(), n1OutDto.GetSign())
		return
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()='%v'.\n"+
			"Instead got '%v'\n",
			nOut2.GetSign(), n2OutDto.GetSign())
		return
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'\n"+
			" Instead got '%v'\n",
			nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
		return
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()='%v'.\n"+
			"Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}
}

func TestNumStrDto_FormatForMathOps_40(t *testing.T) {
	ePrefix := "TestNumStrDto_FormatForMathOps_40() "
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompareResult := 1

	nDto := NumStrDto{}.New()
	var n1, n2, nOut1, nOut2 NumStrDto
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

	nOut1, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nOut2, err = nDto.ParseNumStr(
		nStr4,
		ePrefix+"nStr4 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var n1OutDto, n2OutDto NumStrDto
	var actualCompareResult int
	var actualIsReversed bool

	n1OutDto,
		n2OutDto,
		actualCompareResult,
		actualIsReversed,
		err =
		nDto.FormatForMathOps(n1,
			n2,
			ePrefix+"n1,n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if false != actualIsReversed {
		t.Errorf("Expected isReverse = '%v'.\n"+
			"Instead got '%v'\n",
			false, actualIsReversed)
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Expected actualCompareResult result = '%v'.\n"+
			"Instead got '%v'\n",
			expectedCompareResult, actualCompareResult)
		return
	}

	var nOut1NumStr, n1OutDtoNumStr string

	nOut1NumStr,
		err = nOut1.GetNumStr(ePrefix + "nOut1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n1OutDtoNumStr,
		err = n1OutDto.GetNumStr(ePrefix + "n1OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut1NumStr != n1OutDtoNumStr {
		t.Errorf("Expected nOut1NumStr= '%v'\n"+
			" Instead got '%v'", nOut1NumStr, n1OutDtoNumStr)
		return
	}

	var nOut2NumStr, n2OutDtoNumStr string

	nOut2NumStr,
		err = nOut2.GetNumStr(ePrefix + "nOut2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2OutDtoNumStr,
		err = n2OutDto.GetNumStr(ePrefix + "n2OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut2NumStr != n2OutDtoNumStr {
		t.Errorf("Expected n2OutDtoNumStr= '%v'.\n"+
			"Instead got '%v'\n",
			nOut2NumStr, n2OutDtoNumStr)
		return
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'.\n"+
			"Instead got '%v'\n",
			nOut1.GetSign(), n1OutDto.GetSign())
		return
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()='%v'.\n"+
			"Instead got '%v'\n",
			nOut2.GetSign(), n2OutDto.GetSign())
		return
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'\n"+
			" Instead got '%v'\n",
			nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
		return
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()='%v'.\n"+
			"Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_50(t *testing.T) {
	ePrefix := "TestNumStrDto_FormatForMathOps_50() "
	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-67.521"
	nStr4 := "06.000"
	expectedCompareResult := 1

	nDto := NumStrDto{}.New()
	var n1, n2, nOut1, nOut2 NumStrDto
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

	nOut1, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nOut2, err = nDto.ParseNumStr(
		nStr4,
		ePrefix+"nStr4 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var n1OutDto, n2OutDto NumStrDto
	var actualCompareResult int
	var actualIsReversed bool

	n1OutDto,
		n2OutDto,
		actualCompareResult,
		actualIsReversed,
		err =
		nDto.FormatForMathOps(n1,
			n2,
			ePrefix+"n1,n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if false != actualIsReversed {
		t.Errorf("Expected isReverse = '%v'.\n"+
			"Instead got '%v'\n",
			false, actualIsReversed)
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Expected actualCompareResult result = '%v'.\n"+
			"Instead got '%v'\n",
			expectedCompareResult, actualCompareResult)
		return
	}

	var nOut1NumStr, n1OutDtoNumStr string

	nOut1NumStr,
		err = nOut1.GetNumStr(ePrefix + "nOut1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n1OutDtoNumStr,
		err = n1OutDto.GetNumStr(ePrefix + "n1OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut1NumStr != n1OutDtoNumStr {
		t.Errorf("Expected nOut1NumStr= '%v'\n"+
			" Instead got '%v'", nOut1NumStr, n1OutDtoNumStr)
		return
	}

	var nOut2NumStr, n2OutDtoNumStr string

	nOut2NumStr,
		err = nOut2.GetNumStr(ePrefix + "nOut2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2OutDtoNumStr,
		err = n2OutDto.GetNumStr(ePrefix + "n2OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut2NumStr != n2OutDtoNumStr {
		t.Errorf("Expected n2OutDtoNumStr= '%v'.\n"+
			"Instead got '%v'\n",
			nOut2NumStr, n2OutDtoNumStr)
		return
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'.\n"+
			"Instead got '%v'\n",
			nOut1.GetSign(), n1OutDto.GetSign())
		return
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()='%v'.\n"+
			"Instead got '%v'\n",
			nOut2.GetSign(), n2OutDto.GetSign())
		return
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'\n"+
			" Instead got '%v'\n",
			nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
		return
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()='%v'.\n"+
			"Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}

func TestNumStrDto_FormatForMathOps_60(t *testing.T) {
	ePrefix := "TestNumStrDto_FormatForMathOps_60() "
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "-67.521"
	nStr4 := "67.521"
	expectedCompareResult := 0

	nDto := NumStrDto{}.New()
	var n1, n2, nOut1, nOut2 NumStrDto
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

	nOut1, err = nDto.ParseNumStr(
		nStr3,
		ePrefix+"nStr3 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nOut2, err = nDto.ParseNumStr(
		nStr4,
		ePrefix+"nStr4 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var n1OutDto, n2OutDto NumStrDto
	var actualCompareResult int
	var actualIsReversed bool

	n1OutDto,
		n2OutDto,
		actualCompareResult,
		actualIsReversed,
		err =
		nDto.FormatForMathOps(n1,
			n2,
			ePrefix+"n1,n2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if false != actualIsReversed {
		t.Errorf("Expected isReverse = '%v'.\n"+
			"Instead got '%v'\n",
			false, actualIsReversed)
		return
	}

	if actualCompareResult != expectedCompareResult {
		t.Errorf("Expected actualCompareResult result = '%v'.\n"+
			"Instead got '%v'\n",
			expectedCompareResult, actualCompareResult)
		return
	}

	var nOut1NumStr, n1OutDtoNumStr string

	nOut1NumStr,
		err = nOut1.GetNumStr(ePrefix + "nOut1 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n1OutDtoNumStr,
		err = n1OutDto.GetNumStr(ePrefix + "n1OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut1NumStr != n1OutDtoNumStr {
		t.Errorf("Expected nOut1NumStr= '%v'\n"+
			" Instead got '%v'", nOut1NumStr, n1OutDtoNumStr)
		return
	}

	var nOut2NumStr, n2OutDtoNumStr string

	nOut2NumStr,
		err = nOut2.GetNumStr(ePrefix + "nOut2 ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	n2OutDtoNumStr,
		err = n2OutDto.GetNumStr(ePrefix + "n2OutDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nOut2NumStr != n2OutDtoNumStr {
		t.Errorf("Expected n2OutDtoNumStr= '%v'.\n"+
			"Instead got '%v'\n",
			nOut2NumStr, n2OutDtoNumStr)
		return
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'.\n"+
			"Instead got '%v'\n",
			nOut1.GetSign(), n1OutDto.GetSign())
		return
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()='%v'.\n"+
			"Instead got '%v'\n",
			nOut2.GetSign(), n2OutDto.GetSign())
		return
	}

	if nOut1.GetPrecisionInt() != n1OutDto.GetPrecisionInt() {
		t.Errorf("Expected n1OutDto.GetPrecisionInt()= '%v'\n"+
			" Instead got '%v'\n",
			nOut1.GetPrecisionInt(), n1OutDto.GetPrecisionInt())
		return
	}

	if nOut2.GetPrecisionInt() != n2OutDto.GetPrecisionInt() {
		t.Errorf("Expected n2OutDto.GetPrecisionInt()='%v'.\n"+
			"Instead got '%v'", nOut2.GetPrecisionInt(), n2OutDto.GetPrecisionInt())
	}

}
