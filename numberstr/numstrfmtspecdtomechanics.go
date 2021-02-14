package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDtoMechanics struct {
	lock *sync.Mutex
}

// setCustomFmtSpecDto - Sets the data values of a passed
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
//  numericSeparatorsDto          NumericSeparatorsDto
//     - This instance of 'NumericSeparatorsDto' is
//       used to specify the separator characters which will be
//       included in the number string text display.
//
//        type NumericSeparatorsDto struct {
//         decimalSeparator              rune
//         integerDigitsSeparator        rune
//         integerDigitsGroupingSequence []uint
//        }
//
//        decimalSeparator rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerDigitsSeparator rune
//
//        This type also encapsulates the integer digits separator, often
//        referred to as the 'Thousands Separator'. This is used to
//        separate thousands digits within the integer component of a
//        number string.
//
//        integerDigitsGroupingSequence []uint
//
//        Related to the integer digits separator, the integer digits
//        grouping sequence is also encapsulated in this type. The integer
//        digits grouping sequence is used to identify the digits which
//        will be grouped and separated by the integer digits separator.
//
//        In most western countries integer digits to the left of the
//        decimal separator (a.k.a. decimal point) are separated into
//        groups of three digits representing a grouping of 'thousands'
//        like this: '1,000,000,000,000'. In this case the parameter
//        integer digits grouping sequence would be configured as:
//                     integerDigitsGroupingSequence = []uint{3}
//
//        In some countries and cultures other integer groupings are used.
//        In India, for example, a number might be formatted as like this:
//                      '6,78,90,00,00,00,00,000'
//        The right most group has three digits and all the others are
//        grouped by two. In this case integer digits grouping sequence
//        would be configured as:
//                     integerDigitsGroupingSequence = []uint{3,2}
//
//
//
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

func (nStrFmtSpecDtoMech *numStrFmtSpecDtoMechanics) setCustomFmtSpecDto(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	numSeps NumericSeparatorsDto,
	currencySymbol rune,
	currencyPositiveValueFmt string,
	currencyNegativeValueFmt string,
	signedNumPositiveValueFmt string,
	signedNumNegativeValueFmt string,
	ePrefix *ErrPrefixDto) error {

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

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecDto' is nil pointer!\n",
			ePrefix.String())
	}

	if nStrFmtSpecDto.lock == nil {
		nStrFmtSpecDto.lock = new(sync.Mutex)
	}

	var err error

	err = numSeps.IsValidInstanceError(
		ePrefix.XCtx("Testing validity of 'numSeps'"))

	if err != nil {
		return err
	}

	if currencySymbol == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbol' is invalid!\n"+
			"currencySymbol== '0'\n",
			ePrefix.String())
	}

	nStrCurrencyElectron := numStrFmtSpecCurrencyValueDtoElectron{}

	_,
		err = nStrCurrencyElectron.testCurrencyPositiveValueFormatStr(
		currencyPositiveValueFmt,
		ePrefix.XCtx(
			"currencyPositiveValueFmt"))

	if err != nil {
		return err
	}

	_,
		err = nStrCurrencyElectron.testCurrencyNegativeValueFormatStr(
		currencyNegativeValueFmt,
		ePrefix.XCtx(
			"currencyNegativeValueFmt"))

	if err != nil {
		return err
	}

	nStrSignedNumElectron := numStrSignedNumValElectron{}

	_,
		err = nStrSignedNumElectron.
		testSignedNumValPositiveValueFormatStr(
			signedNumPositiveValueFmt,
			ePrefix.XCtx(
				"signedNumPositiveValueFmt"))

	if err != nil {
		return nil
	}

	_,
		err = nStrSignedNumElectron.
		testSignedNumValNegativeValueFormatStr(
			signedNumNegativeValueFmt,
			ePrefix.XCtx(
				"signedNumNegativeValueFmt"))

	if err != nil {
		return nil
	}

	return nil
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
