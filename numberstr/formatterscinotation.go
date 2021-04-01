package numberstr

import (
	"fmt"
	"reflect"
	"sync"
)

// FormatterSciNotation - Number String Format Specification
// for Scientific Notation. This type contains data fields and
// methods used in formatting number strings as scientific
// notation.
//
// Reference:
//  https://www.mathsisfun.com/numbers/scientific-notation.html
//  https://en.wikipedia.org/wiki/Scientific_notation
//
// Typically a number formatted for scientific notation takes
// the following form:
//      2.652E+8
//
// For background and general information, elements of a scientific
// notation numeric value are described below:
//
//  Significand
//    significand = '2.652'
//      significand integer digit = '2'
//      Integer digit is between 1 - 9, including 1 and 9.
//
//      significand uses leading plus sign
//         - Positive significand integer digits may have a
//           leading plus sign, '+2.652E+8'. The default is
//           no leading plus sign, '2.652E+8'
//
//         - Negative values always contain significand integer
//           digits with a leading minus sign, '-2.652E+8'
//
//  Mantissa
//      mantissa = significand factional digits = '.652'
//      mantissaLength = The number of fractional digits displayed in scientific notation.
//
//  Exponent Character
//      Exponent character or exponentChar = 'E'
//          The character 'E' is used as the default to avoid
//          confusion with Euler's number 'e'.  However the
//          character lower case 'e' is often used and may
//          therefore be specified by the user.
//             Examples: 2.652E+8 or 2.652e+8
//
//      exponent = '8'  (10^8) 2.652E+8
//        '+' exponent leading plus sign = exponentUsesLeadingPlus== true
//        The exponent is often displayed without a leading plus sign.
//             Example: 2.652E+8
//        Of course, depending on the value, the exponent may have a leading
//        minus sign.
//
//
// Data Fields
//
//  numStrFmtType                 NumStrFormatTypeCode
//     - An enumeration value automatically set to:
//           NumStrFormatTypeCode(0).AbsoluteValue()
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
//
type FormatterSciNotation struct {
	numStrFmtType              NumStrFormatTypeCode
	significandUsesLeadingPlus bool
	mantissaLength             uint
	exponentChar               rune
	exponentUsesLeadingPlus    bool
	numFieldDto                NumberFieldDto
	lock                       *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance
// of FormatterSciNotation ('incomingFmtSciNotation') to the
// current FormatterSciNotation instance.
//
// The FormatterSciNotation type encapsulates the Scientific
// Notation format specification used to format number strings for
// text displays.
//
// When this method completes processing current instance of
// FormatterSciNotation and the 'incomingFmtSciNotation' will
// have identical data values.
//
// This method will overwrite the member variable data values of
// the current FormatterSciNotation instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingFmtSciNotation        *FormatterSciNotation
//     - A pointer to an instance of FormatterSciNotation.
//       The member variable data fields from this instance will be
//       copied to those contained in the current instance of
//       FormatterSciNotation.
//
//       If parameter 'incomingFmtSciNotation' is judged to be
//       invalid, this method will return an error.
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
// ------------------------------------------------------------------------
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
func (fmtSciNotation *FormatterSciNotation) CopyIn(
	incomingFmtSciNotation *FormatterSciNotation,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation.CopyIn()")

	return formatterSciNotationElectron{}.ptr().
		copyIn(
			fmtSciNotation,
			incomingFmtSciNotation,
			ePrefix.XCtx("incomingFmtSciNotation->fmtSciNotation"))
}

// CopyInINumStrFormatter - Receives an incoming INumStrFormatter
// object, converts it to a FormatterSciNotation instance and
// proceeds to copy the the data fields into those of the
// current FormatterSciNotation instance.
//
// If the dynamic type of INumStrFormatter is not equal to
// FormatterSciNotation, an error will be returned. Likewise,
// if the data fields of input parameter 'incomingIFormatter' are
// judged to be invalid, an error will be returned.
//
// Be advised, all of the data fields in the current
// FormatterSciNotation instance will be overwritten.
//
// This method is required by interface INumStrFormatter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingIFormatter            INumStrFormatter
//     - An instance of Interface INumStrFormatter. If this the
//       dynamic type is not equal to FormatterSciNotation an error
//       will be returned.
//
//       The data values in this object will be copied to the
//       current FormatterSciNotation instance.
//
//       If input parameter 'incomingIFormatter' is judged to
//       be invalid, this method will return an error.
//
//
//  ePrefix                       *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
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
func (fmtSciNotation *FormatterSciNotation) CopyInINumStrFormatter(
	incomingIFormatter INumStrFormatter,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"CopyInINumStrFormatter()")

	if incomingIFormatter == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIFormatter' is "+
			"'nil'\n",
			ePrefix.String())
	}

	incomingFmtSciNotation,
		isValid := incomingIFormatter.(*FormatterSciNotation)

	if !isValid {

		actualType :=
			reflect.TypeOf(incomingIFormatter)

		typeName := "Unknown"

		if actualType != nil {
			typeName = actualType.Name()
		}

		return fmt.Errorf("%v\n"+
			"Error: 'incomingIFormatter' is NOT Type "+
			"FormatterSciNotation\n"+
			"'incomingIFormatter' is type %v",
			ePrefix.String(),
			typeName)
	}

	return formatterSciNotationElectron{}.ptr().
		copyIn(
			fmtSciNotation,
			incomingFmtSciNotation,
			ePrefix.XCtx("incomingFmtSciNotation->fmtSciNotation"))
}

