package numberstr

import (
	"fmt"
	"sync"
)

type formatterCurrencyUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterCurrencyUtility.
//
func (fmtCurrUtil formatterCurrencyUtility) ptr() *formatterCurrencyUtility {

	if fmtCurrUtil.lock == nil {
		fmtCurrUtil.lock = new(sync.Mutex)
	}

	fmtCurrUtil.lock.Lock()

	defer fmtCurrUtil.lock.Unlock()

	newCurrencyUtility :=
		new(formatterCurrencyUtility)

	newCurrencyUtility.lock = new(sync.Mutex)

	return newCurrencyUtility
}

// setBasicFormatterCurrency - This method will set all of the
// member variable data values for the FormatterCurrency input
// parameter, 'formatterCurrency'. The other input parameters
// represent the minimum information required to configure a
// FormatterCurrency object. Default values are used to supplement
// these input parameters.
//
// This method differs from method
// formatterCurrencyUtility.setBasicRunesFormatterCurrency()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterCurrency, reference method:
//   'FormatterCurrency.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// The FormatterCurrency type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// input parameter 'formatterCurrency'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency             *FormatterCurrency,
//     - A pointer to an instance of FormatterCurrency.
//       All data values in this object will be overwritten and
//       set to new values based on the following input parameters.
//
//
//  decimalSeparatorChars      string
//     - The characters or character used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars    string
//     - The characters or character which will be used to delimit
//       'thousands' in integer number strings. In the United
//       States, the Thousands separator is the comma character
//       (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterCurrency.NewWithComponents().
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
//  requestedNumberFieldLen    int
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
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
func (fmtCurrUtil *formatterCurrencyUtility) setBasicFormatterCurrency(
	formatterCurrency *FormatterCurrency,
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnIntegerDigitsSeparation bool,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrUtil.lock == nil {
		fmtCurrUtil.lock = new(sync.Mutex)
	}

	fmtCurrUtil.lock.Lock()

	defer fmtCurrUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyUtility." +
			"setBasicFormatterCurrency()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	if formatterCurrency.lock == nil {
		formatterCurrency.lock = new(sync.Mutex)
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err = NumberFieldDto{}.NewBasic(
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)

	if err != nil {
		return err
	}

	var numericSeparators NumericSeparators

	numericSeparators,
		err = NumericSeparators{}.NewBasic(
		decimalSeparatorChars,
		thousandsSeparatorChars,
		ePrefix.XCtx(
			fmt.Sprintf("\n"+
				"decimalSeparatorChars='%v'\n"+
				"thousandsSeparatorChars='%v'\n",
				decimalSeparatorChars,
				thousandsSeparatorChars)))

	if err != nil {
		return err
	}

	err = formatterCurrencyMechanics{}.ptr().
		setFormatterCurrencyWithComponents(
			formatterCurrency,
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
			numericSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterCurrency"))

	return err
}

// setBasicRunesFormatterCurrency - This method will set all of the
// member variable data values for the FormatterCurrency input
// parameter, 'formatterCurrency'. The other input parameters
// represent the minimum information required to configure a
// FormatterCurrency object. Default values are used to supplement
// these input parameters.
//
// This method differs from method
// formatterCurrencyUtility.setBasicFormatterCurrency()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterCurrency, reference method:
//   'FormatterCurrency.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// The FormatterCurrency type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// input parameter 'formatterCurrency'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency             *FormatterCurrency,
//     - A pointer to an instance of FormatterCurrency.
//       All data values in this object will be overwritten and
//       set to new values based on the following input parameters.
//
//
//  decimalSeparatorChars      []rune
//     - The characters or character used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars    []rune
//     - The characters or character which will be used to delimit
//       'thousands' in integer number strings. In the United
//       States, the Thousands separator is the comma character
//       (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterCurrency.NewWithComponents().
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
//  requestedNumberFieldLen    int
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
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
func (fmtCurrUtil *formatterCurrencyUtility) setBasicRunesFormatterCurrency(
	formatterCurrency *FormatterCurrency,
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnIntegerDigitsSeparation bool,
	positiveValueFmt string,
	negativeValueFmt string,
	decimalDigits uint,
	currencyCode string,
	currencyCodeNo string,
	currencyName string,
	currencySymbols []rune,
	minorCurrencyName string,
	minorCurrencySymbols []rune,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrUtil.lock == nil {
		fmtCurrUtil.lock = new(sync.Mutex)
	}

	fmtCurrUtil.lock.Lock()

	defer fmtCurrUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyUtility." +
			"setBasicRunesFormatterCurrency()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	if formatterCurrency.lock == nil {
		formatterCurrency.lock = new(sync.Mutex)
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err = NumberFieldDto{}.NewBasic(
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)

	if err != nil {
		return err
	}

	var numericSeparators NumericSeparators

	numericSeparators,
		err = NumericSeparators{}.NewBasicRunes(
		decimalSeparatorChars,
		thousandsSeparatorChars,
		ePrefix.XCtx(
			fmt.Sprintf("\n"+
				"decimalSeparatorChars='%v'\n"+
				"thousandsSeparatorChars='%v'\n",
				decimalSeparatorChars,
				thousandsSeparatorChars)))

	if err != nil {
		return err
	}

	err = formatterCurrencyMechanics{}.ptr().
		setFormatterCurrencyWithComponents(
			formatterCurrency,
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
			numericSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterCurrency"))

	return err
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterSignedNumber instance
// to United States Default values.
//
// In the United States, Currency Number default formatting
// parameters are defined as follows:
//
//   Currency Positive Number Format: "$127.54"
//   Currency Negative Number Format: "($127.54)"
//           Currency Decimal Digits: 2
//                     Currency Code: "USD"
//                    CurrencyCodeNo: "840"
//                     Currency Name: "Dollar"
//                  Currency Symbols: []rune{'\U00000024'}
//               Minor Currency Name: "Cent"
//            Minor Currency Symbols: []rune{'\U000000a2'}
//       Decimal Separator Character: '.'
//     Thousands Separator Character: ','
//       Turn On Thousands Separator: true
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterCurrency             *FormatterCurrency,
//     - A pointer to an instance of FormatterCurrency.
//       All data values in this object will be overwritten and
//       set to United States default values for currency number
//       values displayed in number strings.
//
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
func (fmtCurrUtil *formatterCurrencyUtility) setToUnitedStatesDefaults(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrUtil.lock == nil {
		fmtCurrUtil.lock = new(sync.Mutex)
	}

	fmtCurrUtil.lock.Lock()

	defer fmtCurrUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyUtility." +
			"setToUnitedStatesDefaults()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is invalid!\n"+
			"'formatterCurrency' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err = NumberFieldDto{}.NewBasic(
		-1,
		TextJustify(0).Right(),
		ePrefix)

	if err != nil {
		return err
	}

	var numericSeparators NumericSeparators

	numericSeparators,
		err = NumericSeparators{}.
		NewUnitedStatesDefaults(
			ePrefix.XCtxEmpty())

	if err != nil {
		return err
	}

	err = formatterCurrencyMechanics{}.ptr().
		setFormatterCurrencyWithComponents(
			formatterCurrency,
			"$127.54",
			"($127.54)",
			2,
			"USD",
			"840",
			"Dollar",
			[]rune{'\U00000024'},
			"Cent",
			[]rune{'\U000000a2'},
			true,
			numericSeparators,
			numFieldDto,
			ePrefix)

	return err
}
