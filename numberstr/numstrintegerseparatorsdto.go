package numberstr

import (
	"fmt"
	"sync"
)

// NumStrIntSeparatorsDto - The NumStrIntSeparatorsDto type manages
// an internal collection or array of NumStrIntSeparator objects.
// These NumStrIntSeparator objects, taken as a whole, provide
// formatting specifications for complex integer group separation
// operations.
//
// These NumStrIntSeparator objects are used to control the
// grouping and separation characters employed in separating
// integer digits within a number string.
//
//  United States Example:        1,000,000,000
//  European Example:             1 000 000 000
//  Indian Number System Example: 6,78,90,00,00,00,00,000
//  Chinese Numeral Example:      6789,0000,0000,0000
//
type NumStrIntSeparatorsDto struct {
	intSeparators []NumStrIntSeparator
	lock          *sync.Mutex
}

// Add - Adds one NumStrIntSeparator object to the existing
// collection of NumStrIntSeparator objects encapsulated by the
// current NumStrIntSeparatorsDto instance.
//
// The new NumStrIntSeparator object is created from the following
// input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          []rune
//     - An array of characters used to delimit or separate integer
//       digits within a number string. The most common example is
//       the thousands separator. In the United States, this is a
//       a single comma character ','. United States Example:
//               1,000,000,000
//
//       In other countries and cultures, multiple characters are
//       used to separate integer digits in a number string.
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
//  restartIntGroupingSequence bool
//     - Used in the last NumStrIntSeparatorsDto array element to
//       signal whether the entire array sequence will be restarted
//       from index array element zero.
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
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. An error value of 'nil'
//       signals that the new NumStrIntSeparator object was
//       successfully added to the collection managed by the
//       current NumStrIntSeparatorsDto instance.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) Add(
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool,
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto.Add()")

	if intSeparatorChars == nil {
		intSeparatorChars =
			make([]rune, 0, 5)
	}

	lIntSepChars := len(intSeparatorChars)

	if lIntSepChars == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a zero length rune array.\n",
			ePrefix)
	}

	newIntSep := NumStrIntSeparator{}

	newIntSep.intSeparatorChars =
		make([]rune, lIntSepChars, lIntSepChars+5)

	for i := 0; i < lIntSepChars; i++ {
		newIntSep.intSeparatorChars[i] =
			intSeparatorChars[i]
	}

	newIntSep.intSeparatorGrouping =
		intSeparatorGrouping

	newIntSep.intSeparatorRepetitions =
		intSeparatorRepetitions

	newIntSep.restartIntGroupingSequence =
		restartIntGroupingSequence

	newIntSep.lock = new(sync.Mutex)

	if intSeparatorsDto.intSeparators == nil {
		intSeparatorsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	intSeparatorsDto.intSeparators =
		append(intSeparatorsDto.intSeparators, newIntSep)

	return nil
}

// CopyIn - Copies the data fields from an incoming
// NumStrIntSeparatorsDto instance  to the data fields
// of the current instance of NumStrIntSeparatorsDto
// instance.
//
// If input parameter 'incomingIntSepsDto' is judged to be
// invalid, this method will return an error.
//
// IMPORTANT
//
// Be advised, all of the data fields in the current
// NumStrIntSeparatorsDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingIntSepsDto  *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto.
//       The data values in this object will be copied to the
//       current NumStrIntSeparatorsDto instance.
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) CopyIn(
	incomingIntSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto.CopyIn()")

	return numStrIntSeparatorsDtoElectron{}.ptr().
		copyIn(
			intSeparatorsDto,
			incomingIntSepsDto,
			ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrIntSeparatorsDto instance.
//
// If the current NumStrIntSeparatorsDto instance is judged
// to be invalid, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  NumStrIntSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparatorsDto will be created and returned
//       containing all of the data values copied from the current
//       instance of NumStrIntSeparatorsDto.
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) CopyOut(
	ePrefix *ErrPrefixDto) (
	NumStrIntSeparatorsDto,
	error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto.CopyOut()")

	return numStrIntSeparatorsDtoElectron{}.ptr().
		copyOut(
			intSeparatorsDto,
			ePrefix)
}

// Equal - Receives an incoming NumStrIntSeparatorsDto
// instance and compares it the current NumStrIntSeparatorsDto
// instance. If the two objects have equal data values, this method
// returns 'true'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto2         NumStrIntSeparatorsDto
//     - An instance of NumStrIntSeparatorsDto. The data values in
//       this object will be compared to those contained in the
//       current NumStrIntSeparatorsDto. If the data values are
//       equivalent this method will return 'true'
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the data values contained in input parameters
//       'intSepsDto2' and the current NumStrIntSeparatorsDto
//       instance are equivalent, this boolean return value will be
//       set to 'true'.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) Equal(
	intSepsDto2 NumStrIntSeparatorsDto) bool {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	var isEqual bool

	isEqual,
		_ = numStrIntSeparatorsDtoQuark{}.ptr().isEqual(
		intSeparatorsDto,
		&intSepsDto2,
		nil)

	return isEqual
}

