package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSciNotationDtoMechanics struct {
	lock *sync.Mutex
}

// setSciNotationDto - Transfers new data to an instance of
// NumStrFmtSpecSciNotationDto. After completion, all the data
// fields within input parameter 'nStrFmtSpecSciNotDto' will be
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
//  nStrFmtSpecSciNotDto          *NumStrFmtSpecSciNotationDto
//     - A pointer to a NumStrFmtSpecSciNotationDto object. All
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
//  numFieldDto                NumberFieldDto
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
func (nStrFmtSpecSciNotDtoMech *numStrFmtSpecSciNotationDtoMechanics) setSciNotationDto(
	nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto,
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	numFieldLenDto NumberFieldDto,
	ePrefix string) (
	err error) {

	if nStrFmtSpecSciNotDtoMech.lock == nil {
		nStrFmtSpecSciNotDtoMech.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDtoMech.lock.Lock()

	defer nStrFmtSpecSciNotDtoMech.lock.Unlock()

	ePrefix += "numStrFmtSpecSciNotationDtoMechanics.setSciNotationDto() "

	if nStrFmtSpecSciNotDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
			ePrefix)

		return err
	}

	nStrFmtSpecSciNotDto.significandUsesLeadingPlus =
		significandUsesLeadingPlus

	nStrFmtSpecSciNotDto.mantissaLength =
		mantissaLength

	nStrFmtSpecSciNotDto.exponentChar =
		exponentChar

	nStrFmtSpecSciNotDto.exponentUsesLeadingPlus =
		exponentUsesLeadingPlus

	return nStrFmtSpecSciNotDto.numFieldLenDto.CopyIn(
		&numFieldLenDto,
		ePrefix+
			"numFieldLenDto->nStrFmtSpecSciNotDto.numFieldLenDto\n")

}
