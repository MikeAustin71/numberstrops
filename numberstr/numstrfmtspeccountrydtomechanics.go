package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCountryDtoMechanics struct {
	lock *sync.Mutex
}

// setCountryDto - Transfers new data to an instance of NumStrFmtSpecCountryDto.
// After completion, all the data fields within input parameter 'nStrFmtSpecCntryDto'
// will be overwritten.
//
func (nStrFmtSpecCntryMech *numStrFmtSpecCountryDtoMechanics) setCountryDto(
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

	if nStrFmtSpecCntryMech.lock == nil {
		nStrFmtSpecCntryMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryMech.lock.Lock()

	defer nStrFmtSpecCntryMech.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecCountryDtoMechanics.setCountryDto() "

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is invalid!\n"+
			"'nStrFmtSpecCntryDto' is a 'nil' pointer\n",
			ePrefix)

		return err
	}

	if len(countryCultureName) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'countryCultureName' is invalid!\n"+
			"'countryCultureName' is a an empty string!\n",
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

	nStrFmtSpecCntryDto.alternateCountryNames =
		make([]string, lenAltNames, 10)

	if lenAltNames != 0 {

		_ =
			copy(nStrFmtSpecCntryDto.alternateCountryNames,
				alternateCountryNames)

	}

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	return err
}
