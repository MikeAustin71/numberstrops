package numberstr

import (
	"fmt"
	"reflect"
	"sync"
)

// FormatterBinary - The FormatterBinary type encapsulates the
// formatting parameters necessary to format binary digits for
// display in text number strings.
//
// Reference:
//  https://en.wikipedia.org/wiki/Binary_number
//
// Member data elements are listed as follows:
//
//   type FormatterBinary struct {
//    numStrFmtType                 NumStrFormatTypeCode
//    turnOnIntegerDigitsSeparation bool
//    numericSeparators             NumericSeparators
//    numFieldDto                   NumberFieldDto
//   }
//
//
//  numStrFmtType           NumStrFormatTypeCode
//     - An enumeration value automatically set to:
//           NumStrFormatTypeCode(0).Octal()
//
//
//  turnOnIntegerDigitsSeparation bool
//     - Sets the 'Turn On Integer Digits Separation" flag.
//          FormatterBinary.turnOnIntegerDigitsSeparation
//
//       This boolean flag determines whether integer separation is
//       applied to formatted binary digits in a number
//       string.
//
//       If set to 'true', this flag signals that binary
//       digits will be grouped and separated according to the
//       specification contained in the internal member variable,
//       'integerSeparators'. When set to 'false', this flag
//       ensures that no integer digit separation will occur in
//       formatted binary number strings.
//
//
//  numericSeparators             NumericSeparators
//     - This instance of NumericSeparators is used to specify the
//       separator characters which will be included in number
//       string text displays.
//
//        type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//        }
//
//        decimalSeparators     []rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//        The decimal separator may consist of one or more runes.
//
//        integerSeparatorsDto    NumStrIntSeparatorsDto
//
//        The NumStrIntSeparatorsDto type encapsulates the integer digits
//        separators, often referred to as the 'Thousands Separator'.
//        Integer digit separators are used to separate integers into
//        specific groups within a number string. The
//        NumStrIntSeparatorsDto manages an array or collection of
//        NumStrIntSeparator objects.
//
//               type NumStrIntSeparatorsDto struct {
//                 intSeparators []NumStrIntSeparator
//               }
//
//               type NumStrIntSeparator struct {
//                intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//                intSeparatorGrouping    uint    // Number of integer digits in a group
//                intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                                // A zero value signals unlimited repetitions.
//                restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//               }
//
//               intSeparatorChars          []rune
//                  - A series of runes or characters used to separate integer
//                    digits in a number string. These characters are commonly
//                    known as the 'thousands separator'. A 'thousands
//                    separator' is used to separate groups of integer digits to
//                    the left of the decimal separator (a.k.a. decimal point).
//                    In the United States, the standard integer digits
//                    separator is the single comma character (','). Other
//                    countries and cultures use periods, spaces, apostrophes or
//                    multiple characters to separate integers.
//                          United States Example:  1,000,000,000
//
//               intSeparatorGrouping       uint
//                  - This unsigned integer values specifies the number of
//                    integer digits within a group. This value is used to group
//                    integers within a number string.
//
//                    In most western countries integer digits to the left of
//                    the decimal separator (a.k.a. decimal point) are separated
//                    into groups of three digits representing a grouping of
//                    'thousands' like this: '1,000,000,000'. In this case the
//                    intSeparatorGrouping value would be set to three ('3').
//
//                    In some countries and cultures other integer groupings are
//                    used. In India, for example, a number might be formatted
//                    like this: '6,78,90,00,00,00,00,000'.
//
//               intSeparatorRepetitions    uint
//                  - This unsigned integer value specifies the number of times
//                    this integer grouping is repeated. A value of zero signals
//                    that this integer grouping will be repeated indefinitely.
//
//               restartIntGroupingSequence bool
//                  - If the NumStrIntSeparator is the last element in an array
//                    of NumStrIntSeparator objects, this boolean flag signals
//                    whether the entire integer grouping sequence will be
//                    restarted from array element zero.
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
type FormatterBinary struct {
	numStrFmtType                 NumStrFormatTypeCode
	turnOnIntegerDigitsSeparation bool
	numericSeparators             NumericSeparators
	numFieldDto                   NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming FormatterBinary
// instance to the data fields of the current FormatterBinary
// instance.
//
// If input parameter 'incomingFormatterBinary' is judged to be
// invalid, this method will return an error.
//
// Be advised, all of the data fields in the current
// FormatterBinary instance will be overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFormatterBinary       *FormatterBinary
//     - A pointer to an instance of FormatterBinary.
//       The data values in this object will be copied to the
//       current FormatterBinary instance.
//
//       If input parameter 'incomingFormatterBinary' is judged to
//       be invalid, this method will return an error.
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
func (fmtBinary *FormatterBinary) CopyIn(
	incomingFormatterBinary *FormatterBinary,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"CopyIn()")

	return formatterBinaryAtom{}.ptr().copyIn(
		fmtBinary,
		incomingFormatterBinary,
		ePrefix.XCtx(
			"incomingFormatterBinary->"+
				"fmtBinary"))
}

