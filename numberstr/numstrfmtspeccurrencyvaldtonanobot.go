package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecCurrencyValDto' to input parameter
// 'targetNStrFmtSpecCurrencyValDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecCurrencyValDto'
// will be overwritten.
//
// If input parameter 'inComingNStrFmtSpecCurrencyValDto' is judged
// to be invalid, this method will return an error.
//
func (nStrFmtSpecCurrValNanobot *numStrFmtSpecCurrencyValueDtoNanobot) copyIn(
	targetNStrFmtSpecCurrencyValDto *FormatterCurrency,
	inComingNStrFmtSpecCurrencyValDto *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoNanobot.copyIn()")

	if targetNStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if inComingNStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	nStrFmtSpecCurrDtoMolecule :=
		formatterCurrencyMolecule{}

	_,
		err =
		nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
			inComingNStrFmtSpecCurrencyValDto,
			ePrefix.XCtx("Testing validity of 'inComingNStrFmtSpecCurrencyValDto'"))

	if err != nil {
		return err
	}

	lenCurrencySymbols := len(inComingNStrFmtSpecCurrencyValDto.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'inComingNStrFmtSpecCurrencyValDto.currencySymbols' invalid!\n"+
			"The currencySymbols array is a zero length array.\n",
			ePrefix.XCtxEmpty().String())

		return err
	}

	targetNStrFmtSpecCurrencyValDto.positiveValueFmt =
		inComingNStrFmtSpecCurrencyValDto.positiveValueFmt

	targetNStrFmtSpecCurrencyValDto.negativeValueFmt =
		inComingNStrFmtSpecCurrencyValDto.negativeValueFmt

	targetNStrFmtSpecCurrencyValDto.decimalDigits =
		inComingNStrFmtSpecCurrencyValDto.decimalDigits

	targetNStrFmtSpecCurrencyValDto.currencyCode =
		inComingNStrFmtSpecCurrencyValDto.currencyCode

	targetNStrFmtSpecCurrencyValDto.currencyCodeNo =
		inComingNStrFmtSpecCurrencyValDto.currencyCodeNo

	targetNStrFmtSpecCurrencyValDto.currencyName =
		inComingNStrFmtSpecCurrencyValDto.currencyName

	targetNStrFmtSpecCurrencyValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		targetNStrFmtSpecCurrencyValDto.currencySymbols,
		inComingNStrFmtSpecCurrencyValDto.currencySymbols)

	targetNStrFmtSpecCurrencyValDto.minorCurrencyName =
		inComingNStrFmtSpecCurrencyValDto.minorCurrencyName

	lenCurrencySymbols = len(inComingNStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	targetNStrFmtSpecCurrencyValDto.minorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		targetNStrFmtSpecCurrencyValDto.minorCurrencySymbols,
		inComingNStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	targetNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		inComingNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation

	err =
		targetNStrFmtSpecCurrencyValDto.numericSeparators.CopyIn(
			&inComingNStrFmtSpecCurrencyValDto.numericSeparators,
			ePrefix.XCtx("'inComingNStrFmtSpecCurrencyValDto' -> "+
				"'targetNStrFmtSpecCurrencyValDto'"))

	if err != nil {
		return err
	}

	err =
		targetNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
			&inComingNStrFmtSpecCurrencyValDto.numFieldLenDto,
			ePrefix.XCtx("\ninComingNStrFmtSpecCurrencyValDto->targetNStrFmtSpecCurrencyValDto"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecCurrencyValDto' styled as a new instance
// of FormatterCurrency.
//
// If input parameter 'nStrFmtSpecCurrencyValDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecCurrValNanobot *numStrFmtSpecCurrencyValueDtoNanobot) copyOut(
	nStrFmtSpecCurrencyValDto *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	newNStrFmtSpecCurrencyValDto FormatterCurrency,
	err error) {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoNanobot.copyOut()")

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNStrFmtSpecCurrencyValDto, err
	}

	nStrFmtSpecCurrDtoMolecule :=
		formatterCurrencyMolecule{}

	_,
		err =
		nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
			nStrFmtSpecCurrencyValDto,
			ePrefix.XCtx("Testing validity of 'nStrFmtSpecCurrencyValDto'"))

	if err != nil {
		return newNStrFmtSpecCurrencyValDto, err
	}

	lenCurrencySymbols :=
		len(nStrFmtSpecCurrencyValDto.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'nStrFmtSpecCurrencyValDto.currencySymbols' invalid!\n"+
			"The currencySymbols array is a zero length array.\n",
			ePrefix.XCtxEmpty().String())

		return newNStrFmtSpecCurrencyValDto, err
	}

	newNStrFmtSpecCurrencyValDto.positiveValueFmt =
		nStrFmtSpecCurrencyValDto.positiveValueFmt

	newNStrFmtSpecCurrencyValDto.negativeValueFmt =
		nStrFmtSpecCurrencyValDto.negativeValueFmt

	newNStrFmtSpecCurrencyValDto.decimalDigits =
		nStrFmtSpecCurrencyValDto.decimalDigits

	newNStrFmtSpecCurrencyValDto.currencyCode =
		nStrFmtSpecCurrencyValDto.currencyCode

	newNStrFmtSpecCurrencyValDto.currencyCodeNo =
		nStrFmtSpecCurrencyValDto.currencyCodeNo

	newNStrFmtSpecCurrencyValDto.currencyName =
		nStrFmtSpecCurrencyValDto.currencyName

	newNStrFmtSpecCurrencyValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		newNStrFmtSpecCurrencyValDto.currencySymbols,
		nStrFmtSpecCurrencyValDto.currencySymbols)

	newNStrFmtSpecCurrencyValDto.minorCurrencyName =
		nStrFmtSpecCurrencyValDto.minorCurrencyName

	lenCurrencySymbols =
		len(nStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	newNStrFmtSpecCurrencyValDto.minorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		newNStrFmtSpecCurrencyValDto.minorCurrencySymbols,
		nStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	newNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		nStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation

	err = newNStrFmtSpecCurrencyValDto.numericSeparators.CopyIn(
		&nStrFmtSpecCurrencyValDto.numericSeparators,
		ePrefix.XCtx("nStrFmtSpecCurrencyValDto->newNStrFmtSpecCurrencyValDto"))

	if err != nil {
		return newNStrFmtSpecCurrencyValDto, err
	}

	err =
		newNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
			&nStrFmtSpecCurrencyValDto.numFieldLenDto,
			ePrefix.XCtx("nStrFmtSpecCurrencyValDto->newNStrFmtSpecCurrencyValDto"))

	newNStrFmtSpecCurrencyValDto.lock = new(sync.Mutex)

	return newNStrFmtSpecCurrencyValDto, err
}

