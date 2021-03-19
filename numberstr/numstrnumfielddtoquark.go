package numberstr

import (
	"fmt"
	"sync"
)

type numStrNumFieldDtoQuark struct {
	lock *sync.Mutex
}

func (nStrNFldDtoQuark *numStrNumFieldDtoQuark) empty(
	numFieldDto *NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrNFldDtoQuark.lock == nil {
		nStrNFldDtoQuark.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark.lock.Lock()

	defer nStrNFldDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrNumFieldDtoQuark.empty()")

	if numFieldDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'numFieldDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	numFieldDto.requestedNumFieldLength = 0

	numFieldDto.actualNumFieldLength = 0

	numFieldDto.minimumNumFieldLength = 0

	numFieldDto.textJustifyFormat =
		TextJustify(0).None()

	return err
}

// ptr - Returns a pointer to a new instance of
// numStrNumFieldDtoQuark.
//
func (nStrNFldDtoQuark numStrNumFieldDtoQuark) ptr() *numStrNumFieldDtoQuark {

	if nStrNFldDtoQuark.lock == nil {
		nStrNFldDtoQuark.lock = new(sync.Mutex)
	}

	nStrNFldDtoQuark.lock.Lock()

	defer nStrNFldDtoQuark.lock.Unlock()

	newNStrNumFldDtoQuark :=
		new(numStrNumFieldDtoQuark)

	newNStrNumFldDtoQuark.lock = new(sync.Mutex)

	return newNStrNumFldDtoQuark
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

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

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
