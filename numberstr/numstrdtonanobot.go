package numberstr

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

type numStrDtoNanobot struct {
	lock *sync.Mutex
}

// findNumStrSignificantDigitLimits - Analyzes an array of
// characters, or runes, which constitute a pure number string. The
// method then proceeds calculate and return the significant
// digits from that array.
//
// A pure number string as used here defines an array of runes
// representing all numeric characters or runes. There are no
// non-numeric runes allowed. As such there are no plus or minus
// sign characters, thousands separators, or currency symbols
// allowed. Only numeric runes will be accepted.
//
// Example:
//  absAllRunes  precision   signVal     Result
//  001236700        4          1        123.67
//  000006700        4          1          0.67
//  001230000        4          1        123.0
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  absAllRunes         []rune
//     - An array of runes or characters. This rune must consist of
//       all numeric digits. If any non-numeric characters are
//       found in this array, an error will be returned.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
//
//
//  signVal         int
//     - Valid values for this parameter are plus one (+1) or minus
//       one (-1). This number sign value will determine the number
//       sign of the new NumStrDto instance returned by this method.
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
//  newNumStrDto        NumStrDto
//     - If this method completes successfully, the result of the
//       significant digits operation performed by this method will
//       be returned in the form of a new 'NumStrDto' instance.
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
func (nStrDtoNanobot *numStrDtoNanobot) findNumStrSignificantDigitLimits(
	numStrFormatSpec *NumStrFmtSpecDto,
	absAllRunes []rune,
	precision uint,
	signVal int,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.findNumStrSignificantDigitLimits()")

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.XCtxEmpty().String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	if absAllRunes == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absAllRunes' is invalid!\n"+
			"The rune array of 'absAllRunes' == 'nil'\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	lenAbsAllRunes := len(absAllRunes)

	if lenAbsAllRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absAllRunes' is invalid!\n"+
			"The rune array of 'absAllRunes' is a zero length array.\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	if signVal != 1 &&
		signVal != -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signVal' is invalid!\n"+
			"'signVal' MUST BE EQUAL to +1 or -1.\n"+
			"signVal='%v'\n",
			ePrefix.String(),
			signVal)

		return newNumStrDto, err
	}

	iPrecision := int(precision)
	firstIntIdx := -1
	lastIntIdx := -1
	lastFracIdx := -1

	isFractional := false

	if iPrecision > 0 {
		isFractional = true
	}

	lenAbsFracRunes := iPrecision
	lenAbsIntRunes := lenAbsAllRunes - lenAbsFracRunes

	for i := 0; i < lenAbsAllRunes; i++ {

		if i < lenAbsIntRunes {

			if firstIntIdx == -1 && absAllRunes[i] > '0' &&
				absAllRunes[i] <= '9' {
				firstIntIdx = i
			}

			lastIntIdx = i
		}

		if isFractional && i >= lenAbsIntRunes && absAllRunes[i] > '0' && absAllRunes[i] <= '9' {
			lastFracIdx = i
		}

	}

	if firstIntIdx == -1 {
		firstIntIdx = lastIntIdx
	}

	if isFractional && lastFracIdx == -1 {
		lastFracIdx = lenAbsIntRunes
	}

	numStrOut := ""

	if signVal < 0 {
		numStrOut = "-"
	}

	numStrOut += string(absAllRunes[firstIntIdx : lastIntIdx+1])
	if isFractional {
		numStrOut += string(numStrFormatSpec.GetDecimalSeparator())
		numStrOut += string(absAllRunes[lastIntIdx+1 : lastFracIdx+1])
	}

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto, err = nStrDtoAtom.parseNumStr(
		numStrOut,
		numStrFormatSpec,
		ePrefix)

	return newNumStrDto, err
}

