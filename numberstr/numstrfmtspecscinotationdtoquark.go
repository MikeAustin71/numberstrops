package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSciNotationDtoQuark struct {
	lock *sync.Mutex
}

// testValidityOfNumStrFmtSpecSciNotationDto - Receives an instance
// of NumStrFmtSpecSciNotationDto and proceeds to test the validity
// of the member data fields. The NumStrFmtSpecSciNotationDto
// object encapsulates scientific notation formatting for number
// string text displays.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
// If nStrFmtSpecSicNotDto.mantissaLength is greater than 1000,
// this method will return an error.
//
//
func (nStrFmtSpecSciNotQuark *numStrFmtSpecSciNotationDtoQuark) testValidityOfNumStrFmtSpecSciNotationDto(
	nStrFmtSpecSicNotDto *NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecSciNotQuark.lock == nil {
		nStrFmtSpecSciNotQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotQuark.lock.Lock()

	defer nStrFmtSpecSciNotQuark.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecSciNotationDtoQuark.testValidityOfNumStrFmtSpecSciNotationDto()")

	isValid = false

	if nStrFmtSpecSicNotDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSicNotDto' is invalid!\n"+
			"'nStrFmtSpecSicNotDto' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrFmtSpecSicNotDto.lock == nil {
		nStrFmtSpecSicNotDto.lock = new(sync.Mutex)
	}

	if nStrFmtSpecSicNotDto.mantissaLength > 1000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'nStrFmtSpecSicNotDto.mantissaLength' is invalid!\n"+
			"nStrFmtSpecSicNotDto.mantissaLength is greater than 1000.\n"+
			"nStrFmtSpecSicNotDto.mantissaLength== '%v'\n",
			ePrefix.String(),
			nStrFmtSpecSicNotDto.mantissaLength)

		return isValid, err
	}

	if nStrFmtSpecSicNotDto.mantissaLength == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: 'nStrFmtSpecSicNotDto.mantissaLength' is invalid!\n" +
			"nStrFmtSpecSicNotDto.mantissaLength is equal to zero.\n" +
			ePrefix.String())

		return isValid, err
	}

	if nStrFmtSpecSicNotDto.exponentChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'nStrFmtSpecSicNotDto.exponentChar' is invalid!\n"+
			"nStrFmtSpecSicNotDto.exponentChar== '%v'\n",
			ePrefix.String(),
			nStrFmtSpecSicNotDto.exponentChar)

		return isValid, err
	}

	err = nStrFmtSpecSicNotDto.numFieldLenDto.IsValidInstanceError(
		ePrefix.XCtx(
			"\nTesting Validity of nStrFmtSpecSicNotDto.numFieldLenDto"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
