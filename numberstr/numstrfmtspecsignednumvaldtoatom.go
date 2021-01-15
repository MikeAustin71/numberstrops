package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSignedNumValAtom struct {
	lock *sync.Mutex
}

// testValidityOfSignedNumValDto - Receives an instance of
// NumStrFmtSpecSignedNumValueDto and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (nStrFmtSpecSignedNumValAtom *numStrFmtSpecSignedNumValAtom) testValidityOfSignedNumValDto(
	nStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtSpecSignedNumValAtom.lock == nil {
		nStrFmtSpecSignedNumValAtom.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValAtom.lock.Lock()

	defer nStrFmtSpecSignedNumValAtom.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecSignedNumValAtom.testValidityOfAbsoluteValDto() "

	isValid = false

	if nStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return isValid, err
	}

	nStrSignedNumElectron :=
		numStrSignedNumValElectron{}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValPositiveValueFormat(
		nStrFmtSpecSignedNumValDto,
		ePrefix+
			"\nValidating nStrFmtSpecSignedNumValDto Positive Value Format\n ")

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrSignedNumElectron.testSignedNumValNegativeValueFormat(
		nStrFmtSpecSignedNumValDto,
		ePrefix+
			"\nValidating nStrFmtSpecSignedNumValDto Negative Value Format\n ")

	if err != nil {
		return isValid, err
	}

	err =
		nStrFmtSpecSignedNumValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix +
				"\nValidating nStrFmtSpecSignedNumValDto Number Separators\n ")

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid, err
}
