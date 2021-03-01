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
func (numSepsDtoQuark *numericSeparatorsQuark) numStrSepDtosAreEqual(
	numSep1 *NumericSeparators,
	numSep2 *NumericSeparators) bool {

	if numSepsDtoQuark.lock == nil {
		numSepsDtoQuark.lock = new(sync.Mutex)
	}

	numSepsDtoQuark.lock.Lock()

	defer numSepsDtoQuark.lock.Unlock()

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
func (numSepsDtoQuark numericSeparatorsQuark) ptr() *numericSeparatorsQuark {

	if numSepsDtoQuark.lock == nil {
		numSepsDtoQuark.lock = new(sync.Mutex)
	}

	numSepsDtoQuark.lock.Lock()

	defer numSepsDtoQuark.lock.Unlock()

	newQuark := new(numericSeparatorsQuark)

	newQuark.lock = new(sync.Mutex)

	return newQuark
}

// testValidityOfNumSepsDto - Receives an instance of
// NumericSeparators and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (numSepsDtoQuark *numericSeparatorsQuark) testValidityOfNumSepsDto(
	numSepsDto *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if numSepsDtoQuark.lock == nil {
		numSepsDtoQuark.lock = new(sync.Mutex)
	}

	numSepsDtoQuark.lock.Lock()

	defer numSepsDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsQuark." +
			"testValidityOfNumSepsDto()")

	isValid = false

	if numSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSepsDto' is invalid!\n"+
			"'numSepsDto' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if numSepsDto.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' is empty and has a zero value.\n",
			ePrefix.String())

		return isValid, err
	}

	if numSepsDto.integerSeparators == nil {
		numSepsDto.integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntDigitsSeps := len(numSepsDto.integerSeparators)

	if lenIntDigitsSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member variable 'integerSeparators' is invalid!\n"+
			"'integerSeparators' is a ZERO length array.\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenIntDigitsSeps; i++ {

		err = numSepsDto.integerSeparators[i].IsValidInstanceError(
			ePrefix.XCtx(fmt.Sprintf(
				"Checking numSepsDto.integerSeparators[%v]",
				i)))

		if err != nil {
			return isValid, err
		}

	}

	isValid = true
	return isValid, err
}
