package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCountryDtoQuark struct {
	lock *sync.Mutex
}

// hasData - Examines an instance of NumStrFmtSpecCountryDto.
// If any of the member variables have non-zero values, this
// method returns 'true'. If all member variables are set to
// their zero values, this method returns, 'false'.
//
func (nStrFmtSpecCntryQuark *numStrFmtSpecCountryDtoQuark) hasData(
	nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto,
	ePrefix string) (
	hasData bool,
	err error) {

	if nStrFmtSpecCntryQuark.lock == nil {
		nStrFmtSpecCntryQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryQuark.lock.Lock()

	defer nStrFmtSpecCntryQuark.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecCountryDtoQuark.hasData() "

	hasData = false

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is invalid!\n"+
			"'nStrFmtSpecCntryDto' is a 'nil' pointer\n",
			ePrefix)

		return hasData, err
	}

	hasData = true

	if len(nStrFmtSpecCntryDto.idString) > 0 {
		return hasData, err
	}

	if len(nStrFmtSpecCntryDto.countryCultureName) > 0 {
		return hasData, err
	}

	if len(nStrFmtSpecCntryDto.abbreviatedCountryName) > 0 {
		return hasData, err
	}

	if len(nStrFmtSpecCntryDto.alternateCountryNames) > 0 {
		return hasData, err
	}

	if len(nStrFmtSpecCntryDto.countryCodeTwoChar) > 0 {
		return hasData, err
	}

	if len(nStrFmtSpecCntryDto.countryCodeThreeChar) > 0 {
		return hasData, err
	}

	if len(nStrFmtSpecCntryDto.countryCodeNumber) > 0 {
		return hasData, err
	}

	hasData = false
	return hasData, err
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
