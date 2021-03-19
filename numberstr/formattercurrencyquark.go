package numberstr

import (
	"fmt"
	"sync"
)

type formatterCurrencyQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a FormatterCurrency instance to their initial
// 'zero' values.
//
func (fmtCurrQuark *formatterCurrencyQuark) empty(
	fmtCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrQuark.lock == nil {
		fmtCurrQuark.lock = new(sync.Mutex)
	}

	fmtCurrQuark.lock.Lock()

	defer fmtCurrQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyQuark.empty()")

	if fmtCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCurrency' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	fmtCurrency.numStrFormatterType =
		NumStrFormatTypeCode(0).None()

	fmtCurrency.positiveValueFmt = ""

	fmtCurrency.negativeValueFmt = ""

	fmtCurrency.decimalDigits = 0

	fmtCurrency.currencyCode = ""

	fmtCurrency.currencyCodeNo = ""

	fmtCurrency.currencyName = ""

	fmtCurrency.currencySymbols = nil

	fmtCurrency.minorCurrencyName = ""

	fmtCurrency.minorCurrencySymbols = nil

	fmtCurrency.turnOnIntegerDigitsSeparation = false

	fmtCurrency.numericSeparators.Empty()

	fmtCurrency.numFieldLenDto.Empty()

	return err
}

// getValidCurrencyPositiveValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Positive Value format
// strings for currency displays.
//
func (fmtCurrQuark *formatterCurrencyQuark) getValidCurrencyPositiveValFmtChars() []rune {

	if fmtCurrQuark.lock == nil {
		fmtCurrQuark.lock = new(sync.Mutex)
	}

	fmtCurrQuark.lock.Lock()

	defer fmtCurrQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")
}

// getValidCurrencyNegativeValFmtChars - Returns an array of runes which
// constitute the only valid characters allowed in Negative Value format
// strings for currency displays.
//
func (fmtCurrQuark *formatterCurrencyQuark) getValidCurrencyNegativeValFmtChars() []rune {

	if fmtCurrQuark.lock == nil {
		fmtCurrQuark.lock = new(sync.Mutex)
	}

	fmtCurrQuark.lock.Lock()

	defer fmtCurrQuark.lock.Unlock()

	return []rune("(-) $127.54NUMFIELD")
}

// ptr - Returns a pointer to a new instance of
// formatterCurrencyQuark.
//
func (fmtCurrQuark formatterCurrencyQuark) ptr() *formatterCurrencyQuark {

	if fmtCurrQuark.lock == nil {
		fmtCurrQuark.lock = new(sync.Mutex)
	}

	fmtCurrQuark.lock.Lock()

	defer fmtCurrQuark.lock.Unlock()

	newCurrencyQuark :=
		new(formatterCurrencyQuark)

	newCurrencyQuark.lock = new(sync.Mutex)

	return newCurrencyQuark
}