// GetIntSeparators - Returns the internal member variable:
//    NumStrIntSeparatorsDto.intSeparators
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
//
//  intSeparators       []NumStrIntSeparator
//     - Returns an array of type NumStrIntSeparator containing
//       specifications for integer separation formatting.
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
//             like this: '6,78,90,00,00,00,00,000'.
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) GetIntSeparators(
	ePrefix *ErrPrefixDto) (
	intSeparators []NumStrIntSeparator,
	err error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"GetIntSeparators()")

	if intSeparatorsDto.intSeparators == nil {
		intSeparatorsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSeps := len(intSeparatorsDto.intSeparators)

	if lenIntSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal member variable is invalid!\n"+
			"'intSeparatorsDto.intSeparators' is a zero length array.\n",
			ePrefix.String())

		return intSeparators, err
	}

	intSeparators =
		make([]NumStrIntSeparator, lenIntSeps, lenIntSeps+5)

	for i := 0; i < lenIntSeps; i++ {
		err = intSeparators[i].CopyIn(
			&intSeparatorsDto.intSeparators[i],
			ePrefix.XCtx(
				fmt.Sprintf(
					"intSeparatorsDto.intSeparators[%v]",
					i)))

		if err != nil {
			return intSeparators, err
		}
	}

	return intSeparators, err
}

// GetNumberOfArrayElements - Returns the number of array elements
// in the internal member variable array, 'intSeparators'.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) GetNumberOfArrayElements() int {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if intSeparatorsDto.intSeparators == nil {
		intSeparatorsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	return len(intSeparatorsDto.intSeparators)
}

// GetRestartIntGroupingSequence - Returns the internal member
// variable 'restartIntGroupingSequence' for the last element in
// the internal array of NumStrIntSeparator objects. If the
// internal array of NumStrIntSeparator objects is a zero length
// array, this method will return an error.
//
// The NumStrIntSeparatorsDto type manages an internal collection
// or array of NumStrIntSeparator objects. These NumStrIntSeparator
// objects, taken as a whole, provide formatting specifications for
// complex integer group separation operations.
//
// If the current NumStrIntSeparator is the last element in an
// array of NumStrIntSeparator objects, the 'Restart Integer
// Grouping Sequence' flag signals whether the integer separation
// operation will be restarted from the first NumStrIntSeparator
// object in the array.
//
// Again, the NumStrIntSeparator.restartIntGroupingSequence boolean
// flag only has meaning if the current NumStrIntSeparator object
// is last element in an array of NumStrIntSeparator objects.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) GetRestartIntGroupingSequence(
	ePrefix *ErrPrefixDto) (
	bool, error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"GetRestartIntGroupingSequence()")

	if intSeparatorsDto.intSeparators == nil {
		intSeparatorsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSeps := len(intSeparatorsDto.intSeparators)

	if lenIntSeps == 0 {
		return false,
			fmt.Errorf("%v\n"+
				"Error: Internal 'intSeparators' array is invalid!\n"+
				"'intSeparatorsDto.intSeparators' is a zero length array.\n",
				ePrefix.String())
	}

	return intSeparatorsDto.intSeparators[lenIntSeps-1].
		restartIntGroupingSequence, nil
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrIntSeparatorsDto instance to determine whether the current
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
//     - This returned boolean value will signal whether the
//       current NumStrIntSeparatorsDto is valid, or not. If the
//       current NumStrIntSeparatorsDto contains valid data, this
//       method returns 'true'. If the data is invalid, this method
//       will return 'false'.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) IsValidInstance() (
	isValid bool) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	isValid,
		_ = numStrIntSeparatorsDtoQuark{}.ptr().
		testValidityOfNumStrIntSepsDto(
			intSeparatorsDto,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrIntSeparatorsDto instance to determine whether the
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
//     - If the current instance of NumStrIntSeparatorsDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (intSeparatorsDto *NumStrIntSeparatorsDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"IsValidInstanceError()")

	_,
		err := numStrIntSeparatorsDtoQuark{}.ptr().
		testValidityOfNumStrIntSepsDto(
			intSeparatorsDto,
			ePrefix.XCtx(
				"intSeparatorsDto"))

	return err
}

