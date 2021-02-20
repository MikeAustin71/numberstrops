package numberstr

import (
	"fmt"
	"strconv"
	"sync"
)

/*
 NumStrBasicUtility
 ==================

 Type 'NumStrBasicUtility' provides a basic set of numeric
 conversion and management routines primarily focused on
 number strings. It is now deprecated and superseded by
 type 'NumStrDto' which provides many additional capabilities.

*/

type NumStrBasicUtility struct {
	NumStrFormatSpec NumStrFmtSpecDto
	StrIn            string
	StrOut           string
	IsFractionalVal  bool
	IntegerStr       string
	FractionStr      string
	Int64Val         int64
	Float64Val       float64
	lock             *sync.Mutex
}

// DLimInt - Receives an integer and returns a delimited number
// string with thousands separator (i.e. 1000000 -> 1,000,000).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  num                            int
//     - The integer value which will be converted to a number
//       string and inserted with thousands separators.
//
//
//  integerDigitsSeparator         rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  ePrefix                        ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  numStr                         string
//     - If this method completes successfully, this returned
//       string will contain numeric digits separated into thousands
//       by the delimiter character supplied in input parameter,
//       'integerDigitsSeparator'.
//
//  err                            error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (ns NumStrBasicUtility) DLimInt(
	num int,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix ErrPrefixDto) (
	numStr string,
	err error) {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	ePrefix.SetEPref("NumStrBasicUtility.DLimInt() ")

	pureNumStr := strconv.Itoa(num)

	numStr,
		err = numStrBasicAtom{}.ptr().
		delimitIntSeparators(
			pureNumStr,
			integerDigitsSeparator,
			integerDigitsGroupingSequence,
			ePrefix.XCtx(
				fmt.Sprintf("pureNumStr='%v'", pureNumStr)))

	return numStr, err
}

// DLimI64 - Receives an int64 and returns a delimited number
// string with thousands separator (i.e. 1000000 -> 1,000,000).
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  num                            int64
//     - The integer value which will be converted to a number
//       string and inserted with thousands separators.
//
//
//  integerDigitsSeparator         rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  ePrefix                        ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  numStr                         string
//     - If this method completes successfully, this returned
//       string will contain numeric digits separated into thousands
//       by the delimiter character supplied in input parameter,
//       'integerDigitsSeparator'.
//
//  err                            error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (ns NumStrBasicUtility) DLimI64(
	num int64,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix ErrPrefixDto) (
	numStr string,
	err error) {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	ePrefix.SetEPref("NumStrBasicUtility.DLimI64() ")

	pureNumStr := strconv.FormatInt(num, 10)

	numStr,
		err = numStrBasicAtom{}.ptr().
		delimitIntSeparators(
			pureNumStr,
			integerDigitsSeparator,
			integerDigitsGroupingSequence,
			ePrefix.XCtx(
				fmt.Sprintf("pureNumStr='%v'", pureNumStr)))

	return numStr, err
}