// CopyInINumStrFormatter - Receives an incoming INumStrFormatter
// object, converts it to a FormatterBinary instance and proceeds
// to copy the the data fields into those of the current
// FormatterBinary instance.
//
// If the dynamic type of INumStrFormatter is not equal to
// FormatterBinary, an error will be returned. Likewise,
// if the data fields of input parameter 'incomingIFormatter' are
// judged to be invalid, an error will be returned.
//
// Be advised, all of the data fields in the current
// FormatterBinary instance will be overwritten.
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
//       dynamic type is not equal to FormatterBinary an error
//       will be returned.
//
//       The data values in this object will be copied to the
//       current FormatterBinary instance.
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
func (fmtBinary *FormatterBinary) CopyInINumStrFormatter(
	incomingIFormatter INumStrFormatter,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"CopyInINumStrFormatter()")

	if incomingIFormatter == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIFormatter' is "+
			"'nil'\n",
			ePrefix.String())
	}

	incomingFormatterBinary,
		isValid := incomingIFormatter.(*FormatterBinary)

	if !isValid {

		actualType :=
			reflect.TypeOf(incomingIFormatter)

		typeName := "Unknown"

		if actualType != nil {
			typeName = actualType.Name()
		}

		return fmt.Errorf("%v\n"+
			"Error: 'incomingIFormatter' is NOT Type "+
			"FormatterBinary.\n"+
			"'incomingIFormatter' is type %v",
			ePrefix.String(),
			typeName)
	}

	return formatterBinaryAtom{}.ptr().copyIn(
		fmtBinary,
		incomingFormatterBinary,
		ePrefix.XCtx(
			"incomingFormatterBinary->"+
				"fmtBinary"))
}

// CopyOut - Creates and returns a deep copy of the current
// FormatterBinary instance.
//
// If the current FormatterBinary instance is judged to be
// invalid, this method will return an error.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  FormatterBinary
//     - If this method completes successfully, a new instance of
//       FormatterBinary will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterBinary.
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
func (fmtBinary *FormatterBinary) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterBinary,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"CopyOut()")

	return formatterBinaryAtom{}.ptr().copyOut(
		fmtBinary,
		ePrefix.XCtx(
			"Copy Out from 'fmtBinary'"))
}

// CopyOutINumStrFormatter - Creates and returns a deep copy of the current
// FormatterBinary instance as an INumStrFormatter object.
//
// If the current FormatterBinary instance is judged to be
// invalid, this method will return an error.
//
// This method is required by interface INumStrFormatter.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  INumStrFormatter
//     - If this method completes successfully, a new instance of
//       FormatterBinary will be created, converted to an
//       instance of interface INumStrFormatter and returned
//       to the calling function. The returned data represents a
//       deep copy of the current FormatterBinary instance.
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
func (fmtBinary *FormatterBinary) CopyOutINumStrFormatter(
	ePrefix *ErrPrefixDto) (
	INumStrFormatter,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"CopyOutINumStrFormatter()")

	newFormatterBinary,
		err := formatterBinaryAtom{}.ptr().copyOut(
		fmtBinary,
		ePrefix.XCtx(
			"Copy Out from 'fmtBinary'"))

	return INumStrFormatter(&newFormatterBinary), err
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterBinary instance to
// their initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (fmtBinary *FormatterBinary) Empty() {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	_ = formatterBinaryQuark{}.ptr().
		empty(fmtBinary, nil)

	fmtBinary.lock.Unlock()

	fmtBinary.lock = nil
}

