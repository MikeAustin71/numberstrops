package numberstr

import (
	"fmt"
	"sync"
)

type FormatterCountry struct {
	numStrFmtType          NumStrFormatTypeCode
	idNo                   uint64
	idString               string
	description            string
	tag                    string
	countryCultureName     string
	abbreviatedCountryName string
	alternateCountryNames  []string
	countryCodeTwoChar     string
	countryCodeThreeChar   string
	countryCodeNumber      string

	lock *sync.Mutex
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
	}

	ePrefix.SetEPref("FormatterCountry.CopyIn()")

	nStrFmtSpecCntryElectron :=
		formatterCountryElectron{}

	return nStrFmtSpecCntryElectron.copyIn(
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
	}

	ePrefix.SetEPref("FormatterCountry.CopyOut()")

	nStrFmtSpecCntryElectron :=
		formatterCountryElectron{}

	return nStrFmtSpecCntryElectron.copyOut(
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

	nStrFmtSpecCntryQuark :=
		formatterCountryQuark{}

	isValid,
		_ := nStrFmtSpecCntryQuark.testValidityOfFormatterCountry(
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
	}

	ePrefix.SetEPref(
		"FormatterCountry.IsValidInstanceError()")

	nStrFmtSpecCntryQuark := formatterCountryQuark{}

	_,
		err := nStrFmtSpecCntryQuark.testValidityOfFormatterCountry(
		fmtCountry,
		ePrefix)

	return err
}

// NewWithComponents - Creates and returns a new instance of NumericSeparators.
// This type encapsulates the digit separators used in formatting a
// number string for text display.
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
	}

	ePrefix.SetEPref(
		"FormatterCountry.NewWithComponents()")

	nStrFmtSpecCntryMech :=
		formatterCountryMechanics{}

	newCntryDto := FormatterCountry{}

	err :=
		nStrFmtSpecCntryMech.setWithComponents(
			&newCntryDto,
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

	return newCntryDto, err
}

