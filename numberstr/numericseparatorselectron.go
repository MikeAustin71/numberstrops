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

	targetNumSeps.decimalSeparators =
		inComingNumSeps.decimalSeparators

	lenIntSeparators :=
		len(inComingNumSeps.integerSeparators)

	targetNumSeps.integerSeparators =
		make([]NumStrIntSeparator,
			lenIntSeparators,
			10)

	for i := 0; i < lenIntSeparators; i++ {

		err =
			targetNumSeps.integerSeparators[i].CopyIn(
				&inComingNumSeps.integerSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf("Copying inComingNumSeps.integerSeparators[%v]",
						i)))

		if err != nil {
			return err
		}

	}

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecDigitsSepsDto' styled as a new instance
// of NumericSeparators.
//
// If 'nStrFmtSpecDigitsSepsDto' is judged to be invalid, this
// method will return an error.
//
func (numSepsElectron *numericSeparatorsElectron) copyOut(
	numSeps *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	newNumSepsDto NumericSeparators,
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

		return newNumSepsDto, err
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
		return newNumSepsDto, err
	}

	newNumSepsDto.decimalSeparators =
		numSeps.decimalSeparators

	lenIntSeparators :=
		len(numSeps.integerSeparators)

	newNumSepsDto.integerSeparators =
		make([]NumStrIntSeparator, lenIntSeparators)

	for i := 0; i < lenIntSeparators; i++ {
		err =
			newNumSepsDto.integerSeparators[i].CopyIn(
				&numSeps.integerSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf(
						"Copying numSeps.integerSeparators[%v]",
						i)))

		if err != nil {
			return newNumSepsDto, err
		}

	}

	newNumSepsDto.lock = new(sync.Mutex)

	return newNumSepsDto, err
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