// CopyOut - Returns a deep copy of the current instance of
// FormatterSciNotation.
//
// If the current FormatterSciNotation instance is judged to be
// invalid, this method will return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  FormatterSciNotation
//     - A new instance of FormatterSciNotation containing
//       data values identical to those contained in the current
//       FormatterSciNotation instance. This returned
//       instance of FormatterSciNotation represents a deep
//       copy of the current FormatterSciNotation instance.
//
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
func (fmtSciNotation *FormatterSciNotation) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterSciNotation,
	error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation.CopyOut()")

	return formatterSciNotationElectron{}.ptr().copyOut(
		fmtSciNotation,
		ePrefix.XCtx("fmtSciNotation->"))
}

func (fmtSciNotation *FormatterSciNotation) CopyOutINumStrFormatter(
	ePrefix *ErrPrefixDto) (
	INumStrFormatter,
	error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"CopyOutINumStrFormatter()")

	newFormatterSciNotation,
		err := formatterSciNotationElectron{}.ptr().copyOut(
		fmtSciNotation,
		ePrefix.XCtx("fmtSciNotation->"))

	return INumStrFormatter(&newFormatterSciNotation), err
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterSciNotation instance to
// their initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (fmtSciNotation *FormatterSciNotation) Empty() {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	_ = formatterSciNotationQuark{}.ptr().
		empty(fmtSciNotation,
			nil)

	fmtSciNotation.lock.Unlock()

	fmtSciNotation.lock = nil

	return
}

// GetExponentChar - Returns the exponent character configured for
// the current FormatterSciNotation instance. This character is
// returned as a 'rune' and is displayed in number strings
// formatted with scientific notation.
//
// Background
//
// Scientific Notation Example: 2.652E+8
//
// In the example above, the exponent character is 'E'. This is the
// default. Users may elect to use an exponent character of 'e'
// which would yield the following format result:
//                2.652e+8
//
// 'E' was selected as a default to avoid confusion with Euler's
// number which is also designated with the character 'e'.
//
// The exponent character applied in scientific notation formatting
// is completely controlled by the user. However, a rune or
// exponent character value, of zero will invalidate the format
// specification. An exponent character value of zero ('0') signals
// that the format specification is empty.
//
// This method returns the current scientific notation 'Exponent
// Character' value which is stored in the internal member
// variable:
//   FormatterSciNotation.exponentChar
//
func (fmtSciNotation *FormatterSciNotation) GetExponentChar() rune {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	return fmtSciNotation.exponentChar
}