// NewBigFloat - Creates a new NumStrDto instance from a Big Float value
// (*big.Float) and a precision specification.
//
// For more information on the *big.Float floating point numeric value,
// reference:
//   https://golang.org/pkg/math/big/
//
//
// See the 'Example Usage' section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  bigFloatNum         *big.Float
//     - A type *big.Float floating point numeric value. For details
//       on type *big.Float, reference:
//         https://golang.org/pkg/math/big/
//
//     This floating point numeric value will be converted to a new
//     instance of type NumStrDto.
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. The final value will be
//       rounded to 'precision' digits after the decimal point.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     numSepsDto := NumericSeparatorsDto{}
//     numSepsDto.SetToUSADefaults()
//
//     f64Num := float64(123.456)
//     bigFloatNum := big.NewFloat(f64Num)
//     precision := uint(2)
//     ePrefix := "calling method name "
//
//          nDto, err  :=
//              numStrDtoUtility.newBigFloat(
//              numSepsDto,
//              bigFloatNum,
//              precision,
//              ePrefix)
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//  bigFloatNum     precision        Result
//  -------------   --------------------------
//
//   12.3456            4               12.3456
//   123456.5           0           123457
//   1234.56            1             1234.6
//
func (nStrDtoNanobot *numStrDtoNanobot) newBigFloat(
	numStrFormatSpec *NumStrFmtSpecDto,
	bigFloatNum *big.Float,
	precision uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newBigFloat()")

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	if bigFloatNum == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigFloatNum' has a 'nil' pointer!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	numStr := bigFloatNum.Text('f', int(precision))

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numStrFormatSpec,
		ePrefix.XCtx("numStr"))

	return newNumStrDto, err
}

// newBigInt - receives a signed Big Integer number (type *big.Int)
// and  precision parameter. This method then proceeds to generate
// and return a new NumStrDto type encapsulating the numeric value
// of the passed Big Integer number.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  bigIntNum           *big.Int
//     - This numeric value will be converted to a new instance of
//       type NumStrDto. Type 'big.Int' is designed to handle very
//       large integer values. For more information on type 'big.Int',
//       reference:
//          https://golang.org/pkg/math/big/
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. If 'precision' has a value
//       greater than zero, the returned NumStrDto will be configured
//       as a floating point value.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example #1
//     int64Num := int64(123456)
//     bigIntNum := big.NewInt(int64Num)
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(3)
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//          nDto, err :=
//            nStrDtoNanobot.newBigInt(
//            numSeps,
//            bigIntNum,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Example #2
//     int64Num := int64(123456)
//     bigIntNum := big.NewInt(int64Num)
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(0)
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//          nDto, err :=
//            nStrDtoNanobot.newBigInt(
//            numSeps,
//            bigIntNum,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <--- Output --->
//                                 newNumStrDto
//  bigIntNum    precision            Result
//  ---------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//   123456          7                   0.0123456
//
func (nStrDtoNanobot *numStrDtoNanobot) newBigInt(
	numStrFormatSpec *NumStrFmtSpecDto,
	bigIntNum *big.Int,
	precision uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newBigInt()")

	err = nil
	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	if bigIntNum == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigIntNum' has a 'nil' pointer!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	err = nStrDtoElectron.setFormatSpec(
		&newNumStrDto,
		numStrFormatSpec,
		ePrefix.XCtx("numStrFormatSpec -> newNumStrDto"))

	if err != nil {
		return newNumStrDto, err
	}

	newNumStrDto.precision = precision
	scratchNum := big.NewInt(0).Set(bigIntNum)
	bigZero := big.NewInt(0)
	newNumStrDto.signVal = 1

	if scratchNum.Cmp(bigZero) == -1 {
		scratchNum.Neg(scratchNum)
		newNumStrDto.signVal = -1
	}

	bigTen := big.NewInt(int64(10))
	modulo := big.NewInt(0)
	newNumStrDto.absAllNumRunes =
		make([]rune, 0, 100)

	if scratchNum.Cmp(bigZero) == 0 {

		newNumStrDto.absAllNumRunes =
			append(newNumStrDto.absAllNumRunes,
				'0')

	} else {

		for scratchNum.Cmp(bigZero) == 1 {
			modulo = big.NewInt(0).Rem(scratchNum, bigTen)
			scratchNum = big.NewInt(0).Quo(scratchNum, bigTen)
			newNumStrDto.absAllNumRunes =
				append(newNumStrDto.absAllNumRunes,
					rune(modulo.Int64()+int64(48)))
		}
	}

	lenAllNumRunes :=
		len(newNumStrDto.absAllNumRunes)

	if int(newNumStrDto.precision) >=
		lenAllNumRunes {

		deltaNumRunes :=
			int(newNumStrDto.precision) - lenAllNumRunes + 1

		for k := 0; k < deltaNumRunes; k++ {
			newNumStrDto.absAllNumRunes =
				append(newNumStrDto.absAllNumRunes,
					'0')
			lenAllNumRunes++
		}

	}

	tRune := rune(0)

	if lenAllNumRunes > 1 {
		xLen := lenAllNumRunes - 1
		sortLimit := xLen / 2
		yCnt := 0
		for i := xLen; i > sortLimit; i-- {
			tRune =
				newNumStrDto.absAllNumRunes[yCnt]

			newNumStrDto.absAllNumRunes[yCnt] =
				newNumStrDto.absAllNumRunes[i]

			newNumStrDto.absAllNumRunes[i] = tRune

			yCnt++
		}
	}

	err = nStrDtoElectron.setFormatSpec(
		&newNumStrDto,
		numStrFormatSpec,
		ePrefix.XCtx("newNumStrDto #2"))

	if err != nil {
		return newNumStrDto, err
	}

	_,
		err =
		numStrDtoQuark{}.ptr().testNumStrDtoValidity(
			&newNumStrDto,
			ePrefix.XCtx("Final Validity Check-newNumStrDto"))

	return newNumStrDto, err
}

