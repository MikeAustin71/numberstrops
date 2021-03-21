package numberstr

import (
	"fmt"
	"sync"
)

type formatterCountryMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCountryMechanics.
//
func (fmtCountryMech formatterCountryMechanics) ptr() *formatterCountryMechanics {

	if fmtCountryMech.lock == nil {
		fmtCountryMech.lock = new(sync.Mutex)
	}

	fmtCountryMech.lock.Lock()

	defer fmtCountryMech.lock.Unlock()

	newFormatterCountryMechanics := new(formatterCountryMechanics)

	newFormatterCountryMechanics.lock = new(sync.Mutex)

	return newFormatterCountryMechanics
}

// setWithComponents - Transfers new data to an instance of
// FormatterCountry. After completion, all the data fields within
// input parameter 'formatterCountry' will be overwritten.
//
func (fmtCountryMech *formatterCountryMechanics) setWithComponents(
	formatterCountry *FormatterCountry,
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCultureName string,
	abbreviatedCountryName string,
	alternateCountryNames []string,
	countryCodeTwoChar string,
	countryCodeThreeChar string,
	countryCodeNumber string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCountryMech.lock == nil {
		fmtCountryMech.lock = new(sync.Mutex)
	}

	fmtCountryMech.lock.Lock()

	defer fmtCountryMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("formatterCountryMechanics.setWithComponents()")

	if formatterCountry == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCountry' is invalid!\n"+
			"'formatterCountry' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if len(countryCultureName) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'countryCultureName' is invalid!\n"+
			"'countryCultureName' is a an empty string!\n",
			ePrefix.String())

		return err
	}

	newFormatterCountry := FormatterCountry{}

	newFormatterCountry.numStrFmtType =
		NumStrFormatTypeCode(0).CountryCulture()

	newFormatterCountry.idNo = idNo

	newFormatterCountry.idString = idString

	newFormatterCountry.description = description

	newFormatterCountry.tag = tag

	newFormatterCountry.countryCultureName = countryCultureName

	newFormatterCountry.abbreviatedCountryName = abbreviatedCountryName

	newFormatterCountry.countryCodeTwoChar = countryCodeTwoChar

	newFormatterCountry.countryCodeThreeChar = countryCodeThreeChar

	newFormatterCountry.countryCodeNumber = countryCodeNumber

	lenAltNames := len(alternateCountryNames)

	if alternateCountryNames == nil {
		newFormatterCountry.alternateCountryNames = nil
	} else {
		newFormatterCountry.alternateCountryNames =
			make([]string, lenAltNames)
	}

	if lenAltNames != 0 {
		_ =
			copy(newFormatterCountry.alternateCountryNames,
				alternateCountryNames)
	}

	newFormatterCountry.lock = new(sync.Mutex)

	err = formatterCountryElectron{}.ptr().
		copyIn(
			formatterCountry,
			&newFormatterCountry,
			ePrefix.XCtx(
				"newFormatterCountry->formatterCountry"))

	return err
}
