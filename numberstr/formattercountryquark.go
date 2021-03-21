package numberstr

import (
	"fmt"
	"sync"
)

type formatterCountryQuark struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCountryQuark.
//
func (fmtCountryQuark formatterCountryQuark) ptr() *formatterCountryQuark {

	if fmtCountryQuark.lock == nil {
		fmtCountryQuark.lock = new(sync.Mutex)
	}

	fmtCountryQuark.lock.Lock()

	defer fmtCountryQuark.lock.Unlock()

	newFormatterCountryQuark := new(formatterCountryQuark)

	newFormatterCountryQuark.lock = new(sync.Mutex)

	return newFormatterCountryQuark
}

// testValidityOfFormatterCountry - Receives an instance of
// FormatterCountry and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (fmtCountryQuark *formatterCountryQuark) testValidityOfFormatterCountry(
	formatterCountry *FormatterCountry,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtCountryQuark.lock == nil {
		fmtCountryQuark.lock = new(sync.Mutex)
	}

	fmtCountryQuark.lock.Lock()

	defer fmtCountryQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCountryQuark." +
			"testValidityOfFormatterCountry()")

	isValid = false

	if formatterCountry == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCountry' is invalid!\n"+
			"'formatterCountry' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = false

	if len(formatterCountry.countryCultureName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: countryCultureName is Empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterCountry.numStrFmtType !=
		NumStrFormatTypeCode(0).CountryCulture() {

		err = fmt.Errorf("%v\n"+
			"Error: Number String Format Type Code is invalid!\n"+
			"'formatterCountry.numStrFmtType' IS NOT EQUAL TO \n"+
			"NumStrFormatTypeCode(0).CountryCulture()\n"+
			"'formatterCountry.numStrFmtType' integer value='%v'\n",
			ePrefix.String(),
			formatterCountry.numStrFmtType.XValueInt())

		return isValid, err
	}

	isValid = true
	return isValid, err
}
