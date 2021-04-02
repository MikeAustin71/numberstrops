package numberstr

import (
	"fmt"
	"sync"
)

type formatterBinaryQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a FormatterBinary instance to their initial
// 'zero' values.
//
func (fmtBinaryQuark formatterBinaryQuark) empty(
	formatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryQuark.lock == nil {
		fmtBinaryQuark.lock = new(sync.Mutex)
	}

	fmtBinaryQuark.lock.Lock()

	defer fmtBinaryQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryQuark." +
			"empty()")

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	formatterBinary.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	formatterBinary.turnOnIntegerDigitsSeparation =
		false

	formatterBinary.numericSeparators.Empty()

	formatterBinary.numFieldDto.Empty()

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterBinaryQuark.
//
func (fmtBinaryQuark formatterBinaryQuark) ptr() *formatterBinaryQuark {

	if fmtBinaryQuark.lock == nil {
		fmtBinaryQuark.lock = new(sync.Mutex)
	}

	fmtBinaryQuark.lock.Lock()

	defer fmtBinaryQuark.lock.Unlock()

	newBinaryQuark :=
		new(formatterBinaryQuark)

	newBinaryQuark.lock = new(sync.Mutex)

	return newBinaryQuark
}

// testValidityOfFormatterBinary - Receives an instance of
// FormatterBinary and proceeds to test the validity of the member
// data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (fmtBinaryQuark formatterBinaryQuark) testValidityOfFormatterBinary(
	formatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtBinaryQuark.lock == nil {
		fmtBinaryQuark.lock = new(sync.Mutex)
	}

	fmtBinaryQuark.lock.Lock()

	defer fmtBinaryQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryQuark." +
			"testValidityOfFormatterBinary()")

	isValid = false

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterBinary.numStrFmtType !=
		NumStrFormatTypeCode(0).Binary() {

		err = fmt.Errorf("%v\n"+
			"Error: Data Element Number String Format Type is invalid!\n"+
			"'formatterBinary.numStrFmtType' should BE EQUAL TO "+
			"'NumStrFormatTypeCode(0).Binary()'.\n"+
			"Instead, the 'formatterBinary.numStrFmtType' "+
			"integer value= '%v'\n",
			ePrefix.String(),
			formatterBinary.numStrFmtType.XValueInt())

		return isValid, err
	}

	err =
		formatterBinary.numericSeparators.IsValidInstanceError(
			ePrefix.XCtx(
				"Validating 'formatterHex' Integer Separators"))

	if err != nil {
		return isValid, err
	}

	err =
		formatterBinary.numFieldDto.IsValidInstanceError(
			ePrefix.XCtx("Validating 'formatterHex' Number Field Dto"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
