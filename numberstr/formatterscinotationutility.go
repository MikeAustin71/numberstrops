package numberstr

import (
	"fmt"
	"sync"
)

type formatterSciNotationUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterSciNotationUtility.
//
func (fmtSciNotUtil formatterSciNotationUtility) ptr() *formatterSciNotationUtility {

	if fmtSciNotUtil.lock == nil {
		fmtSciNotUtil.lock = new(sync.Mutex)
	}

	fmtSciNotUtil.lock.Lock()

	defer fmtSciNotUtil.lock.Unlock()

	newFmtSciNotUtility :=
		new(formatterSciNotationUtility)

	newFmtSciNotUtility.lock = new(sync.Mutex)

	return newFmtSciNotUtility
}

// setSciNotWithDefaults - Transfers new data to an instance of
// FormatterSciNotation. After completion, all the data
// fields within input parameter 'formatterSciNotation' will be
// overwritten.
//
// Scientific Notation Format Specification objects encapsulate the
// format specifications used in formatting scientific notation
// numeric values for text display.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterSciNotation          *FormatterSciNotation
//     - A pointer to a FormatterSciNotation object. All
//       of the data fields in this object will overwritten and set
//       to new values based on the following input parameters.
//
//
//  significandUsesLeadingPlus    bool
//     - "Significand uses leading plus sign". This refers to the
//       integer digit in a significand.
//
//       Positive significand integer digits may have a leading
//       plus sign, '+2.652E+8'. The default is no leading plus
//       sign, '2.652E+8'.
//
//       If this value is set to true, positive significand integer
//       digit values will be prefixed with a leading plus sign
//       ('+').
//               Example: '+2.652E+8'
//
//       If this value is set to true, positive significand integer
//       digit values will NOT be prefixed with a leading plus sign
//       ('+').
//               Example: '2.652E+8'
//
//
//  mantissaLength                uint
//     - This parameter sets the length of the mantissa used in the
//       scientific notation format instructions.
//
//       In scientific notation, the term 'mantissa' refers to the
//       fractional digits contained in the significand. In the
//       scientific notation example, '2.652E+8', the 'mantissa'
//       identifies the fractional digits, '.652'.
//
//       The input parameter, 'mantissaLength' controls the number
//       of fractional digits displayed in the mantissa.
//
//
//  exponentChar                  rune
//     - This parameter specifies the exponent character to be used
//       int the scientific notation format instructions.
//
//       In scientific notation example, '2.652E+8', the exponent
//       character is 'E'.  The character 'E' is used as the
//       default to avoid confusion with Euler's number 'e'.
//       However the character 'e' is often used in scientific
//       notation and may therefore be specified by the user.
//
//
//  exponentUsesLeadingPlus       bool
//     - This parameter signals whether a leading plus will be
//       included for positive exponent value in scientific
//       notation displays.
//
//       In scientific notation example, '2.652E8', the exponent
//       value is '8'.
//
//       If input parameter 'exponentUsesLeadingPlus' is set to
//       'true', the scientific notation text display will prefix
//       a leading plus sign ('+') for positive exponent values.
//             Example: '2.652E+8'
//
//       If input parameter 'exponentUsesLeadingPlus' is set to
//       'false', the scientific notation text display will NOT
//       include a leading plus sign ('+') for positive exponent
//       values.
//             Example: '2.652E8'
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
func (fmtSciNotUtil *formatterSciNotationUtility) setSciNotWithDefaults(
	formatterSciNotation *FormatterSciNotation,
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotUtil.lock == nil {
		fmtSciNotUtil.lock = new(sync.Mutex)
	}

	fmtSciNotUtil.lock.Lock()

	defer fmtSciNotUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationUtility.setSciNotWithDefaults()")

	if formatterSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
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

	err = formatterSciNotationMechanics{}.ptr().
		setFmtSciNotWithComponents(
			formatterSciNotation,
			significandUsesLeadingPlus,
			mantissaLength,
			exponentChar,
			exponentUsesLeadingPlus,
			numFieldDto,
			ePrefix.XCtx("formatterSciNotation"))

	return err
}

// setToUnitedStatesDefaults - Sets the member variable data
// values for the incoming FormatterSciNotation instance
// to United States Default values.
//
// In the United States, Scientific Notation default formatting
// parameters are defined as follows:
//
//   SignificandUsesLeadingPlus = false
//   MantissaLength             = 6
//   ExponentChar               = 'E'
//   ExponentUsesLeadingPlus    = true
//   NumFieldLen                = -1
//   NumFieldTextJustify        = TextJustify(0).Right()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  formatterSciNotation          *FormatterSciNotation
//     - A pointer to an instance of FormatterSciNotation.
//       All data values in this object will be overwritten and
//       set to United States default values for scientific
//       notation values displayed in number strings.
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
func (fmtSciNotUtil *formatterSciNotationUtility) setToUnitedStatesDefaults(
	formatterSciNotation *FormatterSciNotation,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotUtil.lock == nil {
		fmtSciNotUtil.lock = new(sync.Mutex)
	}

	fmtSciNotUtil.lock.Lock()

	defer fmtSciNotUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationUtility." +
			"setSciNotWithDefaults()")

	if formatterSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
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

	nStrFmtSpecSciNotDtoMech :=
		formatterSciNotationMechanics{}

	err = nStrFmtSpecSciNotDtoMech.setFmtSciNotWithComponents(
		formatterSciNotation,
		false,
		6,
		'E',
		true,
		numFieldDto,
		ePrefix.XCtx("formatterSciNotation"))

	return err
}
