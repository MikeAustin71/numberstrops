package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsQuark struct {
	lock *sync.Mutex
}

// numericSeparatorsAreEqual - Returns 'true' if the two NumericSeparators
// objects submitted as input parameters have equal data values.
//
func (numSepsQuark *numericSeparatorsQuark) numericSeparatorsAreEqual(
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

	if numSep1.decimalSeparator == nil {
		numSep1.decimalSeparator =
			make([]rune, 0, 2)
	}

	if numSep2.decimalSeparator == nil {
		numSep2.decimalSeparator =
			make([]rune, 0, 2)
	}

	lenDecSep := len(numSep1.decimalSeparator)

	if lenDecSep !=
		len(numSep2.decimalSeparator) {

		return false
	}

	for i := 0; i < lenDecSep; i++ {
		if numSep1.decimalSeparator[i] !=
			numSep2.decimalSeparator[i] {
			return false
		}
	}

	if !numSep1.integerSeparators.Equal(
		numSep2.integerSeparators) {
		return false
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

	if numericSeparators.decimalSeparator == nil {
		numericSeparators.decimalSeparator =
			make([]rune, 0, 2)
	}

	lenDecSeps := len(numericSeparators.decimalSeparator)

	if lenDecSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' is a zero length rune array.\n",
			ePrefix.String())

		return isValid, err
	}

	err =
		numericSeparators.integerSeparators.IsValidInstanceError(
			ePrefix.XCtx(
				"numericSeparators.integerSeparators"))

	if err != nil {
		return isValid, err
	}

	isValid = true
	return isValid, err
}