// GetFmtNumStr - Returns a number string formatted for binary
// digits based on the configuration encapsulated in the current
// instance of FormatterBinary.
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
//       numeric value to be formatted as binary digits. This array
//       of integer digits always represents a positive ('+')
//       numeric value. The array consists entirely of numeric
//       digits.
//
//
//  absValFractionalRunes         []rune
//     - An array of runes containing the fractional component of
//       the numeric value to be formatted as binary digits. This
//       array of numeric digits always represents a positive ('+')
//       numeric value. The array consists entirely of numeric
//       digits.
//
//
//  signVal                       int
//     - The parameter designates the numeric sign of the final
//       formatted binary numeric value returned by this method.
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
//       binary number string will be returned.
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
func (fmtBinary *FormatterBinary) GetFmtNumStr(
	absValIntegerRunes []rune,
	absValFractionalRunes []rune,
	signVal int,
	baseNumSys BaseNumberSystemType,
	ePrefix *ErrPrefixDto) (
	fmtNumStr string,
	err error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"GetFmtNumStr()")

	if baseNumSys != BaseNumberSystemType(0).Binary() {
		err = fmt.Errorf("%v\n"+
			"Error: Base Numbering System is NOT equal to Base-2!\n"+
			"baseNumSys=='%v'\n",
			ePrefix.String(),
			baseNumSys.XValueInt())
		return fmtNumStr, err
	}

	if signVal < 0 {
		fmtNumStr = "-"
	}

	fmtNumStr += string(absValIntegerRunes) +
		string(fmtBinary.numericSeparators.decimalSeparators) +
		string(absValFractionalRunes)

	return fmtNumStr, err
}

// GetNumberFieldDto - Returns a deep copy of the NumberFieldDto
// object configured for the current FormatterBinary instance.
//
// The NumberFieldDto object specifies the length of the number
// field in which the binary numeric value will be displayed
// and justified left, right or center according to the
// specification.
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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
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
//       used to configure the current instance of FormatterBinary.
//
//       The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto object contains specifications for number
//       field length. Typically, the length of a number field is
//       greater than the length of the number string which will be
//       inserted and displayed within the number field.
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
//       (2) 'Left-Justification'  - "NumberString        "
//       (3) 'Centered'            - "    NumberString    "
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
//       judged to be invalid, this method will return an error.
//
func (fmtBinary *FormatterBinary) GetNumberFieldDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"GetNumericSeparators()")

	return fmtBinary.numFieldDto.CopyOut(
		ePrefix.XCtx("fmtBinary.numFieldDto->"))
}

// GetNumericSeparators - Returns a deep copy of the
// NumericSeparators instance currently configured for this
// Binary Format Specification.
//
// The Numeric Separators object contains the decimal separator
// and the integer digit separators. It is used to specify the
// separator characters which will be included in number string
// text displays and is defined as follows:
//
//  type NumericSeparators struct {
//   decimalSeparators    []rune
//   integerSeparatorsDto NumStrIntSeparatorsDto
//  }
//
//  decimalSeparators     []rune
//
//  The 'Decimal Separator' is used to separate integer and
//  fractional digits within a floating point number display.
//  The decimal separator may consist of one or more runes.
//
//  integerSeparatorsDto    NumStrIntSeparatorsDto
//
//  The NumStrIntSeparatorsDto type encapsulates the integer digits
//  separators, often referred to as the 'Thousands Separator'.
//  Integer digit separators are used to separate integers into
//  specific groups within a number string. The
//  NumStrIntSeparatorsDto manages an array or collection of
//  NumStrIntSeparator objects.
//
//         type NumStrIntSeparatorsDto struct {
//           intSeparators []NumStrIntSeparator
//         }
//
//         type NumStrIntSeparator struct {
//          intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//          intSeparatorGrouping    uint    // Number of integer digits in a group
//          intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                          // A zero value signals unlimited repetitions.
//          restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//         }
//
//         intSeparatorChars          []rune
//            - A series of runes or characters used to separate integer
//              digits in a number string. These characters are commonly
//              known as the 'thousands separator'. A 'thousands
//              separator' is used to separate groups of integer digits to
//              the left of the decimal separator (a.k.a. decimal point).
//              In the United States, the standard integer digits
//              separator is the single comma character (','). Other
//              countries and cultures use periods, spaces, apostrophes or
//              multiple characters to separate integers.
//                    United States Example:  1,000,000,000
//
//         intSeparatorGrouping       uint
//            - This unsigned integer values specifies the number of
//              integer digits within a group. This value is used to group
//              integers within a number string.
//
//              In most western countries integer digits to the left of
//              the decimal separator (a.k.a. decimal point) are separated
//              into groups of three digits representing a grouping of
//              'thousands' like this: '1,000,000,000'. In this case the
//              intSeparatorGrouping value would be set to three ('3').
//
//              In some countries and cultures other integer groupings are
//              used. In India, for example, a number might be formatted
//              like this: '6,78,90,00,00,00,00,000'.
//
//         intSeparatorRepetitions    uint
//            - This unsigned integer value specifies the number of times
//              this integer grouping is repeated. A value of zero signals
//              that this integer grouping will be repeated indefinitely.
//
//         restartIntGroupingSequence bool
//            - If the NumStrIntSeparator is the last element in an array
//              of NumStrIntSeparator objects, this boolean flag signals
//              whether the entire integer grouping sequence will be
//              restarted from array element zero.
//
// The returned NumericSeparators object represents the Numeric
// Separator values used to configure the current instance of
// FormatterBinary.
//
// If either the FormatterBinary or NumericSeparators object
// is judged to be invalid, this method will return an error.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be returned through this
//       parameter. This object is a deep copy of the Numeric
//       Separator information used to configure the current
//       instance of FormatterBinary.
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
//       If the FormatterBinary or NumericSeparators object is
//       judged to be invalid, this method will return an error.
//
func (fmtBinary *FormatterBinary) GetNumericSeparators(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"GetNumericSeparators()")

	_,
		err := formatterBinaryQuark{}.ptr().
		testValidityOfFormatterBinary(
			fmtBinary,
			ePrefix.XCtx(
				"fmtBinary"))

	if err != nil {
		return NumericSeparators{}, err
	}

	return fmtBinary.numericSeparators.CopyOut(
		ePrefix.XCtx(
			"fmtBinary.numericSeparators"))
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for FormatterBinary
// objects is NumStrFormatTypeCode(0).Binary().
//
// This method is required by interface INumStrFormatter.
//
func (fmtBinary *FormatterBinary) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	return fmtBinary.numStrFmtType
}

