package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalQuark struct {
	lock *sync.Mutex
}

func (fmtOctalQuark *formatterOctalQuark) empty(
	formatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalQuark.lock == nil {
		fmtOctalQuark.lock = new(sync.Mutex)
	}

	fmtOctalQuark.lock.Lock()

	defer fmtOctalQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterOctalQuark." +
			"empty()")

	if formatterOctal == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	formatterOctal.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	formatterOctal.leftPrefix = ""

	formatterOctal.rightSuffix = ""

	formatterOctal.turnOnIntegerDigitsSeparation = false

	formatterOctal.integerSeparators.Empty()

	formatterOctal.numFieldDto.Empty()

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterOctalQuark.
//
func (fmtOctalQuark formatterOctalQuark) ptr() *formatterOctalQuark {

	if fmtOctalQuark.lock == nil {
		fmtOctalQuark.lock = new(sync.Mutex)
	}

	fmtOctalQuark.lock.Lock()

	defer fmtOctalQuark.lock.Unlock()

	newFmtOctalQuark :=
		new(formatterOctalQuark)

	newFmtOctalQuark.lock = new(sync.Mutex)

	return newFmtOctalQuark
}
