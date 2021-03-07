package numberstr

import "sync"

// NumStrIntSeparatorsDto is used to transport and manage an array
// of NumStrIntSeparator. These integer separators are in turn
// used to control the grouping and separation characters
// employed in separating digits within an integer number.
//
//  United States Example:        1,000,000,000
//  European Example:             1 000 000 000
//  Indian Number System Example: 6,78,90,00,00,00,00,000
//
type NumStrIntSeparatorsDto struct {
	intSeparators []NumStrIntSeparator
	lock          *sync.Mutex
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
			nil)

	return err
}

// NewWithComponents - Creates and returns a new instance of
// NumStrIntSeparatorsDto. The new is generated from the array of
// NumStrIntSeparator objects passed through input parameter,
// 'integerSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerSeparators   []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation in number strings.
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
//  newIntSepsDto       NumStrIntSeparatorsDto
//     - If this method completes successfully, this parameter will
//       return a new, populated instance of NumStrIntSeparatorsDto.
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

// NewWithDefaults - Returns a new instance of
// NumStrIntSeparatorsDto configured for United States default
// values.
//
// The integer separation values used in the United States are
// listed as follows:
//  NumStrIntSeparatorsDto.intSeparators[0].intSeparatorChar = ','
//  NumStrIntSeparatorsDto.intSeparators[0].intSeparatorGrouping = 3
//  NumStrIntSeparatorsDto.intSeparators[0].intSeparatorRepetitions = 0
//
func (intSeparatorsDto NumStrIntSeparatorsDto) NewWithDefaults() NumStrIntSeparatorsDto {

	if intSeparatorsDto.lock == nil {
		intSeparatorsDto.lock = new(sync.Mutex)
	}

	intSeparatorsDto.lock.Lock()

	defer intSeparatorsDto.lock.Unlock()

	newIntSepsDto := NumStrIntSeparatorsDto{}

	_ = numStrIntSeparatorsDtoMechanics{}.ptr().
		setWithUnitedStatesDefaults(
			&newIntSepsDto,
			nil)

	return newIntSepsDto
}

// SetWithComponents - This method will set all of the member
// variable data values for the current instance of
// NumStrIntSeparatorsDto.
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
//  integerSeparators   []NumStrIntSeparator
//     - An array of NumStrIntSeparator elements used to specify
//       the integer separation operation in number strings.
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
