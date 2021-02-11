package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDigitsSeparatorsDtoQuark struct {
	lock *sync.Mutex
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
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecDigitsSepsQuark.lock == nil {
		nStrFmtSpecDigitsSepsQuark.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepsQuark.lock.Lock()

	defer nStrFmtSpecDigitsSepsQuark.lock.Unlock()

	ePrefix.SetEPref("numStrFmtSpecDigitsSeparatorsDtoQuark.testValidityOfNumSepsDto()")

	isValid = false

	if nStrFmtSpecDigitsSepDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDigitsSepDto' is invalid!\n"+
			"'nStrFmtSpecDigitsSepDto' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrFmtSpecDigitsSepDto.decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'decimalSeparator' is invalid!\n"+
			"'decimalSeparator' is empty and has a zero value.\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrFmtSpecDigitsSepDto.integerDigitsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'integerDigitsSeparator' is invalid!\n"+
			"'integerDigitsSeparator' is empty and has a zero value.\n",
			ePrefix.String())

		return isValid, err
	}

	if len(nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member varialbe 'integerDigitsGroupingSequence' is invalid!\n"+
			"'integerDigitsGroupingSequence' is empty and has no member elements.\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true
	return isValid, err
}
