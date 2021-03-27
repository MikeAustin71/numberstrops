package numberstr

import (
	"sync"
)

// FormatterCountry - This structure stores country/culture data
// for associated number string formatting operations.
//
// A Formatter Country object consists of the following data fields.
//
//  numStrFmtType          NumStrFormatTypeCode
//     - An enumeration value automatically set to:
//           NumStrFormatTypeCode(0).CountryCulture()
//
//
//  idNo                   uint64
//     - An arbitrary Id Number
//
//
//  idString               string
//     - An arbitrary Id String
//
//
//  description            string
//     - A user defined description
//
//
//  tag                    string
//     - A user defined tag
//
//  countryCultureName     string
//     - Usually contains the official country name
//
//  abbreviatedCountryName string
//     - The abbreviated country name
//
//
//  alternateCountryNames  []string
//     - A list of alternate country names
//
//
//  countryCodeTwoChar     string
//     - The ISO 3166-1 alpha-2 two character country code.
//       Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  countryCodeThreeChar   string
//     - The ISO 3166-1 alpha-3 three character country code.
//       Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  countryCodeNumber      string
//     - ISO 3166-1 numeric – three-digit country codes which are identical
//       to those developed and maintained by the United Nations Statistics
//       Division.  Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
type FormatterCountry struct {
	numStrFmtType          NumStrFormatTypeCode // An enum set to NumStrFormatTypeCode(0).CountryCulture()
	idNo                   uint64               // An arbitrary Id Number
	idString               string               // An arbitrary Id String
	description            string               // A user defined description
	tag                    string               // A user defined tag
	countryCultureName     string               // Usually contains the official country name
	abbreviatedCountryName string               // The abbreviated country/culture name
	alternateCountryNames  []string             // A list of alternate country names
	countryCodeTwoChar     string               // ISO 3166-1 alpha-2 two character country code
	countryCodeThreeChar   string               // ISO 3166-1 alpha-3 three character country code
	countryCodeNumber      string               // ISO 3166-1 numeric – three-digit country code

	lock *sync.Mutex
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterCountry instance to their
// initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (fmtCountry *FormatterCountry) Empty() {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	_ = formatterCountryQuark{}.ptr().
		empty(fmtCountry,
			nil)

	fmtCountry.lock.Unlock()

	fmtCountry.lock = nil

	return
}

// CopyIn - Copies the data fields from an incoming
// FormatterCountry instance  to the data fields
// of the current instance of FormatterCountry
// instance.
//
func (fmtCountry *FormatterCountry) CopyIn(
	incomingFormatterCountry *FormatterCountry,
	ePrefix *ErrPrefixDto) error {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("FormatterCountry.CopyIn()")

	return formatterCountryElectron{}.ptr().copyIn(
		fmtCountry,
		incomingFormatterCountry,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current
// FormatterCountry instance.
//
func (fmtCountry *FormatterCountry) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterCountry,
	error) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("FormatterCountry.CopyOut()")

	return formatterCountryElectron{}.ptr().
		copyOut(
			fmtCountry,
			ePrefix.XCtx(
				"fmtCountry->"))
}

// GetAbbreviatedCountryName - Returns the value of member variable
// 'abbreviatedCountryName'.
func (fmtCountry *FormatterCountry) GetAbbreviatedCountryName() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.abbreviatedCountryName
}

// GetAlternateCountryNames - Returns the value of member variable
// 'alternateCountryNames'.
func (fmtCountry *FormatterCountry) GetAlternateCountryNames() []string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	var altCountryNames []string

	lenAltCntryNames :=
		len(fmtCountry.alternateCountryNames)

	if lenAltCntryNames == 0 {
		altCountryNames = make([]string, 0, 10)
	} else {

		altCountryNames = make([]string, lenAltCntryNames, 10)

		for i := 0; i < lenAltCntryNames; i++ {
			altCountryNames[i] =
				fmtCountry.alternateCountryNames[i]
		}
	}

	return altCountryNames
}

// GetCountryCodeNumber - Returns the value of member variable
// 'countryCodeThreeChar'.
func (fmtCountry *FormatterCountry) GetCountryCodeNumber() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.countryCodeNumber
}

// GetCountryCodeThreeChar - Returns the value of member variable
// 'countryCodeThreeChar'.
func (fmtCountry *FormatterCountry) GetCountryCodeThreeChar() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.countryCodeThreeChar
}

