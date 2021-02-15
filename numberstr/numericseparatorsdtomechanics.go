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
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
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

	if numSepsDto.lock == nil {
		numSepsDto.lock = new(sync.Mutex)
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

	newDigitsSepsDto := NumericSeparatorsDto{}

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
