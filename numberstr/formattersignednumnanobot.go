package numberstr

import (
	"fmt"
	"sync"
)

type formatterSignedNumberNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecSignedNumValDto' to input parameter
// 'targetNStrFmtSpecSignedNumValDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecSignedNumValDto'
// will be overwritten.
//
// If input parameter 'inComingNStrFmtSpecSignedNumValDto' is judged
// to be invalid, this method will return an error.
//
func (fmtSignedNumNanobot *formatterSignedNumberNanobot) copyIn(
	targetFmtSignedNum *FormatterSignedNumber,
	incomingFmtSignedNum *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSignedNumNanobot.lock == nil {
		fmtSignedNumNanobot.lock = new(sync.Mutex)
	}

	fmtSignedNumNanobot.lock.Lock()

	defer fmtSignedNumNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("formatterSignedNumberNanobot.copyIn()")

	if targetFmtSignedNum == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFmtSignedNum' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetFmtSignedNum.lock == nil {
		targetFmtSignedNum.lock = new(sync.Mutex)
	}

	if incomingFmtSignedNum == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFmtSignedNum' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		formatterSignedNumberMolecule{}.ptr().
			testValidityOfSignedNumValDto(
				incomingFmtSignedNum,
				ePrefix.XCtx("Testing validity of 'incomingFmtSignedNum'"))

	if err != nil {
		return err
	}

	targetFmtSignedNum.numStrFormatterType =
		incomingFmtSignedNum.numStrFormatterType

	targetFmtSignedNum.positiveValueFmt =
		incomingFmtSignedNum.positiveValueFmt

	targetFmtSignedNum.negativeValueFmt =
		incomingFmtSignedNum.negativeValueFmt

	targetFmtSignedNum.turnOnIntegerDigitsSeparation =
		incomingFmtSignedNum.turnOnIntegerDigitsSeparation

	err =
		targetFmtSignedNum.numericSeparators.CopyIn(
			&incomingFmtSignedNum.numericSeparators,
			ePrefix.XCtx("'incomingFmtSignedNum' -> "+
				"'targetFmtSignedNum'"))

	if err != nil {
		return err
	}

	err =
		targetFmtSignedNum.numFieldLenDto.CopyIn(
			&incomingFmtSignedNum.numFieldLenDto,
			ePrefix.XCtx("incomingFmtSignedNum->targetFmtSignedNum"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecSignedNumValDto' styled as a new instance
// of FormatterSignedNumber.
//
// If input parameter 'nStrFmtSpecSignedNumValDto' is judged to be
// invalid, this method will return an error.
//
func (fmtSignedNumNanobot *formatterSignedNumberNanobot) copyOut(
	fmtSignedNum *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	newFmtSignedNum FormatterSignedNumber,
	err error) {

	if fmtSignedNumNanobot.lock == nil {
		fmtSignedNumNanobot.lock = new(sync.Mutex)
	}

	fmtSignedNumNanobot.lock.Lock()

	defer fmtSignedNumNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberNanobot.copyOut()")

	if fmtSignedNum == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSignedNum' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newFmtSignedNum, err
	}

	_,
		err =
		formatterSignedNumberMolecule{}.ptr().
			testValidityOfSignedNumValDto(
				fmtSignedNum,
				ePrefix.XCtx("Testing validity of 'fmtSignedNum'"))

	if err != nil {
		return newFmtSignedNum, err
	}

	newFmtSignedNum.numStrFormatterType =
		fmtSignedNum.numStrFormatterType

	newFmtSignedNum.positiveValueFmt =
		fmtSignedNum.positiveValueFmt

	newFmtSignedNum.negativeValueFmt =
		fmtSignedNum.negativeValueFmt

	newFmtSignedNum.turnOnIntegerDigitsSeparation =
		fmtSignedNum.turnOnIntegerDigitsSeparation

	err = newFmtSignedNum.numericSeparators.CopyIn(
		&fmtSignedNum.numericSeparators,
		ePrefix.XCtx("fmtSignedNum->newFmtSignedNum"))

	if err != nil {
		return newFmtSignedNum, err
	}

	err =
		newFmtSignedNum.numFieldLenDto.CopyIn(
			&fmtSignedNum.numFieldLenDto,
			ePrefix.XCtx("fmtSignedNum->newFmtSignedNum"))

	newFmtSignedNum.lock = new(sync.Mutex)

	return newFmtSignedNum, err
}

// ptr - Returns a pointer to a new instance of
// formatterSignedNumberNanobot.
//
func (fmtSignedNumNanobot formatterSignedNumberNanobot) ptr() *formatterSignedNumberNanobot {

	if fmtSignedNumNanobot.lock == nil {
		fmtSignedNumNanobot.lock = new(sync.Mutex)
	}

	fmtSignedNumNanobot.lock.Lock()

	defer fmtSignedNumNanobot.lock.Unlock()

	newFmtSignedNumNanobot :=
		new(formatterSignedNumberNanobot)

	newFmtSignedNumNanobot.lock = new(sync.Mutex)

	return newFmtSignedNumNanobot
}
