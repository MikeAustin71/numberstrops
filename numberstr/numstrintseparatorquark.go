package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorQuark struct {
	lock *sync.Mutex
}

// testValidityOfNumStrIntSeparator - Tests the validity of a
// NumStrIntSeparator instance.
//
// If the NumStrIntSeparator instance is judged invalid, this
// method will return an error and set 'isValid' to false.
//
func (nStrIntSepQuark *numStrIntSeparatorQuark) testValidityOfNumStrIntSeparator(
	nStrIntSep *NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrIntSepQuark.lock == nil {
		nStrIntSepQuark.lock = new(sync.Mutex)
	}

	nStrIntSepQuark.lock.Lock()

	defer nStrIntSepQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorQuark." +
			"testValidityOfNumStrIntSeparator()")

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return isValid, err
	}

	if nStrIntSep.intSeparatorChars == nil {
		nStrIntSep.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lIntSepChars := len(nStrIntSep.intSeparatorChars)

	if lIntSepChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a ZERO length rune array.\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrIntSep.intSeparatorGrouping > 10000000 {

		err = fmt.Errorf("%v\n"+
			"Error: 'nStrIntSep.intSeparatorGrouping' is invalid!\n"+
			"nStrIntSep.intSeparatorGrouping='%v'\n",
			ePrefix,
			nStrIntSep.intSeparatorGrouping)

		return isValid, err
	}

	if nStrIntSep.intSeparatorRepetitions > 10000000 {

		err = fmt.Errorf("%v\n"+
			"Error: 'nStrIntSep.intSeparatorRepetitions' is invalid!\n"+
			"nStrIntSep.intSeparatorRepetitions='%v'\n",
			ePrefix,
			nStrIntSep.intSeparatorRepetitions)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// Ptr - Returns a pointer to a new instance of
// numStrIntSeparatorQuark
func (nStrIntSepQuark numStrIntSeparatorQuark) Ptr() *numStrIntSeparatorQuark {

	if nStrIntSepQuark.lock == nil {
		nStrIntSepQuark.lock = new(sync.Mutex)
	}

	nStrIntSepQuark.lock.Lock()

	defer nStrIntSepQuark.lock.Unlock()

	newIntSepQuark := new(numStrIntSeparatorQuark)

	newIntSepQuark.lock = new(sync.Mutex)

	return newIntSepQuark
}
