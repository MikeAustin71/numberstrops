package numberstr

import (
	"fmt"
	"sync"
)

type formatterHexadecimalElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFormatterHex' to input parameter
// 'targetFormatterHex'.
//
// Be advised - All data fields in 'targetFormatterHex'
// will be overwritten.
//
// If input parameter 'incomingFormatterHex' is judged
// to be invalid, this method will return an error.
//
func (fmtHexadecimalElectron *formatterHexadecimalElectron) copyIn(
	targetFormatterHex *FormatterHexadecimal,
	incomingFormatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalElectron.lock == nil {
		fmtHexadecimalElectron.lock = new(sync.Mutex)
	}

	fmtHexadecimalElectron.lock.Lock()

	defer fmtHexadecimalElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalElectron.copyIn()")

	if targetFormatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFormatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if incomingFormatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFormatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		formatterHexadecimalQuark{}.ptr().
			testValidityOfFormatterHexadecimal(
				incomingFormatterHex,
				ePrefix.XCtx(
					"Testing validity of 'incomingFormatterHex'"))

	if err != nil {
		return err
	}

	targetFormatterHex.numStrFmtType =
		incomingFormatterHex.numStrFmtType

	targetFormatterHex.useUpperCaseLetters =
		incomingFormatterHex.useUpperCaseLetters

	targetFormatterHex.leftPrefix =
		incomingFormatterHex.leftPrefix

	targetFormatterHex.turnOnIntegerDigitsSeparation =
		incomingFormatterHex.turnOnIntegerDigitsSeparation

	err =
		targetFormatterHex.integerSeparators.CopyIn(
			&incomingFormatterHex.integerSeparators,
			ePrefix.XCtx(
				"incomingFormatterHex.integerSeparators -> "+
					"targetFormatterHex.integerSeparators"))

	if err != nil {
		return err
	}

	err = targetFormatterHex.numFieldDto.CopyIn(
		&incomingFormatterHex.numFieldDto,
		ePrefix.XCtx(
			"incomingFormatterHex.numFieldDto -> "+
				"targetFormatterHex.numFieldDto"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'formatterHex' styled as a new instance of
// FormatterHexadecimal.
//
// If input parameter 'formatterHex' is judged to be invalid, this
// method will return an error.
//
func (fmtHexadecimalElectron *formatterHexadecimalElectron) copyOut(
	formatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	newFormatterHex FormatterHexadecimal,
	err error) {

	if fmtHexadecimalElectron.lock == nil {
		fmtHexadecimalElectron.lock = new(sync.Mutex)
	}

	fmtHexadecimalElectron.lock.Lock()

	defer fmtHexadecimalElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalElectron.copyOut()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newFormatterHex, err
	}

	_,
		err =
		formatterHexadecimalQuark{}.ptr().
			testValidityOfFormatterHexadecimal(
				formatterHex,
				ePrefix.XCtx(
					"Testing validity of 'formatterHex'"))

	if err != nil {
		return newFormatterHex, err
	}

	newFormatterHex.lock = new(sync.Mutex)

	newFormatterHex.numStrFmtType =
		formatterHex.numStrFmtType

	newFormatterHex.useUpperCaseLetters =
		formatterHex.useUpperCaseLetters

	newFormatterHex.leftPrefix =
		formatterHex.leftPrefix

	newFormatterHex.turnOnIntegerDigitsSeparation =
		formatterHex.turnOnIntegerDigitsSeparation

	err =
		newFormatterHex.integerSeparators.CopyIn(
			&formatterHex.integerSeparators,
			ePrefix.XCtx(
				"formatterHex.integerSeparators -> "+
					"newFormatterHex.integerSeparators"))

	if err != nil {
		return newFormatterHex, err
	}

	err = newFormatterHex.numFieldDto.CopyIn(
		&formatterHex.numFieldDto,
		ePrefix.XCtx(
			"formatterHex.numFieldDto -> "+
				"newFormatterHex.numFieldDto"))

	return newFormatterHex, err
}

func (fmtHexadecimalElectron *formatterHexadecimalElectron) equal(
	fmtHexadecimalOne *FormatterHexadecimal,
	fmtHexadecimalTwo *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtHexadecimalElectron.lock == nil {
		fmtHexadecimalElectron.lock = new(sync.Mutex)
	}

	fmtHexadecimalElectron.lock.Lock()

	defer fmtHexadecimalElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalElectron." +
			"equal()")

	isEqual = false

	if fmtHexadecimalOne == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtHexadecimalOne' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtHexadecimalTwo == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtHexadecimalTwo' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtHexadecimalOne.numStrFmtType !=
		fmtHexadecimalTwo.numStrFmtType {

		err = fmt.Errorf("%v\n"+
			"fmtHexadecimalOne.numStrFmtType!=\n"+
			"  fmtHexadecimalTwo.numStrFmtType\n"+
			"fmtHexadecimalOne.numStrFmtType='%v'\n"+
			"fmtHexadecimalTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtHexadecimalOne.numStrFmtType.String(),
			fmtHexadecimalTwo.numStrFmtType.String())

		return isEqual, err
	}

	if fmtHexadecimalOne.useUpperCaseLetters !=
		fmtHexadecimalTwo.useUpperCaseLetters {

		err = fmt.Errorf("%v\n"+
			"fmtHexadecimalOne.useUpperCaseLetters!=\n"+
			"  fmtHexadecimalTwo.useUpperCaseLetters\n"+
			"fmtHexadecimalOne.useUpperCaseLetters='%v'\n"+
			"fmtHexadecimalTwo.useUpperCaseLetters='%v'\n",
			ePrefix.String(),
			fmtHexadecimalOne.useUpperCaseLetters,
			fmtHexadecimalTwo.useUpperCaseLetters)

		return isEqual, err
	}

	if fmtHexadecimalOne.leftPrefix !=
		fmtHexadecimalTwo.leftPrefix {

		err = fmt.Errorf("%v\n"+
			"fmtHexadecimalOne.leftPrefix!=\n"+
			"  fmtHexadecimalTwo.leftPrefix\n"+
			"fmtHexadecimalOne.leftPrefix='%v'\n"+
			"fmtHexadecimalTwo.leftPrefix='%v'\n",
			ePrefix.String(),
			fmtHexadecimalOne.leftPrefix,
			fmtHexadecimalTwo.leftPrefix)

		return isEqual, err
	}

	if fmtHexadecimalOne.turnOnIntegerDigitsSeparation !=
		fmtHexadecimalTwo.turnOnIntegerDigitsSeparation {

		err = fmt.Errorf("%v\n"+
			"fmtHexadecimalOne.turnOnIntegerDigitsSeparation!=\n"+
			"  fmtHexadecimalTwo.turnOnIntegerDigitsSeparation\n"+
			"fmtHexadecimalOne.turnOnIntegerDigitsSeparation='%v'\n"+
			"fmtHexadecimalTwo.turnOnIntegerDigitsSeparation='%v'\n",
			ePrefix.String(),
			fmtHexadecimalOne.turnOnIntegerDigitsSeparation,
			fmtHexadecimalTwo.turnOnIntegerDigitsSeparation)

		return isEqual, err
	}

	_,
		err = fmtHexadecimalOne.integerSeparators.Equal(
		fmtHexadecimalTwo.integerSeparators,
		ePrefix.XCtx(
			"fmtHexadecimalOne vs fmtHexadecimalTwo"))

	if err != nil {
		return isEqual, err
	}

	_,
		err = fmtHexadecimalOne.numFieldDto.Equal(
		fmtHexadecimalTwo.numFieldDto,
		ePrefix.XCtx(
			"fmtHexadecimalOne vs fmtHexadecimalTwo"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
}

// ptr - Returns a pointer to a new instance of
// formatterHexadecimalElectron.
//
func (fmtHexadecimalElectron formatterHexadecimalElectron) ptr() *formatterHexadecimalElectron {

	if fmtHexadecimalElectron.lock == nil {
		fmtHexadecimalElectron.lock = new(sync.Mutex)
	}

	fmtHexadecimalElectron.lock.Lock()

	defer fmtHexadecimalElectron.lock.Unlock()

	newFmtHexadecimalElectron :=
		new(formatterHexadecimalElectron)

	newFmtHexadecimalElectron.lock = new(sync.Mutex)

	return newFmtHexadecimalElectron
}
