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

// AddReplaceCollectionElement - Receives a new INumStrFormatter
// type and either adds this type to the existing
// FormatterCollection or replaces an pre-existing element
// of the same type with this new INumStrFormatter element.
//
// For example, if a new currency formatter (FormatterCurrency)
// object is submitted as an input parameter, the existing
// collection of formatter objects will be searched to determine
// if another currency formatter object previously exists in the
// formatter collection. If a currency object previously exists in
// the collection it will be deleted and replaced by new currency
// formatter object supplied in the input parameter.
//
// If, on the the other hand, a currency formatter does NOT
// previously exist in the collection, the new currency formatter
// object will simply be added to the current collection.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
// newCollectionElement INumStrFormatter
//     - A new INumStrFormatter to be added to the current
//       collection of formatter objects maintained by the current
//       instance of FormatterCollection. If a formatter object
//       of the same type previously exists in the collection, it
//       will be deleted an replaced by a deep copy of this
//       parameter, 'newCollectionElement'.
//
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
func (fmtCollection *FormatterCollection) AddReplaceCollectionElement(
	newCollectionElement INumStrFormatter,
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
			"AddReplaceCollectionElement()")

	return formatterCollectionMechanics{}.ptr().
		addReplaceCollection(
			fmtCollection,
			newCollectionElement,
			ePrefix.XCtx(
				"fmtCollection"))
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

// CopyOut - Returns a deep copy of the current FormatterCollection
// instance.
//
// If the current FormatterCollection instance is judged to be
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
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  FormatterCollection
//     - If this method completes successfully, a new instance of
//       FormatterCollection will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterCollection.
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
func (fmtCollection *FormatterCollection) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterCollection,
	error) {

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
			"CopyOut()")

	return formatterCollectionAtom{}.ptr().copyOut(
		fmtCollection,
		ePrefix.XCtx(
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

// GetFormatter - Receives a formatter type and proceeds to extract
// and return of the matching formatter object from the current
// FormatterCollection instance.
//
// The formatter collection is an array of type INumStrFormatter
// which houses all the currently configured formatter objects. The
// formatters are used to format numeric values within number
// strings.
//
// If the INumStrFormatter formatter specified by input parameter
// 'formatterType' does not exist in the formatter collection,
// return parameter 'formatterDoesExist' will be set to 'false' and
// an error will be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterType       NumStrFormatTypeCode
//     - NumStrFormatTypeCode is an enumeration of valid formatter
//       types. A copy of the formatter housed in the Formatter
//       Collection and matching this formatter type code will be
//       extracted and returned to the calling function.
//
//       If the specified INumStrFormatter formatter does not exist
//       in the formatter collection, return parameter
//       'formatterDoesExist' will be set to 'false' and an error
//       will be returned.
//
//       If the 'formatterType' code is invalid, this method will
//       return an error.
//
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  formatterDoesExist  bool
//     - A boolean flag signaling whether the formatter specified
//       by input parameter 'formatterType' exists in the formatter
//       collection for the current FormatterCollection instance.
//       If set to true, this flag signals that the correct
//       formatter was located in the formatter collection and
//       returned by parameter, 'formatter'.
//
//
//  formatter           INumStrFormatter
//     - If this method completes successfully and a formatter
//       matching the the formatter type specified by input
//       parameter 'formatterType' was located in the formatter
//       collection, an INumStrFormatter formatter object will be
//       returned.
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
func (fmtCollection *FormatterCollection) GetFormatter(
	formatterType NumStrFormatTypeCode,
	ePrefix *ErrPrefixDto) (
	formatterDoesExist bool,
	formatter INumStrFormatter,
	err error) {

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
			"GetFormatter()")

	formatterDoesExist,
		formatter,
		err = formatterCollectionMechanics{}.ptr().
		getFormatter(
			fmtCollection,
			formatterType,
			ePrefix.XCtx(
				"fmtCollection"))

	return formatterDoesExist, formatter, err
}

// FormatterIsInCollection - Receives a formatter type and proceeds
// to determine if that type exists in the formatter collection for
// the current FormatterCollection instance. If the formatter type
// is located in the collection, this method returns 'true'.
//
// If the specified formatter type does NOT exist in the formatter
// collection, this method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterType       NumStrFormatTypeCode
//     - NumStrFormatTypeCode is an enumeration of valid formatter
//       types. If a formatter matching this 'formatterType' code
//       is located in the formatter collection of the current
//       FormatterCollection instance, the return parameter,
//       'formatterDoesExist' will be set to true.
//
//       If the specified INumStrFormatter formatter does not exist
//       in the formatter collection, return parameter
//       'formatterDoesExist' will be set to 'false'.
//
//       If the 'formatterType' code is invalid, this method will
//       return an error.
//
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  formatterDoesExist  bool
//     - A boolean flag signaling whether the formatter specified
//       by input parameter 'formatterType' exists in the formatter
//       collection. If set to 'true', it signals that the correct
//       formatter was located in the formatter collection.
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
func (fmtCollection *FormatterCollection) FormatterIsInCollection(
	formatterType NumStrFormatTypeCode,
	ePrefix *ErrPrefixDto) (
	formatterDoesExist bool,
	err error) {

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
			"FormatterIsInCollection()")

	formatterDoesExist,
		err = formatterCollectionMechanics{}.ptr().
		isFormatterInCollection(
			fmtCollection,
			formatterType,
			ePrefix.XCtx(
				"fmtCollection"))

	return formatterDoesExist, err
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
