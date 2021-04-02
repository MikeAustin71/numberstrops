package numberstr

import (
	"fmt"
	"sync"
)

type formatterCollectionAtom struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingFmtCollection' to input parameter
// 'targetFmtCollection'.
//
// Be advised - All data fields in 'targetFmtCollection' will be
// overwritten.
//
// If the 'incomingFmtCollection' is judged to be invalid, this
// method will return an error.
//
func (fmtCollectionAtom *formatterCollectionAtom) copyIn(
	targetFmtCollection *FormatterCollection,
	incomingFmtCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCollectionAtom.lock == nil {
		fmtCollectionAtom.lock = new(sync.Mutex)
	}

	fmtCollectionAtom.lock.Lock()

	defer fmtCollectionAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionAtom." +
			"copyIn()")

	if targetFmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFmtCollection' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if incomingFmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingFmtCollection' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err = formatterCollectionQuark{}.ptr().
		testValidityOfFormatterCollection(
			incomingFmtCollection,
			ePrefix.XCtx(
				"incomingFmtCollection"))

	if err != nil {
		return err
	}

	err = formatterCollectionQuark{}.ptr().
		empty(
			targetFmtCollection,
			ePrefix.XCtx(
				"targetFmtCollection"))

	if err != nil {
		return err
	}

	if targetFmtCollection.lock == nil {
		targetFmtCollection.lock = new(sync.Mutex)
	}

	if incomingFmtCollection.fmtCollection == nil {

		targetFmtCollection.fmtCollection = nil

		return err
	}

	arrayLen := len(incomingFmtCollection.fmtCollection)

	if arrayLen == 0 {

		targetFmtCollection.fmtCollection =
			make([]INumStrFormatter, 0)

		return err
	}

	targetFmtCollection.fmtCollection = nil

	var newINumStrFormatterObject INumStrFormatter

	for i := 0; i < arrayLen; i++ {

		newINumStrFormatterObject,
			err = incomingFmtCollection.fmtCollection[i].
			CopyOutINumStrFormatter(
				ePrefix.XCtx(
					fmt.Sprintf(
						"incomingFmtCollection.fmtCollection[%v]",
						i)))

		if err != nil {
			return err
		}

		targetFmtCollection.fmtCollection =
			append(targetFmtCollection.fmtCollection,
				newINumStrFormatterObject)
	}

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'formatterCollection' styled as a new instance of
// FormatterCollection.
//
// If input parameter 'formatterCollection' is judged to be
// invalid, this method will return an error.
//
func (fmtCollectionAtom *formatterCollectionAtom) copyOut(
	formatterCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) (
	newFmtCollection FormatterCollection,
	err error) {

	if fmtCollectionAtom.lock == nil {
		fmtCollectionAtom.lock = new(sync.Mutex)
	}

	fmtCollectionAtom.lock.Lock()

	defer fmtCollectionAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionAtom." +
			"copyOut()")

	if formatterCollection == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCollection' is a 'nil' pointer!\n",
			ePrefix.String())

		return newFmtCollection, err
	}

	_,
		err = formatterCollectionQuark{}.ptr().
		testValidityOfFormatterCollection(
			formatterCollection,
			ePrefix.XCtx(
				"formatterCollection"))

	if err != nil {
		return newFmtCollection, err
	}

	newFmtCollection.lock = new(sync.Mutex)

	if formatterCollection.fmtCollection == nil {

		newFmtCollection.fmtCollection = nil

		return newFmtCollection, err
	}

	arrayLen := len(formatterCollection.fmtCollection)

	if arrayLen == 0 {

		newFmtCollection.fmtCollection =
			make([]INumStrFormatter, 0)

		return newFmtCollection, err
	}

	newFmtCollection.fmtCollection = nil

	var newINumStrFormatterObject INumStrFormatter

	for i := 0; i < arrayLen; i++ {

		newINumStrFormatterObject,
			err = formatterCollection.fmtCollection[i].
			CopyOutINumStrFormatter(
				ePrefix.XCtx(
					fmt.Sprintf(
						"formatterCollection.fmtCollection[%v]",
						i)))

		if err != nil {
			return newFmtCollection, err
		}

		newFmtCollection.fmtCollection =
			append(newFmtCollection.fmtCollection,
				newINumStrFormatterObject)
	}

	return newFmtCollection, err
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionAtom.
//
func (fmtCollectionAtom formatterCollectionAtom) ptr() *formatterCollectionAtom {

	if fmtCollectionAtom.lock == nil {
		fmtCollectionAtom.lock = new(sync.Mutex)
	}

	fmtCollectionAtom.lock.Lock()

	defer fmtCollectionAtom.lock.Unlock()

	newFmtCollectionAtom := new(formatterCollectionAtom)

	newFmtCollectionAtom.lock = new(sync.Mutex)

	return newFmtCollectionAtom
}
