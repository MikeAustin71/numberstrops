package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoAtom struct {
	lock *sync.Mutex
}

// testAbsoluteValueFormat - Inspects the format string for an
// Absolute Value number string and returns an error if the format
// string is invalid.
//
func (nStrAbsValDtoAtom *numStrFmtSpecAbsoluteValueDtoAtom) testAbsoluteValueFormat(
	nStrFmtSpecAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrAbsValDtoAtom.lock == nil {
		nStrAbsValDtoAtom.lock = new(sync.Mutex)
	}

	nStrAbsValDtoAtom.lock.Lock()

	defer nStrAbsValDtoAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecAbsoluteValueDtoElectron.testAbsoluteValueFormat()")

	isValid = false

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	if len(nStrFmtSpecAbsValDto.absoluteValFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'absoluteValFmt' is an empty string!\n"+
			"The Absolute Value Dto Format string is missing.\n"+
			"len(nStrFmtSpecAbsValDto.absoluteValFmt) == 0\n",
			ePrefix.String())

		return isValid, err
	}

	nStrAbsValDtoElectron :=
		numStrFmtSpecAbsoluteValueDtoElectron{}

	isValid,
		err = nStrAbsValDtoElectron.testAbsoluteValueFormatStr(
		nStrFmtSpecAbsValDto.absoluteValFmt,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecAbsValDto.absoluteValFmt'"))

	return isValid, err
}
