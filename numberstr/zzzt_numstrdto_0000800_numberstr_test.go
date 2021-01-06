package numberstr

import "testing"

func TestNumStrDto_GetNumStr_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.1234"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_06(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1234567890.25"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_07(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_08(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_09(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_10(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-0.123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_11(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-1234567890.123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator())

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator())

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}
