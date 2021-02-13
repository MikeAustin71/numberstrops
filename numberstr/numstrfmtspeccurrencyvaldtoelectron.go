package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoElectron struct {
	lock *sync.Mutex
}

// testCurrencyPositiveValueFormatStr - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValFmtStr             string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for currency values are listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  ePrefix                       *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid                       bool
//     - If the input parameter 'positiveValFmtStr' is judged to be
//       valid, 'isValid' will be set to true. If 'positiveValFmtStr'
//       is invalid, 'isValid' will be set to false.
//
//
//  err                           error
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
//     - If the input parameter 'positiveValFmtStr' is judged to be
//       valid, 'err' will be set to 'nil'. If 'positiveValFmtStr'
//       is invalid, 'err' will be returned with an appropriate error
//       message.
//
func (nStrCurrencyElectron *numStrFmtSpecCurrencyValueDtoElectron) testCurrencyPositiveValueFormatStr(
	positiveValFmtStr string,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrCurrencyElectron.lock == nil {
		nStrCurrencyElectron.lock = new(sync.Mutex)
	}

	nStrCurrencyElectron.lock.Lock()

	defer nStrCurrencyElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoElectron.testCurrencyPositiveValueFormatStr()")

	isValid = false

	if len(positiveValFmtStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValFmtStr' is an empty string!\n"+
			"The Currency Positive Value Format string is missing.\n"+
			"len(positiveValFmtStr) == 0\n",
			ePrefix.String())
		return isValid, err
	}

	nStrFmtSpecCurrValDtoQuark := numStrFmtSpecCurrencyValueDtoQuark{}

	validFmtChars :=
		nStrFmtSpecCurrValDtoQuark.getValidCurrencyPositiveValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(positiveValFmtStr)
	var lenCurrFmt = len(runesToTest)
	var isRuneValid bool

	for i := 0; i < lenCurrFmt; i++ {

		isRuneValid = false

		for j := 0; j < lenValidFmtChars; j++ {

			if runesToTest[i] != validFmtChars[j] {
				continue
			} else {
				isRuneValid = true
				break
			}
		}

		if !isRuneValid {
			isValid = false
			err = fmt.Errorf("%v\n"+
				"Error: The Number String Format contains an invalid character!\n"+
				"Signed Number Positive Value Formats are NOT allowed to include this character.\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix.String(),
				positiveValFmtStr,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		positiveValFmtStr, "127.54") &&
		!strings.Contains(
			positiveValFmtStr, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Value Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			positiveValFmtStr)

		return isValid, err
	}

	if !strings.Contains(positiveValFmtStr, "$") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Format String is missing a currency placeholder!\n"+
			"Currency Formats require a '$' place holder symbol in order to place\n"+
			"the actual currency in the correction position within a number string.\n"+
			"This Currency Positive Value Format String contains NO currency\n"+
			"placeholder symbol!\n",
			ePrefix)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testCurrencyNegativeValueFormatStr - Inspects the Negative Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  negativeValFmtStr             string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  ePrefix                       *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid                       bool
//     - If the input parameter 'positiveValFmtStr' is judged to be
//       valid, 'isValid' will be set to true. If 'positiveValFmtStr'
//       is invalid, 'isValid' will be set to false.
//
//
//  err                           error
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
//     - If the input parameter 'positiveValFmtStr' is judged to be
//       valid, 'err' will be set to 'nil'. If 'positiveValFmtStr'
//       is invalid, 'err' will be returned with an appropriate error
//       message.
//
func (nStrCurrencyElectron *numStrFmtSpecCurrencyValueDtoElectron) testCurrencyNegativeValueFormatStr(
	negativeValFmtStr string,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrCurrencyElectron.lock == nil {
		nStrCurrencyElectron.lock = new(sync.Mutex)
	}

	nStrCurrencyElectron.lock.Lock()

	defer nStrCurrencyElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoElectron.testCurrencyNegativeValueFormatStr()")

	isValid = false

	if len(negativeValFmtStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValFmtStr' is an empty string!\n"+
			"len(negativeValFmtStr) == 0\n",
			ePrefix.String())
		return isValid, err
	}

	nStrFmtSpecCurrValDtoQuark :=
		numStrFmtSpecCurrencyValueDtoQuark{}

	validFmtChars :=
		nStrFmtSpecCurrValDtoQuark.getValidCurrencyNegativeValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(negativeValFmtStr)
	var lenCurrFmt = len(runesToTest)
	var isRuneValid bool

	for i := 0; i < lenCurrFmt; i++ {

		isRuneValid = false

		for j := 0; j < lenValidFmtChars; j++ {

			if runesToTest[i] != validFmtChars[j] {
				continue
			} else {
				isRuneValid = true
				break
			}
		}

		if !isRuneValid {
			isValid = false
			err = fmt.Errorf("%v\n"+
				"Error: The Number String Format contains an invalid character!\n"+
				"Currency Negative Value Formats are NOT allowed to include this character.\n"+
				"Complete Currency Negative Value String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				negativeValFmtStr,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		negativeValFmtStr, "127.54") &&
		!strings.Contains(
			negativeValFmtStr, "NUMFIELD") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			negativeValFmtStr)

		return isValid, err
	}

	if !strings.Contains(
		negativeValFmtStr, ")") &&
		!strings.Contains(
			negativeValFmtStr, "(") &&
		!strings.Contains(
			negativeValFmtStr, "-") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a negative sign\n"+
			"place holder. The format string MUST contain either a minus\n"+
			"sign '-' or parenthesis '()' to designate a negative value.\n"+
			"This Number String Format does NOT contain a negative value placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			negativeValFmtStr)

		return isValid, err
	}

	if !strings.Contains(negativeValFmtStr, "$") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Format String is missing a currency placeholder!\n"+
			"Currency Formats require a '$' place holder symbol in order to place\n"+
			"the actual currency in the correction position within a number string.\n"+
			"This Currency Negative Value Format String contains NO currency\n"+
			"placeholder symbol!\n",
			ePrefix)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
