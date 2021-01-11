package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatterNanobot struct {
	lock *sync.Mutex
}

// testValidityOfCurrencyFormatter - Analyzes a currency format to
// determine if any of the component data elements are invalid.
//
// If the currency formatter instance is judged to be invalid, a
// boolean value of 'false' is returned along with an error object
// containing a detailed error message.
//
func (nStrFormtrNanobot *numStrFormatterNanobot) testValidityOfCurrencyFormatter(
	currencyValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrNanobot.lock == nil {
		nStrFormtrNanobot.lock = new(sync.Mutex)
	}

	nStrFormtrNanobot.lock.Lock()

	defer nStrFormtrNanobot.lock.Unlock()

	ePrefix += "nStrFormtrNanobot.testValidityOfCurrencyFormatter() "

	isValid = false

	if currencyValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if currencyValueFormat.valueDisplaySpec !=
		NumStrValSpec(0).CurrencyValue() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter, 'currencyValueFormat' has an invalid\n"+
			"Number String Value Specification (NumStrValSpec).\n"+
			"Expected 'currencyValueFormat'== NumStrValSpec(0).CurrencyValue()\n"+
			"Instead, 'currencyValueFormat'== '%v'\n",
			ePrefix,
			currencyValueFormat.valueDisplaySpec.XValueInt())

		return isValid, err
	}

	if currencyValueFormat.currencySymbol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Currency Value Format is missing currency symbol!\n"+
			"currencyValueFormat.currencySymbol == 0\n",
			ePrefix)

		return isValid, err
	}

	if len(currencyValueFormat.countryName) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Currency Value Format 'country name' is missing!\n"+
			"len(currencyValueFormat.countryName) == 0\n",
			ePrefix)

		return isValid, err
	}

	if currencyValueFormat.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Currency Value Format is missing the decimal separator character!\n"+
			"This separates integer and fractional components in a currency number string.\n"+
			"currencyValueFormat.decimalSeparator == 0\n",
			ePrefix)

		return isValid, err
	}

	if currencyValueFormat.integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Currency Value Format is missing the integer digits separator\n"+
			"character (a.k.a.) thousands separator!\n"+
			"currencyValueFormat.integerDigitsSeparator == 0\n",
			ePrefix)

		return isValid, err
	}

	if len(currencyValueFormat.integerDigitsGroupingSequence) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Currency Value Format is missing the integer digits grouping\n"+
			"sequence. This integer array is used to specify the separation\n"+
			"of integer digits into thousands or other groupings depending on\n"+
			"the culture or nationality.\n"+
			"len(currencyValueFormat.integerDigitsGroupingSequence) == 0\n",
			ePrefix)

		return isValid, err
	}

	nStrFormtrAtom := numStrFormatterAtom{}

	isValid,
		err = nStrFormtrAtom.testCurrencyValPositiveValueFormat(
		currencyValueFormat,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrFormtrAtom.testCurrencyValNegativeValueFormat(
		currencyValueFormat,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testValidityOfAbsoluteValueFormatter - Analyzes a currency format to
// determine if any of the component data elements are invalid.
//
// If the currency formatter instance is judged to be invalid, a
// boolean value of 'false' is returned along with an error object
// containing a detailed error message.
//
func (nStrFormtrNanobot *numStrFormatterNanobot) testValidityOfAbsoluteValueFormatter(
	absoluteValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrNanobot.lock == nil {
		nStrFormtrNanobot.lock = new(sync.Mutex)
	}

	nStrFormtrNanobot.lock.Lock()

	defer nStrFormtrNanobot.lock.Unlock()

	ePrefix += "numStrFormatterNanobot.testValidityOfAbsoluteValueFormatter() "

	err = nil
	isValid = false

	if absoluteValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absoluteValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if absoluteValueFormat.valueDisplaySpec !=
		NumStrValSpec(0).AbsoluteValue() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter, 'absoluteValueFormat' has an invalid\n"+
			"Number String Value Specification (NumStrValSpec).\n"+
			"Expected 'absoluteValueFormat'== NumStrValSpec(0).AbsoluteValue()\n"+
			"Instead, 'absoluteValueFormat'== '%v'\n",
			ePrefix,
			absoluteValueFormat.valueDisplaySpec.XValueInt())

		return isValid, err
	}

	if absoluteValueFormat.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Absolute Value Format is missing the decimal separator character!\n"+
			"This separates integer and fractional components in an absolute value number string.\n"+
			"absoluteValueFormat.decimalSeparator == 0\n",
			ePrefix)

		return isValid, err
	}

	if absoluteValueFormat.integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Absolute Value Format is missing the integer digits separator!\n"+
			"character (a.k.a.) thousands separator!\n"+
			"absoluteValueFormat.integerDigitsSeparator == 0\n",
			ePrefix)

		return isValid, err
	}

	if len(absoluteValueFormat.integerDigitsGroupingSequence) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Absolute Value Format is missing the integer digits grouping\n"+
			"sequence. This integer array is used to specify the separation\n"+
			"of integer digits into thousands or other groupings depending on\n"+
			"the culture or nationality.\n"+
			"len(absoluteValueFormat.integerDigitsGroupingSequence) == 0\n",
			ePrefix)

		return isValid, err
	}

	nStrFormtrAtom := numStrFormatterAtom{}

	isValid,
		err = nStrFormtrAtom.testAbsoluteValPositiveValueFormat(
		absoluteValueFormat,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrFormtrAtom.testAbsoluteValNegativeValueFormat(
		absoluteValueFormat,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}

// testValidityOfSignedNumberValueFormatter - Analyzes a currency format to
// determine if any of the component data elements are invalid.
//
// If the currency formatter instance is judged to be invalid, a
// boolean value of 'false' is returned along with an error object
// containing a detailed error message.
//
func (nStrFormtrNanobot *numStrFormatterNanobot) testValidityOfSignedNumberValueFormatter(
	signedNumberFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrNanobot.lock == nil {
		nStrFormtrNanobot.lock = new(sync.Mutex)
	}

	nStrFormtrNanobot.lock.Lock()

	defer nStrFormtrNanobot.lock.Unlock()

	ePrefix += "numStrFormatterNanobot.testValidityOfSignedNumberValueFormatter() "

	err = nil
	isValid = false

	if signedNumberFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumberFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if signedNumberFormat.valueDisplaySpec !=
		NumStrValSpec(0).SignedNumberValue() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter, 'signedNumberFormat' has an invalid\n"+
			"Number String Value Specification (NumStrValSpec).\n"+
			"Expected 'signedNumberFormat'== NumStrValSpec(0).SignedNumberValue()\n"+
			"Instead, 'signedNumberFormat'== '%v'\n",
			ePrefix,
			signedNumberFormat.valueDisplaySpec.XValueInt())

		return isValid, err
	}

	if signedNumberFormat.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Signed Number Value Format is missing the decimal separator character!\n"+
			"This separates integer and fractional components in a signed number string.\n"+
			"signedNumberFormat.decimalSeparator == 0\n",
			ePrefix)

		return isValid, err
	}

	if signedNumberFormat.integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Signed Number Value Format is missing the integer digits separator!\n"+
			"character (a.k.a.) thousands separator!\n"+
			"signedNumberFormat.integerDigitsSeparator == 0\n",
			ePrefix)

		return isValid, err
	}

	if len(signedNumberFormat.integerDigitsGroupingSequence) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Signed Number Value Format is missing the integer digits grouping\n"+
			"sequence. This integer array is used to specify the separation\n"+
			"of integer digits into thousands or other groupings depending on\n"+
			"the culture or nationality.\n"+
			"len(signedNumberFormat.integerDigitsGroupingSequence) == 0\n",
			ePrefix)

		return isValid, err
	}

	nStrFormtrAtom := numStrFormatterAtom{}

	isValid,
		err = nStrFormtrAtom.testSignedNumValPositiveValueFormat(
		signedNumberFormat,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrFormtrAtom.testSignedNumValNegativeValueFormat(
		signedNumberFormat,
		ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid = true
	return isValid, err
}
