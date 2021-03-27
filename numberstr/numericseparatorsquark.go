package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a NumericSeparators instance to their initial
// 'zero' values.
//
func (numSepsQuark *numericSeparatorsQuark) empty(
	numericSeparators *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsQuark.lock == nil {
		numSepsQuark.lock = new(sync.Mutex)
	}

	numSepsQuark.lock.Lock()

	defer numSepsQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsQuark." +
			"testValidityOfNumericSeparators()")

	if numericSeparators == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numericSeparators' is invalid!\n"+
			"'numericSeparators' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	numericSeparators.decimalSeparators = nil

	numericSeparators.integerSeparatorsDto.Empty()

	return err
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

	lenDecSep := len(numSep1.decimalSeparators)

	if lenDecSep !=
		len(numSep2.decimalSeparators) {

		return false
	}

	for i := 0; i < lenDecSep; i++ {
		if numSep1.decimalSeparators[i] !=
			numSep2.decimalSeparators[i] {
			return false
		}
	}

	if !numSep1.integerSeparatorsDto.Equal(
		numSep2.integerSeparatorsDto) {
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

	lenDecSeps := len(numericSeparators.decimalSeparators)

	if lenDecSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparators' is invalid!\n"+
			"'decimalSeparators' is a zero length rune array.\n",
			ePrefix.String())

		return isValid, err
	}

	err =
		numericSeparators.integerSeparatorsDto.IsValidInstanceError(
			ePrefix.XCtx(
				"numericSeparators.integerSeparatorsDto"))

	if err != nil {
		return isValid, err
	}

	isValid = true
	return isValid, err
}
