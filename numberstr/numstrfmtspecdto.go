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

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecDto instance  to the data fields of the current
// instance of NumStrFmtSpecDto instance.
//
// If input parameter 'incomingFmtSpecDto' is judged to be invalid,
// this method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrFmtSpecDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFmtSpecDto         *NumStrFmtSpecDto
//     - A pointer to an instance of NumStrFmtSpecDto.
//       The data values in this object will be copied to the
//       current NumStrFmtSpecDto instance.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
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
func (fmtSpecDto *NumStrFmtSpecDto) CopyIn(
	incomingFmtSpecDto *NumStrFmtSpecDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.CopyIn()")

	nStrFmtSpecDtoElectron :=
		numStrFmtSpecDtoElectron{}

	return nStrFmtSpecDtoElectron.copyIn(
		fmtSpecDto,
		incomingFmtSpecDto,
		ePrefix.XCtx("incomingFmtSpecDto->fmtSpecDto"))
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrFmtSpecDto instance.
//
// If the current NumStrFmtSpecDto instance is judged to be
// invalid, this method will return an error.
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
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDto
//     - If this method completes successfully, a new instance of
//       NumStrFmtSpecDto will be created and returned containing
//       all of the data values copied from the current instance of
//       NumStrFmtSpecDto.
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
func (fmtSpecDto *NumStrFmtSpecDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.CopyOut()")

	nStrFmtSpecDtoElectron :=
		numStrFmtSpecDtoElectron{}

	return nStrFmtSpecDtoElectron.copyOut(
		fmtSpecDto,
		ePrefix.XCtx("Copy Out from 'fmtSpecDto'"))
}

// GetAbsoluteValueSpec - Returns a deep copy of the member variable
// 'absoluteValue', of type NumStrFmtSpecAbsoluteValueDto.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the NumStrFmtSpecAbsoluteValueDto
// object. To validate the current NumStrFmtSpecDto instance
// reference methods NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetAbsoluteValueSpec() NumStrFmtSpecAbsoluteValueDto {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	absVal,
		_ := fmtSpecDto.absoluteValue.CopyOut(
		ErrPrefixDto{}.Ptr())

	return absVal
}

// GetCountryCulture - Returns a deep copy of the member variable
// 'countryCulture', of type NumStrFmtSpecCountryDto.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the NumStrFmtSpecCountryDto object. To
// validate the current NumStrFmtSpecDto instance reference methods
// NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetCountryCulture() NumStrFmtSpecCountryDto {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	countryCulture,
		_ := fmtSpecDto.countryCulture.CopyOut(ErrPrefixDto{}.Ptr())

	return countryCulture
}

// GetCurrencySpec - Returns a deep copy of the member variable
// 'currencyValue', of type NumStrFmtSpecCurrencyValueDto.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the NumStrFmtSpecCountryDto object. To
// validate the current NumStrFmtSpecDto instance reference methods
// NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetCurrencySpec() NumStrFmtSpecCurrencyValueDto {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	currencySpec,
		_ := fmtSpecDto.currencyValue.CopyOut(ErrPrefixDto{}.Ptr())

	return currencySpec
}

// GetDecimalSeparator - Returns the decimal separator
// from 'currency' value.
//
func (fmtSpecDto *NumStrFmtSpecDto) GetDecimalSeparator() rune {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	return fmtSpecDto.currencyValue.numberSeparatorsDto.GetDecimalSeparator()
}

// GetScientificNotationSpec - Returns a deep copy of the member
// variable 'sciNotation', of type NumStrFmtSpecSciNotationDto.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the NumStrFmtSpecSciNotationDto
// object. To validate the current NumStrFmtSpecDto instance
// reference methods NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetScientificNotationSpec() NumStrFmtSpecSciNotationDto {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	scientificNotationSpec,
		_ := fmtSpecDto.sciNotation.CopyOut(ErrPrefixDto{}.Ptr())

	return scientificNotationSpec
}

