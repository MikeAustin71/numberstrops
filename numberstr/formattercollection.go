package numberstr

import (
	"sync"
)

// FormatterCollection - Encapsulates all of the format
// specifications required to format the numeric values contained
// in type NumStrDto. The primary purpose of this type is to
// manage a collection of INumStrFormatter objects which provide
// formatting specifications for numeric values displayed in text
// strings. Member data elements are defined as follows as follows:
//
//  type FormatterCollection struct {
//   fmtCollection []INumStrFormatter
//  }
//
//
//  fmtCollection       []INumStrFormatter
//     - A collection of INumStrFormatter objects. Current
//       supported INumStrFormatter types include:
//          1. FormatterAbsoluteValue
//          2. FormatterBinary
//          3. FormatterCountry
//          4. FormatterCurrency
//          5. FormatterHexadecimal
//          6. FormatterOctal
//          7. FormatterSciNotation // Scientific Notation
//          8. FormatterSignedNumber
//
// An instance of FormatterCollection is maintained as member
// data variable by type NumStrDto.
//
type FormatterCollection struct {
	description   string
	tag           string
	fmtCollection []INumStrFormatter
	lock          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// FormatterCollection ('incomingFmtCollection') to the
// corresponding data fields of the current FormatterCollection
// instance.
//
// If 'incomingFmtCollection' is judged to be invalid, this method
// will return an error.
//
// Be advised, all of the data fields in the current
// FormatterCollection instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFmtCollection      *FormatterCollection
//     - A pointer to an instance of FormatterCollection. The data
//       values in this object will be copied to corresponding data
//       elements in the current FormatterCollection instance.
//
//
//  ePrefix                    *ErrPrefixDto
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
func (fmtCollection *FormatterCollection) CopyIn(
	incomingFmtCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) error {

	if fmtCollection.lock == nil {
		fmtCollection.lock = new(sync.Mutex)
	}

	fmtCollection.lock.Lock()

	defer fmtCollection.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCollection." +
			"CopyIn()")

	return formatterCollectionAtom{}.ptr().
		copyIn(
			fmtCollection,
			incomingFmtCollection,
			ePrefix.XCtx("incomingFmtCollection->"+
				"fmtCollection"))
}

// Empty - Deletes and resets data values for all INumStrFormatter
// collection elements in the current FormatterCollection instance
// to their initial 'zero' values.
//
func (fmtCollection *FormatterCollection) Empty() {

	if fmtCollection.lock == nil {
		fmtCollection.lock = new(sync.Mutex)
	}

	fmtCollection.lock.Lock()

	_ = formatterCollectionQuark{}.ptr().
		empty(
			fmtCollection,
			nil)

	fmtCollection.lock.Unlock()

	fmtCollection.lock = nil
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterCollection instance to determine whether that instance
// is valid in all respects.
//
// The diagnostic analysis is performed on each element in the
// INumStrFormatter collection.
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
//     - If this method completes successfully, the returned
//       boolean value will signal whether the current
//       FormatterCollection is valid, or not. If the current
//       FormatterCollection contains valid data, this method
//       returns 'true'. If the data is invalid, this method
//       returns 'false'.
//
func (fmtCollection *FormatterCollection) IsValidInstance() (
	isValid bool) {

	if fmtCollection.lock == nil {
		fmtCollection.lock = new(sync.Mutex)
	}

	fmtCollection.lock.Lock()

	defer fmtCollection.lock.Unlock()

	isValid,
		_ = formatterCollectionQuark{}.ptr().
		testValidityOfFormatterCollection(
			fmtCollection,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterCollection instance to determine whether that
// instance is valid in all respects.
//
// The diagnostic analysis is performed on each element in the
// INumStrFormatter collection.
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
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//
// -----------------------------------------------------------------
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
func (fmtCollection *FormatterCollection) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtCollection.lock == nil {
		fmtCollection.lock = new(sync.Mutex)
	}

	fmtCollection.lock.Lock()

	defer fmtCollection.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCollection." +
			"IsValidInstanceError()")

	_,
		err := formatterCollectionQuark{}.ptr().
		testValidityOfFormatterCollection(
			fmtCollection,
			ePrefix.XCtx(
				"Testing Validity of 'fmtCollection' "))

	return err
}
