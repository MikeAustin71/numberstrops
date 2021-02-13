package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoMolecule struct {
	lock *sync.Mutex
}

// testValidityOfAbsoluteValDto - Receives an instance of
// NumStrFmtSpecAbsoluteValueDto and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (nStrFmtSpecAbsValDtoMolecule *numStrFmtSpecAbsoluteValueDtoMolecule) testValidityOfAbsoluteValDto(
	nStrFmtSpecAbsoluteValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecAbsValDtoMolecule.lock == nil {
		nStrFmtSpecAbsValDtoMolecule.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoMolecule.lock.Lock()

	defer nStrFmtSpecAbsValDtoMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecAbsoluteValueDtoMolecule.testValidityOfAbsoluteValDto()")

	isValid = false

	if nStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsoluteValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	nStrAbsValDtoAtom :=
		numStrFmtSpecAbsoluteValueDtoAtom{}

	isValid,
		err = nStrAbsValDtoAtom.testAbsoluteValueFormat(
		nStrFmtSpecAbsoluteValDto,
		ePrefix.XCtx("Validating nStrFmtSpecAbsoluteValDto Format"))

	if err != nil {
		return isValid, err
	}

	err =
		nStrFmtSpecAbsoluteValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix.XCtx(
				"Validating nStrFmtSpecAbsoluteValDto Number Separators"))

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid, err
}
