package numberstr

import (
	"fmt"
	"sync"
)

// NumericSeparators - This type encapsulates all the number
// separators required to format numeric values in text strings.
// These separators include the 'Decimal Separator', The 'Integer
// Digits Separators'.
//
// The 'Integer Digits Separators' includes both the character used
// to separate groups of integers and the grouping sequence. This
// operation is most commonly known as 'thousands' grouping.
//      United States Example: 1,000,0000,000
//
// decimalSeparator rune
//
// The 'Decimal Separator' is used to separate integer and
// fractional digits within a floating point number display.
//
// integerDigitsSeparator rune
//
// This type also encapsulates the integer digits separator, often
// referred to as the 'Thousands Separator'. This is used to
// separate thousands digits within the integer component of a
// number string.
//
// integerSeparators             []NumStrIntSeparator
//
// An array of NumStrIntSeparator elements used to specify the
// integer separation operation.
//
//    type NumStrIntSeparator struct {
//      intSeparatorChar     rune // Integer separator character
//      intSeparatorGrouping uint // Number of integers in a group
//      intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                   // A zero value signals unlimited repetitions.
//                                   // A zero value signals unlimited repetitions.
//    }
//
//    intSeparatorChar     rune
//    - This separator is commonly known as the 'thousands'
//      separator. It is used to separate groups of integer
//      digits to the left of the decimal separator (a.k.a.
//      decimal point). In the United States, the standard
//      integer digits separator is the comma (','). Other
//      countries use periods, spaces or apostrophes to
//      separate integers.
//        United States Example:  1,000,000,000
//         numSeps.intSeparators =
//           []NumStrIntSeparator{
//                {
//                intSeparatorChar:   ',',
//                intSeparatorGrouping: 3,
//                intSeparatorRepetitions: 0,
//                },
//             }
//
//    intSeparatorGrouping []uint
//    - In most western countries integer digits to the left
//      of the decimal separator (a.k.a. decimal point) are
//      separated into groups of three digits representing
//      a grouping of 'thousands' like this: '1,000,000,000'.
//      In this case the intSeparatorGrouping value would be
//      set to three ('3').
//
//  In some countries and cultures other integer groupings are
//  used. In India, for example, a number might be formatted
//  like this: '6,78,90,00,00,00,00,000'. The right most group
//  has three digits and all the others are grouped by two. In
//  this case 'integerSeparators' would be configured as
//  follows:
//  as:
//
//  numSeps.intSeparators =
//    []NumStrIntSeparator{
//         {
//         intSeparatorChar:   ',',
//         intSeparatorGrouping: 3,
//         intSeparatorRepetitions: 1,
//         },
//         {
//         intSeparatorChar:     ',',
//         intSeparatorGrouping: 2,
//         intSeparatorRepetitions: 0,
//         },
//      }
//
type NumericSeparators struct {
	decimalSeparator  rune
	integerSeparators []NumStrIntSeparator
	lock              *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumericSeparators instance to the data fields
// of the current instance of NumericSeparators
// instance.
//
// If the incoming NumericSeparators is judged to be
// invalid, this method will return an error.
//
// This method will OVERWRITE the data fields of the current
// NumericSeparators instance.
//
func (numSeps *NumericSeparators) CopyIn(
	incomingSpecDigitsSepDto *NumericSeparators,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("NumericSeparators.CopyIn()")

	nStrFmtSpecDigitsSepsElectron :=
		numericSeparatorsElectron{}

	return nStrFmtSpecDigitsSepsElectron.copyIn(
		numSeps,
		incomingSpecDigitsSepDto,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current
// NumericSeparators instance.
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
	}

	ePrefix.SetEPref("NumericSeparators.CopyOut()")

	nStrFmtSpecDigitsSepsElectron :=
		numericSeparatorsElectron{}

	return nStrFmtSpecDigitsSepsElectron.copyOut(
		numSeps,
		ePrefix.XCtx("numSeps->"))
}

// Equal - Receives an incoming NumericSeparators
// instance and compares it the current NumericSeparators
// instance. If the two objects have equal data values, this method
// returns 'true'
//
func (numSeps *NumericSeparators) Equal(
	numSep2 NumericSeparators) bool {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	return numericSeparatorsQuark{}.
		ptr().numStrSepDtosAreEqual(
		numSeps,
		&numSep2)
}

// GetDecimalSeparator - Returns the decimal separator character.
//
// The decimal separator separates integer and fractional digits
// in a floating point number string. In the United States, the
// decimal separator is the period ('.') or decimal point.
//
//             Example: 1234.456
//
func (numSeps *NumericSeparators) GetDecimalSeparator() rune {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	return numSeps.decimalSeparator
}

// GetIntegerDigitSeparators - Returns the integer digit separators
// array.
//
// The actual value returned is an array of type NumStrIntSeparator
// which contains integer separator characters and grouping
// sequences. This complexity is required in order to support
// countries and cultures with integer groupings other than
// thousands.
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
//  intSeps             []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation in number strings.
//
//        type NumStrIntSeparator struct {
//          intSeparatorChar     rune   // Integer separator character
//          intSeparatorGrouping uint   // Number of integers in a group
//          intSeparatorRepetitions uint   // Number of times this character/group is repeated
//                                         // A zero value signals unlimited repetitions.
//        }
//
//         intSeparatorChar     rune
//         - This separator is commonly known as the 'thousands'
//           separator. It is used to separate groups of integer
//           digits to the left of the decimal separator (a.k.a.
//           decimal point). In the United States, the standard
//           integer digits separator is the comma (','). Other
//           countries use periods, spaces or apostrophes to
//           separate integers.
//             United States Example:  1,000,000,000
//              numSeps.intSeparators =
//                []NumStrIntSeparator{
//                     {
//                     intSeparatorChar:   ',',
//                     intSeparatorGrouping: 3,
//                     intSeparatorRepetitions: 0,
//                     },
//                  }
//
//         intSeparatorGrouping []uint
//         - In most western countries integer digits to the left
//           of the decimal separator (a.k.a. decimal point) are
//           separated into groups of three digits representing
//           a grouping of 'thousands' like this: '1,000,000,000'.
//           In this case the intSeparatorGrouping value would be
//           set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerSeparators' would be configured as
//       follows:
//       as:
//
//       numSeps.intSeparators =
//         []NumStrIntSeparator{
//              {
//              intSeparatorChar:   ',',
//              intSeparatorGrouping: 3,
//              intSeparatorRepetitions: 1,
//              },
//              {
//              intSeparatorChar:     ',',
//              intSeparatorGrouping: 2,
//              intSeparatorRepetitions: 0,
//              },
//           }
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
	intSeps []NumStrIntSeparator,
	err error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.GetIntegerDigitSeparators() ")

	intSeps = make([]NumStrIntSeparator, 0, 5)

	nStrFmtSpecDigitsSepsQuark := numericSeparatorsQuark{}

	_,
		err = nStrFmtSpecDigitsSepsQuark.testValidityOfNumericSeparators(
		numSeps,
		ePrefix.XCtx("numSeps"))

	if err != nil {
		return intSeps, err
	}

	lenIntSeps := len(numSeps.integerSeparators)

	if lenIntSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Internal Integer Separators array is invalid!\n"+
			"'numSeps.integerSeparators' is a ZERO length array.\n",
			ePrefix.String())
		return intSeps, err
	}

	intSeps =
		make([]NumStrIntSeparator, lenIntSeps, 5)

	for i := 0; i < lenIntSeps; i++ {
		err =
			intSeps[i].CopyIn(
				&numSeps.integerSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf("numSeps.integerSeparators[%v]",
						i)))

		if err != nil {
			return intSeps, err
		}
	}

	return intSeps, err
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

	nStrFmtSpecDigitsSepsQuark := numericSeparatorsQuark{}

	_,
		err := nStrFmtSpecDigitsSepsQuark.testValidityOfNumericSeparators(
		numSeps,
		ePrefix)

	return err
}