// GetTurnOnIntegerDigitsSeparationFlag - Returns the boolean flag
// which signals whether integer digits within a number string will
// be grouped by thousands and separated by an integer digits
// separator such as a comma (',').
//
// If this flag is set to 'true', integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1,000,000,000,000
//
// If this flag is set to 'false',  integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1000000000000
//
// Numeric Separators such as decimal and integer digit separators
// are controlled by the member variable,
// FormatterBinary.numericSeparators.
//
func (fmtBinary *FormatterBinary) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	return fmtBinary.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterBinary instance to determine whether that instance is
// valid in all respects.
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
//     - This returned boolean value will signal whether the
//       current FormatterBinary is valid, or not. If the current
//       FormatterBinary instance contains valid data, this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
func (fmtBinary *FormatterBinary) IsValidInstance() bool {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	isValid,
		_ := formatterBinaryQuark{}.ptr().
		testValidityOfFormatterBinary(
			fmtBinary,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterBinary instance to determine whether that
// instance is valid in all respects.
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
//     - If the current instance of FormatterBinary contains invalid
//       data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (fmtBinary *FormatterBinary) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPrefCtx("FormatterBinary."+
		"IsValidInstanceError()",
		"Testing Validity of 'fmtBinary'")

	_,
		err := formatterBinaryQuark{}.ptr().
		testValidityOfFormatterBinary(
			fmtBinary,
			nil)

	return err
}

// NewBasic - Creates and returns a new instance of
// FormatterBinary. This method specifies the minimum number of
// input parameters required to construct a new instance of
// FormatterBinary. Default values are used to supplement these
// input parameters.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.NewBasicRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterBinary.NewWithComponents()'.
//
// The member variable 'FormatterBinary.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         string
//     - The characters or character used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars       string
//     - The characters or character which will be used to delimit
//       'thousands' in integer number strings. In the United
//       States, the Thousands separator is the comma character
//       (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterBinary.NewWithComponents().
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
//  FormatterBinary
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterBinary.
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
func (fmtBinary FormatterBinary) NewBasic(
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterBinary,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"NewBasic()")

	newFormatterBinary := FormatterBinary{}

	err := formatterBinaryUtility{}.ptr().
		setBasic(
			&newFormatterBinary,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterBinary"))

	return newFormatterBinary, err
}

// NewBasicRunes - Creates and returns a new instance of
// FormatterBinary. This method specifies the minimum number of
// input parameters required to construct a new instance of
// FormatterBinary. Default values are used to supplement these
// input parameters.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.NewBasic()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterBinary.NewWithComponents()'.
//
// The member variable 'FormatterBinary.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         []rune
//     - The characters or character used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars       []rune
//     - The characters or character which will be used to delimit
//       'thousands' in integer number strings. In the United
//       States, the Thousands separator is the comma character
//       (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterBinary.NewWithComponents().
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
//  FormatterBinary
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterBinary.
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
func (fmtBinary FormatterBinary) NewBasicRunes(
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterBinary,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"NewBasic()")

	newFormatterBinary := FormatterBinary{}

	err := formatterBinaryUtility{}.ptr().
		setBasicRunes(
			&newFormatterBinary,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterBinary"))

	return newFormatterBinary, err
}

