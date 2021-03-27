package numberstr

import (
	"fmt"
	"sync"
)

type formatterSciNotationQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a FormatterSciNotation instance to their
// initial 'zero' values.
//
func (fmtSciNotQuark *formatterSciNotationQuark) empty(
	formatterSicNot *FormatterSciNotation,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotQuark.lock == nil {
		fmtSciNotQuark.lock = new(sync.Mutex)
	}

	fmtSciNotQuark.lock.Lock()

	defer fmtSciNotQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationQuark." +
			"empty()")

	if formatterSicNot == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterSicNot' is invalid!\n"+
			"'formatterSicNot' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	formatterSicNot.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	formatterSicNot.significandUsesLeadingPlus = false

	formatterSicNot.mantissaLength = 0

	formatterSicNot.exponentChar = 0

	formatterSicNot.exponentUsesLeadingPlus = false

	formatterSicNot.numFieldDto.Empty()

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterSciNotationQuark.
//
func (fmtSciNotQuark formatterSciNotationQuark) ptr() *formatterSciNotationQuark {

	if fmtSciNotQuark.lock == nil {
		fmtSciNotQuark.lock = new(sync.Mutex)
	}

	fmtSciNotQuark.lock.Lock()

	defer fmtSciNotQuark.lock.Unlock()

	newSciNotQuark :=
		new(formatterSciNotationQuark)

	newSciNotQuark.lock = new(sync.Mutex)

	return newSciNotQuark
}

// testValidityOfFormatterSciNotation - Receives an instance
// of FormatterSciNotation and proceeds to test the validity
// of the member data fields. The FormatterSciNotation
// object encapsulates scientific notation formatting for number
// string text displays.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
// If nStrFmtSpecSicNotDto.mantissaLength is greater than 1000,
// this method will return an error.
//
//
func (fmtSciNotQuark *formatterSciNotationQuark) testValidityOfFormatterSciNotation(
	formatterSciNotation *FormatterSciNotation,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtSciNotQuark.lock == nil {
		fmtSciNotQuark.lock = new(sync.Mutex)
	}

	fmtSciNotQuark.lock.Lock()

	defer fmtSciNotQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationQuark." +
			"testValidityOfFormatterSciNotation()")

	isValid = false

	if formatterSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterSciNotation' is invalid!\n"+
			"'formatterSciNotation' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterSciNotation.lock == nil {
		formatterSciNotation.lock = new(sync.Mutex)
	}

	if formatterSciNotation.mantissaLength > 10000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'formatterSciNotation.mantissaLength' is invalid!\n"+
			"formatterSciNotation.mantissaLength is greater than 10,000.\n"+
			"formatterSciNotation.mantissaLength== '%v'\n",
			ePrefix.String(),
			formatterSciNotation.mantissaLength)

		return isValid, err
	}

	if formatterSciNotation.mantissaLength == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'formatterSciNotation.mantissaLength' is invalid!\n"+
			"formatterSciNotation.mantissaLength is equal to ZERO ('0').\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterSciNotation.numStrFmtType !=
		NumStrFormatTypeCode(0).ScientificNotation() {

		err = fmt.Errorf("%v\n"+
			"Error: Number String Format Code is invalid!\n"+
			"'formatterSciNotation.numStrFmtType' IS NOT EQUAL TO\n"+
			"NumStrFormatTypeCode(0).ScientificNotation().\n"+
			"'formatterSciNotation.numStrFmtType' integer value='%v'\n",
			ePrefix.String(),
			formatterSciNotation.numStrFmtType.XValueInt())

		return isValid, err
	}

	if formatterSciNotation.mantissaLength == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: 'formatterSciNotation.mantissaLength' is invalid!\n" +
			"formatterSciNotation.mantissaLength is equal to zero.\n" +
			ePrefix.String())

		return isValid, err
	}

	if formatterSciNotation.exponentChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'formatterSciNotation.exponentChar' is invalid!\n"+
			"formatterSciNotation.exponentChar== '%v'\n",
			ePrefix.String(),
			formatterSciNotation.exponentChar)

		return isValid, err
	}

	err = formatterSciNotation.numFieldDto.IsValidInstanceError(
		ePrefix.XCtx(
			"\nTesting Validity of formatterSciNotation.numFieldDto"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
