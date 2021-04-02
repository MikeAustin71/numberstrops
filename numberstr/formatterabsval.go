package numberstr

import (
	"fmt"
	"reflect"
	"sync"
)

// FormatterAbsoluteValue - This structure stores formatting data
// for Absolute Value numbers displayed in text number strings.
// Member data elements are listed as follows:
//
//  numStrFmtType                 NumStrFormatTypeCode
//     - An enumeration value automatically set to:
//           NumStrFormatTypeCode(0).AbsoluteValue()
//
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//            Example: 1,000,000,000
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: 1000000000
//
//
//  numericSeparators             NumericSeparators
//     - This instance of 'NumericSeparators' is used to specify
//       the separator characters which will be included in number
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
//
type FormatterAbsoluteValue struct {
	numStrFmtType                 NumStrFormatTypeCode
	absoluteValFmt                string
	turnOnIntegerDigitsSeparation bool
	numericSeparators             NumericSeparators
	numFieldDto                   NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// FormatterAbsoluteValue ('incomingFormatterAbsVal') to
// the corresponding data fields of the current
// FormatterAbsoluteValue instance.
//
// If 'incomingFormatterAbsVal' is judged to be invalid, this
// method will return an error.
//
// Be advised, all of the data fields in the current
// FormatterAbsoluteValue instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFormatterAbsVal    *FormatterAbsoluteValue
//     - A pointer to an instance of FormatterAbsoluteValue.
//       The data values in this object will be copied to
//       corresponding data elements in the current
//       FormatterAbsoluteValue instance.
//
//
//  ePrefix                    *ErrPrefixDto
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
func (fmtAbsVal *FormatterAbsoluteValue) CopyIn(
	incomingFormatterAbsVal *FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.CopyIn()")

	return formatterAbsoluteValueNanobot{}.ptr().copyIn(
		fmtAbsVal,
		incomingFormatterAbsVal,
		ePrefix.XCtx(
			"incomingFormatterAbsVal -> fmtAbsVal"))
}

// CopyInINumStrFormatter - Receives an incoming INumStrFormatter
// object, converts it to a FormatterAbsoluteValue instance and
// proceeds to copy the the data fields into those of the
// current FormatterAbsoluteValue instance.
//
// If the dynamic type of INumStrFormatter is not equal to
// FormatterAbsoluteValue, an error will be returned. Likewise,
// if the data fields of input parameter 'incomingIFormatter' are
// judged to be invalid, an error will be returned.
//
// Be advised, all of the data fields in the current
// FormatterAbsoluteValue instance will be overwritten.
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
//       dynamic type is not equal to FormatterAbsoluteValue an
//       error will be returned.
//
//       The data values in this object will be copied to the
//       current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) CopyInINumStrFormatter(
	incomingIFormatter INumStrFormatter,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"CopyInINumStrFormatter()")

	if incomingIFormatter == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIFormatter' is "+
			"'nil'\n",
			ePrefix.String())
	}

	incomingFormatterAbsVal,
		isValid := incomingIFormatter.(*FormatterAbsoluteValue)

	if !isValid {

		actualType :=
			reflect.TypeOf(incomingIFormatter)

		typeName := "Unknown"

		if actualType != nil {
			typeName = actualType.Name()
		}

		return fmt.Errorf("%v\n"+
			"Error: 'incomingIFormatter' is NOT Type "+
			"FormatterAbsoluteValue\n"+
			"'incomingIFormatter' is type %v",
			ePrefix.String(),
			typeName)
	}

	return formatterAbsoluteValueNanobot{}.ptr().copyIn(
		fmtAbsVal,
		incomingFormatterAbsVal,
		ePrefix.XCtx("incomingFormatterAbsVal -> fmtAbsVal"))
}

// CopyOut - Returns a deep copy of the current
// FormatterAbsoluteValue instance.
//
// If the current FormatterAbsoluteValue instance is judged to be
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, a new instance of
//       FormatterAbsoluteValue will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterAbsoluteValue.
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
func (fmtAbsVal *FormatterAbsoluteValue) CopyOut(
	ePrefix *ErrPrefixDto) (FormatterAbsoluteValue, error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.CopyOut()")

	return formatterAbsoluteValueNanobot{}.ptr().
		copyOut(
			fmtAbsVal,
			ePrefix.XCtx(
				"Copy -> 'fmtAbsVal'"))
}