// NewDetail - Creates and returns a new instance of
// FormatterBinary generated from the input parameters described
// below.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.NewDetailRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.NewWithComponents()'
//
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         string
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period (".") or Decimal Point.
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
//  FormatterBinary
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterBinary.
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
func (fmtBinary FormatterBinary) NewDetail(
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterBinary,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"NewDetail()")

	newFormatterBinary := FormatterBinary{}

	err := formatterBinaryUtility{}.ptr().
		setDetail(
			&newFormatterBinary,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterBinary"))

	return newFormatterBinary, err
}

// NewDetailRunes - Creates and returns a new instance of
// FormatterBinary generated from the input parameters described
// below.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.NewDetail()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.NewWithComponents()'
//
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         string
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period (".") or Decimal Point.
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
//  FormatterBinary
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterBinary.
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
func (fmtBinary FormatterBinary) NewDetailRunes(
	decimalSeparatorChars []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterBinary,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"formatterBinaryUtility." +
			"NewDetailRunes()")

	newFormatterBinary := FormatterBinary{}

	err := formatterBinaryUtility{}.ptr().
		setDetailRunes(
			&newFormatterBinary,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterBinary"))

	return newFormatterBinary, err
}

// NewUnitedStatesDefaults - Creates and returns a new instance of
// FormatterBinary.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
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
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         FormatterBinary(0).Binary()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  -- None --
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  FormatterBinary
//     - This parameter will return a new, populated instance of
//       FormatterBinary configured for United States default
//       binary format parameters.
//
func (fmtBinary FormatterBinary) NewUnitedStatesDefaults() FormatterBinary {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	newFormatterBinary := FormatterBinary{}

	_ = formatterBinaryUtility{}.ptr().setToUnitedStatesDefaults(
		&newFormatterBinary,
		nil)

	return newFormatterBinary
}

// NewWithComponents - Creates and returns a new instance of
// FormatterBinary.
//
// The FormatterBinary type encapsulates the configuration
// parameters necessary to format binary numeric values for
// display in text number strings.
//
// This method requires detailed input parameters which provide
// granular control over all data fields contained in the returned
// new instance of FormatterBinary.
//
// For a 'New' method using minimum input parameters coupled
// with default values, see:
//      FormatterBinary.NewBasic()
//
// The member variable 'FormatterBinary.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  numericSeparators        NumericSeparators
//     - This instance of NumericSeparators is used to specify the
//       separator characters which will be included in number
//       string text displays.
//
//        type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//        }
//
//        decimalSeparators     []rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//        The decimal separator may consist of one or more runes.
//
//        integerSeparatorsDto    NumStrIntSeparatorsDto
//
//        The NumStrIntSeparatorsDto type encapsulates the integer digits
//        separators, often referred to as the 'Thousands Separator'.
//        Integer digit separators are used to separate integers into
//        specific groups within a number string. The
//        NumStrIntSeparatorsDto manages an array or collection of
//        NumStrIntSeparator objects.
//
//               type NumStrIntSeparatorsDto struct {
//                 intSeparators []NumStrIntSeparator
//               }
//
//               type NumStrIntSeparator struct {
//                intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//                intSeparatorGrouping    uint    // Number of integer digits in a group
//                intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                                // A zero value signals unlimited repetitions.
//                restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//               }
//
//               intSeparatorChars          []rune
//                  - A series of runes or characters used to separate integer
//                    digits in a number string. These characters are commonly
//                    known as the 'thousands separator'. A 'thousands
//                    separator' is used to separate groups of integer digits to
//                    the left of the decimal separator (a.k.a. decimal point).
//                    In the United States, the standard integer digits
//                    separator is the single comma character (','). Other
//                    countries and cultures use periods, spaces, apostrophes or
//                    multiple characters to separate integers.
//                          United States Example:  1,000,000,000
//
//               intSeparatorGrouping       uint
//                  - This unsigned integer values specifies the number of
//                    integer digits within a group. This value is used to group
//                    integers within a number string.
//
//                    In most western countries integer digits to the left of
//                    the decimal separator (a.k.a. decimal point) are separated
//                    into groups of three digits representing a grouping of
//                    'thousands' like this: '1,000,000,000'. In this case the
//                    intSeparatorGrouping value would be set to three ('3').
//
//                     In some countries and cultures other integer groupings
//                     are used. In India, for example, a number might be
//                     formatted like this: '6,78,90,00,00,00,00,000'. Chinese
//                     Numerals have an integer grouping value of four ('4').
//                         Chinese Numerals Example: '12,3456,7890,2345'
//
//               intSeparatorRepetitions    uint
//                  - This unsigned integer value specifies the number of times
//                    this integer grouping is repeated. A value of zero signals
//                    that this integer grouping will be repeated indefinitely.
//
//               restartIntGroupingSequence bool
//                  - If the NumStrIntSeparator is the last element in an array
//                    of NumStrIntSeparator objects, this boolean flag signals
//                    whether the entire integer grouping sequence will be
//                    restarted from array element zero.
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
//  FormatterBinary
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterBinary.
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
func (fmtBinary FormatterBinary) NewWithComponents(
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	FormatterBinary,
	error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"NewWithComponents()")

	newFormatterBinary := FormatterBinary{}

	err := formatterBinaryMechanics{}.ptr().setWithComponents(
		&newFormatterBinary,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"newFormatterBinary"))

	return newFormatterBinary, err
}

