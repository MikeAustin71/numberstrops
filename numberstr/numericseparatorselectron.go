package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNumSepsDto' to input parameter 'targetNumSepsDto'.
//
// If 'inComingNumSepsDto' is judged to be invalid, this method
// will return an error.
//
// This method will OVERWRITE the data fields of the
// 'targetNumSepsDto' instance.
//
func (numSepsElectron *numericSeparatorsElectron) copyIn(
	targetNumSeps *NumericSeparators,
	inComingNumSeps *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsElectron.lock == nil {
		numSepsElectron.lock = new(sync.Mutex)
	}

	numSepsElectron.lock.Lock()

	defer numSepsElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsElectron.copyIn()")

	if targetNumSeps == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'targetNumSeps' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetNumSeps.lock == nil {
		targetNumSeps.lock = new(sync.Mutex)
	}

	if inComingNumSeps == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'inComingNumSeps' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if inComingNumSeps.lock == nil {
		inComingNumSeps.lock = new(sync.Mutex)
	}

	_,
		err =
		numericSeparatorsQuark{}.ptr().
			testValidityOfNumericSeparators(
				inComingNumSeps,
				ePrefix.XCtx("Testing validity of 'inComingNumSeps'"))

	if err != nil {
		return err
	}

	lenDecSep := len(inComingNumSeps.decimalSeparators)

	if lenDecSep == 0 {
		err = fmt.Errorf("%v"+
			"Error: 'inComingNumSeps.decimalSeparators' is invalid!\n"+
			"'inComingNumSeps.decimalSeparators' is a zero length rune array.\n",
			ePrefix.XCtxEmpty().String())
		return err
	}

	targetNumSeps.decimalSeparators =
		make([]rune, lenDecSep)

	copy(targetNumSeps.decimalSeparators,
		inComingNumSeps.decimalSeparators)

	err = targetNumSeps.integerSeparatorsDto.CopyIn(
		&inComingNumSeps.integerSeparatorsDto,
		ePrefix.XCtx(
			"inComingNumSeps.integerSeparatorsDto->"+
				"targetNumSeps.integerSeparatorsDto"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'numSeps' styled as a new instance of NumericSeparators.
//
// If 'numSeps' is judged to be invalid, this method will return an
// error.
//
func (numSepsElectron *numericSeparatorsElectron) copyOut(
	numSeps *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	newNumSeps NumericSeparators,
	err error) {

	if numSepsElectron.lock == nil {
		numSepsElectron.lock = new(sync.Mutex)
	}

	numSepsElectron.lock.Lock()

	defer numSepsElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsElectron.copyOut()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNumSeps, err
	}

	_,
		err =
		numericSeparatorsQuark{}.ptr().
			testValidityOfNumericSeparators(
				numSeps,
				ePrefix.XCtx(
					"Testing validity of 'numSeps'"))

	if err != nil {
		return newNumSeps, err
	}

	newNumSeps.lock = new(sync.Mutex)

	lenDecSeps := len(numSeps.decimalSeparators)

	if lenDecSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'numSeps.decimalSeparators' is invalid!\n"+
			"'numSeps.decimalSeparators' is a zero length rune array.\n",
			ePrefix.XCtxEmpty())
	}

	newNumSeps.decimalSeparators =
		make([]rune, lenDecSeps)

	copy(newNumSeps.decimalSeparators,
		numSeps.decimalSeparators)

	err =
		newNumSeps.integerSeparatorsDto.
			CopyIn(
				&numSeps.integerSeparatorsDto,
				ePrefix.XCtx(
					"numSeps.integerSeparatorsDto->"+
						"newNumSeps.integerSeparatorsDto"))

	return newNumSeps, err
}

// ptr - Returns a pointer to a new instance of
// numericSeparatorsElectron.
//
func (numSepsElectron numericSeparatorsElectron) ptr() *numericSeparatorsElectron {

	if numSepsElectron.lock == nil {
		numSepsElectron.lock = new(sync.Mutex)
	}

	numSepsElectron.lock.Lock()

	defer numSepsElectron.lock.Unlock()

	newNumSepsElectron := new(numericSeparatorsElectron)

	newNumSepsElectron.lock = new(sync.Mutex)

	return newNumSepsElectron
}
