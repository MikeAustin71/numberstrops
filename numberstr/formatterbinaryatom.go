package numberstr

import (
	"fmt"
	"sync"
)

type formatterBinaryAtom struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFormatterBinary' to input parameter
// 'targetFormatterBinary'.
//
// Be advised - All data fields in 'targetFormatterBinary'
// will be overwritten.
//
// If input parameter 'incomingFormatterBinary' is judged
// to be invalid, this method will return an error.
//
func (fmtBinaryAtom *formatterBinaryAtom) copyIn(
	targetFormatterBinary *FormatterBinary,
	incomingFormatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryAtom.lock == nil {
		fmtBinaryAtom.lock = new(sync.Mutex)
	}

	fmtBinaryAtom.lock.Lock()

	defer fmtBinaryAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryAtom." +
			"copyIn()")

	if targetFormatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFormatterBinary' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetFormatterBinary.lock == nil {
		targetFormatterBinary.lock = new(sync.Mutex)
	}

	if incomingFormatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFormatterBinary' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err = formatterBinaryQuark{}.ptr().
		testValidityOfFormatterBinary(
			incomingFormatterBinary,
			ePrefix.XCtx("Checking incomingFormatterBinary"))

	if err != nil {
		return err
	}

	targetFormatterBinary.numStrFmtType =
		NumStrFormatTypeCode(0).Binary()

	targetFormatterBinary.turnOnIntegerDigitsSeparation =
		incomingFormatterBinary.turnOnIntegerDigitsSeparation

	err = targetFormatterBinary.numericSeparators.CopyIn(
		&incomingFormatterBinary.numericSeparators,
		ePrefix.XCtx(
			"incomingFormatterBinary.numericSeparators->"+
				"targetFormatterBinary.numericSeparators"))

	if err != nil {
		return err
	}

	err = targetFormatterBinary.numFieldDto.CopyIn(
		&incomingFormatterBinary.numFieldDto,
		ePrefix.XCtx(
			"incomingFormatterBinary.numFieldDto->"+
				"targetFormatterBinary.numFieldDto"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'formatterBinary' styled as a new instance of
// FormatterBinary.
//
// If input parameter 'formatterBinary' is judged to be invalid, this
// method will return an error.
//
func (fmtBinaryAtom *formatterBinaryAtom) copyOut(
	formatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	newFormatterBinary FormatterBinary,
	err error) {

	if fmtBinaryAtom.lock == nil {
		fmtBinaryAtom.lock = new(sync.Mutex)
	}

	fmtBinaryAtom.lock.Lock()

	defer fmtBinaryAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryAtom." +
			"copyOut()")

	if formatterBinary == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is a 'nil' pointer!\n",
			ePrefix.String())

		return newFormatterBinary, err
	}

	_,
		err = formatterBinaryQuark{}.ptr().
		testValidityOfFormatterBinary(
			formatterBinary,
			ePrefix.XCtx("Checking formatterBinary"))

	if err != nil {
		return newFormatterBinary, err
	}

	newFormatterBinary.lock = new(sync.Mutex)

	newFormatterBinary.numStrFmtType =
		formatterBinary.numStrFmtType

	newFormatterBinary.turnOnIntegerDigitsSeparation =
		formatterBinary.turnOnIntegerDigitsSeparation

	err = newFormatterBinary.numericSeparators.
		CopyIn(
			&formatterBinary.numericSeparators,
			ePrefix.XCtx(
				"formatterBinary.numericSeparators->"+
					"newFormatterBinary.numericSeparators"))

	if err != nil {
		return newFormatterBinary, err
	}

	err = newFormatterBinary.numFieldDto.CopyIn(
		&formatterBinary.numFieldDto,
		ePrefix.XCtx(
			"formatterBinary.numFieldDto->"+
				"newFormatterBinary.numFieldDto"))

	return newFormatterBinary, err
}

// equal - Receives two FormatterBinary objects and proceeds to
// determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtBinaryOne        *FormatterBinary
//     - A pointer to the first FormatterBinary object. This
//       method will compare all data elements in this object to
//       all corresponding data elements in the second
//       FormatterBinary object to determine if they are
//       equivalent.
//
//
//  fmtBinaryTwo        *FormatterBinary
//     - A pointer to the second FormatterBinary object. This
//       method will compare all data elements in the first
//       FormatterBinary object to all corresponding data
//       elements in this second FormatterBinary object to
//       determine if they are equivalent.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isEqual             bool
//     - If all the data elements in 'fmtBinaryOne' are equal to
//       all the corresponding data elements in 'fmtBinaryTwo',
//       this return parameter will be set to 'true'. If the data
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtBinaryOne' are equal to
//       all the corresponding data elements in 'fmtBinaryTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtBinaryAtom *formatterBinaryAtom) equal(
	fmtBinaryOne *FormatterBinary,
	fmtBinaryTwo *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtBinaryAtom.lock == nil {
		fmtBinaryAtom.lock = new(sync.Mutex)
	}

	fmtBinaryAtom.lock.Lock()

	defer fmtBinaryAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryAtom." +
			"equal()")

	isEqual = false

	if fmtBinaryOne == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtBinaryOne' is a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtBinaryTwo == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtBinaryTwo' is a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtBinaryOne.numStrFmtType !=
		fmtBinaryTwo.numStrFmtType {

		err = fmt.Errorf("%v\n"+
			"fmtBinaryOne.numStrFmtType!=fmtBinaryTwo.numStrFmtType\n"+
			"fmtBinaryOne.numStrFmtType='%v'\n"+
			"fmtBinaryTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtBinaryOne.numStrFmtType.String(),
			fmtBinaryTwo.numStrFmtType.String())

		return isEqual, err
	}

	if fmtBinaryOne.turnOnIntegerDigitsSeparation !=
		fmtBinaryTwo.turnOnIntegerDigitsSeparation {

		err = fmt.Errorf("%v\n"+
			"fmtBinaryOne.turnOnIntegerDigitsSeparation!="+
			"fmtBinaryTwo.turnOnIntegerDigitsSeparation\n"+
			"fmtBinaryOne.turnOnIntegerDigitsSeparation='%v'\n"+
			"fmtBinaryTwo.turnOnIntegerDigitsSeparation='%v'\n",
			ePrefix.String(),
			fmtBinaryOne.turnOnIntegerDigitsSeparation,
			fmtBinaryTwo.turnOnIntegerDigitsSeparation)

		return isEqual, err
	}

	_,
		err = fmtBinaryOne.numericSeparators.Equal(
		fmtBinaryTwo.numericSeparators,
		ePrefix.XCtx(
			"fmtBinaryOne vs fmtBinaryTwo"))

	if err != nil {
		return isEqual, err
	}

	_,
		err = fmtBinaryOne.numFieldDto.Equal(
		fmtBinaryTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtBinaryOne vs fmtBinaryTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
}

// ptr - Returns a pointer to a new instance of
// formatterBinaryAtom.
//
func (fmtBinaryAtom formatterBinaryAtom) ptr() *formatterBinaryAtom {

	if fmtBinaryAtom.lock == nil {
		fmtBinaryAtom.lock = new(sync.Mutex)
	}

	fmtBinaryAtom.lock.Lock()

	defer fmtBinaryAtom.lock.Unlock()

	newBinaryAtom :=
		new(formatterBinaryAtom)

	newBinaryAtom.lock = new(sync.Mutex)

	return newBinaryAtom
}
