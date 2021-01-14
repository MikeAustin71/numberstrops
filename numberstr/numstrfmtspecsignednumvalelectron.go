package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrSignedNumValElectron struct {
	lock *sync.Mutex
}

// testSignedNumValPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (nStrSignedNumElectron *numStrSignedNumValElectron) testSignedNumValPositiveValueFormat(
	nStrFmtSpecSignedNumberValueDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrSignedNumElectron.lock == nil {
		nStrSignedNumElectron.lock = new(sync.Mutex)
	}

	nStrSignedNumElectron.lock.Lock()

	defer nStrSignedNumElectron.lock.Unlock()

	ePrefix += "numStrSignedNumValElectron.testSignedNumValPositiveValueFormat() "

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(nStrFmtSpecSignedNumberValueDto.positiveValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveValueFmt' is an empty string!\n"+
			"The Signed Number Value Positive Value Format string is missing.\n"+
			"len(nStrFmtSpecSignedNumberValueDto.positiveValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFmtSpecSignedNumValQuark := numStrFmtSpecSignedNumValQuark{}

	validFmtChars :=
		nStrFmtSpecSignedNumValQuark.getValidSignedNumPositiveValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(nStrFmtSpecSignedNumberValueDto.positiveValueFmt)
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
				nStrFmtSpecSignedNumberValueDto.positiveValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		nStrFmtSpecSignedNumberValueDto.positiveValueFmt, "127.54") &&
		!strings.Contains(
			nStrFmtSpecSignedNumberValueDto.positiveValueFmt, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			nStrFmtSpecSignedNumberValueDto.positiveValueFmt)

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
func (nStrSignedNumElectron *numStrSignedNumValElectron) testSignedNumValNegativeValueFormat(
	nStrFmtSpecSignedNumberValueDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrSignedNumElectron.lock == nil {
		nStrSignedNumElectron.lock = new(sync.Mutex)
	}

	nStrSignedNumElectron.lock.Lock()

	defer nStrSignedNumElectron.lock.Unlock()

	ePrefix += "numStrSignedNumValElectron.testSignedNumValNegativeValueFormat() "

	isValid = false

	if nStrFmtSpecSignedNumberValueDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumberValueDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(nStrFmtSpecSignedNumberValueDto.negativeValueFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeValueFmt' is an empty string!\n"+
			"len(nStrFmtSpecSignedNumberValueDto.negativeValueFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFmtSpecSignedNumValQuark := numStrFmtSpecSignedNumValQuark{}

	validFmtChars :=
		nStrFmtSpecSignedNumValQuark.getValidSignedNumNegativeValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(nStrFmtSpecSignedNumberValueDto.negativeValueFmt)
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
				nStrFmtSpecSignedNumberValueDto.negativeValueFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		nStrFmtSpecSignedNumberValueDto.negativeValueFmt, "127.54") &&
		!strings.Contains(
			nStrFmtSpecSignedNumberValueDto.negativeValueFmt, "NUMFIELD") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			nStrFmtSpecSignedNumberValueDto.negativeValueFmt)

		return isValid, err
	}

	if !strings.Contains(
		nStrFmtSpecSignedNumberValueDto.negativeValueFmt, ")") &&
		!strings.Contains(
			nStrFmtSpecSignedNumberValueDto.negativeValueFmt, "(") &&
		!strings.Contains(
			nStrFmtSpecSignedNumberValueDto.negativeValueFmt, "-") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a negative sign\n"+
			"place holder. The format string MUST contain either a minus\n"+
			"sign '-' or parenthesis '()' to designate a negative value.\n"+
			"This Number String Format does NOT contain a negative value placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			nStrFmtSpecSignedNumberValueDto.negativeValueFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
