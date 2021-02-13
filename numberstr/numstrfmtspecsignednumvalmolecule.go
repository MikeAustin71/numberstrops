package numberstr

import (
	"fmt"
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
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecSignedNumValMolecule.lock == nil {
		nStrFmtSpecSignedNumValMolecule.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValMolecule.lock.Lock()

	defer nStrFmtSpecSignedNumValMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto()")

	isValid = false

	if nStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

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
		ePrefix.XCtx("Validating nStrFmtSpecSignedNumValDto Positive Value Format"))

	if err != nil {
		return isValid, err
	}

	isValid,
		err = nStrSignedNumAtom.testSignedNumValNegativeValueFormat(
		nStrFmtSpecSignedNumValDto,
		ePrefix.XCtx("Validating nStrFmtSpecSignedNumValDto Negative Value Format"))

	if err != nil {
		return isValid, err
	}

	err =
		nStrFmtSpecSignedNumValDto.numberSeparatorsDto.IsValidInstanceError(
			ePrefix.XCtx("Validating 'nStrFmtSpecSignedNumValDto' Number Separators"))

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid, err
}