// newFloat64 - Creates a new NumStrDto instance from a float64
// and precision specification.
//
// See the 'Example Usage' section below.
//
//     - An instance of NumericSeparatorsDto which will be used to supply
//       the numeric separators for the new NumStrDto instance returned
//       by this method. Numeric separators include the Thousands
//       Separator, Decimal Separator and the Currency Symbol.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  f64                 float64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. The final value will be
//       rounded to 'precision' digits after the decimal point.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  numSepsDto := NumericSeparatorsDto{}
//  numSepsDto.SetToUSADefaults()
//
//           f64 := float64(123.456)
//     precision := uint(2)
//
//          nDto, err  :=
//              numStrDtoUtility.newFloat64(
//              numSepsDto,
//              f64,
//              precision,
//              "calling method name ")
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//    f64        precision           Result
//  ------------------------------------------
//
//   12.3456         4                  12.3456
//   123456.5        0              123457
//   1234.56         1                1234.6
//
func (nStrDtoNanobot *numStrDtoNanobot) newFloat64(
	numStrFormatSpec *NumStrFmtSpecDto,
	f64 float64,
	precision uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newFloat64()")

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	// Number string rounded to 'precision' decimal
	// places.
	numStr := strconv.FormatFloat(f64,
		'f',
		int(precision),
		64)

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numStrFormatSpec,
		ePrefix.XCtx(
			"numStr, numSepsDto"))

	return newNumStrDto, err
}

// newInt64 - Creates a new NumStrDto instance from an int64 and a
// precision specification.
//
//
// --------------------------------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  int64Num            int64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     int64Num := int64(123456)
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(3)
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//          nDto, err :=
//            nStrDtoNanobot.newInt64(
//            numSeps,
//            int64Num,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <--- Output --->
//                                 newNumStrDto
//  int64Num     precision            Result
//  ---------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//   123456          7                   0.0123456
//
func (nStrDtoNanobot *numStrDtoNanobot) newInt64(
	numStrFormatSpec *NumStrFmtSpecDto,
	int64Num int64,
	precision uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newInt64()")

	err = nil
	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.XCtxEmpty().String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	numStr := strconv.FormatInt(int64Num, 10)

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
		err = nStrDtoMolecule.setPrecision(
		numStrFormatSpec,
		numStr,
		precision,
		true,
		ePrefix.XCtx("numSepsDto -> newNumStrDto"))

	return newNumStrDto, err
}

