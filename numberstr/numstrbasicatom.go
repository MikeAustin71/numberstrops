package numberstr

import (
	"fmt"
	"sync"
)

type numStrBasicAtom struct {
	lock *sync.Mutex
}

// delimitIntSeparators - Inserts thousands separators into integer
// number strings.
//
//           Example:
//              pureNumStr = 1000000000
//                  numStr = 1,000,000,000
//
// This method accepts NumStrFmtSpecDto input parameter which may
// be configured to apply integer digit separators and integer
// digits group sequences from any country or culture.
//
// IMPORTANT
//
// The input parameter 'pureNumStr' number string must represent
// an integer value. If a float point value with fractional digits
// is submitted, inaccurate results will be generated.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  pureNumStr          string
//     - A string of numeric digits representing an integer value.
//       Submitting a floating point value will generate inaccurate
//       results. The numeric value MUST BE and integer value.
//
//
//  numStrFmtSpec       *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in number
//       strings.
//
//       This method uses the decimal separator character extracted
//       from 'currencyValue'.
//
//       If this instance is invalid, it will be reset to the
//       default United States Number String Format specification.
//
//       type NumStrFmtSpecDto struct {
//         idNo           uint64
//         idString       string
//         description    string
//         tag            string
//         countryCulture NumStrFmtSpecCountryDto
//         absoluteValue  NumStrFmtSpecAbsoluteValueDto
//         currencyValue  NumStrFmtSpecCurrencyValueDto
//         signedNumValue NumStrFmtSpecSignedNumValueDto
//         sciNotation    NumStrFmtSpecSciNotationDto
//       }
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
//  numStr              string
//     - If this method completes successfully, this returned
//       number string will consist of the numeric digits submitted
//       in input parameter 'pureNumStr' properly configured with
//       thousands separators. In the United States, the thousands
//       separator is the comma character (',').
//           Example:
//              pureNumStr = 1000000000
//                  numStr = 1,000,000,000
//
func (nStrBasicAtom *numStrBasicAtom) delimitIntSeparators(
	pureNumStr string,
	integerDigitsSeparator rune,
	integerGroupingSequence []uint,
	ePrefix *ErrPrefixDto) (
	numStr string,
	err error) {

	if nStrBasicAtom.lock == nil {
		nStrBasicAtom.lock = new(sync.Mutex)
	}

	nStrBasicAtom.lock.Lock()

	defer nStrBasicAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrBasicAtom.delimitIntSeparators()")

	if integerDigitsSeparator == 0 {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'integerDigitsSeparator' is INVALID!\n"+
			"integerDigitsSeparator == ZERO!\n",
			ePrefix.String())

		return numStr, err
	}

	rawNumRunes := []rune(pureNumStr)

	lInStr := len(rawNumRunes)

	if lInStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr is an empty string.\n",
			ePrefix.String())

		return numStr, err
	}

	if integerGroupingSequence == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a nil pointer.\n",
			ePrefix.String())

		return numStr, err
	}

	lenIGrpSeq := len(integerGroupingSequence)

	if lenIGrpSeq == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a ZERO length array.\n",
			ePrefix.String())

		return numStr, err
	}

	pureNumRunes := make([]rune, 0, 256)

	haveFirstNumericDigit := false
	var signChar rune

	for i := 0; i < lInStr; i++ {

		if !haveFirstNumericDigit &&
			signChar == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			signChar = rawNumRunes[i]

			continue
		}

		if rawNumRunes[i] >= '0' &&
			rawNumRunes[i] <= '9' {
			pureNumRunes = append(pureNumRunes,
				rawNumRunes[i])
			haveFirstNumericDigit = true
		}
	}

	lInStr = len(pureNumRunes)

	if lInStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumStr' is invalid!\n"+
			"pureNumStr contains no integer digits.\n",
			ePrefix.String())
	}

	//integerDigitsSeparator :=
	//	numStrFmtSpec.GetIntegerDigitsSeparator()

	groupCnt := uint(0)
	maxGroupCnt := integerGroupingSequence[0]
	currGroupCntIdx := 0
	lastGroupCntIdx := lenIGrpSeq - 1

	numOfDelims := 0
	remainder := 0
	numOfPreDelims := 0
	remainingDigits := lInStr

	for j := 0; j < lenIGrpSeq; j++ {

		if j < lastGroupCntIdx {
			remainingDigits -= int(integerGroupingSequence[j])
			if remainingDigits > 0 {
				numOfPreDelims++
			}

		} else {
			// Must be last index
			if remainingDigits > 0 {
				numOfDelims =
					remainingDigits / int(integerGroupingSequence[j])
				remainder =
					remainingDigits % int(integerGroupingSequence[j])
			}
		}
	}

	if numOfDelims > 0 &&
		remainder == 0 {
		numOfDelims--
	}

	lenOutRunes :=
		lInStr +
			numOfPreDelims +
			numOfDelims

	if signChar != 0 {
		lenOutRunes++
	}

	outRunes := make([]rune, lenOutRunes, 256)
	outIdx := lenOutRunes - 1

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			groupCnt++
			outRunes[outIdx] = pureNumRunes[i]
			outIdx--

			if groupCnt == maxGroupCnt && i != 0 {

				groupCnt = 0

				outRunes[outIdx] = integerDigitsSeparator
				outIdx--

				if currGroupCntIdx+1 > lastGroupCntIdx {

					maxGroupCnt = integerGroupingSequence[currGroupCntIdx]

				} else {

					currGroupCntIdx++

					maxGroupCnt = integerGroupingSequence[currGroupCntIdx]

				}

			} // End of if groupCnt == maxGroupCnt && i != 0

		} // End Of if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9'

	}

	// Check and allow for sign designators
	if signChar != 0 {

		outRunes[0] = signChar

	}

	numStr = string(outRunes)

	return numStr, err
}

