package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoNanobot struct {
	lock *sync.Mutex
}

// setAbsValDto - Sets the data value for incoming parameter
// 'nStrFmtSpecAbValDto', an instance of
// NumStrFmtSpecAbsoluteValueDto.
//
func (nStrFmtSpecAbsValDtoNanobot *numStrFmtSpecAbsoluteValueDtoNanobot) setAbsValDto(
	nStrFmtSpecAbValDto *NumStrFmtSpecAbsoluteValueDto,
	absoluteValueFormat string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	numFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecAbsValDtoNanobot.lock == nil {
		nStrFmtSpecAbsValDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoNanobot.lock.Lock()

	defer nStrFmtSpecAbsValDtoNanobot.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecAbsoluteValueDtoNanobot.setAbsValDto()\n "

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

	newNStrFmtSpecAbsValDto.numFieldLenDto.CopyIn(
		&numFieldLenDto)

	nStrFmtSpecAbsValDtoAtom :=
		numStrFmtSpecAbsoluteValueDtoAtom{}

	_,
		err = nStrFmtSpecAbsValDtoAtom.testValidityOfAbsoluteValDto(
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