// NewInt64Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an int64 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//    numeric value = int64 X 10^exponent
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  int64Num            int64
//     - This numeric value will be multiplied by 10^exponent and
//       converted to a new instance of type NumStrDto.
//
//  exponent            int
//     - 10^exponent is multiplied by input parameter 'int64Num' to
//       generate a new instance of type NumStrDto.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     int64Num := int64(123456)
//     exponent := -3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto, err := nStrDtoNanobot.newInt64Exponent(
//               int64Num,
//               exponent)
//
//  -- nDto is now equal to "123.456", precision = 3
//
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     int64Num := int64(123456)
//     exponent := 3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto, err := nStrDtoNanobot.newInt64Exponent(
//                 int64Num,
//                 exponent)
//
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//                                Decimal
//   int64Num    exponent          Result
//    123456        -3              123.456
//    123456         3           123456.000
//    123456         0           123456
//
func (nStrDtoNanobot *numStrDtoNanobot) newInt64Exponent(
	numStrFormatSpec *NumStrFmtSpecDto,
	int64Num int64,
	exponent int,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newInt64Exponent()")

	err = nil
	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	if numStrFormatSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	numStr := strconv.FormatInt(int64Num, 10)

	if exponent > 0 {
		for i := 0; i < exponent; i++ {
			numStr += "0"
		}
	}

	if exponent < 0 {
		exponent = exponent * -1
	}

	if exponent == 0 {
		nStrDtoAtom := numStrDtoAtom{}

		newNumStrDto,
			err = nStrDtoAtom.parseNumStr(
			numStr,
			numStrFormatSpec,
			ePrefix.XCtx(
				"numStr, numStrFormatSpec"))

	} else {

		nStrDtoMolecule := numStrDtoMolecule{}

		newNumStrDto,
			err = nStrDtoMolecule.shiftPrecisionLeft(
			numStrFormatSpec,
			numStr,
			uint(exponent),
			ePrefix.XCtx(
				"numStrFormatSpec, numStr"))

	}

	return newNumStrDto, err
}

// NewNumStr - Used to create a populated NumStrDto instance using a
// valid number string as an input parameter.
//
// A valid number string 'may' be prefixed with numeric sign value of
// plus ('+') or minus ('-'). The absence of a leading numeric sign
// character will default the numeric value to plus or a positive
// numeric value. A valid number string 'may' also include a decimal
// delimiter such as a decimal point to separate integer and fractional
// digits in the number string. With these two exceptions, all other
// characters in a valid number string must be text characters between
// '0' and '9'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance returned by this method will be
//       configured with this Number String Format Specification.
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
//  numStr              string
//     - A valid number string. A valid number string 'may' be
//       prefixed with a numeric sign value of plus ('+') or
//       minus ('-'). The absence of a leading numeric sign
//       character will default the numeric value to plus or a
//       positive numeric value. A valid number string 'may'
//       also include a decimal delimiter such as a decimal
//       point to separate integer and fractional digits
//       within the number string. With these two exceptions,
//       all other characters in a valid number string must be
//       numeric values represented by text characters between
//       '0' and '9'.
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
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
func (nStrDtoNanobot *numStrDtoNanobot) newNumStr(
	numStrFormatSpec *NumStrFmtSpecDto,
	numStr string,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newNumStr()")

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if len(numStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is a zero length string!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.XCtxEmpty())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numStrFormatSpec,
		ePrefix.XCtx(
			"numStr, numSepsDto"))

	return newNumStrDto, err
}

// newRational - Creates a new NumStrDto instance from a rational
// number and a precision specification.
//
// For information on Big Rational Numbers (*big.Rat), reference:
//    https://golang.org/pkg/math/big/
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  bigRatNum           *big.Rat
//     - This 'big' Rational Number will be converted into a
//       a returned instance of NumStrDto. The numeric value
//       of the big Rational Number will be represented as
//       a fractional or floating point number with a 'precision'
//       number of digits after the decimal point.
//
//       For more information on type *big.Rat, reference:
//         https://golang.org/pkg/math/big/
//
//
//  precision       uint
//     - The number of digits which will be placed to the right
//       of the decimal point in the returned new instance of
//       NumStrDto. This fractional floating point value will be
//       rounded to 'precision' digits.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
func (nStrDtoNanobot *numStrDtoNanobot) newRational(
	numStrFormatSpec *NumStrFmtSpecDto,
	bigRatNum *big.Rat,
	precision uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newRational()")

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if bigRatNum == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigRatNum' has a nil pointer!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.XCtxEmpty().String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	numStr := bigRatNum.FloatString(int(precision))

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numStrFormatSpec,
		ePrefix.XCtx(
			"numStr, numSeps"))

	return newNumStrDto, err
}

