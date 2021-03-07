package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecCountryDtoUtility struct {
	lock *sync.Mutex
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming NumStrFmtSpecCountryDto instance
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
//  nStrCountryDto                *NumStrFmtSpecCountryDto
//     - A pointer to an instance of NumStrFmtSpecCountryDto.
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
func (nStrFmtCntryUtil *numStrFmtSpecCountryDtoUtility) setToUnitedStatesDefaults(
	nStrCountryDto *NumStrFmtSpecCountryDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtCntryUtil.lock == nil {
		nStrFmtCntryUtil.lock = new(sync.Mutex)
	}

	nStrFmtCntryUtil.lock.Lock()

	defer nStrFmtCntryUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecCountryDtoUtility." +
			"setToUnitedStatesDefaults()")

	if nStrCountryDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrCountryDto' is invalid!\n"+
			"'nStrCountryDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	nStrFmtSpecCntryMech :=
		numStrFmtSpecCountryDtoMechanics{}

	return nStrFmtSpecCntryMech.setWithComponents(
		nStrCountryDto,
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
