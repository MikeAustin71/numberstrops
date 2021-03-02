package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsMechanics struct {
	lock *sync.Mutex
}

// setIntDigitsSeps - Transfers new data to an instance of NumericSeparators.
// After completion, all the data fields within input parameter 'nStrFmtSpecDigitsSepDto'
// will be overwritten.
//
func (numSepsMech *numericSeparatorsMechanics) setIntDigitsSeps(
	numSeps *NumericSeparators,
	decimalSeparator rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsMech.lock == nil {
		numSepsMech.lock = new(sync.Mutex)
	}

	numSepsMech.lock.Lock()

	defer numSepsMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numericSeparatorsMechanics.setIntDigitsSeps()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	if decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is invalid!\n"+
			"decimalSeparator == 0\n",
			ePrefix.String())

		return err
	}

	if integerSeparators == nil {
		integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntDigitSeparators :=
		len(integerSeparators)

	if lenIntDigitSeparators == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerSeparators' is invalid!\n"+
			"'integerSeparators' is a zero length or empty array.\n",
			ePrefix.String())

		return err
	}

	newNumericSeps := NumericSeparators{}

	newNumericSeps.decimalSeparator =
		decimalSeparator

	newNumericSeps.integerSeparators =
		make([]NumStrIntSeparator,
			lenIntDigitSeparators,
			10)

	for i := 0; i < lenIntDigitSeparators; i++ {
		err =
			newNumericSeps.integerSeparators[i].
				CopyIn(
					&integerSeparators[i],
					ePrefix.XCtx(
						fmt.Sprintf(
							"integerSeparators[%v]",
							i)))

		if err != nil {
			return err
		}

	}

	newNumericSeps.lock = new(sync.Mutex)

	nStrFmtSpecDigitsSepsQuark := numericSeparatorsQuark{}
	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumericSeparators(
			&newNumericSeps,
			ePrefix.XCtx("Testing validity of preliminary 'newNumericSeps'"))

	if err != nil {
		return err
	}

	nStrFmtSpecDigitsSepsElectron :=
		numericSeparatorsElectron{}

	err =
		nStrFmtSpecDigitsSepsElectron.copyIn(
			numSeps,
			&newNumericSeps,
			ePrefix.XCtx("newNumericSeps->numSeps"))

	return err
}