// NewUint64 - Creates a new NumStrDto instance from an uint64 and a
// precision specification.
//
// See the 'Example Usage' section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  uint64Num           uint64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     uint64Num := uint64(123456)
//     numSepsDto := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(3)
//
//          nDto, err :=
//            numStrDtoUtility.newUint64(
//            numSepsDto,
//            uint64Num,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <-- Output -->
//                                newNumStrDto
//  uint64Num     precision          Result
//  ------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//
func (nStrDtoNanobot *numStrDtoNanobot) newUint64(
	numStrFormatSpec *NumStrFmtSpecDto,
	uint64Num uint64,
	precision uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newUint64()")

	newNumStrDto =
		numStrDtoElectron{}.ptr().newBaseZeroNumStrDto(0)

	err = nil

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.XCtxEmpty().String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	numStr := strconv.FormatUint(uint64Num, 10)

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
		err = nStrDtoMolecule.setPrecision(
		numStrFormatSpec,
		numStr,
		precision,
		true,
		ePrefix.XCtx("numSepsDto -> newNumStrDto "))

	return newNumStrDto, err
}

// newUint64Exponent - Creates a new NumStrDto instance from a type uint64 and a
// precision specification.
//
// See the Example Usage section below.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  uint64Num           uint64
//     - This uint64 is multiplied by 10^exponent (input parameter exponent)
//       to calculate the final numeric value which is returned in a new
//       instance of NumStrDto.
//
//
//  exponent            int
//     - This is the exponent value. Input parameter uint64Num is multiplied
//       by 10 raised to the power of this 'exponent' parameter in order to
//       calculate the numeric value contained in the new instance of
//       NumStrDto returned by this method.
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
//  newNumStrDto        NumStrDto
//    - If this method completes successfully, it will return a new
//       instance of NumStrDto. The numeric value contained in this new
//       instance is calculated by multiplying input parameter
//       'uint64Num' times 10 raised to the power of input parameter
//       'exponent'.
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
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     uint64Num := uint64(123456)
//     exponent := -3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto,err := nStrDtoNanobot.newUint64Exponent(
//          numSeps,
//          uint64Num,
//          exponent,
//          "calling method name ")
//   -- nDto is now equal to "123.456", precision = 3
//
//
//     numSeps := NumericSeparatorsDto{}
//     numSeps.SetToUSADefaults()
//     uint64Num := uint64(123456)
//     exponent := 3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto,err := nStrDtoNanobot.newUint64Exponent(
//          numSeps,
//          uint64Num,
//          exponent,
//          "calling method name ")
//
//  -- nDto is now equal to "123456.000", precision = 3
//
//  Examples:
//  ---------
//
//  <---- Input Parameters ---->       <-- Output -->
//                                       newNumStrDto
//   uint64Num          exponent            Result
//   123456               -3               123.456
//   123456                3            123456.000
//   123456                0            123456
//
func (nStrDtoNanobot *numStrDtoNanobot) newUint64Exponent(
	numStrFormatSpec *NumStrFmtSpecDto,
	uint64Num uint64,
	exponent int,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.newUint64Exponent()")

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	numStr := strconv.FormatUint(uint64Num, 10)

	if exponent > 0 {
		for i := 0; i < exponent; i++ {
			numStr += "0"
		}
	}

	if exponent < 0 {
		exponent = exponent * -1
	}

	if exponent == 0 {

		nStrDtoAtom := numStrDtoAtom{}

		newNumStrDto,
			err = nStrDtoAtom.parseNumStr(
			numStr,
			numStrFormatSpec,
			ePrefix.XCtx("exponent == 0"))

	} else {

		nStrDtoMolecule := numStrDtoMolecule{}

		newNumStrDto,
			err = nStrDtoMolecule.shiftPrecisionLeft(
			numStrFormatSpec,
			numStr,
			uint(exponent),
			ePrefix.XCtx("numStr"))

	}

	return newNumStrDto, err
}