// GetExponentUsesLeadingPlus - Returns the boolean flag which
// determines whether positive exponents in scientific notation
// will be prefixed with a leading plus sign.
//
// Background
//
// Scientific Notation Example: 2.652E+8
//
// In the example above the, 'exponent' is '8'. Positive value
// exponents may or may not be prefixed with a leading plus sign.
//
// When the 'Exponent Uses Leading Plus Sign' flag is set to
// 'true', the scientific notation format will be presented as:
//              2.652E+8
//
// When the 'Exponent Uses Leading Plus Sign' flag is set to
// 'false', the scientific notation format will be presented as:
//              2.652E8
//
// This setting only applies to exponents with a positive value.
// Exponents with a negative value will always be prefixed with a
// minus sign. Example: 2.652E-20
//
// This method returns the current 'Exponent Uses Leading Plus
// Sign' boolean flag value which is stored in the internal member
// variable:
//   FormatterSciNotation.exponentUsesLeadingPlus
//
func (fmtSciNotation *FormatterSciNotation) GetExponentUsesLeadingPlus() bool {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	return fmtSciNotation.exponentUsesLeadingPlus
}

// GetFmtNumStr - Returns a number string formatted for scientific
// notation based on the configuration encapsulated in the current
// instance of FormatterSciNotation.
//
// This method is required by interface INumStrFormatter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absValIntegerRunes            []rune
//     - An array of runes containing the integer component of the
//       numeric value to be formatted as scientific notation. This
//       array of integer digits always represents a positive ('+')
//       numeric value. The array consists entirely of numeric
//       digits.
//
//
//  absValFractionalRunes         []rune
//     - An array of runes containing the fractional component of
//       the numeric value to be formatted as scientific notation.
//       This array of numeric digits always represents a positive
//       ('+') numeric value. The array consists entirely of
//       numeric digits.
//
//
//  signVal                       int
//     - The parameter designates the numeric sign of the final
//       formatted scientific notation value returned by this
//       method.
//
//       Valid values for 'signVal' are listed as follows:
//         -1 = Signals a negative numeric value
//          0 = Signals that the numeric value is zero.
//         +1 = Signals a positive numeric value
//
//
//  baseNumSys                    BaseNumberSystemType
//     - Designates the base number system associated with input
//       parameters 'absValIntegerRunes' and 'absValIntegerRunes'
//       described above.
//
//       Valid values for 'baseNumSys' are listed as follows:
//         BaseNumSys.Binary(),
//         BaseNumSys.Octal(),
//         BaseNumSys.Decimal(),
//         BaseNumSys.Hexadecimal(),
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  fmtNumStr                     string
//     - If this method completes successfully, a formatted
//       scientific notation number string will be returned.
//
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
func (fmtSciNotation *FormatterSciNotation) GetFmtNumStr(
	absValIntegerRunes []rune,
	absValFractionalRunes []rune,
	signVal int,
	baseNumSys BaseNumberSystemType,
	ePrefix *ErrPrefixDto) (
	fmtNumStr string,
	err error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"GetFmtNumStr()")

	if baseNumSys != BaseNumberSystemType(0).Decimal() {
		err = fmt.Errorf("%v\n"+
			"Error: Base Numbering System is NOT equal to Base-10!\n"+
			"baseNumSys=='%v'\n",
			ePrefix.String(),
			baseNumSys.XValueInt())
		return fmtNumStr, err
	}

	if signVal < 0 {
		fmtNumStr = "-"
	}

	fmtNumStr += string(absValIntegerRunes) +
		string(absValFractionalRunes)

	return fmtNumStr, err
}

