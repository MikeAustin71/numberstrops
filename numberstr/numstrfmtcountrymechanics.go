package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatCountryMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numStrFormatCountryMechanics.
//
func (numStrFmtCntryMech numStrFormatCountryMechanics) ptr() *numStrFormatCountryMechanics {

	if numStrFmtCntryMech.lock == nil {
		numStrFmtCntryMech.lock = new(sync.Mutex)
	}

	numStrFmtCntryMech.lock.Lock()

	defer numStrFmtCntryMech.lock.Unlock()

	newNumStrFmtCntryMech :=
		new(numStrFormatCountryMechanics)

	newNumStrFmtCntryMech.lock = new(sync.Mutex)

	return newNumStrFmtCntryMech
}

// addBaseFormatters - Adds Binary, Hexadecimal and Octal default
// formatters to a FormatterCollection instance.
//
func (numStrFmtCntryMech *numStrFormatCountryMechanics) addBaseFormatters(
	fmtCountry *NumStrFormatCountry,
	fmtCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numStrFmtCntryMech.lock == nil {
		numStrFmtCntryMech.lock = new(sync.Mutex)
	}

	numStrFmtCntryMech.lock.Lock()

	defer numStrFmtCntryMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numStrFormatCountryMechanics." +
			"addBaseFormatters()")

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		FormatterBinary{}.NewUnitedStatesDefaultsPtr(),
		ePrefix.XCtx(
			"Binary"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		FormatterHexadecimal{}.NewUnitedStatesDefaultsPtr(),
		ePrefix.XCtx(
			"Hexadecimal"))

	if err != nil {
		return err
	}

}
