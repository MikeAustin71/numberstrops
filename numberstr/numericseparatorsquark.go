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

// equal - Receives two NumericSeparators objects and proceeds to
// determine whether all data elements in the first object are
// equal to corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numSepsOne          *NumericSeparators
//     - A pointer to the first NumericSeparators object. This
//       method will compare all data elements in this object to
//       corresponding data elements in the second
//       NumericSeparators object in order determine equivalency.
//
//
//  numSepsTwo          *NumericSeparators
//     - A pointer to the second NumericSeparators object. This
//       method will compare all data elements in the first
//       NumericSeparators object to corresponding data elements in
//       this second NumericSeparators object in order determine
//       equivalency.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  equal             bool
//     - If all the data elements in 'numSepsOne' are equal to
//       all the corresponding data elements in 'numSepsTwo',
//       this return parameter will be set to 'true'. If all the
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'numSepsOne' are equal to
//       all the corresponding data elements in 'numSepsTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are not equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (numSepsQuark *numericSeparatorsQuark) equal(
	numSepsOne *NumericSeparators,
	numSepsTwo *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
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
		"numericSeparatorsQuark.equal()")

	if numSepsOne == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSepsOne' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if numSepsTwo == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSepsTwo' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if numSepsOne.decimalSeparators == nil &&
		numSepsTwo.decimalSeparators != nil {
		err = fmt.Errorf("%v\n"+
			"numSepsOne.decimalSeparators == nil\n"+
			"numSepsTwo.decimalSeparators != nil\n"+
			"numSepsTwo.decimalSeparators='%v'",
			ePrefix.String(),
			string(numSepsTwo.decimalSeparators))

		return isEqual, err
	}

	if numSepsOne.decimalSeparators != nil &&
		numSepsTwo.decimalSeparators == nil {
		err = fmt.Errorf("%v\n"+
			"numSepsOne.decimalSeparators != nil\n"+
			"numSepsTwo.decimalSeparators == nil\n"+
			"numSepsOne.decimalSeparators='%v'",
			ePrefix.String(),
			string(numSepsOne.decimalSeparators))

		return isEqual, err
	}

	lenDecSep := len(numSepsOne.decimalSeparators)

	if lenDecSep !=
		len(numSepsTwo.decimalSeparators) {

		err = fmt.Errorf("%v\n"+
			"numSepsOne.decimalSeparators !="+
			"numSepsTwo.decimalSeparators\n"+
			"numSepsOne.decimalSeparators='%v'\n"+
			"numSepsTwo.decimalSeparators='%v'\n",
			ePrefix.String(),
			string(numSepsOne.decimalSeparators),
			string(numSepsTwo.decimalSeparators))

		return isEqual, err
	}

	for i := 0; i < lenDecSep; i++ {
		if numSepsOne.decimalSeparators[i] !=
			numSepsTwo.decimalSeparators[i] {
			err = fmt.Errorf("%v\n"+
				"numSepsOne.decimalSeparators !="+
				"numSepsTwo.decimalSeparators\n"+
				"numSepsOne.decimalSeparators='%v'\n"+
				"numSepsTwo.decimalSeparators='%v'\n"+
				"Reference index '%v'\n",
				ePrefix.String(),
				string(numSepsOne.decimalSeparators),
				string(numSepsTwo.decimalSeparators),
				i)

			return isEqual, err
		}
	}

	_,
		err = numSepsOne.integerSeparatorsDto.Equal(
		numSepsTwo.integerSeparatorsDto,
		ePrefix.XCtx(
			"numSepsOne.integerSeparatorsDto vs\n"+
				"numSepsTwo.integerSeparatorsDto\n"))

	if err != nil {
		return isEqual, err
	}

	isEqual = true

	return isEqual, err
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
