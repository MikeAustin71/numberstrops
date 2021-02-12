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
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCntryMech.lock == nil {
		nStrFmtSpecCntryMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryMech.lock.Lock()

	defer nStrFmtSpecCntryMech.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecCountryDtoMechanics.setCountryDto()")

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is invalid!\n"+
			"'nStrFmtSpecCntryDto' is a 'nil' pointer\n",
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

	testStrFmtSpecCntryDto := NumStrFmtSpecCountryDto{}

	testStrFmtSpecCntryDto.idNo = idNo

	testStrFmtSpecCntryDto.idString = idString

	testStrFmtSpecCntryDto.description = description

	testStrFmtSpecCntryDto.tag = tag

	testStrFmtSpecCntryDto.countryCultureName = countryCultureName

	testStrFmtSpecCntryDto.abbreviatedCountryName = abbreviatedCountryName

	testStrFmtSpecCntryDto.countryCodeTwoChar = countryCodeTwoChar

	testStrFmtSpecCntryDto.countryCodeThreeChar = countryCodeThreeChar

	testStrFmtSpecCntryDto.countryCodeNumber = countryCodeNumber

	lenAltNames := len(alternateCountryNames)

	testStrFmtSpecCntryDto.alternateCountryNames =
		make([]string, lenAltNames, 10)

	if lenAltNames != 0 {

		_ =
			copy(testStrFmtSpecCntryDto.alternateCountryNames,
				alternateCountryNames)

	}

	testStrFmtSpecCntryDto.lock = new(sync.Mutex)

	nStrFmtSpecCntryElectron :=
		numStrFmtSpecCountryDtoElectron{}

	err = nStrFmtSpecCntryElectron.copyIn(
		nStrFmtSpecCntryDto,
		&testStrFmtSpecCntryDto,
		ePrefix.XCtx(
			"testStrFmtSpecCntryDto->nStrFmtSpecCntryDto"))

	return err
}
