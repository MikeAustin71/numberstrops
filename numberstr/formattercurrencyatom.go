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
func (fmtCurrAtom formatterCurrencyAtom) ptr() *formatterCurrencyAtom {

	if fmtCurrAtom.lock == nil {
		fmtCurrAtom.lock = new(sync.Mutex)
	}

	fmtCurrAtom.lock.Lock()

	defer fmtCurrAtom.lock.Unlock()

	currencyAtom := new(formatterCurrencyAtom)

	currencyAtom.lock = new(sync.Mutex)

	return currencyAtom
}

// testCurrencyPositiveValueFormat - Inspects the Positive Value
// Number Format string for a Signed Number Value Formatter and returns
// an error if the format string is invalid.
//
//
func (fmtCurrAtom *formatterCurrencyAtom) testCurrencyPositiveValueFormat(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtCurrAtom.lock == nil {
		fmtCurrAtom.lock = new(sync.Mutex)
	}

	fmtCurrAtom.lock.Lock()

	defer fmtCurrAtom.lock.Unlock()

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
func (fmtCurrAtom *formatterCurrencyAtom) testCurrencyNegativeValueFormat(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtCurrAtom.lock == nil {
		fmtCurrAtom.lock = new(sync.Mutex)
	}

	fmtCurrAtom.lock.Lock()

	defer fmtCurrAtom.lock.Unlock()

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