// GetSignedNumSpec - Returns a deep copy of the member variable
// 'signedNumValue', of type NumStrFmtSpecSignedNumValueDto.
//
// IMPORTANT
//
// No validity tests are performed on the current NumStrFmtSpecDto
// instance before returning the NumStrFmtSpecSignedNumValueDto
// object. To validate the current NumStrFmtSpecDto instance
// reference methods NumStrFmtSpecDto.IsValidInstance() and
// NumStrFmtSpecDto.IsValidInstanceError().
//
func (fmtSpecDto *NumStrFmtSpecDto) GetSignedNumSpec() NumStrFmtSpecSignedNumValueDto {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	signedNumSpec,
		_ := fmtSpecDto.signedNumValue.CopyOut(ErrPrefixDto{}.Ptr())

	return signedNumSpec
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecDto instance to determine whether the current
// instance is valid in all respects.
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
//       NumStrFmtSpecDto is valid, or not. If the current
//       NumStrFmtSpecDto contains valid data, this method returns
//       'true'. If the data is invalid, this method will return
//       'false'.
//
func (fmtSpecDto *NumStrFmtSpecDto) IsValidInstance() (
	isValid bool) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	nStrFmtSpecDtoQuark :=
		numStrFmtSpecDtoQuark{}

	isValid,
		_ = nStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto(
		fmtSpecDto,
		ErrPrefixDto{}.Ptr())

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrFmtSpecDto instance to determine whether the
// current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
func (fmtSpecDto *NumStrFmtSpecDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx(
		"NumStrFmtSpecDto.IsValidInstanceError()",
		"Testing Validity of 'fmtSpecDto'")

	nStrFmtSpecDtoQuark :=
		numStrFmtSpecDtoQuark{}

	_,
		err :=
		nStrFmtSpecDtoQuark.testValidityOfNumStrFmtSpecDto(
			fmtSpecDto,
			ePrefix)

	return err
}

// NewFromComponents - Creates and returns a new instance of
// NumStrFmtSpecDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
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
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.NewFromComponents()")

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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (fmtSpecDto NumStrFmtSpecDto) NewFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.NewFromFmtSpecSetupDto()")

	newNumStrFmtSpecDto :=
		NumStrFmtSpecDto{}

	newNumStrFmtSpecDto.lock = new(sync.Mutex)

	if fmtSpecSetupDto == nil {
		return newNumStrFmtSpecDto,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
				"'fmtSpecSetupDto' is a 'nil' pointer!\n",
				ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech :=
		numStrFmtSpecDtoMechanics{}

	err := nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		&newNumStrFmtSpecDto,
		fmtSpecSetupDto,
		ePrefix.XCtx("newNumStrFmtSpecDto"))

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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
func (fmtSpecDto *NumStrFmtSpecDto) SetFromFmtSpecSetupDto(
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumStrFmtSpecDto.NewFromFmtSpecSetupDto()")

	if fmtSpecSetupDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtSpecSetupDto' is invalid!\n"+
			"'fmtSpecSetupDto' is a 'nil' pointer!\n",
			ePrefix.String())
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech :=
		numStrFmtSpecDtoMechanics{}

	ePrefix.SetCtx("fmtSpecDto")

	return nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		fmtSpecDto,
		fmtSpecSetupDto,
		ePrefix)
}

// SetNumStrFmtSpecDto - Overwrites and sets the data values
// for the current NumStrFmtSpecDto instance based on component
// values passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
func (fmtSpecDto *NumStrFmtSpecDto) SetNumStrFmtSpecDto(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture NumStrFmtSpecCountryDto,
	absoluteValue NumStrFmtSpecAbsoluteValueDto,
	currencyValue NumStrFmtSpecCurrencyValueDto,
	signedNumValue NumStrFmtSpecSignedNumValueDto,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrFmtSpecDto.SetNumStrFmtSpecDto()")

	nStrFmtSpecDtoMech := numStrFmtSpecDtoMechanics{}

	err := nStrFmtSpecDtoMech.setNumStrFmtSpecDto(
		fmtSpecDto,
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

	return err
}
