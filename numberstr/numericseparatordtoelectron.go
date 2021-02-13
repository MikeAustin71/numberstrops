package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'inComingNumSepsDto' to input parameter 'targetNumSepsDto'.
//
// If 'inComingNumSepsDto' is judged to be invalid, this method
// will return an error.
//
// This method will OVERWRITE the data fields of the
// 'targetNumSepsDto' instance.
//
func (numSepsDtoElectron *numericSeparatorDtoElectron) copyIn(
	targetNumSepsDto *NumericSeparatorDto,
	inComingNumSepsDto *NumericSeparatorDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsDtoElectron.lock == nil {
		numSepsDtoElectron.lock = new(sync.Mutex)
	}

	numSepsDtoElectron.lock.Lock()

	defer numSepsDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numericSeparatorDtoElectron.copyIn()")

	if targetNumSepsDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'targetNumSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if inComingNumSepsDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'inComingNumSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetNumSepsDto.lock == nil {
		targetNumSepsDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark :=
		numericSeparatorDtoQuark{}

	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
			inComingNumSepsDto,
			ePrefix.XCtx("Testing validity of 'inComingNumSepsDto'"))

	if err != nil {
		return err
	}

	targetNumSepsDto.decimalSeparator =
		inComingNumSepsDto.decimalSeparator

	targetNumSepsDto.integerDigitsSeparator =
		inComingNumSepsDto.integerDigitsSeparator

	lenIntDigitsGroupSeq :=
		len(inComingNumSepsDto.integerDigitsGroupingSequence)

	targetNumSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupSeq, lenIntDigitsGroupSeq+5)

	if lenIntDigitsGroupSeq == 0 {
		return
	}

	_ =
		copy(targetNumSepsDto.integerDigitsGroupingSequence,
			inComingNumSepsDto.integerDigitsGroupingSequence)

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecDigitsSepsDto' styled as a new instance
// of NumericSeparatorDto.
//
func (numSepsDtoElectron *numericSeparatorDtoElectron) copyOut(
	numSepsDto *NumericSeparatorDto,
	ePrefix *ErrPrefixDto) (
	newNumSepsDto NumericSeparatorDto,
	err error) {

	if numSepsDtoElectron.lock == nil {
		numSepsDtoElectron.lock = new(sync.Mutex)
	}

	numSepsDtoElectron.lock.Lock()

	defer numSepsDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numericSeparatorDtoElectron.copyOut()")

	if numSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNumSepsDto, err
	}

	numSepsDtoQuark :=
		numericSeparatorDtoQuark{}

	_,
		err =
		numSepsDtoQuark.testValidityOfNumSepsDto(
			numSepsDto,
			ePrefix.XCtx("Testing validity of 'numSepsDto'"))

	if err != nil {
		return newNumSepsDto, err
	}

	newNumSepsDto.decimalSeparator =
		numSepsDto.decimalSeparator

	newNumSepsDto.integerDigitsSeparator =
		numSepsDto.integerDigitsSeparator

	lenIntDigitsGroupingSequence :=
		len(numSepsDto.integerDigitsGroupingSequence)

	newNumSepsDto.integerDigitsGroupingSequence =
		make([]uint, lenIntDigitsGroupingSequence)

	if lenIntDigitsGroupingSequence == 0 {
		return newNumSepsDto, err
	}

	_ = copy(
		newNumSepsDto.integerDigitsGroupingSequence,
		numSepsDto.integerDigitsGroupingSequence)

	newNumSepsDto.lock = new(sync.Mutex)

	return newNumSepsDto, err
}

// ptr - Returns a pointer to a new instance of
// numericSeparatorDtoElectron.
//
func (numSepsDtoElectron numericSeparatorDtoElectron) ptr() *numericSeparatorDtoElectron {

	if numSepsDtoElectron.lock == nil {
		numSepsDtoElectron.lock = new(sync.Mutex)
	}

	numSepsDtoElectron.lock.Lock()

	defer numSepsDtoElectron.lock.Unlock()

	newNumSepsElectron := new(numericSeparatorDtoElectron)

	newNumSepsElectron.lock = new(sync.Mutex)

	return newNumSepsElectron
}
