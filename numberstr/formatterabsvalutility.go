package numberstr

import (
	"fmt"
	"sync"
)

type formatterAbsoluteValueUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterAbsoluteValueUtility.
func (fmtAbsValUtil formatterAbsoluteValueUtility) ptr() *formatterAbsoluteValueUtility {

	if fmtAbsValUtil.lock == nil {
		fmtAbsValUtil.lock = new(sync.Mutex)
	}

	fmtAbsValUtil.lock.Lock()

	defer fmtAbsValUtil.lock.Unlock()

	fmtUtility := new(formatterAbsoluteValueUtility)

	fmtUtility.lock = new(sync.Mutex)

	return fmtUtility
}

// setBasic() - Sets the data values for an instance of
// FormatterAbsoluteValue passed as an input parameter.
//
// This method differs from formatterAbsoluteValueUtility.SetBasicRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// The FormatterAbsoluteValue type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsoluteValue              *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All of the data values in this object will be overwritten
//       and set to new values based on the following input
//       parameters.
//
//
//  decimalSeparatorChars         string
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars       string
//     - The character or characters which will be used to delimit
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
//       FormatterAbsoluteValue.SetWithComponents().
//
//
//  turnOnThousandsSeparator      bool
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
//       are listed as follows:
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
//               "127.54" THE DEFAULT Absolute Value Format
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
func (fmtAbsValUtil *formatterAbsoluteValueUtility) setBasic(
	fmtAbsoluteValue *FormatterAbsoluteValue,
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValUtil.lock == nil {
		fmtAbsValUtil.lock = new(sync.Mutex)
	}

	fmtAbsValUtil.lock.Lock()

	defer fmtAbsValUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueUtility." +
			"setBasic()")

	if fmtAbsoluteValue == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteValue' is invalid!\n"+
			"'fmtAbsoluteValue' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	if len(decimalSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorChars' is invalid!\n"+
			"'decimalSeparatorChars' is a zero length string.\n",
			ePrefix.String())
		return err
	}

	if len(thousandsSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparatorChars' is invalid!\n"+
			"'thousandsSeparatorChars' is a zero length string.\n",
			ePrefix.String())
		return err
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

	var newNumericSeparators NumericSeparators

	newNumericSeparators,
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

	err = formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			fmtAbsoluteValue,
			absoluteValFmt,
			turnOnIntegerDigitsSeparation,
			newNumericSeparators,
			numFieldDto,
			ePrefix)

	return err
}

// setBasicRunes() - Sets the data values for an instance of
// FormatterAbsoluteValue passed as an input parameter.
//
// The FormatterAbsoluteValue type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsoluteValue              *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All of the data values in this object will be overwritten
//       and set to new values based on the following input
//       parameters.
//
//
//  decimalSeparatorChars         []rune
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars       []rune
//     - The character or characters which will be used to delimit
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
//       FormatterAbsoluteValue.SetWithComponents().
//
//
//  turnOnThousandsSeparator      bool
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
//       are listed as follows:
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
//               "127.54" THE DEFAULT Absolute Value Format
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
func (fmtAbsValUtil *formatterAbsoluteValueUtility) setBasicRunes(
	fmtAbsoluteValue *FormatterAbsoluteValue,
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValUtil.lock == nil {
		fmtAbsValUtil.lock = new(sync.Mutex)
	}

	fmtAbsValUtil.lock.Lock()

	defer fmtAbsValUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueUtility." +
			"setBasicRunes()")

	if fmtAbsoluteValue == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteValue' is invalid!\n"+
			"'fmtAbsoluteValue' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	if decimalSeparatorChars == nil {
		decimalSeparatorChars =
			make([]rune, 0)
	}

	if len(decimalSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorChars' is invalid!\n"+
			"'decimalSeparatorChars' is a zero length rune array.\n",
			ePrefix.String())
		return err
	}

	if thousandsSeparatorChars == nil {
		thousandsSeparatorChars =
			make([]rune, 0)
	}

	if len(thousandsSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparatorChars' is invalid!\n"+
			"'thousandsSeparatorChars' is a zero length rune array.\n",
			ePrefix.String())
		return err
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

	var newNumericSeparators NumericSeparators

	newNumericSeparators,
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

	err = formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			fmtAbsoluteValue,
			absoluteValFmt,
			turnOnIntegerDigitsSeparation,
			newNumericSeparators,
			numFieldDto,
			ePrefix)

	return err
}

