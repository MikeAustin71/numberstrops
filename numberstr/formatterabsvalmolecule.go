package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueMolecule struct {
	lock *sync.Mutex
}

// testValidityOfAbsoluteValDto - Receives an instance of
// FormatterAbsoluteValue and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (fmtAbsValMolecule *formatterAbsoluteValueMolecule) testValidityOfAbsoluteValDto(
	fmtAbsoluteVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtAbsValMolecule.lock == nil {
		fmtAbsValMolecule.lock = new(sync.Mutex)
	}

	fmtAbsValMolecule.lock.Lock()

	defer fmtAbsValMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueMolecule." +
			"testValidityOfAbsoluteValDto()")

	isValid = false

	if fmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	nStrAbsValDtoAtom :=
		numStrFmtSpecAbsoluteValueDtoAtom{}

	isValid,
		err = nStrAbsValDtoAtom.testAbsoluteValueFormat(
		fmtAbsoluteVal,
		ePrefix.XCtx("Validating fmtAbsoluteVal Format"))

	if err != nil {
		return isValid, err
	}

	err =
		fmtAbsoluteVal.numericSeparators.IsValidInstanceError(
			ePrefix.XCtx(
				"Validating fmtAbsoluteVal Number Separators"))

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid, err
}
