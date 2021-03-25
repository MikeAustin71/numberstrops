package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterOctalMechanics.
//
func (fmtOctalMech formatterOctalMechanics) ptr() *formatterOctalMechanics {

	if fmtOctalMech.lock == nil {
		fmtOctalMech.lock = new(sync.Mutex)
	}

	fmtOctalMech.lock.Lock()

	defer fmtOctalMech.lock.Unlock()

	newFmtOctalMech :=
		new(formatterOctalMechanics)

	newFmtOctalMech.lock = new(sync.Mutex)

	return newFmtOctalMech
}

func (fmtOctalMech *formatterOctalMechanics) setFmtOctalWithComponents(
	formatterOctal *FormatterOctal,
	leftPrefix string,
	rightSuffix string,
	turnOnIntegerDigitsSeparation bool,
	integerSeparators NumStrIntSeparatorsDto,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalMech.lock == nil {
		fmtOctalMech.lock = new(sync.Mutex)
	}

	fmtOctalMech.lock.Lock()

	defer fmtOctalMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterOctalMechanics." +
			"setFmtOctalWithComponents()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterOctal.lock == nil {
		formatterOctal.lock = new(sync.Mutex)
	}

	err = integerSeparators.IsValidInstanceError(
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}

	err = numFieldDto.IsValidInstanceError(
		ePrefix.XCtx(
			"numFieldDto"))

	if err != nil {
		return err
	}

	formatterOctal.numStrFmtType =
		NumStrFormatTypeCode(0).Octal()

	formatterOctal.leftPrefix =
		leftPrefix

	formatterOctal.rightSuffix =
		rightSuffix

	formatterOctal.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		formatterOctal.integerSeparators.CopyIn(
			&integerSeparators,
			ePrefix.XCtx(
				"integerSeparators->"+
					"formatterOctal"))

	if err != nil {
		return err
	}

	err =
		formatterOctal.numFieldDto.CopyIn(
			&numFieldDto,
			ePrefix.XCtx(
				"numFieldDto->"+
					"formatterOctal"))

	return err
}
