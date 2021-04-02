package numberstr

import (
	"fmt"
	"sync"
)

// NumericSeparators - This type encapsulates all the number
// separators required to format numeric values in text strings.
// These separators include the 'Decimal Separator' and the
// 'Integer Digits Separators'.
//
// The 'Integer Digits Separators' includes both the character or
// characters used to separate groups of integers and the grouping
// sequence. In Western Countries, integer grouping is most
// commonly known as 'thousands' grouping.
//      United States Example: 1,000,0000,000
//
//
// decimalSeparators    []rune
//
// The 'Decimal Separator' is used to separate integer and
// fractional digits within a floating point number display.
// The decimal separator may consist of one or more runes.
//
//
// integerSeparatorsDto    NumStrIntSeparatorsDto
//
// The NumStrIntSeparatorsDto type encapsulates the integer digits
// separators, often referred to as the 'Thousands Separator'.
// Integer digit separators are used to separate integers into
// specific groups within a number string. The
// NumStrIntSeparatorsDto manages an array or collection of
// NumStrIntSeparator objects. Taken as a whole, these
// NumStrIntSeparator objects define the integer separation
// operation used in formatting number strings.
//
//        type NumStrIntSeparatorsDto struct {
//          intSeparators []NumStrIntSeparator
//        }
//
//        type NumStrIntSeparator struct {
//         intSeparatorChars       []rune  // A series of runes used to separate integer digits.
//         intSeparatorGrouping    uint    // Number of integer digits in a group
//         intSeparatorRepetitions uint    // Number of times this character/group sequence is repeated
//                                         // A zero value signals unlimited repetitions.
//         restartIntGroupingSequence bool // If true, the grouping sequence starts over at index zero.
//        }
//
//
//    intSeparatorChars          []rune
//       - A series of runes or characters used to separate integer
//         digits in a number string. These characters are commonly
//         known as the 'thousands separator'. A 'thousands
//         separator' is used to separate groups of integer digits to
//         the left of the decimal separator (a.k.a. decimal point).
//         In the United States, the standard integer digits
//         separator is the single comma character (','). Other
//         countries and cultures use periods, spaces, apostrophes or
//         multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//    intSeparatorGrouping       uint
//       - This unsigned integer values specifies the number of
//         integer digits within a group. This value is used to group
//         integers within a number string.
//
//         In most western countries integer digits to the left of
//         the decimal separator (a.k.a. decimal point) are separated
//         into groups of three digits representing a grouping of
//         'thousands' like this: '1,000,000,000'. In this case the
//         intSeparatorGrouping value would be set to three ('3').
//
//         In some countries and cultures other integer groupings
//         are used. In India, for example, a number might be
//         formatted like this: '6,78,90,00,00,00,00,000'. Chinese
//         Numerals have an integer grouping value of four ('4').
//             Chinese Numerals Example: '12,3456,7890,2345'
//
//    intSeparatorRepetitions    uint
//       - This unsigned integer value specifies the number of times
//         this integer grouping is repeated. A value of zero signals
//         that this integer grouping will be repeated indefinitely.
//
//    restartIntGroupingSequence bool
//       - If the NumStrIntSeparator is the last element in an array
//         of NumStrIntSeparator objects, this boolean flag signals
//         whether the entire integer grouping sequence will be
//         restarted from array element zero.
//
type NumericSeparators struct {
	decimalSeparators    []rune
	integerSeparatorsDto NumStrIntSeparatorsDto
	lock                 *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumericSeparators instance to the data fields of the current
// NumericSeparators instance.
//
// If the incoming NumericSeparators instance is judged to be
// invalid, this method will return an error.
//
// This method will OVERWRITE the data fields of the current
// NumericSeparators instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNumericSeparators     *NumStrIntSeparator
//     - A pointer to an instance of NumStrIntSeparator.
//       The data values in this object will be copied to the
//       data fields of the current NumStrIntSeparator instance.
//
//       If input parameter 'incomingNumericSeparators' is judged
//       to be invalid, this method will return an error.
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
func (numSeps *NumericSeparators) CopyIn(
	incomingNumericSeparators *NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.CopyIn()")

	return numericSeparatorsElectron{}.ptr().
		copyIn(
			numSeps,
			incomingNumericSeparators,
			ePrefix)
}

// CopyOut - Returns a deep copy of the current NumericSeparators
// instance.
//
// If the current NumericSeparators instance is judged to be
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
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned
//       containing all of the data values copied from the current
//       instance of NumericSeparators.
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
func (numSeps *NumericSeparators) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.CopyOut()")

	return numericSeparatorsElectron{}.ptr().
		copyOut(
			numSeps,
			ePrefix.XCtx("numSeps->"))
}