// New - Returns a new instance of NumericSeparators.
//
// The values are automatically set to USA defaults.
//
//  Decimal Separator = '.'
//  Thousands Separator (a.k.a Integer Digits Separator) = ','
//  Integer Digits Grouping Sequence = []uint{3}
//
//
func (numSeps NumericSeparators) New() NumericSeparators {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	newNumSep := NumericSeparators{}

	newNumSep.decimalSeparator = '.'

	newNumSep.integerSeparators =
		make([]NumStrIntSeparator, 1, 2)

	newNumSep.integerSeparators[0].intSeparatorChar = ','
	newNumSep.integerSeparators[0].intSeparatorGrouping = 3

	newNumSep.lock = new(sync.Mutex)

	return newNumSep
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
//  decimalSeparator              rune
//     - A text character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the standard decimal separator is the period
//       ('.') or decimal point.
//
//
//  integerSeparators             []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation.
//
//        type NumStrIntSeparator struct {
//          intSeparatorChar     rune   // Integer separator character
//          intSeparatorGrouping uint   // Number of integers in a group
//          intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                       // A zero value signals unlimited repetitions.
//        }
//
//         intSeparatorChar     rune
//         - This separator is commonly known as the 'thousands'
//           separator. It is used to separate groups of integer
//           digits to the left of the decimal separator (a.k.a.
//           decimal point). In the United States, the standard
//           integer digits separator is the comma (','). Other
//           countries use periods, spaces or apostrophes to
//           separate integers.
//             United States Example:  1,000,000,000
//              numSeps.intSeparators =
//                []NumStrIntSeparator{
//                     {
//                     intSeparatorChar:   ',',
//                     intSeparatorGrouping: 3,
//                     intSeparatorRepetitions: 0,
//                     },
//                  }
//
//         intSeparatorGrouping []uint
//         - In most western countries integer digits to the left
//           of the decimal separator (a.k.a. decimal point) are
//           separated into groups of three digits representing
//           a grouping of 'thousands' like this: '1,000,000,000'.
//           In this case the intSeparatorGrouping value would be
//           set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerSeparators' would be configured as
//       follows:
//       as:
//
//       numSeps.intSeparators =
//         []NumStrIntSeparator{
//              {
//              intSeparatorChar:   ',',
//              intSeparatorGrouping: 3,
//              intSeparatorRepetitions: 1,
//              },
//              {
//              intSeparatorChar:     ',',
//              intSeparatorGrouping: 2,
//              intSeparatorRepetitions: 0,
//              },
//           }
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
	decimalSeparator rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref("NumericSeparators.NewWithComponents()")

	newDigitsSepsDto := NumericSeparators{}

	numericSeparatorDtoMech :=
		numericSeparatorsMechanics{}

	err := numericSeparatorDtoMech.setIntDigitsSeps(
		&newDigitsSepsDto,
		decimalSeparator,
		integerSeparators,
		ePrefix)

	return newDigitsSepsDto, err
}

