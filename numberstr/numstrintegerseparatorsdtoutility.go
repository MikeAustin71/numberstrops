package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorsDtoUtility struct {
	lock *sync.Mutex
}

// addBasic - Adds one NumStrIntSeparator object to the existing
// collection of NumStrIntSeparator objects encapsulated by the
// input NumStrIntSeparatorsDto instance, 'intSep' .
//
// The new NumStrIntSeparator object is generated from default
// values and a string specifying the integer digits separator.
//
// For the newly added NumStrIntSeparator array element, the input
// parameter 'integerDigitsSeparators' specifies the integer
// separator character or characters. The integer digit grouping is
// defaulted to a value of three (3). The 'separator repetitions'
// value is defaulted to zero (0), signaling unlimited
// repetitions. Finally, the 'restartIntGroupingSequence' flag will
// be defaulted to 'false'.
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
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// Again, this method will automatically set the 'integer digits
// grouping sequence' to a default of 3-digits.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto                 *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. A new
//       instance NumStrIntSeparator will be created and added to
//       the collection of NumStrIntSeparator objects encapsulated
//       by this NumStrIntSeparatorsDto instance.
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
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. An error value of 'nil'
//       signals that the new NumStrIntSeparator object was
//       successfully added to the collection managed by the
//       'intSepsDto' NumStrIntSeparatorsDto instance.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (intSepsDtoUtil *numStrIntSeparatorsDtoUtility) addBasic(
	intSepsDto *NumStrIntSeparatorsDto,
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) (
	err error) {

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorUtility." +
			"addBasic()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	var newNumStrIntSep NumStrIntSeparator

	newNumStrIntSep,
		err = NumStrIntSeparator{}.NewBasic(
		integerDigitsSeparators,
		ePrefix.XCtx(
			"newNumStrIntSep"))

	if err != nil {
		return err
	}

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	intSepsDto.intSeparators =
		append(intSepsDto.intSeparators, newNumStrIntSep)

	return err
}

// addBasicRunes - Adds one NumStrIntSeparator object to the
// existing collection of NumStrIntSeparator objects encapsulated
// by the input NumStrIntSeparatorsDto instance, 'intSep' .
//
// The new NumStrIntSeparator object is generated from default
// values and a string specifying the integer digits separator.
//
// For the newly added NumStrIntSeparator array element, the input
// parameter 'integerDigitsSeparators' specifies the integer
// separator character or characters. The integer digit grouping is
// defaulted to a value of three (3). The 'separator repetitions'
// value is defaulted to zero (0), signaling unlimited
// repetitions. Finally, the 'restartIntGroupingSequence' flag will
// be defaulted to 'false'.
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
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//    Chinese Numerals Example: '12,3456,7890,2345'
//
// Again, this method will automatically set the 'integer digits
// grouping sequence' to a default of 3-digits.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto                 *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. A new
//       instance NumStrIntSeparator will be created and added to
//       the collection of NumStrIntSeparator objects encapsulated
//       by this NumStrIntSeparatorsDto instance.
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
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. An error value of 'nil'
//       signals that the new NumStrIntSeparator object was
//       successfully added to the collection managed by the
//       'intSepsDto' NumStrIntSeparatorsDto instance.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (intSepsDtoUtil *numStrIntSeparatorsDtoUtility) addBasicRunes(
	intSepsDto *NumStrIntSeparatorsDto,
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) (
	err error) {

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorUtility." +
			"addBasicRunes()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	var newNumStrIntSep NumStrIntSeparator

	newNumStrIntSep,
		err = NumStrIntSeparator{}.NewBasicRunes(
		integerDigitsSeparators,
		ePrefix.XCtx(
			"newNumStrIntSep"))

	if err != nil {
		return err
	}

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	intSepsDto.intSeparators =
		append(intSepsDto.intSeparators, newNumStrIntSep)

	return err
}

// addWithComponents - Adds one NumStrIntSeparator object to the
// existing collection of NumStrIntSeparator objects encapsulated
// by the input NumStrIntSeparatorsDto parameter, 'intSepsDto'.
//
// The new NumStrIntSeparator member of the collection is created
// from the following input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto                 *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. A new
//       instance NumStrIntSeparator will be created and added to
//       the collection of NumStrIntSeparator objects encapsulated
//       by this NumStrIntSeparatorsDto instance.
//
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
//       'intSepsDto' NumStrIntSeparatorsDto instance.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (intSepsDtoUtil *numStrIntSeparatorsDtoUtility) addWithComponents(
	intSepsDto *NumStrIntSeparatorsDto,
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool,
	ePrefix *ErrPrefixDto) (
	err error) {

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorUtility." +
			"addWithComponents()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	var newNumStrIntSep NumStrIntSeparator

	newNumStrIntSep,
		err = NumStrIntSeparator{}.NewWithComponents(
		intSeparatorChars,
		intSeparatorGrouping,
		intSeparatorRepetitions,
		restartIntGroupingSequence,
		ePrefix.XCtx(
			"newNumStrIntSep"))

	if err != nil {
		return err
	}

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	intSepsDto.intSeparators =
		append(intSepsDto.intSeparators, newNumStrIntSep)

	return err
}

