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