// parseIntFracRunesFromNumStr - Receives a string of numeric digits
// ('numStr') and proceeds to parse the integer and fractional
// digits returning them in separate rune arrays.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStr              string
//     - A string of numeric digits which may or may not contain
//       a series of fractional digits to the right of a decimal
//       separator character.
//
//
//  numStrFmtSpec       *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in number
//       strings.
//
//       This method uses the decimal separator character extracted
//       from 'currencyValue'.
//
//       If this instance is invalid, it will be reset to the
//       default United States Number String Format specification.
//
//       type NumStrFmtSpecDto struct {
//         idNo           uint64
//         idString       string
//         description    string
//         tag            string
//         countryCulture NumStrFmtSpecCountryDto
//         absoluteValue  NumStrFmtSpecAbsoluteValueDto
//         currencyValue  NumStrFmtSpecCurrencyValueDto
//         signedNumValue NumStrFmtSpecSignedNumValueDto
//         sciNotation    NumStrFmtSpecSciNotationDto
//       }
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
//  signChar            rune
//     - If a leading numeric sign character exists in input
//       parameter 'numStr', it will be extracted and returned in
//       the 'signChar' parameter. This returned rune will be set
//       to one of three values.
//
//       '+' (ascii 43)   Signals a plus or positive sign value.
//                        This means that a leading plus sign was
//                        present in 'numStr'.
//
//       '-' (ascii 45)   Signals a minus or negative sign value.
//                        This means that a leading minus sign was
//                        present in 'numStr'.
//
//       0 integer value  Signals that neither a plus ('+') or
//                        minus ('-') was present in 'numStr'.
//                        This means that by default, the 'numStr'
//                        value is positive.
//
//
//  intNumRunes         []runes
//     - An array of runes containing the numeric digits which
//       comprise the integer component of input parameter
//       'numStr'.
//
//
//  lenIntNumRunes      int
//     - This value is represents the length of rune array
//       'intNumRunes'.
//
//
//  fracNumRunes        []rune
//     - An array of runes containing the numeric digits which
//       comprise the fractional component of input parameter
//       'numStr'.
//
//
//  lenFracNumRunes     int
//     - This value is represents the length of rune array
//       'fracNumRunes'.
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
func (nStrBasicAtom *numStrBasicAtom) parseIntFracRunesFromNumStr(
	numStr string,
	decimalSeparator rune,
	ePrefix *ErrPrefixDto) (
	signChar rune,
	intNumRunes []rune,
	lenIntNumRunes int,
	fracNumRunes []rune,
	lenFracNumRunes int,
	err error) {

	if nStrBasicAtom.lock == nil {
		nStrBasicAtom.lock = new(sync.Mutex)
	}

	nStrBasicAtom.lock.Lock()

	defer nStrBasicAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrBasicAtom.parseIntFracRunesFromNumStr()")

	if decimalSeparator == 0 {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'decimalSeparator' is INVALID!\n"+
			"decimalSeparator == ZERO!\n",
			ePrefix.String())

		return signChar, intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	signChar = 0
	intNumRunes = make([]rune, 0, 50)
	fracNumRunes = make([]rune, 0, 50)

	lenNumStr := len(numStr)

	if lenNumStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter numStr is invalid!\n"+
			"'numStr' is an empty or zero length string.\n",
			ePrefix.String())
		return signChar, intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	rawNumRunes := []rune(numStr)

	lenNumStr = len(rawNumRunes)

	haveFirstNumericDigit := false
	haveFracSeparator := false

	for i := 0; i < lenNumStr; i++ {

		if !haveFirstNumericDigit &&
			signChar == 0 &&
			(rawNumRunes[i] == '+' ||
				rawNumRunes[i] == '-') {

			signChar = rawNumRunes[i]

			continue
		}

		if !haveFracSeparator &&
			rawNumRunes[i] == decimalSeparator {

			haveFracSeparator = true

			continue
		}

		if rawNumRunes[i] >= '0' &&
			rawNumRunes[i] <= '9' {

			haveFirstNumericDigit = true

			if !haveFracSeparator {

				intNumRunes = append(intNumRunes,
					rawNumRunes[i])

				lenIntNumRunes++

			} else {

				// Must Have the Fractional Separator
				fracNumRunes = append(fracNumRunes,
					rawNumRunes[i])

				lenFracNumRunes++

			}
		} // End of Numeric Digit If statement

	} // End of 'i' loop

	if lenIntNumRunes == 0 &&
		lenFracNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is invalid!\n"+
			"'numStr' contained ZERO Numeric Digits.\n",
			ePrefix.String())

		signChar = 0
	}

	return signChar, intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
}

