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
//  numSepsDto          NumericSeparatorDto
//     - An instance of NumericSeparatorDto which will be used to supply
//       the numeric separators for the new NumStrDto instance returned
//       by this method. Numeric separators include the Thousands
//       Separator, Decimal Separator and the Currency Symbol.
//
//       The data fields included in the NumericSeparatorDto are
//       listed as follows:
//
//          type NumericSeparatorDto struct {
//
//            DecimalSeparator   rune // Character used to separate
//                                    //  integer and fractional digits ('.')
//
//            ThousandsSeparator rune // Character used to separate thousands
//                                    //  (1,000,000,000
//
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//       If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), they will
//       be reset to USA default values. USA default numeric
//       separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
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
//  err                error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will encapsulate
//       an error message. Note that this error message will
//       incorporate the method chain and text passed by input
//       parameter, 'ePrefix'. The 'ePrefix' text will be prefixed
//       at the beginning of the returned error message.
//
func (nStrDtoMech *numStrDtoMechanics) findIntArraySignificantDigitLimits(
	numSepsDto NumericSeparatorDto,
	intArray []int,
	precision uint,
	signVal int,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoMech.lock == nil {
		nStrDtoMech.lock = new(sync.Mutex)
	}

	nStrDtoMech.lock.Lock()

	defer nStrDtoMech.lock.Unlock()

	ePrefix += "numStrDtoMechanics.findIntArraySignificantDigitLimits() "

	lenIntArray := len(intArray)

	absNumStr := make([]rune, 0, 20)

	for i := 0; i < lenIntArray; i++ {

		if intArray[i] < 0 ||
			intArray[i] > 9 {

			err = fmt.Errorf(ePrefix+"\n"+
				"Error: Input parameter 'intArray' contains an invalid value!\n"+
				"Members of 'intArray' must be integer values between '0' and '9', inclusive.\n"+
				"intArray[%v]='%v'\n",
				i,
				intArray[i])

			return newNumStrDto, err
		}

		absNumStr = append(absNumStr, rune(intArray[i]+48))
	}

	numSepsDto.SetToUSADefaults()

	nStrDtoNanobot := numStrDtoNanobot{}

	newNumStrDto,
		err = nStrDtoNanobot.findNumStrSignificantDigitLimits(
		numSepsDto,
		absNumStr,
		precision,
		signVal,
		ePrefix)

	return newNumStrDto, err
}
