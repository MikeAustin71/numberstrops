package numberstr

import (
	"fmt"
	"testing"
)

/*
 These tests are associated with type NumStrBasicUtility
 located in source file numstrbasicutility.go.

*/

func TestNumStrBasicUtility_DlimDecStr(t *testing.T) {

	ns := NumStrBasicUtility{}
	n := "1234567890"
	expected := "1,234,567,890"

	result := ns.DlimCurrencyStr(n, ',', '.', '$')

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}

}

func TestNumStrBasicUtility_DlimDecStr_With_DecimalCurrency(t *testing.T) {
	ns := NumStrBasicUtility{}
	n := "$1234567890.25"

	result := ns.DlimCurrencyStr(n, ',', '.', '$')
	expected := "$1,234,567,890.25"

	if result != expected {
		t.Errorf("Expected result = %v\n"+
			"Instead got: %v", expected, result)
	}

}

func TestNumStrBasicUtility_DNumI64(t *testing.T) {
	ns := NumStrBasicUtility{}
	n := int64(1234567890)
	expected := "1,234,567,890"

	result := ns.DLimI64(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v\n"+
			"Instead got: %v", expected, result)
	}

}

func TestNumStrBasicUtility_DNumI64_EvenThousands(t *testing.T) {
	ns := NumStrBasicUtility{}
	n := int64(123456)
	expected := "123,456"

	result := ns.DLimI64(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v\n"+
			"Instead got: %v", expected, result)
	}

}

func TestNumStrBasicUtility_DLimInt(t *testing.T) {
	ns := NumStrBasicUtility{}
	n := 1234567
	expected := "1,234,567"

	result := ns.DLimInt(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v\n"+
			"Instead got: %v", expected, result)
	}
}

func TestNumStrBasicUtility_ConvertStrToInt64_01(t *testing.T) {
	numStr := "-12314617914"

	ns := NumStrBasicUtility{}

	result, err := ns.ConvertNumStrToInt64(numStr)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertNumStrToInt64(%v).\n"+
			"Error='%v'\n", numStr, err.Error())
		return
	}

	var s string
	s, err = ns.ConvertInt64ToStr(result)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertInt64ToStr(%v)\n"+
			"Error='%v'\n", result, err.Error())
		return
	}

	if s != numStr {
		t.Errorf("Error: Expected value '%v'.\n"+
			"Instead, got '%v'", numStr, s)
	}

}

func TestNumStrBasicUtility_ConvertStrToInt64_02(t *testing.T) {
	numStr := "+12314617914"

	ns := NumStrBasicUtility{}

	result, err := ns.ConvertNumStrToInt64(numStr)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertNumStrToInt64(%v).\n"+
			"Error='%v'\n", numStr, err.Error())
		return
	}

	var s string

	s, err = ns.ConvertInt64ToStr(result)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertInt64ToStr(%v)\n"+
			"Error='%v'\n",
			result, err.Error())
		return
	}

	s = "+" + s

	if s != numStr {
		t.Errorf("Error. Expected value %v.\n"+
			"Instead, got %v\n", numStr, s)
	}
}

func TestNumStrBasicUtility_ConvertStrToInt64_03(t *testing.T) {

	numStr := "12314617914"

	ns := NumStrBasicUtility{}

	result, err := ns.ConvertNumStrToInt64(numStr)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertNumStrToInt64(%v).\n"+
			"Error='%v'\n", numStr, err.Error())
		return
	}

	var s string

	s, err = ns.ConvertInt64ToStr(result)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertInt64ToStr(%v)\n"+
			"Error='%v'\n", result, err.Error())
		return
	}

	if s != numStr {
		t.Errorf("Error. Expected value %v.\n"+
			"Instead, got %v\n", numStr, s)
	}
}

func TestNumStrBasicUtility_ConvertStrToIntNumStr(t *testing.T) {
	str := "-12,314,617,914"
	expected := "-12314617914"
	ns := NumStrBasicUtility{}
	ns.ThousandsSeparator = ','

	ePrefix := "TestNumStrBasicUtility_ConvertStrToIntNumStr()"

	var err error

	var numStr string
	numStr, err = ns.ConvertStrToIntNumStr(
		str,
		ePrefix+
			fmt.Sprintf("str='%v' ", str))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expected != numStr {
		t.Errorf("Expected number string %v.\n"+
			"Instead got %v", expected, numStr)
	}

}

