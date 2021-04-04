package numberstr

import (
	"fmt"
	"sync"
)

type formatterCollectionMechanics struct {
	lock *sync.Mutex
}

// addReplaceCollection - Receives a new INumStrFormatter type and
// either adds this type to the FormatterCollection or replaces
// an existing element of the same type with this new
// INumStrFormatter element.
//
func (fmtCollectionMech *formatterCollectionMechanics) addReplaceCollection(
	formatterCollection *FormatterCollection,
	newCollectionElement INumStrFormatter,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCollectionMech.lock == nil {
		fmtCollectionMech.lock = new(sync.Mutex)
	}

	fmtCollectionMech.lock.Lock()

	defer fmtCollectionMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionMechanics." +
			"addReplaceCollection()")

	if formatterCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCollection' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if newCollectionElement == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newCollectionElement' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	err = newCollectionElement.IsValidInstanceError(
		ePrefix.XCtx(
			"newCollectionElement"))

	if err != nil {
		return err
	}

	collectionType := newCollectionElement.GetNumStrFormatTypeCode()

	if !collectionType.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Collection 'NumStrFormatTypeCode' is invalid!\n"+
			"Collection 'NumStrFormatTypeCode' integer value='%v'\n",
			ePrefix.XCtxEmpty().String(),
			collectionType.XValueInt())
		return err
	}

	arrayLen := len(formatterCollection.fmtCollection)

	if arrayLen == 0 {

		formatterCollection.fmtCollection = nil

		formatterCollection.fmtCollection =
			append(formatterCollection.fmtCollection,
				newCollectionElement)

		return err
	}

	var elementType NumStrFormatTypeCode

	for i := 0; i < arrayLen; i++ {

		elementType =
			formatterCollection.fmtCollection[i].
				GetNumStrFormatTypeCode()

		if elementType == collectionType {
			err =
				formatterCollection.fmtCollection[i].
					CopyInINumStrFormatter(
						newCollectionElement,
						ePrefix.XCtx(
							"newCollectionElement"))

			return err
		}
	}

	formatterCollection.fmtCollection =
		append(formatterCollection.fmtCollection,
			newCollectionElement)

	return err
}

// isFormatterInCollection - Receives a formatter collection and a
// formatter type. If the formatter type exists in the collection,
// this method returns 'true'.
//
// If the specified formatter type does NOT exist in the
// collection, this method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCollection *FormatterCollection
//     - A pointer to an instance of FormatterCollection. The
//       Formatter Collection houses an array of INumStrFormatter
//       objects used in formatting different types of number
//       strings.
//
//
//  formatterType       NumStrFormatTypeCode
//     - NumStrFormatTypeCode is an enumeration of valid formatter
//       types. If a formatter matching this 'formatterType' code
//       is located in the formatter collection of input parameter,
//       'formatterCollection', the return parameter,
//       'formatterDoesExist' will be set to true.
//
//       If the specified INumStrFormatter formatter does not exist
//       in the formatter collection, return parameter
//       'formatterDoesExist' will be set to false.
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
//       collection. If set to true, it signals that the correct
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
func (fmtCollectionMech *formatterCollectionMechanics) isFormatterInCollection(
	formatterCollection *FormatterCollection,
	formatterType NumStrFormatTypeCode,
	ePrefix *ErrPrefixDto) (
	formatterDoesExist bool,
	err error) {

	if fmtCollectionMech.lock == nil {
		fmtCollectionMech.lock = new(sync.Mutex)
	}

	fmtCollectionMech.lock.Lock()

	defer fmtCollectionMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionMechanics." +
			"isFormatterInCollection()")

	formatterDoesExist = false

	if formatterCollection == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCollection' is a 'nil' pointer!\n",
			ePrefix.String())

		return formatterDoesExist, err
	}

	if !formatterType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterType' is invalid!\n"+
			"Integer formatterType value='%v'\n"+
			"String formatterType value='%v'\n",
			ePrefix.String(),
			formatterType.XValueInt(),
			formatterType.String())

		return formatterDoesExist, err
	}

	lenFmtCol := len(formatterCollection.fmtCollection)

	if lenFmtCol == 0 {

		return formatterDoesExist, err
	}

	var actFmtType NumStrFormatTypeCode

	for i := 0; i < lenFmtCol; i++ {
		actFmtType =
			formatterCollection.fmtCollection[i].
				GetNumStrFormatTypeCode()

		if actFmtType == formatterType {

			formatterDoesExist = true

			return formatterDoesExist, err
		}

	}

	return formatterDoesExist, err
}

