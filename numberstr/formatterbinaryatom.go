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
