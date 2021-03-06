package numberstr

import (
	"sync"
)

// NumStrIntSeparator is designed to be used as an array
// element in an array which will be used to insert integer
// separators, primarily thousands separators, into number
// strings. Some countries/cultures do not use thousands
// separation and instead rely on multiple integer separation
// characters and grouping sequences for a single number
// string. One notable example of this found in the 'Indian
// Number System'.
//  Reference:
//  https://en.wikipedia.org/wiki/Indian_numbering_system
//
// An array of NumStrIntSeparator elements provides the flexibility
// necessary to process these complex number formats.
//
type NumStrIntSeparator struct {
	intSeparatorChar        rune
	intSeparatorGrouping    uint
	intSeparatorRepetitions uint
	lock                    *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrIntSeparator instance to the data fields of the current
// NumStrIntSeparator instance.
//
// If input parameter 'incomingNStrIntSeparator' is judged to be
// invalid, this method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrIntSeparator instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNStrIntSeparator     *NumStrIntSeparator
//     - A pointer to an instance of NumStrIntSeparator.
//       The data values in this object will be copied to the
//       current NumStrIntSeparator instance.
//
//       If input parameter 'incomingNStrIntSeparator' is judged
//       to be invalid, this method will return an error.
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
func (nStrIntSep *NumStrIntSeparator) CopyIn(
	incomingNStrIntSeparator *NumStrIntSeparator,
	ePrefix *ErrPrefixDto) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrIntSeparator.CopyIn()")

	return numStrIntSeparatorMolecule{}.Ptr().
		copyIn(
			nStrIntSep,
			incomingNStrIntSeparator,
			ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrIntSeparator instance.
//
// If the current NumStrIntSeparator instance is judged to be
// invalid, this method will return an error.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrIntSeparator
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparator will be created and returned containing
//       all of the data values copied from the current instance of
//       NumStrIntSeparator.
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
func (nStrIntSep *NumStrIntSeparator) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparator,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrIntSeparator.CopyOut()")
	return numStrIntSeparatorMolecule{}.Ptr().
		copyOut(
			nStrIntSep,
			ePrefix)
}

// Equal - Receives an incoming NumStrIntSeparator
// instance and compares it the current NumStrIntSeparator
// instance. If the two objects have equal data values, this method
// returns 'true'
//
func (nStrIntSep *NumStrIntSeparator) Equal(
	nStrIntSep2 NumStrIntSeparator) bool {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	return numStrIntSeparatorMolecule{}.Ptr().
		equal(nStrIntSep, &nStrIntSep2)
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrIntSeparator instance to determine whether the current
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
//       current NumStrIntSeparator is valid, or not. If the
//       current NumStrIntSeparator contains valid data, this
//       method returns 'true'. If the data is invalid, this method
//       will return 'false'.
//
func (nStrIntSep *NumStrIntSeparator) IsValidInstance() (
	isValid bool) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	isValid,
		_ =
		numStrIntSeparatorQuark{}.Ptr().
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrIntSeparator instance to determine whether the
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
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the current instance of NumStrIntSeparator contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nStrIntSep *NumStrIntSeparator) IsValidInstanceError(
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx(
		"NumStrIntSeparator.IsValidInstanceError()",
		"Testing Validity of 'nStrIntSep'")
	_,
		err =
		numStrIntSeparatorQuark{}.Ptr().
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				ePrefix)

	return err
}
