package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecSciNotationDtoUtility struct {
	lock *sync.Mutex
}

// setSciNotationDtoWithDefaults - Transfers new data to an instance of
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
func (nStrFmtSpecSciNotDtoUtil *numStrFmtSpecSciNotationDtoUtility) setSciNotationDtoWithDefaults(
	nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto,
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	requestedNumberFieldLen int,
	ePrefix string) (
	err error) {

	if nStrFmtSpecSciNotDtoUtil.lock == nil {
		nStrFmtSpecSciNotDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDtoUtil.lock.Lock()

	defer nStrFmtSpecSciNotDtoUtil.lock.Unlock()

	ePrefix += "numStrFmtSpecSciNotationDtoUtility.setSciNotationDto() "

	if nStrFmtSpecSciNotDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
			ePrefix)

		return err
	}

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	nStrFmtSpecSciNotDtoMech :=
		numStrFmtSpecSciNotationDtoMechanics{}

	err = nStrFmtSpecSciNotDtoMech.setSciNotationDto(
		nStrFmtSpecSciNotDto,
		significandUsesLeadingPlus,
		mantissaLength,
		exponentChar,
		exponentUsesLeadingPlus,
		numFieldDto,
		ePrefix+
			"nStrFmtSpecSciNotDto\n")

	return err
}