// GetMantissaLength - Returns the Mantissa Length as an unsigned
// integer.
//
// Background
//
// Typically a number formatted for scientific notation takes
// the following form:
//      2.652E+8
//
// In the example shown above, the mantissa is ".652"
// The length of the mantissa in this example is '3'.
//
// The length of the mantissa is used in formatting a numeric value
// in scientific notation. This value is usually fairly small in
// the range of 1 - 8 digits. However, this length is completely
// controlled by the user. A value of zero invalidates the format
// specification.
//
//  A zero value mantissa length signals that the specification is
//  empty.
//
// This method returns the current 'Mantissa Length' value which is
// stored in the internal member variable:
//   FormatterSciNotation.mantissaLength
//
func (fmtSciNotation *FormatterSciNotation) GetMantissaLength() (mantissaLength uint) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	mantissaLength = fmtSciNotation.mantissaLength

	return mantissaLength
}

// GetNumberFieldDto - Returns the NumberFieldDto object
// currently configured for this Scientific Notation Number String
// Format Specification.
//
// The NumberFieldDto details the length of the number field in
// which the scientific notation value will be displayed and
// justified left, right or center according to the specification.
//
// If the NumberFieldDto object is judged to be invalid, this
// method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumberFieldDto
//     - If this method completes successfully, a new instance of
//       NumberFieldDto will be returned through this parameter.
//       This object is deep copy of the Number Field information
//       used to configure the current instance of
//       FormatterSciNotation.
//
//       The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
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
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
//
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
//       Be advised that if the returned 'NumberFieldDto' object is
//       judged invalid, this method will return an error.
//
func (fmtSciNotation *FormatterSciNotation) GetNumberFieldDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"fmtSciNotation." +
			"GetNumberFieldDto()")

	return fmtSciNotation.numFieldDto.CopyOut(
		ePrefix.XCtx(
			"fmtSciNotation"))
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for
// FormatterSciNotation objects is
//NumStrFormatTypeCode(0).ScientificNotation().
//
// This method is required by interface INumStrFormatter.
//
func (fmtSciNotation *FormatterSciNotation) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	return fmtSciNotation.numStrFmtType
}

// GetSignificandUsesLeadingPlus - Returns the 'Significand Uses
// Leading Plus' flag configured for the current instance of
// FormatterSciNotation. This flag determines whether positive
// 'significand' integer digits in scientific notation will be
// prefixed with a leading plus sign.
//
//
// Background
//
// Typically a number formatted for scientific notation takes
// the following form:
//      2.652E+8
// In this example, the 'significand' is '2.652'. The designation
// 'Significand Uses Leading Plus' refers to the following
// examples showing positive significand integer digits :
//
//  - Positive significand integer digits may have a
//    leading plus sign, '+2.652E+8'.
//
//  - The default is no leading plus sign, '2.652E+8'
//
// If the 'Significand Uses Leading Plus' flag is set to 'true', it
// signals that positive significand integer digits WILL HAVE a
// leading plus sign: '+2.652E+8'.
//
// Conversely, if the 'Significand Uses Leading Plus' flag is set
// to 'false', it signals that positive significand integer digits
// WILL NOT HAVE a leading plus sign: '2.652E+8'.
//
// This method returns the current 'Significand Uses Leading Plus'
// flag which is stored in the internal member variable:
//   FormatterSciNotation.significandUsesLeadingPlus
//
func (fmtSciNotation *FormatterSciNotation) GetSignificandUsesLeadingPlus() bool {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	return fmtSciNotation.significandUsesLeadingPlus
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterSciNotation instance to determine whether
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
//       FormatterSciNotation is valid, or not. If the
//       current FormatterSciNotation contains valid data,
//       this method returns 'true'. If the data is invalid, this
//       method will return 'false'.
//
func (fmtSciNotation *FormatterSciNotation) IsValidInstance() (isValid bool) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	isValid,
		_ = formatterSciNotationQuark{}.ptr().
		testValidityOfFormatterSciNotation(
			fmtSciNotation,
			new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// FormatterSciNotation instance to determine whether the
// current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             *ErrPrefixDto
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
func (fmtSciNotation *FormatterSciNotation) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"" +
			"FormatterSciNotation." +
			"IsValidInstanceError()")

	_,
		err := formatterSciNotationQuark{}.ptr().
		testValidityOfFormatterSciNotation(
			fmtSciNotation,
			ePrefix.XCtx("Testing validity of fmtSciNotation"))

	return err
}

