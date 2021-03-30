package numberstr

import (
	"fmt"
	"sync"
)

type formatterBinaryUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterBinaryUtility.
//
func (fmtBinaryUtil formatterBinaryUtility) ptr() *formatterBinaryUtility {

	if fmtBinaryUtil.lock == nil {
		fmtBinaryUtil.lock = new(sync.Mutex)
	}

	fmtBinaryUtil.lock.Lock()

	defer fmtBinaryUtil.lock.Unlock()

	newBinaryUtility :=
		new(formatterBinaryUtility)

	newBinaryUtility.lock = new(sync.Mutex)

	return newBinaryUtility
}

// setDetail - This method will set all of the member variable data
// values for the FormatterBinary input parameter,
// 'formatterBinary'. New data values will be generated from the
// input parameters described below.
//
// This method differs from formatterBinaryUtility.setDetailRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.SetWithComponents()'
//
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
// The FormatterBinary type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// input parameter 'formatterBinary'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterBinary               *FormatterBinary,
//     - A pointer to an instance of FormatterBinary. All data
//       values in this object will be overwritten and set to new
//       values based on the following input parameters.
//
//
//  decimalSeparatorChars         string
//     - The characters or character used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
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
//       FormatterBinary.NewWithComponents().
//
//
//  intSeparatorGrouping          uint
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
//  intSeparatorRepetitions       uint
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
func (fmtBinaryUtil *formatterBinaryUtility) setDetail(
	formatterBinary *FormatterBinary,
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryUtil.lock == nil {
		fmtBinaryUtil.lock = new(sync.Mutex)
	}

	fmtBinaryUtil.lock.Lock()

	defer fmtBinaryUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"setDetail()")

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
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
	var numericSeparators NumericSeparators

	numericSeparators,
		err = NumericSeparators{}.NewIntSepsDetail(
		decimalSeparatorChars,
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			fmt.Sprintf(
				"decimalSeparatorChars='%v'\n"+
					"integerDigitsSeparators='%v'\n",
				decimalSeparatorChars,
				integerDigitsSeparators)))

	if err != nil {
		return err
	}

	err = formatterBinaryMechanics{}.ptr().setWithComponents(
		formatterBinary,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"formatterBinary"))

	return err
}

// setDetailRunes - This method will set all of the member variable data
// values for the FormatterBinary input parameter,
// 'formatterBinary'. New data values will be generated from the
// input parameters described below.
//
// This method differs from formatterBinaryUtility.setDetail()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.SetWithComponents()'
//
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
// The FormatterBinary type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// input parameter 'formatterBinary'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterBinary               *FormatterBinary,
//     - A pointer to an instance of FormatterBinary. All data
//       values in this object will be overwritten and set to new
//       values based on the following input parameters.
//
//
//  decimalSeparatorChars         []rune
//     - The characters or character used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
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
//       FormatterBinary.NewWithComponents().
//
//
//  intSeparatorGrouping          uint
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
//  intSeparatorRepetitions       uint
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
func (fmtBinaryUtil *formatterBinaryUtility) setDetailRunes(
	formatterBinary *FormatterBinary,
	decimalSeparatorChars []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryUtil.lock == nil {
		fmtBinaryUtil.lock = new(sync.Mutex)
	}

	fmtBinaryUtil.lock.Lock()

	defer fmtBinaryUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"setDetailRunes()")

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
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
	var numericSeparators NumericSeparators

	numericSeparators,
		err = NumericSeparators{}.NewIntSepsDetailRunes(
		decimalSeparatorChars,
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			fmt.Sprintf(
				"decimalSeparatorChars='%v'\n"+
					"integerDigitsSeparators='%v'\n",
				string(decimalSeparatorChars),
				string(integerDigitsSeparators))))

	if err != nil {
		return err
	}

	err = formatterBinaryMechanics{}.ptr().setWithComponents(
		formatterBinary,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"formatterBinary"))

	return err
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterBinary instance to United
// States Default values.
//
// In the United States, Binary Number default formatting
// parameters are defined as follows:
//
//    Turn On Integer Digits Separation: false
//         Decimal Separator Character: '.'
//                 Number Field Length: -1
//          Number Field Justification: Right-Justified
//
// Note: With 'Turn On Integer Digits Separation' set to false,
// integer digit separation is not applied to binary digits
// displayed in number strings.
//
func (fmtBinaryUtil *formatterBinaryUtility) setToUnitedStatesDefaults(
	formatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinaryUtil.lock == nil {
		fmtBinaryUtil.lock = new(sync.Mutex)
	}

	fmtBinaryUtil.lock.Lock()

	defer fmtBinaryUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"setToUnitedStatesDefaults()")

	if formatterBinary == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterBinary' is invalid!\n"+
			"'formatterBinary' is a 'nil' pointer.\n",
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
		err = NumericSeparators{}.NewIntSepsDetailRunes(
		[]rune{'.'}, // Decimal Separators
		[]rune{','}, // Integer Digits Separators
		3,           // Integer Digit Grouping
		0,           // Integer Separator Repetitions
		ePrefix.XCtx(
			"decimalSeparatorChars='.'\n"+
				"integerDigitsSeparators=','\n"))

	if err != nil {
		return err
	}

	err = formatterBinaryMechanics{}.ptr().setWithComponents(
		formatterBinary,
		false, // turnOnIntegerDigitsSeparation
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"formatterBinary"))

	return err
}
