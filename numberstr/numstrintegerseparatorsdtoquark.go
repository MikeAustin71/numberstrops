package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorsDtoQuark struct {
	lock *sync.Mutex
}

// isEmpty - Returns a boolean flag signaling whether the incoming
// NumStrIntSeparatorsDto is populated with data.
//
// If the incoming instance, 'NumStrIntSeparatorsDto', is invalid,
// not populated with data or simply empty, this method will return
// 'true'.
//
// If the incoming instance, 'NumStrIntSeparatorsDto', contains
// data, this method will return false.
//
func (intSepsDtoQuark *numStrIntSeparatorsDtoQuark) isEmpty(
	intSepsDto *NumStrIntSeparatorsDto) bool {

	if intSepsDtoQuark.lock == nil {
		intSepsDtoQuark.lock = new(sync.Mutex)
	}

	intSepsDtoQuark.lock.Lock()

	defer intSepsDtoQuark.lock.Unlock()

	if intSepsDto == nil {
		return true
	}

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	if len(intSepsDto.intSeparators) == 0 {
		return true
	}

	return false
}

// isEqual - Receives an incoming NumStrIntSeparatorsDto
// instance and compares it to a second NumStrIntSeparatorsDto
// instance. If the two objects have equal data values, this method
// returns 'true'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSepDto01     *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. The
//       data values in this object will be compared to those
//       contained in input parameter 'nStrIntSepDto02' to
//       determine equivalency.
//
//
//  nStrIntSepDto02     *NumStrIntSeparatorsDto
//     - A pointer to a second instance of NumStrIntSeparatorsDto.
//       The data values in this object will be compared to those
//       contained in input parameter 'nStrIntSepDto01' to
//       determine equivalency.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isEqual             bool
//     - If the data values contained in input parameters
//       'nStrIntSepDto01' and 'nStrIntSepDto02' are equivalent,
//       this boolean return value will be set to 'true'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (intSepsDtoQuark *numStrIntSeparatorsDtoQuark) isEqual(
	nStrIntSepDto01 *NumStrIntSeparatorsDto,
	nStrIntSepDto02 *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	isEqual bool,
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
			"isEqual()")

	isEqual = false

	if nStrIntSepDto01 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSepDto01' is invalid!\n"+
			"'nStrIntSepDto01' is a 'nil' pointer\n",
			ePrefix.String())

		return isEqual, err
	}

	if nStrIntSepDto01.intSeparators == nil {
		nStrIntSepDto01.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	if nStrIntSepDto02 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSepDto02' is invalid!\n"+
			"'nStrIntSepDto02' is a 'nil' pointer\n",
			ePrefix.String())

		return isEqual, err
	}

	if nStrIntSepDto02.intSeparators == nil {
		nStrIntSepDto02.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSeps01 := len(nStrIntSepDto01.intSeparators)

	if lenIntSeps01 != len(nStrIntSepDto02.intSeparators) {
		return isEqual, err
	}

	for i := 0; i < lenIntSeps01; i++ {

		isEqual = nStrIntSepDto01.intSeparators[i].Equal(
			nStrIntSepDto02.intSeparators[i])

		if !isEqual {
			return isEqual, err
		}

	}

	return isEqual, err
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

	lastIdx := lenIntSeps - 1

	for i := 0; i < lenIntSeps; i++ {

		if i < lastIdx &&
			intSepsDto.intSeparators[i].restartIntGroupingSequence == true {

			err = fmt.Errorf("%v\n"+
				"Error: NumStrIntSeparator is not the last array element.\n"+
				"However, NumStrIntSeparator.restartIntGroupingSequence=='true'\n"+
				"i='%v'\n",
				ePrefix.XCtxEmpty(),
				i)

			return isValid, err
		}

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
