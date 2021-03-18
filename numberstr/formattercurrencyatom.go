package numberstr

import (
	"fmt"
	"sync"
)

type formatterCurrencyAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCurrencyAtom.
func (fmtCurrencyAtom formatterCurrencyAtom) ptr() *formatterCurrencyAtom {

	if fmtCurrencyAtom.lock == nil {
		fmtCurrencyAtom.lock = new(sync.Mutex)
	}

	fmtCurrencyAtom.lock.Lock()

	defer fmtCurrencyAtom.lock.Unlock()

	currencyAtom := new(formatterCurrencyAtom)

	currencyAtom.lock = new(sync.Mutex)

	return currencyAtom
}

// testCurrencyPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (fmtCurrencyAtom *formatterCurrencyAtom) testCurrencyPositiveValueFormat(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtCurrencyAtom.lock == nil {
		fmtCurrencyAtom.lock = new(sync.Mutex)
	}

	fmtCurrencyAtom.lock.Lock()

	defer fmtCurrencyAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("formatterCurrencyAtom.testCurrencyPositiveValueFormat()")

	isValid = false

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	nStrCurrencyElectron :=
		formatterCurrencyElectron{}

	isValid,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		formatterCurrency.positiveValueFmt,
		ePrefix.XCtx("Testing validity of 'formatterCurrency.positiveValueFmt'"))

	isValid = true
	return isValid, err
}

// testCurrencyNegativeValueFormat - Inspects the Negative Value
// Number Format string for a Currency Value Formatter and returns
// an error if the format string is invalid.
//
//
func (fmtCurrencyAtom *formatterCurrencyAtom) testCurrencyNegativeValueFormat(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtCurrencyAtom.lock == nil {
		fmtCurrencyAtom.lock = new(sync.Mutex)
	}

	fmtCurrencyAtom.lock.Lock()

	defer fmtCurrencyAtom.lock.Unlock()

	ePrefix.SetEPref("formatterCurrencyAtom.testCurrencyNegativeValueFormat()")

	isValid = false

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	nStrCurrencyElectron :=
		formatterCurrencyElectron{}

	isValid,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		formatterCurrency.negativeValueFmt,
		ePrefix.XCtx("Testing validity of 'formatterCurrency.negativeValueFmt'"))

	isValid = true
	return isValid, err
}
