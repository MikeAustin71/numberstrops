package numberstr

import (
	"fmt"
	"sync"
)

type numStrDtoMechanics struct {
	lock *sync.Mutex
}

// findIntArraySignificantDigitLimits - Receives an array of integers and
// converts them to a number string consisting of significant digits.
// Leading and trailing zeros are eliminated.
//
// See related method  numStrDtoNanobot.findNumStrSignificantDigitLimits()
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
//         currencyValue  FormatterCurrency
//         signedNumValue NumStrFmtSpecSignedNumValueDto
//         sciNotation    NumStrFmtSpecSciNotationDto
//       }
//
//
//  intArray            []int
//     - An array of integer representing single numeric digits in a number
//       string. If any member values are less than zero ('0'), an error will
//       be returned. Likewise, if any member values are greater than nine
//       ('9'), an error will be returned.
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
func (nStrDtoMech *numStrDtoMechanics) findIntArraySignificantDigitLimits(
	numStrFormatSpec *NumStrFmtSpecDto,
	intArray []int,
	precision uint,
	signVal int,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoMech.lock == nil {
		nStrDtoMech.lock = new(sync.Mutex)
	}

	nStrDtoMech.lock.Lock()

	defer nStrDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoMechanics.findIntArraySignificantDigitLimits()")

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

	ePrefix.XCtxEmpty()

	lenIntArray := len(intArray)
	//
	//if lenIntArray == 0 {
	//	newNumStrDto = nStrDtoElectron.newBaseZeroNumStrDto(
	//		precision)
	//
	//	return newNumStrDto, err
	//}

	absNumStr := make([]rune, 0, 20)

	for i := 0; i < lenIntArray; i++ {

		if intArray[i] < 0 ||
			intArray[i] > 9 {

			err = fmt.Errorf("%v\n"+
				"Error: Input parameter 'intArray' contains an invalid value!\n"+
				"Members of 'intArray' must be integer values between '0' and '9', inclusive.\n"+
				"intArray[%v]='%v'\n",
				ePrefix.String(),
				i,
				intArray[i])

			return newNumStrDto, err
		}

		absNumStr = append(absNumStr, rune(intArray[i]+48))
	}

	nStrDtoNanobot := numStrDtoNanobot{}

	newNumStrDto,
		err = nStrDtoNanobot.findNumStrSignificantDigitLimits(
		numStrFormatSpec,
		absNumStr,
		precision,
		signVal,
		ePrefix.XCtx("numSepsDto, absNumStr"))

	return newNumStrDto, err
}
