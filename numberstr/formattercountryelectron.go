package numberstr

import (
	"fmt"
	"sync"
)

type formatterCountryElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFormatterCountry' to input parameter
// 'targetFormatterCountry'.
//
// Be advised - All data fields in 'targetFormatterCountry'
// will be overwritten.
//
func (fmtCountryElectron *formatterCountryElectron) copyIn(
	targetFormatterCountry *FormatterCountry,
	incomingFormatterCountry *FormatterCountry,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCountryElectron.lock == nil {
		fmtCountryElectron.lock = new(sync.Mutex)
	}

	fmtCountryElectron.lock.Lock()

	defer fmtCountryElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCountryElectron.copyIn()")

	if targetFormatterCountry == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFormatterCountry' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if incomingFormatterCountry == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFormatterCountry' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err = formatterCountryQuark{}.ptr().
		testValidityOfFormatterCountry(
			incomingFormatterCountry,
			ePrefix.XCtx(
				"Testing validity of 'incomingFormatterCountry'"))

	if err != nil {
		return err
	}

	targetFormatterCountry.numStrFmtType =
		incomingFormatterCountry.numStrFmtType

	targetFormatterCountry.idNo =
		incomingFormatterCountry.idNo

	targetFormatterCountry.idString =
		incomingFormatterCountry.idString

	targetFormatterCountry.description =
		incomingFormatterCountry.description

	targetFormatterCountry.tag =
		incomingFormatterCountry.tag

	targetFormatterCountry.countryCultureName =
		incomingFormatterCountry.countryCultureName

	targetFormatterCountry.abbreviatedCountryName =
		incomingFormatterCountry.abbreviatedCountryName

	lenAltNames :=
		len(incomingFormatterCountry.alternateCountryNames)

	if incomingFormatterCountry.alternateCountryNames == nil {

		targetFormatterCountry.alternateCountryNames = nil

	} else {

		targetFormatterCountry.alternateCountryNames =
			make([]string, lenAltNames)
	}

	if lenAltNames != 0 {
		_ = copy(targetFormatterCountry.alternateCountryNames,
			incomingFormatterCountry.alternateCountryNames)
	}

	targetFormatterCountry.countryCodeTwoChar =
		incomingFormatterCountry.countryCodeTwoChar

	targetFormatterCountry.countryCodeThreeChar =
		incomingFormatterCountry.countryCodeThreeChar

	targetFormatterCountry.countryCodeNumber =
		incomingFormatterCountry.countryCodeNumber

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'formatterCountry' styled as a new instance
// of FormatterCountry.
//
func (fmtCountryElectron *formatterCountryElectron) copyOut(
	formatterCountry *FormatterCountry,
	ePrefix *ErrPrefixDto) (
	newFmtSpecCntryDto FormatterCountry,
	err error) {

	if fmtCountryElectron.lock == nil {
		fmtCountryElectron.lock = new(sync.Mutex)
	}

	fmtCountryElectron.lock.Lock()

	defer fmtCountryElectron.lock.Unlock()

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
		return newFmtSpecCntryDto, err
	}

	_,
		err = formatterCountryQuark{}.ptr().
		testValidityOfFormatterCountry(
			formatterCountry,
			ePrefix.XCtx("Testing validity of 'formatterCountry'"))

	if err != nil {
		return newFmtSpecCntryDto, err
	}

	newFmtSpecCntryDto.numStrFmtType =
		formatterCountry.numStrFmtType

	newFmtSpecCntryDto.idNo =
		formatterCountry.idNo

	newFmtSpecCntryDto.idString =
		formatterCountry.idString

	newFmtSpecCntryDto.description =
		formatterCountry.description

	newFmtSpecCntryDto.tag =
		formatterCountry.tag

	newFmtSpecCntryDto.countryCultureName =
		formatterCountry.countryCultureName

	newFmtSpecCntryDto.abbreviatedCountryName =
		formatterCountry.abbreviatedCountryName

	lenAltNames :=
		len(formatterCountry.alternateCountryNames)

	if formatterCountry.alternateCountryNames == nil {

		newFmtSpecCntryDto.alternateCountryNames = nil

	} else {

		newFmtSpecCntryDto.alternateCountryNames =
			make([]string, lenAltNames)

	}

	if lenAltNames != 0 {
		_ = copy(newFmtSpecCntryDto.alternateCountryNames,
			formatterCountry.alternateCountryNames)
	}

	newFmtSpecCntryDto.countryCodeTwoChar =
		formatterCountry.countryCodeTwoChar

	newFmtSpecCntryDto.countryCodeThreeChar =
		formatterCountry.countryCodeThreeChar

	newFmtSpecCntryDto.countryCodeNumber =
		formatterCountry.countryCodeNumber

	newFmtSpecCntryDto.lock = new(sync.Mutex)

	return newFmtSpecCntryDto, err
}

// ptr - Returns a pointer to a new instance of
// formatterCountryElectron.
//
func (fmtCountryElectron formatterCountryElectron) ptr() *formatterCountryElectron {

	if fmtCountryElectron.lock == nil {
		fmtCountryElectron.lock = new(sync.Mutex)
	}

	fmtCountryElectron.lock.Lock()

	defer fmtCountryElectron.lock.Unlock()

	newFormatterCountryElectron := new(formatterCountryElectron)

	newFormatterCountryElectron.lock = new(sync.Mutex)

	return newFormatterCountryElectron
}
