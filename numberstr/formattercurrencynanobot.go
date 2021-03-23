package numberstr

import (
	"fmt"
	"sync"
)

type formatterCurrencyNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecCurrencyValDto' to input parameter
// 'targetNStrFmtSpecCurrencyValDto'.
//
// Be advised - All data fields in 'targetFormatterCurrency'
// will be overwritten.
//
// If input parameter 'inComingFormatterCurrency' is judged
// to be invalid, this method will return an error.
//
func (fmtCurrNanobot *formatterCurrencyNanobot) copyIn(
	targetFormatterCurrency *FormatterCurrency,
	inComingFormatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrNanobot.lock == nil {
		fmtCurrNanobot.lock = new(sync.Mutex)
	}

	fmtCurrNanobot.lock.Lock()

	defer fmtCurrNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyNanobot.copyIn()")

	if targetFormatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFormatterCurrency' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if inComingFormatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingFormatterCurrency' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		formatterCurrencyMolecule{}.ptr().
			testValidityOfFormatterCurrency(
				inComingFormatterCurrency,
				ePrefix.XCtx("Testing validity of 'inComingFormatterCurrency'"))

	if err != nil {
		return err
	}

	lenCurrencySymbols := len(inComingFormatterCurrency.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'inComingFormatterCurrency.currencySymbols' invalid!\n"+
			"The currencySymbols array is a zero length array.\n",
			ePrefix.XCtxEmpty().String())

		return err
	}

	targetFormatterCurrency.numStrFmtType =
		inComingFormatterCurrency.numStrFmtType

	targetFormatterCurrency.positiveValueFmt =
		inComingFormatterCurrency.positiveValueFmt

	targetFormatterCurrency.negativeValueFmt =
		inComingFormatterCurrency.negativeValueFmt

	targetFormatterCurrency.decimalDigits =
		inComingFormatterCurrency.decimalDigits

	targetFormatterCurrency.currencyCode =
		inComingFormatterCurrency.currencyCode

	targetFormatterCurrency.currencyCodeNo =
		inComingFormatterCurrency.currencyCodeNo

	targetFormatterCurrency.currencyName =
		inComingFormatterCurrency.currencyName

	targetFormatterCurrency.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		targetFormatterCurrency.currencySymbols,
		inComingFormatterCurrency.currencySymbols)

	targetFormatterCurrency.minorCurrencyName =
		inComingFormatterCurrency.minorCurrencyName

	lenCurrencySymbols = len(inComingFormatterCurrency.minorCurrencySymbols)

	targetFormatterCurrency.minorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		targetFormatterCurrency.minorCurrencySymbols,
		inComingFormatterCurrency.minorCurrencySymbols)

	targetFormatterCurrency.turnOnIntegerDigitsSeparation =
		inComingFormatterCurrency.turnOnIntegerDigitsSeparation

	err =
		targetFormatterCurrency.numericSeparators.CopyIn(
			&inComingFormatterCurrency.numericSeparators,
			ePrefix.XCtx("'inComingFormatterCurrency' -> "+
				"'targetFormatterCurrency'"))

	if err != nil {
		return err
	}

	err =
		targetFormatterCurrency.numFieldDto.CopyIn(
			&inComingFormatterCurrency.numFieldDto,
			ePrefix.XCtx("\ninComingFormatterCurrency->targetFormatterCurrency"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecCurrencyValDto' styled as a new instance
// of FormatterCurrency.
//
// If input parameter 'formatterCurrency' is judged to be
// invalid, this method will return an error.
//
func (fmtCurrNanobot *formatterCurrencyNanobot) copyOut(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	newFormatterCurrency FormatterCurrency,
	err error) {

	if fmtCurrNanobot.lock == nil {
		fmtCurrNanobot.lock = new(sync.Mutex)
	}

	fmtCurrNanobot.lock.Lock()

	defer fmtCurrNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyNanobot.copyOut()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newFormatterCurrency, err
	}

	_,
		err =
		formatterCurrencyMolecule{}.ptr().
			testValidityOfFormatterCurrency(
				formatterCurrency,
				ePrefix.XCtx("Testing validity of 'formatterCurrency'"))

	if err != nil {
		return newFormatterCurrency, err
	}

	lenCurrencySymbols :=
		len(formatterCurrency.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'formatterCurrency.currencySymbols' invalid!\n"+
			"The currencySymbols array is a zero length array.\n",
			ePrefix.XCtxEmpty().String())

		return newFormatterCurrency, err
	}

	newFormatterCurrency.numStrFmtType =
		formatterCurrency.numStrFmtType

	newFormatterCurrency.positiveValueFmt =
		formatterCurrency.positiveValueFmt

	newFormatterCurrency.negativeValueFmt =
		formatterCurrency.negativeValueFmt

	newFormatterCurrency.decimalDigits =
		formatterCurrency.decimalDigits

	newFormatterCurrency.currencyCode =
		formatterCurrency.currencyCode

	newFormatterCurrency.currencyCodeNo =
		formatterCurrency.currencyCodeNo

	newFormatterCurrency.currencyName =
		formatterCurrency.currencyName

	newFormatterCurrency.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		newFormatterCurrency.currencySymbols,
		formatterCurrency.currencySymbols)

	newFormatterCurrency.minorCurrencyName =
		formatterCurrency.minorCurrencyName

	lenCurrencySymbols =
		len(formatterCurrency.minorCurrencySymbols)

	newFormatterCurrency.minorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		newFormatterCurrency.minorCurrencySymbols,
		formatterCurrency.minorCurrencySymbols)

	newFormatterCurrency.turnOnIntegerDigitsSeparation =
		formatterCurrency.turnOnIntegerDigitsSeparation

	err = newFormatterCurrency.numericSeparators.CopyIn(
		&formatterCurrency.numericSeparators,
		ePrefix.XCtx("formatterCurrency->newFormatterCurrency"))

	if err != nil {
		return newFormatterCurrency, err
	}

	err =
		newFormatterCurrency.numFieldDto.CopyIn(
			&formatterCurrency.numFieldDto,
			ePrefix.XCtx("formatterCurrency->newFormatterCurrency"))

	newFormatterCurrency.lock = new(sync.Mutex)

	return newFormatterCurrency, err
}

