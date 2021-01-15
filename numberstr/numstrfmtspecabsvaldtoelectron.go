package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoElectron struct {
	lock *sync.Mutex
}

// testAbsoluteValueFormat - Inspects the format string for an
// Absolute Value number string and returns an error if the format
// string is invalid.
//
func (nStrAbsValDtoElectron *numStrFmtSpecAbsoluteValueDtoElectron) testAbsoluteValueFormat(
	nStrFmtSpecAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrAbsValDtoElectron.lock == nil {
		nStrAbsValDtoElectron.lock = new(sync.Mutex)
	}

	nStrAbsValDtoElectron.lock.Lock()

	defer nStrAbsValDtoElectron.lock.Unlock()

	ePrefix += "numStrFmtSpecAbsoluteValueDtoElectron.testAbsoluteValueFormat() "

	isValid = false

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is a 'nil' pointer!\n",
			ePrefix)
		return isValid, err
	}

	if len(nStrFmtSpecAbsValDto.absoluteValFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absoluteValFmt' is an empty string!\n"+
			"The Absolute Value Dto Format string is missing.\n"+
			"len(nStrFmtSpecAbsValDto.absoluteValFmt) == 0\n",
			ePrefix)
		return isValid, err
	}

	nStrFmtSpecAbsValueDtoQuark := numStrFmtSpecAbsoluteValueDtoQuark{}

	validFmtChars :=
		nStrFmtSpecAbsValueDtoQuark.getValidAbsoluteValFmtChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(nStrFmtSpecAbsValDto.absoluteValFmt)
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
				"Absolute Value Format Strings are NOT allowed to include this character.\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				ePrefix,
				nStrFmtSpecAbsValDto.absoluteValFmt,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if !strings.Contains(
		nStrFmtSpecAbsValDto.absoluteValFmt, "127.54") &&
		!strings.Contains(
			nStrFmtSpecAbsValDto.absoluteValFmt, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format string MUST contain either\n"+
			"'127.54' or 'NUMFIELD' to designate a place holder for the\n"+
			"numeric value. This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			nStrFmtSpecAbsValDto.absoluteValFmt)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
