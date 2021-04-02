package numberstr

import (
	"fmt"
	"sync"
)

type formatterCollectionQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values of all member
// variables within a FormatterCollection instance to their
// initial 'zero' values.
//
func (fmtCollectionQuark *formatterCollectionQuark) empty(
	formatterCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) (
	err error) {
	if fmtCollectionQuark.lock == nil {
		fmtCollectionQuark.lock = new(sync.Mutex)
	}

	fmtCollectionQuark.lock.Lock()

	defer fmtCollectionQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionQuark." +
			"empty()")

	if formatterCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCollection' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	arrayLen := len(formatterCollection.fmtCollection)

	if arrayLen == 0 {
		return err
	}

	for i := 0; i < arrayLen; i++ {
		formatterCollection.fmtCollection[i].Empty()
	}

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterCollectionQuark.
//
func (fmtCollectionQuark formatterCollectionQuark) ptr() *formatterCollectionQuark {

	if fmtCollectionQuark.lock == nil {
		fmtCollectionQuark.lock = new(sync.Mutex)
	}

	fmtCollectionQuark.lock.Lock()

	defer fmtCollectionQuark.lock.Unlock()

	newFmtCollectionQuark := new(formatterCollectionQuark)

	newFmtCollectionQuark.lock = new(sync.Mutex)

	return newFmtCollectionQuark
}

// testValidityOfFormatterCollection - Receives an instance of
// FormatterBinary and proceeds to test the validity of the member
// data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (fmtCollectionQuark *formatterCollectionQuark) testValidityOfFormatterCollection(
	formatterCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	fmtCollectionQuark.lock.Lock()

	defer fmtCollectionQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCollectionQuark." +
			"testValidityOfFormatterCollection()")

	isValid = false

	if formatterCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCollection' is a 'nil' pointer!\n",
			ePrefix.String())
		return isValid, err
	}

	arrayLen := len(formatterCollection.fmtCollection)

	if arrayLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Formatter Collection of INumStrFormatter objects\n"+
			"is a zero length array!\n",
			ePrefix.XCtx(
				"formatterCollection.fmtCollection"))
		return isValid, err
	}

	for i := 0; i < arrayLen; i++ {
		err = formatterCollection.fmtCollection[i].IsValidInstanceError(
			ePrefix.XCtx(
				fmt.Sprintf(
					"formatterCollection.fmtCollection[%v]",
					i)))

		if err != nil {
			return isValid, err
		}
	}

	isValid = true

	return isValid, err
}
