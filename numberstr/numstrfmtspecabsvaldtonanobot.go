package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecAbsoluteValDto' to input parameter
// 'targetNStrFmtSpecAbsoluteValDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecAbsoluteValDto'
// will be overwritten.
//
// If the 'inComingNStrFmtSpecAbsoluteValDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecAbsValDtoNanobot *numStrFmtSpecAbsoluteValueDtoNanobot) copyIn(
	targetNStrFmtSpecAbsoluteValDto *NumStrFmtSpecAbsoluteValueDto,
	inComingNStrFmtSpecAbsoluteValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecAbsValDtoNanobot.lock == nil {
		nStrFmtSpecAbsValDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoNanobot.lock.Lock()

	defer nStrFmtSpecAbsValDtoNanobot.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecAbsoluteValueDtoNanobot.copyIn() "

	if targetNStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecAbsoluteValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if inComingNStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecAbsoluteValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	_,
		err = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		inComingNStrFmtSpecAbsoluteValDto,
		ePrefix+
			"\nTesting validity of 'inComingNStrFmtSpecAbsoluteValDto'\n ")

	if err != nil {
		return err
	}

	targetNStrFmtSpecAbsoluteValDto.absoluteValFmt =
		inComingNStrFmtSpecAbsoluteValDto.absoluteValFmt

	targetNStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation =
		inComingNStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation

	err =
		targetNStrFmtSpecAbsoluteValDto.numberSeparatorsDto.CopyIn(
			&inComingNStrFmtSpecAbsoluteValDto.numberSeparatorsDto,
			ePrefix+
				"Copying inComingNStrFmtSpecAbsoluteValDto->"+
				"targetNStrFmtSpecAbsoluteValDto\n ")

	if err != nil {
		return err
	}

	err =
		targetNStrFmtSpecAbsoluteValDto.numFieldLenDto.CopyIn(
			&inComingNStrFmtSpecAbsoluteValDto.numFieldLenDto,
			ePrefix+
				"\ninComingNStrFmtSpecAbsoluteValDto->targetNStrFmtSpecAbsoluteValDto\n")

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecAbsoluteValDto' styled as a new instance
// of NumStrFmtSpecAbsoluteValueDto.
//
// If input parameter 'nStrFmtSpecAbsoluteValDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecAbsValDtoNanobot *numStrFmtSpecAbsoluteValueDtoNanobot) copyOut(
	nStrFmtSpecAbsoluteValDto *NumStrFmtSpecAbsoluteValueDto,
	ePrefix string) (
	newNStrFmtSpecAbsoluteValDto NumStrFmtSpecAbsoluteValueDto,
	err error) {

	if nStrFmtSpecAbsValDtoNanobot.lock == nil {
		nStrFmtSpecAbsValDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoNanobot.lock.Lock()

	defer nStrFmtSpecAbsValDtoNanobot.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecAbsoluteValueDtoNanobot.copyOut() "

	if nStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return newNStrFmtSpecAbsoluteValDto, err
	}

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	_,
		err = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		nStrFmtSpecAbsoluteValDto,
		ePrefix+
			"\nTesting validity of 'nStrFmtSpecAbsoluteValDto'\n ")

	if err != nil {
		return newNStrFmtSpecAbsoluteValDto, err
	}

	newNStrFmtSpecAbsoluteValDto.absoluteValFmt =
		nStrFmtSpecAbsoluteValDto.absoluteValFmt

	newNStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation =
		nStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation

	err = newNStrFmtSpecAbsoluteValDto.numberSeparatorsDto.CopyIn(
		&nStrFmtSpecAbsoluteValDto.numberSeparatorsDto,
		ePrefix+
			"\nnStrFmtSpecAbsoluteValDto->newNStrFmtSpecAbsoluteValDto ")

	if err != nil {
		return newNStrFmtSpecAbsoluteValDto, err
	}

	newNStrFmtSpecAbsoluteValDto.lock = new(sync.Mutex)

	err =
		newNStrFmtSpecAbsoluteValDto.numFieldLenDto.CopyIn(
			&nStrFmtSpecAbsoluteValDto.numFieldLenDto,
			ePrefix+
				"\nStrFmtSpecAbsoluteValDto->newNStrFmtSpecAbsoluteValDto\n")

	return newNStrFmtSpecAbsoluteValDto, err
}