// DlimCurrencyStr - Converts a number string to a formatted
// currency string by inserting a Currency Symbol and a
// Thousands Separator. This method processes both integer
// and floating point numeric values.
//
// Examples:
//   rawNumStr == 1234567890 -> Output == "$1,234,567,890".
//   rawNumStr == 1234567.89 -> Output == "$1,234,567.89".
//   rawNumStr == -1234567890 -> Output == "$-1,234,567,890".
//   rawNumStr == -1234567.89 -> Output == "$-1,234,567.89".
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  rawNumStr                      string
//     - A number string which will be injected with thousands
//       separators, prefixed with a currency symbol and
//       correctly formatted for any included fractional digits.
//
//
//  integerDigitsSeparator         rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  decimalSeparator               rune
//     - A unicode character inserted into a number string to
//       separate integer and fractional digits. In the United
//       States, the decimal separator is the period character
//       ('.') and it is known as the decimal point.
//
//
//  currencySymbol                 rune
//     - A unicode character used for the currency symbol. This
//       currency symbol will be prefixed or attached to the
//       beginning of the final, returned number string. In the
//       United States, the currency symbol is the dollar sign,
//       '$'.
//
//
//  ePrefix                        ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - Formatted Number String. If this method completes
//       successfully, this returned string will contain numeric
//       digits separated into thousands by the delimiter character
//       supplied in input parameter, 'integerDigitsSeparator'. It
//       will also be prefixed with the currency symbol supplied by
//       input parameter, 'currencySymbol'. Floating point numeric
//       values will be properly formatted and delimited using the
//       'decimalSeparator' character.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (ns NumStrBasicUtility) DlimCurrencyStr(
	rawNumStr string,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	decimalSeparator rune,
	currencySymbol rune,
	ePrefix ErrPrefixDto) (
	string,
	error) {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	ePrefix.SetEPref("NumStrBasicUtility.DlimCurrencyStr() ")

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.delimitCurrencyStr(
		rawNumStr,
		integerDigitsSeparator,
		integerDigitsGroupingSequence,
		decimalSeparator,
		currencySymbol,
		ePrefix.XCtx(
			fmt.Sprintf(
				"rawNumStr='%v'\n", rawNumStr)))
}

// DelimitNumStr - is designed to delimit or format a number string
// with a thousands separator (i.e. ',') and, in the case of
// floating point numeric values, a decimal separator (a.k.a.
// decimal point).
//
// Examples:
//   rawNumStr == 1234567890 -> Output == "1,234,567,890".
//   rawNumStr == 1234567.89 -> Output == "1,234,567.89".
//   rawNumStr == -1234567890 -> Output == "-1,234,567,890".
//   rawNumStr == -1234567.89 -> Output == "-1,234,567.89".
//   rawNumStr == +1234567890 -> Output == "+1,234,567,890".
//   rawNumStr == +1234567.89 -> Output == "+1,234,567.89".
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  rawNumStr                      string
//     - A number string which will be formatted with thousands
//       separators. Floating point numeric values will be
//       correctly formatted with a decimal separator separating
//       integer and fractional digits.
//
//
//  integerDigitsSeparator         rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  decimalSeparator               rune
//     - A unicode character inserted into a number string to
//       separate integer and fractional digits. In the United
//       States, the decimal separator is the period character
//       ('.') and it is known as the decimal point.
//
//
//  ePrefix                        ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - Formatted Number String. If this method completes
//       successfully, this returned string will contain numeric
//       digits separated into thousands by the delimiter character
//       supplied in input parameter, 'integerDigitsSeparator'.
//       Floating point numeric values will be properly formatted
//       and delimited with the 'decimalSeparator' character.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (ns NumStrBasicUtility) DelimitNumStr(
	rawNumStr string,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	decimalSeparator rune,
	ePrefix ErrPrefixDto) (
	fmtNumStr string,
	err error) {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	ePrefix.SetEPref("NumStrBasicUtility.DelimitNumStr() ")

	return numStrBasicMechanics{}.ptr().delimitNumberStr(
		rawNumStr,
		integerDigitsSeparator,
		integerDigitsGroupingSequence,
		decimalSeparator,
		&ePrefix)
}

