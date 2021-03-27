package numberstr

import (
	"fmt"
	"sync"
)

type formatterCurrencyMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCurrencyMechanics.
//
func (fmtCurrMech formatterCurrencyMechanics) ptr() *formatterCurrencyMechanics {

	if fmtCurrMech.lock == nil {
		fmtCurrMech.lock = new(sync.Mutex)
	}

	fmtCurrMech.lock.Lock()

	defer fmtCurrMech.lock.Unlock()

	newCurrencyValDtoMechanics :=
		new(formatterCurrencyMechanics)

	newCurrencyValDtoMechanics.lock = new(sync.Mutex)

	return newCurrencyValDtoMechanics
}

// setFormatterCurrencyWithComponents - Transfers new data to an
// instance of FormatterCurrency. After completion, all the data
// fields within input parameter 'formatterCurrency' will be
// overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency             *FormatterCurrency,
//
//     - A pointer to an instance of FormatterCurrency.
//       All of the data values in this object will be overwritten
//       set to new values based on the following input parameters.
//
//
//  positiveValueFmt              string
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
//  negativeValueFmt              string
//     - This format string will be used to format negative currency
//       values in text number strings. Valid negative currency value
//       format strings must comply with the following constraints.
//
//       Negative Currency Value Formatting Terminology and Placeholders:
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
//  decimalDigits                 uint
//     - The standard number of digits to the right of the decimal
//       place which is expected for currency values. In the United
//       States, currency is typically formatted with two digits to
//       the right of the decimal.
//         Example:  $24.92
//
//
//  currencyCode                  string
//     - The ISO 4217 Currency Code associated with this currency
//       specification. Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyCodeNo                string
//     - The ISO 4217 Currency Code Number associated with this
//       currency specification. The Currency Code Number is stored
//       as a string per ISO 4217.
//
//       Reference:
//        https://en.wikipedia.org/wiki/ISO_4217
//
//
//  currencyName                  string
//     - The official name for this currency.
//
//
//  currencySymbols                []rune
//     - The authorized character symbol associated with this
//       currency specification. In the United States, the currency
//       symbol is the dollar sign ('$'). Some countries and
//       cultures have currency symbols consisting of two or more
//       characters.
//
//
//  minorCurrencyName             string
//     - The minor currency name. In the United States, the minor
//       currency name is 'Cent'.
//
//
//  minorCurrencySymbols            []rune
//     - These are the unicode characters for minor currency
//       symbols. In the United States, the minor currency symbol
//       is the cent sign (¢), represented by a single unicode
//       character ('\U000000a2'). Some countries and cultures have
//       currency symbols consisting of two or more characters.
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
//  numericSeparators        NumericSeparators
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       including in the number string text display.
//
//        type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//        }
//
//        decimalSeparators     []rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//        The decimal separator may consist of one or more runes.
//
//        integerSeparatorsDto    NumStrIntSeparatorsDto
//
//        The NumStrIntSeparatorsDto type encapsulates the integer digits
//        separators, often referred to as the 'Thousands Separator'.
//        Integer digit separators are used to separate integers into
//        specific groups within a number string. The
//        NumStrIntSeparatorsDto manages an array or collection of
//        NumStrIntSeparator objects.
//
//               type NumStrIntSeparatorsDto struct {
//                 intSeparators []NumStrIntSeparator
//               }
//
//               type NumStrIntSeparator struct {
//                intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//                intSeparatorGrouping    uint    // Number of integer digits in a group
//                intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                                // A zero value signals unlimited repetitions.
//                restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//               }
//
//               intSeparatorChars          []rune
//                  - A series of runes or characters used to separate integer
//                    digits in a number string. These characters are commonly
//                    known as the 'thousands separator'. A 'thousands
//                    separator' is used to separate groups of integer digits to
//                    the left of the decimal separator (a.k.a. decimal point).
//                    In the United States, the standard integer digits
//                    separator is the single comma character (','). Other
//                    countries and cultures use periods, spaces, apostrophes or
//                    multiple characters to separate integers.
//                          United States Example:  1,000,000,000
//
//               intSeparatorGrouping       uint
//                  - This unsigned integer values specifies the number of
//                    integer digits within a group. This value is used to group
//                    integers within a number string.
//
//                    In most western countries integer digits to the left of
//                    the decimal separator (a.k.a. decimal point) are separated
//                    into groups of three digits representing a grouping of
//                    'thousands' like this: '1,000,000,000'. In this case the
//                    intSeparatorGrouping value would be set to three ('3').
//
//                     In some countries and cultures other integer groupings
//                     are used. In India, for example, a number might be
//                     formatted like this: '6,78,90,00,00,00,00,000'. Chinese
//                     Numerals have an integer grouping value of four ('4').
//                         Chinese Numerals Example: '12,3456,7890,2345'
//
//               intSeparatorRepetitions    uint
//                  - This unsigned integer value specifies the number of times
//                    this integer grouping is repeated. A value of zero signals
//                    that this integer grouping will be repeated indefinitely.
//
//               restartIntGroupingSequence bool
//                  - If the NumStrIntSeparator is the last element in an array
//                    of NumStrIntSeparator objects, this boolean flag signals
//                    whether the entire integer grouping sequence will be
//                    restarted from array element zero.
//
//
//  numFieldDto                NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
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
//  err                           error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (fmtCurrMech *formatterCurrencyMechanics) setFormatterCurrencyWithComponents(
	formatterCurrency *FormatterCurrency,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrMech.lock == nil {
		fmtCurrMech.lock = new(sync.Mutex)
	}

	fmtCurrMech.lock.Lock()

	defer fmtCurrMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyMechanics." +
			"setFormatterCurrencyWithComponents()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if formatterCurrency.lock == nil {
		formatterCurrency.lock = new(sync.Mutex)
	}

	newFormatterCurrency := FormatterCurrency{}

	newFormatterCurrency.lock = new(sync.Mutex)

	err =
		formatterCurrencyNanobot{}.ptr().
			setCurrencyData(
				&newFormatterCurrency,
				positiveValueFmt,
				negativeValueFmt,
				decimalDigits,
				currencyCode,
				currencyCodeNo,
				currencyName,
				currencySymbols,
				minorCurrencyName,
				minorCurrencySymbols,
				turnOnIntegerDigitsSeparation,
				ePrefix.XCtx(
					"newFormatterCurrency"))

	if err != nil {
		return err
	}

	err =
		newFormatterCurrency.numericSeparators.CopyIn(
			&numericSeparators,
			ePrefix.XCtx("numericSeparators->newFormatterCurrency"))

	if err != nil {
		return err
	}

	err =
		newFormatterCurrency.numFieldDto.CopyIn(
			&numFieldLenDto,
			ePrefix.XCtx(" numFieldDto->newFormatterCurrency"))

	if err != nil {
		return err
	}

	err =
		formatterCurrencyNanobot{}.ptr().
			copyIn(
				formatterCurrency,
				&newFormatterCurrency,
				ePrefix.XCtx("newFormatterCurrency-> "+
					"formatterCurrency"))

	return err
}

// setNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current FormatterCurrency instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency             *FormatterCurrency,
//
//     - A pointer to an instance of FormatterCurrency.
//       All of the data values in this object will be overwritten
//       set to new values based on the following input parameters.
//
//
//  numberFieldLenDto             NumberFieldDto
//     - The NumberFieldDto details the length of the number field
//       in which the signed numeric value will be displayed and
//       right justified.
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int
//         actualNumFieldLength    int
//         minimumNumFieldLength   int
//       }
//
//
//  ePrefix                       *ErrPrefixDto
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
func (fmtCurrMech *formatterCurrencyMechanics) setNumberFieldLengthDto(
	formatterCurrency *FormatterCurrency,
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtCurrMech.lock == nil {
		fmtCurrMech.lock = new(sync.Mutex)
	}

	fmtCurrMech.lock.Lock()

	defer fmtCurrMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyMechanics." +
			"setNumberFieldLengthDto()")

	if formatterCurrency == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer\n",
			ePrefix.String())

	}

	return formatterCurrency.
		numFieldDto.CopyIn(
		&numberFieldLenDto,
		ePrefix.XCtx(
			"formatterCurrency"))
}

// setNumericSeparators - Sets the Number Separators object
// for the FormatterCurrency instance,
// 'nStrFmtSpecCurrencyValDto'.
//
// The Number Separators object is used to specify the Decimal
// Separators Character(s) and the Integer Digits Separator
// Character(s).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency          *FormatterCurrency,
//
//     - A pointer to an instance of FormatterCurrency.
//       The Numeric Separator data values in this object will be
//       overwritten and set to new values based on the following
//       input parameters.
//
//
//  numericSeparators          NumericSeparators
//     - This instance of 'NumericSeparators' encapsulates all the
//       number separators required to format numeric values in
//       text strings. These separators include the 'Decimal
//       Separator(s)' and the 'Integer Digits Separators'.
//
//       The 'Integer Digits Separators' includes both the character or
//       characters used to separate groups of integers and the grouping
//       sequence. In Western Countries, integer grouping is most
//       commonly known as 'thousands' grouping.
//            United States Example: 1,000,0000,000
//
//       type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//       }
//
// decimalSeparators    []rune
//
// The 'Decimal Separator' is used to separate integer and
// fractional digits within a floating point number display.
// The decimal separator may consist of one or more runes.
//
//
// integerSeparatorsDto    NumStrIntSeparatorsDto
//
// The NumStrIntSeparatorsDto type encapsulates the integer digits
// separators, often referred to as the 'Thousands Separator'.
// Integer digit separators are used to separate integers into
// specific groups within a number string. The
// NumStrIntSeparatorsDto manages an array or collection of
// NumStrIntSeparator objects. Taken as a whole, these
// NumStrIntSeparator objects define the integer separation
// operation used in formatting number strings.
//
//        type NumStrIntSeparatorsDto struct {
//          intSeparators []NumStrIntSeparator
//        }
//
//        type NumStrIntSeparator struct {
//         intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//         intSeparatorGrouping    uint    // Number of integer digits in a group
//         intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                         // A zero value signals unlimited repetitions.
//         restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//        }
//
//        intSeparatorChars          []rune
//           - A series of runes or characters used to separate integer
//             digits in a number string. These characters are commonly
//             known as the 'thousands separator'. A 'thousands
//             separator' is used to separate groups of integer digits to
//             the left of the decimal separator (a.k.a. decimal point).
//             In the United States, the standard integer digits
//             separator is the single comma character (','). Other
//             countries and cultures use periods, spaces, apostrophes or
//             multiple characters to separate integers.
//                   United States Example:  1,000,000,000
//
//        intSeparatorGrouping       uint
//           - This unsigned integer values specifies the number of
//             integer digits within a group. This value is used to group
//             integers within a number string.
//
//             In most western countries integer digits to the left of
//             the decimal separator (a.k.a. decimal point) are separated
//             into groups of three digits representing a grouping of
//             'thousands' like this: '1,000,000,000'. In this case the
//             intSeparatorGrouping value would be set to three ('3').
//
//             In some countries and cultures other integer groupings
//             are used. In India, for example, a number might be
//             formatted like this: '6,78,90,00,00,00,00,000'. Chinese
//             Numerals have an integer grouping value of four ('4').
//                 Chinese Numerals Example: '12,3456,7890,2345'
//
//        intSeparatorRepetitions    uint
//           - This unsigned integer value specifies the number of times
//             this integer grouping is repeated. A value of zero signals
//             that this integer grouping will be repeated indefinitely.
//
//        restartIntGroupingSequence bool
//           - If the NumStrIntSeparator is the last element in an array
//             of NumStrIntSeparator objects, this boolean flag signals
//             whether the entire integer grouping sequence will be
//             restarted from array element zero.
//
//
//
//  ePrefix                    *ErrPrefixDto
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
func (fmtCurrMech *formatterCurrencyMechanics) setNumericSeparators(
	formatterCurrency *FormatterCurrency,
	numericSeparators NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if fmtCurrMech.lock == nil {
		fmtCurrMech.lock = new(sync.Mutex)
	}

	fmtCurrMech.lock.Lock()

	defer fmtCurrMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyMechanics." +
			"setNumericSeparators()")

	if formatterCurrency == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer\n",
			ePrefix.String())
	}

	return formatterCurrency.numericSeparators.CopyIn(
		&numericSeparators,
		ePrefix)
}
