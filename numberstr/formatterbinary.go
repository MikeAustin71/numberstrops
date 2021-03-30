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

// NewUnitedStatesDefaults - Creates and returns a new instance of
// FormatterBinary. This method specifies the United States
// default values for binary number string formatting.
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
// for the current FormatterSignedNumber instance to
// United States Default values.
//
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

// SetWithComponents - This method will overwrite and set all data
// data values for the current instance of FormatterBinary.
//
// The FormatterBinary type encapsulates the formatting parameters
// necessary to format numeric currency values for display in text
// number strings.
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
// current FormatterCurrency instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