// Empty - Deletes and resets the data values of all member
// variables within the current NumericSeparators instance to their
// initial 'zero' values.
//
func (numSeps *NumericSeparators) Empty() {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	_ = numericSeparatorsQuark{}.ptr().empty(
		numSeps,
		nil)

	numSeps.lock.Unlock()

	numSeps.lock = nil
}

// Equal - Receives an incoming NumericSeparators
// instance and compares it the current NumericSeparators
// instance. If the two objects have equal data values, this method
// returns 'true'
//

// Equal - Receives an NumStrIntSeparatorsDto object and proceeds
// to determine whether all data elements in this object are equal
// to all corresponding data elements in the current instance of
// NumStrIntSeparatorsDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDtoTwo       NumStrIntSeparatorsDto
//     - This method will compare all data elements in the current
//       NumStrIntSeparatorsDto instance to corresponding data
//       elements in this second NumStrIntSeparatorsDto object in
//       order determine equivalency.
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
//  bool
//     - If all the data elements in the current
//       NumStrIntSeparatorsDto instance are equal to all the
//       corresponding data elements in 'intSepsDtoTwo', this
//       return parameter will be set to 'true'. If all the data
//       elements are NOT equal, this return parameter will be set
//       to 'false'.
//
//
//  error
//     - If all the data elements in the current
//       NumStrIntSeparatorsDto are equal to all the corresponding
//       data elements in 'intSepsDtoTwo', this return parameter
//       will be set to 'nil'.
//
//       If the corresponding data elements are not equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (numSeps *NumericSeparators) Equal(
	numSepsTwo NumericSeparators,
	ePrefix *ErrPrefixDto) (
	bool,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.Equal()")

	return numericSeparatorsQuark{}.
		ptr().equal(
		numSeps,
		&numSepsTwo,
		ePrefix.XCtx(
			"numSeps vs numSepsTwo"))
}

// GetDecimalSeparators - Returns the decimal separator characters.
//
// The decimal separator separates integer and fractional digits
// in a floating point number string. In the United States, the
// decimal separator is the period ('.') or decimal point.
//
//             Example: 1234.456
//
func (numSeps *NumericSeparators) GetDecimalSeparators() []rune {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if numSeps.decimalSeparators == nil {
		numSeps.decimalSeparators =
			make([]rune, 0, 2)
	}

	lenNumSepDecSeps := len(numSeps.decimalSeparators)

	newDecSeps := make(
		[]rune,
		lenNumSepDecSeps,
		lenNumSepDecSeps+5)

	for i := 0; i < lenNumSepDecSeps; i++ {
		newDecSeps[i] = numSeps.decimalSeparators[i]
	}

	return newDecSeps
}

// GetIntegerDigitSeparators - Returns the integer digit separators
// as an NumStrIntSeparatorsDto object.
//
// The returned NumStrIntSeparatorsDto object encapsulates an array
// of type NumStrIntSeparator which contains integer separator
// characters, grouping sequences, sequence repetitions and
// the 'restartIntGroupingSequence' flag.
//
// The integer digit separators is also known as the 'thousands'
// separator. In the United States the standard integer digit
// separator character is the comma (',') and integers are shown
// in groups of three ('3').
//
//    United States Example: 1,000,000,000,000
//
// If the current NumericSeparators instance is judged to be
// invalid, this method will return an error.
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
//     - The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       the integer separation operation in number strings.
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
func (numSeps *NumericSeparators) GetIntegerDigitSeparators(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.GetIntegerDigitSeparators() ")

	return numSeps.integerSeparatorsDto.CopyOut(
		ePrefix.XCtx(
			"numSeps.integerSeparatorsDto"))
}

