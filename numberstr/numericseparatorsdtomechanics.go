package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorDtoMechanics struct {
	lock *sync.Mutex
}

// setDigitsSeps - Transfers new data to an instance of NumericSeparatorsDto.
// After completion, all the data fields within input parameter 'nStrFmtSpecDigitsSepDto'
// will be overwritten.
//
func (numSepsDtoMech *numericSeparatorDtoMechanics) setDigitsSeps(
	numSepsDto *NumericSeparatorsDto,
	decimalSeparator rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsDtoMech.lock == nil {
		numSepsDtoMech.lock = new(sync.Mutex)
	}

	numSepsDtoMech.lock.Lock()

	defer numSepsDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numericSeparatorDtoMechanics.setDigitsSeps()")

	if numSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSepsDto' is invalid!\n"+
			"'numSepsDto' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if numSepsDto.lock == nil {
		numSepsDto.lock = new(sync.Mutex)
	}

	if decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is invalid!\n"+
			"decimalSeparator == 0\n",
			ePrefix.String())

		return err
	}

	lenIntDigitSeparators :=
		len(integerSeparators)

	if lenIntDigitSeparators == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerSeparators' is invalid!\n"+
			"'integerSeparators' is a zero length or empty array.\n",
			ePrefix.String())

		return err
	}

	newDigitsSepsDto := NumericSeparatorsDto{}

	newDigitsSepsDto.decimalSeparator =
		decimalSeparator

	newDigitsSepsDto.integerSeparators =
		make([]NumStrIntSeparator,
			lenIntDigitSeparators,
			10)

	elementsCopied :=
		copy(newDigitsSepsDto.integerSeparators,
			integerSeparators)

	if elementsCopied != lenIntDigitSeparators {
		err = fmt.Errorf("%v\n"+
			"Error: Copy of integerSeparators -> "+
			"newDigitsSepsDto.integerSeparators failed!\n"+
			"Expected to copy %v elements.\n"+
			"Actually copied %v elements.\n",
			ePrefix,
			lenIntDigitSeparators,
			elementsCopied)

		return err
	}

	newDigitsSepsDto.lock = new(sync.Mutex)

	nStrFmtSpecDigitsSepsQuark := numericSeparatorDtoQuark{}
	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
			&newDigitsSepsDto,
			ePrefix.XCtx("Testing validity of preliminary 'newDigitsSepsDto'"))

	if err != nil {
		return err
	}

	nStrFmtSpecDigitsSepsElectron :=
		numericSeparatorDtoElectron{}

	err =
		nStrFmtSpecDigitsSepsElectron.copyIn(
			numSepsDto,
			&newDigitsSepsDto,
			ePrefix.XCtx("newDigitsSepsDto->numSepsDto"))

	return err
}
