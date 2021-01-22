package numberstr

import "sync"

// Number String Format Specification for Scientific Notation
// This type contains data fields and methods used in formatting
// number strings as scientific notation.
// Reference:
//  https://www.mathsisfun.com/numbers/scientific-notation.html
//  https://en.wikipedia.org/wiki/Scientific_notation
//
// Typically a number formatted for scientific notation takes
// the following form:
//      2.652E+8
//
// The components which comprise this number format are
// detailed below:
//
//  Significand
//    significand = '2.652'
//      significand integer digit = '2'
//      Integer digit is between 1 - 9, including 1 and 9.
//
//      significand uses leading plus sign
//      Positive significand integer digits may have a
//       leading plus sign, '+2.652E+8'. The default is
//       no leading plus sign, '2.652E8'
//
//  Mantissa
//      mantissa = significand factional digits = '.652'
//      mantissaLength = length of fractional digits displayed in scientific notation.
//
//  Exponent
//      exponent character or exponentChar = 'E'
//        The character 'E' is used as the default to avoid confusion
//          with Euler's number 'e'.  However the character 'e' is
//          often used and may therefore be specified by the user.
//           Examples: 2.652E+8 or 2.652e+8
//
//      exponent = '8'  (10^8) 2.652E+8
//      '+' exponent leading plus sign = exponentUsesLeadingPlus== true
//          The exponent is often displayed without a leading plus sign.
//          Example: 2.652E+8
//
//  Number Field
//      number field  length = Length of string in which
//        scientific notation is right justified.
//        Example: Number Field Length = 9
//                                         123456789
//        Formatted Scientific Notation = " 2.652E+8"
//
type NumStrFmtSpecSciNotationDto struct {
	significandUsesLeadingPlus bool
	mantissaLength             uint
	exponentChar               rune
	exponentUsesLeadingPlus    bool
	numFieldLenDto             NumberFieldDto
	lock                       *sync.Mutex
}

func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) CopyIn(
	incomingSciNotDto *NumStrFmtSpecSciNotationDto,
	ePrefix string) error {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecSciNotationDto.CopyIn()\n"

	nStrFmtSpecSciNotDto.numFieldLenDto.CopyIn(
		&incomingSciNotDto.numFieldLenDto)

	return nil
}

// SetExponentChar - Sets the exponent character which will be
// displayed in number strings formatted with scientific notation.
//
// Terminology
//
// Scientific Notation Example: 2.652E+8
//
// In the example above the exponent character is 'E'. This is the
// default. Users may wish to use an exponent character of 'e' which
// would result in the following format result:
//                2.652e+8
//
// 'E' was selected as a default to avoid confusion with Euler's
// number which is also designated with the character 'e'.
//
// The exponent character applied in scientific notation formatting
// is completely controlled by the user. However, a rune or
// character value, of zero will invalidate the format
// specification. An exponent character value of zero signals that
// the format specification is empty.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  exponentChar        rune
//     - Sets the exponent character used in formatting a numeric
//       value in scientific notation. In the following examples
//       the exponent character is formatted as 'e' and 'E'.
//       Examples:
//                  2.652e+8
//                  2.652E+8
//
//       Setting exponentChar = 0 invalidates the format specification
//       and signals that the format specification is empty.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetExponentChar(
	exponentChar rune) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	nStrFmtSpecSciNotDto.exponentChar = exponentChar
}

// SetExponentUsesLeadingPlus - Sets the boolean flag which
// determines whether positive exponents will be prefixed with a
// plus sign in scientific notation.
//
// Terminology
//
// Scientific Notation Example: 2.652E+8
//
// In the example above the 'exponent' is '8'. Positive value
// exponents may or may not be prefixed with a leading plus sign.
//
// When input parameter 'exponentUsesLeadingPlusSign' is set to
// 'true', the scientific notation format will be presented as:
//              2.652E+8
//
// When input parameter 'exponentUsesLeadingPlusSign' is set to
// 'false', the scientific notation format will be presented as:
//              2.652E8
//
// This setting only applies to exponents with a positive value.
// Exponents with a negative value will always be prefixed with a
// minus sign. Example: 2.652E-20
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  exponentUsesLeadingPlusSign   bool
//     - Sets the boolean flag which determines whether positive
//       value exponents will be preceded with a plus sign.
//
//       If this value is set to 'true', the resulting formatted
//       scientific notation will include an exponent leading plus
//       sign:
//                 '2.652E+8'
//
//       If this value is set to 'false', the resulting formatted
//       scientific notation will NOT include an exponent leading
//       plus sign:
//                 '2.652E8'
//
//       This setting only applies to exponents with a positive
//       value. Exponents with a negative value will always be
//       prefixed with a minus sign. Example: 2.652E-20
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetExponentUsesLeadingPlus(
	exponentUsesLeadingPlusSign bool) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	nStrFmtSpecSciNotDto.exponentUsesLeadingPlus = exponentUsesLeadingPlusSign
}

// SetMantissaLength - Sets the mantissa length used in
// formatting number strings using Scientific Notation.
//
// Terminology
//
// Scientific Notation Example: 2.652E+8
// The mantissa in this example is ".652"
// The length of the mantissa in this example is '3'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  mantissaLength      uint
//     - Sets the length of the mantissa used in formatting a
//       numeric value in scientific notation. This value is
//       usually fairly small in the range of 1 - 8 digits.
//       However, this length is completely controlled by the
//       user. A value of zero invalidates the format specification.
//       A zero value signals that the specification is empty.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetMantissaLength(
	mantissaLength uint) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	nStrFmtSpecSciNotDto.mantissaLength = mantissaLength
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current NumStrFmtSpecSciNotationDto instance.
//
// The Number Separators Dto object is used to specify the Decimal
// Separators Character and the Integer Digits Separator Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldLenDto  NumberFieldDto
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
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	nStrFmtSpecSciNotDto.numFieldLenDto.CopyIn(
		&numberFieldLenDto)
}

// SetSignificandUsesLeadingPlus - Sets the boolean flag which
// determines whether positive significand integer digit values
// will be prefixed with a leading plus sign.
//
// Terminology
//
// Scientific Notation Example: 2.652E+8
//
// The 'significand' in the example above is '2.652'. The
// significand integer digit is '2'.
//
// Setting input parameter 'significandUsesLeadingPlus' to 'true'
// will result in the following format output: '+2.652E+8'
//
// Setting input parameter 'significandUsesLeadingPlus' to 'false'
// will result in the following format output: '2.652E+8'
//
// This setting only applies to significand positive integer digit
// values. If the significand integer digit value is negative,
// this setting has no impact because negative values are always
// prefixed with a leading minus sign like this: '-2.652E+8'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  significandUsesLeadingPlus    bool
//     - Sets the boolean flag which determines whether positive
//       significand integer digit values will be prefixed with a
//       leading plus sign.
//
//       If this value is set to 'true', the resulting formatted
//       scientific notation will include the leading plus sign:
//             '+2.652E+8'
//
//       If this value is set to 'false', the resulting formatted
//       scientific notation will NOT include the leading plus sign:
//             '2.652E+8'
//
//       This setting only applies to positive significand integer
//       digit values. If the significand integer digit value is
//       negative, this setting has no impact because negative
//       values are always prefixed with a leading minus sign as
//       shown in this example:
//             '-2.652E+8'
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetSignificandUsesLeadingPlus(
	significandUsesLeadingPlus bool) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	nStrFmtSpecSciNotDto.significandUsesLeadingPlus = significandUsesLeadingPlus
}
