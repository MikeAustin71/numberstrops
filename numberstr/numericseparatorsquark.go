package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsQuark struct {
	lock *sync.Mutex
}

// numStrSepDtosAreEqual - Returns 'true' if the two NumericSeparators
// objects submitted as input parameters have equal values.
//
func (numSepsQuark *numericSeparatorsQuark) numStrSepDtosAreEqual(
	numSep1 *NumericSeparators,
	numSep2 *NumericSeparators) bool {

	if numSepsQuark.lock == nil {
		numSepsQuark.lock = new(sync.Mutex)
	}

	numSepsQuark.lock.Lock()

	defer numSepsQuark.lock.Unlock()

	if numSep1 == nil ||
		numSep2 == nil {
		return false
	}

	if numSep1.decimalSeparator !=
		numSep2.decimalSeparator {

		return false
	}

	if numSep1.integerSeparators == nil {
		numSep1.integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	if numSep2.integerSeparators == nil {
		numSep2.integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenNumSep1IntSeps := len(numSep1.integerSeparators)

	if len(numSep2.integerSeparators) !=
		lenNumSep1IntSeps {
		return false
	}

	for i := 0; i < lenNumSep1IntSeps; i++ {

		if !numSep1.integerSeparators[i].Equal(
			numSep2.integerSeparators[i]) {

			return false

		}
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// numericSeparatorsQuark.
//
func (numSepsQuark numericSeparatorsQuark) ptr() *numericSeparatorsQuark {

	if numSepsQuark.lock == nil {
		numSepsQuark.lock = new(sync.Mutex)
	}

	numSepsQuark.lock.Lock()

	defer numSepsQuark.lock.Unlock()

	newQuark := new(numericSeparatorsQuark)

	newQuark.lock = new(sync.Mutex)

	return newQuark
}

// testValidityOfNumericSeparators - Receives an instance of
// NumericSeparators and proceeds to test the validity of the
// member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (numSepsQuark *numericSeparatorsQuark) testValidityOfNumericSeparators(
	numericSeparators *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if numSepsQuark.lock == nil {
		numSepsQuark.lock = new(sync.Mutex)
	}

	numSepsQuark.lock.Lock()

	defer numSepsQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsQuark." +
			"testValidityOfNumericSeparators()")

	isValid = false

	if numericSeparators == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numericSeparators' is invalid!\n"+
			"'numericSeparators' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if numericSeparators.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' is empty and has a zero value.\n",
			ePrefix.String())

		return isValid, err
	}

	if numericSeparators.integerSeparators == nil {
		numericSeparators.integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntDigitsSeps := len(numericSeparators.integerSeparators)

	if lenIntDigitsSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member variable 'integerSeparators' is invalid!\n"+
			"'integerSeparators' is a ZERO length array.\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenIntDigitsSeps; i++ {

		err = numericSeparators.integerSeparators[i].IsValidInstanceError(
			ePrefix.XCtx(fmt.Sprintf(
				"Checking numericSeparators.integerSeparators[%v]",
				i)))

		if err != nil {
			return isValid, err
		}

	}

	isValid = true
	return isValid, err
}
