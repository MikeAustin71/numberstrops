package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDtoMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of numStrFmtSpecDtoUtility.
//
func (nStrFmtSpecDtoMech numStrFmtSpecDtoMechanics) ptr() *numStrFmtSpecDtoMechanics {

	if nStrFmtSpecDtoMech.lock == nil {
		nStrFmtSpecDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech.lock.Lock()

	defer nStrFmtSpecDtoMech.lock.Unlock()

	newNumStrFmtSpecDtoMechanics := new(numStrFmtSpecDtoMechanics)

	newNumStrFmtSpecDtoMechanics.lock = new(sync.Mutex)

	return newNumStrFmtSpecDtoMechanics
}

// setFromFmtSpecSetupDto - Sets the data values of a passed
// NumStrFmtSpecDto instance based on the data values contained
// in a NumStrFmtSpecSetupDto structure.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtSpecDto      *NumStrFmtSpecDto
//     - A pointer to an instance of NumStrFmtSpecDto. The data
//      values contained in this object will be overwritten and
//      set to new values based on NumStrFmtSpecSetupDto instance
//      passed as an input parameter.
//
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
//         CurrencySymbols                            rune
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
func (nStrFmtSpecDtoMech *numStrFmtSpecDtoMechanics) setFromFmtSpecSetupDto(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	fmtSpecSetupDto *NumStrFmtSpecSetupDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecDtoMech.lock == nil {
		nStrFmtSpecDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech.lock.Lock()

	defer nStrFmtSpecDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecDtoMechanics.setFromFmtSpecSetupDto()")

	if nStrFmtSpecDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'nStrFmtSpecDto' is nil pointer!\n",
			ePrefix.String())
		return err
	}

	if nStrFmtSpecDto.lock == nil {
		nStrFmtSpecDto.lock = new(sync.Mutex)
	}

	if fmtSpecSetupDto == nil {
		err = fmt.Errorf("%v"+
			"Error: Input parameter 'fmtSpecSetupDto' is nil pointer!\n",
			ePrefix.String())
		return err
	}

	if fmtSpecSetupDto.Lock == nil {
		fmtSpecSetupDto.Lock = new(sync.Mutex)
	}

	fmtSpecSetupDto.Lock.Lock()

	nStrFmtSpecDto.idNo = fmtSpecSetupDto.IdNo

	nStrFmtSpecDto.idString = fmtSpecSetupDto.IdString

	nStrFmtSpecDto.description =
		fmtSpecSetupDto.Description

	nStrFmtSpecDto.tag =
		fmtSpecSetupDto.Tag

	fmtSpecSetupDto.Lock.Unlock()

	err = nStrFmtSpecDto.countryCulture.SetFromFmtSpecSetupDto(
		fmtSpecSetupDto,
		ePrefix.XCtx("fmtSpecSetupDto->CountryCulture"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.absoluteValue.SetFromFmtSpecSetupDto(
		fmtSpecSetupDto,
		ePrefix.XCtx("fmtSpecSetupDto->AbsoluteValue"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.currencyValue.SetFromFmtSpecSetupDto(
		fmtSpecSetupDto,
		ePrefix.XCtx("fmtSpecSetupDto"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.signedNumValue.SetFromFmtSpecSetupDto(
		fmtSpecSetupDto,
		ePrefix.XCtx("fmtSpecSetupDto"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.sciNotation.SetFromFmtSpecSetupDto(
		fmtSpecSetupDto,
		ePrefix.XCtx("fmtSpecSetupDto"))

	return err
}

// setNumStrFmtSpecDto - Sets the data values for the designated
// NumStrFmtSpecDto object, 'nStrFmtSpecDto', based on components
// passed as input parameters.
//
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
func (nStrFmtSpecDtoMech *numStrFmtSpecDtoMechanics) setNumStrFmtSpecDto(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture NumStrFmtSpecCountryDto,
	absoluteValue NumStrFmtSpecAbsoluteValueDto,
	currencyValue NumStrFmtSpecCurrencyValueDto,
	signedNumValue NumStrFmtSpecSignedNumValueDto,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix *ErrPrefixDto) (err error) {

	if nStrFmtSpecDtoMech.lock == nil {
		nStrFmtSpecDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoMech.lock.Lock()

	defer nStrFmtSpecDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrFmtSpecDtoMechanics.setNumStrFmtSpecDto()")

	if nStrFmtSpecDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDto' is nil pointer!\n",
			ePrefix.String())

		return err
	}

	if nStrFmtSpecDto.lock == nil {
		nStrFmtSpecDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDto.idNo = idNo
	nStrFmtSpecDto.idString = idString
	nStrFmtSpecDto.description = description
	nStrFmtSpecDto.tag = tag

	err = nStrFmtSpecDto.countryCulture.CopyIn(
		&countryCulture,
		ePrefix)

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.absoluteValue.CopyIn(
		&absoluteValue,
		ePrefix.XCtx(
			"absoluteValue->nStrFmtSpecDto.absoluteValue"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.currencyValue.CopyIn(
		&currencyValue,
		ePrefix.XCtx(
			"currencyValue->nStrFmtSpecDto.currencyValue"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.signedNumValue.CopyIn(
		&signedNumValue,
		ePrefix.XCtx(
			"signedNumValue->nStrFmtSpecDto.signedNumValue"))

	if err != nil {
		return err
	}

	err = nStrFmtSpecDto.sciNotation.CopyIn(
		&sciNotation,
		ePrefix.XCtx(
			"sciNotation->nStrFmtSpecDto.sciNotation"))

	return err
}
