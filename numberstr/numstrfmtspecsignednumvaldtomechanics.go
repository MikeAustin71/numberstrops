package numberstr

import (
	"fmt"
	"sync"
)

type nStrFmtSpecSignedNumValMechanics struct {
	lock *sync.Mutex
}

// setSignedNumValDto - Transfers new data to an instance of
// NumStrFmtSpecSignedNumValueDto. After completion, all the data
// fields within input parameter 'nStrFmtSpecDigitsSepDto' will be
// overwritten.
//
func (nStrFmtSpecSignedNumValMech *nStrFmtSpecSignedNumValMechanics) setSignedNumValDto(
	nStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	numFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecSignedNumValMech.lock == nil {
		nStrFmtSpecSignedNumValMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValMech.lock.Lock()

	defer nStrFmtSpecSignedNumValMech.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "nStrFmtSpecSignedNumValMech.setSignedNumValDto() "

	if nStrFmtSpecSignedNumValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is invalid!\n"+
			"'nStrFmtSpecSignedNumValDto' is a 'nil' pointer\n",
			ePrefix)

		return err
	}

	newNStrFmtSpecSignedNumValDto := NumStrFmtSpecSignedNumValueDto{}

	newNStrFmtSpecSignedNumValDto.positiveValueFmt =
		positiveValueFmt

	newNStrFmtSpecSignedNumValDto.negativeValueFmt =
		negativeValueFmt

	newNStrFmtSpecSignedNumValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newNStrFmtSpecSignedNumValDto.numberSeparatorsDto.CopyIn(
			&numberSeparatorsDto,
			ePrefix+
				"\nnumberSeparatorsDto->newNStrFmtSpecSignedNumValDto\n ")

	if err != nil {
		return err
	}

	newNStrFmtSpecSignedNumValDto.numFieldLenDto.CopyIn(
		&numFieldLenDto)

	nStrFmtSpecSignedNumValNanobot :=
		numStrFmtSpecSignedNumValNanobot{}

	err =
		nStrFmtSpecSignedNumValNanobot.copyIn(
			nStrFmtSpecSignedNumValDto,
			&newNStrFmtSpecSignedNumValDto,
			ePrefix+
				"\nnewNStrFmtSpecSignedNumValDto-> "+
				"nStrFmtSpecSignedNumValDto\n ")

	return err
}