// GetCountryCodeTwoChar - Returns the value of member variable
// 'countryCodeTwoChar'.
func (fmtCountry *FormatterCountry) GetCountryCodeTwoChar() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.countryCodeTwoChar
}

// GetCountryCultureName - Returns the value of member variable
// 'countryCultureName'.
func (fmtCountry *FormatterCountry) GetCountryCultureName() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.countryCultureName
}

// GetDescription - Returns the value of member variable 'description'.
func (fmtCountry *FormatterCountry) GetDescription() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.description
}

// GetIdNo - Returns the value of member variable 'idNo'.
func (fmtCountry *FormatterCountry) GetIdNo() uint64 {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.idNo
}

// GetIdString - Returns the value of member variable 'idString'.
func (fmtCountry *FormatterCountry) GetIdString() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.idString
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for FormatterCountry
// objects is NumStrFormatTypeCode(0).CountryCulture().
//
// This method is required by interface INumStrFormatter.
//
func (fmtCountry *FormatterCountry) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.numStrFmtType
}

// GetTag - Returns the value of member variable 'tag'.
func (fmtCountry *FormatterCountry) GetTag() string {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	return fmtCountry.tag
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterCountry instance to determine whether
// the current instance is valid in all respects.
//
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
//     - This returned boolean value will signal whether the current
//       FormatterCountry is valid, or not. If the current
//       FormatterCountry contains valid data, this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
func (fmtCountry *FormatterCountry) IsValidInstance() bool {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	isValid,
		_ := formatterCountryQuark{}.ptr().
		testValidityOfFormatterCountry(
			fmtCountry,
			new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// FormatterCountry instance to determine whether the current
// instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix             *ErrPrefixDto
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
//       If this instance of FormatterCountry contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (fmtCountry *FormatterCountry) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCountry.IsValidInstanceError()")

	_,
		err := formatterCountryQuark{}.ptr().
		testValidityOfFormatterCountry(
			fmtCountry,
			ePrefix)

	return err
}

// NewUnitedStatesDefaults - Creates and returns a new instance of
// FormatterCountry. This method specifies the United States
// default values for Country/Culture.
//
// United States Country/Culture default parameters are defined as
// follows:
//
//    numStrFmtType         =
//          NumStrFormatTypeCode(0).CountryCulture()
//
//    idNo                   = 840
//    idString               = "840"
//    description            = "Country Setup - United States"
//    tag                    = ""
//    countryCultureName     = "United States"
//    abbreviatedCountryName = "USA"
//    alternateCountryNames  =
//          "The United States of America"
//          "United States of America"
//          "America"
//
//    countryCodeTwoChar     = "US"
//    countryCodeThreeChar   = "USA"
//    countryCodeNumber      = "840"
//
// The member variable 'FormatterCountry.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).CountryCulture()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  -- None --
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  FormatterCountry
//     - This parameter will return a new, populated instance of
//       FormatterCountry configured United States default
//       Country/Culture parameters.
//
func (fmtCountry FormatterCountry) NewUnitedStatesDefaults() FormatterCountry {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	newFormatterCountry := FormatterCountry{}

	_ = formatterCountryUtility{}.ptr().
		setToUnitedStatesDefaults(
			&newFormatterCountry,
			nil)

	return newFormatterCountry
}

// NewWithComponents - Creates and returns a new instance of NumericSeparators.
// This type encapsulates the digit separators used in formatting a
// number string for text display.
//
// The member variable 'FormatterCountry.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).CountryCulture()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  idNo                   uint64
//     - An arbitrary Id Number
//
//
//  idString               string
//     - An arbitrary Id String
//
//
//  description            string
//     - A user defined description
//
//
//  tag                    string
//     - A user defined tag
//
//  countryCultureName     string
//     - Usually contains the official country name
//
//  abbreviatedCountryName string
//     - The abbreviated country name
//
//
//  alternateCountryNames  []string
//     - A list of alternate country names
//
//
//  countryCodeTwoChar     string
//     - The ISO 3166-1 alpha-2 two character country code.
//       Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  countryCodeThreeChar   string
//     - The ISO 3166-1 alpha-3 three character country code.
//       Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  countryCodeNumber      string
//     - ISO 3166-1 numeric – three-digit country codes which are identical
//       to those developed and maintained by the United Nations Statistics
//       Division.  Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  ePrefix             *ErrPrefixDto
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
//  FormatterCountry
//     - If this method completes successfully, new instance of
//       FormatterCountry will be created and
//       returned.
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
func (fmtCountry FormatterCountry) NewWithComponents(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCultureName string,
	abbreviatedCountryName string,
	alternateCountryNames []string,
	countryCodeTwoChar string,
	countryCodeThreeChar string,
	countryCodeNumber string,
	ePrefix *ErrPrefixDto) (
	FormatterCountry,
	error) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCountry.NewWithComponents()")

	newFormatterCountry := FormatterCountry{}

	err :=
		formatterCountryMechanics{}.ptr().setWithComponents(
			&newFormatterCountry,
			idNo,
			idString,
			description,
			tag,
			countryCultureName,
			abbreviatedCountryName,
			alternateCountryNames,
			countryCodeTwoChar,
			countryCodeThreeChar,
			countryCodeNumber,
			ePrefix)

	return newFormatterCountry, err
}

// SetAbbreviatedCountryName - Sets the value of member variable
// 'abbreviatedCountryName'.
func (fmtCountry *FormatterCountry) SetAbbreviatedCountryName(
	abbreviatedCountryName string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.abbreviatedCountryName =
		abbreviatedCountryName
}

// SetAlternateCountryNames - Sets the value of member variable
// 'alternateCountryNames'.
func (fmtCountry *FormatterCountry) SetAlternateCountryNames(
	alternateCountryNames []string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	lenAltCntryNames :=
		len(alternateCountryNames)

	if lenAltCntryNames == 0 {

		fmtCountry.alternateCountryNames =
			make([]string, 0, 10)

	} else {

		fmtCountry.alternateCountryNames = make([]string, lenAltCntryNames, 10)

		for i := 0; i < lenAltCntryNames; i++ {
			fmtCountry.alternateCountryNames[i] =
				alternateCountryNames[i]
		}
	}

}

// SetCountryCodeNumber - Sets the value of member variable
// 'countryCodeNumber'.
//
// ISO 3166-1 numeric – three-digit country codes which are identical
// to those developed and maintained by the United Nations Statistics
// Division.  Reference:
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (fmtCountry *FormatterCountry) SetCountryCodeNumber(
	countryCodeNumber string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.countryCodeNumber =
		countryCodeNumber
}

// SetCountryCodeThreeChar - Sets the value of member variable
// 'countryCodeThreeChar'.
//
// The ISO 3166-1 alpha-3 three character country code.
//  Reference:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (fmtCountry *FormatterCountry) SetCountryCodeThreeChar(
	countryCodeThreeChar string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.countryCodeThreeChar =
		countryCodeThreeChar
}