// getFormatter - Receives a formatter type and proceeds to extract
// and return of the matching formatter object from the
// FormatterCollection.
//
// The formatter collection is an array of type INumStrFormatter
// which houses all the currently configured formatter objects.
//
// If the INumStrFormatter formatter specified by input parameter
// 'formatterType' does not exist in the formatter collection,
// return parameter 'formatterDoesExist' will be set to false and
// an error will be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCollection *FormatterCollection
//     - A pointer to an instance of FormatterCollection. The
//       Formatter Collection houses an array of INumStrFormatter
//       objects used in formatting different types of number
//       strings.
//
//
//  formatterType       NumStrFormatTypeCode
//     - NumStrFormatTypeCode is an enumeration of valid formatter
//       types. A copy of the formatter housed in the Formatter
//       Collection and matching this formatter type code will be
//       extracted and returned to the calling function.
//
//       If the specified INumStrFormatter formatter does not exist
//       in the formatter collection, return parameter
//       'formatterDoesExist' will be set to false and an error will
//       be returned.
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
//       collection. If set to true, it signals that the correct
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
func (fmtCollectionMech *formatterCollectionMechanics) getFormatter(
	formatterCollection *FormatterCollection,
	formatterType NumStrFormatTypeCode,
	ePrefix *ErrPrefixDto) (
	formatterDoesExist bool,
	formatter INumStrFormatter,
	err error) {

	if fmtCollectionMech.lock == nil {
		fmtCollectionMech.lock = new(sync.Mutex)
	}

	fmtCollectionMech.lock.Lock()

	defer fmtCollectionMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionMechanics." +
			"getFormatter()")

	formatterDoesExist = false

	if formatterCollection == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCollection' is a 'nil' pointer!\n",
			ePrefix.String())

		return formatterDoesExist, formatter, err
	}

	if !formatterType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterType' is invalid!\n"+
			"Integer formatterType value='%v'\n"+
			"String formatterType value='%v'\n",
			ePrefix.String(),
			formatterType.XValueInt(),
			formatterType.String())

		return formatterDoesExist, formatter, err
	}

	lenFmtCol := len(formatterCollection.fmtCollection)

	if lenFmtCol == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Formatter Collection is empty!\n"+
			"len(formatterCollection.fmtCollection)==0\n",
			ePrefix.String())

		return formatterDoesExist, formatter, err
	}

	var actFmtType NumStrFormatTypeCode

	for i := 0; i < lenFmtCol; i++ {
		actFmtType =
			formatterCollection.fmtCollection[i].
				GetNumStrFormatTypeCode()

		if actFmtType == formatterType {
			formatter,
				err = formatterCollection.fmtCollection[i].
				CopyOutINumStrFormatter(
					ePrefix.XCtx(
						"formatterCollection." +
							"fmtCollection[i]"))

			if err != nil {
				formatterDoesExist = true
			}

			return formatterDoesExist, formatter, err
		}
	}

	err = fmt.Errorf("%v\n"+
		"Error: Could NOT located Formatter Type '%v'\n"+
		"in formatter collection!\n",
		ePrefix.XCtxEmpty().String(),
		formatterType.String())

	return formatterDoesExist, formatter, err
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionMechanics.
//
func (fmtCollectionMech formatterCollectionMechanics) ptr() *formatterCollectionMechanics {

	if fmtCollectionMech.lock == nil {
		fmtCollectionMech.lock = new(sync.Mutex)
	}

	fmtCollectionMech.lock.Lock()

	defer fmtCollectionMech.lock.Unlock()

	newFmtCollectionMechanics :=
		new(formatterCollectionMechanics)

	newFmtCollectionMechanics.lock = new(sync.Mutex)

	return newFmtCollectionMechanics
}
