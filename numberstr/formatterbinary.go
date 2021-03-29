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
//     - This instance of 'NumericSeparators' is
//       used to specify the separator characters which will be
//       including in the number string text display.
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
// If input parameter 'incomingFormatterCurrency' is judged to be
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
//  incomingFormatterCurrency     *FormatterCurrency
//     - A pointer to an instance of FormatterBinary.
//       The data values in this object will be copied to the
//       current FormatterBinary instance.
//
//       If input parameter 'incomingFormatterCurrency' is judged to
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