// NewWithFmtSpecSetupDto - Creates and returns a new
// FormatterCountry instance based on input received from
// an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       FormatterCountry object. Only the following data
//       fields with a prefix of "Country" are used.
//
//       type NumStrFmtSpecSetupDto struct {
//         IdNo                                      uint64
//         IdString                                  string
//         Description                               string
//         Tag                                       string
//         CountryIdNo                               uint64
//         CountryIdString                           string
//         CountryDescription                        string
//         CountryTag                                string
//         CountryCultureName                        string
//         CountryAbbreviatedName                    string
//         CountryAlternateNames                     []string
//         CountryCodeTwoChar                        string
//         CountryCodeThreeChar                      string
//         CountryCodeNumber                         string
//         AbsoluteValFmt                            string
//         AbsoluteValTurnOnIntegerDigitsSeparation  bool
//         AbsoluteValNumSeps                        NumericSeparators
//         AbsoluteValNumField                       NumberFieldDto
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     uint
//         CurrencyCode                              string
//         CurrencyCodeNo                            string
//         CurrencyName                              string
//         CurrencySymbols                           []rune
//         MinorCurrencyName                         string
//         MinorCurrencySymbols                      []rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumSeps                           NumericSeparators
//         CurrencyNumField                          NumberFieldDto
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumSeps                       NumericSeparators
//         SignedNumValNumField                      NumberFieldDto
//         SciNotSignificandUsesLeadingPlus          bool
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//         SciNotNumFieldTextJustify                 TextJustify
//         Lock                                      *sync.Mutex
//       }
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
//     - If this method completes successfully, a new instance of
//       FormatterCountry will be returned to the caller.
//
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
func (fmtCountry FormatterCountry) NewWithFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
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
	}

	ePrefix.SetEPref(
		"FormatterCountry.NewWithFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return FormatterCountry{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryMech :=
		formatterCountryMechanics{}

	newCountryDto := FormatterCountry{}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	err := nStrFmtSpecCntryMech.setWithComponents(
		&newCountryDto,
		fmtSpecSetupDto.CountryIdNo,
		fmtSpecSetupDto.CountryIdString,
		fmtSpecSetupDto.CountryDescription,
		fmtSpecSetupDto.CountryTag,
		fmtSpecSetupDto.CountryCultureName,
		fmtSpecSetupDto.CountryAbbreviatedName,
		fmtSpecSetupDto.CountryAlternateNames,
		fmtSpecSetupDto.CountryCodeTwoChar,
		fmtSpecSetupDto.CountryCodeThreeChar,
		fmtSpecSetupDto.CountryCodeNumber,
		ePrefix)

	return newCountryDto, err
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

// SetWithSpecSetupDto - Sets the data values for current
// FormatterCountry instance based on input received from
// an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       FormatterCountry object. Only the following data
//       fields with a prefix of "Country" are used.
//
//       type NumStrFmtSpecSetupDto struct {
//         IdNo                                      uint64
//         IdString                                  string
//         Description                               string
//         Tag                                       string
//         CountryIdNo                               uint64
//         CountryIdString                           string
//         CountryDescription                        string
//         CountryTag                                string
//         CountryCultureName                        string
//         CountryAbbreviatedName                    string
//         CountryAlternateNames                     []string
//         CountryCodeTwoChar                        string
//         CountryCodeThreeChar                      string
//         CountryCodeNumber                         string
//         AbsoluteValFmt                            string
//         AbsoluteValTurnOnIntegerDigitsSeparation  bool
//         AbsoluteValNumSeps                        NumericSeparators
//         AbsoluteValNumField                       NumberFieldDto
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     uint
//         CurrencyCode                              string
//         CurrencyCodeNo                            string
//         CurrencyName                              string
//         CurrencySymbols                           []rune
//         MinorCurrencyName                         string
//         MinorCurrencySymbols                      []rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumSeps                           NumericSeparators
//         CurrencyNumField                          NumberFieldDto
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumSeps                       NumericSeparators
//         SignedNumValNumField                      NumberFieldDto
//         SciNotSignificandUsesLeadingPlus          bool
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//         SciNotNumFieldTextJustify                 TextJustify
//         Lock                                      *sync.Mutex
//       }
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
func (fmtCountry *FormatterCountry) SetWithSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if fmtCountry.lock == nil {
		fmtCountry.lock = new(sync.Mutex)
	}

	fmtCountry.lock.Lock()

	defer fmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterCountry." +
			"SetWithFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	fmtSpecSetupDto.Lock.Lock()

	defer fmtSpecSetupDto.Lock.Unlock()

	nStrFmtSpecCntryMech :=
		formatterCountryMechanics{}

	return nStrFmtSpecCntryMech.setWithComponents(
		fmtCountry,
		fmtSpecSetupDto.CountryIdNo,
		fmtSpecSetupDto.CountryIdString,
		fmtSpecSetupDto.CountryDescription,
		fmtSpecSetupDto.CountryTag,
		fmtSpecSetupDto.CountryCultureName,
		fmtSpecSetupDto.CountryAbbreviatedName,
		fmtSpecSetupDto.CountryAlternateNames,
		fmtSpecSetupDto.CountryCodeTwoChar,
		fmtSpecSetupDto.CountryCodeThreeChar,
		fmtSpecSetupDto.CountryCodeNumber,
		ePrefix)
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
	}

	ePrefix.SetEPref(
		"FormatterCountry." +
			"SetToUnitedStatesDefaults()")

	countryUtil :=
		formatterCountryUtility{}

	return countryUtil.setToUnitedStatesDefaults(
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
	}

	ePrefix.SetEPref(
		"FormatterCountry." +
			"SetToUnitedStatesDefaults()")

	nStrFmtSpecCntryQuark :=
		formatterCountryQuark{}

	isValid,
		_ := nStrFmtSpecCntryQuark.testValidityOfFormatterCountry(
		fmtCountry,
		ePrefix)

	if isValid {
		return nil
	}

	countryUtil :=
		formatterCountryUtility{}

	return countryUtil.setToUnitedStatesDefaults(
		fmtCountry,
		ePrefix)
}

// SetWithComponents - Sets the data values for current
// FormatterCountry instance.
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
	}

	ePrefix.SetEPref(
		"FormatterCountry.SetWithComponents()")

	nStrFmtSpecCntryMech :=
		formatterCountryMechanics{}

	return nStrFmtSpecCntryMech.setWithComponents(
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
