package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFmtSpecSignedNumValMolecule struct {
	lock *sync.Mutex
}

// testValidityOfSignedNumValDto - Receives an instance of
// NumStrFmtSpecSignedNumValueDto and proceeds to test the validity
// of the member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (nStrFmtSpecSignedNumValMolecule *numStrFmtSpecSignedNumValMolecule) testValidityOfSignedNumValDto(
	nStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtSpecSignedNumValMolecule.lock == nil {
		nStrFmtSpecSignedNumValMolecule.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValMolecule.lock.Lock()

	defer nStrFmtSpecSignedNumValMolecule.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto() "

	isValid = false

	if nStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return isValid, err
	}

	if nStrFmtSpecSignedNumValDto.lock == nil {
		nStrFmtSpecSignedNumValDto.lock = new(sync.Mutex)
	}

	nStrSignedNumAtom :=
		numStrFmtSpecSignedNumValAtom{}

	isValid,
		err = nStrSignedNumAtom.testSignedNumValPositiveValueFormat(
		nStrFmtSpecSignedNumValDto,
		ePrefix+
			"\nValidating nStrFmtSpecSignedNumValDto Positive Value Format\n ")

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrSignedNumAtom.testSignedNumValNegativeValueFormat(
		nStrFmtSpecSignedNumValDto,
		ePrefix+
			"\nValidating nStrFmtSpecSignedNumValDto Negative Value Format\n ")

	if err != nil {
		return isValid, err
	}

	err =
		nStrFmtSpecSignedNumValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix +
				"\nValidating 'nStrFmtSpecSignedNumValDto' Number Separators\n ")

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid, err
}
