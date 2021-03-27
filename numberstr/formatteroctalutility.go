package numberstr

import (
	"fmt"
	"sync"
)

type formatterOctalUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterOctalUtility.
//
func (fmtOctalUtil formatterOctalUtility) ptr() *formatterOctalUtility {

	if fmtOctalUtil.lock == nil {
		fmtOctalUtil.lock = new(sync.Mutex)
	}

	fmtOctalUtil.lock.Lock()

	defer fmtOctalUtil.lock.Unlock()

	newFmtOctalUtil :=
		new(formatterOctalUtility)

	newFmtOctalUtil.lock = new(sync.Mutex)

	return newFmtOctalUtil
}

// setFmtOctalDetail - This method will set all of the member
// variable data values for the FormatterOctal input parameter,
// 'formatterOctal'. New data values will be generated from the
// input parameters described below.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// This method differs from formatterOctalUtility.setFmtOctalDetailRunes()
// in that this method accepts a string for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterHexadecimal, reference method:
//   'FormatterHexadecimal.SetWithComponents()'
//
// The member variable 'FormatterOctal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Octal()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterOctal                *FormatterOctal
//     - A pointer to an instance of FormatterOctal. Member
//       variable data values within this instance will be
//       overwritten and reset to new values based on the input
//       parameters listed below.
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
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
func (fmtOctalUtil *formatterOctalUtility) setFmtOctalDetail(
	formatterOctal *FormatterOctal,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	leftPrefix string,
	rightSuffix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (err error) {

	if fmtOctalUtil.lock == nil {
		fmtOctalUtil.lock = new(sync.Mutex)
	}

	fmtOctalUtil.lock.Lock()

	defer fmtOctalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalUtility." +
			"setFmtOctalDetail()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterOctal.lock == nil {
		formatterOctal.lock = new(sync.Mutex)
	}

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

	return formatterOctalMechanics{}.ptr().
		setFmtOctalWithComponents(
			formatterOctal,
			leftPrefix,
			rightSuffix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterOctal"))
}

// setFmtOctalDetailRunes - This method will set all of the member
// variable data values for the FormatterOctal input parameter,
// 'formatterOctal'. New data values will be generated from the
// input parameters described below.
//
// The FormatterOctal type encapsulates the formatting parameters
// necessary to format octal digits for display in text number
// strings.
//
// This method differs from formatterOctalUtility.setFmtOctalDetail()
// in that this method accepts a rune array for input parameter,
// 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterOctal, reference method:
//   'FormatterOctal.SetWithComponents()'
//
// The member variable 'FormatterOctal.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Octal()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterOctal                *FormatterOctal
//     - A pointer to an instance of FormatterOctal. Member
//       variable data values within this instance will be
//       overwritten and reset to new values based on the input
//       parameters listed below.
//
//
//  integerDigitsSeparators       []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (',').
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
//  leftPrefix                    string
//     - Sets the 'Left Prefix' string for the new returned
//       instance of FormatterOctal.
//
//       'Left Prefix' string characters are appended or
//        concatenated to the beginning of octal digits formatted
//        in a text number string.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are prefixed with leading characters such as the
//       digit zero ('0'), the letters 'o' or 'q', the digit–letter
//       combination zero-'o' ('0o'), the ampersand symbol ('&'),
//       the dollar sign symbol ('$'), the At Sign ('@') or the
//       back slash symbol ('\').
//
//       The 'Left Prefix' parameter allows for custom
//       configuration of these leading prefix characters.
//
//       An empty 'leftPrefix' string signals that no characters
//       will be prefixed to the beginning of octal number strings.
//
//       Reference:
//          https://en.wikipedia.org/wiki/Octal
//
//
//  rightSuffix                   string
//     - Sets the 'Right Suffix' string for the new returned
//       instance of FormatterOctal. 'rightSuffix' specifies the
//       characters which will be appended to the end of octal
//       number strings.
//
//       Occasionally, octal numeric digits displayed in a text
//       string are suffixed with trailing characters such as the
//       the upper and lower case letters ('o','O', 'q' or 'Q').
//
//       The 'Right Suffix' parameter allows for custom
//       configuration of these trailing suffix characters.
//
//       An empty 'rightSuffix' string signals that no characters
//       will be appended to the end of octal number strings.
//
//       Reference:
//         https://en.wikipedia.org/wiki/Octal
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
func (fmtOctalUtil *formatterOctalUtility) setFmtOctalDetailRunes(
	formatterOctal *FormatterOctal,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	leftPrefix string,
	rightSuffix string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (err error) {

	if fmtOctalUtil.lock == nil {
		fmtOctalUtil.lock = new(sync.Mutex)
	}

	fmtOctalUtil.lock.Lock()

	defer fmtOctalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalUtility." +
			"setFmtOctalDetailRunes()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterOctal.lock == nil {
		formatterOctal.lock = new(sync.Mutex)
	}

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

	return formatterOctalMechanics{}.ptr().
		setFmtOctalWithComponents(
			formatterOctal,
			leftPrefix,
			rightSuffix,
			turnOnIntegerDigitsSeparation,
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterOctal"))
}

// setFmtOctalWithDefaults - Overwrites and resets the data fields
// within input parameter 'formatterOctal'. The new data values are
// created from default parameters specified below.
//
// Member variable data fields in the input parameter
// FormatterOctal instance 'formatterOctal' are configured as
// follows:
//
// Number String Format Type (numStrFmtType) =
//     NumStrFormatTypeCode(0).Octal()
//
// Octal Number String Left Prefix String (leftPrefix) = ""
//
// Octal Number String Right Suffix String (rightSuffix) = ""
//
// Turn On Integer Digits Separation
//  (turnOnIntegerDigitsSeparation) = false
//
// Integer Separators (integerSeparators) =
//   Integer Digit Separator  = ","
//   Integer Digit Grouping Sequence = 3
//   Integer Group Repetitions = 0
//
// Number Field Dto = Off
//   Requested Number Field Length = -1
//   Text Justify Format (textJustifyFormat) =
//     TextJustify(0).Right
//
func (fmtOctalUtil *formatterOctalUtility) setFmtOctalWithDefaults(
	formatterOctal *FormatterOctal,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtOctalUtil.lock == nil {
		fmtOctalUtil.lock = new(sync.Mutex)
	}

	fmtOctalUtil.lock.Lock()

	defer fmtOctalUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterOctalUtility." +
			"setFmtOctalWithDefaults()")

	if formatterOctal == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterOctal' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if formatterOctal.lock == nil {
		formatterOctal.lock = new(sync.Mutex)
	}

	var integerSeparators NumStrIntSeparatorsDto

	integerSeparators,
		err = NumStrIntSeparatorsDto{}.NewBasic(
		",",
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

	return formatterOctalMechanics{}.ptr().
		setFmtOctalWithComponents(
			formatterOctal,
			"",    // leftPrefix
			"",    // rightSuffix
			false, //turnOnIntegerDigitsSeparation
			integerSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"formatterOctal"))
}