// setDetail() - This method will set all of the member variable
// data values for the FormatterAbsoluteValue input parameter,
// 'fmtAbsoluteValue'. New data values will be generated from the
// input parameters described below.
//
// This method differs from formatterAbsoluteValueUtility.SetDetailRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// The FormatterAbsoluteValue type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsoluteValue              *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All of the data values in this object will be overwritten
//       and set to new values based on the following input
//       parameters.
//
//
//  decimalSeparatorChars         string
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period (".") or Decimal Point.
//           United States Example: '123.45678'
//
//
//  integerDigitsSeparators       string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//       For custom integer digit grouping, use method
//       FormatterAbsoluteValue.SetWithComponents().
//
//
//  intSeparatorGrouping       uint
//     - The number of integer digits in group to be separated by
//       separator characters. The most common grouping is the
//       thousands grouping consisting of 3-digits. United States
//       Example:
//               1,000,000,000
//
//       Other countries and cultures use different grouping sequences.
//       used to separate integer digits in a number string.
//       Indian Number System Example: 6,78,90,00,00,00,00,000
//       Chinese Numeral Example:      6789,0000,0000,0000
//
//
//  intSeparatorRepetitions    uint
//     - Number of times this character/group sequence is repeated.
//       A zero value signals unlimited repetitions.
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
//       are listed as follows:
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
//               "127.54" THE DEFAULT Absolute Value Format
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
func (fmtAbsValUtil *formatterAbsoluteValueUtility) setDetail(
	fmtAbsoluteValue *FormatterAbsoluteValue,
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValUtil.lock == nil {
		fmtAbsValUtil.lock = new(sync.Mutex)
	}

	fmtAbsValUtil.lock.Lock()

	defer fmtAbsValUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueUtility." +
			"setDetail()")

	if fmtAbsoluteValue == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteValue' is invalid!\n"+
			"'fmtAbsoluteValue' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	if len(decimalSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorChars' is invalid!\n"+
			"'decimalSeparatorChars' is a zero length string.\n",
			ePrefix.String())
		return err
	}

	if len(integerDigitsSeparators) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is a zero length string.\n",
			ePrefix.String())
		return err
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

	var newNumericSeparators NumericSeparators

	newNumericSeparators,
		err = NumericSeparators{}.NewIntSepsDetail(
		decimalSeparatorChars,
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			fmt.Sprintf("\n"+
				"decimalSeparatorChars='%v'\n"+
				"integerDigitsSeparators='%v'\n",
				decimalSeparatorChars,
				integerDigitsSeparators)))

	if err != nil {
		return err
	}

	err = formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			fmtAbsoluteValue,
			absoluteValFmt,
			turnOnIntegerDigitsSeparation,
			newNumericSeparators,
			numFieldDto,
			ePrefix)

	return err
}

// setDetailRunes() - This method will set all of the member variable
// data values for the FormatterAbsoluteValue input parameter,
// 'fmtAbsoluteValue'. New data values will be generated from the
// input parameters described below.
//
// This method differs from formatterAbsoluteValueUtility.SetDetail()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// The FormatterAbsoluteValue type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsoluteValue              *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All of the data values in this object will be overwritten
//       and set to new values based on the following input
//       parameters.
//
//
//  decimalSeparatorChars         []rune
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period (".") or Decimal Point.
//           United States Example: '123.45678'
//
//
//  integerDigitsSeparators       []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//       For custom integer digit grouping, use method
//       FormatterAbsoluteValue.SetWithComponents().
//
//
//  intSeparatorGrouping       uint
//     - The number of integer digits in group to be separated by
//       separator characters. The most common grouping is the
//       thousands grouping consisting of 3-digits. United States
//       Example:
//               1,000,000,000
//
//       Other countries and cultures use different grouping sequences.
//       used to separate integer digits in a number string.
//       Indian Number System Example: 6,78,90,00,00,00,00,000
//       Chinese Numeral Example:      6789,0000,0000,0000
//
//
//  intSeparatorRepetitions    uint
//     - Number of times this character/group sequence is repeated.
//       A zero value signals unlimited repetitions.
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
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
//       Valid format strings for absolute value number strings
//       are listed as follows:
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
//               "127.54" THE DEFAULT Absolute Value Format
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
func (fmtAbsValUtil *formatterAbsoluteValueUtility) setDetailRunes(
	fmtAbsoluteValue *FormatterAbsoluteValue,
	decimalSeparatorChars []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValUtil.lock == nil {
		fmtAbsValUtil.lock = new(sync.Mutex)
	}

	fmtAbsValUtil.lock.Lock()

	defer fmtAbsValUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueUtility." +
			"setDetailRunes()")

	if fmtAbsoluteValue == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtAbsoluteValue' is invalid!\n"+
			"'fmtAbsoluteValue' is a 'nil' pointer.\n",
			ePrefix.String())
		return err
	}

	if len(decimalSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorChars' is invalid!\n"+
			"'decimalSeparatorChars' is a zero length string.\n",
			ePrefix.String())
		return err
	}

	if len(integerDigitsSeparators) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is a zero length string.\n",
			ePrefix.String())
		return err
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

	var newNumericSeparators NumericSeparators

	newNumericSeparators,
		err = NumericSeparators{}.NewIntSepsDetailRunes(
		decimalSeparatorChars,
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			fmt.Sprintf("\n"+
				"decimalSeparatorChars='%v'\n"+
				"integerDigitsSeparators='%v'\n",
				decimalSeparatorChars,
				integerDigitsSeparators)))

	if err != nil {
		return err
	}

	err = formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			fmtAbsoluteValue,
			absoluteValFmt,
			turnOnIntegerDigitsSeparation,
			newNumericSeparators,
			numFieldDto,
			ePrefix)

	return err
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterAbsoluteValue instance
// to United States Default values.
//
// In the United States, Absolute Value default formatting
// parameters are defined as follows:
//
//    Absolute Value Number format: "127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtSpecAbsValDto          *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       All data values in this object will be overwritten and
//       set to United States default values for absolute numeric
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
func (fmtAbsValUtil *formatterAbsoluteValueUtility) setToUnitedStatesDefaults(
	nStrFmtSpecAbsValDto *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsValUtil.lock == nil {
		fmtAbsValUtil.lock = new(sync.Mutex)
	}

	fmtAbsValUtil.lock.Lock()

	defer fmtAbsValUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterAbsoluteValueUtility." +
			"setToUnitedStatesDefaults()")

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is invalid!\n"+
			"'nStrFmtSpecAbsValDto' is a 'nil' pointer.\n",
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

	err = formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			nStrFmtSpecAbsValDto,
			"127.54",
			true,
			numericSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"nStrFmtSpecAbsValDto"))

	return err
}
