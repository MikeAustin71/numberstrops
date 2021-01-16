package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCurrencyValueDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'targetNStrFmtSpecCurrencyValDto' to input parameter
// 'inComingNStrFmtSpecCurrencyValDto'.
//
// Be advised - All data fields in 'targetNStrFmtSpecCurrencyValDto'
// will be overwritten.
//
// If input parameter 'inComingNStrFmtSpecCurrencyValDto' is judged
// to be invalid, this method will return an error.
//
func (nStrFmtSpecCurrValNanobot *numStrFmtSpecCurrencyValueDtoNanobot) copyIn(
	targetNStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	inComingNStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecCurrencyValueDtoNanobot.copyIn()\n "

	if targetNStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if inComingNStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err =
		nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
			inComingNStrFmtSpecCurrencyValDto,
			ePrefix+
				"\nTesting validity of 'inComingNStrFmtSpecCurrencyValDto'\n ")

	if err != nil {
		return err
	}

	targetNStrFmtSpecCurrencyValDto.positiveValueFmt =
		inComingNStrFmtSpecCurrencyValDto.positiveValueFmt

	targetNStrFmtSpecCurrencyValDto.negativeValueFmt =
		inComingNStrFmtSpecCurrencyValDto.negativeValueFmt

	targetNStrFmtSpecCurrencyValDto.decimalDigits =
		inComingNStrFmtSpecCurrencyValDto.decimalDigits

	targetNStrFmtSpecCurrencyValDto.currencyCode =
		inComingNStrFmtSpecCurrencyValDto.currencyCode

	targetNStrFmtSpecCurrencyValDto.currencyName =
		inComingNStrFmtSpecCurrencyValDto.currencyName

	targetNStrFmtSpecCurrencyValDto.currencySymbol =
		inComingNStrFmtSpecCurrencyValDto.currencySymbol

	targetNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		inComingNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation

	err =
		targetNStrFmtSpecCurrencyValDto.numberSeparatorsDto.CopyIn(
			&inComingNStrFmtSpecCurrencyValDto.numberSeparatorsDto,
			ePrefix+
				"'inComingNStrFmtSpecCurrencyValDto' -> "+
				"'targetNStrFmtSpecCurrencyValDto'\n ")

	if err != nil {
		return err
	}

	targetNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
		&inComingNStrFmtSpecCurrencyValDto.numFieldLenDto)

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecCurrencyValDto' styled as a new instance
// of NumStrFmtSpecCurrencyValueDto.
//
// If input parameter 'nStrFmtSpecCurrencyValDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecCurrValNanobot *numStrFmtSpecCurrencyValueDtoNanobot) copyOut(
	nStrFmtSpecCurrencyValDto *NumStrFmtSpecCurrencyValueDto,
	ePrefix string) (
	newNStrFmtSpecCurrencyValDto NumStrFmtSpecCurrencyValueDto,
	err error) {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecCurrencyValueDtoNanobot.copyOut() "

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return newNStrFmtSpecCurrencyValDto, err
	}

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err =
		nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
			nStrFmtSpecCurrencyValDto,
			ePrefix+
				"Testing validity of 'nStrFmtSpecCurrencyValDto'\n ")

	if err != nil {
		return newNStrFmtSpecCurrencyValDto, err
	}

	newNStrFmtSpecCurrencyValDto.positiveValueFmt =
		nStrFmtSpecCurrencyValDto.positiveValueFmt

	newNStrFmtSpecCurrencyValDto.negativeValueFmt =
		nStrFmtSpecCurrencyValDto.negativeValueFmt

	newNStrFmtSpecCurrencyValDto.decimalDigits =
		nStrFmtSpecCurrencyValDto.decimalDigits

	newNStrFmtSpecCurrencyValDto.currencyCode =
		nStrFmtSpecCurrencyValDto.currencyCode

	newNStrFmtSpecCurrencyValDto.currencyName =
		nStrFmtSpecCurrencyValDto.currencyName

	newNStrFmtSpecCurrencyValDto.currencySymbol =
		nStrFmtSpecCurrencyValDto.currencySymbol

	newNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		nStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation

	err = newNStrFmtSpecCurrencyValDto.numberSeparatorsDto.CopyIn(
		&nStrFmtSpecCurrencyValDto.numberSeparatorsDto,
		ePrefix+
			"nStrFmtSpecCurrencyValDto->newNStrFmtSpecCurrencyValDto\n ")

	if err != nil {
		return newNStrFmtSpecCurrencyValDto, err
	}

	newNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
		&nStrFmtSpecCurrencyValDto.numFieldLenDto)

	newNStrFmtSpecCurrencyValDto.lock = new(sync.Mutex)

	return newNStrFmtSpecCurrencyValDto, err
}
