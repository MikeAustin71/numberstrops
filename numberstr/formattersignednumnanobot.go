package numberstr

import (
	"fmt"
	"sync"
)

type formatterSignedNumberNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFmtSignedNum' to input parameter 'targetFmtSignedNum'.
//
// Be advised - All data fields in 'targetFmtSignedNum'
// will be overwritten.
//
// If input parameter 'targetFmtSignedNum' is judged
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
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberNanobot.copyIn()")

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
				ePrefix.XCtx(
					"Testing validity of 'incomingFmtSignedNum'"))

	if err != nil {
		return err
	}

	targetFmtSignedNum.numStrFmtType =
		incomingFmtSignedNum.numStrFmtType

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
		targetFmtSignedNum.numFieldDto.CopyIn(
			&incomingFmtSignedNum.numFieldDto,
			ePrefix.XCtx("incomingFmtSignedNum->targetFmtSignedNum"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'fmtSignedNum' styled as a new instance of
// FormatterSignedNumber.
//
// If input parameter 'fmtSignedNum' is judged to be invalid, this
// method will return an error.
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

	newFmtSignedNum.numStrFmtType =
		fmtSignedNum.numStrFmtType

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
		newFmtSignedNum.numFieldDto.CopyIn(
			&fmtSignedNum.numFieldDto,
			ePrefix.XCtx("fmtSignedNum->newFmtSignedNum"))

	newFmtSignedNum.lock = new(sync.Mutex)

	return newFmtSignedNum, err
}

// equal - Receives two FormatterSignedNumber objects and proceeds
// to determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSignedNumberOne  *FormatterSignedNumber
//     - A pointer to the first FormatterSignedNumber object. This
//       method will compare all data elements in this object to
//       all corresponding data elements in the second
//       FormatterSignedNumber object to determine if they are
//       equivalent.
//
//
//  fmtSignedNumberTwo  *FormatterSignedNumber
//     - A pointer to the second FormatterSignedNumber object. This
//       method will compare all data elements in the first
//       FormatterSignedNumber object to all corresponding data
//       elements in this second FormatterSignedNumber object to
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
//     - If all the data elements in 'fmtSignedNumberOne' are equal
//       to all the corresponding data elements in
//       'fmtSignedNumberTwo', this return parameter will be set to
//       'true'. If the data elements are NOT equal, this return
//       parameter will be set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtSignedNumberOne' are equal
//       to all the corresponding data elements in
//       'fmtSignedNumberTwo', this return parameter will be set to
//       'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtSignedNumNanobot *formatterSignedNumberNanobot) equal(
	fmtSignedNumberOne *FormatterSignedNumber,
	fmtSignedNumberTwo *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtSignedNumNanobot.lock == nil {
		fmtSignedNumNanobot.lock = new(sync.Mutex)
	}

	fmtSignedNumNanobot.lock.Lock()

	defer fmtSignedNumNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberNanobot." +
			"equal()")

	isEqual = false

	if fmtSignedNumberOne == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSignedNumberOne' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtSignedNumberTwo == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSignedNumberTwo' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtSignedNumberOne.numStrFmtType !=
		fmtSignedNumberTwo.numStrFmtType {

		err = fmt.Errorf("%v\n"+
			"fmtSignedNumberOne.numStrFmtType!=\n"+
			"  fmtSignedNumberTwo.numStrFmtType\n"+
			"fmtSignedNumberOne.numStrFmtType='%v'\n"+
			"fmtSignedNumberTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtSignedNumberOne.numStrFmtType.String(),
			fmtSignedNumberTwo.numStrFmtType.String())

		return isEqual, err
	}

	if fmtSignedNumberOne.positiveValueFmt !=
		fmtSignedNumberTwo.positiveValueFmt {

		err = fmt.Errorf("%v\n"+
			"fmtSignedNumberOne.positiveValueFmt!=\n"+
			"  fmtSignedNumberTwo.positiveValueFmt\n"+
			"fmtSignedNumberOne.positiveValueFmt='%v'\n"+
			"fmtSignedNumberTwo.positiveValueFmt='%v'\n",
			ePrefix.String(),
			fmtSignedNumberOne.positiveValueFmt,
			fmtSignedNumberTwo.positiveValueFmt)

		return isEqual, err
	}

	if fmtSignedNumberOne.negativeValueFmt !=
		fmtSignedNumberTwo.negativeValueFmt {

		err = fmt.Errorf("%v\n"+
			"fmtSignedNumberOne.negativeValueFmt!=\n"+
			"  fmtSignedNumberTwo.negativeValueFmt\n"+
			"fmtSignedNumberOne.negativeValueFmt='%v'\n"+
			"fmtSignedNumberTwo.negativeValueFmt='%v'\n",
			ePrefix.String(),
			fmtSignedNumberOne.negativeValueFmt,
			fmtSignedNumberTwo.negativeValueFmt)

		return isEqual, err
	}

	if fmtSignedNumberOne.turnOnIntegerDigitsSeparation !=
		fmtSignedNumberTwo.turnOnIntegerDigitsSeparation {

		err = fmt.Errorf("%v\n"+
			"fmtSignedNumberOne.turnOnIntegerDigitsSeparation!=\n"+
			"  fmtSignedNumberTwo.turnOnIntegerDigitsSeparation\n"+
			"fmtSignedNumberOne.turnOnIntegerDigitsSeparation='%v'\n"+
			"fmtSignedNumberTwo.turnOnIntegerDigitsSeparation='%v'\n",
			ePrefix.String(),
			fmtSignedNumberOne.turnOnIntegerDigitsSeparation,
			fmtSignedNumberTwo.turnOnIntegerDigitsSeparation)

		return isEqual, err
	}

	_,
		err = fmtSignedNumberOne.numericSeparators.Equal(
		fmtSignedNumberTwo.numericSeparators,
		ePrefix.XCtx(
			"fmtSignedNumberOne vs "+
				"fmtSignedNumberTwo"))

	if err != nil {
		return isEqual, err
	}

	_,
		err = fmtSignedNumberOne.numFieldDto.Equal(
		fmtSignedNumberTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtSignedNumberOne vs "+
				"fmtSignedNumberTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
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