// New - Returns a new instance of NumStrIntSeparatorsDto
// configured for United States default integer separation values.
//
// The integer separation values used in the United States are
// listed as follows:
//
//  Thousands Separator (a.k.a Integer Digits Separator) = ","
//  Integer Digits Grouping Sequence = 3
//  Example Floating Point Number String: 1,000,000,000.456
//
func (intSeparatorsDto NumStrIntSeparatorsDto) New() NumStrIntSeparatorsDto {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	newIntSepsDto := NumStrIntSeparatorsDto{}

	_ = numStrIntSeparatorsDtoMechanics{}.ptr().
		setToUSADefaults(
			&newIntSepsDto,
			nil)

	return newIntSepsDto
}

// NewBasic - Creates and returns a new instance of
// NumStrIntSeparatorsDto. This method is intended to configure a
// basic or simple NumStrIntSeparatorsDto instance. Data values for
// the new instance are generated from default values and a minimum
// number of input parameters.
//
// The NumStrIntSeparatorsDto type encapsulates the integer digit
// separators used in formatting a number string for text display.
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
// Again, this method will automatically set the 'integer digits
// grouping sequence' to a default of 3-digits.
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumStrIntSeparatorsDto.NewWithComponents()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  newIntSepsDto              NumStrIntSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparatorsDto will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  err                        error
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
func (intSeparatorsDto NumStrIntSeparatorsDto) NewBasic(
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) (
	newIntSepsDto NumStrIntSeparatorsDto,
	err error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"NewBasic()")

	err =
		numStrIntSeparatorsDtoUtility{}.ptr().
			setBasic(
				&newIntSepsDto,
				integerDigitsSeparators,
				ePrefix.XCtx("newIntSepsDto"))

	return newIntSepsDto, err
}

// NewBasicRunes - Creates and returns a new instance of
// NumStrIntSeparatorsDto. This method is intended to configure a
// basic or simple NumStrIntSeparatorsDto instance. Data values for
// the new instance are generated from default values and a minimum
// number of input parameters.
//
// The NumStrIntSeparatorsDto type encapsulates the integer digit
// separators used in formatting a number string for text display.
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
// Again, this method will automatically set the 'integer digits
// grouping sequence' to a default of 3-digits.
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumStrIntSeparatorsDto.NewWithComponents()
//
// This method is an alternative to method
// NumStrIntSeparatorsDto.NewBasic() in that this method
// accepts integer separator characters as an array of runes
// instead of a string.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  newIntSepsDto              NumStrIntSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparatorsDto will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  err                        error
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
func (intSeparatorsDto NumStrIntSeparatorsDto) NewBasicRunes(
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) (
	newIntSepsDto NumStrIntSeparatorsDto,
	err error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"NewBasicRunes()")

	err =
		numStrIntSeparatorsDtoUtility{}.ptr().
			setBasicRunes(
				&newIntSepsDto,
				integerDigitsSeparators,
				ePrefix.XCtx("newIntSepsDto"))

	return newIntSepsDto, err
}

// NewWithComponents - Creates and returns a new instance of
// NumStrIntSeparatorsDto. The new is generated from the array of
// NumStrIntSeparator objects passed through input parameter,
// 'integerSeparatorsDto'.
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
//  ePrefix                    *ErrPrefixDto
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
//  newIntSepsDto              NumStrIntSeparatorsDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrIntSeparatorsDto.
//
//
//  err                        error
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
func (intSeparatorsDto NumStrIntSeparatorsDto) NewWithComponents(
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	newIntSepsDto NumStrIntSeparatorsDto,
	err error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"NewWithComponents()")

	err =
		numStrIntSeparatorsDtoMechanics{}.ptr().
			setWithComponents(
				&newIntSepsDto,
				integerSeparators,
				ePrefix.XCtx("newIntSepsDto"))

	return newIntSepsDto, err
}

