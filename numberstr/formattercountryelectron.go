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
	} else {
		ePrefix = ePrefix.CopyPtr()
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
	} else {
		ePrefix = ePrefix.CopyPtr()
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

// equal - Receives two FormatterCountry objects and proceeds to
// determine whether all data elements in the first object are
// equal to all corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtCountryOne       *FormatterCountry
//     - A pointer to the first FormatterCountry object. This
//       method will compare all data elements in this object to
//       all corresponding data elements in the second
//       FormatterCountry object to determine if they are
//       equivalent.
//
//
//  fmtCountryTwo       *FormatterCountry
//     - A pointer to the second FormatterCurrency object. This
//       method will compare all data elements in the first
//       FormatterCountry object to all corresponding data
//       elements in this second FormatterCountry object to
//       determine if they are equivalent.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isEqual             bool
//     - If all the data elements in 'fmtCountryOne' are equal to
//       all the corresponding data elements in 'fmtCountryTwo',
//       this return parameter will be set to 'true'. If the data
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'fmtCountryOne' are equal to
//       all the corresponding data elements in 'fmtCountryTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtCountryElectron *formatterCountryElectron) equal(
	fmtCountryOne *FormatterCountry,
	fmtCountryTwo *FormatterCountry,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
	err error) {

	if fmtCountryElectron.lock == nil {
		fmtCountryElectron.lock = new(sync.Mutex)
	}

	fmtCountryElectron.lock.Lock()

	defer fmtCountryElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCountryElectron." +
			"equal()")

	isEqual = false

	if fmtCountryOne == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCountryOne' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtCountryTwo == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCountryTwo' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if fmtCountryOne.numStrFmtType !=
		fmtCountryTwo.numStrFmtType {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.numStrFmtType!=fmtCountryTwo.numStrFmtType\n"+
			"fmtCountryOne.numStrFmtType='%v'\n"+
			"fmtCountryTwo.numStrFmtType='%v'\n",
			ePrefix.String(),
			fmtCountryOne.numStrFmtType.String(),
			fmtCountryTwo.numStrFmtType.String())

		return isEqual, err
	}

	if fmtCountryOne.idNo !=
		fmtCountryTwo.idNo {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.idNo!=fmtCountryTwo.idNo\n"+
			"fmtCountryOne.idNo='%v'\n"+
			"fmtCountryTwo.idNo='%v'\n",
			ePrefix.String(),
			fmtCountryOne.idNo,
			fmtCountryTwo.idNo)

		return isEqual, err
	}

	if fmtCountryOne.idString !=
		fmtCountryTwo.idString {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.idString!=fmtCountryTwo.idString\n"+
			"fmtCountryOne.idString='%v'\n"+
			"fmtCountryTwo.idString='%v'\n",
			ePrefix.String(),
			fmtCountryOne.idString,
			fmtCountryTwo.idString)

		return isEqual, err
	}

	if fmtCountryOne.description !=
		fmtCountryTwo.description {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.description!=fmtCountryTwo.description\n"+
			"fmtCountryOne.description='%v'\n"+
			"fmtCountryTwo.description='%v'\n",
			ePrefix.String(),
			fmtCountryOne.description,
			fmtCountryTwo.description)

		return isEqual, err
	}

	if fmtCountryOne.tag !=
		fmtCountryTwo.tag {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.tag!=fmtCountryTwo.tag\n"+
			"fmtCountryOne.tag='%v'\n"+
			"fmtCountryTwo.tag='%v'\n",
			ePrefix.String(),
			fmtCountryOne.tag,
			fmtCountryTwo.tag)

		return isEqual, err
	}

	if fmtCountryOne.countryCultureName !=
		fmtCountryTwo.countryCultureName {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.countryCultureName!=fmtCountryTwo.countryCultureName\n"+
			"fmtCountryOne.countryCultureName='%v'\n"+
			"fmtCountryTwo.countryCultureName='%v'\n",
			ePrefix.String(),
			fmtCountryOne.countryCultureName,
			fmtCountryTwo.countryCultureName)

		return isEqual, err
	}

	if fmtCountryOne.abbreviatedCountryName !=
		fmtCountryTwo.abbreviatedCountryName {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.abbreviatedCountryName!=fmtCountryTwo.abbreviatedCountryName\n"+
			"fmtCountryOne.abbreviatedCountryName='%v'\n"+
			"fmtCountryTwo.abbreviatedCountryName='%v'\n",
			ePrefix.String(),
			fmtCountryOne.abbreviatedCountryName,
			fmtCountryTwo.abbreviatedCountryName)

		return isEqual, err
	}

	lenArray := len(fmtCountryOne.alternateCountryNames)

	if lenArray != len(fmtCountryTwo.alternateCountryNames) {

		err = fmt.Errorf("%v\n"+
			"Length of 'alternateCountryNames' is NOT equal!\n"+
			"Length fmtCountryOne.alternateCountryNames='%v'\n"+
			"Length fmtCountryTwo.alternateCountryNames='%v'\n",
			ePrefix.String(),
			lenArray,
			len(fmtCountryTwo.alternateCountryNames))

		return isEqual, err
	}

	for i := 0; i < lenArray; i++ {
		if fmtCountryOne.alternateCountryNames[i] !=
			fmtCountryTwo.alternateCountryNames[i] {

			err = fmt.Errorf("%v\n"+
				"fmtCountryOne.alternateCountryNames !=\n"+
				" fmtCountryTwo.alternateCountryNames at\n"+
				"index='%v'\n"+
				"fmtCountryOne.alternateCountryNames='%v'\n"+
				"fmtCountryTwo.alternateCountryNames='%v'\n",
				ePrefix.String(),
				i,
				fmtCountryOne.alternateCountryNames[i],
				fmtCountryTwo.alternateCountryNames[i])

			return isEqual, err
		}
	}

	if fmtCountryOne.countryCodeTwoChar !=
		fmtCountryTwo.countryCodeTwoChar {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.countryCodeTwoChar!=\n"+
			"  fmtCountryTwo.countryCodeTwoChar\n"+
			"fmtCountryOne.countryCodeTwoChar='%v'\n"+
			"fmtCountryTwo.countryCodeTwoChar='%v'\n",
			ePrefix.String(),
			fmtCountryOne.countryCodeTwoChar,
			fmtCountryTwo.countryCodeTwoChar)

		return isEqual, err
	}

	if fmtCountryOne.countryCodeThreeChar !=
		fmtCountryTwo.countryCodeThreeChar {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.countryCodeThreeChar!=\n"+
			"  fmtCountryTwo.countryCodeThreeChar\n"+
			"fmtCountryOne.countryCodeThreeChar='%v'\n"+
			"fmtCountryTwo.countryCodeThreeChar='%v'\n",
			ePrefix.String(),
			fmtCountryOne.countryCodeThreeChar,
			fmtCountryTwo.countryCodeThreeChar)

		return isEqual, err
	}

	if fmtCountryOne.countryCodeNumber !=
		fmtCountryTwo.countryCodeNumber {

		err = fmt.Errorf("%v\n"+
			"fmtCountryOne.countryCodeNumber!=\n"+
			"  fmtCountryTwo.countryCodeNumber\n"+
			"fmtCountryOne.countryCodeNumber='%v'\n"+
			"fmtCountryTwo.countryCodeNumber='%v'\n",
			ePrefix.String(),
			fmtCountryOne.countryCodeNumber,
			fmtCountryTwo.countryCodeNumber)

		return isEqual, err
	}

	isEqual = true

	return isEqual, err
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