// NewWithComponents - Creates and returns a new instance of
// FormatterSciNotation.
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
//  numFieldDto                   NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
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
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
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
//  FormatterSciNotation
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterSciNotation.
//
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
func (fmtSciNotation FormatterSciNotation) NewWithComponents(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	FormatterSciNotation,
	error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(

		"FormatterSciNotation." +
			"NewWithComponents()")

	newNStrFmtSpecSciNotationDto := FormatterSciNotation{}

	err := formatterSciNotationMechanics{}.ptr().setFmtSciNotWithComponents(
		&newNStrFmtSpecSciNotationDto,
		significandUsesLeadingPlus,
		mantissaLength,
		exponentChar,
		exponentUsesLeadingPlus,
		numFieldDto,
		ePrefix.XCtx(
			"newNStrFmtSpecSciNotationDto"))

	return newNStrFmtSpecSciNotationDto, err
}

// NewWithDefaults - Creates and returns a new instance of
// FormatterSciNotation.
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
//       The 'NumberFieldDto' object created from this parameter
//       will be configured with default 'Right-Justification' text
//       alignment.
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
//  FormatterSciNotation
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterSciNotation.
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
func (fmtSciNotation FormatterSciNotation) NewWithDefaults(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterSciNotation,
	error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation.NewBasic()")

	newNStrFmtSpecSciNotationDto :=
		FormatterSciNotation{}

	err := formatterSciNotationUtility{}.ptr().setSciNotWithDefaults(
		&newNStrFmtSpecSciNotationDto,
		significandUsesLeadingPlus,
		mantissaLength,
		exponentChar,
		exponentUsesLeadingPlus,
		requestedNumberFieldLen,
		numberFieldTextJustify,
		ePrefix)

	return newNStrFmtSpecSciNotationDto, err
}

// SetExponentChar - Sets the exponent character which will be
// displayed in number strings formatted with scientific notation.
//
// Background
//
// Scientific Notation Example: 2.652E+8
//
// In the example above the exponent character is 'E'. This is the
// default. Users may wish to use an exponent character of 'e'
// which would yield the following format result:
//                2.652e+8
//
// 'E' was selected as a default to avoid confusion with Euler's
// number which is also designated with the character 'e'.
//
// The exponent character applied in scientific notation formatting
// is completely controlled by the user. However, a rune or
// character value, of zero will invalidate the format
// specification. An exponent character value of zero ('0') signals
// that the format specification is empty and will trigger an
// error.
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
//       Setting exponentChar = 0 invalidates the format
//       specification, signals that the format specification is
//       empty and triggers an error return.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
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
func (fmtSciNotation *FormatterSciNotation) SetExponentChar(
	exponentChar rune,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"SetExponentChar()")

	if exponentChar == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'exponentChar' is invalid!\n"+
			"'exponentChar' is equal to ZERO ('0').\n",
			ePrefix)
	}

	fmtSciNotation.exponentChar = exponentChar

	return nil
}

// SetExponentUsesLeadingPlus - Sets the boolean flag which
// determines whether positive exponents in scientific notation
// will be prefixed with a leading plus sign.
//
// Background
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
func (fmtSciNotation *FormatterSciNotation) SetExponentUsesLeadingPlus(
	exponentUsesLeadingPlusSign bool) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	fmtSciNotation.exponentUsesLeadingPlus = exponentUsesLeadingPlusSign
}

