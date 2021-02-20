package numberstr

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type numStrBasicMechanics struct {
	lock *sync.Mutex
}

// convertInt64ToFractionalValue - Converts an int64 value to a float64 with
// all digits to the right of the decimal place.
func (nStrBasicMech *numStrBasicMechanics) convertInt64ToFractionalValue(
	i64 int64) float64 {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	ex := 1
	f64 := float64(i64)
	exp := math.Pow10(ex)
	for f64 > exp {
		ex++
		exp = math.Pow10(ex)
	}

	r64 := f64 / math.Pow10(ex)

	return r64

}

// convertRunesToInt64 - Converts a rune array to an
// int64 value.
//
// Input parameter 'signVal' must be equal to +1 or -1.
// Any other value will trigger an error.
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
func (nStrBasicMech *numStrBasicMechanics) convertRunesToInt64(
	rAry []rune,
	signVal int,
	ePrefix *ErrPrefixDto) (
	int64,
	error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrBasicMechanics.convertRunesToInt64()")

	var numRunes []rune
	var lenNumRunes int
	var err error

	_,
		numRunes,
		lenNumRunes,
		err = numStrBasicAtom{}.ptr().
		parseIntRunesFromNumStr(
			string(rAry),
			ePrefix.XCtx(
				"string(rAry)"))

	if err != nil {
		return 0, err
	}

	if lenNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Rune array 'rAry' is invalid!\n"+
			"'rAry' contains ZERO numeric digits.\n",
			ePrefix.XCtxEmpty().String())
	}

	var numStr string

	if signVal == -1 {
		numStr = "-" + string(numRunes)
	} else {
		numStr = string(numRunes)
	}

	var numVal int64

	numVal,
		err =
		strconv.ParseInt(numStr, 10, 64)

	if err != nil {
		return numVal,
			fmt.Errorf("%v\n"+
				"Error returned by strconv.ParseInt(numStr, 10, 64).\n"+
				"numStr='%v'\n"+
				"Error= '%v'\n",
				ePrefix.XCtxEmpty().String(),
				numStr,
				err.Error())
	}

	return numVal, err
}

// convertStrToFloat64 - Receives a floating point value in the
// form of a number string and converts it to a float64 value.
//
// If the input value exceeds the capacity of a float64, an error
// will be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  rawNumStr           string
//     - A number string which will be converted to a float64
//       value and returned to the calling function.
//
//
//  decimalSeparator    rune
//     - A unicode character inserted into a number string to
//       separate integer and fractional digits. In the United
//       States, the decimal separator is the period character
//       ('.') and it is known as the decimal point.
//
//       This text character will be used to identify integer and
//       fractional digits in input parameter, 'rawNumStr'.
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
//  resultFloat64       float64
//     - The result generated from converting input parameter
//       'rawNumStr' to a float64.
//
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
func (nStrBasicMech *numStrBasicMechanics) convertStrToFloat64(
	rawNumStr string,
	decimalSeparator rune,
	ePrefix *ErrPrefixDto) (
	resultFloat64 float64,
	err error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrBasicMechanics.convertStrToFloat64()")

	var floatingPtNumStr string

	floatingPtNumStr,
		err = numStrBasicNanobot{}.
		ptr().convertNumStrToFloatingPointNumStr(
		rawNumStr,
		decimalSeparator,
		ePrefix.XCtx("rawNumStr"))

	if err != nil {
		return resultFloat64, err
	}

	floatingPtNumStr = strings.Replace(
		floatingPtNumStr,
		string(decimalSeparator),
		".",
		1)

	var err2 error

	resultFloat64,
		err2 = strconv.ParseFloat(
		floatingPtNumStr,
		64)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from strconv.ParseFloat()\n"+
			"Input floatingPtNumStr='%v'\n"+
			"Error= '%v'\n",
			ePrefix.XCtxEmpty().String(),
			floatingPtNumStr,
			err2.Error())
	}

	return resultFloat64, err
}

