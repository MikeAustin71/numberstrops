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

	if targetNumSeps.lock == nil {
		targetNumSeps.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark :=
		numericSeparatorsQuark{}

	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumericSeparators(
			inComingNumSeps,
			ePrefix.XCtx("Testing validity of 'inComingNumSeps'"))

	if err != nil {
		return err
	}

	if inComingNumSeps.decimalSeparators == nil {
		inComingNumSeps.decimalSeparators =
			make([]rune, 0, 2)
	}

	lenDecSep := len(inComingNumSeps.decimalSeparators)

	targetNumSeps.decimalSeparators =
		make([]rune, lenDecSep, lenDecSep+5)

	for i := 0; i < lenDecSep; i++ {

		targetNumSeps.decimalSeparators[i] =
			inComingNumSeps.decimalSeparators[i]
	}

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
	}

	ePrefix.SetEPref("numericSeparatorsElectron.copyOut()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNumSeps, err
	}

	numSepsDtoQuark :=
		numericSeparatorsQuark{}

	_,
		err =
		numSepsDtoQuark.testValidityOfNumericSeparators(
			numSeps,
			ePrefix.XCtx(
				"Testing validity of 'numSeps'"))

	if err != nil {
		return newNumSeps, err
	}

	newNumSeps.lock = new(sync.Mutex)

	if numSeps.decimalSeparators == nil {
		numSeps.decimalSeparators =
			make([]rune, 0, 2)
	}

	lenDecSeps := len(numSeps.decimalSeparators)

	newNumSeps.decimalSeparators =
		make([]rune, lenDecSeps, lenDecSeps+5)

	for i := 0; i < lenDecSeps; i++ {
		newNumSeps.decimalSeparators[i] =
			numSeps.decimalSeparators[i]
	}

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
