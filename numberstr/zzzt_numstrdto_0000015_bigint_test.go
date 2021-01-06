package numberstr

import (
	"math/big"
	"testing"
)

func TestNumStrDto_GetBigIntNum_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetBigIntNum_10() "
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
		ePrefix+"bigI ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		actualBigIntNDto.GetNumStr(ePrefix + "actualBigIntNDto ")

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
		err = actualBigIntNDto.GetBigInt(ePrefix + "actualBigIntNDto ")

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

	ePrefix := "TestNumStrDto_GetSignedBigInt_10"

	var n1, nXDto NumStrDto
	var err error

	nXDto = NumStrDto{}.New()

	n1, err =
		nXDto.ParseNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var signedBigInt *big.Int

	signedBigInt, err = n1.GetBigInt(
		ePrefix + "n1 ")

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
