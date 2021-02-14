package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorDtoQuark struct {
	lock *sync.Mutex
}

// numStrSepDtosAreEqual - Returns 'true' if the two NumericSeparatorsDto
// objects submitted as input parameters have equal values.
//
func (numSepsDtoQuark *numericSeparatorDtoQuark) numStrSepDtosAreEqual(
	numSep1 *NumericSeparatorsDto,
	numSep2 *NumericSeparatorsDto) bool {

	if numSepsDtoQuark.lock == nil {
		numSepsDtoQuark.lock = new(sync.Mutex)
	}

	numSepsDtoQuark.lock.Lock()

	defer numSepsDtoQuark.lock.Unlock()

	if numSep1 == nil ||
		numSep2 == nil {
		return false
	}

	if numSep1.decimalSeparator !=
		numSep2.decimalSeparator {

		return false
	}

	if numSep1.integerDigitsSeparator !=
		numSep2.integerDigitsSeparator {

		return false
	}

	if numSep1.integerDigitsGroupingSequence == nil {
		numSep1.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
	}

	if numSep2.integerDigitsGroupingSequence == nil {
		numSep2.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
	}

	lenGrpSeq := len(numSep1.integerDigitsGroupingSequence)

	if lenGrpSeq !=
		len(numSep1.integerDigitsGroupingSequence) {
		return false
	}

	for i := 0; i < lenGrpSeq; i++ {
		if numSep1.integerDigitsGroupingSequence[i] !=
			numSep2.integerDigitsGroupingSequence[i] {
			return false
		}
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// numericSeparatorDtoQuark.
//
func (numSepsDtoQuark numericSeparatorDtoQuark) ptr() *numericSeparatorDtoQuark {

	if numSepsDtoQuark.lock == nil {
		numSepsDtoQuark.lock = new(sync.Mutex)
	}

	numSepsDtoQuark.lock.Lock()

	defer numSepsDtoQuark.lock.Unlock()

	newQuark := new(numericSeparatorDtoQuark)

	newQuark.lock = new(sync.Mutex)

	return newQuark
}

// testValidityOfNumSepsDto - Receives an instance of
// NumericSeparatorsDto and proceeds to test the
// validity of the member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (numSepsDtoQuark *numericSeparatorDtoQuark) testValidityOfNumSepsDto(
	numSepsDto *NumericSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if numSepsDtoQuark.lock == nil {
		numSepsDtoQuark.lock = new(sync.Mutex)
	}

	numSepsDtoQuark.lock.Lock()

	defer numSepsDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorDtoQuark." +
			"testValidityOfNumSepsDto()")

	isValid = false

	if numSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSepsDto' is invalid!\n"+
			"'numSepsDto' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if numSepsDto.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' is empty and has a zero value.\n",
			ePrefix.String())

		return isValid, err
	}

	if numSepsDto.integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'integerDigitsSeparator' is invalid!\n"+
			"'integerDigitsSeparator' is empty and has a zero value.\n",
			ePrefix.String())

		return isValid, err
	}

	if numSepsDto.integerDigitsGroupingSequence == nil {
		numSepsDto.integerDigitsGroupingSequence =
			make([]uint, 0, 10)
	}

	if len(numSepsDto.integerDigitsGroupingSequence) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'integerDigitsGroupingSequence' is invalid!\n"+
			"'integerDigitsGroupingSequence' is empty and has no member elements.\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true
	return isValid, err
}
