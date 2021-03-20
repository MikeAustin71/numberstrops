package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueMolecule struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueQuark.
//
func (fmtAbsValMolecule formatterAbsoluteValueMolecule) ptr() *formatterAbsoluteValueMolecule {

	if fmtAbsValMolecule.lock == nil {
		fmtAbsValMolecule.lock = new(sync.Mutex)
	}

	fmtAbsValMolecule.lock.Lock()

	defer fmtAbsValMolecule.lock.Unlock()

	newAbsValMolecule :=
		new(formatterAbsoluteValueMolecule)

	newAbsValMolecule.lock = new(sync.Mutex)

	return newAbsValMolecule
}

// testValidityOfFormatterAbsoluteValue - Receives an instance of
// FormatterAbsoluteValue and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (fmtAbsValMolecule *formatterAbsoluteValueMolecule) testValidityOfFormatterAbsoluteValue(
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
			"testValidityOfFormatterAbsoluteValue()")

	isValid = false

	if fmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	nStrAbsValDtoAtom :=
		formatterAbsoluteValueAtom{}

	isValid,
		err = nStrAbsValDtoAtom.testAbsoluteValueFormat(
		fmtAbsoluteVal,
		ePrefix.XCtx("Validating fmtAbsoluteVal Format"))

	if err != nil {
		return isValid, err
	}

	if fmtAbsoluteVal.numStrFmtType !=
		NumStrFormatTypeCode(0).AbsoluteValue() {
		err = fmt.Errorf("%v\n"+
			"Error: Number String Format Type Code is invalid!\n"+
			"integer value of fmtAbsoluteVal.numStrFmtType= '%v'\n",
			ePrefix.XCtxEmpty().String(),
			fmtAbsoluteVal.numStrFmtType.XValueInt())
		return isValid, err
	}

	err =
		fmtAbsoluteVal.numericSeparators.IsValidInstanceError(
			ePrefix.XCtx(
				"Validating fmtAbsoluteVal"))

	if err != nil {
		return isValid, err
	}

	err =
		fmtAbsoluteVal.numFieldLenDto.IsValidInstanceError(
			ePrefix)

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
