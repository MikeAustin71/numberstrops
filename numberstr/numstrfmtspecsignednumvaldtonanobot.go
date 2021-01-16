package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSignedNumValNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecSignedNumValDto' to input parameter
// 'targetNStrFmtSpecSignedNumValDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecSignedNumValDto'
// will be overwritten.
//
// If input parameter 'inComingNStrFmtSpecSignedNumValDto' is judged
// to be invalid, this method will return an error.
//
func (nStrFmtSpecSignedNumValNanobot *numStrFmtSpecSignedNumValNanobot) copyIn(
	targetNStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	inComingNStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecSignedNumValNanobot.lock == nil {
		nStrFmtSpecSignedNumValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValNanobot.lock.Lock()

	defer nStrFmtSpecSignedNumValNanobot.lock.Unlock()

	ePrefix += "\nnStrFmtSpecSignedNumValNanobot.copyIn()\n "

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

	nStrFmtSpecSignedNumValMolecule :=
		numStrFmtSpecSignedNumValMolecule{}

	_,
		err =
		nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
			inComingNStrFmtSpecSignedNumValDto,
			ePrefix+
				"Testing validity of 'inComingNStrFmtSpecSignedNumValDto'\n ")

	if err != nil {
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
				"'inComingNStrFmtSpecSignedNumValDto' -> "+
				"'targetNStrFmtSpecSignedNumValDto'\n ")

	if err != nil {
		return err
	}

	targetNStrFmtSpecSignedNumValDto.numFieldLenDto.CopyIn(
		&inComingNStrFmtSpecSignedNumValDto.numFieldLenDto)

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecSignedNumValDto' styled as a new instance
// of NumStrFmtSpecSignedNumValueDto.
//
// If input parameter 'nStrFmtSpecSignedNumValDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecSignedNumValNanobot *numStrFmtSpecSignedNumValNanobot) copyOut(
	nStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	ePrefix string) (
	newNStrFmtSpecSignedNumValDto NumStrFmtSpecSignedNumValueDto,
	err error) {

	if nStrFmtSpecSignedNumValNanobot.lock == nil {
		nStrFmtSpecSignedNumValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValNanobot.lock.Lock()

	defer nStrFmtSpecSignedNumValNanobot.lock.Unlock()

	ePrefix += "\nnStrFmtSpecSignedNumValNanobot.copyOut()\n "

	if nStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return newNStrFmtSpecSignedNumValDto, err
	}

	nStrFmtSpecSignedNumValMolecule :=
		numStrFmtSpecSignedNumValMolecule{}

	_,
		err =
		nStrFmtSpecSignedNumValMolecule.testValidityOfSignedNumValDto(
			nStrFmtSpecSignedNumValDto,
			ePrefix+
				"Testing validity of 'nStrFmtSpecSignedNumValDto'\n ")

	if err != nil {
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
			"nStrFmtSpecSignedNumValDto->newNStrFmtSpecSignedNumValDto\n ")

	if err != nil {
		return newNStrFmtSpecSignedNumValDto, err
	}

	newNStrFmtSpecSignedNumValDto.numFieldLenDto.CopyIn(
		&nStrFmtSpecSignedNumValDto.numFieldLenDto)

	newNStrFmtSpecSignedNumValDto.lock = new(sync.Mutex)

	return newNStrFmtSpecSignedNumValDto, err
}