// SetBasic - This method will set all of the member variable data
// values for the current instance of FormatterBinary. The input
// parameters represent the minimum information required to
// configure a FormatterBinary object.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.SetBasicRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// The member variable 'FormatterBinary.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterBinary instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         string
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars       string
//     - The character or characters used to delimit 'thousands' in
//       integer number strings. In the United States, the
//       Thousands separator is the comma character (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterBinary.SetWithComponents().
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
func (fmtBinary *FormatterBinary) SetBasic(
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetBasic()")

	return formatterBinaryUtility{}.ptr().
		setBasic(
			fmtBinary,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"fmtBinary"))
}

// SetBasicRunes - This method will set all of the member variable data
// values for the current instance of FormatterBinary. The input
// parameters represent the minimum information required to
// configure a FormatterBinary object.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.SetBasic()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterBinary, reference method:
//   'FormatterBinary.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// The member variable 'FormatterBinary.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterBinary instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChars         []rune
//     - The character or characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the Decimal Separator character is the
//       period ('.') or Decimal Point.
//           United States Example: '123.45678'
//
//
//  thousandsSeparatorChars       []rune
//     - The character or characters used to delimit 'thousands' in
//       integer number strings. In the United States, the
//       Thousands separator is the comma character (',').
//           United States Example: '1,000,000,000'
//
//       The default integer digit grouping of three ('3') digits
//       is applied with this separator character. An integer digit
//       grouping of three ('3') results in thousands grouping.
//           United States Example: '1,000,000,000'
//
//       For custom integer digit grouping, use method
//       FormatterBinary.SetWithComponents().
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
func (fmtBinary *FormatterBinary) SetBasicRunes(
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetBasicRunes()")

	return formatterBinaryUtility{}.ptr().
		setBasicRunes(
			fmtBinary,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"fmtBinary"))
}

