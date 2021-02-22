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
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoNanobot.copyIn()")

	if targetNStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if inComingNStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err =
		nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
			inComingNStrFmtSpecCurrencyValDto,
			ePrefix.XCtx("Testing validity of 'inComingNStrFmtSpecCurrencyValDto'"))

	if err != nil {
		return err
	}

	lenCurrencySymbols := len(inComingNStrFmtSpecCurrencyValDto.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'inComingNStrFmtSpecCurrencyValDto.currencySymbols' invalid!\n"+
			"The currencySymbols array is a zero length array.\n",
			ePrefix.XCtxEmpty().String())

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

	targetNStrFmtSpecCurrencyValDto.currencyCodeNo =
		inComingNStrFmtSpecCurrencyValDto.currencyCodeNo

	targetNStrFmtSpecCurrencyValDto.currencyName =
		inComingNStrFmtSpecCurrencyValDto.currencyName

	targetNStrFmtSpecCurrencyValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		targetNStrFmtSpecCurrencyValDto.currencySymbols,
		inComingNStrFmtSpecCurrencyValDto.currencySymbols)

	targetNStrFmtSpecCurrencyValDto.minorCurrencyName =
		inComingNStrFmtSpecCurrencyValDto.minorCurrencyName

	lenCurrencySymbols = len(inComingNStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	targetNStrFmtSpecCurrencyValDto.minorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		targetNStrFmtSpecCurrencyValDto.minorCurrencySymbols,
		inComingNStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	targetNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		inComingNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation

	err =
		targetNStrFmtSpecCurrencyValDto.numberSeparatorsDto.CopyIn(
			&inComingNStrFmtSpecCurrencyValDto.numberSeparatorsDto,
			ePrefix.XCtx("'inComingNStrFmtSpecCurrencyValDto' -> "+
				"'targetNStrFmtSpecCurrencyValDto'"))

	if err != nil {
		return err
	}

	err =
		targetNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
			&inComingNStrFmtSpecCurrencyValDto.numFieldLenDto,
			ePrefix.XCtx("\ninComingNStrFmtSpecCurrencyValDto->targetNStrFmtSpecCurrencyValDto"))

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
	ePrefix *ErrPrefixDto) (
	newNStrFmtSpecCurrencyValDto NumStrFmtSpecCurrencyValueDto,
	err error) {

	if nStrFmtSpecCurrValNanobot.lock == nil {
		nStrFmtSpecCurrValNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecCurrValNanobot.lock.Lock()

	defer nStrFmtSpecCurrValNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecCurrencyValueDtoNanobot.copyOut()")

	if nStrFmtSpecCurrencyValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecCurrencyValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNStrFmtSpecCurrencyValDto, err
	}

	nStrFmtSpecCurrDtoMolecule :=
		numStrFmtSpecCurrencyValueDtoMolecule{}

	_,
		err =
		nStrFmtSpecCurrDtoMolecule.testValidityOfCurrencyValDto(
			nStrFmtSpecCurrencyValDto,
			ePrefix.XCtx("Testing validity of 'nStrFmtSpecCurrencyValDto'"))

	if err != nil {
		return newNStrFmtSpecCurrencyValDto, err
	}

	lenCurrencySymbols :=
		len(nStrFmtSpecCurrencyValDto.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'nStrFmtSpecCurrencyValDto.currencySymbols' invalid!\n"+
			"The currencySymbols array is a zero length array.\n",
			ePrefix.XCtxEmpty().String())

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

	newNStrFmtSpecCurrencyValDto.currencyCodeNo =
		nStrFmtSpecCurrencyValDto.currencyCodeNo

	newNStrFmtSpecCurrencyValDto.currencyName =
		nStrFmtSpecCurrencyValDto.currencyName

	newNStrFmtSpecCurrencyValDto.currencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		newNStrFmtSpecCurrencyValDto.currencySymbols,
		nStrFmtSpecCurrencyValDto.currencySymbols)

	newNStrFmtSpecCurrencyValDto.minorCurrencyName =
		nStrFmtSpecCurrencyValDto.minorCurrencyName

	lenCurrencySymbols =
		len(nStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	newNStrFmtSpecCurrencyValDto.minorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		newNStrFmtSpecCurrencyValDto.minorCurrencySymbols,
		nStrFmtSpecCurrencyValDto.minorCurrencySymbols)

	newNStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation =
		nStrFmtSpecCurrencyValDto.turnOnIntegerDigitsSeparation

	err = newNStrFmtSpecCurrencyValDto.numberSeparatorsDto.CopyIn(
		&nStrFmtSpecCurrencyValDto.numberSeparatorsDto,
		ePrefix.XCtx("nStrFmtSpecCurrencyValDto->newNStrFmtSpecCurrencyValDto"))

	if err != nil {
		return newNStrFmtSpecCurrencyValDto, err
	}

	err =
		newNStrFmtSpecCurrencyValDto.numFieldLenDto.CopyIn(
			&nStrFmtSpecCurrencyValDto.numFieldLenDto,
			ePrefix.XCtx("nStrFmtSpecCurrencyValDto->newNStrFmtSpecCurrencyValDto"))

	newNStrFmtSpecCurrencyValDto.lock = new(sync.Mutex)

	return newNStrFmtSpecCurrencyValDto, err
}
