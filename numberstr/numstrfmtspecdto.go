package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFmtSpecDto struct {
	idNo           uint64
	idString       string
	description    string
	tag            string
	countryCulture NumStrFmtSpecCountryDto
	absoluteValue  NumStrFmtSpecAbsoluteValueDto
	currencyValue  NumStrFmtSpecCurrencyValueDto
	signedNumValue NumStrFmtSpecSignedNumValueDto
	sciNotation    NumStrFmtSpecSciNotationDto
	lock           *sync.Mutex
}

// NewFromComponents - Creates and returns a new instance of
// NumStrFmtSpecDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtSpecDto                *NumStrFmtSpecDto
//     - A pointer to an instance of NumStrFmtSpecDto. The data
//       values for this object will be overwritten and set
//       according to information passed in the following input'
//       parameters.
//
//  idNo                          uint64
//     - A user created identification number associated with the
//       'nStrFmtSpecDto' object.
//
//
//  idString                      string
//     - User defined text identification associated with the
//       'nStrFmtSpecDto' object.
//
//  description                   string
//     - User defined text description associated with the
//       'nStrFmtSpecDto' object.
//
//
//  tag                           string
//     - User defined tag associated with the 'nStrFmtSpecDto'
//       object.
//
//
//  countryCulture                NumStrFmtSpecCountryDto
//     - A valid and fully populated NumStrFmtSpecCountryDto
//       object. This object contains information on the country
//       or culture related to this number string format.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  absoluteValue                 NumStrFmtSpecAbsoluteValueDto
//     - A valid and fully populated NumStrFmtSpecAbsoluteValueDto
//       object. This object contains formatting specifications
//       controlling the text display of absolute numeric values.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  currencyValue                 NumStrFmtSpecCurrencyValueDto
//     - A valid and fully populated NumStrFmtSpecCurrencyValueDto
//       object. This object contains formatting specifications
//       controlling the text display of currency number strings.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  signedNumValue                NumStrFmtSpecSignedNumValueDto
//     - A valid and fully populated NumStrFmtSpecSignedNumValueDto
//       object. This object contains formatting specifications
//       controlling the text display of signed numeric values.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  sciNotation                   NumStrFmtSpecSciNotationDto
//     - A valid and fully populated NumStrFmtSpecCountryDto
//       object. This object contains formatting specifications
//       controlling the text display of scientific notation.
//
//       If this object is judged to be invalid, an error
//       will be returned.
//
//
//  ePrefix                    string
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
//  NumStrFmtSpecDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecDto.
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
func (fmtSpecDto NumStrFmtSpecDto) NewFromComponents(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture NumStrFmtSpecCountryDto,
	absoluteValue NumStrFmtSpecAbsoluteValueDto,
	currencyValue NumStrFmtSpecCurrencyValueDto,
	signedNumValue NumStrFmtSpecSignedNumValueDto,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix string) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecDto.NewFromComponents() "

	newFmtSpecDto := NumStrFmtSpecDto{}

	nStrFmtSpecDtoMech := numStrFmtSpecDtoMechanics{}

	newFmtSpecDto.lock = new(sync.Mutex)

	err := nStrFmtSpecDtoMech.setNumStrFmtSpecDto(
		&newFmtSpecDto,
		idNo,
		idString,
		description,
		tag,
		countryCulture,
		absoluteValue,
		currencyValue,
		signedNumValue,
		sciNotation,
		ePrefix)

	return newFmtSpecDto, err
}

// NewFromFmtSpecSetupDto - Creates and returns a new
// NumStrFmtSpecDto instance based on input received
// from an instance of NumStrFmtSpecSetupDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     *NumStrFmtSpecSetupDto
//     - A data structure conveying setup information for a
//       NumStrFmtSpecDto object.
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
//         AbsoluteValNumFieldLen                    int
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     int
//         CurrencyCode                              string
//         CurrencyName                              string
//         CurrencySymbol                            rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumFieldLen                       int
//         DecimalSeparator                          rune
//         IntegerDigitsSeparator                    rune
//         IntegerDigitsGroupingSequence             []uint
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumFieldLen                   int
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//         Lock                                      *sync.Mutex
//       }
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
//  NumStrFmtSpecDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecDto will be returned to the caller.
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
func (fmtSpecDto NumStrFmtSpecDto) NewFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix string) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecDto.NewFromFmtSpecSetupDto() "

	newNumStrFmtSpecDto :=
		NumStrFmtSpecDto{}

	newNumStrFmtSpecDto.lock = new(sync.Mutex)

	if fmtSpecSetupDto == nil {
		return newNumStrFmtSpecDto,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix)
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech :=
		numStrFmtSpecDtoMechanics{}

	err := nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		&newNumStrFmtSpecDto,
		fmtSpecSetupDto,
		ePrefix+
			"newNumStrFmtSpecDto\n")

	return newNumStrFmtSpecDto, err
}

// SetFromFmtSpecSetupDto - Overwrites and sets the data values
// for the current NumStrFmtSpecDto instance based on the input
// data values passed in a NumStrFmtSpecSetupDto structure.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtSpecSetupDto     NumStrFmtSpecSetupDto
//     - A data structure conveying setup and configuration
//       information for a NumStrFmtSpecDto object.
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
//         AbsoluteValNumFieldLen                    int
//         CurrencyPositiveValueFmt                  string
//         CurrencyNegativeValueFmt                  string
//         CurrencyDecimalDigits                     int
//         CurrencyCode                              string
//         CurrencyName                              string
//         CurrencySymbol                            rune
//         CurrencyTurnOnIntegerDigitsSeparation     bool
//         CurrencyNumFieldLen                       int
//         DecimalSeparator                          rune
//         IntegerDigitsSeparator                    rune
//         IntegerDigitsGroupingSequence             []uint
//         SignedNumValPositiveValueFmt              string
//         SignedNumValNegativeValueFmt              string
//         SignedNumValTurnOnIntegerDigitsSeparation bool
//         SignedNumValNumFieldLen                   int
//         SciNotMantissaLength                      uint
//         SciNotExponentChar                        rune
//         SciNotExponentUsesLeadingPlus             bool
//         SciNotNumFieldLen                         int
//       }
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
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (fmtSpecDto *NumStrFmtSpecDto) SetFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix string) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecDto.NewFromFmtSpecSetupDto() "

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix)
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech :=
		numStrFmtSpecDtoMechanics{}

	return nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		fmtSpecDto,
		fmtSpecSetupDto,
		ePrefix+
			"fmtSpecDto\n")
}
