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