// ptr - Returns a pointer to a new instance of
// numStrFmtSpecCurrencyValueDtoNanobot.
//
func (nStrFmtSpecCurrValNanobot numStrFmtSpecCurrencyValueDtoNanobot) ptr() *numStrFmtSpecCurrencyValueDtoNanobot {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	newCurrencyValDtoNanobot :=
		new(numStrFmtSpecCurrencyValueDtoNanobot)

	newCurrencyValDtoNanobot.lock = new(sync.Mutex)

	return newCurrencyValDtoNanobot
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
//  nStrFmtSpecCurrencyValDto     *FormatterCurrency,
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
func (nStrFmtSpecCurrValNanobot *numStrFmtSpecCurrencyValueDtoNanobot) setCurrencyData(
	nStrFmtSpecCurrencyValDto *FormatterCurrency,
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

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecCurrencyValueDtoNanobot." +
			"setCurrencyData()")

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is invalid!\n"+
			"'nStrFmtSpecCurrencyValDto' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if currencySymbols == nil {
		currencySymbols = make([]rune, 0, 5)
	}

	lenCurrencySymbols := len(currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"currencySymbols is a zero length rune array!\n",
			ePrefix.String())

		return err
	}

	if minorCurrencySymbols == nil {
		minorCurrencySymbols = make([]rune, 0, 5)
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

	lenMinorCurrSymbols := len(minorCurrencySymbols)

	newNStrFmtSpecCurrencyValDto := FormatterCurrency{}

	newNStrFmtSpecCurrencyValDto.lock =
		new(sync.Mutex)

	newNStrFmtSpecCurrencyValDto.positiveValueFmt =
		positiveValueFmt

	newNStrFmtSpecCurrencyValDto.negativeValueFmt =
		negativeValueFmt

	newNStrFmtSpecCurrencyValDto.decimalDigits =
		decimalDigits

	newNStrFmtSpecCurrencyValDto.currencyCode =
		currencyCode

	newNStrFmtSpecCurrencyValDto.currencyCodeNo =
		currencyCodeNo

	newNStrFmtSpecCurrencyValDto.currencyName =
		currencyName

	newNStrFmtSpecCurrencyValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	runesCopied := copy(newNStrFmtSpecCurrencyValDto.currencySymbols,
		currencySymbols)

	if runesCopied != lenCurrencySymbols {
		err = fmt.Errorf("%v\n"+
			"Error copying currency symbols!\n"+
			"Expected to copy %v currency symbol runes.\n"+
			"However, only %v currency symbol runes were copied.\n"+
			"Source currencySymbols = '%v'\n"+
			"Destination newNStrFmtSpecCurrencyValDto.currencySymbols = '%v'\n",
			ePrefix.String(),
			lenCurrencySymbols,
			runesCopied,
			string(currencySymbols),
			string(newNStrFmtSpecCurrencyValDto.currencySymbols))

		return err
	}

	newNStrFmtSpecCurrencyValDto.minorCurrencyName =
		minorCurrencyName

	newNStrFmtSpecCurrencyValDto.minorCurrencySymbols =
		make([]rune, lenMinorCurrSymbols, 10)

	runesCopied = copy(newNStrFmtSpecCurrencyValDto.minorCurrencySymbols,
		minorCurrencySymbols)

	if runesCopied != lenMinorCurrSymbols {
		err = fmt.Errorf("%v\n"+
			"Error copying minor currency symbols!\n"+
			"Expected to copy %v minor currency symbol runes.\n"+
			"However, only %v minor currency symbol runes were copied.\n"+
			"Source minorCurrencySymbols = '%v'\n"+
			"Destination newNStrFmtSpecCurrencyValDto.minorCurrencySymbols = '%v'\n",
			ePrefix.String(),
			lenMinorCurrSymbols,
			runesCopied,
			string(minorCurrencySymbols),
			string(newNStrFmtSpecCurrencyValDto.minorCurrencySymbols))

		return err
	}

	newNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	return err
}
