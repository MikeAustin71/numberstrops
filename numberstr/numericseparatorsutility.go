package numberstr

import (
	"fmt"
	"sync"
)

type numericSeparatorsUtility struct {
	lock *sync.Mutex
}

// setWithComponents - Transfers new data to an instance of
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
			"setWithUSADefaults()")

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