// SetDetail - This method will set all of the member variable data
// values for the current instance of FormatterBinary. The new
// data values are derived from the input parameters described
// below.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.SetDetailRunes()
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
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterBinary instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
func (fmtBinary *FormatterBinary) SetDetail(
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetDetail()")

	return formatterBinaryUtility{}.ptr().
		setDetail(
			fmtBinary,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterBinary"))
}

// SetDetailRunes - This method will set all of the member variable
// data values for the current instance of FormatterBinary. The new
// data values are derived from the input parameters described
// below.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format binary digits for display in text number
// strings.
//
// This method differs from method FormatterBinary.SetDetail()
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
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterBinary instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
func (fmtBinary *FormatterBinary) SetDetailRunes(
	decimalSeparatorChars []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetDetailRunes()")

	return formatterBinaryUtility{}.ptr().
		setDetailRunes(
			fmtBinary,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix.XCtx(
				"newFormatterBinary"))
}

// SetNumberFieldDto - Sets the Number Field Dto object configured
// for the current FormatterBinary instance.
//
// The Number Field Dto object is used to specify the length and
// string justification characteristics used to display binary
// number strings within a number field.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldDto      NumberFieldDto
//     - The NumberFieldDto object contains formatting instructions
//       for the creation and implementation of a number field.
//       Number fields are text strings which contain number strings
//       for use in text displays.
//
//       The NumberFieldDto object contains specifications for number
//       field length. Typically, the length of a number field is
//       greater than the length of the number string which will be
//       inserted and displayed within the number field.
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
//       (2) 'Left-Justification'  - "NumberString        "
//       (3) 'Centered'            - "    NumberString    "
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
func (fmtBinary *FormatterBinary) SetNumberFieldDto(
	numberFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetNumberFieldDto()")

	return fmtBinary.numFieldDto.CopyIn(
		&numberFieldDto,
		ePrefix.XCtx(
			"numberFieldDto->"+
				"fmtCurr.numFieldDto"))
}

// SetNumericSeparators - Sets the Numeric Separators object
// for the current FormatterBinary instance.
//
// The Numeric Separators object is used to specify the Decimal
// Separator Characters and the Integer Digits Separator
// Characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numericSeparators        NumericSeparators
//     - This instance of NumericSeparators is used to specify the
//       separator characters which will be included in number
//       string text displays.
//
//        type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//        }
//
//        decimalSeparators     []rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//        The decimal separator may consist of one or more runes.
//
//        integerSeparatorsDto    NumStrIntSeparatorsDto
//
//        The NumStrIntSeparatorsDto type encapsulates the integer digits
//        separators, often referred to as the 'Thousands Separator'.
//        Integer digit separators are used to separate integers into
//        specific groups within a number string. The
//        NumStrIntSeparatorsDto manages an array or collection of
//        NumStrIntSeparator objects.
//
//               type NumStrIntSeparatorsDto struct {
//                 intSeparators []NumStrIntSeparator
//               }
//
//               type NumStrIntSeparator struct {
//                intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//                intSeparatorGrouping    uint    // Number of integer digits in a group
//                intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                                // A zero value signals unlimited repetitions.
//                restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//               }
//
//               intSeparatorChars          []rune
//                  - A series of runes or characters used to separate integer
//                    digits in a number string. These characters are commonly
//                    known as the 'thousands separator'. A 'thousands
//                    separator' is used to separate groups of integer digits to
//                    the left of the decimal separator (a.k.a. decimal point).
//                    In the United States, the standard integer digits
//                    separator is the single comma character (','). Other
//                    countries and cultures use periods, spaces, apostrophes or
//                    multiple characters to separate integers.
//                          United States Example:  1,000,000,000
//
//               intSeparatorGrouping       uint
//                  - This unsigned integer values specifies the number of
//                    integer digits within a group. This value is used to group
//                    integers within a number string.
//
//                    In most western countries integer digits to the left of
//                    the decimal separator (a.k.a. decimal point) are separated
//                    into groups of three digits representing a grouping of
//                    'thousands' like this: '1,000,000,000'. In this case the
//                    intSeparatorGrouping value would be set to three ('3').
//
//                    In some countries and cultures other integer groupings are
//                    used. In India, for example, a number might be formatted
//                    like this: '6,78,90,00,00,00,00,000'.
//
//               intSeparatorRepetitions    uint
//                  - This unsigned integer value specifies the number of times
//                    this integer grouping is repeated. A value of zero signals
//                    that this integer grouping will be repeated indefinitely.
//
//               restartIntGroupingSequence bool
//                  - If the NumStrIntSeparator is the last element in an array
//                    of NumStrIntSeparator objects, this boolean flag signals
//                    whether the entire integer grouping sequence will be
//                    restarted from array element zero.
//
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
func (fmtBinary *FormatterBinary) SetNumericSeparators(
	numericSeparators NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetNumericSeparators()")

	return formatterBinaryMechanics{}.ptr().setNumericSeparators(
		fmtBinary,
		numericSeparators,
		ePrefix.XCtx(
			"fmtBinary"))
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterBinary object. For Binary formatters,
// the Number String Format Type Code is set to
// NumStrFormatTypeCode(0).Binary().
//
// This method is required by interface INumStrFormatter.
//
func (fmtBinary *FormatterBinary) SetNumStrFormatTypeCode() {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	fmtBinary.numStrFmtType = NumStrFormatTypeCode(0).Binary()
}

// SetToUnitedStatesDefaults - Sets the member variable data values
// for the current FormatterBinary instance to United States
// Default values.
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
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         FormatterBinary(0).Binary()
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
func (fmtBinary *FormatterBinary) SetToUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetToUnitedStatesDefaults()")

	err = formatterBinaryUtility{}.ptr().
		setToUnitedStatesDefaults(
			fmtBinary,
			nil)

	return err
}

// SetToUnitedStatesDefaultsIfEmpty -If the current FormatterBinary
// instance is empty or invalid, this method will set the member
// variable data values to United States default values.
//
// If the current FormatterBinary instance is valid and NOT empty,
// this method will take no action and exit.
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
// The member variable 'FormatterBinary.numStrFmtType' is
// automatically defaulted to:
//         FormatterBinary(0).Binary()
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
func (fmtBinary *FormatterBinary) SetToUnitedStatesDefaultsIfEmpty(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetToUnitedStatesDefaultsIfEmpty()")

	isValid,
		_ := formatterBinaryQuark{}.ptr().
		testValidityOfFormatterBinary(
			fmtBinary,
			nil)

	if isValid {
		return err
	}

	err = formatterBinaryUtility{}.ptr().
		setToUnitedStatesDefaults(
			fmtBinary,
			nil)

	return err
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// FormatterBinary.
//
// This boolean flag signals whether integer digits within a number
// string will be grouped by thousands and separated by an integer
// digits separator such as a comma (',').
//
// If this flag is set to 'true', integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1,000,000,000,000
//
// If this flag is set to 'false',  integer digits will be presented
// in text number strings as shown in the following example.
//  Example: 1000000000000
//
// Numeric Separators such as decimal and integer digit separators
// are controlled by the member variable,
// FormatterBinary.numericSeparators.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//            Example: 1,000,000,000
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: 1000000000
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (fmtBinary *FormatterBinary) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	fmtBinary.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}

// SetWithComponents - This method will overwrite and set all data
// data values for the current instance of FormatterBinary.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format numeric binary numeric values for display in
// text number strings.
//
// This method requires detailed input parameters to control
// configuration for all member variables in the current instance
// of FormatterBinary. For a similar method using minimum input
// parameters coupled with default values, see:
//      FormatterBinary.SetBasic()
//
// The member variable 'FormatterBinary.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).Binary()
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterBinary instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  numericSeparators        NumericSeparators
//     - This instance of NumericSeparators is used to specify the
//       separator characters which will be included in the number
//       string text displays.
//
//        type NumericSeparators struct {
//         decimalSeparators    []rune
//         integerSeparatorsDto NumStrIntSeparatorsDto
//        }
//
//        decimalSeparators     []rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//        The decimal separator may consist of one or more runes.
//
//        integerSeparatorsDto    NumStrIntSeparatorsDto
//
//        The NumStrIntSeparatorsDto type encapsulates the integer digits
//        separators, often referred to as the 'Thousands Separator'.
//        Integer digit separators are used to separate integers into
//        specific groups within a number string. The
//        NumStrIntSeparatorsDto manages an array or collection of
//        NumStrIntSeparator objects.
//
//               type NumStrIntSeparatorsDto struct {
//                 intSeparators []NumStrIntSeparator
//               }
//
//               type NumStrIntSeparator struct {
//                intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//                intSeparatorGrouping    uint    // Number of integer digits in a group
//                intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                                // A zero value signals unlimited repetitions.
//                restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//               }
//
//               intSeparatorChars          []rune
//                  - A series of runes or characters used to separate integer
//                    digits in a number string. These characters are commonly
//                    known as the 'thousands separator'. A 'thousands
//                    separator' is used to separate groups of integer digits to
//                    the left of the decimal separator (a.k.a. decimal point).
//                    In the United States, the standard integer digits
//                    separator is the single comma character (','). Other
//                    countries and cultures use periods, spaces, apostrophes or
//                    multiple characters to separate integers.
//                          United States Example:  1,000,000,000
//
//               intSeparatorGrouping       uint
//                  - This unsigned integer values specifies the number of
//                    integer digits within a group. This value is used to group
//                    integers within a number string.
//
//                    In most western countries integer digits to the left of
//                    the decimal separator (a.k.a. decimal point) are separated
//                    into groups of three digits representing a grouping of
//                    'thousands' like this: '1,000,000,000'. In this case the
//                    intSeparatorGrouping value would be set to three ('3').
//
//                     In some countries and cultures other integer groupings
//                     are used. In India, for example, a number might be
//                     formatted like this: '6,78,90,00,00,00,00,000'. Chinese
//                     Numerals have an integer grouping value of four ('4').
//                         Chinese Numerals Example: '12,3456,7890,2345'
//
//               intSeparatorRepetitions    uint
//                  - This unsigned integer value specifies the number of times
//                    this integer grouping is repeated. A value of zero signals
//                    that this integer grouping will be repeated indefinitely.
//
//               restartIntGroupingSequence bool
//                  - If the NumStrIntSeparator is the last element in an array
//                    of NumStrIntSeparator objects, this boolean flag signals
//                    whether the entire integer grouping sequence will be
//                    restarted from array element zero.
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
func (fmtBinary *FormatterBinary) SetWithComponents(
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtBinary.lock == nil {
		fmtBinary.lock = new(sync.Mutex)
	}

	fmtBinary.lock.Lock()

	defer fmtBinary.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterBinary." +
			"SetWithComponents()")

	return formatterBinaryMechanics{}.ptr().setWithComponents(
		fmtBinary,
		turnOnIntegerDigitsSeparation,
		numericSeparators,
		numFieldDto,
		ePrefix.XCtx(
			"fmtBinary"))
}