// IsValidInstance - Performs a diagnostic review of the current
// NumericSeparators instance to determine whether
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
//     - If this method completes successfully, the returned boolean
//       value will signal whether the current NumericSeparators
//       is valid, or not. If the current NumericSeparators
//       contains valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
func (numSeps *NumericSeparators) IsValidInstance() (
	isValid bool) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	nStrFmtSpecDigitsSepsQuark :=
		numericSeparatorsQuark{}

	isValid,
		_ = nStrFmtSpecDigitsSepsQuark.testValidityOfNumericSeparators(
		numSeps,
		new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumericSeparators instance to determine
// whether the current instance is valid in all respects.
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
//       If this instance of NumericSeparators contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (numSeps *NumericSeparators) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPrefCtx(
		"NumericSeparators.IsValidInstanceError()",
		"Testing Validity of 'numSeps'")

	numSepsQuark := numericSeparatorsQuark{}

	_,
		err := numSepsQuark.testValidityOfNumericSeparators(
		numSeps,
		ePrefix)

	return err
}

// NewBasic - Creates and returns a new instance of
// NumericSeparators. The returned NumericSeparators instance
// represents a basic or simple numeric separators object using
// default values and a minimum number of input parameters.
//
// Input parameter 'integerDigitsSeparators' specifies  the integer
// separator character or characters. The integer digit grouping is
// defaulted to a value of three (3). The 'separator repetitions'
// value is defaulted to zero (0), signaling unlimited repetitions.
// Finally, the 'restartIntGroupingSequence' flag will be defaulted
// to 'false'.
//
//This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// This method defaults the integer digits grouping sequence to
// that most commonly used across the world. Namely, this is a
// constant thousands grouping of three '3' digits.
//      Example: 1,000,000,000,000
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// Again, this method applies the 3-digit integer digits grouping
// sequence by default.
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumericSeparators.NewWithComponents()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          string
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    string
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
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
func (numSeps NumericSeparators) NewBasic(
	decimalSeparators string,
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.NewBasic()")

	newNumericSeps := NumericSeparators{}

	err := numericSeparatorsUtility{}.ptr().
		setBasic(
			&newNumericSeps,
			decimalSeparators,
			integerDigitsSeparators,
			ePrefix)

	return newNumericSeps, err
}

// NewBasicRunes - Creates and returns a new instance of
// NumericSeparators. The returned NumericSeparators instance
// represents a basic or simple numeric separators object using
// default values and a minimum number of input parameters.
//
// This method is an alternative to NumericSeparators.NewBasic()
// in that this method accepts decimal separators and integer digit
// separators as rune arrays.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// This method defaults the integer digits grouping sequence to
// that most commonly used across the world. Namely, this is a
// constant thousands grouping of three '3' digits.
//      Example: 1,000,000,000,000
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// Again, this method applies the 3-digit integer digits grouping
// sequence by default.
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumericSeparators.NewWithComponents()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator (a.k.a. decimal
//       point). In the United States, the standard integer digits
//       separator is the comma (",").
//
//             United States Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
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
func (numSeps NumericSeparators) NewBasicRunes(
	decimalSeparators []rune,
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.NewBasic()")

	newNumericSeps := NumericSeparators{}

	err := numericSeparatorsUtility{}.ptr().
		setBasicRunes(
			&newNumericSeps,
			decimalSeparators,
			integerDigitsSeparators,
			ePrefix)

	return newNumericSeps, err
}

// NewDetail - Creates and returns a new a full featured instance
// of NumericSeparators based on the input parameters listed below.
// Like method NumericSeparators.NewWithComponents(), this method
// provides detailed control over all aspects of NumericSeparators
// configuration.
//
// This method differs from NumericSeparators.NewDetailRunes() in
// that this method accepts a string for input parameter,
// 'decimalSeparators'.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          string
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerSeparatorsDto       NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned.
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
func (numSeps NumericSeparators) NewDetail(
	decimalSeparators string,
	integerSeparatorsDto NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.NewDetail()")

	newNumericSeps := NumericSeparators{}

	newNumericSeps.lock = new(sync.Mutex)

	err := numericSeparatorsMechanics{}.ptr().
		setDetail(
			&newNumericSeps,
			decimalSeparators,
			&integerSeparatorsDto,
			ePrefix)

	return newNumericSeps, err
}

// NewDetailRunes - Creates and returns a new a full featured instance
// of NumericSeparators based on the input parameters listed below.
// Like method NumericSeparators.NewWithComponents(), this method
// provides detailed control over all aspects of NumericSeparators
// configuration.
//
// This method differs from NumericSeparators.NewDetail() in that
// this method accepts a rune array for input parameter,
// 'decimalSeparators'.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerSeparators          NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned.
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
func (numSeps NumericSeparators) NewDetailRunes(
	decimalSeparators []rune,
	integerSeparators NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.NewDetailRunes()")

	newNumericSeps := NumericSeparators{}

	newNumericSeps.lock = new(sync.Mutex)

	err := numericSeparatorsMechanics{}.ptr().
		setDetailRunes(
			&newNumericSeps,
			decimalSeparators,
			&integerSeparators,
			ePrefix)

	return newNumericSeps, err
}

// NewIntSepsDetail - Creates and returns a new a full featured
// instance of NumericSeparators based on the input parameters
// listed below.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// This method differs from NumericSeparators.NewIntSepsDetailRunes()
// in that this method accepts strings for input parameters
// 'decimalSeparators' and 'integerDigitsSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          string
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    string
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
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned.
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
func (numSeps NumericSeparators) NewIntSepsDetail(
	decimalSeparators string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators." +
			"NewIntSepsDetail()")

	newNumericSeps := NumericSeparators{}

	newNumericSeps.lock = new(sync.Mutex)

	err := numericSeparatorsUtility{}.ptr().
		setIntSepsDetail(
			&newNumericSeps,
			decimalSeparators,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			ePrefix)

	return newNumericSeps, err
}

