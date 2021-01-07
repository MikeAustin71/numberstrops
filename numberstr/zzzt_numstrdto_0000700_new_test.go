package numberstr

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
)

func TestNumStrDto_NewBigInt_0010(t *testing.T) {
	expectedNumStr := "123.456"
	ePrefix := "TestNumStrDto_NewBigInt_0010() "

	intNumStr := "123456"

	bigIntNum, oK := big.NewInt(0).
		SetString(intNumStr, 10)

	if !oK {
		t.Errorf(ePrefix+"\n"+
			"bigIntNum conversion failed!\n"+
			"big.NewInt(0).SetString(\"%v\", 10)\n",
			intNumStr)
		return
	}

	nDto, err := NumStrDto{}.NewBigInt(
		bigIntNum,
		3,
		ePrefix+
			fmt.Sprintf("bigIntNum=%v ",
				bigIntNum.Text(10)))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. \n"+
			"Instead, nDto.GetNumStr()='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewBigInt_0020(t *testing.T) {
	expectedNumStr := "-123.456"
	ePrefix := "TestNumStrDto_NewBigInt_0020() "

	intNumStr := "123456"

	bigIntNum, oK := big.NewInt(0).
		SetString(intNumStr, 10)

	if !oK {
		t.Errorf(ePrefix+"\n"+
			"bigIntNum conversion failed!\n"+
			"big.NewInt(0).SetString(\"%v\", 10)\n",
			intNumStr)
		return
	}

	nDto, err := NumStrDto{}.NewBigInt(
		bigIntNum,
		3,
		ePrefix+
			fmt.Sprintf("bigIntNum=%v ",
				bigIntNum.Text(10)))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. \n"+
			"Instead, nDto.GetNumStr()='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt_0010() "

	intNum := 7
	precision := uint(0)

	expectedNumStr := "7"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt(
			intNum,
			precision,
			ePrefix+"intNum ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt_0020() "

	intNum := 7
	precision := uint(1)

	expectedNumStr := "7.0"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt(
			intNum,
			precision,
			ePrefix+"intNum ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt_0030() "

	intNum := 7
	precision := uint(3)

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt(
			intNum,
			precision,
			ePrefix+"intNum ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt_0040() "

	intNum := -7
	precision := uint(3)

	expectedNumStr := "-7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt(
			intNum,
			precision,
			ePrefix+"intNum ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt_0050() "

	intNum := -7
	precision := uint(0)

	expectedNumStr := "-7"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt(
			intNum,
			precision,
			ePrefix+"intNum ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewIntExponent_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewIntExponent_0010() "

	intNum := 7
	exponent := 3

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewIntExponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewIntExponent_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewIntExponent_0020() "

	intNum := 7123
	exponent := -3

	expectedNumStr := "7.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewIntExponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewIntExponent_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewIntExponent_0030() "

	intNum := -72
	exponent := 3

	expectedNumStr := "-72.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewIntExponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewIntExponent_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewIntExponent_0040() "

	intNum := -72123
	exponent := -3

	expectedNumStr := "-72.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewIntExponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewIntExponent_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewIntExponent_0050() "

	intNum := 72
	exponent := 0

	expectedNumStr := "72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewIntExponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewIntExponent_0060(t *testing.T) {

	ePrefix := "TestNumStrDto_NewIntExponent_0060() "

	intNum := -72
	exponent := 0

	expectedNumStr := "-72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewIntExponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32_0010() "

	intNum := int32(7)
	precision := uint(0)

	expectedNumStr := "7"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32_0020() "

	intNum := int32(7)
	precision := uint(1)

	expectedNumStr := "7.0"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32_0030() "

	intNum := int32(7)
	precision := uint(3)

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32_0040() "

	intNum := int32(-7)
	precision := uint(3)

	expectedNumStr := "-7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32_0050() "

	intNum := int32(-7)
	precision := uint(0)

	expectedNumStr := "-7"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32Exponent_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32Exponent_0010() "

	intNum := int32(7)
	exponent := 3

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32Exponent_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32Exponent_0020() "

	intNum := int32(7123)
	exponent := -3

	expectedNumStr := "7.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32Exponent_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32Exponent_0030() "

	intNum := int32(-72)
	exponent := 3

	expectedNumStr := "-72.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32Exponent_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32Exponent_0040() "

	intNum := int32(-72123)
	exponent := -3

	expectedNumStr := "-72.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32Exponent_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32Exponent_0050() "

	intNum := int32(72)
	exponent := 0

	expectedNumStr := "72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt32Exponent_0060(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt32Exponent_0060() "

	intNum := int32(-72)
	exponent := 0

	expectedNumStr := "-72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt32Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64_0010() "

	intNum := int64(7)
	precision := uint(0)

	expectedNumStr := "7"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64_0020() "

	intNum := int64(7)
	precision := uint(1)

	expectedNumStr := "7.0"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64_0030() "

	intNum := int64(7)
	precision := uint(3)

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64_0040() "

	intNum := int64(-7)
	precision := uint(3)

	expectedNumStr := "-7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64_0050() "

	intNum := int64(-7)
	precision := uint(0)

	expectedNumStr := "-7"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64(
			intNum,
			precision,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64Exponent_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64Exponent_0010() "

	intNum := int64(7)
	exponent := 3

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64Exponent_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64Exponent_0020() "

	intNum := int64(7123)
	exponent := -3

	expectedNumStr := "7.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64Exponent_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64Exponent_0030() "

	intNum := int64(-72)
	exponent := 3

	expectedNumStr := "-72.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64Exponent_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64Exponent_0040() "

	intNum := int64(-72123)
	exponent := -3

	expectedNumStr := "-72.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64Exponent_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64Exponent_0050() "

	intNum := int64(72)
	exponent := 0

	expectedNumStr := "72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewInt64Exponent_0060(t *testing.T) {

	ePrefix := "TestNumStrDto_NewInt64Exponent_0060() "

	intNum := int64(-72)
	exponent := 0

	expectedNumStr := "-72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewInt64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewNumStr_0010(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	ePrefix := "TestNumStrDto_NewNumStr_0010() "

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+
				fmt.Sprintf("nStr=%v", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
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
		t.Errorf("Expected nDto.GetSign()= '%v'.\n"+
			"Instead, gotnDto.GetSign()= '%v'\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got nDto.HasNumericDigist= '%v'\n",
			nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected nDto.IsFractionalValue= 'true'.\n"+
			"Instead, got nDto.IsFractionalValue= '%v'.\n",
			nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected nDto.precision= '%v'.\n"+
			" Instead, got nDto.precision= '%v'\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(ePrefix +
		"Testing 'nDto' validity. ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError()\n"+
			"Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_0020(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	ePrefix := "TestNumStrDto_NewNumStr_0020() "

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+
				fmt.Sprintf("nStr=%v", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
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
		t.Errorf("Expected nDto.GetSign()= '%v'.\n"+
			"Instead, gotnDto.GetSign()= '%v'\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got nDto.HasNumericDigist= '%v'\n",
			nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected nDto.IsFractionalValue= 'true'.\n"+
			"Instead, got nDto.IsFractionalValue= '%v'.\n",
			nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected nDto.precision= '%v'.\n"+
			" Instead, got nDto.precision= '%v'\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity. ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError()\n"+
			"Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_0030(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	ePrefix := "TestNumStrDto_NewNumStr_0030() "

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+
				fmt.Sprintf("nStr=%v", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
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
		t.Errorf("Expected nDto.GetSign()= '%v'.\n"+
			"Instead, gotnDto.GetSign()= '%v'\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got nDto.HasNumericDigist= '%v'\n",
			nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected nDto.IsFractionalValue= 'true'.\n"+
			"Instead, got nDto.IsFractionalValue= '%v'.\n",
			nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected nDto.precision= '%v'.\n"+
			" Instead, got nDto.precision= '%v'\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity. ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError()\n"+
			"Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_0040(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	ePrefix := "TestNumStrDto_NewNumStr_0040() "

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+
				fmt.Sprintf("nStr=%v", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
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
		t.Errorf("Expected nDto.GetSign()= '%v'.\n"+
			"Instead, gotnDto.GetSign()= '%v'\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got nDto.HasNumericDigist= '%v'\n",
			nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected nDto.IsFractionalValue= 'true'.\n"+
			"Instead, got nDto.IsFractionalValue= '%v'.\n",
			nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected nDto.precision= '%v'.\n"+
			" Instead, got nDto.precision= '%v'\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity. ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError()\n"+
			"Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_0050(t *testing.T) {
	nStr := "-000.000"
	expectedNumStr := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	ePrefix := "TestNumStrDto_NewNumStr_0050() "

	isFractionalValue := true

	if !strings.Contains(nStr, ".") {
		isFractionalValue = false
	}

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+
				fmt.Sprintf("nStr=%v", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected NumStrOut = '%v'.\n"+
			"Instead, got %v\n",
			expectedNumStr, actualNumStr)
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
		t.Errorf("Expected nDto.GetSign()= '%v'.\n"+
			"Instead, gotnDto.GetSign()= '%v'\n",
			signVal, nDto.GetSign())
		return
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected nDto.HasNumericDigist= 'true'.\n"+
			"Instead, got nDto.HasNumericDigist= '%v'\n",
			nDto.HasNumericDigits())
		return
	}

	if isFractionalValue != nDto.IsFractionalValue() {
		t.Errorf("Expected nDto.IsFractionalValue= 'true'.\n"+
			"Instead, got nDto.IsFractionalValue= '%v'.\n",
			nDto.IsFractionalValue())
		return
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected nDto.precision= '%v'.\n"+
			" Instead, got nDto.precision= '%v'\n",
			precision, nDto.GetPrecisionUint())
		return
	}

	err = nDto.IsValidInstanceError(
		ePrefix + "Testing 'nDto' validity. ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValidInstanceError()\n"+
			"Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStrWithNumSeps_0010() "

	nStr := "123,456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStrWithNumSeps(
		nStr,
		expectedNumSeps,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'.\n",
			nStr, actualNumStr)
		return
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Number Separators='%v'.\n"+
			"Instead, Number Separators='%v'.\n",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStrWithNumSeps_0020() "

	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.SetToUSADefaultsIfEmpty()

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStrWithNumSeps(
		nStr,
		expectedNumSeps,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'.\n",
			nStr, actualNumStr)
		return
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Number Separators='%v'.\n"+
			"Instead, Number Separators='%v'.\n",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewNumStrWithNumSeps_0020() "

	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.SetToUSADefaultsIfEmpty()

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStrWithNumSeps(
		nStr,
		expectedNumSeps,
		ePrefix+
			fmt.Sprintf("nStr='%v' ", nStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.\n"+
			"Instead, NumStr='%v'.\n",
			nStr, actualNumStr)
		return
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected Number Separators='%v'.\n"+
			"Instead, Number Separators='%v'.\n",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestNumStrDto_NewUint_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint_0010() "

	intNum := uint(7)
	precision := uint(0)

	expectedNumStr := "7"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint_0020() "

	intNum := uint(7)
	precision := uint(1)

	expectedNumStr := "7.0"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint_0030() "

	intNum := uint(7)
	precision := uint(3)

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint_0040() "

	intNum := uint(792)
	precision := uint(3)

	expectedNumStr := "792.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint_0050() "

	intNum := uint(792)
	precision := uint(0)

	expectedNumStr := "792"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewUintExponent_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0010() "

	intNum := uint(7)
	exponent := 3

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUintExponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUintExponent_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0020() "

	intNum := uint(7123)
	exponent := -3

	expectedNumStr := "7.123"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUintExponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUintExponent_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0030() "

	intNum := uint(872)
	exponent := 3

	expectedNumStr := "872.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUintExponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUintExponent_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0040() "

	intNum := uint(872123)
	exponent := -3

	expectedNumStr := "872.123"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUintExponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUintExponent_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0050() "

	intNum := uint(72)
	exponent := 0

	expectedNumStr := "72"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUintExponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUintExponent_0060(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0060() "

	intNum := uint(472)
	exponent := 0

	expectedNumStr := "472"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUintExponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint32_0010() "

	intNum := uint32(7)
	precision := uint(0)

	expectedNumStr := "7"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint32_0020() "

	intNum := uint32(7)
	precision := uint(1)

	expectedNumStr := "7.0"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint32_0030() "

	intNum := uint32(7)
	precision := uint(3)

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint32_0040() "

	intNum := uint32(792)
	precision := uint(3)

	expectedNumStr := "792.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint32_0050() "

	intNum := uint32(792)
	precision := uint(0)

	expectedNumStr := "792"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32Exponent_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0010() "

	intNum := uint32(7)
	exponent := 3

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32Exponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32Exponent_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0020() "

	intNum := uint32(7123)
	exponent := -3

	expectedNumStr := "7.123"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32Exponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32Exponent_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0030() "

	intNum := uint32(872)
	exponent := 3

	expectedNumStr := "872.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32Exponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32Exponent_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0040() "

	intNum := uint32(872123)
	exponent := -3

	expectedNumStr := "872.123"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32Exponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32Exponent_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0050() "

	intNum := uint32(72)
	exponent := 0

	expectedNumStr := "72"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32Exponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint32Exponent_0060(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUintExponent_0060() "

	intNum := uint32(472)
	exponent := 0

	expectedNumStr := "472"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint32Exponent(
		intNum,
		exponent,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint64_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64_0010() "

	intNum := uint64(7)
	precision := uint(0)

	expectedNumStr := "7"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint64(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint64_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64_0020() "

	intNum := uint64(7)
	precision := uint(1)

	expectedNumStr := "7.0"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint64(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint64_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64_0030() "

	intNum := uint64(7)
	precision := uint(3)

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint64(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint64_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64_0040() "

	intNum := uint64(792)
	precision := uint(3)

	expectedNumStr := "792.000"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint64(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint64_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64_0050() "

	intNum := uint64(792)
	precision := uint(0)

	expectedNumStr := "792"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewUint64(
		intNum,
		precision,
		ePrefix+
			fmt.Sprintf("intNum='%v' ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
		return
	}

}

func TestNumStrDto_NewUint64Exponent_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64Exponent_0010() "

	intNum := uint64(7)
	exponent := 3

	expectedNumStr := "7.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewUint64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewUint64Exponent_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64Exponent_0020() "

	intNum := uint64(7123)
	exponent := -3

	expectedNumStr := "7.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewUint64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewUint64Exponent_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64Exponent_0030() "

	intNum := uint64(872)
	exponent := 3

	expectedNumStr := "872.000"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewUint64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewUint64Exponent_0040(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64Exponent_0040() "

	intNum := uint64(872123)
	exponent := -3

	expectedNumStr := "872.123"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewUint64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewUint64Exponent_0050(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64Exponent_0050() "

	intNum := uint64(72)
	exponent := 0

	expectedNumStr := "72"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewUint64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewUint64Exponent_0060(t *testing.T) {

	ePrefix := "TestNumStrDto_NewUint64Exponent_0060() "

	intNum := uint64(472)
	exponent := 0

	expectedNumStr := "472"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewUint64Exponent(
			intNum,
			exponent,
			ePrefix+
				fmt.Sprintf("intNum=%v ", intNum))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewZero_0010(t *testing.T) {

	ePrefix := "TestNumStrDto_NewZero_0010() "

	expectedNumStr := "0"

	nDto := NumStrDto{}.NewZero(0)

	var actualNumStr string
	var err error
	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewZero_0020(t *testing.T) {

	ePrefix := "TestNumStrDto_NewZero_0020() "

	expectedNumStr := "0.00"

	nDto := NumStrDto{}.NewZero(2)

	var actualNumStr string
	var err error
	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_NewZero_0030(t *testing.T) {

	ePrefix := "TestNumStrDto_NewZero_0030() "

	expectedNumStr := "0.0000"

	nDto := NumStrDto{}.NewZero(4)

	var actualNumStr string
	var err error
	actualNumStr, err =
		nDto.GetNumStr(
			ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'.\n"+
			"Instead, nDto.GetNumStr()='%v'.\n",
			expectedNumStr, actualNumStr)
	}

}