// ConvertStrToIntNumRunes - Receives an integer string and returns
// a slice of runes. Note, thousands separator (',') and currency signs ('$')
// will be ignored and excluded from the array of runes returned by this
// method. In order to take advantage of this exclusion, the thousands
// separator and the currency symbol must be previously initialized in
// the NumStrBasicUtility.NumStrFormatSpec data field.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStr              string
//     - A string of numeric digits. This number string will be
//       converted to an array of runes and returned to the calling
//       function.
//
//
//  nStrFmtSpecDto      *NumStrFmtSpecDto
//     - A pointer to an instance of NumStrFmtSpecDto. This object
//       contains the numeric separators such as the thousands
//       separator (','), the decimal separator ('.') and the
//       currency symbol. These separator characters will be used
//       in processing input parameter 'numStr' and converting
//       it to an array of runes.
//
//       If this parameter is judged to be empty or invalid, it
//       will be set to default values.
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
func (nStrBasicMech *numStrBasicMechanics) convertStrToIntNumRunes(
	numStr string,
	nStrFmtSpecDto *NumStrFmtSpecDto,
	ePrefix *ErrPrefixDto) (
	runes []rune,
	err error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrBasicMechanics.convertStrToIntNumRunes()")

	if nStrFmtSpecDto == nil {
		nStrFmtSpecDto = &NumStrFmtSpecDto{}
	}

	err = nStrFmtSpecDto.SetToDefaultIfEmpty(
		ePrefix.XCtx("nStrFmtSpecDto"))

	if err != nil {
		return runes, err
	}

	thousandsSeparator :=
		nStrFmtSpecDto.GetThousandsSeparator()

	currencySymbol :=
		nStrFmtSpecDto.GetCurrencySymbol()

	baseRunes := []rune(numStr)
	runes = make([]rune, 0, 100)
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false

	for i := 0; i < lBaseRunes; i++ {

		if (baseRunes[i] == '-' || baseRunes[i] == '+') &&
			len(runes) == 0 && isEndRunes == false &&
			i < lBaseRunes-2 &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' {

			runes = append(runes, baseRunes[i])
			isStartRunes = true

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			runes = append(runes, baseRunes[i])
			isStartRunes = true

		} else if baseRunes[i] == thousandsSeparator || baseRunes[i] == currencySymbol {

			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}
	}

	return runes, err
}

// delimitCurrencyStr - Inserts a Currency Symbol and a Thousands
// Separator in a number string containing a decimal point.
func (nStrBasicMech *numStrBasicMechanics) delimitCurrencyStr(
	rawNumStr string,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	decimalSeparator rune,
	currencySymbol rune,
	ePrefix *ErrPrefixDto) (
	numStr string,
	err error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	ePrefix.SetEPref(
		"numStrBasicMechanics.delimitCurrencyStr() ")

	nStrBasicAtom := numStrBasicAtom{}

	var signChar rune
	var intNumRunes, fracNumRunes []rune
	var lenIntNumRunes, lenFracNumRunes int

	signChar,
		intNumRunes,
		lenIntNumRunes,
		fracNumRunes,
		lenFracNumRunes,
		err = nStrBasicAtom.parseIntFracRunesFromNumStr(
		rawNumStr,
		decimalSeparator,
		ePrefix.XCtx("rawNumStr"))

	if err != nil {
		return numStr, err
	}

	if lenIntNumRunes == 0 &&
		lenFracNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: There are no viable numeric digits\n"+
			"int input parameter 'rawNumStr'.\n"+
			"rawNumStr='%v'\n",
			ePrefix.XCtxEmpty().String(),
			rawNumStr)

		return numStr, err
	}

	outStr := make([]rune, 1, 30)

	outStr[0] = currencySymbol

	// Exclude leading plus ('+') sign
	if signChar == '-' {
		outStr = append(outStr, signChar)
	} else if signChar == '+' {
		signChar = 0
	}

	if lenIntNumRunes > 0 {

		var delimitedNumStr string

		delimitedNumStr,
			err = nStrBasicAtom.delimitIntSeparators(
			string(intNumRunes),
			integerDigitsSeparator,
			integerDigitsGroupingSequence,
			ePrefix.XCtx(
				"string(intNumRunes)"))

		if err != nil {
			return numStr, err
		}

		numStr = string(outStr) + delimitedNumStr

	} else {
		// lenIntNumRunes == 0

		outStr = append(outStr, '0')

		numStr = string(outStr)

	}

	if lenFracNumRunes > 0 {
		numStr += string(decimalSeparator)
		numStr += string(fracNumRunes)
	}

	return numStr, err
}

