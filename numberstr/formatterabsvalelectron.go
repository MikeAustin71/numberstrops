package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type formatterAbsoluteValueElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueElectron.
//
func (fmtAbsValElectron formatterAbsoluteValueElectron) ptr() *formatterAbsoluteValueElectron {

	if fmtAbsValElectron.lock == nil {
		fmtAbsValElectron.lock = new(sync.Mutex)
	}

	fmtAbsValElectron.lock.Lock()

	defer fmtAbsValElectron.lock.Unlock()

	newAbsValElectron :=
		new(formatterAbsoluteValueElectron)

	newAbsValElectron.lock = new(sync.Mutex)

	return newAbsValElectron
}

// testAbsoluteValueFormatStr - Inspects the format string for an
// Absolute Value number string and returns an error if the format
// string is invalid.
//
func (fmtAbsValElectron *formatterAbsoluteValueElectron) testAbsoluteValueFormatStr(
	absoluteValueFormatStr string,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtAbsValElectron.lock == nil {
		fmtAbsValElectron.lock = new(sync.Mutex)
	}

	fmtAbsValElectron.lock.Lock()

	defer fmtAbsValElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueElectron." +
			"testAbsoluteValueFormatStr()")

	isValid = false

	if len(absoluteValueFormatStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absoluteValueFormatStr' is an empty string!\n"+
			"The Absolute Value Dto Format String is missing.\n"+
			"len(absoluteValueFormatStr) == 0\n",
			ePrefix.String())

		return isValid, err
	}

	nStrFmtSpecAbsValueDtoQuark := numStrFmtSpecAbsoluteValueDtoQuark{}

	validFmtChars :=
		nStrFmtSpecAbsValueDtoQuark.getValidAbsoluteValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(absoluteValueFormatStr)
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
				"Error: The Absolute Value Number String Format contains an invalid character!\n"+
				"Absolute Value Format Strings are NOT allowed to include this character.\n"+
				"Complete Absolute Value Format String= '%v'\n"+
				"invalid character == '%v' at Index [%v] \n",
				ePrefix.String(),
				absoluteValueFormatStr,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		absoluteValueFormatStr, "127.54") &&
		!strings.Contains(
			absoluteValueFormatStr, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Absolute Value Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either '127.54' or\n"+
			"'NUMFIELD' to designate a place holder for the numeric value. This Number\n"+
			"String Format has neither placeholder!\n"+
			"Complete Absolute Value Number String Format= '%v'\n",
			ePrefix.String(),
			absoluteValueFormatStr)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