// SetCountryCodeTwoChar - Sets the value of member variable
// 'countryCodeTwoChar'.
//
// The ISO 3166-1 alpha-2 two character country code.
// Reference:
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (fmtCountry *FormatterCountry) SetCountryCodeTwoChar(
	countryCodeTwoChar string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.countryCodeTwoChar =
		countryCodeTwoChar

}

// SetCountryCultureName - Sets the value of member variable
// 'countryCultureName'.
//
// Official name of the country.
//
func (fmtCountry *FormatterCountry) SetCountryCultureName(
	countryCultureName string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.countryCultureName =
		countryCultureName
}

// SetDescription - Sets the value of member variable 'description'.
// This is a user defined description.
//
func (fmtCountry *FormatterCountry) SetDescription(
	description string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.description =
		description
}

// SetIdNo - Sets the value of member variable 'idNo'.
func (fmtCountry *FormatterCountry) SetIdNo(
	idNo uint64) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.idNo = idNo
}

// SetIdString - Sets the value of member variable 'idString'.
func (fmtCountry *FormatterCountry) SetIdString(
	idString string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.idString = idString
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterCountry object. For Country/Culture
// formatters, the Number String Format Type Code is set to
// NumStrFormatTypeCode(0).CountryCulture().
//
// This method is required by interface INumStrFormatter.
//
func (fmtCountry *FormatterCountry) SetNumStrFormatTypeCode() {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.numStrFmtType =
		NumStrFormatTypeCode(0).CountryCulture()
}

// SetTag - Sets the value of member variable 'tag'.
// This is a user defined description.
//
func (fmtCountry *FormatterCountry) SetTag(
	tag string) {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	fmtCountry.tag = tag
}

// SetToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterCountry instance
// to United States Default values.
//
// United States Country/Culture default parameters are defined as
// follows:
//
//    numStrFmtType         =
//          NumStrFormatTypeCode(0).CountryCulture()
//
//    idNo                   = 840
//    idString               = "840"
//    description            = "Country Setup - United States"
//    tag                    = ""
//    countryCultureName     = "United States"
//    abbreviatedCountryName = "USA"
//    alternateCountryNames  =
//          "The United States of America"
//          "United States of America"
//          "America"
//
//    countryCodeTwoChar     = "US"
//    countryCodeThreeChar   = "USA"
//    countryCodeNumber      = "840"
//
// The member variable 'FormatterCountry.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).CountryCulture()
//
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterCountry instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
func (fmtCountry *FormatterCountry) SetToUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) error {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCountry." +
			"SetToUnitedStatesDefaults()")

	return formatterCountryUtility{}.ptr().
		setToUnitedStatesDefaults(
			fmtCountry,
			ePrefix)
}

