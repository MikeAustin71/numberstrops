package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalAtom struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFormatterOctal' to input parameter
// 'targetFormatterOctal'.
//
// Be advised - All data fields in 'targetFormatterOctal'
// will be overwritten.
//
// If input parameter 'incomingFormatterOctal' is judged
// to be invalid, this method will return an error.
//
func (fmtOctalAtom *formatterOctalAtom) copyIn(
	targetFormatterOctal *FormatterOctal,
	incomingFormatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalAtom.lock == nil {
		fmtOctalAtom.lock = new(sync.Mutex)
	}

	fmtOctalAtom.lock.Lock()

	defer fmtOctalAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalAtom." +
			"copyIn()")

	if targetFormatterOctal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFormatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetFormatterOctal.lock == nil {
		targetFormatterOctal.lock = new(sync.Mutex)
	}

	if incomingFormatterOctal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFormatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err = formatterOctalQuark{}.ptr().
		testValidityOfFormatterOctal(
			incomingFormatterOctal,
			ePrefix.XCtx(
				"Testing validity of 'incomingFormatterOctal'"))

	if err != nil {
		return err
	}

	targetFormatterOctal.numStrFmtType =
		NumStrFormatTypeCode(0).Octal()

	targetFormatterOctal.leftPrefix =
		incomingFormatterOctal.leftPrefix

	targetFormatterOctal.leftPrefix =
		incomingFormatterOctal.leftPrefix

	targetFormatterOctal.rightSuffix =
		incomingFormatterOctal.rightSuffix

	targetFormatterOctal.turnOnIntegerDigitsSeparation =
		incomingFormatterOctal.turnOnIntegerDigitsSeparation

	err =
		targetFormatterOctal.numericSeparators.CopyIn(
			&incomingFormatterOctal.numericSeparators,
			ePrefix.XCtx(
				"incomingFormatterOctal.integerSeparators->"+
					"targetFormatterOctal.integerSeparators"))

	if err != nil {
		return err
	}

	err =
		targetFormatterOctal.numFieldDto.CopyIn(
			&incomingFormatterOctal.numFieldDto,
			ePrefix.XCtx(
				"incomingFormatterOctal.numFieldDto->"+
					"targetFormatterOctal.numFieldDto"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'formatterOctal' styled as a new instance of
// FormatterOctal.
//
// If input parameter 'formatterOctal' is judged to be invalid, this
// method will return an error.
//
func (fmtOctalAtom *formatterOctalAtom) copyOut(
	formatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	newFormatterOctal FormatterOctal,
	err error) {

	if fmtOctalAtom.lock == nil {
		fmtOctalAtom.lock = new(sync.Mutex)
	}

	fmtOctalAtom.lock.Lock()

	defer fmtOctalAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalAtom." +
			"copyOut()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return newFormatterOctal, err
	}

	_,
		err = formatterOctalQuark{}.ptr().
		testValidityOfFormatterOctal(
			formatterOctal,
			ePrefix.XCtx(
				"Testing validity of "+
					"'formatterOctal'"))

	if err != nil {
		return newFormatterOctal, err
	}

	newFormatterOctal.lock = new(sync.Mutex)

	newFormatterOctal.numStrFmtType =
		formatterOctal.numStrFmtType

	newFormatterOctal.leftPrefix =
		formatterOctal.leftPrefix

	newFormatterOctal.rightSuffix =
		formatterOctal.rightSuffix

	newFormatterOctal.turnOnIntegerDigitsSeparation =
		formatterOctal.turnOnIntegerDigitsSeparation

	err =
		newFormatterOctal.numericSeparators.CopyIn(
			&formatterOctal.numericSeparators,
			ePrefix.XCtx(""+
				"formatterOctal.integerSeparators->"+
				"newFormatterOctal.integerSeparators"))

	if err != nil {
		return newFormatterOctal, err
	}

	err =
		newFormatterOctal.numFieldDto.CopyIn(
			&formatterOctal.numFieldDto,
			ePrefix.XCtx(""+
				"formatterOctal.numFieldDto->"+
				"newFormatterOctal.numFieldDto"))

	return newFormatterOctal, err
}

// equal - Receives two FormatterOctal objects and proceeds to
// determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtOctalOne         *FormatterOctal
//     - A pointer to the first FormatterOctal object. This
//       method will compare all data elements in this object to
//       all corresponding data elements in the second
//       FormatterOctal object to determine if they are equivalent.
//
//
//  fmtOctalTwo         *FormatterOctal
//     - A pointer to the second FormatterOctal object. This method
//       will compare all data elements in the first FormatterOctal
//       object to all corresponding data elements in this second
//       FormatterOctal object to determine if they are equivalent.
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
//     - If all the data elements in 'fmtOctalOne' are equal to
//       all the corresponding data elements in 'fmtOctalTwo',
//       this return parameter will be set to 'true'. If the data
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtOctalOne' are equal to
//       all the corresponding data elements in 'fmtOctalTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtOctalAtom *formatterOctalAtom) equal(
	fmtOctalOne *FormatterOctal,
	fmtOctalTwo *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtOctalAtom.lock == nil {
		fmtOctalAtom.lock = new(sync.Mutex)
	}

	fmtOctalAtom.lock.Lock()

	defer fmtOctalAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalAtom." +
			"equal()")

	isEqual = false

	if fmtOctalOne == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtOctalOne' is a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtOctalTwo == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtOctalTwo' is a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtOctalOne.numStrFmtType !=
		fmtOctalTwo.numStrFmtType {

		err = fmt.Errorf("%v\n"+
			"fmtOctalOne.numStrFmtType!=fmtOctalTwo.numStrFmtType\n"+
			"fmtOctalOne.numStrFmtType='%v'\n"+
			"fmtOctalTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtOctalOne.numStrFmtType.String(),
			fmtOctalTwo.numStrFmtType.String())

		return isEqual, err
	}

	if fmtOctalOne.leftPrefix !=
		fmtOctalTwo.leftPrefix {

		err = fmt.Errorf("%v\n"+
			"fmtOctalOne.leftPrefix!=fmtOctalTwo.leftPrefix\n"+
			"fmtOctalOne.leftPrefix='%v'\n"+
			"fmtOctalTwo.leftPrefix='%v'\n",
			ePrefix.String(),
			fmtOctalOne.leftPrefix,
			fmtOctalTwo.leftPrefix)

		return isEqual, err
	}

	if fmtOctalOne.rightSuffix !=
		fmtOctalTwo.rightSuffix {

		err = fmt.Errorf("%v\n"+
			"fmtOctalOne.rightSuffix!=fmtOctalTwo.rightSuffix\n"+
			"fmtOctalOne.rightSuffix='%v'\n"+
			"fmtOctalTwo.rightSuffix='%v'\n",
			ePrefix.String(),
			fmtOctalOne.rightSuffix,
			fmtOctalTwo.rightSuffix)

		return isEqual, err
	}

	if fmtOctalOne.turnOnIntegerDigitsSeparation !=
		fmtOctalTwo.turnOnIntegerDigitsSeparation {

		err = fmt.Errorf("%v\n"+
			"fmtOctalOne.turnOnIntegerDigitsSeparation!=\n"+
			"  fmtOctalTwo.turnOnIntegerDigitsSeparation\n"+
			"fmtOctalOne.turnOnIntegerDigitsSeparation='%v'\n"+
			"fmtOctalTwo.turnOnIntegerDigitsSeparation='%v'\n",
			ePrefix.String(),
			fmtOctalOne.turnOnIntegerDigitsSeparation,
			fmtOctalTwo.turnOnIntegerDigitsSeparation)

		return isEqual, err
	}

	_,
		err = fmtOctalOne.numericSeparators.Equal(
		fmtOctalTwo.numericSeparators,
		ePrefix.XCtx(
			"fmtOctalOne vs fmtOctalTwo"))

	if err != nil {
		return isEqual, err
	}

	_,
		err = fmtOctalOne.numFieldDto.Equal(
		fmtOctalTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtOctalOne vs fmtOctalTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
}

// ptr - Returns a pointer to a new instance of
// formatterOctalAtom.
//
func (fmtOctalAtom formatterOctalAtom) ptr() *formatterOctalAtom {

	if fmtOctalAtom.lock == nil {
		fmtOctalAtom.lock = new(sync.Mutex)
	}

	fmtOctalAtom.lock.Lock()

	defer fmtOctalAtom.lock.Unlock()

	newFmtOctalAtom :=
		new(formatterOctalAtom)

	newFmtOctalAtom.lock = new(sync.Mutex)

	return newFmtOctalAtom
}