// NewIntSepsDetailRunes - Creates and returns a new a full featured
// instance of NumericSeparators based on the input parameters
// listed below.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// This method differs from NumericSeparators.NewIntSepsDetail()
// in that this method accepts rune arrays for input parameters
// 'decimalSeparators' and 'integerDigitsSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character ('.') otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (',').
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
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
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - If this method completes successfully, a new instance of
//       NumericSeparators will be created and returned.
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
func (numSeps NumericSeparators) NewIntSepsDetailRunes(
	decimalSeparators []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators." +
			"NewIntSepsDetail()")

	newNumericSeps := NumericSeparators{}

	newNumericSeps.lock = new(sync.Mutex)

	err := numericSeparatorsUtility{}.ptr().
		setIntSepsDetailRunes(
			&newNumericSeps,
			decimalSeparators,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			ePrefix)

	return newNumericSeps, err
}

// NewUnitedStatesDefaults - Returns a new instance of
// NumericSeparators.
//
// The values are automatically set to United States default
// numeric separator values. These values are listed as follows:
//
//  Decimal Separator Character = '.'
//  Thousands Separator (a.k.a Integer Digits Separator) = ','
//  Integer Digits (a.k.a. Thousands) Grouping Sequence = 3
//
//  United States Example Floating Point Number String:
//           1,000,000,000.456
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumericSeparators
//     - This parameter will return a new, populated instance of
//       NumericSeparators.
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
func (numSeps NumericSeparators) NewUnitedStatesDefaults(
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators." +
			"NewUnitedStatesDefaults()")

	newNumSep := NumericSeparators{}

	newNumSep.lock = new(sync.Mutex)

	err :=
		numericSeparatorsUtility{}.ptr().
			setWithUSADefaults(
				&newNumSep,
				ePrefix.XCtx(
					"newNumSep"))

	return newNumSep, err
}

// NewWithComponents - Creates and returns a new instance of NumericSeparators.
// This type encapsulates the digit separators used in formatting a
// number string for text display.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators             []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period ('.') character otherwise known as a decimal point.
//
//
//  integerSeparators             []NumStrIntSeparator
//     - Encapsulates an array of type NumStrIntSeparator
//       containing specifications for integer separation
//       formatting.
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
//             used. In India, for example, a number might be formatted
//             like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//             have an integer grouping value of four ('4').
//               Chinese Numerals Example: '12,3456,7890,2345'
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
//  NumericSeparators
//     - If this method completes successfully, new instance of
//       NumericSeparators will be created and
//       returned.
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
func (numSeps NumericSeparators) NewWithComponents(
	decimalSeparators []rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.NewWithComponents()")

	newDigitsSepsDto := NumericSeparators{}

	numericSeparatorDtoMech :=
		numericSeparatorsMechanics{}

	err := numericSeparatorDtoMech.setWithComponents(
		&newDigitsSepsDto,
		decimalSeparators,
		integerSeparators,
		ePrefix)

	return newDigitsSepsDto, err
}