// NewWithDefaults - Creates and returns a new instance of
// NumericSeparators. This type encapsulates the
// digit separators used in formatting a number string for text
// display. Digit separators are commonly referred to as the
// decimal point and thousands separator characters. In addition,
// Digit Separators include the integer digits grouping sequence.
//
// This method default the integer digits grouping sequence to that
// used in most of the world.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = []uint{3}
// This method applies the 3-digit integer digits grouping sequence
// by default.
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted as like this:
// '6,78,90,00,00,00,00,000'. The right most group has three digits
// and all the others are grouped by two. In this case 'integer
// digits grouping sequence' would be configured as:
//        integerDigitsGroupingSequence = []uint{3,2}
//
// Again, this method will automatically set the 'integer digits
// grouping sequence' to a default of 3-digits or []uint{3}.
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
//  decimalSeparator           rune
//     - A text character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the standard decimal separator is the period
//       ('.') or decimal point.
//
//       If this input parameter is set to zero, an error will be
//       returned.
//
//
//  integerDigitsSeparator     rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//       If this input parameter is set to zero, an error will be
//       returned.
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
//       returned. The 'integer digits grouping sequence' will be
//       automatically set to default of 3-digits or []uint{3}.
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
func (numSeps NumericSeparators) NewWithDefaults(
	ePrefix *ErrPrefixDto,
	decimalSeparator rune,
	integerDigitsSeparator rune) (
	NumericSeparators,
	error) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"NumericSeparators.NewWithDefaults()")

	newDigitsSepsDto := NumericSeparators{}

	nStrFmtSpecDigitsSepsDtoMech :=
		numericSeparatorsMechanics{}

	err := nStrFmtSpecDigitsSepsDtoMech.setIntDigitsSeps(
		&newDigitsSepsDto,
		decimalSeparator,
		[]NumStrIntSeparator{
			{
				intSeparatorChar:     integerDigitsSeparator,
				intSeparatorGrouping: 3,
			},
		},
		ePrefix)

	return newDigitsSepsDto, err

}

