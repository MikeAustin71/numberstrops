package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalQuark struct {
	lock *sync.Mutex
}

func (fmtOctalQuark *formatterOctalQuark) empty(
	formatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalQuark.lock == nil {
		fmtOctalQuark.lock = new(sync.Mutex)
	}

	fmtOctalQuark.lock.Lock()

	defer fmtOctalQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalQuark." +
			"empty()")

	if formatterOctal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	formatterOctal.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	formatterOctal.leftPrefix = ""

	formatterOctal.rightSuffix = ""

	formatterOctal.turnOnIntegerDigitsSeparation = false

	formatterOctal.integerSeparators.Empty()

	formatterOctal.numFieldDto.Empty()

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterOctalQuark.
//
func (fmtOctalQuark formatterOctalQuark) ptr() *formatterOctalQuark {

	if fmtOctalQuark.lock == nil {
		fmtOctalQuark.lock = new(sync.Mutex)
	}

	fmtOctalQuark.lock.Lock()

	defer fmtOctalQuark.lock.Unlock()

	newFmtOctalQuark :=
		new(formatterOctalQuark)

	newFmtOctalQuark.lock = new(sync.Mutex)

	return newFmtOctalQuark
}

// testValidityOfFormatterOctal - Receives an instance of type
// FormatterOctal and proceeds to test the validity of the member
// data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (fmtOctalQuark *formatterOctalQuark) testValidityOfFormatterOctal(
	formatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtOctalQuark.lock == nil {
		fmtOctalQuark.lock = new(sync.Mutex)
	}

	fmtOctalQuark.lock.Lock()

	defer fmtOctalQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalQuark." +
			"testValidityOfFormatterOctal()")

	isValid = false

	if formatterOctal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	if formatterOctal.numStrFmtType !=
		NumStrFormatTypeCode(0).Octal() {

		err = fmt.Errorf("%v\n"+
			"Error: Data Element Number String Format Type is invalid!\n"+
			"'formatterHex.numStrFmtType' should BE EQUAL TO "+
			"'NumStrFormatTypeCode(0).Octal()'.\n"+
			"Instead, the 'formatterOctal.numStrFmtType' "+
			"integer value= '%v'\n",
			ePrefix.String(),
			formatterOctal.numStrFmtType.XValueInt())

		return isValid, err
	}

	err =
		formatterOctal.integerSeparators.IsValidInstanceError(
			ePrefix.XCtx(
				"Validating 'formatterOctal' Integer Separators"))

	if err != nil {
		return isValid, err
	}

	err =
		formatterOctal.numFieldDto.IsValidInstanceError(
			ePrefix.XCtx("Validating 'formatterOctal' Number Field Dto"))

	if err != nil {
		return isValid, err
	}

	return isValid, err
}
