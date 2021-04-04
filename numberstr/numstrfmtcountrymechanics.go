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

// addBaseFormatters - Adds Binary, Hexadecimal, Octal and
// Scientific Notation default formatters to a
// FormatterCollection instance.
//
func (numStrFmtCntryMech *numStrFormatCountryMechanics) addBaseFormatters(
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

	err = fmtCollection.AddReplaceCollectionElement(
		FormatterOctal{}.NewUnitedStatesDefaultsPtr(),
		ePrefix.XCtx(
			"Octal"))

	if err != nil {
		return err
	}

	var fmtSciNotation FormatterSciNotation

	fmtSciNotation,
		err = FormatterSciNotation{}.NewWithDefaults(
		false,                  // significandUsesLeadingPlus
		6,                      // mantissaLength
		'E',                    // exponentChar
		true,                   // exponentUsesLeadingPlus,
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"Scientific Notation"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&fmtSciNotation,
		ePrefix.XCtx(
			"Scientific Notation"))

	return err
}
