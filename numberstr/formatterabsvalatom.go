package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueAtom.
//
func (fmtAbsValAtom formatterAbsoluteValueAtom) ptr() *formatterAbsoluteValueAtom {

	if fmtAbsValAtom.lock == nil {
		fmtAbsValAtom.lock = new(sync.Mutex)
	}

	fmtAbsValAtom.lock.Lock()

	defer fmtAbsValAtom.lock.Unlock()

	newAbsValAtom :=
		new(formatterAbsoluteValueAtom)

	newAbsValAtom.lock = new(sync.Mutex)

	return newAbsValAtom
}

// testAbsoluteValueFormat - Inspects the format string for an
// Absolute Value number string and returns an error if the format
// string is invalid.
//
func (fmtAbsValAtom *formatterAbsoluteValueAtom) testAbsoluteValueFormat(
	fmtAbsoluteVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtAbsValAtom.lock == nil {
		fmtAbsValAtom.lock = new(sync.Mutex)
	}

	fmtAbsValAtom.lock.Lock()

	defer fmtAbsValAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(

		"formatterAbsoluteValueAtom." +
			"testAbsoluteValueFormat()")

	isValid = false

	if fmtAbsoluteVal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	if len(fmtAbsoluteVal.absoluteValFmt) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal.absoluteValFmt' "+
			"is invalid!\n"+
			"The Absolute Value Dto Format is a zero length string.\n"+
			"len(fmtAbsoluteVal.absoluteValFmt) == 0\n",
			ePrefix.String())

		return isValid, err
	}

	isValid,
		err = formatterAbsoluteValueElectron{}.ptr().
		testAbsoluteValueFormatStr(
			fmtAbsoluteVal.absoluteValFmt,
			ePrefix.XCtx(
				"Testing validity of 'fmtAbsoluteVal.absoluteValFmt'"))

	return isValid, err
}