// ptr - Returns a pointer to a new instance of
// formatterCurrencyNanobot.
//
func (fmtCurrNanobot formatterCurrencyNanobot) ptr() *formatterCurrencyNanobot {

	if fmtCurrNanobot.lock == nil {
		fmtCurrNanobot.lock = new(sync.Mutex)
	}

	fmtCurrNanobot.lock.Lock()

	defer fmtCurrNanobot.lock.Unlock()

	newFmtCurrencyNanobot :=
		new(formatterCurrencyNanobot)

	newFmtCurrencyNanobot.lock = new(sync.Mutex)

	return newFmtCurrencyNanobot
}

// setCurrencyData - Sets the basic currency data elements of
// an instance of FormatterCurrency. The only two
// elements of a FormatterCurrency instance which are
// NOT set by this method are the NumericSeparators and the
// NumberFieldDto data elements.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency             *FormatterCurrency,
//
//     - A pointer to an instance of FormatterCurrency.
//       All of the currency data values in this object will be
//       overwritten and set to new values based on the following
//       input parameters.
//
//
//  positiveValueFmt              string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
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
//       Valid format strings for positive currency values are
//       listed as follows:
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
//  negativeValueFmt              string
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
//  decimalDigits                 uint
//     - The standard number of digits to the right of the decimal
//       place which is expected for currency values. In the United
//       States, currency is typically formatted with two digits to
//       the right of the decimal.
//         Example:  $24.92
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbols                []rune
//     - The authorized character symbol associated with this
//       currency specification. In the United States, the currency
//       symbol is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
//
//  minorCurrencyName             string
//     - The minor currency name. In the United States, the minor
//       currency name is 'Cent'.
//
//
//  minorCurrencySymbols            []rune
//     - These are the unicode characters for minor currency
//       symbols. In the United States, the minor currency symbol
//       is the cent sign (¢), represented by a single unicode
//       character ('\U000000a2'). Some countries and cultures have
//       currency symbols consisting of two or more characters.
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnIntegerDigitsSeparation' is a boolean
//       flag used to control the 'Thousands Separator'. When set
//       to 'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  ePrefix             *ErrPrefixDto
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
//  err                           error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (fmtCurrNanobot *formatterCurrencyNanobot) setCurrencyData(
	formatterCurrency *FormatterCurrency,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	turnOnIntegerDigitsSeparation bool,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrNanobot.lock == nil {
		fmtCurrNanobot.lock = new(sync.Mutex)
	}

	fmtCurrNanobot.lock.Lock()

	defer fmtCurrNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyNanobot." +
			"setCurrencyData()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	lenCurrencySymbols := len(currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"currencySymbols is a zero length rune array!\n",
			ePrefix.String())

		return err
	}

	nStrCurrencyElectron :=
		formatterCurrencyElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx(
			fmt.Sprintf("Input parameter"+
				" 'positiveValueFmt' = '%v'",
				positiveValueFmt)))

	if err != nil {
		return err
	}

	_,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		positiveValueFmt,
		ePrefix.XCtx(
			fmt.Sprintf("Input parameter"+
				" 'negativeValueFmt' = '%v'",
				positiveValueFmt)))

	if err != nil {
		return err
	}

	formatterCurrency.numStrFmtType =
		NumStrFormatTypeCode(0).Currency()

	formatterCurrency.positiveValueFmt =
		positiveValueFmt

	formatterCurrency.negativeValueFmt =
		negativeValueFmt

	formatterCurrency.decimalDigits =
		decimalDigits

	formatterCurrency.currencyCode =
		currencyCode

	formatterCurrency.currencyCodeNo =
		currencyCodeNo

	formatterCurrency.currencyName =
		currencyName

	formatterCurrency.currencySymbols =
		make([]rune, lenCurrencySymbols)

	copy(formatterCurrency.currencySymbols,
		currencySymbols)

	formatterCurrency.minorCurrencyName =
		minorCurrencyName

	lenMinorCurrSymbols := len(minorCurrencySymbols)

	if lenMinorCurrSymbols == 0 {

		formatterCurrency.minorCurrencySymbols = nil

	} else {

		formatterCurrency.minorCurrencySymbols =
			make([]rune, lenMinorCurrSymbols)

		copy(formatterCurrency.minorCurrencySymbols,
			minorCurrencySymbols)
	}

	formatterCurrency.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	return err
}
