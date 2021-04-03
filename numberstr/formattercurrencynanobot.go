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
	} else {
		ePrefix = ePrefix.CopyPtr()
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// equal - Receives two FormatterCurrency objects and proceeds to
// determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtCurrencyOne      *FormatterCurrency
//     - A pointer to the first FormatterCurrency object. This
//       method will compare all data elements in this object to
//       all corresponding data elements in the second
//       FormatterCurrency object to determine if they are
//       equivalent.
//
//
//  fmtCurrencyTwo      *FormatterCurrency
//     - A pointer to the second FormatterCurrency object. This
//       method will compare all data elements in the first
//       FormatterCurrency object to all corresponding data
//       elements in this second FormatterCurrency object to
//       determine if they are equivalent.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  isEqual             bool
//     - If all the data elements in 'fmtCurrencyOne' are equal to
//       all the corresponding data elements in 'fmtCurrencyTwo',
//       this return parameter will be set to 'true'. If the data
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtCurrencyOne' are equal to
//       all the corresponding data elements in 'fmtCurrencyTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtCurrNanobot *formatterCurrencyNanobot) equal(
	fmtCurrencyOne *FormatterCurrency,
	fmtCurrencyTwo *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtCurrNanobot.lock == nil {
		fmtCurrNanobot.lock = new(sync.Mutex)
	}

	fmtCurrNanobot.lock.Lock()

	defer fmtCurrNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyNanobot." +
			"equal()")

	isEqual = false

	if fmtCurrencyOne == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCurrencyOne' is invalid!\n"+
			"'fmtCurrencyOne' is a 'nil' pointer\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtCurrencyTwo == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCurrencyTwo' is invalid!\n"+
			"'fmtCurrencyTwo' is a 'nil' pointer\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtCurrencyOne.numStrFmtType !=
		fmtCurrencyTwo.numStrFmtType {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.numStrFmtType!=fmtCurrencyTwo.numStrFmtType\n"+
			"fmtCurrencyOne.numStrFmtType='%v'\n"+
			"fmtCurrencyTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.numStrFmtType.String(),
			fmtCurrencyTwo.numStrFmtType.String())

		return isEqual, err
	}

	if fmtCurrencyOne.positiveValueFmt !=
		fmtCurrencyTwo.positiveValueFmt {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.positiveValueFmt!=fmtCurrencyTwo.positiveValueFmt\n"+
			"fmtCurrencyOne.positiveValueFmt='%v'\n"+
			"fmtCurrencyTwo.positiveValueFmt='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.positiveValueFmt,
			fmtCurrencyTwo.positiveValueFmt)

		return isEqual, err
	}

	if fmtCurrencyOne.negativeValueFmt !=
		fmtCurrencyTwo.negativeValueFmt {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.negativeValueFmt!=fmtCurrencyTwo.negativeValueFmt\n"+
			"fmtCurrencyOne.negativeValueFmt='%v'\n"+
			"fmtCurrencyTwo.negativeValueFmt='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.negativeValueFmt,
			fmtCurrencyTwo.negativeValueFmt)

		return isEqual, err
	}

	if fmtCurrencyOne.decimalDigits !=
		fmtCurrencyTwo.decimalDigits {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.decimalDigits!=fmtCurrencyTwo.decimalDigits\n"+
			"fmtCurrencyOne.decimalDigits='%v'\n"+
			"fmtCurrencyTwo.decimalDigits='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.decimalDigits,
			fmtCurrencyTwo.decimalDigits)

		return isEqual, err
	}

	if fmtCurrencyOne.currencyCode !=
		fmtCurrencyTwo.currencyCode {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.currencyCode!=fmtCurrencyTwo.currencyCode\n"+
			"fmtCurrencyOne.currencyCode='%v'\n"+
			"fmtCurrencyTwo.currencyCode='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.currencyCode,
			fmtCurrencyTwo.currencyCode)

		return isEqual, err
	}

	if fmtCurrencyOne.currencyCodeNo !=
		fmtCurrencyTwo.currencyCodeNo {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.currencyCodeNo!=fmtCurrencyTwo.currencyCodeNo\n"+
			"fmtCurrencyOne.currencyCodeNo='%v'\n"+
			"fmtCurrencyTwo.currencyCodeNo='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.currencyCodeNo,
			fmtCurrencyTwo.currencyCodeNo)

		return isEqual, err
	}

	if fmtCurrencyOne.currencyName !=
		fmtCurrencyTwo.currencyName {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.currencyName!=fmtCurrencyTwo.currencyName\n"+
			"fmtCurrencyOne.currencyName='%v'\n"+
			"fmtCurrencyTwo.currencyName='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.currencyName,
			fmtCurrencyTwo.currencyName)

		return isEqual, err
	}

	if fmtCurrencyOne.currencySymbols == nil &&
		fmtCurrencyTwo.currencySymbols != nil {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.currencySymbols!=fmtCurrencyTwo.currencySymbols\n"+
			"fmtCurrencyOne.currencySymbols=='nil'\n"+
			"fmtCurrencyTwo.currencySymbols !='nil'\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtCurrencyOne.currencySymbols != nil &&
		fmtCurrencyTwo.currencySymbols == nil {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.currencySymbols!=fmtCurrencyTwo.currencySymbols\n"+
			"fmtCurrencyOne.currencySymbols !='nil'\n"+
			"fmtCurrencyTwo.currencySymbols =='nil'\n",
			ePrefix.String())

		return isEqual, err
	}

	lenSymbols := len(fmtCurrencyOne.currencySymbols)

	if lenSymbols != len(fmtCurrencyTwo.currencySymbols) {

		err = fmt.Errorf("%v\n"+
			"Length fmtCurrencyOne.currencySymbols != "+
			"Length fmtCurrencyTwo.currencySymbols\n"+
			"Length fmtCurrencyOne.currencySymbols='%v'\n"+
			"Length fmtCurrencyTwo.currencySymbols='%v'\n",
			ePrefix.String(),
			string(fmtCurrencyOne.currencySymbols),
			string(fmtCurrencyTwo.currencySymbols))

		return isEqual, err
	}

	for i := 0; i < lenSymbols; i++ {

		if fmtCurrencyOne.currencySymbols[i] !=

			fmtCurrencyTwo.currencySymbols[i] {
			err = fmt.Errorf("%v\n"+
				"Length fmtCurrencyOne.currencySymbols != "+
				"Length fmtCurrencyTwo.currencySymbols\n"+
				"fmtCurrencyOne.currencySymbols='%v'\n"+
				"fmtCurrencyTwo.currencySymbols='%v'\n",
				ePrefix.String(),
				string(fmtCurrencyOne.currencySymbols),
				string(fmtCurrencyTwo.currencySymbols))

			return isEqual, err
		}

	}

	if fmtCurrencyOne.minorCurrencyName !=
		fmtCurrencyTwo.minorCurrencyName {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.minorCurrencyName!="+
			"fmtCurrencyTwo.minorCurrencyName\n"+
			"fmtCurrencyOne.minorCurrencyName='%v'\n"+
			"fmtCurrencyTwo.minorCurrencyName='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.minorCurrencyName,
			fmtCurrencyTwo.minorCurrencyName)

		return isEqual, err
	}

	if fmtCurrencyOne.minorCurrencySymbols == nil &&
		fmtCurrencyTwo.minorCurrencySymbols != nil {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.minorCurrencySymbols != fmtCurrencyTwo.minorCurrencySymbols\n"+
			"fmtCurrencyOne.minorCurrencySymbols=='nil'\n"+
			"fmtCurrencyTwo.minorCurrencySymbols !='nil'\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtCurrencyOne.minorCurrencySymbols != nil &&
		fmtCurrencyTwo.minorCurrencySymbols == nil {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.minorCurrencySymbols!=fmtCurrencyTwo.minorCurrencySymbols\n"+
			"fmtCurrencyOne.minorCurrencySymbols !='nil'\n"+
			"fmtCurrencyTwo.minorCurrencySymbols =='nil'\n",
			ePrefix.String())

		return isEqual, err
	}

	lenSymbols = len(fmtCurrencyOne.minorCurrencySymbols)

	if lenSymbols != len(fmtCurrencyTwo.minorCurrencySymbols) {

		err = fmt.Errorf("%v\n"+
			"Length fmtCurrencyOne.minorCurrencySymbols != "+
			"Length fmtCurrencyTwo.minorCurrencySymbols\n"+
			"Length fmtCurrencyOne.minorCurrencySymbols='%v'\n"+
			"Length fmtCurrencyTwo.minorCurrencySymbols='%v'\n",
			ePrefix.String(),
			string(fmtCurrencyOne.minorCurrencySymbols),
			string(fmtCurrencyTwo.minorCurrencySymbols))

		return isEqual, err
	}

	for i := 0; i < lenSymbols; i++ {

		if fmtCurrencyOne.minorCurrencySymbols[i] !=
			fmtCurrencyTwo.minorCurrencySymbols[i] {

			err = fmt.Errorf("%v\n"+
				"Length fmtCurrencyOne.minorCurrencySymbols != "+
				"Length fmtCurrencyTwo.minorCurrencySymbols\n"+
				"fmtCurrencyOne.minorCurrencySymbols='%v'\n"+
				"fmtCurrencyTwo.minorCurrencySymbols='%v'\n",
				ePrefix.String(),
				string(fmtCurrencyOne.minorCurrencySymbols),
				string(fmtCurrencyTwo.minorCurrencySymbols))

			return isEqual, err
		}

	}

	if fmtCurrencyOne.turnOnIntegerDigitsSeparation !=
		fmtCurrencyTwo.turnOnIntegerDigitsSeparation {

		err = fmt.Errorf("%v\n"+
			"fmtCurrencyOne.turnOnIntegerDigitsSeparation!="+
			"fmtCurrencyTwo.turnOnIntegerDigitsSeparation\n"+
			"fmtCurrencyOne.turnOnIntegerDigitsSeparation='%v'\n"+
			"fmtCurrencyTwo.turnOnIntegerDigitsSeparation='%v'\n",
			ePrefix.String(),
			fmtCurrencyOne.turnOnIntegerDigitsSeparation,
			fmtCurrencyTwo.turnOnIntegerDigitsSeparation)

		return isEqual, err
	}

	_,
		err = fmtCurrencyOne.numericSeparators.Equal(
		fmtCurrencyTwo.numericSeparators,
		ePrefix.XCtx(
			"fmtCurrencyOne vs fmtCurrencyTwo"))

	if err != nil {
		return isEqual, err
	}

	_,
		err = fmtCurrencyOne.numFieldDto.Equal(
		fmtCurrencyTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtCurrencyOne vs fmtCurrencyTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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
