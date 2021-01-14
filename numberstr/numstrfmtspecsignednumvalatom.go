package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSignedNumValAtom struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecDigitsSepsDto' to input parameter
// 'targetNStrFmtSpecDigitsSepsDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecDigitsSepsDto'
// will be overwritten.
//
func (nStrFmtSpecSignedNumValAtom *numStrFmtSpecSignedNumValAtom) copyIn(
	targetNStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	inComingNStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecSignedNumValAtom.lock == nil {
		nStrFmtSpecSignedNumValAtom.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValAtom.lock.Lock()

	defer nStrFmtSpecSignedNumValAtom.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecSignedNumValAtom.copyIn() "

	if targetNStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if inComingNStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	targetNStrFmtSpecSignedNumValDto.positiveValueFmt =
		inComingNStrFmtSpecSignedNumValDto.positiveValueFmt

	targetNStrFmtSpecSignedNumValDto.negativeValueFmt =
		inComingNStrFmtSpecSignedNumValDto.negativeValueFmt

	targetNStrFmtSpecSignedNumValDto.turnOnIntegerDigitsSeparation =
		inComingNStrFmtSpecSignedNumValDto.turnOnIntegerDigitsSeparation

	err =
		targetNStrFmtSpecSignedNumValDto.numberSeparatorsDto.CopyIn(
			&inComingNStrFmtSpecSignedNumValDto.numberSeparatorsDto,
			ePrefix+
				"\nCopying from inComingNStrFmtSpecSignedNumValDto\n"+
				"to targetNStrFmtSpecSignedNumValDto ")

	if err != nil {
		return err
	}

	targetNStrFmtSpecSignedNumValDto.numFieldLenDto.CopyIn(
		&inComingNStrFmtSpecSignedNumValDto.numFieldLenDto)

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecDigitsSepsDto' styled as a new instance
// of NumStrFmtSpecDigitsSeparatorsDto.
//
func (nStrFmtSpecSignedNumValAtom *numStrFmtSpecSignedNumValAtom) copyOut(
	nStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	newNStrFmtSpecSignedNumValDto NumStrFmtSpecSignedNumValueDto,
	err error) {

	if nStrFmtSpecSignedNumValAtom.lock == nil {
		nStrFmtSpecSignedNumValAtom.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValAtom.lock.Lock()

	defer nStrFmtSpecSignedNumValAtom.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecSignedNumValAtom.copyOut() "

	if nStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return newNStrFmtSpecSignedNumValDto, err
	}

	newNStrFmtSpecSignedNumValDto.positiveValueFmt =
		nStrFmtSpecSignedNumValDto.positiveValueFmt

	newNStrFmtSpecSignedNumValDto.negativeValueFmt =
		nStrFmtSpecSignedNumValDto.negativeValueFmt

	newNStrFmtSpecSignedNumValDto.turnOnIntegerDigitsSeparation =
		nStrFmtSpecSignedNumValDto.turnOnIntegerDigitsSeparation

	err = newNStrFmtSpecSignedNumValDto.numberSeparatorsDto.CopyIn(
		&nStrFmtSpecSignedNumValDto.numberSeparatorsDto,
		ePrefix+
			"\nnStrFmtSpecSignedNumValDto->newNStrFmtSpecSignedNumValDto ")

	if err != nil {
		return newNStrFmtSpecSignedNumValDto, err
	}

	newNStrFmtSpecSignedNumValDto.numFieldLenDto.CopyIn(
		&nStrFmtSpecSignedNumValDto.numFieldLenDto)

	newNStrFmtSpecSignedNumValDto.lock = new(sync.Mutex)

	return newNStrFmtSpecSignedNumValDto, err
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

	ePrefix += "\nnumStrFmtSpecSignedNumValAtom.testValidityOfSignedNumValDto() "

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