// ptr - Returns a pointer to a new instance of
// numStrIntSeparatorsDtoUtility.
//
func (intSepsDtoUtil numStrIntSeparatorsDtoUtility) ptr() *numStrIntSeparatorsDtoUtility {

	if intSepsDtoUtil.lock == nil {
		intSepsDtoUtil.lock = new(sync.Mutex)
	}

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	newMech := new(numStrIntSeparatorsDtoUtility)

	newMech.lock = new(sync.Mutex)

	return newMech
}

// setBasic - Overwrites all the member variable data values for
// the input parameter 'intSepsDto', an instance of
// NumStrIntSeparatorsDto. This method is intended to configure a
// basic or simple NumStrIntSeparatorsDto object using a minimum
// number of input parameters and specified default values.
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
//     numStrIntSeparatorsDtoMechanics.setWithComponents()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSepsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto                 *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. Member
//       variable data values will be overwritten and reset to new
//       values generated from default values and input parameter
//       'integerDigitsSeparators'.
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
func (intSepsDtoUtil *numStrIntSeparatorsDtoUtility) setBasic(
	intSepsDto *NumStrIntSeparatorsDto,
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if intSepsDtoUtil.lock == nil {
		intSepsDtoUtil.lock = new(sync.Mutex)
	}

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorsDtoUtility." +
			"setBasic()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	if len(integerDigitsSeparators) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is zero length string.\n",
			ePrefix.String())

		return err
	}

	err = numStrIntSeparatorsDtoMechanics{}.ptr().
		setWithComponents(
			intSepsDto,
			[]NumStrIntSeparator{
				{
					intSeparatorChars:          []rune(integerDigitsSeparators),
					intSeparatorGrouping:       3,
					intSeparatorRepetitions:    0,
					restartIntGroupingSequence: false,
				},
			},
			ePrefix.XCtx("intSeparatorsDto"))

	return err
}

// setBasicRunes - Overwrites all the member variable data values for
// the input parameter 'intSepsDto', an instance of
// NumStrIntSeparatorsDto. This method is intended to configure a
// basic or simple NumStrIntSeparatorsDto object using a minimum
// number of input parameters and specified default values.
//
// The NumStrIntSeparatorsDto type encapsulates the integer digit
// separators used in formatting a number string for text display.
// It serves to manage a collection of NumStrIntSeparator objects.
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
//     numStrIntSeparatorsDtoMechanics.setWithComponents()
//
// This method is an alternative to method
// numStrIntSeparatorsDtoUtility.SetBasic() in that this method
// accepts integer separator characters as an array of runes
// instead of a string.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSepsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto                 *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. Member
//       variable data values will be overwritten and reset to new
//       values generated from default values and input parameter
//       'integerDigitsSeparators'.
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
func (intSepsDtoUtil *numStrIntSeparatorsDtoUtility) setBasicRunes(
	intSepsDto *NumStrIntSeparatorsDto,
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) (
	err error) {

	if intSepsDtoUtil.lock == nil {
		intSepsDtoUtil.lock = new(sync.Mutex)
	}

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorsDtoUtility." +
			"setBasicRunes()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	if len(integerDigitsSeparators) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is zero length string.\n",
			ePrefix.String())

		return err
	}

	err = numStrIntSeparatorsDtoMechanics{}.ptr().
		setWithComponents(
			intSepsDto,
			[]NumStrIntSeparator{
				{
					intSeparatorChars:          integerDigitsSeparators,
					intSeparatorGrouping:       3,
					intSeparatorRepetitions:    0,
					restartIntGroupingSequence: false,
				},
			},
			ePrefix.XCtx("intSeparatorsDto"))

	return err
}

// setToUSADefaultsIfEmpty - If any of the NumStrIntSeparatorsDto
// data values are zero or invalid, this method will reset ALL data
// elements to United States default values.
//
// If the input parameter 'intSepsDto' object is valid and
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
// input parameter, 'intSepsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSepsDto                 *NumStrIntSeparatorsDto
//     - A pointer to an instance of NumStrIntSeparatorsDto. If
//       this object is invalid or contains zero data values, all
//       member variable data values will be overwritten and reset
//       to United States default integer separator values.
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
func (intSepsDtoUtil *numStrIntSeparatorsDtoUtility) setToUSADefaultsIfEmpty(
	intSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if intSepsDtoUtil.lock == nil {
		intSepsDtoUtil.lock = new(sync.Mutex)
	}

	intSepsDtoUtil.lock.Lock()

	defer intSepsDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorsDtoUtility." +
			"setToUSADefaultsIfEmpty()")

	if intSepsDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	_,
		err := numStrIntSeparatorsDtoQuark{}.ptr().
		testValidityOfNumStrIntSepsDto(
			intSepsDto,
			ePrefix.XCtx("intSepsDto"))

	if err == nil {
		return err
	}

	return numStrIntSeparatorsDtoMechanics{}.ptr().
		setToUSADefaults(
			intSepsDto,
			ePrefix.XCtx("intSeparatorsDto"))
}