func TestNumStrBasicUtility_SetCountryAndCurrency(t *testing.T) {

	ns := NumStrBasicUtility{}

	err := ns.SetCountryAndCurrency("United States")

	if err != nil {
		t.Errorf("Received Error from ns.SetCountryAndCurrency(\"United States\").\n"+
			"Error: %v\n", err.Error())
		return
	}

	if ns.CurrencySymbol != '$' {
		t.Errorf("Expected '$' currency symbol,\n"+
			"Instead got %v.\n", ns.CurrencySymbol)
		return
	}

	if ns.Nation != "United States" {
		t.Errorf("Expected Nation = 'United States'.\n"+
			"Instead, got %v.\n", ns.Nation)
		return
	}

	if ns.ThousandsSeparator != ',' {
		t.Errorf("Expected thousands separator = ','.\n"+
			"Instead, got %v.\n", ns.ThousandsSeparator)
		return
	}

	if ns.DecimalSeparator != '.' {
		t.Errorf("Expected decimal separator = ','.\n"+
			"Instead, got %v.\n", ns.DecimalSeparator)
	}

}

func TestNumStrBasicUtility_ConvertInt64ToFractionalValue(t *testing.T) {

	ns := NumStrBasicUtility{}

	i64 := int64(123456)

	f64, err := ns.ConvertInt64ToFractionalValue(i64)

	if err != nil {
		t.Errorf("Error received from ns.ConvertInt64ToFractionalValue(i64).\n"+
			"Error: %v\n", err.Error())
		return
	}

	result := fmt.Sprintf("%v", f64)
	expected := "0.123456"

	if result != expected {
		t.Errorf("Error on ConvertInt64ToFractionalValue(). Expected '%v'.\n"+
			"Instead, got '%v'.", expected, result)
	}
}

func TestNumStrBasicUtility_ParseNumString_01(t *testing.T) {

	ePrefix :=
		ErrPrefixDto{}.NewEPrefOld(
			"TestNumStrBasicUtility_ParseNumString_01()")

	rawStr := "123456.654321"
	expected := "123456.654321"
	ns := NumStrBasicUtility{}

	nsDto, err := ns.ParseNumString(
		rawStr,
		ePrefix.ZCtx(
			fmt.Sprintf("rawStr='%v' ", rawStr)))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr, err = nsDto.GetNumStr(
		ePrefix + "nsDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != expected {
		t.Errorf("Expected Actual NumStr= '%v'.\n"+
			"Instead, got %v\n",
			expected, actualNumStr)
		return
	}

	if nsDto.GetSign() != 1 {
		t.Errorf("Expected signVal=1. Instead, got %v", nsDto.GetSign())
		return
	}

	if nsDto.GetPrecisionInt() != 6 {
		t.Errorf("Expected precision=6. Instead, got %v", nsDto.GetPrecisionInt())
		return
	}

	if nsDto.HasNumericDigits() == false {
		t.Errorf("Expected HasNumericDigits=true. Instead, got %v", nsDto.HasNumericDigits())
		return
	}

	if nsDto.IsFractionalValue() == false {
		t.Errorf("Expected IsFractionalValue=true. Instead, got %v", nsDto.IsFractionalValue())
	}

}

func TestNumStrBasicUtility_ParseNumString_02(t *testing.T) {

	ePrefix := "TestNumStrBasicUtility_ParseNumString_02() "
	rawStr := "-123456.654321"
	expected := "-123456.654321"
	ns := NumStrBasicUtility{}

	nsDto, err := ns.ParseNumString(
		rawStr,
		ePrefix+
			fmt.Sprintf("rawStr='%v' ", rawStr))

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

	if actualNumStr != expected {
		t.Errorf("Expected NumStr= '%v'.\n"+
			"Instead, got %v\n", expected, actualNumStr)
		return
	}

	if nsDto.GetSign() != -1 {
		t.Errorf("Expected signVal= -1.\n"+
			"Instead, got %v\n", nsDto.GetSign())
		return
	}

	if nsDto.GetPrecisionInt() != 6 {
		t.Errorf("Expected precision=6.\n"+
			"Instead, got %v\n", nsDto.GetPrecisionInt())
		return
	}

	if nsDto.HasNumericDigits() == false {
		t.Errorf("Expected HasNumericDigits=true.\n"+
			"Instead, got %v\n", nsDto.HasNumericDigits())
		return
	}

	if nsDto.IsFractionalValue() == false {
		t.Errorf("Expected IsFractionalValue=true.\n"+
			"Instead, got %v\n", nsDto.IsFractionalValue())
	}
}

func TestNumStrBasicUtility_ParseNumString_03(t *testing.T) {

	ePrefix := "TestNumStrBasicUtility_ParseNumString_03() "
	rawStr := "Nothing"
	expected := "0"
	ns := NumStrBasicUtility{}

	nsDto, err := ns.ParseNumString(
		rawStr,
		ePrefix+
			fmt.Sprintf("rawStr='%v' ", rawStr))

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = nsDto.GetNumStr(ePrefix + "nsDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if actualNumStr != expected {
		t.Errorf("Expected NumStrOut= '%v'.\n"+
			"Instead, got %v\n",
			expected, actualNumStr)
		return
	}

	if nsDto.HasNumericDigits() == false {
		t.Errorf("Expected HasNumericDigits=true.\n"+
			"Instead, got %v\n", nsDto.HasNumericDigits())
	}

}
