package numberstr

import (
	"fmt"
	"sync"
)

type formatterHexadecimalUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterHexadecimalUtility.
//
func (fmtHexadecimalUtil formatterHexadecimalUtility) ptr() *formatterHexadecimalUtility {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	newFmtHexadecimalUtility :=
		new(formatterHexadecimalUtility)

	newFmtHexadecimalUtility.lock = new(sync.Mutex)

	return newFmtHexadecimalUtility
}

// setFmtHexDetail - This method will set all of the member
// variable data values for the FormatterHexadecimal input
// parameter, 'formatterHex'. New data values will be generated
// from the input parameters described below.
//
// The FormatterHexadecimal type encapsulates the formatting
// parameters necessary to format hexadecimal digits for display in
// text number strings.
//
// This method differs from formatterHexadecimalUtility.setFmtHexDetailRunes()
// in that this method accepts a string for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterHexadecimal, reference method:
//   'FormatterHexadecimal.SetWithComponents()'
//
// The member variable 'FormatterHexadecimal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Hexadecimal()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterHex                  *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal. All the
//       member data variable data values in this instance will
//       overwritten and reset according to the following input
//       parameters.
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
//       FormatterCurrency.NewWithComponents().
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
//  useUpperCaseLetters           bool
//     - Sets the 'Use Upper Case Letters' flag. This boolean flag
//       determines whether Hexadecimal Digits 'A' through 'F' will
//       be expressed as Upper Case Letters or Lower Case Letters
//       ('a' through 'f').
//
//       If input parameter 'useUpperCaseLetters' is set to 'true',
//       hexadecimal alphabetic digits will be expressed as upper
//       case letters. Conversely, if the parameter is 'false',
//       hexadecimal alphabetic digits will be expressed as lower
//       case letters.
//
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string which is concatenated to the
//       beginning of hexadecimal digits formatted in a number
//       string.
//
//       Often, hexadecimal digits displayed in a text string are
//       prefixed with the characters, "0x". Example:  '0x14AF'
//
//       Input parameter 'leftPrefix' controls this prefix string.
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
func (fmtHexadecimalUtil *formatterHexadecimalUtility) setFmtHexDetail(
	formatterHex *FormatterHexadecimal,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	useUpperCaseLetters bool,
	leftPrefix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalUtility." +
			"setFmtHexDetail()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewDetail(
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}
	var numFieldDto NumberFieldDto

	numFieldDto,
		err =
		NumberFieldDto{}.NewBasic(
			requestedNumberFieldLen,
			numberFieldTextJustify,
			nil)

	if err != nil {
		return err
	}

	return formatterHexadecimalMechanics{}.ptr().
		setFmtHexadecimalWithComponents(
			formatterHex,
			useUpperCaseLetters,
			leftPrefix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterHex"))
}

// setFmtHexDetailRunes - This method will overwrite and set all of
// the member variable data values for the FormatterHexadecimal
// input parameter, 'formatterHex'. New data values will be
// generated from the input parameters described below.
//
// The FormatterHexadecimal type encapsulates the formatting
// parameters necessary to format hexadecimal digits for display in
// text number strings.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterHexadecimal, reference method:
//   'FormatterHexadecimal.SetWithComponents()'
//
// This method differs from formatterHexadecimalUtility.setFmtHexDetail()
// in that this method accepts a rune array for input parameter,
// 'integerDigitsSeparators'.
//
// The member variable 'FormatterHexadecimal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Hexadecimal()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterHex                  *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal. All the
//       member data variable data values in this instance will
//       overwritten and reset according to the following input
//       parameters.
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
//       FormatterCurrency.NewWithComponents().
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
//  useUpperCaseLetters           bool
//     - Sets the 'Use Upper Case Letters' flag. This boolean flag
//       determines whether Hexadecimal Digits 'A' through 'F' will
//       be expressed as Upper Case Letters or Lower Case Letters
//       ('a' through 'f').
//
//       If input parameter 'useUpperCaseLetters' is set to 'true',
//       hexadecimal alphabetic digits will be expressed as upper
//       case letters. Conversely, if the parameter is 'false',
//       hexadecimal alphabetic digits will be expressed as lower
//       case letters.
//
//
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string which is concatenated to the
//       beginning of hexadecimal digits formatted in a number
//       string.
//
//       Often, hexadecimal digits displayed in a text string are
//       prefixed with the characters, "0x". Example:  '0x14AF'
//
//       Input parameter 'leftPrefix' controls this prefix string.
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
func (fmtHexadecimalUtil *formatterHexadecimalUtility) setFmtHexDetailRunes(
	formatterHex *FormatterHexadecimal,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	useUpperCaseLetters bool,
	leftPrefix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalUtility." +
			"setFmtHexDetailRunes()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewDetailRunes(
		integerDigitsSeparators,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		ePrefix.XCtx(
			"integerSeparators"))

	if err != nil {
		return err
	}
	var numFieldDto NumberFieldDto

	numFieldDto,
		err =
		NumberFieldDto{}.NewBasic(
			requestedNumberFieldLen,
			numberFieldTextJustify,
			nil)

	if err != nil {
		return err
	}

	return formatterHexadecimalMechanics{}.ptr().
		setFmtHexadecimalWithComponents(
			formatterHex,
			useUpperCaseLetters,
			leftPrefix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterHex"))
}

// setFmtHexWithDefaults - Receives a pointer to an
// instance of FormatterHexadecimal and proceeds to set the data
// values to standard defaults.
//
// The FormatterHexadecimal type encapsulates the formatting
// parameters necessary to format hexadecimal digits for display in
// text number strings.
//
// Member variable data fields within input parameter
// 'formatterHex' are configured as follows:
//
//  numStrFmtType = NumStrFormatTypeCode(0).Hexadecimal()
//
//  useUpperCaseLetters = false (Alphabetic letters will be
//                              displayed as 'a' through 'f')
//
//  leftPrefix = "0x" All hexadecimal number strings will be
//                    prefixed with "0x". Example: '0x2b'
//
//  turnOnIntegerDigitsSeparation = false No integer digit
//                                        separation will be
//                                        applied.
//
//  integerSeparators = Set to blank space (" ")
//
//  numFieldDto.requestedNumFieldLength = -1
//                   Number Field = Number String Length
//
//  numFieldDto.textJustifyFormat = TextJustify(0).Right()
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterHexadecimal, reference method:
//   'FormatterHexadecimal.SetWithComponents()'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterHex                  *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal. All the
//       member data variable data values in this instance will
//       overwritten and reset with default values described above.
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
func (fmtHexadecimalUtil *formatterHexadecimalUtility) setFmtHexWithDefaults(
	formatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtHexadecimalUtil.lock == nil {
		fmtHexadecimalUtil.lock = new(sync.Mutex)
	}

	fmtHexadecimalUtil.lock.Lock()

	defer fmtHexadecimalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterHexadecimalUtility." +
			"setFmtHexWithDefaults()")

	if formatterHex == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterHex' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewBasic(
		" ",
		nil)

	if err != nil {
		return err
	}

	var numFieldDto NumberFieldDto

	numFieldDto,
		err =
		NumberFieldDto{}.NewBasic(
			-1,
			TextJustify(0).Right(),
			nil)

	if err != nil {
		return err
	}

	return formatterHexadecimalMechanics{}.ptr().
		setFmtHexadecimalWithComponents(
			formatterHex,
			false, // useUpperCaseLetters
			"0x",  // leftPrefix
			false, // turnOnIntegerDigitsSeparation
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterHex"))
}
