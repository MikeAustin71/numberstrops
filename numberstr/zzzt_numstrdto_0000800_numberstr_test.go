package numberstr

import "testing"

func TestNumStrDto_GetNumStr_010(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "123456.97"

	ePrefix := "TestNumStrDto_GetNumStr_010() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_020(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "123.45"

	ePrefix := "TestNumStrDto_GetNumStr_020() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}
}

func TestNumStrDto_GetNumStr_030(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "12345.29"

	ePrefix := "TestNumStrDto_GetNumStr_030() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_040(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "12345"

	ePrefix := "TestNumStrDto_GetNumStr_040() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_050(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "12345.1234"

	ePrefix := "TestNumStrDto_GetNumStr_050() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_060(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "1234567890.25"

	ePrefix := "TestNumStrDto_GetNumStr_060() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_070(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "-12345.29"

	ePrefix := "TestNumStrDto_GetNumStr_070() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_080(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "-12345"

	ePrefix := "TestNumStrDto_GetNumStr_080() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_090(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "-123"

	ePrefix := "TestNumStrDto_GetNumStr_090() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_100(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "-0.123"

	ePrefix := "TestNumStrDto_GetNumStr_100() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}

func TestNumStrDto_GetNumStr_110(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedNumStr := "-1234567890.123"

	ePrefix := "TestNumStrDto_GetNumStr_110() "

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v' .\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v' .\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualNumStr string

	actualNumStr, err = nDto.GetNumStr(
		ePrefix + "nDto ")

	if expectedNumStr != actualNumStr {
		t.Errorf("Expected Number String='%v'.\n"+
			"Instead, Number String='%v'\n",
			expectedNumStr, actualNumStr)
	}

}
