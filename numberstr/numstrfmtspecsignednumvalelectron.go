package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrSignedNumValElectron struct {
	lock *sync.Mutex
}

// testSignedNumValNegativeValueFormatStr - Inspects the Negative Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
func (nStrSignedNumElectron *numStrSignedNumValElectron) testSignedNumValNegativeValueFormatStr(
	negativeValueFmtStr string,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrSignedNumElectron.lock == nil {
		nStrSignedNumElectron.lock = new(sync.Mutex)
	}

	nStrSignedNumElectron.lock.Lock()

	defer nStrSignedNumElectron.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrSignedNumValElectron.testSignedNumValNegativeValueFormat() "

	isValid = false

	if len(negativeValueFmtStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValueFmtStr' is an empty string!\n"+
			"len(negativeValueFmtStr) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFmtSpecSignedNumValQuark := numStrFmtSpecSignedNumValQuark{}

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
				ePrefix,
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
			ePrefix,
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
			ePrefix,
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
func (nStrSignedNumElectron *numStrSignedNumValElectron) testSignedNumValPositiveValueFormatStr(
	positiveValueFmtStr string,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrSignedNumElectron.lock == nil {
		nStrSignedNumElectron.lock = new(sync.Mutex)
	}

	nStrSignedNumElectron.lock.Lock()

	defer nStrSignedNumElectron.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrSignedNumValElectron.testSignedNumValPositiveValueFormatStr() "

	isValid = false

	if len(positiveValueFmtStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValueFmtStr' is an empty string!\n"+
			"The Signed Number Value Positive Value Format string is missing.\n"+
			"len(positiveValueFmtStr) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFmtSpecSignedNumValQuark := numStrFmtSpecSignedNumValQuark{}

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
