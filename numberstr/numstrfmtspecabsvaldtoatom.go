package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoAtom struct {
	lock *sync.Mutex
}

// testValidityOfAbsoluteValDto - Receives an instance of
// numStrFmtSpecAbsoluteValueDtoAtom and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (nStrFmtSpecAbsValDtoAtom *numStrFmtSpecAbsoluteValueDtoAtom) testValidityOfAbsoluteValDto(
	nStrFmtSpecAbsoluteValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtSpecAbsValDtoAtom.lock == nil {
		nStrFmtSpecAbsValDtoAtom.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoAtom.lock.Lock()

	defer nStrFmtSpecAbsValDtoAtom.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecAbsoluteValueDtoAtom.testValidityOfAbsoluteValDto() "

	isValid = false

	if nStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsoluteValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return isValid, err
	}

	nStrAbsValDtoElectron :=
		numStrFmtSpecAbsoluteValueDtoElectron{}

	isValid,
		err = nStrAbsValDtoElectron.testAbsoluteValueFormat(
		nStrFmtSpecAbsoluteValDto,
		ePrefix+
			"\nValidating nStrFmtSpecAbsoluteValDto Format\n ")

	if err != nil {
		return isValid, err
	}

	err =
		nStrFmtSpecAbsoluteValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix +
				"\nValidating nStrFmtSpecAbsoluteValDto Number Separators\n ")

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid, err
}