// SetMantissaLength - Sets the mantissa length used in
// formatting number strings using Scientific Notation.
//
// Background
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
//       If a zero value is submitted, it signals that the
//       specification is empty and an error will be returned.
//       Likewise, if 'mantissaLength' is greater than 10,000, an
//       error will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// ------------------------------------------------------------------------
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
func (fmtSciNotation *FormatterSciNotation) SetMantissaLength(
	mantissaLength uint,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"SetMantissaLength()")

	if mantissaLength == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'mantissaLength' is invalid!\n"+
			"'mantissaLength' is equal to ZERO ('0').\n",
			ePrefix)
	}

	if mantissaLength > 10000 {
		return fmt.Errorf("%v\n"+
			"Error: 'formatterSciNotation.mantissaLength' is invalid!\n"+
			"formatterSciNotation.mantissaLength is greater than 10,000.\n"+
			"formatterSciNotation.mantissaLength== '%v'\n",
			ePrefix.String(),
			mantissaLength)
	}

	fmtSciNotation.mantissaLength = mantissaLength

	return nil
}

// SetNumberFieldLengthDto - Sets the Number Field Length Dto object
// for the current FormatterSciNotation instance.
//
// The Number Field Length Dto object is used to specify the length
// and string justification characteristics used to display
// scientific notation number strings within a number field.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldLenDto   NumberFieldDto
//     - If this method completes successfully, a new instance of
//       NumberFieldDto will be returned through this parameter.
//       This object is deep copy of the Number Field information
//       used to configure the current instance of
//       FormatterSciNotation.
//
//       The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
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
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
//
//
//  ePrefix             *ErrPrefixDto
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
func (fmtSciNotation *FormatterSciNotation) SetNumberFieldLengthDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation.SetNumberFieldDto()")

	return fmtSciNotation.numFieldDto.CopyIn(
		&numberFieldLenDto,
		ePrefix.XCtx("fmtSciNotation\n"))
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterSciNotation object. For Scientific
// Notation formatters, the Number String Format Type Code is set
// to NumStrFormatTypeCode(0).ScientificNotation().
//
// This method is required by interface INumStrFormatter.
//
func (fmtSciNotation *FormatterSciNotation) SetNumStrFormatTypeCode() {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	fmtSciNotation.numStrFmtType =
		NumStrFormatTypeCode(0).ScientificNotation()

}

// SetSignificandUsesLeadingPlus - Sets the boolean flag which
// determines whether positive significand integer digit values
// in scientific notation will be prefixed with a leading plus
// sign.
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
func (fmtSciNotation *FormatterSciNotation) SetSignificandUsesLeadingPlus(
	significandUsesLeadingPlus bool) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	fmtSciNotation.significandUsesLeadingPlus = significandUsesLeadingPlus
}

// SetToUnitedStatesDefaults - Sets the member variable data values
// for this FormatterSciNotation instance to United
// States default values.
//
// In the United States, Signed Number default formatting
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
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterSciNotation instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
func (fmtSciNotation *FormatterSciNotation) SetToUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"SetToUnitedStatesDefaults()")

	if fmtSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err =
		formatterSciNotationUtility{}.ptr().
			setToUnitedStatesDefaults(
				fmtSciNotation,
				ePrefix)

	return err
}

