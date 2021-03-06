package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorsDtoQuark struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numStrIntSeparatorsDtoQuark.
//
func (intSepsDtoQuark numStrIntSeparatorsDtoQuark) ptr() *numStrIntSeparatorsDtoQuark {

	if intSepsDtoQuark.lock == nil {
		intSepsDtoQuark.lock = new(sync.Mutex)
	}

	intSepsDtoQuark.lock.Lock()

	defer intSepsDtoQuark.lock.Unlock()

	newQuark := new(numStrIntSeparatorsDtoQuark)

	newQuark.lock = new(sync.Mutex)

	return newQuark
}

// testValidityOfNumStrIntSepsDto - Receives an instance of
// NumStrIntSeparatorsDto and proceeds to test the validity of the
// member data fields.
//
// If one or more data elements are found to be invalid, an
// error is returned and the return boolean parameter, 'isValid',
// is set to 'false'.
//
func (intSepsDtoQuark *numStrIntSeparatorsDtoQuark) testValidityOfNumStrIntSepsDto(
	intSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if intSepsDtoQuark.lock == nil {
		intSepsDtoQuark.lock = new(sync.Mutex)
	}

	intSepsDtoQuark.lock.Lock()

	defer intSepsDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorsDtoQuark." +
			"testValidityOfNumStrIntSepsDto()")

	isValid = false

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer\n",
			ePrefix.String())

		return isValid, err
	}

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSeps := len(intSepsDto.intSeparators)

	if lenIntSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Integer Separators array is empty!\n"+
			"'intSepsDto.intSeparators' is a zero length array.\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenIntSeps; i++ {
		err =
			intSepsDto.intSeparators[i].
				IsValidInstanceError(ePrefix.XCtx(
					fmt.Sprintf("intSepsDto.intSeparators[%v]",
						i)))

		if err != nil {
			return isValid, err
		}
	}

	isValid = true

	return isValid, err
}