// CopyOutINumStrFormatter - Creates and returns a deep copy of the
// current FormatterAbsoluteValue instance as an INumStrFormatter
// object.
//
// If the current FormatterAbsoluteValue instance is judged to be
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
//       FormatterAbsoluteValue will be created, converted to an
//       instance of interface INumStrFormatter and returned
//       to the calling function. The returned data represents a
//       deep copy of the current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) CopyOutINumStrFormatter(
	ePrefix *ErrPrefixDto) (INumStrFormatter, error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"CopyOutINumStrFormatter()")

	newFormatterAbsVal,
		err := formatterAbsoluteValueNanobot{}.ptr().
		copyOut(
			fmtAbsVal,
			ePrefix.XCtx(
				"Copy -> 'fmtAbsVal'"))

	return INumStrFormatter(&newFormatterAbsVal), err
}

// Empty - Deletes and resets data values for all member variables
// in the current FormatterAbsoluteValue instance to their initial
// 'zero' values.
//
func (fmtAbsVal *FormatterAbsoluteValue) Empty() {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	_ = formatterAbsoluteValueQuark{}.ptr().
		empty(fmtAbsVal, nil)

	fmtAbsVal.lock.Unlock()

	fmtAbsVal.lock = nil
}

// Equal - Receives a FormatterAbsoluteValue object
// ('fmtAbsValTwo') and proceeds to determine whether all data
// elements in this object are equal to all corresponding data
// elements in the current instance of FormatterAbsoluteValue.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtAbsValTwo        FormatterAbsoluteValue
//     - This method will compare all data elements in the current
//       FormatterAbsoluteValue instance to corresponding data elements
//       for this second FormatterAbsoluteValue object to determine
//       if they are equivalent.
//
//
//  ePrefix             *ErrPrefixDto
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
//  bool
//     - If all the data elements in the current
//       FormatterAbsoluteValue instance are equal to all the
//       corresponding data elements in 'fmtAbsValTwo', this return
//       parameter will be set to 'true'. If all the data elements
//       are NOT equal, this return parameter will be set to
//       'false'.
//
//
//  error
//     - If all the data elements in the current
//       FormatterAbsoluteValue are equal to all the corresponding
//       data elements in 'fmtAbsValTwo', this return parameter
//       will be set to 'nil'.
//
//       If the corresponding data elements are NOT equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (fmtAbsVal *FormatterAbsoluteValue) Equal(
	fmtAbsValTwo FormatterAbsoluteValue,
	ePrefix *ErrPrefixDto) (
	bool,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"Equal()")

	return formatterAbsoluteValueNanobot{}.ptr().
		equal(
			fmtAbsVal,
			&fmtAbsValTwo,
			ePrefix.XCtx(
				"fmtAbsVal vs fmtAbsValTwo"))
}

// GetAbsoluteValueFormat - Returns the formatting string used to
// format absolute numeric values in text number strings.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetAbsoluteValueFormat() string {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.absoluteValFmt
}

// GetDecimalSeparators - Returns the unicode character(s) (runes)
// used to separate integer and fractional digits in a floating
// point number.
//
// In the United States, the Decimal Separator character is the
// decimal point or period ('.').
//
//  Example:   123.45
//
// Decimal Separator is extracted from the underlying member
// variable, 'nStrFmtAbsValDto.numericSeparators'.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetDecimalSeparators() []rune {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.
		numericSeparators.
		GetDecimalSeparators()
}

// GetFmtNumStr - Returns a number string formatted for absolute
// values based on the configuration encapsulated in the current
// instance of FormatterAbsoluteValue.
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
//       numeric value to be formatted as an absolute value. This
//       array of integer digits always represents a positive ('+')
//       numeric value. The array consists entirely of numeric
//       digits.
//
//
//  absValFractionalRunes         []rune
//     - An array of runes containing the fractional component of
//       the numeric value to be formatted as an absolute value.
//       This array of numeric digits always represents a positive
//       ('+') numeric value. The array consists entirely of
//       numeric digits.
//
//
//  signVal                       int
//     - The parameter designates the numeric sign of the final
//       formatted absolute value returned by this method.
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
//       absolute value number string will be returned.
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
func (fmtAbsVal *FormatterAbsoluteValue) GetFmtNumStr(
	absValIntegerRunes []rune,
	absValFractionalRunes []rune,
	signVal int,
	baseNumSys BaseNumberSystemType,
	ePrefix *ErrPrefixDto) (
	fmtNumStr string,
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
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
		string(fmtAbsVal.numericSeparators.decimalSeparators) +
		string(absValFractionalRunes)

	return fmtNumStr, err
}

