package numberstr

import (
	"fmt"
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
	targetNStrFmtSpecAbsoluteValDto *FormatterAbsoluteValue,
	inComingNStrFmtSpecAbsoluteValDto *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecAbsValDtoNanobot.lock == nil {
		nStrFmtSpecAbsValDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoNanobot.lock.Lock()

	defer nStrFmtSpecAbsValDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecAbsoluteValueDtoNanobot.copyIn()")

	if targetNStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecAbsoluteValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if inComingNStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecAbsoluteValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	nStrFmtSpecAbsValDtoMolecule :=
		formatterAbsoluteValueMolecule{}

	_,
		err = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		inComingNStrFmtSpecAbsoluteValDto,
		ePrefix.XCtx(
			"Testing validity of 'inComingNStrFmtSpecAbsoluteValDto'"))

	if err != nil {
		return err
	}

	targetNStrFmtSpecAbsoluteValDto.absoluteValFmt =
		inComingNStrFmtSpecAbsoluteValDto.absoluteValFmt

	targetNStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation =
		inComingNStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation

	err =
		targetNStrFmtSpecAbsoluteValDto.numericSeparators.CopyIn(
			&inComingNStrFmtSpecAbsoluteValDto.numericSeparators,
			ePrefix.XCtx(
				"Copying inComingNStrFmtSpecAbsoluteValDto->"+
					"targetNStrFmtSpecAbsoluteValDto"))

	if err != nil {
		return err
	}

	err =
		targetNStrFmtSpecAbsoluteValDto.numFieldLenDto.CopyIn(
			&inComingNStrFmtSpecAbsoluteValDto.numFieldLenDto,
			ePrefix.XCtx(
				"inComingNStrFmtSpecAbsoluteValDto->"+
					"targetNStrFmtSpecAbsoluteValDto"))

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecAbsoluteValDto' styled as a new instance
// of FormatterAbsoluteValue.
//
// If input parameter 'nStrFmtSpecAbsoluteValDto' is judged to be
// invalid, this method will return an error.
//
func (nStrFmtSpecAbsValDtoNanobot *numStrFmtSpecAbsoluteValueDtoNanobot) copyOut(
	nStrFmtSpecAbsoluteValDto *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	newNStrFmtSpecAbsoluteValDto FormatterAbsoluteValue,
	err error) {

	if nStrFmtSpecAbsValDtoNanobot.lock == nil {
		nStrFmtSpecAbsValDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoNanobot.lock.Lock()

	defer nStrFmtSpecAbsValDtoNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecAbsoluteValueDtoNanobot.copyOut()")

	if nStrFmtSpecAbsoluteValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecSignedNumValDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNStrFmtSpecAbsoluteValDto, err
	}

	nStrFmtSpecAbsValDtoMolecule :=
		formatterAbsoluteValueMolecule{}

	_,
		err = nStrFmtSpecAbsValDtoMolecule.testValidityOfAbsoluteValDto(
		nStrFmtSpecAbsoluteValDto,
		ePrefix.XCtx("\nTesting validity of 'nStrFmtSpecAbsoluteValDto'"))

	if err != nil {
		return newNStrFmtSpecAbsoluteValDto, err
	}

	newNStrFmtSpecAbsoluteValDto.absoluteValFmt =
		nStrFmtSpecAbsoluteValDto.absoluteValFmt

	newNStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation =
		nStrFmtSpecAbsoluteValDto.turnOnIntegerDigitsSeparation

	err = newNStrFmtSpecAbsoluteValDto.numericSeparators.CopyIn(
		&nStrFmtSpecAbsoluteValDto.numericSeparators,
		ePrefix.XCtx(
			"nStrFmtSpecAbsoluteValDto->"+
				"newNStrFmtSpecAbsoluteValDto"))

	if err != nil {
		return newNStrFmtSpecAbsoluteValDto, err
	}

	newNStrFmtSpecAbsoluteValDto.lock = new(sync.Mutex)

	err =
		newNStrFmtSpecAbsoluteValDto.numFieldLenDto.CopyIn(
			&nStrFmtSpecAbsoluteValDto.numFieldLenDto,
			ePrefix.XCtx(
				"StrFmtSpecAbsoluteValDto->newNStrFmtSpecAbsoluteValDto"))

	return newNStrFmtSpecAbsoluteValDto, err
}
