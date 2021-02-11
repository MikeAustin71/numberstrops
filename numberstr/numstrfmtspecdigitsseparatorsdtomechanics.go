package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDigitsSepsDtoMechanics struct {
	lock *sync.Mutex
}

// setDigitsSeps - Transfers new data to an instance of NumStrFmtSpecDigitsSeparatorsDto.
// After completion, all the data fields within input parameter 'nStrFmtSpecDigitsSepDto'
// will be overwritten.
//
func (nStrFmtSpecDigitsSepsDtoMech *numStrFmtSpecDigitsSepsDtoMechanics) setDigitsSeps(
	nStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	decimalSeparator rune,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecDigitsSepsDtoMech.lock == nil {
		nStrFmtSpecDigitsSepsDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsDtoMech.lock.Lock()

	defer nStrFmtSpecDigitsSepsDtoMech.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecDigitsSepsDtoMechanics.setDigitsSeps()")

	if nStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepsDto' is invalid!\n"+
			"'nStrFmtSpecDigitsSepsDto' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is invalid!\n"+
			"decimalSeparator == 0\n",
			ePrefix.String())

		return err
	}

	if integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparator' is invalid!\n"+
			"integerDigitsSeparator == 0\n",
			ePrefix.String())

		return err
	}

	if nStrFmtSpecDigitsSepsDto.lock == nil {
		nStrFmtSpecDigitsSepsDto.lock = new(sync.Mutex)
	}

	lenIntDigitsGroupingSequence :=
		len(integerDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsGroupingSequence' is invalid!\n"+
			"'integerDigitsGroupingSequence' is a zero length array or empty array.\n",
			ePrefix.String())

		return err
	}

	newDigitsSepsDto := NumStrFmtSpecDigitsSeparatorsDto{}

	newDigitsSepsDto.decimalSeparator =
		decimalSeparator

	newDigitsSepsDto.integerDigitsSeparator =
		integerDigitsSeparator

	newDigitsSepsDto.integerDigitsGroupingSequence =
		make([]uint, 0, 10)

	newDigitsSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence, 10)

	for i := 0; i < lenIntDigitsGroupingSequence; i++ {
		newDigitsSepsDto.integerDigitsGroupingSequence[i] =
			integerDigitsGroupingSequence[i]
	}

	newDigitsSepsDto.lock = new(sync.Mutex)

	nStrFmtSpecDigitsSepsQuark := numStrFmtSpecDigitsSeparatorsDtoQuark{}
	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
			&newDigitsSepsDto,
			ePrefix.XCtx("Testing validity of preliminary 'newDigitsSepsDto'"))

	if err != nil {
		return err
	}

	nStrFmtSpecDigitsSepsElectron :=
		numStrFmtSpecDigitsSeparatorsDtoElectron{}

	err =
		nStrFmtSpecDigitsSepsElectron.copyIn(
			nStrFmtSpecDigitsSepsDto,
			&newDigitsSepsDto,
			ePrefix.XCtx("newDigitsSepsDto->nStrFmtSpecDigitsSepsDto"))

	return err
}
