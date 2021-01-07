package numberstr

import "testing"

func TestNumStrDto_GetNumParen_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetNumParen_10() "
	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123456.97"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_20(t *testing.T) {

	ePrefix := "TestNumStrDto_GetNumParen_20() "
	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_30(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.29"

	ePrefix := "TestNumStrDto_GetNumParen_30() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_40(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345"

	ePrefix := "TestNumStrDto_GetNumParen_40() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}
}

func TestNumStrDto_GetNumParen_50(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.1234"

	ePrefix := "TestNumStrDto_GetNumParen_50() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}
}

func TestNumStrDto_GetNumParen_60(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1234567890.25"

	ePrefix := "TestNumStrDto_GetNumParen_60() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}
}

func TestNumStrDto_GetNumParen_70(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12345.29)"

	ePrefix := "TestNumStrDto_GetNumParen_70() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}
}

func TestNumStrDto_GetNumParen_80(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12345)"

	ePrefix := "TestNumStrDto_GetNumParen_80() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_90(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(123)"

	ePrefix := "TestNumStrDto_GetNumParen_90() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_100(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(0.123)"

	ePrefix := "TestNumStrDto_GetNumParen_100() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_110(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(1234567890.123)"

	ePrefix := "TestNumStrDto_GetNumParen_110() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetNumParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Number Prenthesis Str='%v'.\n"+
			"Instead, Number Prenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0010(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123,456.97"

	ePrefix := "TestNumStrDto_GetThouParen_0010() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0020(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	ePrefix := "TestNumStrDto_GetThouParen_0020() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0030(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345.29"

	ePrefix := "TestNumStrDto_GetThouParen_0030() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0040(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345"

	ePrefix := "TestNumStrDto_GetThouParen_0040() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0050(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12,345.1234"

	ePrefix := "TestNumStrDto_GetThouParen_0050() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0060(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1,234,567,890.25"

	ePrefix := "TestNumStrDto_GetThouParen_0060() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0070(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12,345.29)"

	ePrefix := "TestNumStrDto_GetThouParen_0070() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0080(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12,345)"

	ePrefix := "TestNumStrDto_GetThouParen_0080() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0090(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(123)"

	ePrefix := "TestNumStrDto_GetThouParen_0090() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0100(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(0.123)"

	ePrefix := "TestNumStrDto_GetThouParen_0100() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetThouParen_0110(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(1,234,567,890.123)"

	ePrefix := "TestNumStrDto_GetThouParen_0110() "

	var nDto NumStrDto
	var err error

	nDto, err =
		NumStrDto{}.NewNumStr(
			nStr,
			ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v'.\n",
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

	var actualStr string

	actualStr, err = nDto.GetThouParen(
		ePrefix + "nDto ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("Expected Thousands Parenthesis Str='%v'.\n"+
			"Instead, Thousands Parenthesis Str='%v'\n",
			expectedStr, actualStr)
	}

}
