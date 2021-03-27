package numberstr

import (
	"fmt"
	"sync"
)

type formatterCountryUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCountryUtility.
//
func (fmtCountryUtil formatterCountryUtility) ptr() *formatterCountryUtility {

	if fmtCountryUtil.lock == nil {
		fmtCountryUtil.lock = new(sync.Mutex)
	}

	fmtCountryUtil.lock.Lock()

	defer fmtCountryUtil.lock.Unlock()

	newFormatterCountryUtility := new(formatterCountryUtility)

	newFormatterCountryUtility.lock = new(sync.Mutex)

	return newFormatterCountryUtility
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterCountry instance
// to United States Default values.
//
// For the United States, the country specification parameters are
// defined as follows:
//
//   CountryIdNo            = 840
//   CountryIdString        = "840"
//   CountryDescription     = "Country Setup - United States"
//   CountryTag             = ""
//   CountryCultureName     = "United States"
//   CountryAbbreviatedName = "USA"
//   CountryAlternateNames  = []string{
//                            "The United States of America",
//                            "United States of America",
//                            "America"}
//   CountryCodeTwoChar     = "US"
//   CountryCodeThreeChar   = "USA"
//   CountryCodeNumber      = "840"
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCountry              *FormatterCountry
//     - A pointer to an instance of FormatterCountry.
//       All data values in this object will be overwritten and
//       set to United States default values for Country
//       specification codes.
//
//
//  ePrefix                       *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                           error
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
func (fmtCountryUtil *formatterCountryUtility) setToUnitedStatesDefaults(
	formatterCountry *FormatterCountry,
	ePrefix *ErrPrefixDto) error {

	if fmtCountryUtil.lock == nil {
		fmtCountryUtil.lock = new(sync.Mutex)
	}

	fmtCountryUtil.lock.Lock()

	defer fmtCountryUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCountryUtility." +
			"setToUnitedStatesDefaults()")

	if formatterCountry == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCountry' is invalid!\n"+
			"'formatterCountry' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	return formatterCountryMechanics{}.ptr().
		setWithComponents(
			formatterCountry,
			840,
			"840",
			"Country Setup - United States",
			"",
			"United States",
			"USA",
			[]string{
				"The United States of America",
				"United States of America",
				"America"},
			"US",
			"USA",
			"840",
			ePrefix)
}
