package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCountryDtoQuark struct {
	lock *sync.Mutex
}

// testValidityOfCountryDto - Receives an instance of
// NumStrFmtSpecCountryDto and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (nStrFmtSpecCntryQuark *numStrFmtSpecCountryDtoQuark) testValidityOfCountryDto(
	nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtSpecCntryQuark.lock == nil {
		nStrFmtSpecCntryQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryQuark.lock.Lock()

	defer nStrFmtSpecCntryQuark.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecCountryDtoQuark.testValidityOfCountryDto() "

	isValid = false

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is invalid!\n"+
			"'nStrFmtSpecCntryDto' is a 'nil' pointer\n",
			ePrefix)

		return isValid, err
	}

	isValid = false

	if len(nStrFmtSpecCntryDto.countryCultureName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: countryCultureName is Empty!\n",
			ePrefix)

		return isValid, err
	}

	isValid = true
	return isValid, err
}

// setCountryDto - Transfers new data to an instance of NumStrFmtSpecCountryDto.
// After completion, all the data fields within input parameter 'nStrFmtSpecCntryDto'
// will be overwritten.
//
func (nStrFmtSpecCntryQuark *numStrFmtSpecCountryDtoQuark) setCountryDto(
	nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto,
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
	ePrefix string) (
	err error) {

	if nStrFmtSpecCntryQuark.lock == nil {
		nStrFmtSpecCntryQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryQuark.lock.Lock()

	defer nStrFmtSpecCntryQuark.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecCountryDtoQuark.setCountryDto() "

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is invalid!\n"+
			"'nStrFmtSpecCntryDto' is a 'nil' pointer\n",
			ePrefix)

		return err
	}

	nStrFmtSpecCntryDto.idNo = idNo

	nStrFmtSpecCntryDto.idString = idString

	nStrFmtSpecCntryDto.description = description

	nStrFmtSpecCntryDto.tag = tag

	nStrFmtSpecCntryDto.countryCultureName = countryCultureName

	nStrFmtSpecCntryDto.abbreviatedCountryName = abbreviatedCountryName

	nStrFmtSpecCntryDto.countryCodeTwoChar = countryCodeTwoChar

	nStrFmtSpecCntryDto.countryCodeThreeChar = countryCodeThreeChar

	nStrFmtSpecCntryDto.countryCodeNumber = countryCodeNumber

	lenAltNames := len(alternateCountryNames)

	if lenAltNames == 0 {

		nStrFmtSpecCntryDto.alternateCountryNames =
			make([]string, 0, 10)

	} else {

		nStrFmtSpecCntryDto.alternateCountryNames =
			make([]string, lenAltNames, 10)

		for i := 0; i < lenAltNames; i++ {
			nStrFmtSpecCntryDto.alternateCountryNames[i] =
				alternateCountryNames[i]
		}

	}

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	return err
}
