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
		targetFormatterOctal.integerSeparators.CopyIn(
			&incomingFormatterOctal.integerSeparators,
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
		newFormatterOctal.integerSeparators.CopyIn(
			&formatterOctal.integerSeparators,
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
