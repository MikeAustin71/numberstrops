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

// empty - Deletes and resets the data values of all member
// variables within a FormatterCountry instance to their initial
// 'zero' values.
//
func (fmtCountryQuark *formatterCountryQuark) empty(
	formatterCountry *FormatterCountry,
	ePrefix *ErrPrefixDto) (
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
		"formatterCountryElectron.copyOut()")

	if formatterCountry == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCountry' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	formatterCountry.numStrFmtType =
		NumStrFormatTypeCode(0).None()

	formatterCountry.idNo = 0

	formatterCountry.idString = ""

	formatterCountry.description = ""

	formatterCountry.tag = ""

	formatterCountry.countryCultureName = ""

	formatterCountry.abbreviatedCountryName = ""

	formatterCountry.alternateCountryNames = nil

	formatterCountry.countryCodeTwoChar = ""

	formatterCountry.countryCodeThreeChar = ""

	formatterCountry.countryCodeNumber = ""

	return err
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
