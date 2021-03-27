package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFmtAbsoluteVal' to input parameter
// 'targetFmtAbsoluteVal'.
//
// Be advised - All data fields in 'targetFmtAbsoluteVal'
// will be overwritten.
//
// If the 'incomingFmtAbsoluteVal' is judged to be
// invalid, this method will return an error.
//
func (fmtAbsValNanobot *formatterAbsoluteValueNanobot) copyIn(
	targetFmtAbsoluteVal *FormatterAbsoluteValue,
	incomingFmtAbsoluteVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValNanobot.lock == nil {
		fmtAbsValNanobot.lock = new(sync.Mutex)
	}

	fmtAbsValNanobot.lock.Lock()

	defer fmtAbsValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueNanobot.copyIn()")

	if targetFmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingFmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = formatterAbsoluteValueMolecule{}.ptr().
		testValidityOfFormatterAbsoluteValue(
			incomingFmtAbsoluteVal,
			ePrefix.XCtx(
				"Testing validity of 'incomingFmtAbsoluteVal'"))

	if err != nil {
		return err
	}

	targetFmtAbsoluteVal.numStrFmtType =
		incomingFmtAbsoluteVal.numStrFmtType

	targetFmtAbsoluteVal.absoluteValFmt =
		incomingFmtAbsoluteVal.absoluteValFmt

	targetFmtAbsoluteVal.turnOnIntegerDigitsSeparation =
		incomingFmtAbsoluteVal.turnOnIntegerDigitsSeparation

	err =
		targetFmtAbsoluteVal.numericSeparators.CopyIn(
			&incomingFmtAbsoluteVal.numericSeparators,
			ePrefix.XCtx(
				"Copying incomingFmtAbsoluteVal->"+
					"targetFmtAbsoluteVal"))

	if err != nil {
		return err
	}

	err =
		targetFmtAbsoluteVal.numFieldDto.CopyIn(
			&incomingFmtAbsoluteVal.numFieldDto,
			ePrefix.XCtx(
				"incomingFmtAbsoluteVal->"+
					"targetFmtAbsoluteVal"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'fmtAbsoluteVal' styled as a new instance
// of FormatterAbsoluteValue.
//
// If input parameter 'fmtAbsoluteVal' is judged to be
// invalid, this method will return an error.
//
func (fmtAbsValNanobot *formatterAbsoluteValueNanobot) copyOut(
	fmtAbsoluteVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	newFmtAbsoluteVal FormatterAbsoluteValue,
	err error) {

	if fmtAbsValNanobot.lock == nil {
		fmtAbsValNanobot.lock = new(sync.Mutex)
	}

	fmtAbsValNanobot.lock.Lock()

	defer fmtAbsValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueNanobot.copyOut()")

	if fmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newFmtAbsoluteVal, err
	}

	_,
		err = formatterAbsoluteValueMolecule{}.ptr().
		testValidityOfFormatterAbsoluteValue(
			fmtAbsoluteVal,
			ePrefix.XCtx("\nTesting validity of 'fmtAbsoluteVal'"))

	if err != nil {
		return newFmtAbsoluteVal, err
	}

	newFmtAbsoluteVal.numStrFmtType =
		fmtAbsoluteVal.numStrFmtType

	newFmtAbsoluteVal.absoluteValFmt =
		fmtAbsoluteVal.absoluteValFmt

	newFmtAbsoluteVal.turnOnIntegerDigitsSeparation =
		fmtAbsoluteVal.turnOnIntegerDigitsSeparation

	err = newFmtAbsoluteVal.numericSeparators.CopyIn(
		&fmtAbsoluteVal.numericSeparators,
		ePrefix.XCtx(
			"fmtAbsoluteVal->"+
				"newFmtAbsoluteVal"))

	if err != nil {
		return newFmtAbsoluteVal, err
	}

	newFmtAbsoluteVal.lock = new(sync.Mutex)

	err =
		newFmtAbsoluteVal.numFieldDto.CopyIn(
			&fmtAbsoluteVal.numFieldDto,
			ePrefix.XCtx(
				"StrFmtSpecAbsoluteValDto->newFmtAbsoluteVal"))

	return newFmtAbsoluteVal, err
}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueNanobot.
//
func (fmtAbsValNanobot formatterAbsoluteValueNanobot) ptr() *formatterAbsoluteValueNanobot {

	if fmtAbsValNanobot.lock == nil {
		fmtAbsValNanobot.lock = new(sync.Mutex)
	}

	fmtAbsValNanobot.lock.Lock()

	defer fmtAbsValNanobot.lock.Unlock()

	newAbsValNanobot :=
		new(formatterAbsoluteValueNanobot)

	newAbsValNanobot.lock = new(sync.Mutex)

	return newAbsValNanobot
}