// SetBasic - Overwrites all the member variable data values for
// the current NumericSeparators instance with numeric separator
// values generated from a minimum number of input parameters and
// specified default values.
//
// The resulting NumericSeparators values represent a basic numeric
// separators object with a three digit integer grouping sequence.
// This means that integer digits will be separated into
// 'thousands' with each group containing three digits each.
//
//     Example: 1,000,000,000
//
// Users have the option of specifying integer separator and
// decimal separator characters.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// This method defaults the integer digits grouping sequence to
// that most commonly used across the world. Namely, this is the
// the thousands grouping of three '3' digits.
//      Example: 1,000,000,000,000
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// Again, this method applies the 3-digit integer digits grouping
// sequence by default.
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumericSeparators.SetWithComponents()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          string
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    string
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetBasic(
	decimalSeparators string,
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.SetBasic()")

	return numericSeparatorsUtility{}.ptr().
		setBasic(
			numSeps,
			decimalSeparators,
			integerDigitsSeparators,
			ePrefix.XCtx("numSeps"))
}

// SetBasicRunes - Overwrites all the member variable data values for
// the current NumericSeparators instance with numeric separator
// values generated from a minimum number of input parameters and
// specified default values.
//
// The resulting NumericSeparators values represent a basic numeric
// separators object with a three digit integer grouping sequence.
// This means that integer digits will be separated into
// 'thousands' with each group containing three digits each.
//
//     Example: 1,000,000,000
//
// Users have the option of specifying integer separator and
// decimal separator characters.
//
// This method is an alternative to NumericSeparators.SetBasic()
// in that this method accepts decimal separators and integer digit
// separators as rune arrays.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// This method defaults the integer digits grouping sequence to
// that most commonly used across the world. Namely, this is the
// the thousands grouping of three '3' digits.
//      Example: 1,000,000,000,000
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// Again, this method applies the 3-digit integer digits grouping
// sequence by default.
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumericSeparators.SetWithComponents()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    []rune
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetBasicRunes(
	decimalSeparators []rune,
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.SetBasicRunes()")

	return numericSeparatorsUtility{}.ptr().
		setBasicRunes(
			numSeps,
			decimalSeparators,
			integerDigitsSeparators,
			ePrefix.XCtx("numSeps"))
}

// SetDecimalSeparators - Sets the value of the decimal separators
// for the current NumericSeparators instance.
//
// The decimal separator separates integer and fractional digits
// in a floating point number string. In the United States, the
// decimal separator is the period (".") or decimal point.
//
//             Example: 1234.456
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators   string
//     - Contains the character or characters which will be used to
//       separate integer and fractional digits in a number string.
//       The decimal separators value for the current
//       NumericSeparators instance will be overwritten and
//       replaced by the characters contained in this string.
//
//       If a zero length string is is passed for this parameter,
//       an error will be returned.
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
func (numSeps *NumericSeparators) SetDecimalSeparators(
	decimalSeparators string,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators." +
			"SetDecimalSeparators()")

	if len(decimalSeparators) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparators' is invalid!\n"+
			"'decimalSeparators' is a zero length string.\n",
			ePrefix.String())
	}

	decSepRunes := []rune(decimalSeparators)

	lenDecSepRunes := len(decSepRunes)

	if lenDecSepRunes == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparators' is invalid!\n"+
			"Converting 'decimalSeparators' to a rune array produced\n"+
			"a zero length rune array.\n",
			ePrefix.String())
	}

	numSeps.decimalSeparators =
		make([]rune, lenDecSepRunes, lenDecSepRunes+5)

	for i := 0; i < lenDecSepRunes; i++ {
		numSeps.decimalSeparators[i] =
			decSepRunes[i]
	}

	return nil
}

