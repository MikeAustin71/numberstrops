package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDigitsSeparatorsDtoQuark struct {
	lock *sync.Mutex
}

// CopyIn - Copies the data fields from input parameter
// 'inComingNStrFmtSpecDigitsSepsDto' to input parameter
// 'targetNStrFmtSpecDigitsSepsDto'.
//
func (nStrFmtSpecDigitsSepsQuark *numStrFmtSpecDigitsSeparatorsDtoQuark) copyIn(
	targetNStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	inComingNStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecDigitsSepsQuark.lock == nil {
		nStrFmtSpecDigitsSepsQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark.lock.Lock()

	defer nStrFmtSpecDigitsSepsQuark.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecDigitsSeparatorsDtoQuark.copyIn() "

	if targetNStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrFmtSpecDigitsSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if inComingNStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingNStrFmtSpecDigitsSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	targetNStrFmtSpecDigitsSepsDto.decimalSeparator =
		inComingNStrFmtSpecDigitsSepsDto.decimalSeparator

	targetNStrFmtSpecDigitsSepsDto.integerDigitsSeparator =
		inComingNStrFmtSpecDigitsSepsDto.integerDigitsSeparator

	lenIntDigitsGroupingSequence :=
		len(inComingNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		targetNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
		return
	}

	targetNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence)

	for i := 0; i < lenIntDigitsGroupingSequence; i++ {
		targetNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence[i] =
			inComingNStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence[i]
	}

	return err
}

// CopyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecDigitsSepsDto' styled as a new instance
// of NumStrFmtSpecDigitsSeparatorsDto.
//
func (nStrFmtSpecDigitsSepsQuark *numStrFmtSpecDigitsSeparatorsDtoQuark) copyOut(
	nStrFmtSpecDigitsSepsDto *NumStrFmtSpecDigitsSeparatorsDto,
	ePrefix string) (
	newDigitsSepsDto NumStrFmtSpecDigitsSeparatorsDto,
	err error) {

	ePrefix += "\nnumStrFmtSpecDigitsSeparatorsDtoQuark.copyOut() "

	if nStrFmtSpecDigitsSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix)

		return newDigitsSepsDto, err
	}

	newDigitsSepsDto.decimalSeparator =
		nStrFmtSpecDigitsSepsDto.decimalSeparator

	newDigitsSepsDto.integerDigitsSeparator =
		nStrFmtSpecDigitsSepsDto.integerDigitsSeparator

	lenIntDigitsGroupingSequence :=
		len(nStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		newDigitsSepsDto.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
		return newDigitsSepsDto, err
	}

	newDigitsSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence)

	for i := 0; i < lenIntDigitsGroupingSequence; i++ {
		newDigitsSepsDto.integerDigitsGroupingSequence[i] =
			nStrFmtSpecDigitsSepsDto.integerDigitsGroupingSequence[i]
	}

	newDigitsSepsDto.lock = new(sync.Mutex)

	return newDigitsSepsDto, err
}

// setDigitsSeps - Transfers new data to an instance of NumStrFmtSpecDigitsSeparatorsDto.
// When complete all the data fields within input parameter 'nStrFmtSpecDigitsSepDto'
// will be overwritten.
//
func (nStrFmtSpecDigitsSepsQuark *numStrFmtSpecDigitsSeparatorsDtoQuark) setDigitsSeps(
	nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto,
	decimalSeparator rune,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix string) (
	err error) {

	if nStrFmtSpecDigitsSepsQuark.lock == nil {
		nStrFmtSpecDigitsSepsQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark.lock.Lock()

	defer nStrFmtSpecDigitsSepsQuark.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "numStrFmtSpecDigitsSeparatorsDtoQuark.setDigitsSeps() "

	if nStrFmtSpecDigitsSepDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepDto' is invalid!\n"+
			"'nStrFmtSpecDigitsSepDto' is a 'nil' pointer\n",
			ePrefix)

		return err
	}

	nStrFmtSpecDigitsSepDto.decimalSeparator =
		decimalSeparator

	nStrFmtSpecDigitsSepDto.integerDigitsSeparator =
		integerDigitsSeparator

	nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence =
		make([]uint, 0, 10)

	lenIntDigitsGroupingSequence :=
		len(integerDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		return err
	}

	nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence, 10)

	for i := 0; i < lenIntDigitsGroupingSequence; i++ {
		nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence[i] =
			integerDigitsGroupingSequence[i]
	}

	return err
}

// testValidityOfNumSepsDto - Receives an instance of
// NumStrFmtSpecDigitsSeparatorsDto and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (nStrFmtSpecDigitsSepsQuark *numStrFmtSpecDigitsSeparatorsDtoQuark) testValidityOfNumSepsDto(
	nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtSpecDigitsSepsQuark.lock == nil {
		nStrFmtSpecDigitsSepsQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark.lock.Lock()

	defer nStrFmtSpecDigitsSepsQuark.lock.Unlock()

	ePrefix += "numStrFmtSpecDigitsSeparatorsDtoQuark.testValidityOfNumSepsDto() "

	isValid = false

	if nStrFmtSpecDigitsSepDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepDto' is invalid!\n"+
			"'nStrFmtSpecDigitsSepDto' is a 'nil' pointer\n",
			ePrefix)

		return isValid, err
	}

	if nStrFmtSpecDigitsSepDto.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' is empty and has a zero value.\n",
			ePrefix)

		return isValid, err
	}

	if nStrFmtSpecDigitsSepDto.integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'integerDigitsSeparator' is invalid!\n"+
			"'integerDigitsSeparator' is empty and has a zero value.\n",
			ePrefix)

		return isValid, err
	}

	if len(nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'integerDigitsGroupingSequence' is invalid!\n"+
			"'integerDigitsGroupingSequence' is empty and has no member elements.\n",
			ePrefix)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