// ConvertNumStrToInt64 - Converts string of numeric digits to an
// int64 value. The 'rawNumStr' parameter may include a leading
// number sign value of plus or minus ('+' or '-'). If any digits
// following the number sign are non-numeric, they will be
// excluded. All numeric digits in 'rawNumStr' will be consolidated
// to form the return 'int64' value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  rawNumStr                      string
//     - A number string which will be formatted with thousands
//       separators. Floating point numeric values will be
//       correctly formatted with a decimal separator separating
//       integer and fractional digits.
//
//
//  ePrefix                        ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int64
//     - If this method completes successfully, it well return an
//       int64 numeric value equivalent of the input parameter,
//       'rawNumStr'.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (ns *NumStrBasicUtility) ConvertNumStrToInt64(
	rawNumStr string,
	ePrefix ErrPrefixDto) (
	int64,
	error) {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	ePrefix.SetEPref(
		"NumSt rBasicUtility.ConvertNumStrToInt64() ")

	signChar,
		numRunes,
		lenNumRunes,
		err := numStrBasicAtom{}.ptr().
		parseIntRunesFromNumStr(
			rawNumStr,
			ePrefix.XCtx(
				fmt.Sprintf("rawNumStr='%v'",
					rawNumStr)))

	if err != nil {
		return int64(0), err
	}

	if lenNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rawNumStr' contained\n"+
			"zero numeric digits!\n",
			ePrefix.XCtxEmpty().String())

		return int64(0), err
	}

	signVal := 1

	if signChar == '-' {
		signVal = -1
	}

	var newInt64 int64

	newInt64,
		err =
		numStrBasicMechanics{}.ptr().convertRunesToInt64(
			numRunes,
			signVal,
			ePrefix.XCtx(
				"numRunes, signVal"))

	return newInt64, err
}

// ConvertStrToIntNumStr - Converts a string of characters to
// a string consisting of a pure number string consisting of
// an optional leading sign character ('+' or '-') and a
// series of numeric digits.
//
func (ns *NumStrBasicUtility) ConvertStrToIntNumStr(
	rawNumStr string,
	ePrefix ErrPrefixDto) (
	intNumStr string,
	err error) {

	ePrefixLocal := ePrefix.Copy()

	ePrefixLocal.SetEPref("NumStrBasicUtility.ConvertStrToIntNumStr()")

	intNumStr = ""
	err = nil

	var signChar rune
	var numRunes []rune
	var lenNumRunes int

	signChar,
		numRunes,
		lenNumRunes,
		err = numStrBasicAtom{}.ptr().
		parseIntRunesFromNumStr(
			rawNumStr,
			ePrefix.XCtx(
				fmt.Sprintf("rawNumStr='%v'",
					rawNumStr)))

	if err != nil {
		return "", err
	}

	if lenNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rawNumStr' contained\n"+
			"zero numeric digits!\n",
			ePrefix.XCtxEmpty().String())

		return "", err
	}

	if signChar != 0 {
		intNumStr =
			string(signChar) + string(numRunes)
	} else {
		intNumStr = string(numRunes)
	}

	return intNumStr, err
}

// ConvertInt64ToStr - Converts an int64 to a string of numeric
// characters. If the original number is less than zero, the first
// character of the numeric string is a minus sign ('-').
//
func (ns NumStrBasicUtility) ConvertInt64ToStr(
	num int64) string {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	return strconv.FormatInt(num, 10)
}

// ConvertRunesToInt64 - Converts a rune array of integer values to
// an int64 value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  rAry                []rune
//     - An array of runes containing numeric text characters
//       each with a value between '0-9' (0 and 9 included).
//       Combined, these characters represent an integer value
//       which be returned as an int64.
//
//  signVal             int
//     - This parameter must be set to one of two values: +1 or -1.
//
//       +1 signals that 'rAry' is a value Greater Than Or Equal To
//       Zero.
//
//       -1 signals that 'rAry' is a value Less Than Zero.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int64
//     - If this method completes successfully, an int64 will be
//       returned. This value is equivalent to the numeric digits
//       passed through input parameter 'rAry'.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (ns *NumStrBasicUtility) ConvertRunesToInt64(
	rAry []rune,
	signVal int,
	ePrefix ErrPrefixDto) (int64, error) {

	nStrBasicMech := numStrBasicMechanics{}

	ePrefix.SetEPref(
		"numStrBasicMechanics.ConvertRunesToInt64()")

	return nStrBasicMech.convertRunesToInt64(
		rAry,
		signVal,
		&ePrefix)
}

