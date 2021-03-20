package numberstr

import (
	"fmt"
	"sync"
)

type formatterSignedNumberMolecule struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterSignedNumberMolecule.
//
func (fmtSignedNumMolecule formatterSignedNumberMolecule) ptr() *formatterSignedNumberMolecule {

	if fmtSignedNumMolecule.lock == nil {
		fmtSignedNumMolecule.lock = new(sync.Mutex)
	}

	fmtSignedNumMolecule.lock.Lock()

	defer fmtSignedNumMolecule.lock.Unlock()

	newFmtSignedNumMolecule :=
		new(formatterSignedNumberMolecule)

	newFmtSignedNumMolecule.lock = new(sync.Mutex)

	return newFmtSignedNumMolecule
}

// testValidityOfSignedNumValDto - Receives an instance of
// FormatterSignedNumber and proceeds to test the validity
// of the member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (fmtSignedNumMolecule *formatterSignedNumberMolecule) testValidityOfSignedNumValDto(
	fmtSignedNum *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtSignedNumMolecule.lock == nil {
		fmtSignedNumMolecule.lock = new(sync.Mutex)
	}

	fmtSignedNumMolecule.lock.Lock()

	defer fmtSignedNumMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberMolecule." +
			"testValidityOfSignedNumValDto()")

	isValid = false

	if fmtSignedNum == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSignedNum' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if fmtSignedNum.lock == nil {
		fmtSignedNum.lock = new(sync.Mutex)
	}

	nStrSignedNumAtom :=
		formatterSignedNumberAtom{}

	isValid,
		err = nStrSignedNumAtom.testSignedNumValPositiveValueFormat(
		fmtSignedNum,
		ePrefix.XCtx("Validating fmtSignedNum Positive Value Format"))

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrSignedNumAtom.testSignedNumValNegativeValueFormat(
		fmtSignedNum,
		ePrefix.XCtx("Validating fmtSignedNum Negative Value Format"))

	if err != nil {
		return isValid, err
	}

	if fmtSignedNum.numStrFormatterType !=
		NumStrFormatTypeCode(0).SignedNumber() {

		err = fmt.Errorf("%v\n"+
			"Error: Number String Format Type Code is invalid!\n"+
			"'fmtSignedNum.numStrFormatterType' should be equal to NumStrFormatTypeCode(0).SignedNumber().\n"+
			"Instead, 'fmtSignedNum.numStrFormatterType' is equal to integer value '%v'\n",
			ePrefix.XCtxEmpty().String(),
			fmtSignedNum.numStrFormatterType.XValueInt())

		return isValid, err
	}

	err =
		fmtSignedNum.numericSeparators.IsValidInstanceError(
			ePrefix.XCtx("Validating 'fmtSignedNum' Number Separators"))

	if err != nil {
		return isValid, err
	}

	err =
		fmtSignedNum.numFieldLenDto.IsValidInstanceError(
			ePrefix.XCtx("Validating 'fmtSignedNum' Number Field Dto"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
