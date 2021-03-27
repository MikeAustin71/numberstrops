package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type formatterSignedNumberElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterSignedNumberElectron.
//
func (fmtSignedNumElectron formatterSignedNumberElectron) ptr() *formatterSignedNumberElectron {

	if fmtSignedNumElectron.lock == nil {
		fmtSignedNumElectron.lock = new(sync.Mutex)
	}

	fmtSignedNumElectron.lock.Lock()

	defer fmtSignedNumElectron.lock.Unlock()

	newFmtSignedNumElectron :=
		new(formatterSignedNumberElectron)

	newFmtSignedNumElectron.lock = new(sync.Mutex)

	return newFmtSignedNumElectron
}

// testSignedNumValNegativeValueFormatStr - Inspects the Negative Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
func (fmtSignedNumElectron *formatterSignedNumberElectron) testSignedNumValNegativeValueFormatStr(
	negativeValueFmtStr string,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtSignedNumElectron.lock == nil {
		fmtSignedNumElectron.lock = new(sync.Mutex)
	}

	fmtSignedNumElectron.lock.Lock()

	defer fmtSignedNumElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberElectron." +
			"testSignedNumValNegativeValueFormat()")

	isValid = false

	if len(negativeValueFmtStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValueFmtStr' is an empty string!\n"+
			"len(negativeValueFmtStr) == 0\n",
			ePrefix.String())

		return isValid, err
	}

	nStrFmtSpecSignedNumValQuark := formatterSignedNumberQuark{}

	validFmtChars :=
		nStrFmtSpecSignedNumValQuark.getValidSignedNumNegativeValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(negativeValueFmtStr)
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
				"Error: The Signed Value Negative Number Format String contains an invalid character!\n"+
				"Signed Number Negative Value Formats are NOT allowed to include this character.\n"+
				"Complete Signed Value Negative Number Format String= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix.String(),
				negativeValueFmtStr,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		negativeValueFmtStr, "127.54") &&
		!strings.Contains(
			negativeValueFmtStr, "NUMFIELD") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Signed Value Negative Number Format String is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either '127.54' or 'NUMFIELD'\n"+
			"to designate a place holder for the numeric value. This Number String Format has\n"+
			"neither placeholder!\n"+
			"Complete Signed Value Negative Number Format String= '%v'\n",
			ePrefix.String(),
			negativeValueFmtStr)

		return isValid, err
	}

	if !strings.Contains(
		negativeValueFmtStr, ")") &&
		!strings.Contains(
			negativeValueFmtStr, "(") &&
		!strings.Contains(
			negativeValueFmtStr, "-") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Signed Number Negative Value Format String is missing a negative sign\n"+
			"place holder. The Signed Number Negative Value Format String MUST contain either\n"+
			"a minus sign '-' or parenthesis '()' to designate a negative value.\n This Signed\n"+
			"Number Negative Value Format String does NOT contain a negative value placeholder!\n"+
			"Complete Signed Number Negative Value Format String= '%v'\n",
			ePrefix.String(),
			negativeValueFmtStr)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// testSignedNumValPositiveValueFormatStr - Inspects the Positive
// Value Number Format string for a Signed Number Value and returns
// an error if the format string is invalid.
//
//
func (fmtSignedNumElectron *formatterSignedNumberElectron) testSignedNumValPositiveValueFormatStr(
	positiveValueFmtStr string,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtSignedNumElectron.lock == nil {
		fmtSignedNumElectron.lock = new(sync.Mutex)
	}

	fmtSignedNumElectron.lock.Lock()

	defer fmtSignedNumElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberElectron." +
			"testSignedNumValPositiveValueFormatStr()")

	isValid = false

	if len(positiveValueFmtStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValueFmtStr' is an empty string!\n"+
			"The Signed Number Value Positive Value Format string is missing.\n"+
			"len(positiveValueFmtStr) == 0\n",
			ePrefix.String())
		return isValid, err
	}

	nStrFmtSpecSignedNumValQuark := formatterSignedNumberQuark{}

	validFmtChars :=
		nStrFmtSpecSignedNumValQuark.getValidSignedNumPositiveValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(positiveValueFmtStr)
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
				"Error: The Signed Number Format String ('positiveValueFmtStr') contains an invalid character!\n"+
				"Signed Number Positive Value Format Strings are NOT allowed to include this character.\n"+
				"Complete Signed Number Positive Value Format String= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				positiveValueFmtStr,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		positiveValueFmtStr, "127.54") &&
		!strings.Contains(
			positiveValueFmtStr, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Signed Number Positive Value Format String is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either '127.54' or\n"+
			"'NUMFIELD' to designate a place holder for the numeric value. This Signed Number\n"+
			"Positive Value Format String has neither placeholder!\n"+
			"Complete Signed Number Positive Value Format String= '%v'\n",
			ePrefix,
			positiveValueFmtStr)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
