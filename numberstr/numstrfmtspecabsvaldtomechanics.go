package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoMechanics struct {
	lock *sync.Mutex
}

// setAbsValDto - Sets the data value for incoming parameter
// 'nStrFmtSpecAbValDto', an instance of
// NumStrFmtSpecAbsoluteValueDto.
//
func (nStrFmtSpecAbsValDtoMech *numStrFmtSpecAbsoluteValueDtoMechanics) setAbsValDto(
	nStrFmtSpecAbValDto *NumStrFmtSpecAbsoluteValueDto,
	absoluteValueFormat string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	numFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecAbsValDtoMech.lock == nil {
		nStrFmtSpecAbsValDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoMech.lock.Lock()

	defer nStrFmtSpecAbsValDtoMech.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecAbsoluteValueDtoMechanics.setAbsValDto()\n "

	if nStrFmtSpecAbValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	newNStrFmtSpecAbsValDto :=
		NumStrFmtSpecAbsoluteValueDto{}

	newNStrFmtSpecAbsValDto.absoluteValFmt =
		absoluteValueFormat

	newNStrFmtSpecAbsValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newNStrFmtSpecAbsValDto.numberSeparatorsDto.CopyIn(
			&numberSeparatorsDto,
			ePrefix+
				"'numberSeparatorsDto'->'newNStrFmtSpecAbsValDto'\n ")

	if err != nil {
		return err
	}

	err =
		newNStrFmtSpecAbsValDto.numFieldLenDto.CopyIn(
			&numFieldLenDto,
			ePrefix+"numFieldLenDto")

	nStrFmtSpecAbsValDtoMolecule :=
		numStrFmtSpecAbsoluteValueDtoMolecule{}

	_,
		err = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		&newNStrFmtSpecAbsValDto,
		ePrefix+
			"Testing validity of 'newNStrFmtSpecAbsValDto'\n ")

	if err != nil {
		return err
	}

	err =
		nStrFmtSpecAbValDto.CopyIn(
			&newNStrFmtSpecAbsValDto,
			ePrefix+
				"'newNStrFmtSpecAbsValDto'->'nStrFmtSpecAbValDto'\n ")

	return err
}
