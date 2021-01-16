package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDigitsSeparatorsDtoElectron struct {
	lock *sync.Mutex
}

// setDigitsSeps - Transfers new data to an instance of NumStrFmtSpecDigitsSeparatorsDto.
// After completion, all the data fields within input parameter 'nStrFmtSpecDigitsSepDto'
// will be overwritten.
//
func (nStrFmtSpecDigitsSepsDtoElectron *numStrFmtSpecDigitsSeparatorsDtoElectron) setDigitsSeps(
	nStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	decimalSeparator rune,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix string) (
	err error) {

	if nStrFmtSpecDigitsSepsDtoElectron.lock == nil {
		nStrFmtSpecDigitsSepsDtoElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsDtoElectron.lock.Lock()

	defer nStrFmtSpecDigitsSepsDtoElectron.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecDigitsSeparatorsDtoElectron.setDigitsSeps() "

	if nStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepsDto' is invalid!\n"+
			"'nStrFmtSpecDigitsSepsDto' is a 'nil' pointer\n",
			ePrefix)

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
			ePrefix)

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
			ePrefix+
				"\nTesting validity of preliminary 'newDigitsSepsDto'\n ")

	if err != nil {
		return err
	}

	err =
		nStrFmtSpecDigitsSepsQuark.copyIn(
			nStrFmtSpecDigitsSepsDto,
			&newDigitsSepsDto,
			ePrefix+
				"\nnewDigitsSepsDto->nStrFmtSpecDigitsSepsDto\n ")

	return err
}