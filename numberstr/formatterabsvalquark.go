package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets data values for all member variables
// in a FormatterAbsoluteValue instance to their initial 'zero'
// values.
//
func (fmtAbsValQuark *formatterAbsoluteValueQuark) empty(
	fmtAbsoluteVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValQuark.lock == nil {
		fmtAbsValQuark.lock = new(sync.Mutex)
	}

	fmtAbsValQuark.lock.Lock()

	defer fmtAbsValQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueQuark." +
			"empty()")

	if fmtAbsoluteVal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteVal' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	fmtAbsoluteVal.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	fmtAbsoluteVal.absoluteValFmt = ""

	fmtAbsoluteVal.turnOnIntegerDigitsSeparation = false

	fmtAbsoluteVal.numericSeparators.Empty()

	fmtAbsoluteVal.numFieldLenDto = NumberFieldDto{}

	return err
}

// getValidAbsoluteValFmtChars- Returns an array of runes which
// constitute the only valid characters allowed in Absolute Value
// format strings.
//
func (fmtAbsValQuark *formatterAbsoluteValueQuark) getValidAbsoluteValFmtChars() []rune {

	if fmtAbsValQuark.lock == nil {
		fmtAbsValQuark.lock = new(sync.Mutex)
	}

	fmtAbsValQuark.lock.Lock()

	defer fmtAbsValQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")

}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueQuark.
//
func (fmtAbsValQuark formatterAbsoluteValueQuark) ptr() *formatterAbsoluteValueQuark {

	if fmtAbsValQuark.lock == nil {
		fmtAbsValQuark.lock = new(sync.Mutex)
	}

	fmtAbsValQuark.lock.Lock()

	defer fmtAbsValQuark.lock.Unlock()

	newAbsValQuark :=
		new(formatterAbsoluteValueQuark)

	newAbsValQuark.lock = new(sync.Mutex)

	return newAbsValQuark
}
