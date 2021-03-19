package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numericSeparatorsUtility.
//
func (numSepsUtility numericSeparatorsUtility) ptr() *numericSeparatorsUtility {

	if numSepsUtility.lock == nil {
		numSepsUtility.lock = new(sync.Mutex)
	}

	numSepsUtility.lock.Lock()

	defer numSepsUtility.lock.Unlock()

	newUtility := new(numericSeparatorsUtility)

	newUtility.lock = new(sync.Mutex)

	return newUtility
}

// setBasic - Overwrites all the member variable data values
// for the input parameter 'numSeps' with numeric separator values
// generated from a minimum number of input parameters and
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
//     numericSeparatorsMechanics.setWithComponents()
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
//  numSeps             *NumericSeparators
//     - A pointer to an instance of NumericSeparators. The
//       internal member variable data values for this object will
//       be overwritten and set to new values based on the
//       following input parameters.
//
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
func (numSepsUtility *numericSeparatorsUtility) setBasic(
	numSeps *NumericSeparators,
	decimalSeparators string,
	integerDigitsSeparators string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsUtility.lock == nil {
		numSepsUtility.lock = new(sync.Mutex)
	}

	numSepsUtility.lock.Lock()

	defer numSepsUtility.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsUtility." +
			"setBasicRunes()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(decimalSeparators) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparators' is invalid!\n"+
			"'decimalSeparators' is zero length string.\n",
			ePrefix.String())
	}

	var integerSeparatorsDto NumStrIntSeparatorsDto

	integerSeparatorsDto,
		err =
		NumStrIntSeparatorsDto{}.NewBasic(
			integerDigitsSeparators,
			ePrefix.XCtx(
				fmt.Sprintf("integerDigitsSeparators='%v'",
					integerDigitsSeparators)))

	if err != nil {
		return err
	}

	err = numericSeparatorsMechanics{}.ptr().setDetail(
		numSeps,
		decimalSeparators,
		&integerSeparatorsDto,
		ePrefix.XCtx(
			"numSeps"))

	return err
}

// setBasicRunes - Overwrites all the member variable data values
// for the input parameter 'numSeps' with numeric separator values
// generated from a minimum number of input parameters and
// specified default values.
//
// This method is an alternative to numericSeparatorsUtility.SetBasic()
// in that this method accepts decimal separators and integer digit
// separators as rune arrays.
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
//     numericSeparatorsMechanics.setWithComponents()
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter, 'numSeps'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numSeps                    *NumericSeparators
//     - A pointer to an instance of NumericSeparators. The
//       internal member variable data values for this object will
//       be overwritten and set to new values based on the
//       following input parameters.
//
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
func (numSepsUtility *numericSeparatorsUtility) setBasicRunes(
	numSeps *NumericSeparators,
	decimalSeparators []rune,
	integerDigitsSeparators []rune,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsUtility.lock == nil {
		numSepsUtility.lock = new(sync.Mutex)
	}

	numSepsUtility.lock.Lock()

	defer numSepsUtility.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsUtility." +
			"setBasicRunes()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if decimalSeparators == nil {
		decimalSeparators =
			make([]rune, 0)
	}

	if len(decimalSeparators) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparators' is invalid!\n"+
			"'decimalSeparators' is zero length rune array.\n",
			ePrefix.String())
	}

	if integerDigitsSeparators == nil {
		integerDigitsSeparators =
			make([]rune, 0)
	}

	if len(integerDigitsSeparators) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is a zero length rune array.\n",
			ePrefix.String())
	}

	var integerSeparatorsDto NumStrIntSeparatorsDto

	integerSeparatorsDto,
		err =
		NumStrIntSeparatorsDto{}.NewBasicRunes(
			integerDigitsSeparators,
			ePrefix.XCtx(
				fmt.Sprintf("integerDigitsSeparators='%v'",
					string(integerDigitsSeparators))))

	err = numericSeparatorsMechanics{}.ptr().setDetailRunes(
		numSeps,
		decimalSeparators,
		&integerSeparatorsDto,
		ePrefix.XCtx(
			"numSeps"))

	return err
}

// setWithUSADefaults - Transfers new data to an instance of
// NumericSeparators. After completion, all the data fields within
// input parameter 'numSeps' will be overwritten.
//
// The values numeric separator values configured for the 'numSeps'
// object will be United States default numeric separation values.
//
// These default values are defined as follows:
//
//  Decimal Separator: '.'
//  Integer Separator: ','
//  Integer Grouping:  3
//  Example Floating Point Number String: 1,000,000,000.456
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numSeps             *NumericSeparators
//     - A pointer to an instance of NumericSeparators. The
//       internal member variable data values for this object will
//       be overwritten and set to default United States numeric
//       separation values.
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
func (numSepsUtility *numericSeparatorsUtility) setWithUSADefaults(
	numSeps *NumericSeparators,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsUtility.lock == nil {
		numSepsUtility.lock = new(sync.Mutex)
	}

	numSepsUtility.lock.Lock()

	defer numSepsUtility.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsUtility." +
			"setToUSADefaults()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	decimalSeparators := []rune{'.'}

	integerSeparators := []NumStrIntSeparator{
		{
			intSeparatorChars:          []rune{','},
			intSeparatorGrouping:       3,
			intSeparatorRepetitions:    0,
			restartIntGroupingSequence: false,
			lock:                       new(sync.Mutex),
		},
	}

	err =
		numericSeparatorsMechanics{}.ptr().
			setWithComponents(
				numSeps,
				decimalSeparators,
				integerSeparators,
				ePrefix.XCtx(
					"numSeps"))

	return err
}
