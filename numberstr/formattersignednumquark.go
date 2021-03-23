package numberstr

import (
	"fmt"
	"sync"
)

type formatterSignedNumberQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a FormatterSignedNumber instance to their
// initial 'zero' values.
//
func (fmtSignedNumQuark *formatterSignedNumberQuark) empty(
	fmtSignedNum *FormatterSignedNumber,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSignedNumQuark.lock == nil {
		fmtSignedNumQuark.lock = new(sync.Mutex)
	}

	fmtSignedNumQuark.lock.Lock()

	defer fmtSignedNumQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSignedNumberMolecule." +
			"testValidityOfSignedNumValDto()")

	if fmtSignedNum == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSignedNum' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	fmtSignedNum.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	fmtSignedNum.positiveValueFmt = ""

	fmtSignedNum.negativeValueFmt = ""

	fmtSignedNum.turnOnIntegerDigitsSeparation = false

	fmtSignedNum.numericSeparators.Empty()

	fmtSignedNum.numFieldDto.Empty()

	return err
}

// getValidSignedNumPositiveValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Positive Value format
// strings for signed number values.
//
func (fmtSignedNumQuark *formatterSignedNumberQuark) getValidSignedNumPositiveValFmtChars() []rune {

	if fmtSignedNumQuark.lock == nil {
		fmtSignedNumQuark.lock = new(sync.Mutex)
	}

	fmtSignedNumQuark.lock.Lock()

	defer fmtSignedNumQuark.lock.Unlock()

	return []rune(" +127.54NUMFIELD")
}

// getValidSignedNumNegativeValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Negative Value format
// strings for signed number values.
//
func (fmtSignedNumQuark *formatterSignedNumberQuark) getValidSignedNumNegativeValFmtChars() []rune {

	if fmtSignedNumQuark.lock == nil {
		fmtSignedNumQuark.lock = new(sync.Mutex)
	}

	fmtSignedNumQuark.lock.Lock()

	defer fmtSignedNumQuark.lock.Unlock()

	return []rune("(-) 127.54NUMFIELD")
}

// ptr - Returns a pointer to a new instance of
// formatterSignedNumberQuark.
//
func (fmtSignedNumQuark formatterSignedNumberQuark) ptr() *formatterSignedNumberQuark {

	if fmtSignedNumQuark.lock == nil {
		fmtSignedNumQuark.lock = new(sync.Mutex)
	}

	fmtSignedNumQuark.lock.Lock()

	defer fmtSignedNumQuark.lock.Unlock()

	newFmtSignedNumQuark :=
		new(formatterSignedNumberQuark)

	newFmtSignedNumQuark.lock = new(sync.Mutex)

	return newFmtSignedNumQuark
}