// ParseNumString - Parses a number string and returns
// the value as a Type 'NumStrDto'.
func (ns *NumStrBasicUtility) ParseNumString(
	numStr string,
	ePrefix ErrPrefixDto) (NumStrDto, error) {

	ePrefix.SetEPref("NumStrBasicUtility.ParseNumString()")
	nDto := NumStrDto{}

	return nDto.ParseNumStr(numStr, ePrefix)
}

// ConvertStrToIntNumRunes - Receives an integer string and returns
// a slice of runes.
//
func (ns *NumStrBasicUtility) ConvertStrToIntNumRunes(
	str string) []rune {

	return numStrBasicAtom{}.ptr().
		parseIntRunesFromNumStr(
			str,
		)
}

// ConvertStrToFloat64 - Converts a string of numbers to a float64 value.
//
func (ns *NumStrBasicUtility) ConvertStrToFloat64(
	str string,
	ePrefix string) (
	resultFloat64 float64,
	err error) {

	resultFloat64 = 0.0
	err = nil

	var nDto NumStrDto

	nStrDto := NumStrDto{}

	nDto, err = nStrDto.ParseNumStr(
		str,
		ePrefix)

	if err != nil {
		return resultFloat64, err
	}

	var nDtoNumStr string

	nDtoNumStr,
		err = nDto.GetNumStr(ePrefix)

	resultFloat64, err = strconv.ParseFloat(nDtoNumStr, 64)

	if err != nil {
		resultFloat64 = 0.0
		return resultFloat64,
			fmt.Errorf(ePrefix+
				"Error returned from strconv.ParseFloat(nDto.GetNumStr(), 64)\n"+
				"nDto.GetNumStr()= '%v'\n"+
				"Error = %v",
				nDtoNumStr,
				err.Error())
	}

	return resultFloat64, err
}

// ConvertInt64ToIntegerFloat64Value - Receives an int64 value and converts to a
// float64 value.  All of the digits are positioned to the right of the decimal
// place.
func (ns *NumStrBasicUtility) ConvertInt64ToIntegerFloat64Value(
	i64 int64) (
	resultFloat64 float64) {

	resultFloat64 = float64(i64)

	return resultFloat64

}

// ConvertInt64ToFractionalValue - Converts an int64 value to a float64 with
// all digits to the right of the decimal place.
func (ns *NumStrBasicUtility) ConvertInt64ToFractionalValue(i64 int64) (float64, error) {

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.convertInt64ToFractionalValue(i64)
}

// SetPrecision - Receives a string consisting of numeric
// digits. All numeric punctuation  (i.e. '.', ',', '$' is
// automatically removed. The string is converted to a
// floating point number string and the decimal placement
// is made according to the 'precision' parameter.  For example,
// a number string of '123456' with a precision parameter of '3'
// is converted to the number string '123.456'.
//
// Note: A sign character may be placed in the first character
// position before the first digit. Example: '-' minus or
// '+' plus.
//
func (ns *NumStrBasicUtility) ScaleNumStr(
	str string,
	precision uint,
	roundResult bool) (NumStrDto, error) {

	ePrefix := "NumStrBasicUtility.ScaleNumStr() "

	nStrDto := NumStrDto{}

	return nStrDto.SetPrecision(str, precision, roundResult, ePrefix)

}

// SetCountryAndCurrency - Sets the Country and Currency flags for the
// current NumStrBasicUtility values.
//
func (ns *NumStrBasicUtility) SetCountryAndCurrency(country string) error {

	nStrBasicMech := numStrBasicMechanics{}

	xNS, err := nStrBasicMech.getCountryAndCurrency(
		country,
		"NumStrBasicUtility.SetCountryAndCurrency() ")

	if err != nil {
		return err
	}

	ns.CurrencySymbol = xNS.CurrencySymbol
	return nil
}
