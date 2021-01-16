package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoMechanics struct {
	lock *sync.Mutex
}

// setCurrencyValDto - Transfers new data to an instance of
// NumStrFmtSpecCurrencyValueDto. After completion, all the data
// fields within input parameter 'nStrFmtSpecCurrencyValDto' will be
// overwritten.
//
func (nStrFmtSpecCurrValMech *numStrFmtSpecCurrencyValueDtoMechanics) setCurrencyValDto(
	nStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyName string,
	currencySymbol rune,
	turnOnIntegerDigitsSeparation bool,
	numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto,
	numFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValMech.lock == nil {
		nStrFmtSpecCurrValMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValMech.lock.Lock()

	defer nStrFmtSpecCurrValMech.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecCurrencyValueDtoMechanics.setCurrencyValDto() "

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is invalid!\n"+
			"'nStrFmtSpecCurrencyValDto' is a 'nil' pointer\n",
			ePrefix)

		return err
	}

	newNStrFmtSpecCurrencyValDto := NumStrFmtSpecCurrencyValueDto{}

	newNStrFmtSpecCurrencyValDto.positiveValueFmt =
		positiveValueFmt

	newNStrFmtSpecCurrencyValDto.negativeValueFmt =
		negativeValueFmt

	newNStrFmtSpecCurrencyValDto.decimalDigits =
		decimalDigits

	newNStrFmtSpecCurrencyValDto.currencyCode =
		currencyCode

	newNStrFmtSpecCurrencyValDto.currencyName =
		currencyName

	newNStrFmtSpecCurrencyValDto.currencySymbol =
		currencySymbol

	newNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	err =
		newNStrFmtSpecCurrencyValDto.numberSeparatorsDto.CopyIn(
			&numberSeparatorsDto,
			ePrefix+
				"\nnumberSeparatorsDto->newNStrFmtSpecCurrencyValDto\n ")

	if err != nil {
		return err
	}

	newNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
		&numFieldLenDto)

	nStrFmtSpecCurrValNanobot :=
		numStrFmtSpecCurrencyValueDtoNanobot{}

	err =
		nStrFmtSpecCurrValNanobot.copyIn(
			nStrFmtSpecCurrencyValDto,
			&newNStrFmtSpecCurrencyValDto,
			ePrefix+
				"\nnewNStrFmtSpecCurrencyValDto-> "+
				"nStrFmtSpecCurrencyValDto\n ")

	return err
}
