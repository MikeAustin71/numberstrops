package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecDtoUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of numStrFmtSpecDtoUtility.
//
func (nStrFmtSpecDtoUtil numStrFmtSpecDtoUtility) ptr() *numStrFmtSpecDtoUtility {

	if nStrFmtSpecDtoUtil.lock == nil {
		nStrFmtSpecDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoUtil.lock.Lock()

	defer nStrFmtSpecDtoUtil.lock.Unlock()

	newNumStrFmtSpecDtoUtility := new(numStrFmtSpecDtoUtility)

	newNumStrFmtSpecDtoUtility.lock = new(sync.Mutex)

	return newNumStrFmtSpecDtoUtility
}

// setDefaultFormatSpec - Sets the default Number String Format
// Specification. The default Number String Format Specification is
// the United States decimal separator ('.'), thousands separator
// (',') and currency symbol ('$').
//
// The input parameter 'numStrFmtSpec' will be configured with the
// United States Number String Format Specification.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrFmtSpec       *NumStrFmtSpecDto
//     - This object contains all the formatting specifications
//       required to format numeric values contained in type
//       NumStrDto. This method will configure this object for
//       the Unites States default Number String Format
//       Specification.
//
//       type NumStrFmtSpecDto struct {
//         idNo           uint64
//         idString       string
//         description    string
//         tag            string
//         countryCulture NumStrFmtSpecCountryDto
//         absoluteValue  FormatterAbsoluteValue
//         currencyValue  FormatterCurrency
//         signedNumValue NumStrFmtSpecSignedNumValueDto
//         sciNotation    NumStrFmtSpecSciNotationDto
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
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  err                 error
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
func (nStrFmtSpecDtoUtil *numStrFmtSpecDtoUtility) setDefaultFormatSpec(
	numStrFmtSpec *NumStrFmtSpecDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtSpecDtoUtil.lock == nil {
		nStrFmtSpecDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoUtil.lock.Lock()

	defer nStrFmtSpecDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrFmtSpecDtoUtility.setDefaultFormatSpec()")

	if numStrFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"numStrFmtSpec is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	numStrFmtSpecSetupDto := NumStrFormatCountry{}.Ptr().UnitedStates()

	err =
		numStrFmtSpecDtoMechanics{}.ptr().setFromFmtSpecSetupDto(
			numStrFmtSpec,
			&numStrFmtSpecSetupDto,
			ePrefix.XCtx(
				"numStrFmtSpec, numStrFmtSpecSetupDto"))

	return err
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
//  nStrFmtSpecDto                *NumStrFmtSpecDto
//     - A pointer to an instance of NumStrFmtSpecDto. The data
//       values contained in this object will be overwritten and
//       set to new values based on input parameters passed to this
//       method.
//
//
//  decimalSeparatorChar          rune
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar        rune
//     - The character which will be used to delimit 'thousands' in
//       integer number strings. In the United States, the Thousands
//       separator is the comma character (',').
//           Example: '1,000,000,000'
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnIntegerDigitsSeparation' is a boolean
//       flag used to control the 'Thousands Separator'. When set
//       to 'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbols               []rune
//     - The authorized unicode character symbols associated with
//       this currency specification. The currency symbol for the
//       United States is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
//
//  minorCurrencyName             string
//     - The minor currency name. In the United States, the minor
//       currency name is 'Cent'.
//
//
//  minorCurrencySymbols          []rune
//     - The unicode character for minor currency symbol. In the
//       United States, the minor currency symbol is the cent sign
//       (¢). Some countries and cultures have currency symbols
//       consisting of two or more characters.
//
//
//  currencyPositiveValueFmt      string
//     - This format string will be used to format positive currency
//       value in text number strings. Valid positive currency value
//       format strings must comply with the following constraints.
//
//       Positive Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the positive value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for positive currency values are
//       listed as follows:
//
//          $127.54
//          $ 127.54
//          127.54$
//          127.54 $
//
//          $127.54+
//          $127.54 +
//          $ 127.54+
//          $ 127.54 +
//
//          $+127.54
//          $ +127.54
//          $+ 127.54
//          $ + 127.54
//
//          127.54$+
//          127.54$ +
//          127.54 $+
//          127.54$ +
//          127.54 $ +
//
//          127.54+$
//          127.54+ $
//          127.54 +$
//          127.54+ $
//          127.54 + $
//
//          +127.54$
//          +127.54 $
//          + 127.54$
//          + 127.54 $
//
//
//  currencyNegativeValueFmt      string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
//
//               "$" - Placeholder for the previously selected currency
//                     symbol associated with the user's preferred country
//                     or culture. This placeholder symbol, '$', MUST BE
//                     present in the negative value format string in order
//                     to correctly position the actual currency symbol
//                     relative to the currency numeric value.
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//       Valid format strings for negative currency values are
//       listed as follows:
//
//          ( $127.54 )
//          ( $ 127.54 )
//          ($ 127.54)
//          ($127.54)
//
//          $(127.54)
//          $ (127.54)
//          $( 127.54 )
//          $ ( 127.54 )
//
//          ( 127.54$ )
//          ( 127.54 $ )
//          ( 127.54 $)
//          (127.54$)
//
//          (127.54)$
//          (127.54) $
//          ( 127.54 )$
//          ( 127.54 ) $
//
//          (-) $127.54
//          (-) $ 127.54
//          (-)$127.54
//          (-)$ 127.54
//
//          $ (-)127.54
//          $ (-) 127.54
//          $(-)127.54
//          $(-) 127.54
//
//          (-) 127.54$
//          (-) 127.54 $
//          (-)127.54$
//          (-)127.54 $
//
//          127.54(-) $
//          127.54 (-) $
//          127.54(-)$
//          127.54 (-)$
//
//          127.54$(-)
//          127.54$ (-)
//          127.54 $ (-)
//          127.54 $(-)
//
//          $127.54(-)
//          $127.54 (-)
//          $ 127.54(-)
//          $ 127.54 (-)
//
//          - $127.54
//          - $ 127.54
//          -$127.54
//          -$ 127.54
//
//          $ -127.54
//          $ - 127.54
//          $-127.54
//          $- 127.54
//
//          - 127.54$
//          - 127.54 $
//          -127.54$
//          -127.54 $
//
//          127.54- $
//          127.54 - $
//          127.54-$
//          127.54 -$
//
//          127.54$-
//          127.54$ -
//          127.54 $ -
//          127.54 $-
//
//          $127.54-
//          $127.54 -
//          $ 127.54-
//          $ 127.54 -
//
//
//  signedNumPositiveValueFmt     string
//     - A string specifying the number string format to be used in
//       formatting positive numeric values for signed numbers in
//       text strings.
//
//       Positive Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for positive number formatting.
//
//       Valid format strings for positive numeric values
//       (NOT Currency) are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Positive Value Format
//
//
//  signedNumNegativeValueFmt     string
//     - A string specifying the number string format to be used in
//       formatting negative numeric values in text strings.
//
//       Negative Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators, Decimal
//                     Separators and Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the
//                     format string, the minus sign ('-') specifies
//                     where the minus sign will be positioned
//                     relative to the numeric value in the text
//                     number string.
//
//             "(-)" - These three characters are often used in
//                     Europe and the United Kingdom to classify
//                     a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are
//                     frequently used in the United States
//                     to classify a numeric value as negative.
//
//
//       Valid format strings for negative numeric values
//       (NOT Currency) are listed as follows:
//
//               -127.54   The Default Negative Value Format String
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -127.54
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -NUMFIELD
//               - NUMFIELD
//               NUMFIELD-
//               NUMFIELD -
//               (-) NUMFIELD
//               (-)NUMFIELD
//               NUMFIELD(-)
//               NUMFIELD (-)
//               (NUMFIELD)
//               ( NUMFIELD )
//               (NUMFIELD)
//               ( NUMFIELD )
//
//
//  requestedNumberFieldLen       int
//     - This is the requested length of the number field in which
//       the number string will be displayed. If this field length
//       is greater than the actual length of the number string,
//       the number string will be right justified within the the
//       number field. If the actual number string length is greater
//       than the requested number field length, the number field
//       length will be automatically expanded to display the entire
//       number string. The 'requested' number field length is used
//       to create number fields of standard lengths for text
//       presentations.
//
//
//  numberFieldTextJustify        TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to a number string when
//       it is positioned inside of a number field. This
//       enumeration value must be one of the three following
//       format specifications:
//
//       1. Left   - Signals that the text justification format is
//                   set to 'Left-Justify'. Strings within text
//                   fields will be flush with the left margin.
//                          Example: "TextString      "
//
//       2. Right  - Signals that the text justification format is
//                   set to 'Right-Justify'. Strings within text
//                   fields will terminate at the right margin.
//                          Example: "      TextString"
//
//       3. Center - Signals that the text justification format is
//                   is set to 'Centered'. Strings will be positioned
//                   in the center of the text field equidistant
//                   from the left and right margins.
//                           Example: "   TextString   "
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
func (nStrFmtSpecDtoUtil *numStrFmtSpecDtoUtility) setCustomFmtSpecDto(
	nStrFmtSpecDto *NumStrFmtSpecDto,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	integerDigitsGroupingSequence []uint,
	turnOnIntegerDigitsSeparation bool,
	currencyCode string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	currencyPositiveValueFmt string,
	currencyNegativeValueFmt string,
	signedNumPositiveValueFmt string,
	signedNumNegativeValueFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecDtoUtil.lock == nil {
		nStrFmtSpecDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecDtoUtil.lock.Lock()

	defer nStrFmtSpecDtoUtil.lock.Unlock()

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

	if decimalSeparatorChar == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorChar' is invalid!\n"+
			"decimalSeparatorChar== '0'\n",
			ePrefix.String())
	}

	if thousandsSeparatorChar == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparatorChar' is invalid!\n"+
			"thousandsSeparatorChar== '0'\n",
			ePrefix.String())
	}

	if currencySymbols == nil {
		currencySymbols = make([]rune, 0, 10)
	}

	lenCurrencySymbols := len(currencySymbols)

	if lenCurrencySymbols == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"'currencySymbols' is missing.\n"+
			"'currencySymbols' is a zero length array.\n",
			ePrefix.String())
	}

	if minorCurrencySymbols == nil {
		minorCurrencySymbols = make([]rune, 0, 10)
	}

	nStrCurrencyElectron := formatterCurrencyElectron{}

	var err error

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

	if requestedNumberFieldLen < 1 {
		requestedNumberFieldLen = -1
	}

	if !numberFieldTextJustify.XIsValid() {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberFieldTextJustify' is invalid!\n"+
			"numberFieldTextJustify== '%v'\n",
			ePrefix.String(),
			numberFieldTextJustify.XValueInt())
	}

	setupDto := NumStrFmtSpecSetupDto{}

	setupDto.IdNo = 9999
	setupDto.IdString = "9999"
	setupDto.Description = "Custom Format Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 9999
	setupDto.CountryIdString = "9999"
	setupDto.CountryDescription = "Custom Format Setup"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Custom Format Setup"
	setupDto.CountryAbbreviatedName = "Custom Format"

	setupDto.CountryAlternateNames =
		[]string{
			"Custom Format Setup"}

	setupDto.CountryCodeTwoChar = "XX"
	setupDto.CountryCodeThreeChar = "XXX"
	setupDto.CountryCodeNumber = "9999"

	setupDto.AbsoluteValFmt = signedNumPositiveValueFmt
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = turnOnIntegerDigitsSeparation
	setupDto.AbsoluteValNumFieldLen = requestedNumberFieldLen

	setupDto.AbsoluteValNumFieldTextJustify =
		numberFieldTextJustify

	setupDto.CurrencyPositiveValueFmt = currencyPositiveValueFmt
	setupDto.CurrencyNegativeValueFmt = currencyNegativeValueFmt
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = currencyCode
	setupDto.CurrencyName = currencyName
	setupDto.CurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		setupDto.CurrencySymbols,
		currencySymbols)

	setupDto.MinorCurrencyName = minorCurrencyName

	lenCurrencySymbols = len(minorCurrencySymbols)

	setupDto.MinorCurrencySymbols =
		make([]rune, lenCurrencySymbols, 10)

	_ = copy(
		setupDto.MinorCurrencySymbols,
		minorCurrencySymbols)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation

	setupDto.CurrencyNumFieldLen = requestedNumberFieldLen

	setupDto.CurrencyNumFieldTextJustify =
		numberFieldTextJustify

	setupDto.DecimalSeparator = decimalSeparatorChar
	setupDto.IntegerDigitsSeparator = thousandsSeparatorChar
	setupDto.IntegerDigitsGroupingSequence =
		integerDigitsGroupingSequence

	setupDto.SignedNumValPositiveValueFmt = signedNumPositiveValueFmt
	setupDto.SignedNumValNegativeValueFmt = signedNumNegativeValueFmt
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = turnOnIntegerDigitsSeparation
	setupDto.SignedNumValNumFieldLen = requestedNumberFieldLen

	setupDto.SignedNumValNumFieldTextJustify =
		numberFieldTextJustify

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.Lock = new(sync.Mutex)

	nStrFmtSpecDtoMech := numStrFmtSpecDtoMechanics{}

	return nStrFmtSpecDtoMech.setFromFmtSpecSetupDto(
		nStrFmtSpecDto,
		&setupDto,
		ePrefix.XCtx(
			"nStrFmtSpecDto, setupDto"))
}