// ptr - Returns a pointer to a new instance of numStrDtoAtom.
//
func (nStrDtoNanobot numStrDtoNanobot) ptr() *numStrDtoNanobot {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	newNumStrDtoNanobot := new(numStrDtoNanobot)

	newNumStrDtoNanobot.lock = new(sync.Mutex)

	return newNumStrDtoNanobot
}

// scaleNumStr - Receives a signed number string and proceeds
// to shifts the position of the decimal point left or right
// depending on the value of input parameter 'scaleMode'.
//
// The scaled number string will then be converted to a new
// instance of NumStrDto and returned to the caller.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatSpec    *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto.
//
//       The NumStrDto instance ('newNumStrDto') returned by this
//       method will be configured with this Number String Format
//       Specification.
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
//  signedNumStr        string
//     - A valid number string. The leading digit may optionally
//       be a '+' or '-' indicating numeric sign value. If '+'
//       or '-' characters are not present in the first character
//       position, the number is assumed to represent a positive
//       numeric value ('+'). In addition to leading plus or minus
//       characters, the number string may contain a decimal point
//       separating integer and fractional digits. All other
//       characters in this number string must be numeric digits.
//
//
//  shiftPrecision      uint
//     - The number of positions which the decimal point will be
//       shifted. If 'shiftPrecision is Equal to zero, no action
//       will be taken, no error will be issued and the original
//       signedNumStr will be converted to a NumStrDto instance
//       and returned to the caller.
//
//
//  scaleMode           PrecisionScaleMode
//     - A constant with one of two Scale Mode values. These
//       constant values are located in source code file:
//             datetime/numstrdtoconstants.go
//
//       SCALEPRECISIONLEFT - Shifts the decimal point
//                            from its current position to the left.
//
//       SCALEPRECISIONRIGHT - Shifts the decimal point from its current
//                             position to the right.
//
//       Note: See Methods numStrDtoMolecule.shiftPrecisionRight() and
//       numStrDtoMolecule.shiftPrecisionLeft() for additional
//       information.
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
//  newNumStrDto        NumStrDto
//     - If this method completes successfully, the result of the numeric
//       scaling operation performed by this method will be returned in
//       the form of a new 'NumStrDto' instance.
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
func (nStrDtoNanobot *numStrDtoNanobot) scaleNumStr(
	numStrFormatSpec *NumStrFmtSpecDto,
	signedNumStr string,
	shiftPrecision uint,
	scaleMode PrecisionScaleMode,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.scaleNumStr()")

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	if len(signedNumStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumStr' is INVALID!\n"+
			"'signedNumStr' is a zero length number string!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	if numStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFormatSpec' is invalid!\n"+
			"numStrFormatSpec is a 'nil' pointer.\n",
			ePrefix.XCtxEmpty().String())
		return newNumStrDto, err
	}

	err = numStrFormatSpec.IsValidInstanceError(
		ePrefix.XCtx("numStrFormatSpec"))

	if err != nil {
		return newNumStrDto, err
	}

	err = nStrDtoElectron.setFormatSpec(
		&newNumStrDto,
		numStrFormatSpec,
		ePrefix.XCtx("Setting 'newNumStrDto' numeric separators"))

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoMolecule := numStrDtoMolecule{}

	var err2 error

	if scaleMode == SCALEPRECISIONLEFT {

		newNumStrDto,
			err2 = nStrDtoMolecule.shiftPrecisionLeft(
			numStrFormatSpec,
			signedNumStr,
			shiftPrecision,
			ePrefix.XCtx(
				"numSeparators, signedNumStr"))

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error returned from nStrDtoMolecule.ShiftPrecisionLeft"+
				"(signedNumStr, shiftPrecision)\n"+
				"signedNumStr='%v'\n"+
				"shiftPrecision='%v'\n"+
				"scaleMode='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				signedNumStr,
				shiftPrecision,
				scaleMode.String(),
				err2.Error())

			return newNumStrDto, err
		}

	} else if scaleMode == SCALEPRECISIONRIGHT {

		newNumStrDto, err2 =
			nStrDtoMolecule.shiftPrecisionRight(
				numStrFormatSpec,
				signedNumStr,
				shiftPrecision,
				ePrefix.XCtx(
					"numSeparators, signedNumStr"))

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error returned from nStrDtoMolecule.ShiftPrecisionRight"+
				"(signedNumStr, shiftPrecision)\n"+
				"signedNumStr='%v'\n"+
				"shiftPrecision='%v'\n"+
				"scaleMode='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				signedNumStr,
				shiftPrecision,
				scaleMode.String(),
				err2.Error())

			return newNumStrDto, err
		}

	} else {

		err = fmt.Errorf("%v\n"+
			"Error! Scale Mode is INVALID!\n"+
			"Scale Mode is NOT Equal to SCALEPRECISIONLEFT or SCALEPRECISIONRIGHT.\n"+
			"scaleMode='%v'\n",
			ePrefix.String(),
			scaleMode.String())

		return newNumStrDto, err
	}

	return newNumStrDto, err
}

