package numberstr

import (
	"fmt"
	"strconv"
	"strings"
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
// This number string must be pure number string with NO decimal
// points.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  num                 int
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
//  ePrefix             ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  numStr              string
//     - If this method completes successfully, this returned
//       string will contain numeric digits separated into thousands
//       by the delimiter character supplied in input parameter,
//       'numStrFmtSpec'.
//
//  err                 error
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
				fmt.Sprintf("numStr='%v'", numStr)))

	return numStr, err
}

// DLimI64 - Receives an int64 and returns a delimited number
// string with thousands separator (i.e. 1000000 -> 1,000,000).
// This number string must be pure number string with NO decimal
// points.
//
func (ns NumStrBasicUtility) DLimI64(num int64, delimiter byte) string {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.delimitThousands(
		fmt.Sprintf("%v", num),
		delimiter)
}

// DlimDecCurrStr - Inserts a Currency Symbol and a Thousands
// Separator in a number string containing a decimal point.
func (ns NumStrBasicUtility) DlimDecCurrStr(
	rawStr string,
	thousandsSeparator rune,
	decimal rune,
	currency rune) string {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.delimitDecimalCurrencyStr(
		rawStr,
		thousandsSeparator,
		decimal,
		currency)
}

// DelimitNumStr - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
// NOTE: This method will not handle number strings containing decimal fractions
// and currency characters. For these options see method ns.DlimDecCurrStr(),
// above.
func (ns NumStrBasicUtility) DelimitNumStr(pureNumStr string, thousandsSeparator byte) string {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.delimitThousands(
		pureNumStr,
		thousandsSeparator)
}

// ConvertNumStrToInt64 - Converts string of numeric digits to an int64 value.
// The str parameter may including a leading sign value ('-' or '+'). If any
// digits following the number sign are non-numeric, an error will be
// generated.
//
func (ns *NumStrBasicUtility) ConvertNumStrToInt64(
	numStr string,
	ePrefix ErrPrefixDto) (
	int64,
	error) {

	if ns.lock == nil {
		ns.lock = new(sync.Mutex)
	}

	ns.lock.Lock()

	defer ns.lock.Unlock()

	ePrefix := ErrPrefixDto{}.Ptr()

	ePrefix.SetEPref(
		"NumStrBasicUtility.ConvertNumStrToInt64() ")

	numStr = strings.TrimLeft(strings.TrimRight(numStr, " "), " ")

	nStrBasicMech := numStrBasicMechanics{}

	numRunes,
		err := nStrBasicMech.convertStrToIntNumRunes(
		numStr,
		&ns.NumStrFormatSpec,
		ePrefix)

	if err != nil {
		return int64(0), err
	}

	signVal := 1

	if numStr[0] == '-' {
		signVal = -1
	}

	var newInt64 int64

	newInt64,
		err =
		nStrBasicMech.convertRunesToInt64(
			numRunes,
			signVal,
			ePrefix.XCtx(
				"numRunes, signVal"))

	return newInt64, err
}

// ConvertStrToIntNumStr - Converts a string of characters to
// a string consisting of a sign character ('-') followed by
// a string of numeric characters.
//
func (ns *NumStrBasicUtility) ConvertStrToIntNumStr(
	numStr string,
	ePrefix ErrPrefixDto) (
	intNumStr string,
	err error) {

	ePrefixLocal := ePrefix.Copy()

	ePrefixLocal.SetEPref("NumStrBasicUtility.ConvertStrToIntNumStr()")

	intNumStr = ""
	err = nil

	err =
		ns.NumStrFormatSpec.SetToDefaultIfEmpty(
			&ePrefixLocal)

	if err != nil {
		return intNumStr, err
	}

	var outputNDto NumStrDto

	outputNDto,
		err = numStrDtoAtom{}.ptr().parseNumStr(
		numStr,
		&ns.NumStrFormatSpec,
		&ePrefixLocal)

	if err != nil {
		return intNumStr, err
	}

	intNumStr, err = outputNDto.GetNumStr(ePrefixLocal)

	return intNumStr, err
}

// ConvertInt64ToStr - Converts an int64 to a string of numeric
// characters. If the original number is less than zero, the first
// character of the numeric string is a minus sign ('-').
func (ns NumStrBasicUtility) ConvertInt64ToStr(
	num int64) (string, error) {

	ePrefixLocal :=
		ErrPrefixDto{}.NewEPrefOld("NumStrBasicUtility.ConvertInt64ToStr()")

	nStrDtoAtom := numStrDtoAtom{}

	numSepsDto := NumericSeparatorsDto{}

	numSepsDto.DecimalSeparator = ns.DecimalSeparator
	numSepsDto.CurrencySymbol = ns.CurrencySymbol
	numSepsDto.ThousandsSeparator = ns.ThousandsSeparator

	var nDto NumStrDto
	var err error
	var numStr string

	nDto,
		err = nStrDtoAtom.parseNumStr(
		strconv.FormatInt(num, 10),
		numSepsDto,
		ePrefix)

	if err != nil {
		return numStr, err
	}

	numStr, err = nDto.GetNumStr(ePrefix)

	return numStr, err
}

// ConvertRunesToInt64 - Converts a rune array to an int64 value.
//
func (ns *NumStrBasicUtility) ConvertRunesToInt64(
	rAry []rune, signVal int) (int64, error) {

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.convertRunesToInt64(rAry, signVal)
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
// a slice of runes. Note, thousands separator (',') and currency signs ('$')
// will be ignored and excluded from the array of runes returned by this
// method. In order to take advantage of this exclusion, the thousands
// separator and the currency symbol must be previously initialized in
// the NumStrBasicUtility data fields.
func (ns *NumStrBasicUtility) ConvertStrToIntNumRunes(
	str string) []rune {

	nStrBasicMech := numStrBasicMechanics{}

	return nStrBasicMech.convertStrToIntNumRunes(
		str,
		ns.ThousandsSeparator,
		ns.CurrencySymbol)
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
