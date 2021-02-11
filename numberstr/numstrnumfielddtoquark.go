package numberstr

import (
	"fmt"
	"sync"
)

type numStrNumFieldDtoQuark struct {
	lock *sync.Mutex
}

// testValidityNumberFieldDto - Performs validity tests on an
// instance of NumberFieldDto. If the instance is found to be
// invalid, return parameter 'isValid' is set to false and
// 'err' is configured with an appropriate error message.
//
// If the NumberFieldDto instance qualifies as valid, 'isValid'
// is set to 'true' and 'err' is set to 'nil'.
//
func (nStrNFldDtoQuark *numStrNumFieldDtoQuark) testValidityNumberFieldDto(
	numFieldDto *NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrNFldDtoQuark.lock == nil {
		nStrNFldDtoQuark.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark.lock.Lock()

	defer nStrNFldDtoQuark.lock.Unlock()

	ePrefix.SetEPref("numStrNumFieldDtoQuark.testValidityNumberFieldDto()")
	isValid = false

	if numFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if numFieldDto.lock == nil {
		numFieldDto.lock = new(sync.Mutex)
	}

	if numFieldDto.requestedNumFieldLength == 0 {
		err = fmt.Errorf("%v"+
			"Error: numFieldDto.requestedNumFieldLength == 0\n",
			ePrefix.String())

		return isValid, err
	}

	if numFieldDto.requestedNumFieldLength < -1 {
		numFieldDto.requestedNumFieldLength = -1
	}

	if numFieldDto.actualNumFieldLength < -1 {
		numFieldDto.requestedNumFieldLength = -1
	}

	if numFieldDto.minimumNumFieldLength < -1 {
		numFieldDto.requestedNumFieldLength = -1
	}

	if !numFieldDto.textJustifyFormat.XIsValid() {
		err = fmt.Errorf("%v"+
			"Error: numFieldDto.textJustifyFormat is invalid\n"+
			"numFieldDto.textJustifyFormat integer value=='%v'\n",
			ePrefix.String(),
			numFieldDto.textJustifyFormat.XValueInt())

		return isValid, err
	}

	isValid = true
	return isValid, err
}