// SetDetail - Overwrites and resets all the data values in the
// current NumericSeparators instance. The new data values are
// generated from the input parameters listed below.
//
// This method differs from NumericSeparators.SetDetailRunes() in
// that this method accepts a string for input parameter,
// 'decimalSeparators'.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          string
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerSeparatorsDto       NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetDetail(
	decimalSeparators string,
	integerSeparatorsDto NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.SetDetail()")

	return numericSeparatorsMechanics{}.ptr().
		setDetail(
			numSeps,
			decimalSeparators,
			&integerSeparatorsDto,
			ePrefix)
}

// SetDetailRunes - Overwrites and resets all the data values in
// the current NumericSeparators instance. The new data values are
// generated from the input parameters listed below.
//
// This method differs from NumericSeparators.SetDetail() in that
// this method accepts an array of runes for input parameter,
// 'decimalSeparators'.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerSeparatorsDto       NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetDetailRunes(
	decimalSeparators []rune,
	integerSeparatorsDto NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.SetDetailRunes()")

	return numericSeparatorsMechanics{}.ptr().
		setDetailRunes(
			numSeps,
			decimalSeparators,
			&integerSeparatorsDto,
			ePrefix)
}

// SetIntegerSeparatorsArray - Sets the value of the internal
// integer digit separators array for this NumericSeparators
// instance. The integer digit separators array controls the
// formatting of integer digits in a number string.
//
// This method is very similar to method
// NumericSeparators.SetIntegerSeparatorsDto() which takes an input
// parameter of NumStrIntSeparatorsDto. NumStrIntSeparatorsDto
// objects encapsulate the integer digit separators array.
//
// In the United States, the standard integer digits separator is
// the comma (','). Other countries and cultures use spaces (' '),
// apostrophes ('), or periods ('.') to separate integer digits.
//
//             United States Example:  1,000,000,000
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// The integer digit separators array provides full support for all
// known types of integer grouping sequences.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerSeparators          []NumStrIntSeparator
//     - Encapsulates an array of type NumStrIntSeparator
//       containing specifications for integer separation
//       formatting.
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
func (numSeps *NumericSeparators) SetIntegerSeparatorsArray(
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"" +
			"NumericSeparators." +
			"SetIntegerSeparatorsArray()")

	if integerSeparators == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerSeparatorsDto' is invalid!\n"+
			"'integerSeparatorsDto' is 'nil'.",
			ePrefix.String())

		return err
	}

	lenIntSeparators := len(integerSeparators)

	if lenIntSeparators == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerSeparatorsDto' is invalid!\n"+
			"'integerSeparatorsDto' is a zero length array.",
			ePrefix.String())

		return err
	}

	err = numSeps.integerSeparatorsDto.SetWithComponents(
		integerSeparators,
		ePrefix)

	return
}

// SetIntegerSeparatorsDto - Sets member data variable controlling
// integer digit separation. This method will delete and replace
// the integer separation specification for the current
// NumericSeparators instance
//
// In the United States, the standard integer digits separator is
// the comma (','). Other countries and cultures use spaces,
// apostrophes, or periods ('.') to separate integer digits.
//
//             United States Example:  1,000,000,000
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// This method receives an input parameter of type
// NumStrIntSeparatorsDto which provides full support for all known
// types of integer grouping sequences.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerSeparatorsDto       NumStrIntSeparatorsDto
//     - The NumStrIntSeparatorsDto type manages an internal
//       collection or array of NumStrIntSeparator objects.
//       Taken as a whole, these NumStrIntSeparator objects define
//       the integer separation operation used in formatting number strings.
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
//
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetIntegerSeparatorsDto(
	integerSeparatorsDto NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumericSeparators." +
			"SetIntegerSeparators()")

	err = numericSeparatorsNanobot{}.ptr().
		setIntegerSeparatorsDto(
			numSeps,
			&integerSeparatorsDto,
			ePrefix.XCtx(
				"numSeps"))

	return err
}