// GetIntegerDigitSeparators - Returns an array of type
// NumStrIntSeparator. The data contained in type
// NumStrIntSeparator is used to separate integer digits.
//
// The returned integer digit separators are those configured
// for the current instance of FormatterAbsoluteValue.
//
// The integer digit separators is also known as the 'thousands'
// separator. In the United States the standard integer digit
// separator character is the comma (',') and integers are shown
// in groups of three ('3').
//
//    United States Example: 1,000,000,000,000
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
//  NumStrIntSeparatorsDto
//     - Returns a deep copy of the NumStrIntSeparatorsDto object
//       configured for the current FormatterAbsoluteValue
//       instance.
//
//       The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       Taken as a whole, this collection represents the
//       specification controlling the integer separation
//       operation in absolute value number strings.
//
//        type NumStrIntSeparator struct {
//            intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//            intSeparatorGrouping    uint    // Number of integer digits in a group
//            intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                            // A zero value signals unlimited repetitions.
//            restartIntGroupingSequence bool // If true, the entire grouping sequence is repeated
//                                            //  beginning at array index zero.
//        }
//
//        intSeparatorChars          []rune
//           - A series of runes or characters used to separate integer
//             digits in a number string. These characters are commonly
//             known as the 'thousands separator'. A 'thousands
//             separator' is used to separate groups of integer digits to
//             the left of the decimal separator (a.k.a. decimal point).
//             In the United States, the standard integer digits
//             separator is the single comma character (','). Other
//             countries and cultures use periods, spaces, apostrophes or
//             multiple characters to separate integers.
//                   United States Example:  1,000,000,000
//
//        intSeparatorGrouping       uint
//           - This unsigned integer values specifies the number of
//             integer digits within a group. This value is used to group
//             integers within a number string.
//
//             In most western countries integer digits to the left of
//             the decimal separator (a.k.a. decimal point) are separated
//             into groups of three digits representing a grouping of
//             'thousands' like this: '1,000,000,000'. In this case the
//             intSeparatorGrouping value would be set to three ('3').
//
//             In some countries and cultures other integer groupings are
//             used. In India, for example, a number might be formatted like
//             this: '6,78,90,00,00,00,00,000'. Chinese Numerals have an
//             integer grouping value of four ('4').
//                Chinese Numerals Example: '12,3456,7890,2345'
//
//        intSeparatorRepetitions    uint
//           - This unsigned integer value specifies the number of times
//             this integer grouping is repeated. A value of zero signals
//             that this integer grouping will be repeated indefinitely.
//
//        restartIntGroupingSequence bool
//           - If the NumStrIntSeparator is the last element in an array
//             of NumStrIntSeparator objects, this boolean flag signals
//             whether the entire integer grouping sequence will be
//             restarted from array element zero.
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
func (fmtAbsVal *FormatterAbsoluteValue) GetIntegerDigitSeparators(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"GetIntegerDigitSeparators() ")

	return fmtAbsVal.
		numericSeparators.
		GetIntegerDigitSeparators(
			ePrefix.XCtx(
				"fmtAbsVal.numericSeparators"))
}

// GetNumberFieldDto - Returns the NumberFieldDto object
// currently configured for this Number String Format Specification
// Absolute Value Dto.
//
// The NumberFieldDto details the length of the number field in
// which the absolute numeric value will be displayed and justified
// left, right or center according to the specification.
//
// If the NumberFieldDto object is judged to be invalid, this
// method will return an error.
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
//       used to configure the current instance of
//       FormatterAbsoluteValue.
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
//       Be advised - If the NumberFieldDto object is judged to be
//       invalid, this method will return an error.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetNumberFieldDto(
	ePrefix *ErrPrefixDto) (
	NumberFieldDto,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"GetNumberFieldDto()")

	return fmtAbsVal.numFieldDto.CopyOut(
		ePrefix.XCtx(
			"fmtAbsVal.numFieldDto->"))
}