// SetToUnitedStatesDefaultsIfEmpty - If the current
// FormatterCountry instance is empty or invalid,
// this method will set the member variable data values to United
// States default values.
//
// If the current FormatterCountry instance is valid
// and NOT empty, this method will take no action and exit.
//
// United States Country/Culture default parameters are defined as
// follows:
//
//    numStrFmtType         =
//          NumStrFormatTypeCode(0).CountryCulture()
//
//    idNo                   = 840
//    idString               = "840"
//    description            = "Country Setup - United States"
//    tag                    = ""
//    countryCultureName     = "United States"
//    abbreviatedCountryName = "USA"
//    alternateCountryNames  =
//          "The United States of America"
//          "United States of America"
//          "America"
//
//    countryCodeTwoChar     = "US"
//    countryCodeThreeChar   = "USA"
//    countryCodeNumber      = "840"
//
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterCountry instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
func (fmtCountry *FormatterCountry) SetToUnitedStatesDefaultsIfEmpty(
	ePrefix *ErrPrefixDto) error {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCountry." +
			"SetToUnitedStatesDefaults()")

	isValid,
		_ := formatterCountryQuark{}.ptr().
		testValidityOfFormatterCountry(
			fmtCountry,
			ePrefix)

	if isValid {
		return nil
	}

	return formatterCountryUtility{}.ptr().
		setToUnitedStatesDefaults(
			fmtCountry,
			ePrefix)
}

// SetWithComponents - Sets the data values for current
// FormatterCountry instance.
//
// The member variable 'FormatterCountry.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).CountryCulture()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  idNo                   uint64
//     - An arbitrary Id Number
//
//
//  idString               string
//     - An arbitrary Id String
//
//
//  description            string
//     - A user defined description
//
//
//  tag                    string
//     - A user defined tag
//
//  countryCultureName     string
//     - Usually contains the official country name
//
//  abbreviatedCountryName string
//     - The abbreviated country name
//
//
//  alternateCountryNames  []string
//     - A list of alternate country names
//
//
//  countryCodeTwoChar     string
//     - The ISO 3166-1 alpha-2 two character country code.
//       Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  countryCodeThreeChar   string
//     - The ISO 3166-1 alpha-3 three character country code.
//       Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  countryCodeNumber      string
//     - ISO 3166-1 numeric – three-digit country codes which are identical
//       to those developed and maintained by the United Nations Statistics
//       Division.  Reference:
//        https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
//
//  ePrefix             *ErrPrefixDto
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
func (fmtCountry *FormatterCountry) SetWithComponents(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCultureName string,
	abbreviatedCountryName string,
	alternateCountryNames []string,
	countryCodeTwoChar string,
	countryCodeThreeChar string,
	countryCodeNumber string,
	ePrefix *ErrPrefixDto) error {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterCountry.SetWithComponents()")

	return formatterCountryMechanics{}.ptr().
		setWithComponents(
			fmtCountry,
			idNo,
			idString,
			description,
			tag,
			countryCultureName,
			abbreviatedCountryName,
			alternateCountryNames,
			countryCodeTwoChar,
			countryCodeThreeChar,
			countryCodeNumber,
			ePrefix)
}
