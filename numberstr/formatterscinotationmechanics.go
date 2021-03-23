package numberstr

import (
	"fmt"
	"sync"
)

type formatterSciNotationMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// formatterSciNotationMechanics.
//
func (fmtSciNotationMech formatterSciNotationMechanics) ptr() *formatterSciNotationMechanics {

	if fmtSciNotationMech.lock == nil {
		fmtSciNotationMech.lock = new(sync.Mutex)
	}

	fmtSciNotationMech.lock.Lock()

	defer fmtSciNotationMech.lock.Unlock()

	newFmtSciNotMechanics :=
		new(formatterSciNotationMechanics)

	newFmtSciNotMechanics.lock = new(sync.Mutex)

	return newFmtSciNotMechanics
}

// setFmtSciNotWithComponents - Transfers new data to an instance of
// FormatterSciNotation. After completion, all the data
// fields within input parameter 'nStrFmtSpecSciNotDto' will be
// overwritten.
//
// Scientific Notation Format Specification objects encapsulate the
// format specifications used in formatting scientific notation
// numeric values for text display.
//
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
//  numFieldDto                   NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto object contains specifications for number
//       field length. Typically, the length of a number field is
//       greater than the length of the number string, or scientific
//       notation string, which will be inserted and displayed within
//       the number field.
//
//       The NumberFieldDto object also contains specifications
//       for positioning or alignment of the number string within
//       the number field. This alignment dynamic is described as
//       text justification. The member variable '
//       NumberFieldDto.textJustifyFormat' is used to specify one
//       of three possible alignment formats. One of these formats
//       will be selected to control the alignment of the number
//       string within the number field. These optional alignment
//       formats are shown below with examples:
//
//       (1) 'Right-Justification' - "       NumberString"
//       (2) 'Left-Justification' - "NumberString        "
//       (3) 'Centered'           - "    NumberString    "
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
func (fmtSciNotationMech *formatterSciNotationMechanics) setFmtSciNotWithComponents(
	formatterSciNotation *FormatterSciNotation,
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotationMech.lock == nil {
		fmtSciNotationMech.lock = new(sync.Mutex)
	}

	fmtSciNotationMech.lock.Lock()

	defer fmtSciNotationMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterSciNotationMechanics.setFmtSciNotWithComponents()")

	if formatterSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	newFmtSciNotation := FormatterSciNotation{}

	newFmtSciNotation.numStrFmtType =
		NumStrFormatTypeCode(0).ScientificNotation()

	newFmtSciNotation.significandUsesLeadingPlus =
		significandUsesLeadingPlus

	newFmtSciNotation.mantissaLength =
		mantissaLength

	newFmtSciNotation.exponentChar =
		exponentChar

	newFmtSciNotation.exponentUsesLeadingPlus =
		exponentUsesLeadingPlus

	err = newFmtSciNotation.numFieldDto.CopyIn(
		&numFieldDto,
		ePrefix.XCtx(
			"numFieldDto->formatterSciNotation.numFieldDto"))

	if err != nil {
		return err
	}

	err = formatterSciNotationElectron{}.ptr().copyIn(
		formatterSciNotation,
		&newFmtSciNotation,
		ePrefix.XCtx(
			"newFmtSciNotation->formatterSciNotation"))

	return err
}
