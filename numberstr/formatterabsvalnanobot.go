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

	if targetFmtAbsoluteVal.lock == nil {
		targetFmtAbsoluteVal.lock = new(sync.Mutex)
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

// equal - Receives two FormatterAbsoluteValue objects and proceeds
// to determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsoluteValOne   *FormatterAbsoluteValue
//     - A pointer to the first FormatterAbsoluteValue object. This
//       method will compare all data elements in this object to all
//       corresponding data elements in the second
//       FormatterAbsoluteValue object to determine if they are
//       equivalent.
//
//
//  fmtAbsoluteValTwo      *FormatterAbsoluteValue
//     - A pointer to the second FormatterAbsoluteValue object.
//       This method will compare all data elements in the first
//       FormatterAbsoluteValue object to all corresponding data
//       elements in this second FormatterAbsoluteValue object to
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
//     - If all the data elements in 'fmtAbsoluteValOne' are equal
//       to all the corresponding data elements in
//       'fmtAbsoluteValTwo', this return parameter will be set to
//       'true'. If the data data elements are NOT equal, this
//       return parameter will be set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtAbsoluteValOne' are equal
//       to all the corresponding data elements in
//       'fmtAbsoluteValTwo', this return parameter will be set to
//       'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtAbsValNanobot *formatterAbsoluteValueNanobot) equal(
	fmtAbsoluteValOne *FormatterAbsoluteValue,
	fmtAbsoluteValTwo *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
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
		"formatterAbsoluteValueNanobot." +
			"equal()")

	isEqual = false

	if fmtAbsoluteValOne == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteValOne' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtAbsoluteValTwo == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteValTwo' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtAbsoluteValOne.numStrFmtType !=
		fmtAbsoluteValTwo.numStrFmtType {
		err = fmt.Errorf("%v\n"+
			"fmtAbsoluteValOne.numStrFmtType!=fmtAbsoluteValTwo.numStrFmtType\n"+
			"fmtAbsoluteValOne.numStrFmtType='%v'\n"+
			"fmtAbsoluteValTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtAbsoluteValOne.numStrFmtType.String(),
			fmtAbsoluteValTwo.numStrFmtType.String())
		return isEqual, err
	}

	if fmtAbsoluteValOne.absoluteValFmt !=
		fmtAbsoluteValTwo.absoluteValFmt {
		err = fmt.Errorf("%v\n"+
			"fmtAbsoluteValOne.absoluteValFmt!=fmtAbsoluteValTwo.absoluteValFmt\n"+
			"fmtAbsoluteValOne.absoluteValFmt='%v'\n"+
			"fmtAbsoluteValTwo.absoluteValFmt='%v'\n",
			ePrefix.String(),
			fmtAbsoluteValOne.absoluteValFmt,
			fmtAbsoluteValTwo.absoluteValFmt)
		return isEqual, err
	}

	if fmtAbsoluteValOne.turnOnIntegerDigitsSeparation !=
		fmtAbsoluteValTwo.turnOnIntegerDigitsSeparation {
		err = fmt.Errorf("%v\n"+
			"fmtAbsoluteValOne.turnOnIntegerDigitsSeparation!=\n"+
			"  fmtAbsoluteValTwo.turnOnIntegerDigitsSeparation\n"+
			"fmtAbsoluteValOne.turnOnIntegerDigitsSeparation='%v'\n"+
			"fmtAbsoluteValTwo.turnOnIntegerDigitsSeparation='%v'\n",
			ePrefix.String(),
			fmtAbsoluteValOne.turnOnIntegerDigitsSeparation,
			fmtAbsoluteValTwo.turnOnIntegerDigitsSeparation)
		return isEqual, err
	}

	_,
		err = fmtAbsoluteValOne.numericSeparators.Equal(
		fmtAbsoluteValTwo.numericSeparators,
		ePrefix.XCtx(
			"fmtAbsoluteValOne vs fmtAbsoluteValTwo"))

	if err != nil {
		return isEqual, err
	}

	_,
		err = fmtAbsoluteValOne.numFieldDto.Equal(
		fmtAbsoluteValTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtAbsoluteValOne vs fmtAbsoluteValTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
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