// delimitNumberStr - Inserts a Currency Symbol and a Thousands
// Separator in a number string containing a decimal point.
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
//  ePrefix                        *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  fmtNumStr                      string
//     - Formatted Number String. If this method completes
//       successfully, this returned string will contain numeric
//       digits separated into thousands by the delimiter character
//       supplied in input parameter, 'integerDigitsSeparator'.
//       Floating point numeric values will be properly formatted
//       and delimited with the 'decimalSeparator' character.
//
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
func (nStrBasicMech *numStrBasicMechanics) delimitNumberStr(
	rawNumStr string,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	decimalSeparator rune,
	ePrefix *ErrPrefixDto) (
	fmtNumStr string,
	err error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	ePrefix.SetEPref(
		"numStrBasicMechanics.delimitNumberStr() ")

	nStrBasicAtom := numStrBasicAtom{}

	var signChar rune
	var intNumRunes, fracNumRunes []rune
	var lenIntNumRunes, lenFracNumRunes int

	signChar,
		intNumRunes,
		lenIntNumRunes,
		fracNumRunes,
		lenFracNumRunes,
		err = nStrBasicAtom.parseIntFracRunesFromNumStr(
		rawNumStr,
		decimalSeparator,
		ePrefix.XCtx("rawNumStr"))

	if err != nil {
		return fmtNumStr, err
	}

	if lenIntNumRunes == 0 &&
		lenFracNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: There are no viable numeric digits\n"+
			"int input parameter 'rawNumStr'.\n"+
			"rawNumStr='%v'\n",
			ePrefix.XCtxEmpty().String(),
			rawNumStr)

		return fmtNumStr, err
	}

	outStr := make([]rune, 0, 30)

	if signChar != 0 {
		outStr = append(outStr, signChar)
	}

	if lenIntNumRunes > 0 {

		var delimitedNumStr string

		delimitedNumStr,
			err = nStrBasicAtom.delimitIntSeparators(
			string(intNumRunes),
			integerDigitsSeparator,
			integerDigitsGroupingSequence,
			ePrefix.XCtx(
				"string(intNumRunes)"))

		if err != nil {
			return fmtNumStr, err
		}

		fmtNumStr = string(outStr) + delimitedNumStr

	} else {
		// lenIntNumRunes == 0

		if lenFracNumRunes > 0 {

			outStr = append(outStr, '0')

		} else {

			outStr = make([]rune, 1, 5)
			outStr[0] = '0'
		}

		fmtNumStr = string(outStr)
	}

	if lenFracNumRunes > 0 {
		fmtNumStr += string(decimalSeparator)
		fmtNumStr += string(fracNumRunes)
	}

	return fmtNumStr, err
}

