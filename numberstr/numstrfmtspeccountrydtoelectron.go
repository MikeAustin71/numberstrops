package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCountryDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecCntryDto' to input parameter
// 'targetNStrFmtSpecCntryDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecCntryDto'
// will be overwritten.
//
func (nStrFmtSpecCntryElectron *numStrFmtSpecCountryDtoElectron) copyIn(
	targetNStrFmtSpecCntryDto *NumStrFmtSpecCountryDto,
	inComingNStrFmtSpecCntryDto *NumStrFmtSpecCountryDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCntryElectron.lock == nil {
		nStrFmtSpecCntryElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryElectron.lock.Lock()

	defer nStrFmtSpecCntryElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecCountryDtoElectron.copyIn()")

	if targetNStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecCntryDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if inComingNStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecCntryDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	nStrFmtSpecCntryQuark := numStrFmtSpecCountryDtoQuark{}

	_,
		err = nStrFmtSpecCntryQuark.testValidityOfCountryDto(
		inComingNStrFmtSpecCntryDto,
		ePrefix.XCtx(
			"Testing validity of 'inComingNStrFmtSpecCntryDto'"))

	if err != nil {
		return err
	}

	targetNStrFmtSpecCntryDto.idNo =
		inComingNStrFmtSpecCntryDto.idNo

	targetNStrFmtSpecCntryDto.idString =
		inComingNStrFmtSpecCntryDto.idString

	targetNStrFmtSpecCntryDto.description =
		inComingNStrFmtSpecCntryDto.description

	targetNStrFmtSpecCntryDto.tag =
		inComingNStrFmtSpecCntryDto.tag

	targetNStrFmtSpecCntryDto.countryCultureName =
		inComingNStrFmtSpecCntryDto.countryCultureName

	targetNStrFmtSpecCntryDto.abbreviatedCountryName =
		inComingNStrFmtSpecCntryDto.abbreviatedCountryName

	lenAltNames :=
		len(inComingNStrFmtSpecCntryDto.alternateCountryNames)

	targetNStrFmtSpecCntryDto.alternateCountryNames =
		make([]string, lenAltNames, lenAltNames+10)

	_ = copy(targetNStrFmtSpecCntryDto.alternateCountryNames,
		inComingNStrFmtSpecCntryDto.alternateCountryNames)

	targetNStrFmtSpecCntryDto.countryCodeTwoChar =
		inComingNStrFmtSpecCntryDto.countryCodeTwoChar

	targetNStrFmtSpecCntryDto.countryCodeThreeChar =
		inComingNStrFmtSpecCntryDto.countryCodeThreeChar

	targetNStrFmtSpecCntryDto.countryCodeNumber =
		inComingNStrFmtSpecCntryDto.countryCodeNumber

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecCntryDto' styled as a new instance
// of NumStrFmtSpecCountryDto.
//
func (nStrFmtSpecCntryElectron *numStrFmtSpecCountryDtoElectron) copyOut(
	nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto,
	ePrefix *ErrPrefixDto) (
	newFmtSpecCntryDto NumStrFmtSpecCountryDto,
	err error) {

	if nStrFmtSpecCntryElectron.lock == nil {
		nStrFmtSpecCntryElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryElectron.lock.Lock()

	defer nStrFmtSpecCntryElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecCountryDtoElectron.copyOut()")

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return newFmtSpecCntryDto, err
	}

	nStrFmtSpecCntryQuark := numStrFmtSpecCountryDtoQuark{}

	_,
		err = nStrFmtSpecCntryQuark.testValidityOfCountryDto(
		nStrFmtSpecCntryDto,
		ePrefix.XCtx("Testing validity of 'nStrFmtSpecCntryDto'"))

	if err != nil {
		return newFmtSpecCntryDto, err
	}

	newFmtSpecCntryDto.idNo =
		nStrFmtSpecCntryDto.idNo

	newFmtSpecCntryDto.idString =
		nStrFmtSpecCntryDto.idString

	newFmtSpecCntryDto.description =
		nStrFmtSpecCntryDto.description

	newFmtSpecCntryDto.tag =
		nStrFmtSpecCntryDto.tag

	newFmtSpecCntryDto.countryCultureName =
		nStrFmtSpecCntryDto.countryCultureName

	newFmtSpecCntryDto.abbreviatedCountryName =
		nStrFmtSpecCntryDto.abbreviatedCountryName

	lenAltNames :=
		len(nStrFmtSpecCntryDto.alternateCountryNames)

	if lenAltNames == 0 {
		newFmtSpecCntryDto.alternateCountryNames =
			make([]string, 0, 10)
	} else {
		newFmtSpecCntryDto.alternateCountryNames =
			make([]string, lenAltNames, 10)

		for i := 0; i < lenAltNames; i++ {
			newFmtSpecCntryDto.alternateCountryNames[i] =
				nStrFmtSpecCntryDto.alternateCountryNames[i]
		}
	}

	newFmtSpecCntryDto.countryCodeTwoChar =
		nStrFmtSpecCntryDto.countryCodeTwoChar

	newFmtSpecCntryDto.countryCodeThreeChar =
		nStrFmtSpecCntryDto.countryCodeThreeChar

	newFmtSpecCntryDto.countryCodeNumber =
		nStrFmtSpecCntryDto.countryCodeNumber

	newFmtSpecCntryDto.lock = new(sync.Mutex)

	return newFmtSpecCntryDto, err
}