// SetIntSepsDetail - Overwrites and resets all the data values in
// the current NumericSeparators instance. The new data values are
// generated from the input parameters listed below.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// This method differs from NumericSeparators.SetIntSepsDetailRunes()
// in that this method accepts strings for input parameters
// 'decimalSeparators' and 'integerDigitsSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          string
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character (".") otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    string
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
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetIntSepsDetail(
	decimalSeparators string,
	integerDigitsSeparators string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.SetIntSepsDetail()")

	return numericSeparatorsUtility{}.ptr().
		setIntSepsDetail(
			numSeps,
			decimalSeparators,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			ePrefix)
}

// SetIntSepsDetailRunes - Overwrites and resets all the data values in
// the current NumericSeparators instance. The new data values are
// generated from the input parameters listed below.
//
// The NumericSeparators type encapsulates the numeric digit
// separators used in formatting a number string for text display.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator and decimal separator characters.
//
// Digit separators are commonly referred to as the decimal point
// and thousands separator characters. In addition, Digit
// Separators include the integer digits grouping sequence.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = 3
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// This method differs from NumericSeparators.SetIntSepsDetail()
// in that this method accepts rune arrays for input parameters
// 'decimalSeparators' and 'integerDigitsSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators          []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period character ('.') otherwise known as a decimal point.
//
//            Example: 12.345
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  integerDigitsSeparators    []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (',').
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
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
//  ePrefix                    *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (numSeps *NumericSeparators) SetIntSepsDetailRunes(
	decimalSeparators []rune,
	integerDigitsSeparators []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators." +
			"SetIntSepsDetailRunes()")

	return numericSeparatorsUtility{}.ptr().
		setIntSepsDetailRunes(
			numSeps,
			decimalSeparators,
			integerDigitsSeparators,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			ePrefix)
}

// SetToUSADefaultsIfEmpty - If any of the NumericSeparators
// data values are zero or invalid, this method will set ALL data
// elements to United States default values.
//
// If the current NumericSeparators instance is valid and populated
// with data, this method will take no action and exit.
//
// United States default numeric separators are listed as follows:
//
//  Decimal Separator = '.'
//  Thousands Separator (a.k.a Integer Digits Separator) = ','
//  Integer Digits Grouping Sequence = 3
//  Example Floating Point Number String: 1,000,000,000.456
//
// IMPORTANT
//
// This method MAY overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
func (numSeps *NumericSeparators) SetToUSADefaultsIfEmpty() {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	isValid,
		_ := numericSeparatorsQuark{}.
		ptr().testValidityOfNumericSeparators(
		numSeps,
		new(ErrPrefixDto))

	if isValid {
		return
	}

	_ =
		numericSeparatorsUtility{}.ptr().
			setWithUSADefaults(
				numSeps,
				nil)
}

// SetToUSADefaults - Sets all of the numeric separator
// member variables to default United States values.
//
// United States default numeric separators are listed as follows:
//
//  Decimal Separator = '.'
//  Thousands Separator (a.k.a Integer Digits Separator) = ','
//  Integer Digits Grouping Sequence = 3
//  Example Floating Point Number String: 1,000,000,000.456
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
func (numSeps *NumericSeparators) SetToUSADefaults() {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	_ =
		numericSeparatorsUtility{}.ptr().
			setWithUSADefaults(
				numSeps,
				nil)
}

// SetWithComponents - This method will set all of the member variable
// data values for the current instance of NumericSeparators.
//
// The NumericSeparators type encapsulates the digit separators
// used in formatting a number string for text display.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumericSeparators instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparators             []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period ('.') character otherwise known as a decimal point.
//
//
//  integerSeparators          []NumStrIntSeparator
//     - Encapsulates an array of type NumStrIntSeparator
//       containing specifications for integer separation
//       formatting.
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
func (numSeps *NumericSeparators) SetWithComponents(
	decimalSeparators []rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref(
		"NumericSeparators.SetWithComponents()")

	return numericSeparatorsMechanics{}.ptr().
		setWithComponents(
			numSeps,
			decimalSeparators,
			integerSeparators,
			ePrefix.XCtx(
				"Setting Data Values for current instance 'numSeps'"))
}

// String - Returns a string detailing numeric separator
//information.
//
func (numSeps *NumericSeparators) String() string {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	str := fmt.Sprintf("Decimal Separators: %v\n",
		string(numSeps.decimalSeparators))

	str += numSeps.integerSeparatorsDto.String()

	str += "\n"

	return str
}
