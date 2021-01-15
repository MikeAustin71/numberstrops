package numberstr

import "sync"

type NumStrFmtSpecCurrencyValueDto struct {
	positiveValueFmt              string
	negativeValueFmt              string
	decimalDigits                 int
	currencyCode                  string
	currencyName                  string
	currencySymbol                rune
	turnOnIntegerDigitsSeparation bool
	numberSeparators              NumStrFmtSpecDigitsSeparatorsDto
	numFieldLen                   NumberFieldDto
	lock                          *sync.Mutex
}

// SetPositiveValueFormat - Sets the currency positive value format
// string. This strings contains the formatting symbols used to
// format positive currency values in number strings.
//
// If the positive value string input parameter, 'positiveValueFmt',
// is invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt    string
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCurrValDto *NumStrFmtSpecCurrencyValueDto) SetPositiveValueFormat(
	positiveValueFmt string,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValDto.lock == nil {
		nStrFmtSpecCurrValDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValDto.lock.Lock()

	defer nStrFmtSpecCurrValDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecCurrencyValueDto.SetPositiveValueFormat()\n "

	nStrCurrencyElectron :=
		numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix+
			"\nTesting validity of 'positiveValueFmt'\n ")

	if err != nil {
		return err
	}

	nStrFmtSpecCurrValDto.positiveValueFmt =
		positiveValueFmt

	return err
}
