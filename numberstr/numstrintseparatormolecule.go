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
		numStrIntSeparatorQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				incomingNStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'incomingNStrIntSeparator'."))

	if err != nil {
		return err
	}

	if incomingNStrIntSeparator.intSeparatorChars == nil {
		incomingNStrIntSeparator.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lIntSepChars :=
		len(incomingNStrIntSeparator.intSeparatorChars)

	targetNStrIntSeparator.intSeparatorChars =
		make([]rune, lIntSepChars, lIntSepChars+5)

	for i := 0; i < lIntSepChars; i++ {
		targetNStrIntSeparator.intSeparatorChars[i] =
			incomingNStrIntSeparator.intSeparatorChars[i]
	}

	targetNStrIntSeparator.intSeparatorGrouping =
		incomingNStrIntSeparator.intSeparatorGrouping

	targetNStrIntSeparator.intSeparatorRepetitions =
		incomingNStrIntSeparator.intSeparatorRepetitions

	targetNStrIntSeparator.restartIntGroupingSequence =
		incomingNStrIntSeparator.restartIntGroupingSequence

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

	ePrefix.SetEPref(
		"numStrIntSeparatorMolecule.copyOut()")

	if numStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNumSrIntSeparator, err
	}

	_,
		err =
		numStrIntSeparatorQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				numStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'numStrIntSeparator'."))

	if err != nil {
		return newNumSrIntSeparator, err
	}

	if numStrIntSeparator.intSeparatorChars == nil {
		numStrIntSeparator.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lIntSepChars :=
		len(numStrIntSeparator.intSeparatorChars)

	newNumSrIntSeparator.intSeparatorChars =
		make([]rune, lIntSepChars, lIntSepChars+5)

	for i := 0; i < lIntSepChars; i++ {
		newNumSrIntSeparator.intSeparatorChars[i] =
			numStrIntSeparator.intSeparatorChars[i]
	}

	newNumSrIntSeparator.intSeparatorGrouping =
		numStrIntSeparator.intSeparatorGrouping

	newNumSrIntSeparator.intSeparatorRepetitions =
		numStrIntSeparator.intSeparatorRepetitions

	newNumSrIntSeparator.restartIntGroupingSequence =
		numStrIntSeparator.restartIntGroupingSequence

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

	if nStrIntSep1.intSeparatorChars == nil {
		nStrIntSep1.intSeparatorChars =
			make([]rune, 0, 5)
	}

	if nStrIntSep2.intSeparatorChars == nil {
		nStrIntSep2.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIntSeps1 := len(nStrIntSep1.intSeparatorChars)

	if lenIntSeps1 != len(nStrIntSep2.intSeparatorChars) {
		areEqual = false
		return areEqual
	}

	for i := 0; i < lenIntSeps1; i++ {
		if nStrIntSep1.intSeparatorChars[i] !=
			nStrIntSep2.intSeparatorChars[i] {
			areEqual = false
		}
	}

	if nStrIntSep1.intSeparatorGrouping !=
		nStrIntSep2.intSeparatorGrouping {
		areEqual = false
	}

	if nStrIntSep1.intSeparatorRepetitions !=
		nStrIntSep2.intSeparatorRepetitions {
		areEqual = false
	}

	if nStrIntSep1.restartIntGroupingSequence !=
		nStrIntSep2.restartIntGroupingSequence {
		areEqual = false
	}

	return areEqual
}

// ptr - Returns a pointer to a new instance of
// numStrIntSeparatorMolecule.
func (nStrIntSepMolecule numStrIntSeparatorMolecule) ptr() *numStrIntSeparatorMolecule {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	newIntSepMolecule := new(numStrIntSeparatorMolecule)

	newIntSepMolecule.lock = new(sync.Mutex)

	return newIntSepMolecule
}