// SetDecimalSeparator - Sets the value of the decimal separator.
//
// The decimal separator separates integer and fractional digits
// in a floating point number string. In the United States, the
// decimal separator is the period ('.') or decimal point.
//
//             Example: 1234.456
//
//
func (numSeps *NumericSeparators) SetDecimalSeparator(
	decimalSeparator rune) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	numSeps.decimalSeparator = decimalSeparator

}

// SetNumSeps - This method will set all of the member variable
// data values for the current instance of NumericSeparators.
//
// The NumericSeparators type encapsulates the digit separators
// used in formatting a number string for text display.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparator               rune
//     - A text character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the standard decimal separator is the period
//       ('.') or decimal point.
//
//
//  integerSeparators             []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation.
//
//        type NumStrIntSeparator struct {
//          intSeparatorChar     rune   // Integer separator character
//          intSeparatorGrouping uint   // Number of integers in a group
//          intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                       // A zero value signals unlimited repetitions.
//        }
//
//         intSeparatorChar     rune
//         - This separator is commonly known as the 'thousands'
//           separator. It is used to separate groups of integer
//           digits to the left of the decimal separator (a.k.a.
//           decimal point). In the United States, the standard
//           integer digits separator is the comma (','). Other
//           countries use periods, spaces or apostrophes to
//           separate integers.
//             United States Example:  1,000,000,000
//              numSeps.intSeparators =
//                []NumStrIntSeparator{
//                     {
//                     intSeparatorChar:   ',',
//                     intSeparatorGrouping: 3,
//                     intSeparatorRepetitions: 0,
//                     },
//                  }
//
//
//         intSeparatorGrouping []uint
//         - In most western countries integer digits to the left
//           of the decimal separator (a.k.a. decimal point) are
//           separated into groups of three digits representing
//           a grouping of 'thousands' like this: '1,000,000,000'.
//           In this case the intSeparatorGrouping value would be
//           set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerSeparators' would be configured as
//       follows:
//       as:
//
//       numSeps.intSeparators =
//         []NumStrIntSeparator{
//              {
//              intSeparatorChar:   ',',
//              intSeparatorGrouping: 3,
//              intSeparatorRepetitions: 1,
//              },
//              {
//              intSeparatorChar:     ',',
//              intSeparatorGrouping: 2,
//              intSeparatorRepetitions: 0,
//              },
//           }
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
func (numSeps *NumericSeparators) SetNumSeps(
	decimalSeparator rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) error {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	ePrefix.SetEPref("NumericSeparators.SetNumSeps()")

	nStrFmtSpecDigitsSepsDtoMech :=
		numericSeparatorsMechanics{}

	return nStrFmtSpecDigitsSepsDtoMech.setIntDigitsSeps(
		numSeps,
		decimalSeparator,
		integerSeparators,
		ePrefix.XCtx("Setting Data Values for current instance 'numSeps'"))
}

