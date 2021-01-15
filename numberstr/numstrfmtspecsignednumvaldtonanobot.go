package numberstr

import (
	"fmt"
	"sync"
)

type nStrFmtSpecSignedNumValNanobot struct {
	lock *sync.Mutex
}

// setSignedNumValDto - Transfers new data to an instance of
// NumStrFmtSpecSignedNumValueDto. After completion, all the data
// fields within input parameter 'nStrFmtSpecDigitsSepDto' will be
// overwritten.
//
func (nStrFmtSpecSignedNumValNanobot *nStrFmtSpecSignedNumValNanobot) setSignedNumValDto(
	nStrFmtSpecSignedNumValDto *NumStrFmtSpecSignedNumValueDto,
	positiveValueFmt string,
	negativeValueFmt string,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	numFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecSignedNumValNanobot.lock == nil {
		nStrFmtSpecSignedNumValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecSignedNumValNanobot.lock.Lock()

	defer nStrFmtSpecSignedNumValNanobot.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "nStrFmtSpecSignedNumValNanobot.setSignedNumValDto() "

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

	nStrFmtSpecSignedNumValMolecule :=
		nStrFmtSpecSignedNumValMolecule{}

	err =
		nStrFmtSpecSignedNumValMolecule.copyIn(
			nStrFmtSpecSignedNumValDto,
			&newNStrFmtSpecSignedNumValDto,
			ePrefix+
				"\nnewNStrFmtSpecSignedNumValDto-> "+
				"nStrFmtSpecSignedNumValDto\n ")

	return err
}