// parseIntFracRunesFromRuneArray - Receives an array of runes
// representing the absolute value of number (always positive).
// This method then proceeds to extract the integer and fractional
// components of this absolute value and return them as separate
// rune arrays.
//
// The input parameter 'precision' is used to determine the number
// of fractional digits to the right of the decimal separator
// (a.k.a. decimal point).
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absNumValRuneArray  []rune
//     - An array of runes containing the numeric digits in the
//       absolute value of a number. This number may be an integer
//       or a floating point value with fractional digits.
//
//       As a absolute numeric value, this number is always
//       positive.
//
//
//  precision           uint
//     - As used here, 'precision' defines the number of fractional
//       numeric digits to the right of the decimal separator in
//       numeric value contained in input parameter,
//       'absNumValRuneArray'.
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
//
//
//  intNumRunes         []runes
//     - An array of runes containing the numeric digits which
//       comprise the integer component of input parameter
//       'absNumValRuneArray'.
//
//
//  lenIntNumRunes      int
//     - This value is represents the length of rune array
//       'intNumRunes'.
//
//
//  fracNumRunes        []rune
//     - An array of runes containing the numeric digits which
//       comprise the fractional component of input parameter
//       'absNumValRuneArray'.
//
//
//  lenFracNumRunes     int
//     - This value is represents the length of rune array
//       'fracNumRunes'.
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
func (nStrBasicAtom *numStrBasicAtom) parseIntFracRunesFromRuneArray(
	absNumValRuneArray []rune,
	precision uint,
	ePrefix *ErrPrefixDto) (
	intNumRunes []rune,
	lenIntNumRunes int,
	fracNumRunes []rune,
	lenFracNumRunes int,
	err error) {

	if nStrBasicAtom.lock == nil {
		nStrBasicAtom.lock = new(sync.Mutex)
	}

	nStrBasicAtom.lock.Lock()

	defer nStrBasicAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrBasicAtom.parseIntFracRunesFromRuneArray()")

	intNumRunes = make([]rune, 0, 50)
	fracNumRunes = make([]rune, 0, 50)
	intPrecision := int(precision)

	lenAbsValArray := len(absNumValRuneArray)

	if lenAbsValArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absNumValRuneArray' is invalid!\n"+
			"absNumValRuneArray is a ZERO Length Array.\n",
			ePrefix.String())

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	if intPrecision > lenAbsValArray {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'precision' is invalid!\n"+
			"precision is greater than the Length of Array 'absNumValRuneArray'.\n"+
			"precision=='%v'\n"+
			"Length Of Absolute Value Array=='%v'\n",
			ePrefix.String(),
			precision,
			lenAbsValArray)

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err

	}

	pureAbsValRuneArray := make([]rune, 0, lenAbsValArray)

	for i := 0; i < lenAbsValArray; i++ {

		if absNumValRuneArray[i] >= '0' &&
			absNumValRuneArray[i] <= '9' {

			pureAbsValRuneArray = append(pureAbsValRuneArray, absNumValRuneArray[i])
		}

	}

	lenAbsValArray = len(pureAbsValRuneArray)

	if lenAbsValArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absNumValRuneArray' is invalid!\n"+
			"absNumValRuneArray contains ZERO Numeric Digits.\n",
			ePrefix.String())

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	if intPrecision == lenAbsValArray {

		fracNumRunes = make([]rune, lenAbsValArray, 10)

		lenFracNumRunes = copy(fracNumRunes, pureAbsValRuneArray)

		if lenFracNumRunes != lenAbsValArray {
			err = fmt.Errorf("%v\n"+
				"Error: 'precision' is equal to the number of numeric\n"+
				"digits in absNumValRuneArray. However, the copy operation\n"+
				"failed. runes copied='%v'  runes to copy='%v'\n",
				ePrefix.String(),
				lenFracNumRunes,
				lenAbsValArray)

		}

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	// precision < lenAbsValArray
	lenIAry := lenAbsValArray - intPrecision

	intNumRunes = make([]rune, lenIAry, 50)

	lenIntNumRunes = copy(intNumRunes, pureAbsValRuneArray[0:lenIAry])

	if lenIntNumRunes != lenIAry {
		err = fmt.Errorf("%v\n"+
			"Error: 'precision' is less than the number of numeric digits\n"+
			"in absNumValRuneArray. However, the integer copy operation\n"+
			"failed. integer runes copied='%v'  integer runes to copy='%v'\n",
			ePrefix.String(),
			lenIntNumRunes,
			lenIAry)

		return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
	}

	fracNumRunes = make([]rune, intPrecision, 50)

	lenFracNumRunes = copy(fracNumRunes, pureAbsValRuneArray[lenIntNumRunes:])

	if lenFracNumRunes != intPrecision {
		err = fmt.Errorf("%v\n"+
			"Error: 'precision' is less than the number of numeric digits\n"+
			"in absNumValRuneArray. However, the fractional copy operation\n"+
			"failed. fractional runes copied='%v'  fractional runes to copy='%v'\n",
			ePrefix.String(),
			lenFracNumRunes,
			intPrecision)
	}

	return intNumRunes, lenIntNumRunes, fracNumRunes, lenFracNumRunes, err
}

// ptr - Returns a pointer to a new instance of numStrBasicAtom.
//
func (nStrBasicAtom numStrBasicAtom) ptr() *numStrBasicAtom {

	if nStrBasicAtom.lock == nil {
		nStrBasicAtom.lock = new(sync.Mutex)
	}

	nStrBasicAtom.lock.Lock()

	defer nStrBasicAtom.lock.Unlock()

	newNumStrBasicAtom := new(numStrBasicAtom)

	newNumStrBasicAtom.lock = new(sync.Mutex)

	return newNumStrBasicAtom
}
