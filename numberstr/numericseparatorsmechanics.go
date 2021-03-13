package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numericSeparatorsMechanics.
//
func (numSepsMech numericSeparatorsMechanics) ptr() *numericSeparatorsMechanics {

	if numSepsMech.lock == nil {
		numSepsMech.lock = new(sync.Mutex)
	}

	numSepsMech.lock.Lock()

	defer numSepsMech.lock.Unlock()

	newMech := new(numericSeparatorsMechanics)

	newMech.lock = new(sync.Mutex)

	return newMech
}

// setIntDigitsSeps - Transfers new data to an instance of
// NumericSeparators. After completion, all the data fields within
// input parameter 'numSeps' will be overwritten.
//
// New data values for 'numSeps' will be generated from component
// elements passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numSeps                       *NumericSeparators
//     - A pointer to an instance of NumericSeparators. The
//       internal member variable data values for this object will
//       be overwritten and set to new values generated from the
//       following input parameters.
//
//
//  decimalSeparators             []rune
//     - One or more text characters used to separate integer and
//       fractional digits in a floating point number string. In
//       the United States, the standard decimal separator is the
//       period ('.') character otherwise known as a decimal point.
//
//
//  integerSeparators             []NumStrIntSeparator
//     - An array of type NumStrIntSeparator containing
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
func (numSepsMech *numericSeparatorsMechanics) setIntDigitsSeps(
	numSeps *NumericSeparators,
	decimalSeparators []rune,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if numSepsMech.lock == nil {
		numSepsMech.lock = new(sync.Mutex)
	}

	numSepsMech.lock.Lock()

	defer numSepsMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numericSeparatorsMechanics." +
			"setIntDigitsSeps()")

	if numSeps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSeps' is invalid!\n"+
			"'numSeps' is a 'nil' pointer\n",
			ePrefix.String())

		return err
	}

	if numSeps.lock == nil {
		numSeps.lock = new(sync.Mutex)
	}

	if decimalSeparators == nil {
		decimalSeparators =
			make([]rune, 0, 5)
	}

	lenDecSeps := len(decimalSeparators)

	if lenDecSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparators' is invalid!\n"+
			"decimalSeparators is a zero length rune array\n",
			ePrefix.String())

		return err
	}

	if integerSeparators == nil {
		integerSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIntDigitSeparators :=
		len(integerSeparators)

	if lenIntDigitSeparators == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerSeparatorsDto' is invalid!\n"+
			"'integerSeparatorsDto' is a zero length or empty array.\n",
			ePrefix.String())

		return err
	}

	newNumericSeps := NumericSeparators{}

	newNumericSeps.decimalSeparators =
		make([]rune, lenDecSeps, lenDecSeps+5)

	for i := 0; i < lenDecSeps; i++ {
		newNumericSeps.decimalSeparators[i] =
			decimalSeparators[i]
	}

	newNumericSeps.integerSeparatorsDto.intSeparators =
		make([]NumStrIntSeparator,
			lenIntDigitSeparators,
			lenIntDigitSeparators+5)

	for i := 0; i < lenIntDigitSeparators; i++ {
		err =
			newNumericSeps.integerSeparatorsDto.
				intSeparators[i].
				CopyIn(
					&integerSeparators[i],
					ePrefix.XCtx(
						fmt.Sprintf(
							"integerSeparatorsDto[%v]",
							i)))

		if err != nil {
			return err
		}

	}

	newNumericSeps.lock = new(sync.Mutex)

	nStrFmtSpecDigitsSepsQuark := numericSeparatorsQuark{}
	_,
		err =
		nStrFmtSpecDigitsSepsQuark.testValidityOfNumericSeparators(
			&newNumericSeps,
			ePrefix.XCtx("Testing validity of preliminary 'newNumericSeps'"))

	if err != nil {
		return err
	}

	nStrFmtSpecDigitsSepsElectron :=
		numericSeparatorsElectron{}

	err =
		nStrFmtSpecDigitsSepsElectron.copyIn(
			numSeps,
			&newNumericSeps,
			ePrefix.XCtx("newNumericSeps->numSeps"))

	return err
}
