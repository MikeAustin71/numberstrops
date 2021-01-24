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
//         Positive significand integer digits may have a
//         leading plus sign, '+2.652E+8'. The default is
//         no leading plus sign, '2.652E+8'
//
//  Mantissa
//      mantissa = significand factional digits = '.652'
//      mantissaLength = length of fractional digits displayed in scientific notation.
//
//  Exponent
//      Exponent character or exponentChar = 'E'
//          The character 'E' is used as the default to avoid confusion
//          with Euler's number 'e'.  However the character 'e' is
//          often used and may therefore be specified by the user.
//             Examples: 2.652E+8 or 2.652e+8
//
//      exponent = '8'  (10^8) 2.652E+8
//        '+' exponent leading plus sign = exponentUsesLeadingPlus== true
//        The exponent is often displayed without a leading plus sign.
//             Example: 2.652E+8
//
//  Number Field Length
//      number field  length = Length of string in which
//        scientific notation is right justified.
//        Example: Number Field Length = 9
//  -------------------------------------- 123456789
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

// CopyIn - Copies the data fields from an incoming instance
// of NumStrFmtSpecSciNotationDto ('incomingSciNotDto') to the
// current NumStrFmtSpecSciNotationDto instance.
//
// The NumStrFmtSpecSciNotationDto type encapsulates the Scientific
// Notation format specification used to format number strings for
// text displays.
//
// When this method completes processing current instance of
// NumStrFmtSpecSciNotationDto and the 'incomingSciNotDto' will
// have identical data values.
//
// This method will overwrite the member variable data values of
// the current NumStrFmtSpecSciNotationDto instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingSciNotDto        *NumStrFmtSpecSciNotationDto
//     - A pointer to an instance of NumStrFmtSpecSciNotationDto.
//       The member variable data fields from this instance will be
//       copied to those contained in the current instance of
//       NumStrFmtSpecSciNotationDto.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) CopyIn(
	incomingSciNotDto *NumStrFmtSpecSciNotationDto,
	ePrefix string) error {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecSciNotationDto.CopyIn()\n"

	nStrFmtSpecSciNotDtoElectron :=
		numStrFmtSpecSciNotationDtoElectron{}

	return nStrFmtSpecSciNotDtoElectron.copyIn(
		nStrFmtSpecSciNotDto,
		incomingSciNotDto,
		ePrefix+
			"incomingSciNotDto->nStrFmtSpecSciNotDto\n")
}

// CopyOut - Returns a deep copy of the current instance of
// NumStrFmtSpecSciNotationDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecSciNotationDto
//     - A new instance of NumStrFmtSpecSciNotationDto containing
//       data values identical to those contained in the current
//       NumStrFmtSpecSciNotationDto instance. This returned
//       instance of NumStrFmtSpecSciNotationDto represents a deep
//       copy of the current NumStrFmtSpecSciNotationDto instance.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) CopyOut(
	ePrefix string) (
	NumStrFmtSpecSciNotationDto,
	error) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecSciNotationDto.CopyOut() "

	nStrFmtSpecSciNotDtoElectron :=
		numStrFmtSpecSciNotationDtoElectron{}

	return nStrFmtSpecSciNotDtoElectron.copyOut(
		nStrFmtSpecSciNotDto,
		ePrefix+
			"nStrFmtSpecSciNotDto->\n")
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecSciNotationDto instance to determine whether
// the current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - This returned boolean value will signal whether the current
//       NumStrFmtSpecSciNotationDto is valid, or not. If the
//       current NumStrFmtSpecSciNotationDto contains valid data,
//       this method returns 'true'. If the data is invalid, this
//       method will return 'false'.
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) IsValidInstance() (isValid bool) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	nStrFmtSpecSciNotQuark :=
		numStrFmtSpecSciNotationDtoQuark{}

	isValid,
		_ = nStrFmtSpecSciNotQuark.
		testValidityOfNumStrFmtSpecSciNotationDto(
			nStrFmtSpecSciNotDto,
			"")

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFmtSpecSciNotationDto instance to determine whether the
// current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current instance of NumStrFmtSpecSciNotationDto
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) IsValidInstanceError(
	ePrefix string) error {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecSciNotationDto.IsValidInstanceError() "

	nStrFmtSpecSciNotQuark :=
		numStrFmtSpecSciNotationDtoQuark{}

	_,
		err := nStrFmtSpecSciNotQuark.
		testValidityOfNumStrFmtSpecSciNotationDto(
			nStrFmtSpecSciNotDto,
			ePrefix+
				"\nnStrFmtSpecSciNotDto ")

	return err
}

