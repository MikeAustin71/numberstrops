package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDtoQuark struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of numStrFmtSpecDtoQuark.
//
func (nStrFmtSpecDtoQuark numStrFmtSpecDtoQuark) ptr() *numStrFmtSpecDtoQuark {

	if nStrFmtSpecDtoQuark.lock == nil {
		nStrFmtSpecDtoQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoQuark.lock.Lock()

	defer nStrFmtSpecDtoQuark.lock.Unlock()

	newNumStrFmtSpecDtoQuark := new(numStrFmtSpecDtoQuark)

	newNumStrFmtSpecDtoQuark.lock = new(sync.Mutex)

	return newNumStrFmtSpecDtoQuark
}

// testValidityOfNumStrFmtSpecDto - Receives an instance of
// NumStrFmtSpecDto and proceeds to test the validity of the
// member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (nStrFmtSpecDtoQuark *numStrFmtSpecDtoQuark) testValidityOfNumStrFmtSpecDto(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecDtoQuark.lock == nil {
		nStrFmtSpecDtoQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoQuark.lock.Lock()

	defer nStrFmtSpecDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto()")

	isValid = false

	if nStrFmtSpecDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDto' is nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrFmtSpecDto.lock == nil {
		nStrFmtSpecDto.lock = new(sync.Mutex)
	}

	err = nStrFmtSpecDto.countryCulture.IsValidInstanceError(
		ePrefix.XCtx(
			"Testing validity of " +
				"nStrFmtSpecDto.countryCulture"))

	if err != nil {
		return isValid, err
	}

	err = nStrFmtSpecDto.absoluteValue.IsValidInstanceError(
		ePrefix.XCtx(
			"Testing validity of nStrFmtSpecDto" +
				".absoluteValue "))

	if err != nil {
		return isValid, err
	}

	err = nStrFmtSpecDto.currencyValue.IsValidInstanceError(
		ePrefix.XCtx("Testing validity of nStrFmtSpecDto.currencyValue"))

	if err != nil {
		return isValid, err
	}

	err = nStrFmtSpecDto.signedNumValue.IsValidInstanceError(
		ePrefix.XCtx("Testing validity of nStrFmtSpecDto.signedNumValue"))

	if err != nil {
		return isValid, err
	}

	err = nStrFmtSpecDto.sciNotation.IsValidInstanceError(
		ePrefix.XCtx("Testing validity of nStrFmtSpecDto.sciNotation"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
