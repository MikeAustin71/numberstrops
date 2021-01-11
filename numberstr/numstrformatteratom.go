package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFormatterAtom struct {
	lock *sync.Mutex
}

// testCurrencyValPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrFormtrAtom *numStrFormatterAtom) testCurrencyValPositiveValueFormat(
	currencyValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrAtom.lock == nil {
		nStrFormtrAtom.lock = new(sync.Mutex)
	}

	nStrFormtrAtom.lock.Lock()

	defer nStrFormtrAtom.lock.Unlock()

	ePrefix += "numStrFormatterAtom.testCurrencyValPositiveValueFormat() "

	isValid = false

	if currencyValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(currencyValueFormat.positiveValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValueFmt' is an empty string!\n"+
			"The Currency Value Positive Value Format string is missing.\n"+
			"len(currencyValueFormat.positiveValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrQuark := numStrFormatterQuark{}

	validFmtChars := nStrFormtrQuark.getValidCurrencyPositiveValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(currencyValueFormat.positiveValueFmt)
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
				"Currency Positive Value Formats are NOT allowed to include this character.\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				currencyValueFormat.positiveValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		currencyValueFormat.positiveValueFmt, "127.54") &&
		!strings.Contains(
			currencyValueFormat.positiveValueFmt, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			currencyValueFormat.positiveValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testCurrencyValNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrFormtrAtom *numStrFormatterAtom) testCurrencyValNegativeValueFormat(
	currencyValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrAtom.lock == nil {
		nStrFormtrAtom.lock = new(sync.Mutex)
	}

	nStrFormtrAtom.lock.Lock()

	defer nStrFormtrAtom.lock.Unlock()

	ePrefix += "numStrFormatterAtom.testCurrencyValNegativeValueFormat() "

	isValid = false

	if currencyValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(currencyValueFormat.negativeValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValueFmt' is an empty string!\n"+
			"len(currencyValueFormat.negativeValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrQuark := numStrFormatterQuark{}

	validFmtChars := nStrFormtrQuark.getValidCurrencyNegativeValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(currencyValueFormat.negativeValueFmt)
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
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				currencyValueFormat.negativeValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		currencyValueFormat.negativeValueFmt, "127.54") &&
		!strings.Contains(
			currencyValueFormat.negativeValueFmt, "NUMFIELD") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			currencyValueFormat.negativeValueFmt)

		return isValid, err
	}

	if !strings.Contains(
		currencyValueFormat.negativeValueFmt, ")") &&
		!strings.Contains(
			currencyValueFormat.negativeValueFmt, "(") &&
		!strings.Contains(
			currencyValueFormat.negativeValueFmt, "-") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a negative sign\n"+
			"place holder. The format string MUST contain either a minus\n"+
			"sign '-' or parenthesis '()' to designate a negative value.\n"+
			"This Number String Format does NOT contain a negative value placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			currencyValueFormat.negativeValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testAbsoluteValPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrFormtrAtom *numStrFormatterAtom) testAbsoluteValPositiveValueFormat(
	absoluteValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrAtom.lock == nil {
		nStrFormtrAtom.lock = new(sync.Mutex)
	}

	nStrFormtrAtom.lock.Lock()

	defer nStrFormtrAtom.lock.Unlock()

	ePrefix += "numStrFormatterAtom.testAbsoluteValPositiveValueFormat() "

	isValid = false

	if absoluteValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absoluteValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(absoluteValueFormat.positiveValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValueFmt' is an empty string!\n"+
			"The Absolute Value Positive Value Format string is missing.\n"+
			"len(absoluteValueFormat.positiveValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrQuark := numStrFormatterQuark{}

	validFmtChars := nStrFormtrQuark.getValidAbsolutePositiveValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(absoluteValueFormat.positiveValueFmt)
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
				"Absolute Positive Value Formats are NOT allowed to include this character.\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				absoluteValueFormat.positiveValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		absoluteValueFormat.positiveValueFmt, "127.54") &&
		!strings.Contains(
			absoluteValueFormat.positiveValueFmt, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Absolute Value Number String Format is missing a\n"+
			"place holder for the numeric value. The format string MUST contain\n"+
			"either '127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			absoluteValueFormat.positiveValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testAbsoluteValNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrFormtrAtom *numStrFormatterAtom) testAbsoluteValNegativeValueFormat(
	absoluteValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrAtom.lock == nil {
		nStrFormtrAtom.lock = new(sync.Mutex)
	}

	nStrFormtrAtom.lock.Lock()

	defer nStrFormtrAtom.lock.Unlock()

	ePrefix += "numStrFormatterAtom.testAbsoluteValNegativeValueFormat() "

	isValid = false

	if absoluteValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absoluteValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(absoluteValueFormat.negativeValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValueFmt' is an empty string!\n"+
			"The Absolute Value Negative Value Format string is missing.\n"+
			"len(absoluteValueFormat.negativeValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrQuark := numStrFormatterQuark{}

	validFmtChars := nStrFormtrQuark.getValidAbsoluteNegativeValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(absoluteValueFormat.negativeValueFmt)
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
				"Absolute Negative Value Formats are NOT allowed to include this character.\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				absoluteValueFormat.negativeValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		absoluteValueFormat.negativeValueFmt, "127.54") &&
		!strings.Contains(
			absoluteValueFormat.negativeValueFmt, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Absolute Value Number String Format is missing a\n"+
			"place holder for the numeric value. The format string MUST contain\n"+
			"either '127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			absoluteValueFormat.negativeValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testSignedNumValPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrFormtrAtom *numStrFormatterAtom) testSignedNumValPositiveValueFormat(
	signedNumberValueFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrAtom.lock == nil {
		nStrFormtrAtom.lock = new(sync.Mutex)
	}

	nStrFormtrAtom.lock.Lock()

	defer nStrFormtrAtom.lock.Unlock()

	ePrefix += "numStrFormatterAtom.testSignedNumValPositiveValueFormat() "

	isValid = false

	if signedNumberValueFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumberValueFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(signedNumberValueFormat.positiveValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValueFmt' is an empty string!\n"+
			"The Signed Number Value Positive Value Format string is missing.\n"+
			"len(signedNumberValueFormat.positiveValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrQuark := numStrFormatterQuark{}

	validFmtChars := nStrFormtrQuark.getValidSignedNumPositiveValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(signedNumberValueFormat.positiveValueFmt)
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
				ePrefix,
				signedNumberValueFormat.positiveValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		signedNumberValueFormat.positiveValueFmt, "127.54") &&
		!strings.Contains(
			signedNumberValueFormat.positiveValueFmt, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			signedNumberValueFormat.positiveValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testSignedNumValNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrFormtrAtom *numStrFormatterAtom) testSignedNumValNegativeValueFormat(
	signedNumberValFormat *NumStrFormatter,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFormtrAtom.lock == nil {
		nStrFormtrAtom.lock = new(sync.Mutex)
	}

	nStrFormtrAtom.lock.Lock()

	defer nStrFormtrAtom.lock.Unlock()

	ePrefix += "numStrFormatterAtom.testSignedNumValNegativeValueFormat() "

	isValid = false

	if signedNumberValFormat == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumberValFormat' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(signedNumberValFormat.negativeValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValueFmt' is an empty string!\n"+
			"len(signedNumberValFormat.negativeValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFormtrQuark := numStrFormatterQuark{}

	validFmtChars := nStrFormtrQuark.getValidSignedNumNegativeValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(signedNumberValFormat.negativeValueFmt)
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
				"Signed Number Negative Value Formats are NOT allowed to include this character.\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				signedNumberValFormat.negativeValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		signedNumberValFormat.negativeValueFmt, "127.54") &&
		!strings.Contains(
			signedNumberValFormat.negativeValueFmt, "NUMFIELD") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			signedNumberValFormat.negativeValueFmt)

		return isValid, err
	}

	if !strings.Contains(
		signedNumberValFormat.negativeValueFmt, ")") &&
		!strings.Contains(
			signedNumberValFormat.negativeValueFmt, "(") &&
		!strings.Contains(
			signedNumberValFormat.negativeValueFmt, "-") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a negative sign\n"+
			"place holder. The format string MUST contain either a minus\n"+
			"sign '-' or parenthesis '()' to designate a negative value.\n"+
			"This Number String Format does NOT contain a negative value placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			signedNumberValFormat.negativeValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