// SetBasic - Overwrites all the member variable data values for
// the current NumStrIntSeparatorsDto instance. This method is
// intended to configure a basic or simple NumStrIntSeparatorsDto
// object using a minimum number of input parameters and specified
// default values.
//
// The NumStrIntSeparatorsDto type encapsulates the integer digit
// separators used in formatting a number string for text display.
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
//     NumStrIntSeparatorsDto.SetWithComponents()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparatorsDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  newIntSepsDto              NumStrIntSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparatorsDto will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  err                        error
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) SetBasic(
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"SetBasic()")

	if len(integerDigitsSeparators) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is zero length string.\n",
			ePrefix.String())

	}

	return numStrIntSeparatorsDtoUtility{}.ptr().
		setBasic(
			intSeparatorsDto,
			integerDigitsSeparators,
			ePrefix.XCtx("intSeparatorsDto"))
}

// SetBasicRunes - Overwrites all the member variable data values for
// the current NumStrIntSeparatorsDto instance. This method is
// intended to configure a basic or simple NumStrIntSeparatorsDto
// object using a minimum number of input parameters and specified
// default values.
//
// The NumStrIntSeparatorsDto type encapsulates the integer digit
// separators used in formatting a number string for text display.
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
// This method applies the 3-digit integer digits grouping
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
//     NumStrIntSeparatorsDto.SetWithComponents()
//
// This method is an alternative to method
// numStrIntSeparatorsDtoUtility.SetBasic() in that this method
// accepts integer separator characters as an array of runes
// instead of a string.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparatorsDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  newIntSepsDto              NumStrIntSeparatorsDto
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparatorsDto will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  err                        error
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) SetBasicRunes(
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"SetBasicRunes()")

	if len(integerDigitsSeparators) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is zero length string.\n",
			ePrefix.String())

	}

	return numStrIntSeparatorsDtoUtility{}.ptr().
		setBasicRunes(
			intSeparatorsDto,
			integerDigitsSeparators,
			ePrefix.XCtx("intSeparatorsDto"))
}

// SetToUSADefaults - This method will overwrite and set the all
// the internal member variable data values to default values used
// in the United States. Integer separator values used in the
// United States consist of the comma character (','), an integer
// grouping of three ('3') and unlimited repetitions of this
// sequence.
//
//   United States Integer Separation Example:
//         '1,000,000,000,000'
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparatorsDto instance.
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) SetToUSADefaults(
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"SetToUSADefaults()")

	return numStrIntSeparatorsDtoMechanics{}.ptr().
		setToUSADefaults(
			intSeparatorsDto,
			ePrefix.XCtx(
				"intSeparatorsDto"))
}

// SetToUSADefaultsIfEmpty - If any of the NumStrIntSeparatorsDto
// data values are zero or invalid, this method will reset ALL data
// elements to United States default values.
//
// If the current NumStrIntSeparatorsDto instance is valid and
// populated with data, this method will take no action and exit.
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
// the current NumStrIntSeparatorsDto instance.
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) SetToUSADefaultsIfEmpty(
	ePrefix *ErrPrefixDto) error {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"SetToUSADefaultsIfEmpty()")

	return numStrIntSeparatorsDtoMechanics{}.ptr().
		setToUSADefaults(
			intSeparatorsDto,
			ePrefix.XCtx(
				"intSeparatorsDto"))
}

// SetWithComponents - This method will reset all of the member
// variable data values for the current instance of
// NumStrIntSeparatorsDto. The new data values will be generated
// from the input parameters listed below.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparatorsDto instance.
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
//  ePrefix                    *ErrPrefixDto
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
func (intSeparatorsDto *NumStrIntSeparatorsDto) SetWithComponents(
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumStrIntSeparatorsDto." +
			"SetWithComponents()")

	err =
		numStrIntSeparatorsDtoMechanics{}.ptr().
			setWithComponents(
				intSeparatorsDto,
				integerSeparators,
				ePrefix.XCtx("intSeparatorsDto"))

	return err
}
