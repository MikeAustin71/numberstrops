package numberstr

import "sync"

type NumStrFmtSpecCountryDto struct {
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
// NumStrFmtSpecCountryDto instance  to the data fields
// of the current instance of NumStrFmtSpecCountryDto
// instance.
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) CopyIn(
	inComingSpecCntryDto *NumStrFmtSpecCountryDto,
	ePrefix string) error {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecCountryDto.CopyIn() "

	nStrFmtSpecCntryQuark :=
		numStrFmtSpecCountryDtoQuark{}

	return nStrFmtSpecCntryQuark.copyIn(
		nStrFmtSpecCntryDto,
		inComingSpecCntryDto,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current
// NumStrFmtSpecCountryDto instance.
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) CopyOut() NumStrFmtSpecCountryDto {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryQuark :=
		numStrFmtSpecCountryDtoQuark{}

	newFmtSpecCntryDto,
		_ :=
		nStrFmtSpecCntryQuark.copyOut(
			nStrFmtSpecCntryDto,
			"")

	return newFmtSpecCntryDto
}

// GetIdNo - Returns the value of member variable 'idNo'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetIdNo() uint64 {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.idNo
}

// GetIdString - Returns the value of member variable 'idString'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetIdString() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.idString
}

// GetAbbreviatedCountryName - Returns the value of member variable
// 'abbreviatedCountryName'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetAbbreviatedCountryName() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.abbreviatedCountryName
}

// GetAlternateCountryNames - Returns the value of member variable
// 'alternateCountryNames'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetAlternateCountryNames() []string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	var altCountryNames []string

	lenAltCntryNames :=
		len(nStrFmtSpecCntryDto.alternateCountryNames)

	if lenAltCntryNames == 0 {
		altCountryNames = make([]string, 0, 10)
	} else {

		altCountryNames = make([]string, lenAltCntryNames, 10)

		for i := 0; i < lenAltCntryNames; i++ {
			altCountryNames[i] =
				nStrFmtSpecCntryDto.alternateCountryNames[i]
		}
	}

	return altCountryNames
}

// GetCountryCodeNumber - Returns the value of member variable
// 'countryCodeThreeChar'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetCountryCodeNumber() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.countryCodeNumber
}

// GetCountryCodeThreeChar - Returns the value of member variable
// 'countryCodeThreeChar'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetCountryCodeThreeChar() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.countryCodeThreeChar
}

// GetCountryCodeTwoChar - Returns the value of member variable
// 'countryCodeTwoChar'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetCountryCodeTwoChar() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.countryCodeTwoChar
}

// GetCountryCultureName - Returns the value of member variable
// 'countryCultureName'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetCountryCultureName() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.countryCultureName
}

// GetDescription - Returns the value of member variable 'description'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetDescription() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.description
}

// GetTag - Returns the value of member variable 'tag'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) GetTag() string {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	return nStrFmtSpecCntryDto.tag
}

// HasData - Returns 'true' if any of the significant data fields has
// a non-zero value. If all of the data fields have zero value, this
// method returns 'false'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) HasData() bool {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryQuark :=
		numStrFmtSpecCountryDtoQuark{}

	hasData,
		_ := nStrFmtSpecCntryQuark.hasData(
		nStrFmtSpecCntryDto,
		"")

	return hasData
}

// New() - Creates and returns a new instance of NumStrFmtSpecDigitsSeparatorsDto.
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecCountryDto
//     - If this method completes successfully, new instance of
//       NumStrFmtSpecCountryDto will be created and
//       returned.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecCntryDto NumStrFmtSpecCountryDto) New(
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
	ePrefix string) (
	NumStrFmtSpecCountryDto,
	error) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecCountryDto.New() "

	nStrFmtSpecCntryQuark :=
		numStrFmtSpecCountryDtoQuark{}

	newCntryDto := NumStrFmtSpecCountryDto{}

	err :=
		nStrFmtSpecCntryQuark.setCountryDto(
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

// SetIdNo - Sets the value of member variable 'idNo'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetIdNo(
	idNo uint64) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.idNo = idNo
}

// SetIdString - Sets the value of member variable 'idString'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetIdString(
	idString string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.idString = idString
}

// SetAbbreviatedCountryName - Sets the value of member variable
// 'abbreviatedCountryName'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetAbbreviatedCountryName(
	abbreviatedCountryName string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.abbreviatedCountryName =
		abbreviatedCountryName
}

// SetAlternateCountryNames - Sets the value of member variable
// 'alternateCountryNames'.
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetAlternateCountryNames(
	alternateCountryNames []string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	lenAltCntryNames :=
		len(alternateCountryNames)

	if lenAltCntryNames == 0 {

		nStrFmtSpecCntryDto.alternateCountryNames =
			make([]string, 0, 10)

	} else {

		nStrFmtSpecCntryDto.alternateCountryNames = make([]string, lenAltCntryNames, 10)

		for i := 0; i < lenAltCntryNames; i++ {
			nStrFmtSpecCntryDto.alternateCountryNames[i] =
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
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetCountryCodeNumber(
	countryCodeNumber string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.countryCodeNumber =
		countryCodeNumber
}

// SetCountryCodeThreeChar - Sets the value of member variable
// 'countryCodeThreeChar'.
//
// The ISO 3166-1 alpha-3 three character country code.
//  Reference:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetCountryCodeThreeChar(
	countryCodeThreeChar string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.countryCodeThreeChar =
		countryCodeThreeChar
}

// SetCountryCodeTwoChar - Sets the value of member variable
// 'countryCodeTwoChar'.
//
// The ISO 3166-1 alpha-2 two character country code.
// Reference:
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetCountryCodeTwoChar(
	countryCodeTwoChar string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.countryCodeTwoChar =
		countryCodeTwoChar

}

// SetCountryCultureName - Sets the value of member variable
// 'countryCultureName'.
//
// Official name of the country.
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetCountryCultureName(
	countryCultureName string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.countryCultureName =
		countryCultureName
}

// SetDescription - Sets the value of member variable 'description'.
// This is a user defined description.
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetDescription(
	description string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.description =
		description
}

// SetTag - Sets the value of member variable 'tag'.
// This is a user defined description.
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) SetTag(
	tag string) {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryDto.tag = tag
}