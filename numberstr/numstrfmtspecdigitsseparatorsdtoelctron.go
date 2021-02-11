package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDigitsSeparatorsDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecDigitsSepsDto' to input parameter
// 'targetNStrFmtSpecDigitsSepsDto'.
//
// If 'inComingNStrFmtSpecDigitsSepsDto' is judged to be invalid,
// this method will return an error.
//
// This method will OVERWRITE the data fields of the
// 'targetNStrFmtSpecDigitsSepsDto' instance.
//
func (nStrFmtSpecDigitsSepsElectron *numStrFmtSpecDigitsSeparatorsDtoElectron) copyIn(
	targetNStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	inComingNStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecDigitsSepsElectron.lock == nil {
		nStrFmtSpecDigitsSepsElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsElectron.lock.Lock()

	defer nStrFmtSpecDigitsSepsElectron.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecDigitsSeparatorsDtoElectron.copyIn()")

	if targetNStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'targetNStrFmtSpecDigitsSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if inComingNStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'inComingNStrFmtSpecDigitsSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetNStrFmtSpecDigitsSepsDto.lock == nil {
		targetNStrFmtSpecDigitsSepsDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark :=
		numStrFmtSpecDigitsSeparatorsDtoQuark{}

	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
			inComingNStrFmtSpecDigitsSepsDto,
			ePrefix.XCtx("Testing validity of 'inComingNStrFmtSpecDigitsSepsDto'"))

	if err != nil {
		return err
	}

	targetNStrFmtSpecDigitsSepsDto.decimalSeparator =
		inComingNStrFmtSpecDigitsSepsDto.decimalSeparator

	targetNStrFmtSpecDigitsSepsDto.integerDigitsSeparator =
		inComingNStrFmtSpecDigitsSepsDto.integerDigitsSeparator

	lenIntDigitsGroupSeq :=
		len(inComingNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence)

	targetNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupSeq, lenIntDigitsGroupSeq+5)

	if lenIntDigitsGroupSeq == 0 {
		return
	}

	_ =
		copy(targetNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence,
			inComingNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence)

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecDigitsSepsDto' styled as a new instance
// of NumStrFmtSpecDigitsSeparatorsDto.
//
func (nStrFmtSpecDigitsSepsElectron *numStrFmtSpecDigitsSeparatorsDtoElectron) copyOut(
	nStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	newDigitsSepsDto NumStrFmtSpecDigitsSeparatorsDto,
	err error) {

	if nStrFmtSpecDigitsSepsElectron.lock == nil {
		nStrFmtSpecDigitsSepsElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsElectron.lock.Lock()

	defer nStrFmtSpecDigitsSepsElectron.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecDigitsSeparatorsDtoElectron.copyOut()")

	if nStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newDigitsSepsDto, err
	}

	nStrFmtSpecDigitsSepsQuark :=
		numStrFmtSpecDigitsSeparatorsDtoQuark{}

	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
			nStrFmtSpecDigitsSepsDto,
			ePrefix.XCtx("Testing validity of 'nStrFmtSpecDigitsSepsDto'"))

	if err != nil {
		return newDigitsSepsDto, err
	}

	newDigitsSepsDto.decimalSeparator =
		nStrFmtSpecDigitsSepsDto.decimalSeparator

	newDigitsSepsDto.integerDigitsSeparator =
		nStrFmtSpecDigitsSepsDto.integerDigitsSeparator

	lenIntDigitsGroupingSequence :=
		len(nStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence)

	newDigitsSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		return newDigitsSepsDto, err
	}

	_ = copy(
		newDigitsSepsDto.integerDigitsGroupingSequence,
		nStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence)

	newDigitsSepsDto.lock = new(sync.Mutex)

	return newDigitsSepsDto, err
}