// getCountryAndCurrency - Returns the Country and Currency runes as
// an instance of NumStrBasicUtility.
//
func (nStrBasicMech *numStrBasicMechanics) getCountryAndCurrency(
	country string,
	ePrefix string) (
	NumStrBasicUtility,
	error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	ePrefix += "numStrBasicMechanics.getCountryAndCurrency() "
	ns := NumStrBasicUtility{}

	lcStr := strings.ToLower(country)

	if strings.Contains(lcStr, "united states") {
		ns.Nation = "United States"
		ns.CurrencySymbol = NumStrCurrencySymbols[28]
		ns.ThousandsSeparator = ','
		ns.DecimalSeparator = '.'
		return ns, nil
	}

	if strings.Contains(lcStr, "united kingdom") {
		ns.Nation = "United Kingdom"
		ns.CurrencySymbol = NumStrCurrencySymbols[29]
		ns.DecimalSeparator = '.'
		return ns, nil
	}

	if strings.Contains(lcStr, "australia") {
		ns.Nation = "Australia"
		ns.CurrencySymbol = NumStrCurrencySymbols[0]
		return ns, nil
	}

	if strings.Contains(lcStr, "brazil") {
		ns.Nation = "Brazil"
		ns.CurrencySymbol = NumStrCurrencySymbols[1]
		return ns, nil
	}

	if strings.Contains(lcStr, "canada") {
		ns.Nation = "Canada"
		ns.CurrencySymbol = NumStrCurrencySymbols[2]
		ns.ThousandsSeparator = ','
		ns.DecimalSeparator = '.'
		return ns, nil
	}

	if strings.Contains(lcStr, "china") {
		ns.Nation = "China"
		ns.CurrencySymbol = NumStrCurrencySymbols[3]
		return ns, nil
	}

	if strings.Contains(lcStr, "colombia") {
		ns.Nation = "Colombia"
		ns.CurrencySymbol = NumStrCurrencySymbols[4]
		return ns, nil
	}

	if strings.Contains(lcStr, "czech") {
		ns.Nation = "Czechoslovakia"
		ns.CurrencySymbol = NumStrCurrencySymbols[5]
		return ns, nil
	}

	if strings.Contains(lcStr, "egypt") {
		ns.Nation = "Egypt"
		ns.CurrencySymbol = NumStrCurrencySymbols[6]
		return ns, nil
	}

	if strings.Contains(lcStr, "euro") {
		ns.Nation = "Euro"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "germany") {
		ns.Nation = "Germany"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "france") {
		ns.Nation = "France"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "italy") {
		ns.Nation = "Italy"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "spain") {
		ns.Nation = "Spain"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "hungary") {
		ns.Nation = "Hungary"
		ns.CurrencySymbol = NumStrCurrencySymbols[8]
		return ns, nil
	}

	if strings.Contains(lcStr, "iceland") {
		ns.Nation = "Iceland"
		ns.CurrencySymbol = NumStrCurrencySymbols[9]
		return ns, nil
	}

	if strings.Contains(lcStr, "indonesia") {
		ns.Nation = "Indonesia"
		ns.CurrencySymbol = NumStrCurrencySymbols[10]
		return ns, nil
	}

	if strings.Contains(lcStr, "israel") {
		ns.Nation = "Israel"
		ns.CurrencySymbol = NumStrCurrencySymbols[11]
		return ns, nil
	}

	if strings.Contains(lcStr, "japan") {
		ns.Nation = "Japan"
		ns.CurrencySymbol = NumStrCurrencySymbols[12]
		return ns, nil
	}

	if strings.Contains(lcStr, "korea") {
		ns.Nation = "Korea"
		ns.CurrencySymbol = NumStrCurrencySymbols[13]
		return ns, nil
	}

	if strings.Contains(lcStr, "malaysia") {
		ns.Nation = "Malaysia"
		ns.CurrencySymbol = NumStrCurrencySymbols[14]
		return ns, nil
	}

	if strings.Contains(lcStr, "mexico") {
		ns.Nation = "Mexico"
		ns.CurrencySymbol = NumStrCurrencySymbols[15]
		return ns, nil
	}

	if strings.Contains(lcStr, "norway") {
		ns.Nation = "Norway"
		ns.CurrencySymbol = NumStrCurrencySymbols[16]
		return ns, nil
	}

	if strings.Contains(lcStr, "netherlands") {
		ns.Nation = "Netherlands"
		ns.CurrencySymbol = NumStrCurrencySymbols[17]
		return ns, nil
	}

	if strings.Contains(lcStr, "pakistan") {
		ns.Nation = "Pakistan"
		ns.CurrencySymbol = NumStrCurrencySymbols[18]
		return ns, nil
	}

	if strings.Contains(lcStr, "russia") {
		ns.Nation = "Russia"
		ns.CurrencySymbol = NumStrCurrencySymbols[19]
		return ns, nil
	}

	if strings.Contains(lcStr, "saudi") {
		ns.Nation = "Saudi Arabia"
		ns.CurrencySymbol = NumStrCurrencySymbols[20]
		return ns, nil
	}

	if strings.Contains(lcStr, "south africa") {
		ns.Nation = "South Africa"
		ns.CurrencySymbol = NumStrCurrencySymbols[21]
		return ns, nil
	}

	if strings.Contains(lcStr, "sweden") {
		ns.Nation = "Sweden"
		ns.CurrencySymbol = NumStrCurrencySymbols[22]
		return ns, nil
	}

	if strings.Contains(lcStr, "switzerland") {
		ns.Nation = "Switzerland"
		ns.CurrencySymbol = NumStrCurrencySymbols[23]
		return ns, nil
	}

	if strings.Contains(lcStr, "taiwan") {
		ns.Nation = "Taiwan"
		ns.CurrencySymbol = NumStrCurrencySymbols[24]
		return ns, nil
	}

	if strings.Contains(lcStr, "turkey") {
		ns.Nation = "Turkey"
		ns.CurrencySymbol = NumStrCurrencySymbols[25]
		return ns, nil
	}

	if strings.Contains(lcStr, "venezuela") {
		ns.Nation = "Venezuela"
		ns.CurrencySymbol = NumStrCurrencySymbols[26]
		return ns, nil
	}

	if strings.Contains(lcStr, "viet nam") {
		ns.Nation = "Viet Nam"
		ns.CurrencySymbol = NumStrCurrencySymbols[27]
		return ns, nil
	}

	return ns,
		fmt.Errorf(ePrefix+"\n"+
			"Failed to initialize country, %v.\n", country)

}

// ptr - Returns a pointer to a new instance of
// numStrBasicMechanics.
//
func (nStrBasicMech numStrBasicMechanics) ptr() *numStrBasicMechanics {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	newNumStrBasicMechanics := new(numStrBasicMechanics)

	newNumStrBasicMechanics.lock = new(sync.Mutex)

	return newNumStrBasicMechanics
}
