package numberstr

import "testing"

func TestNumStrDto_GetCurrencyStr_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_10() "
	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$123,456.97"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_20(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_20() "
	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$123.45"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_30(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_30() "
	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$12,345.29"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_40(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_40() "
	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$12,345"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_50(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_50() "
	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$12,345.1234"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_60(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_60() "
	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "-$12,345.29"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_70(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_70() "
	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "-$12,345"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}
}

func TestNumStrDto_GetCurrencyStr_80(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_80() "
	nStr := "-12345"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "-€12,345"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyStr_90(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyStr_90() "
	nStr := "12345.12"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "€12,345.12"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_10() "
	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$123,456.97"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_20(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_20() "
	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$123.45"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_30(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_30() "
	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$12,345.29"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_40(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_40() "
	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "$12,345"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_50(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_50() "
	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "($12,345.29)"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_60(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_60() "
	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "($12,345)"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_70(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_70() "
	nStr := "-12345"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "(€12,345)"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}

func TestNumStrDto_GetCurrencyParen_80(t *testing.T) {

	ePrefix := "TestNumStrDto_GetCurrencyParen_80() "
	nStr := "12345.12"
	// '\U000020ac', // Euro €
	currencySymbol := '€'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedCurrencyStr := "€12,345.12"

	var nDto NumStrDto
	var err error

	nDto, err = NumStrDto{}.NewNumStr(
		nStr,
		ePrefix+"nStr ")

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	nDto.SetNumericSeparators(
		decimalSeparator,
		thousandsSeparator,
		currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.\n"+
			"Instead, Currency Symbol='%v' .\n",
			currencySymbol, nDto.GetCurrencySymbol())
		return
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.\n"+
			"Instead, Decimal Separator='%v'.\n",
			decimalSeparator, nDto.GetDecimalSeparator())
		return
	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.\n"+
			"Instead, Thousands Separator='%v'.\n",
			thousandsSeparator, nDto.GetThousandsSeparator())
		return
	}

	var actualCurrencyStr string

	actualCurrencyStr, err =
		nDto.GetCurrencyStr(ePrefix + "nDto ")

	if expectedCurrencyStr != actualCurrencyStr {
		t.Errorf("Expected Currency Str='%v'.\n"+
			"Instead, Currency Str='%v'\n",
			expectedCurrencyStr, actualCurrencyStr)
		return
	}

}