// NewWithDefaults() - Creates and returns a new instance of
// NumStrFmtSpecSciNotationDto.
//
// Scientific Notation Format Specification objects encapsulate the
// format specifications used in formatting scientific notation
// numeric values for text display.
//
// This method will apply default "Right-Justification" for the
// Number Field specification created from input parameter,
// 'requestedNumberFieldLen'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//       The 'NumberFieldDto' object created from this parameter
//       will be configured with default 'Right-Justification' text
//       alignment.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecSciNotationDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecSciNotationDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecSciNotDto NumStrFmtSpecSciNotationDto) NewWithDefaults(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	requestedNumberFieldLen int,
	ePrefix string) (
	NumStrFmtSpecSciNotationDto,
	error) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecSciNotationDto.NewWithDefaults() "

	newNStrFmtSpecSciNotationDto :=
		NumStrFmtSpecSciNotationDto{}

	nStrFmtSpecSciNotDtoUtil :=
		numStrFmtSpecSciNotationDtoUtility{}

	err := nStrFmtSpecSciNotDtoUtil.setSciNotationDtoWithDefaults(
		&newNStrFmtSpecSciNotationDto,
		significandUsesLeadingPlus,
		mantissaLength,
		exponentChar,
		exponentUsesLeadingPlus,
		requestedNumberFieldLen,
		ePrefix)

	return newNStrFmtSpecSciNotationDto, err
}

// NewFromComponents() - Creates and returns a new instance of
// NumStrFmtSpecSciNotationDto.
//
// Scientific Notation Format Specification objects encapsulate the
// format specifications used in formatting scientific notation
// numeric values for text display.
//
// This method requires a properly configured NumberFieldDto object
// as an input parameter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecSciNotationDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrFmtSpecSciNotationDto.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecSciNotDto NumStrFmtSpecSciNotationDto) NewFromComponents(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	numFieldDto NumberFieldDto,
	ePrefix string) (
	NumStrFmtSpecSciNotationDto,
	error) {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecSciNotationDto.NewFromComponents() "

	newNStrFmtSpecSciNotationDto := NumStrFmtSpecSciNotationDto{}

	nStrFmtSpecSciNotDtoMech :=
		numStrFmtSpecSciNotationDtoMechanics{}

	err := nStrFmtSpecSciNotDtoMech.setSciNotationDto(
		&newNStrFmtSpecSciNotationDto,
		significandUsesLeadingPlus,
		mantissaLength,
		exponentChar,
		exponentUsesLeadingPlus,
		numFieldDto,
		ePrefix+
			"\nnewNStrFmtSpecSciNotationDto\n")

	return newNStrFmtSpecSciNotationDto, err
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix string) error {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	ePrefix += "\nNumStrFmtSpecSciNotationDto.SetNumberFieldLengthDto()\n"

	return nStrFmtSpecSciNotDto.numFieldLenDto.CopyIn(
		&numberFieldLenDto,
		ePrefix+
			"nStrFmtSpecSciNotDto\n")
}

// SetSciNotDto() - This method will set all of the member
// variable data values for the current instance of
// NumStrFmtSpecSciNotationDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (nStrFmtSpecSciNotDto *NumStrFmtSpecSciNotationDto) SetSciNotDto(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	numFieldDto NumberFieldDto,
	ePrefix string) error {

	if nStrFmtSpecSciNotDto.lock == nil {
		nStrFmtSpecSciNotDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecSciNotDto.lock.Lock()

	defer nStrFmtSpecSciNotDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecSciNotationDto.SetSciNotDto() "

	nStrFmtSpecSciNotDtoMech :=
		numStrFmtSpecSciNotationDtoMechanics{}

	return nStrFmtSpecSciNotDtoMech.setSciNotationDto(
		nStrFmtSpecSciNotDto,
		significandUsesLeadingPlus,
		mantissaLength,
		exponentChar,
		exponentUsesLeadingPlus,
		numFieldDto,
		ePrefix+
			"nStrFmtSpecSciNotDto\n")
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
