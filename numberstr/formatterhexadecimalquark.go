package numberstr

import (
	"fmt"
	"sync"
)

type formatterHexadecimalQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a FormatterHexadecimal instance to their
// initial 'zero' values.
//
func (fmtHexadecimalQuark *formatterHexadecimalQuark) empty(
	formatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalQuark.lock == nil {
		fmtHexadecimalQuark.lock = new(sync.Mutex)
	}

	fmtHexadecimalQuark.lock.Lock()

	defer fmtHexadecimalQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalQuark." +
			"empty()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	formatterHex.useUpperCaseLetters = false
	formatterHex.leftPrefix = ""
	formatterHex.turnOnIntegerDigitsSeparation = false
	formatterHex.integerSeparators.Empty()
	formatterHex.numFieldDto.Empty()

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterHexadecimalQuark.
//
func (fmtHexadecimalQuark formatterHexadecimalQuark) ptr() *formatterHexadecimalQuark {

	if fmtHexadecimalQuark.lock == nil {
		fmtHexadecimalQuark.lock = new(sync.Mutex)
	}

	fmtHexadecimalQuark.lock.Lock()

	defer fmtHexadecimalQuark.lock.Unlock()

	newFmtHexadecimalQuark :=
		new(formatterHexadecimalQuark)

	newFmtHexadecimalQuark.lock = new(sync.Mutex)

	return newFmtHexadecimalQuark
}

// testValidityOfFormatterHexadecimal - Receives an instance of
// FormatterHexadecimal and proceeds to test the validity
// of the member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (fmtHexadecimalQuark *formatterHexadecimalQuark) testValidityOfFormatterHexadecimal(
	formatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtHexadecimalQuark.lock == nil {
		fmtHexadecimalQuark.lock = new(sync.Mutex)
	}

	fmtHexadecimalQuark.lock.Lock()

	defer fmtHexadecimalQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalQuark." +
			"testValidityOfFormatterHexadecimal()")

	isValid = false

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is invalid!\n"+
			"'formatterHex' is a 'nil' pointer.\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterHex.numStrFmtType !=
		NumStrFormatTypeCode(0).Hexadecimal() {

		err = fmt.Errorf("%v\n"+
			"Error: Data Element Number String Format Type is invalid!\n"+
			"'formatterHex.numStrFmtType' should BE EQUAL TO "+
			"'NumStrFormatTypeCode(0).Hexadecimal()'.\n"+
			"Instead, the 'formatterHex.numStrFmtType' "+
			"integer value= '%v'\n",
			ePrefix.String(),
			formatterHex.numStrFmtType.XValueInt())

		return isValid, err
	}

	err =
		formatterHex.integerSeparators.IsValidInstanceError(
			ePrefix.XCtx(
				"Validating 'formatterHex' Integer Separators"))

	if err != nil {
		return isValid, err
	}

	err =
		formatterHex.numFieldDto.IsValidInstanceError(
			ePrefix.XCtx("Validating 'formatterHex' Number Field Dto"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
