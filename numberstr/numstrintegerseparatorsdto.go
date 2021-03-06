package numberstr

import "sync"

// NumStrIntSeparatorsDto is used to transport and manage an array
// of NumStrIntSeparator. These integer separators are in turn
// used to control the grouping and separation characters
// employed in separating digits within an integer number.
//
//  United States Example:        1,000,000,000
//  European Example:             1 000 000 000
//  Indian Number System Example: 6,78,90,00,00,00,00,000
//
type NumStrIntSeparatorsDto struct {
	intSeparators []NumStrIntSeparator
	lock          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrIntSeparatorsDto instance  to the data fields
// of the current instance of NumStrIntSeparatorsDto
// instance.
//
// If input parameter 'incomingIntSepsDto' is judged to be
// invalid, this method will return an error.
//
// IMPORTANT
//
// Be advised, all of the data fields in the current
// NumStrIntSeparatorsDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingIntSepsDto  *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto.
//       The data values in this object will be copied to the
//       current NumStrIntSeparatorsDto instance.
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
//  error
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) CopyIn(
	incomingIntSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto.CopyIn()")

	return numStrIntSeparatorsDtoElectron{}.ptr().
		copyIn(
			intSeparatorsDto,
			incomingIntSepsDto,
			ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrIntSeparatorsDto instance.
//
// If the current NumStrIntSeparatorsDto instance is judged
// to be invalid, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  NumStrIntSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparatorsDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrIntSeparatorsDto.
//
//
//  error
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto.CopyOut()")

	return numStrIntSeparatorsDtoElectron{}.ptr().
		copyOut(
			intSeparatorsDto,
			ePrefix)
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrIntSeparatorsDto instance to determine whether the current
// instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - This returned boolean value will signal whether the
//       current NumStrIntSeparatorsDto is valid, or not. If the
//       current NumStrIntSeparatorsDto contains valid data, this
//       method returns 'true'. If the data is invalid, this method
//       will return 'false'.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) IsValidInstance() (
	isValid bool) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	isValid,
		_ = numStrIntSeparatorsDtoQuark{}.ptr().
		testValidityOfNumStrIntSepsDto(
			intSeparatorsDto,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrIntSeparatorsDto instance to determine whether the
// current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current instance of NumStrIntSeparatorsDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"IsValidInstanceError()")

	_,
		err := numStrIntSeparatorsDtoQuark{}.ptr().
		testValidityOfNumStrIntSepsDto(
			intSeparatorsDto,
			nil)

	return err
}
