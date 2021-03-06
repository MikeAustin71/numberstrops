package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingNStrIntSeparator' to input parameter
// 'targetNStrIntSeparator'.
//
// Be advised - All data fields in 'targetNStrIntSeparator'
// will be overwritten.
//
// If input parameter 'incomingNStrIntSeparator' is judged
// to be invalid, this method will return an error.
//
func (nStrIntSepMolecule *numStrIntSeparatorMolecule) copyIn(
	targetNStrIntSeparator *NumStrIntSeparator,
	incomingNStrIntSeparator *NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrIntSeparatorMolecule.copyIn()")

	if targetNStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingNStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		numStrIntSeparatorQuark{}.Ptr().
			testValidityOfNumStrIntSeparator(
				incomingNStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'incomingNStrIntSeparator'."))

	if err != nil {
		return err
	}

	targetNStrIntSeparator.intSeparatorChar =
		incomingNStrIntSeparator.intSeparatorChar

	targetNStrIntSeparator.intSeparatorGrouping =
		incomingNStrIntSeparator.intSeparatorGrouping

	targetNStrIntSeparator.intSeparatorRepetitions =
		incomingNStrIntSeparator.intSeparatorRepetitions

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'numStrIntSeparator' styled as a new instance
// of NumStrIntSeparator.
//
// If input parameter 'numStrIntSeparator' is judged to be
// invalid, this method will return an error.
//
func (nStrIntSepMolecule *numStrIntSeparatorMolecule) copyOut(
	numStrIntSeparator *NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	newNumSrIntSeparator NumStrIntSeparator,
	err error) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrIntSeparatorMolecule.copyOut()")

	if numStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNumSrIntSeparator, err
	}

	_,
		err =
		numStrIntSeparatorQuark{}.Ptr().
			testValidityOfNumStrIntSeparator(
				numStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'numStrIntSeparator'."))

	if err != nil {
		return newNumSrIntSeparator, err
	}

	newNumSrIntSeparator.intSeparatorChar =
		numStrIntSeparator.intSeparatorChar

	newNumSrIntSeparator.intSeparatorGrouping =
		numStrIntSeparator.intSeparatorGrouping

	newNumSrIntSeparator.intSeparatorRepetitions =
		numStrIntSeparator.intSeparatorRepetitions

	newNumSrIntSeparator.lock = new(sync.Mutex)

	return newNumSrIntSeparator, err
}

// equal - Receives two instances of NumStrIntSeparator and
// proceeds to compare the internal data values. If all the
// data values in 'nStrIntSep1' and 'nStrIntSep2' are equal, this
// method returns 'true'.
//
// If any of the data values in 'nStrIntSep1' and 'nStrIntSep2' are
// NOT equal, this method returns 'false'.
//
func (nStrIntSepMolecule *numStrIntSeparatorMolecule) equal(
	nStrIntSep1 *NumStrIntSeparator,
	nStrIntSep2 *NumStrIntSeparator) (
	areEqual bool) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	areEqual = true

	if nStrIntSep1.intSeparatorChar !=
		nStrIntSep2.intSeparatorChar {
		areEqual = false
	}

	if nStrIntSep1.intSeparatorGrouping !=
		nStrIntSep2.intSeparatorGrouping {
		areEqual = false
	}

	if nStrIntSep1.intSeparatorRepetitions !=
		nStrIntSep2.intSeparatorRepetitions {
		areEqual = false
	}

	return areEqual
}

// Ptr - Returns a pointer to a new instance of
// numStrIntSeparatorMolecule.
func (nStrIntSepMolecule numStrIntSeparatorMolecule) Ptr() *numStrIntSeparatorMolecule {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	newIntSepMolecule := new(numStrIntSeparatorMolecule)

	newIntSepMolecule.lock = new(sync.Mutex)

	return newIntSepMolecule
}