// setNumStrDtoPrecision - Sets or resets the precision for a passed instance
// of NumStrDto (Input parameter 'numStrDto').
//
// Input parameter 'precision' identifies the number of decimal places to the
// right of the decimal point which will be configured in 'numStrDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE the values of internal member variables to achieve
//       the method's objectives. Member variables will be tested for
//       validity.
//
//       This method will set or reset the 'precision' of the numeric
//       value encapsulated by this instance of NumStrDto. 'precision',
//       as defined here, specifies the number of digits to the right
//       of the decimal point which will be formatted in the numeric
//       value encapsulated in parameter, 'numStrDto'.
//
//
//  precision           uint
//     - The number of numeric digits to the right of the decimal place
//       which will be configured in the numeric value encapsulated within
//       input parameter 'numStrDto'.
//
//
//  roundResult         bool
//     - If the 'precision' value is less than the current number of places
//       to the	right of the decimal point, this method will truncate the
//       existing fractional digits. If 'roundResult' is set to true, this
//       truncation operation will include rounding the last digit.
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
func (nStrDtoNanobot *numStrDtoNanobot) setNumStrDtoPrecision(
	numStrDto *NumStrDto,
	precision uint,
	roundResult bool,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoNanobot.setNumStrDtoPrecision()")

	err = nil

	var numStr string

	nStrDtoMolecule := numStrDtoMolecule{}

	numStr,
		err = nStrDtoMolecule.getNumStr(
		numStrDto,
		ePrefix.XCtx("Input parameter 'numStrDto' "))

	if err != nil {
		return err
	}

	nStrDtoElectron := numStrDtoElectron{}

	var numStrFormatSpec NumStrFmtSpecDto

	numStrFormatSpec,
		err =
		nStrDtoElectron.getFormatSpec(
			numStrDto,
			ePrefix.XCtx("numStrDto -> numSepsDto"))

	if err != nil {
		return err
	}

	var n2 NumStrDto

	n2,
		err = nStrDtoMolecule.setPrecision(
		&numStrFormatSpec,
		numStr,
		precision,
		roundResult,
		ePrefix.XCtx("numStr"))

	if err != nil {
		return err
	}

	err = nStrDtoElectron.setFormatSpec(
		&n2,
		&numStrFormatSpec,
		ePrefix.XCtx("n2"))

	if err != nil {
		return err
	}

	err = nStrDtoElectron.copyIn(
		numStrDto,
		&n2,
		ePrefix.XCtx("n2 -> numStrDto"))

	if err != nil {
		return err
	}

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			numStrDto,
			ePrefix.XCtx(
				"numStrDto"))

	return err
}