// SetIntegerSeparators - Sets the value of the integer digit
// separators.
//
// This separator is used in formatting a string of numeric digits.
// It is also known as the 'thousands' separator. The integer
// digits separator is used to separate groups of integer digits to
// the left of the decimal separator (a.k.a. decimal point). In the
// United States, the standard integer digits separator is the comma
// (',').
//
//        Example:  1,000,000,000
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerSeparators             []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation.
//
//        type NumStrIntSeparator struct {
//          intSeparatorChar     rune   // Integer separator character
//          intSeparatorGrouping uint   // Number of integers in a group
//          intSeparatorRepetitions uint // Number of times this character/group is repeated
//                                       // A zero value signals unlimited repetitions.
//        }
//
//         intSeparatorChar     rune
//         - This separator is commonly known as the 'thousands'
//           separator. It is used to separate groups of integer
//           digits to the left of the decimal separator (a.k.a.
//           decimal point). In the United States, the standard
//           integer digits separator is the comma (','). Other
//           countries use periods, spaces or apostrophes to
//           separate integers.
//             United States Example:  1,000,000,000
//              numSeps.intSeparators =
//                []NumStrIntSeparator{
//                     {
//                     intSeparatorChar:   ',',
//                     intSeparatorGrouping: 3,
//                     intSeparatorRepetitions: 0,
//                     },
//                  }
//
//
//         intSeparatorGrouping []uint
//         - In most western countries integer digits to the left
//           of the decimal separator (a.k.a. decimal point) are
//           separated into groups of three digits representing
//           a grouping of 'thousands' like this: '1,000,000,000'.
//           In this case the intSeparatorGrouping value would be
//           set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerSeparators' would be configured as
//       follows:
//       as:
//
//       numSeps.intSeparators =
//         []NumStrIntSeparator{
//              {
//              intSeparatorChar:   ',',
//              intSeparatorGrouping: 3,
//              intSeparatorRepetitions: 1,
//              },
//              {
//              intSeparatorChar:     ',',
//              intSeparatorGrouping: 2,
//              intSeparatorRepetitions: 0,
//              },
//           }
//
func (numSeps *NumericSeparators) SetIntegerSeparators(
	integerSeparators []NumStrIntSeparator) {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	if integerSeparators == nil {
		return
	}

	lenIntSeparators := len(integerSeparators)

	if lenIntSeparators == 0 {
		return
	}

	for i := 0; i < lenIntSeparators; i++ {
		if !integerSeparators[i].IsValidInstance() {
			return
		}
	}

	numSeps.integerSeparators =
		make([]NumStrIntSeparator, lenIntSeparators, 10)

	var err error

	for i := 0; i < lenIntSeparators; i++ {

		err =
			numSeps.integerSeparators[i].CopyIn(
				&integerSeparators[i],
				nil)

		if err != nil {
			return
		}

	}

	return
}

// SetToUSADefaultsIfEmpty - If any of the NumericSeparators
// data values are zero, this method will set ALL data elements to USA
// default values.
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

	if !isValid {

		numSeps.decimalSeparator = '.'

		numSeps.integerSeparators =
			make([]NumStrIntSeparator, 1, 5)

		numSeps.integerSeparators[0].
			intSeparatorChar = ','

		numSeps.integerSeparators[0].
			intSeparatorGrouping = 3

	}

}

// SetToUSADefaults - Sets all of the numeric separator
// member variables to default USA values.
//
// USA default numeric separators are listed as follows:
//
//   Thousands Separator
//    (a.k.a.) Integer Digits Separator: ','
//
//   Decimal Separator: '.'
//
func (numSeps *NumericSeparators) SetToUSADefaults() {

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	numSeps.lock.Lock()

	defer numSeps.lock.Unlock()

	numSeps.decimalSeparator = '.'

	numSeps.integerSeparators =
		make([]NumStrIntSeparator, 1, 5)

	numSeps.integerSeparators[0].
		intSeparatorChar = ','

	numSeps.integerSeparators[0].
		intSeparatorGrouping = 3

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

	str := fmt.Sprintf("Decimal Separator: %v\n",
		string(numSeps.decimalSeparator))

	if numSeps.integerSeparators == nil {
		numSeps.integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntSeps := len(numSeps.integerSeparators)

	for i := 0; i < lenIntSeps; i++ {
		str += fmt.Sprintf("Integer Separator Char= '%v' "+
			"Integer Grouping= '%v'\n",
			string(numSeps.integerSeparators[i].intSeparatorChar),
			numSeps.integerSeparators[i].intSeparatorGrouping)

	}

	return str
}
