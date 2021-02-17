package numberstr

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNumStrDto_GetBigIntNum_10(t *testing.T) {

	ePrefix := ErrPrefixDto{}.NewEPrefOld(
		"TestNumStrDto_GetBigIntNum_10()")

	bigI := big.NewInt(123456123456)
	expectedPrecision := uint(0)

	expectedBigIntStr := "123456123456"

	expectedBigIntNum :=
		big.NewInt(0).Set(bigI)

	var actualBigIntNDto NumStrDto
	var err error

	actualBigIntNDto,
		err = NumStrDto{}.NewBigInt(
		bigI,
		0,
		ePrefix.ZCtx("bigI"))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		actualBigIntNDto.GetNumStr(
			ePrefix.ZCtx("actualBigIntNDto"))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedBigIntStr != actualNumStr {
		t.Errorf("Error: Expected Big Int Number String ='%v'.\n"+
			"Instead, Big Int Number String ='%v'",
			expectedBigIntStr, actualNumStr)
		return
	}

	var actualBigIntNum *big.Int

	actualBigIntNum,
		err = actualBigIntNDto.GetBigInt(
		ePrefix.ZCtx("actualBigIntNDto"))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	compareResult :=
		actualBigIntNum.Cmp(expectedBigIntNum)

	if compareResult != 0 {
		t.Errorf("Error: Expected actualBigIntNum='%v'\n"+
			"Instead, actualBigIntNum='%v'\n",
			expectedBigIntNum.Text(10),
			actualBigIntNum.Text(10))
		return
	}

	actualPrecision := actualBigIntNDto.GetPrecisionUint()

	if expectedPrecision != actualPrecision {
		t.Errorf("Error: Expected 'actualBigIntNDto' Precision='%v'\n"+
			"Instead, 'actualBigIntNDto' Precision='%v'\n",
			expectedPrecision, actualPrecision)
	}

}

func TestNumStrDto_GetSignedBigInt_10(t *testing.T) {

	nStr := "-123.456"
	signedNumStr := "-123456"
	expected, isOk := big.NewInt(0).SetString(signedNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(signedNumStr,10) Failed!.\n"+
			"signedNumStr= '%v'\n", signedNumStr)
		return
	}

	ePrefix := ErrPrefixDto{}.NewEPrefOld("TestNumStrDto_GetSignedBigInt_10")

	var n1, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}.New()

	n1, err =
		nXDto.ParseNumStr(
			nStr,
			ePrefix.ZCtx("nStr"))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var signedBigInt *big.Int

	signedBigInt, err = n1.GetBigInt(
		ePrefix.ZCtx("n1"))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if signedBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected signedBigInt= %v .\n"+
			"Instead got, %v\n",
			expected.String(), signedBigInt.String())
	}

}

func TestNumStrDto_NewBigInt_10(t *testing.T) {

	ePrefix := ErrPrefixDto{}.NewEPrefOld("TestNumStrDto_NewBigInt_10()")

	signedAbsNumStr := "-123456789"
	absAllNumStr := "123456789"
	nStr := "-123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := -1
	sBigInt, isOk :=
		big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("bigInt.SetString(signedAbsNumStr,10) conversion failed!\n"+
			"signedAbsNumStr= '%v'\n", signedAbsNumStr)
		return
	}

	var n1 NumStrDto
	var err error

	n1, err = NumStrDto{}.NewBigInt(
		sBigInt,
		precision,
		ePrefix.ZCtx(
			fmt.Sprintf("sBigInt='%v' ",
				sBigInt.Text(10))))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var nDto NumStrDto

	nDto,
		err = n1.CopyOut(
		ePrefix.ZCtxEmpty())

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix.ZCtx("nDto"))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if nStr != actualNumStr {
		t.Errorf("Expected Number String = '%v'.\n"+
			"Instead, got Number String = '%v'\n",
			nStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != actualNumStr {
		t.Errorf("Expected AbsAllRunes = '%v'.\n"+
			" Instead, got %v\n",
			absAllNumStr, actualNumStr)
		return

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v", iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'\n"+
			"Instead, got %v\n", nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix.ZCtx("Testing 'nDto' validity"))

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError()\n"+
			"Error='%v'\n", err.Error())
		return
	}

}

func TestNumStrDto_NewBigInt_20(t *testing.T) {

	ePrefix := ErrPrefixDto{}.NewEPrefOld(
		"TestNumStrDto_NewBigInt_20()")

	signedAbsNumStr := "123456789"
	absAllNumStr := "123456789"
	nStr := "123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := 1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("bigInt.SetString(signedAbsNumStr,10) conversion failed!\n"+
			"signedAbsNumStr= '%v'\n", signedAbsNumStr)
		return
	}

	var n1, nDto NumStrDto
	var err error

	n1, err = NumStrDto{}.NewBigInt(
		sBigInt,
		precision,
		ePrefix.ZCtx(
			fmt.Sprintf("sBigInt='%v' ",
				sBigInt.Text(10))))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	nDto,
		err = n1.CopyOut(
		ePrefix.ZCtx(
			"n1->nDto"))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(ePrefix.ZCtx("nDto"))

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if nStr != actualNumStr {
		t.Errorf("Expected Number String = '%v'.\n"+
			"Instead, got Number String = '%v'\n",
			nStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != actualNumStr {
		t.Errorf("Expected AbsAllRunes = '%v'.\n"+
			" Instead, got %v\n",
			absAllNumStr, actualNumStr)
		return

	}

	actualNumStr = string(nDto.GetAbsIntRunes())

	if iStr != actualNumStr {
		t.Errorf("Expected AbsIntRunes = '%v'.\n"+
			"Instead, got %v", iStr, actualNumStr)
		return
	}

	actualNumStr = string(nDto.GetAbsFracRunes())

	if fracStr != actualNumStr {
		t.Errorf("Expected AbsFracRunes = '%v'.\n"+
			"Instead, got %v\n", fracStr, actualNumStr)
		return
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'.\n"+
			"Instead, got %v\n", signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got %v\n", nDto.HasNumericDigits())
		return
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'\n"+
			"Instead, got %v\n", nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'.\n"+
			"Instead, got %v\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix.ZCtx("Testing 'nDto' validity"))

	if err != nil {
		t.Errorf("%v", err)
	}

}
