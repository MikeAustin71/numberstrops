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
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecCntryQuark.lock == nil {
		nStrFmtSpecCntryQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryQuark.lock.Lock()

	defer nStrFmtSpecCntryQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecCountryDtoQuark.testValidityOfCountryDto()")

	isValid = false

	if nStrFmtSpecCntryDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCntryDto' is invalid!\n"+
			"'nStrFmtSpecCntryDto' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = false

	if len(nStrFmtSpecCntryDto.countryCultureName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: countryCultureName is Empty!\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true
	return isValid, err
}