// SetToUnitedStatesDefaultsIfEmpty - If the current
// FormatterSciNotation instance is empty or invalid,
// this method will set the member variable data values to United
// States default values.
//
// If the current FormatterSciNotation instance is valid
// and NOT empty, this method will take no action and exit.
//
// In the United States, Signed Number default formatting
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
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterSciNotation instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
func (fmtSciNotation *FormatterSciNotation) SetToUnitedStatesDefaultsIfEmpty(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation." +
			"SetToUnitedStatesDefaultsIfEmpty()")

	if fmtSciNotation == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSciNotDto' is invalid!\n"+
			"'targetSciNotDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	isValid,
		_ := formatterSciNotationQuark{}.ptr().
		testValidityOfFormatterSciNotation(
			fmtSciNotation,
			new(ErrPrefixDto))

	if isValid {
		return err
	}

	err =
		formatterSciNotationUtility{}.ptr().
			setToUnitedStatesDefaults(
				fmtSciNotation,
				ePrefix)

	return err
}

// SetWithComponents - This method will set all of the member
// variable data values for the current instance of
// FormatterSciNotation.
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterSciNotation instance.
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
//       The NumberFieldDto type is detailed as follows:
//
//       type NumberFieldDto struct {
//         requestedNumFieldLength int // User requested number field length
//         actualNumFieldLength    int // Machine generated actual number field Length
//         minimumNumFieldLength   int // Machine generated minimum number field length
//         textJustifyFormat       TextJustify // User specified text justification
//       }
//
//       requestedNumberFieldLen    int
//
//       - This is the requested length of the number field in which
//         the number string will be displayed.
//
//         If this field length is greater than the actual length of
//         the number string, the number string will be justified
//         within the number field. If the actual number string
//         length is greater than the requested number field length,
//         the number field length will be automatically expanded
//         to display the entire number string. The 'requested'
//         number field length is used to create number fields
//         of standard lengths for text presentations.
//
//         If 'requestedNumberFieldLen' is set to a value of minus
//         one (-1), the final number field length will be set to
//         the length of the actual number string.
//
//       numberFieldTextJustify        TextJustify
//       - An enumeration value used to specify the type of text
//         formatting which will be applied to a number string when
//         it is positioned inside of a number field. This
//         enumeration value must be one of the three following
//         format specifications:
//
//         1. Left   - Signals that the text justification format is
//                     set to 'Left-Justify'. Strings within text
//                     fields will be flush with the left margin.
//                            Example: "TextString      "
//
//         2. Right  - Signals that the text justification format is
//                     set to 'Right-Justify'. Strings within text
//                     fields will terminate at the right margin.
//                            Example: "      TextString"
//
//         3. Center - Signals that the text justification format is
//                     is set to 'Centered'. Strings will be positioned
//                     in the center of the text field equidistant
//                     from the left and right margins.
//                             Example: "   TextString   "
//
//
//  ePrefix             *ErrPrefixDto
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
func (fmtSciNotation *FormatterSciNotation) SetWithComponents(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(

		"FormatterSciNotation." +
			"SetWithComponents()")

	return formatterSciNotationMechanics{}.ptr().
		setFmtSciNotWithComponents(
			fmtSciNotation,
			significandUsesLeadingPlus,
			mantissaLength,
			exponentChar,
			exponentUsesLeadingPlus,
			numFieldDto,
			ePrefix.XCtx("fmtSciNotation"))
}

// SetWithDefaults - This method will set all of the member
// variable data values for the current instance of
// FormatterSciNotation. The input parameters
// represent the minimum information required to set the data
// values for a FormatterSciNotation object.
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
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
func (fmtSciNotation *FormatterSciNotation) SetWithDefaults(
	significandUsesLeadingPlus bool,
	mantissaLength uint,
	exponentChar rune,
	exponentUsesLeadingPlus bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtSciNotation.lock == nil {
		fmtSciNotation.lock = new(sync.Mutex)
	}

	fmtSciNotation.lock.Lock()

	defer fmtSciNotation.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterSciNotation.SetWithDefaults()")

	return formatterSciNotationUtility{}.ptr().
		setSciNotWithDefaults(
			fmtSciNotation,
			significandUsesLeadingPlus,
			mantissaLength,
			exponentChar,
			exponentUsesLeadingPlus,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx("fmtSciNotation"))
}