// GetNumericSeparators - Returns a deep copy of the
// NumericSeparators object configured for the current
// FormatterAbsoluteValue instance.
//
// The Numeric Separators object contains the decimal separator
// and the integer digit separators. It is used to specify the
// separator characters which will be included in the number
// string text display and is defined as follows:
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
// FormatterAbsoluteValue.
//
// If either the FormatterAbsoluteValue or NumericSeparators object
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
//       instance of FormatterAbsoluteValue.
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
//       If the FormatterAbsoluteValue or NumericSeparators
//       object is judged to be invalid, this method will return
//       an error.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetNumericSeparators(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue.GetNumericSeparators()")

	return fmtAbsVal.numericSeparators.CopyOut(
		ePrefix.XCtx("fmtAbsVal.numericSeparators->"))
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for
// FormatterAbsoluteValue objects is:
//     NumStrFormatTypeCode(0).AbsoluteValue().
//
// This method is required by interface INumStrFormatter.
//
func (fmtAbsVal *FormatterAbsoluteValue) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.numStrFmtType
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
func (fmtAbsVal *FormatterAbsoluteValue) GetTurnOnIntegerDigitsSeparationFlag() bool {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	return fmtAbsVal.turnOnIntegerDigitsSeparation
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterAbsoluteValue instance to determine whether that
// instance is valid in all respects.
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
//     - If this method completes successfully, the returned
//       boolean value will signal whether the current
//       FormatterAbsoluteValue is valid, or not. If the current
//       FormatterAbsoluteValue contains valid data, this method
//       returns 'true'. If the data is invalid, this method
//       returns 'false'.
//
func (fmtAbsVal *FormatterAbsoluteValue) IsValidInstance() (
	isValid bool) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	isValid,
		_ = formatterAbsoluteValueMolecule{}.ptr().
		testValidityOfFormatterAbsoluteValue(
			fmtAbsVal,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterAbsoluteValue instance to determine whether
// that instance is valid in all respects.
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
//       If no error prefix information is needed, set this
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
func (fmtAbsVal *FormatterAbsoluteValue) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"IsValidInstanceError()")

	_,
		err := formatterAbsoluteValueMolecule{}.ptr().
		testValidityOfFormatterAbsoluteValue(
			fmtAbsVal,
			ePrefix.XCtx(
				"Testing Validity of 'fmtAbsVal'"))

	return err
}

