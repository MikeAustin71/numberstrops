package numberstr

import (
	"fmt"
	"sync"
)

type numStrBasicQuark struct {
	lock *sync.Mutex
}

// getCountryFormatters - Populates a FormatterCollection
// collection with a series of number string formatter objects
// associated with the country or culture specified by input
// parameter, 'countryCulture'.
//
func (nStrBasicQuark numStrBasicQuark) getCountryFormatters(
	fmtCollection *FormatterCollection,
	countryCulture CountryCultureId,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrBasicQuark.lock == nil {
		nStrBasicQuark.lock = new(sync.Mutex)
	}

	nStrBasicQuark.lock.Lock()

	defer nStrBasicQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numStrBasicQuark." +
			"getCountryFormatters()")

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if !countryCulture.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'countryCulture' is invalid!\n"+
			"'countryCulture' string value='%v'\n"+
			"'countryCulture' integer value='%v'\n",
			ePrefix.String(),
			countryCulture.String(),
			countryCulture.XValueInt())

		return err
	}

	fmtCntry := NumStrFormatCountry{}

	switch countryCulture {

	case CountryCultureId(0).Albania():
		err = fmtCntry.Albania(
			fmtCollection,
			ePrefix.XCtx("Albania"))

	case CountryCultureId(0).Argentina():
		err = fmtCntry.Argentina(
			fmtCollection,
			ePrefix.XCtx("Argentina"))

	case CountryCultureId(0).Australia():
		err = fmtCntry.Australia(
			fmtCollection,
			ePrefix.XCtx("Australia"))

	case CountryCultureId(0).Austria():
		err = fmtCntry.Austria(
			fmtCollection,
			ePrefix.XCtx("Austria"))

	case CountryCultureId(0).Bahrain():
		err = fmtCntry.Bahrain(
			fmtCollection,
			ePrefix.XCtx("Bahrain"))

	case CountryCultureId(0).Bangladesh():
		err = fmtCntry.Bangladesh(
			fmtCollection,
			ePrefix.XCtx("Bangladesh"))

	default:
		err = fmt.Errorf("%v\n"+
			"Error: Unsupported CountryCultureId!\n"+
			"'countryCulture' string value='%v'\n"+
			"'countryCulture' integer value='%v'\n",
			ePrefix.String(),
			countryCulture.String(),
			countryCulture.XValueInt())
	}

	return err
}

// Ptr - Returns a pointer to a new instance of numStrBasicQuark.
//
func (nStrBasicQuark numStrBasicQuark) ptr() *numStrBasicQuark {

	if nStrBasicQuark.lock == nil {
		nStrBasicQuark.lock = new(sync.Mutex)
	}

	nStrBasicQuark.lock.Lock()

	defer nStrBasicQuark.lock.Unlock()

	return &numStrBasicQuark{
		lock: new(sync.Mutex),
	}
}
