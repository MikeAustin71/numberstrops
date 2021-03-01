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
	targetNumSepsDto *NumericSeparatorsDto,
	inComingNumSepsDto *NumericSeparatorsDto,
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

	lenIntSeparators :=
		len(inComingNumSepsDto.integerSeparators)

	targetNumSepsDto.integerSeparators =
		make([]NumStrIntSeparator,
			lenIntSeparators,
			10)

	for i := 0; i < lenIntSeparators; i++ {

		err =
			targetNumSepsDto.integerSeparators[i].CopyIn(
				&inComingNumSepsDto.integerSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf("Copying inComingNumSepsDto.integerSeparators[%v]",
						i)))

		if err != nil {
			return err
		}

	}

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'nStrFmtSpecDigitsSepsDto' styled as a new instance
// of NumericSeparatorsDto.
//
// If 'nStrFmtSpecDigitsSepsDto' is judged to be invalid, this
// method will return an error.
//
func (numSepsDtoElectron *numericSeparatorDtoElectron) copyOut(
	numSepsDto *NumericSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	newNumSepsDto NumericSeparatorsDto,
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

	lenIntSeparators :=
		len(numSepsDto.integerSeparators)

	newNumSepsDto.integerSeparators =
		make([]NumStrIntSeparator, lenIntSeparators)

	for i := 0; i < lenIntSeparators; i++ {
		err =
			newNumSepsDto.integerSeparators[i].CopyIn(
				&numSepsDto.integerSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf(
						"Copying numSepsDto.integerSeparators[%v]",
						i)))

		if err != nil {
			return newNumSepsDto, err
		}

	}

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