// NewBasic - Creates and returns a new instance of
// FormatterAbsoluteValue. This method specifies the minimum
// number of input parameters required to construct a new instance
// of FormatterAbsoluteValue. Default values are used to
// supplement these input parameters.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterAbsoluteValue.NewWithComponents()'.
//
// The FormatterAbsoluteValue type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChar          string
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar        string
//     - The character which will be used to delimit 'thousands' in
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
//       FormatterAbsoluteValue.NewWithComponents().
//
//
//  turnOnThousandsSeparator   bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnThousandsSeparator' is a boolean flag
//       used to control the 'Thousands Separator'. When set to
//       'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewBasic(
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewBasic()")

	newNumStrFmtSpecAbsValueDto :=
		FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	err := formatterAbsoluteValueUtility{}.ptr().
		setBasic(
			&newNumStrFmtSpecAbsValueDto,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewBasicRunes - Creates and returns a new instance of
// FormatterAbsoluteValue. This method specifies the minimum
// number of input parameters required to construct a new instance
// of FormatterAbsoluteValue. Default values are used to
// supplement these input parameters.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.NewWithComponents()'
//
// This method automatically sets a default integer digits
// grouping sequence of '3'. This means that integers will
// be grouped by thousands.
//
//     Example: '1,000,000,000'
//
// To control and specify alternative integer digit groupings, use
// method 'FormatterAbsoluteValue.NewWithComponents()'.
//
// The FormatterAbsoluteValue type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorChar          []rune
//     - The character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the Decimal Separator character is the period
//       ('.') or Decimal Point.
//           Example: '123.45678'
//
//
//  thousandsSeparatorChar        []rune
//     - The character which will be used to delimit 'thousands' in
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
//       FormatterAbsoluteValue.NewWithComponents().
//
//
//  turnOnThousandsSeparator   bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnThousandsSeparator' is a boolean flag
//       used to control the 'Thousands Separator'. When set to
//       'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewBasicRunes(
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewBasic()")

	newNumStrFmtSpecAbsValueDto :=
		FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	err := formatterAbsoluteValueUtility{}.ptr().
		setBasicRunes(
			&newNumStrFmtSpecAbsValueDto,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewDetail - Creates and returns a new instance of
// FormatterAbsoluteValue generated from the input parameters
// described below.
//
// This method differs from method FormatterAbsoluteValue.NewDetailRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.NewWithComponents()'
//
// The FormatterAbsoluteValue type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
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
//       FormatterAbsoluteValue.NewWithComponents().
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewDetail(
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewDetail()")

	newNumStrFmtSpecAbsValueDto :=
		FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	err := formatterAbsoluteValueUtility{}.ptr().
		setDetail(
			&newNumStrFmtSpecAbsValueDto,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewDetailRunes - Creates and returns a new instance of
// FormatterAbsoluteValue generated from the input parameters
// described below.
//
// This method differs from method FormatterAbsoluteValue.NewDetail()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.NewWithComponents()'
//
// The FormatterAbsoluteValue type encapsulates the
// formatting parameters necessary to format absolute numeric
// values for display in text number strings.
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
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
//       period (".") or Decimal Point.
//           United States Example: '123.45678'
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
//       FormatterAbsoluteValue.NewWithComponents().
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewDetailRunes(
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewDetailRunes()")

	newNumStrFmtSpecAbsValueDto :=
		FormatterAbsoluteValue{}

	newNumStrFmtSpecAbsValueDto.lock = new(sync.Mutex)

	err := formatterAbsoluteValueUtility{}.ptr().
		setDetail(
			&newNumStrFmtSpecAbsValueDto,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)

	return newNumStrFmtSpecAbsValueDto, err
}

// NewUnitedStatesDefaults - Creates and returns a new instance of
// FormatterSignedNumber. This method specifies the United States
// default values for signed number string formatting.
//
// United States Signed Number default formatting parameters are
// defined as follows:
//
//                 Absolute Value format: "127.54"
//           Decimal Separator Character: '.'
//         Thousands Separator Character: ','
//      Thousands Integer Digit Grouping: 3
//           Turn On Thousands Separator: true
//  United States Absolute Value Example: 2,354.92846
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
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
//  FormatterAbsoluteValue
//     - This parameter will return a new, populated instance of
//       FormatterAbsoluteValue configured United States default
//       absolute value number formatting parameters.
//
func (fmtAbsVal FormatterAbsoluteValue) NewUnitedStatesDefaults() FormatterAbsoluteValue {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	newFmtAbsoluteValue := FormatterAbsoluteValue{}

	newFmtAbsoluteValue.lock = new(sync.Mutex)

	_ = formatterAbsoluteValueUtility{}.ptr().
		setToUnitedStatesDefaults(
			&newFmtAbsoluteValue,
			nil)

	return newFmtAbsoluteValue
}

// NewWithComponents - Creates and returns a new instance of
// FormatterAbsoluteValue.
//
// The FormatterAbsoluteValue type encapsulates the
// configuration parameters necessary to format absolute numeric
// values for display in text number strings.
//
// This method requires detailed input parameters which provide
// granular control over all data fields contained in the returned
// new instance of FormatterAbsoluteValue.
//
// For a 'New' method using minimum input parameters coupled
// with default values, see:
//      FormatterAbsoluteValue.NewBasicRunes()
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//            Example: 1,000,000,000
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: 1000000000
//
//
//  numericSeparators             NumericSeparators
//     - This NumericSeparators object is used to specify the
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
//  FormatterAbsoluteValue
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of
//       FormatterAbsoluteValue.
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
func (fmtAbsVal FormatterAbsoluteValue) NewWithComponents(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) (
	FormatterAbsoluteValue,
	error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"NewWithComponents()")

	newFmtAbsoluteValue := FormatterAbsoluteValue{}

	newFmtAbsoluteValue.lock = new(sync.Mutex)

	err := formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			&newFmtAbsoluteValue,
			absoluteValFmt,
			turnOnIntegerDigitsSeparation,
			numericSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"Setting 'newFmtAbsoluteValue'"))

	return newFmtAbsoluteValue, err
}

// SetAbsoluteValueFormat - Sets the formatting string used to
// format absolute numeric values in text number strings.
//
// If the absolute value format string input parameter,
// 'absoluteValueFormatStr', is invalid, this method will return
// an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValueFormatStr     string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
func (fmtAbsVal *FormatterAbsoluteValue) SetAbsoluteValueFormat(
	absoluteValueFormatStr string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetAbsoluteValueFormat()")

	nStrAbsValDtoElectron :=
		formatterAbsoluteValueElectron{}

	_,
		err = nStrAbsValDtoElectron.testAbsoluteValueFormatStr(
		absoluteValueFormatStr,
		ePrefix.XCtx(
			"Testing validity of 'absoluteValueFormatStr'"))

	if err != nil {
		return err
	}

	fmtAbsVal.absoluteValFmt = absoluteValueFormatStr

	return err
}

// SetBasic - This method will set all of the member variable
// data values for the current instance of FormatterAbsoluteValue.
// The input parameters represent the minimum information required
// to configure the data values for a FormatterAbsoluteValue
// object. Default values are used to supplement these input
// parameters.
//
// This method differs from FormatterAbsoluteValue.SetBasicRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
//           Example: '123.45678'
//
//
//  thousandsSeparatorChars       string
//     - The character or characters which will be used to delimit
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
//       FormatterAbsoluteValue.SetWithComponents().
//
//
//  turnOnThousandsSeparator      bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnThousandsSeparator' is a boolean flag
//       used to control the 'Thousands Separator'. When set to
//       'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
//     - The character which will be used to delimit 'thousands' in
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
//       FormatterAbsoluteValue.NewWithComponents().
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
func (fmtAbsVal *FormatterAbsoluteValue) SetBasic(
	decimalSeparatorChars string,
	thousandsSeparatorChars string,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetBasic()")

	return formatterAbsoluteValueUtility{}.ptr().
		setBasic(
			fmtAbsVal,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)
}

// SetBasicRunes - This method will set all of the member variable
// data values for the current instance of FormatterAbsoluteValue.
// The input parameters represent the minimum information required
// to configure the data values for a FormatterAbsoluteValue
// object. Default values are used to supplement these input
// parameters.
//
// This method differs from FormatterAbsoluteValue.SetBasic() in
// that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'thousandsSeparatorChars'.
//
// To exercise granular control over all parameters needed to
// construct an instance of FormatterAbsoluteValue,
// reference method:
//   'FormatterAbsoluteValue.SetWithComponents()'
//
// This method automatically sets a default integer digits grouping
// sequence of '3'. This means that integers will be grouped by
// thousands.
//
//        Example: '1,000,000,000'
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
//           Example: '123.45678'
//
//
//  thousandsSeparatorChars       []rune
//     - The character or characters which will be used to delimit
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
//       FormatterAbsoluteValue.SetWithComponents().
//
//
//  turnOnThousandsSeparator      bool
//     - Inter digits separation is also known as the 'Thousands
//       Separator". Often a single character is used to separate
//       thousands within the integer component of a numeric value
//       in number strings. In the United States, the comma
//       character (',') is used to separate thousands.
//            Example: 1,000,000,000
//
//       The parameter 'turnOnThousandsSeparator' is a boolean flag
//       used to control the 'Thousands Separator'. When set to
//       'true', integer number strings will be separated into
//       thousands for text presentation.
//            Example: '1,000,000,000'
//
//       When this parameter is set to 'false', the 'Thousands
//       Separator' will NOT be inserted into text number strings.
//            Example: '1000000000'
//
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
//     - The character which will be used to delimit 'thousands' in
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
//       FormatterAbsoluteValue.NewWithComponents().
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
func (fmtAbsVal *FormatterAbsoluteValue) SetBasicRunes(
	decimalSeparatorChars []rune,
	thousandsSeparatorChars []rune,
	turnOnThousandsSeparator bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetBasicRunes()")

	return formatterAbsoluteValueUtility{}.ptr().
		setBasicRunes(
			fmtAbsVal,
			decimalSeparatorChars,
			thousandsSeparatorChars,
			turnOnThousandsSeparator,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)
}

// SetDetail - This method will set all of the member variable
// data values for the current instance of FormatterAbsoluteValue.
// New data values will be generated from the input parameters
// described below.
//
// This method differs from FormatterAbsoluteValue.SetDetailRunes()
// in that this method accepts strings for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// The FormatterAbsoluteValue type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
//       FormatterAbsoluteValue.SetWithComponents().
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
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
func (fmtAbsVal *FormatterAbsoluteValue) SetDetail(
	decimalSeparatorChars string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetDetail()")

	return formatterAbsoluteValueUtility{}.ptr().
		setDetail(
			fmtAbsVal,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)
}

// SetDetailRunes - This method will set all of the member variable
// data values for the current instance of FormatterAbsoluteValue.
// New data values will be generated from the input parameters
// described below.
//
// This method differs from FormatterAbsoluteValue.SetDetail()
// in that this method accepts rune arrays for input parameters,
// 'decimalSeparatorChars' and 'integerDigitsSeparators'.
//
// The FormatterAbsoluteValue type encapsulates the formatting
// parameters necessary to format absolute numeric values for
// display in text number strings.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
//       period (".") or Decimal Point.
//           United States Example: '123.45678'
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
//       FormatterAbsoluteValue.SetWithComponents().
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
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
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
func (fmtAbsVal *FormatterAbsoluteValue) SetDetailRunes(
	decimalSeparatorChars []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	turnOnIntegerDigitsSeparation bool,
	absoluteValFmt string,
	requestedNumberFieldLen int,
	numberFieldTextJustify TextJustify,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetDetailRunes()")

	return formatterAbsoluteValueUtility{}.ptr().
		setDetailRunes(
			fmtAbsVal,
			decimalSeparatorChars,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			turnOnIntegerDigitsSeparation,
			absoluteValFmt,
			requestedNumberFieldLen,
			numberFieldTextJustify,
			ePrefix)
}

// SetNumberFieldDto - Sets the Number Field Length Dto object
// for the current FormatterAbsoluteValue instance.
//
// The Number Field Length Dto object is used to specify the length
// and string justification characteristics used to display
// absolute value number strings within a number field.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldLenDto   NumberFieldDto
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
func (fmtAbsVal *FormatterAbsoluteValue) SetNumberFieldDto(
	numberFieldLenDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetNumberFieldDto()")

	return fmtAbsVal.numFieldDto.CopyIn(
		&numberFieldLenDto,
		ePrefix)
}

// SetNumericSeparators - Sets the Numeric Separators object for
// the current FormatterAbsoluteValue instance.
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
//     - This instance of 'NumericSeparators' is used to specify
//       the separator characters which will be included in number
//       string text displays.
//
//        type NumericSeparators struct {
//         decimalSeparators              rune
//         integerDigitsSeparator        rune
//         integerDigitsGroupingSequence []uint
//        }
//
//        decimalSeparators rune
//
//        The 'Decimal Separator' is used to separate integer and
//        fractional digits within a floating point number display.
//
//        integerDigitsSeparator rune
//
//        This type also encapsulates the integer digits separator, often
//        referred to as the 'Thousands Separator'. This is used to
//        separate thousands digits within the integer component of a
//        number string.
//
//        integerDigitsGroupingSequence []uint
//
//        Related to the integer digits separator, the integer digits
//        grouping sequence is also encapsulated in this type. The integer
//        digits grouping sequence is used to identify the digits which
//        will be grouped and separated by the integer digits separator.
//
//        In most western countries integer digits to the left of the
//        decimal separator (a.k.a. decimal point) are separated into
//        groups of three digits representing a grouping of 'thousands'
//        like this: '1,000,000,000,000'. In this case the parameter
//        integer digits grouping sequence would be configured as:
//                     integerDigitsGroupingSequence = []uint{3}
//
//        In some countries and cultures other integer groupings are used.
//        In India, for example, a number might be formatted as like this:
//                      '6,78,90,00,00,00,00,000'
//        The right most group has three digits and all the others are
//        grouped by two. In this case integer digits grouping sequence
//        would be configured as:
//                     integerDigitsGroupingSequence = []uint{3,2}
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
//  --- NONE ---
//
func (fmtAbsVal *FormatterAbsoluteValue) SetNumericSeparators(
	numberSeparators NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetNumericSeparators()")

	return fmtAbsVal.numericSeparators.CopyIn(
		&numberSeparators,
		ePrefix)
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterAbsoluteValue object. For Absolute Value
// formatters the Number String Format Type Code is automatically
// set to:
//    NumStrFormatTypeCode(0).AbsoluteValue().
//
// This method is required by interface INumStrFormatter.
//
func (fmtAbsVal *FormatterAbsoluteValue) SetNumStrFormatTypeCode() {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	fmtAbsVal.numStrFmtType =
		NumStrFormatTypeCode(0).AbsoluteValue()
}

// SetToUnitedStatesDefaults - Sets the member variable data
// values for the current FormatterAbsoluteValue instance
// to United States Default values.
//
// In the United States, Absolute Value default formatting
// parameters are defined as follows:
//
//    Absolute Value Number format: "127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetToUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetToUnitedStatesDefaults()")

	err = formatterAbsoluteValueUtility{}.ptr().
		setToUnitedStatesDefaults(
			fmtAbsVal,
			ePrefix.XCtx("fmtAbsVal"))

	return err
}

// SetToUnitedStatesDefaultsIfEmpty - If the current
// FormatterAbsoluteValue instance is empty or invalid,
// this method will set the member variable data values to United
// States default values.
//
// If the current FormatterAbsoluteValue instance is valid
// and NOT empty, this method will take no action and exit.
//
// In the United States, Absolute Value default formatting
// parameters are defined as follows:
//
//    Absolute Value Number format: "127.54"
//     Decimal Separator Character: '.'
//   Thousands Separator Character: ','
//     Turn On Thousands Separator: true
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetToUnitedStatesDefaultsIfEmpty(
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetToUnitedStatesDefaultsIfEmpty()")

	isValid,
		_ := formatterAbsoluteValueMolecule{}.ptr().testValidityOfFormatterAbsoluteValue(
		fmtAbsVal,
		ErrPrefixDto{}.Ptr())

	absValUtil := formatterAbsoluteValueUtility{}

	if isValid {
		return err
	}

	err = absValUtil.setToUnitedStatesDefaults(
		fmtAbsVal,
		ePrefix.XCtx("fmtAbsVal"))

	return err
}

// SetTurnOnIntegerDigitsSeparationFlag - Sets the
// 'turnOnIntegerDigitsSeparation' flag for the current instance of
// FormatterAbsoluteValue.
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
func (fmtAbsVal *FormatterAbsoluteValue) SetTurnOnIntegerDigitsSeparationFlag(
	turnOnIntegerDigitsSeparation bool) {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	fmtAbsVal.turnOnIntegerDigitsSeparation =
		turnOnIntegerDigitsSeparation
}

// SetWithComponents - This method will set all of the member
// variable data values for the current instance of
// FormatterAbsoluteValue.
//
// The FormatterAbsoluteValue type encapsulates the
// configuration parameters necessary to format absolute numeric
// values for display in text number strings.
//
// The member variable 'FormatterAbsoluteValue.numStrFmtType' is
// defaulted to:
//         NumStrFormatTypeCode(0).AbsoluteValue()
//
// IMPORTANT
// This method will overwrite all pre-existing data values in the
// current FormatterAbsoluteValue instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  absoluteValFmt                string
//     - A string specifying the number string format to be used in
//       formatting absolute numeric values in text number strings.
//
//       Absolute Value Formatting Terminology and Placeholders:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the actual numeric value of
//                     a number string. This place holder signals
//                     that the actual length of the numeric value
//                     including formatting characters and symbols
//                     such as Thousands Separators and Decimal
//                     Separators.
//
//               "+" - The Plus Sign ('+'). If present in the format
//                     string, the plus sign ('+') specifies  where
//                     the plus sign will be placed in relation to
//                     the positive numeric value.
//
//       Absence of "+" - The absence of a plus sign ('+') means that
//                        the positive numeric value will be displayed
//                        in text with out a plus sign ('+'). This is
//                        the default for absolute value number formatting.
//
//       Valid format strings for absolute value number strings
//       are listed as follows:
//
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Absolute Value Format
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
//  numericSeparators             NumericSeparators
//     - This instance of 'NumericSeparators' is used to specify
//       the separator characters which will be included in number
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
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note that this error message will incorporate the
//       method chain and text passed by input parameter, 'ePrefix'.
//       The 'ePrefix' text will be prefixed to the beginning of the
//       error message.
//
func (fmtAbsVal *FormatterAbsoluteValue) SetWithComponents(
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	numericSeparators NumericSeparators,
	numFieldDto NumberFieldDto,
	ePrefix *ErrPrefixDto) error {

	if fmtAbsVal.lock == nil {
		fmtAbsVal.lock = new(sync.Mutex)
	}

	fmtAbsVal.lock.Lock()

	defer fmtAbsVal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"FormatterAbsoluteValue." +
			"SetWithComponents()")

	return formatterAbsoluteValueMechanics{}.ptr().
		setAbsValDtoWithComponents(
			fmtAbsVal,
			absoluteValFmt,
			turnOnIntegerDigitsSeparation,
			numericSeparators,
			numFieldDto,
			ePrefix.XCtx(
				"Setting 'fmtAbsVal'"))
}
