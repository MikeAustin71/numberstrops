package numberstr

import (
	"fmt"
	"sync"
)

type numStrDtoUtility struct {
	lock *sync.Mutex
}

// addNumStrs - Adds the numeric values represented by two
// NumStrDto objects and returns the sum of these two numeric
// values as another NumStrDto instance.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSepsDto          NumericSeparatorsDto
//     - An instance of NumericSeparatorsDto which will be used to supply
//       the numeric separators for the new NumStrDto instance returned
//       by this method. Numeric separators include the Thousands
//       Separator, Decimal Separator and the Currency Symbol.
//
//       The data fields included in the NumericSeparatorsDto are
//       listed as follows:
//
//          type NumericSeparatorsDto struct {
//
//            DecimalSeparator   rune // Character used to separate
//                                    //  integer and fractional digits ('.')
//
//            ThousandsSeparator rune // Character used to separate thousands
//                                    //  (1,000,000,000
//
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//       If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), they will
//       be reset to USA default values. USA default numeric
//       separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
//
//
//  addend1             *NumStrDto
//     - A pointer to a NumStrDto instance. The numeric value
//       encapsulated by this NumStrDto instance, 'addend1',
//       will be added to the second input parameter, 'addend2'.
//
//       This method WILL NOT change the values of internal member
//       variables in order to achieve the method's objectives.
//
//      If this instance of NumStrDto proves to be invalid, an
//      error will be returned.
//
//
//  addend2             *NumStrDto
//     -  A pointer to a NumStrDto instance. The numeric value
//        encapsulated by this NumStrDto instance, 'addend2',
//        will be added to the first input parameter, 'addend1'.
//
//       This method WILL NOT change the values of internal member
//       variables in order to achieve the method's objectives.
//
//      If this instance of NumStrDto proves to be invalid, an error
//      will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  sum                NumStrDto
//     - If this method completes successfully, this parameter will
//       encapsulate the numeric sum obtained by adding the numeric
//       values of input parameters, 'addend1' and 'addend2'.
//
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
func (nStrDtoUtil *numStrDtoUtility) addNumStrs(
	numSepsDto NumericSeparatorsDto,
	addend1 *NumStrDto,
	addend2 *NumStrDto,
	ePrefix *ErrPrefixDto) (
	sum NumStrDto,
	err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoUtility.addNumStrs()")

	nStrDtoElectron := numStrDtoElectron{}

	sum = nStrDtoElectron.newBaseZeroNumStrDto(0)
	err = nil

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			addend1,
			ePrefix.XCtx("Initial validity test for 'addend1'"))

	if err != nil {
		return sum, err
	}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			addend2,
			ePrefix.XCtx("Initial validity test for 'addend2'"))

	if err != nil {
		return sum, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	var n1DtoSetup, n2DtoSetup NumStrDto
	var compare int
	var isOrderReversed bool

	nStrDtoMolecule := numStrDtoMolecule{}

	n1DtoSetup,
		n2DtoSetup,
		compare,
		isOrderReversed,
		err =
		nStrDtoMolecule.formatForMathOps(
			addend1,
			addend2,
			ePrefix)

	if err != nil {
		return sum, err
	}

	newSignVal := n1DtoSetup.signVal
	nStrDtoHelper := numStrDtoHelper{}

	if n1DtoSetup.signVal != n2DtoSetup.signVal {
		// Sign Values ARE NOT Equal!
		nStrDtoElectron := numStrDtoElectron{}

		err = nStrDtoElectron.setSignValue(
			&n1DtoSetup,
			1,
			ePrefix.XCtx("n1DtoSetup Sign=1"))

		if err != nil {
			return sum, err
		}

		err = nStrDtoElectron.setSignValue(
			&n2DtoSetup,
			1,
			ePrefix.XCtx("n1DtoSetup Sign=1"))

		if err != nil {
			return sum, err
		}

		if compare == 0 {
			// Numeric Values are Equal!

			sum = nStrDtoElectron.newBaseZeroNumStrDto(
				n1DtoSetup.precision)

			err = nStrDtoElectron.setNumericSeparatorsDto(
				&sum,
				numSepsDto,
				ePrefix.XCtx("sum from values are equal"))

			return sum, err
		}

		// Sign Values are NOW Equal!

		sum,
			err = nStrDtoHelper.signValuesAreEqualSubtractNumStrs(
			numSepsDto,
			&n1DtoSetup,
			&n2DtoSetup,
			isOrderReversed,
			ePrefix.XCtx("n1DtoSetup, n2DtoSetup "))

		if err != nil {
			return sum, err
		}

		var isZeroValue bool

		isZeroValue,
			err = nStrDtoElectron.isNumStrZeroValue(
			&sum,
			ePrefix.XCtx("sum"))

		if err != nil {
			return sum, err
		}

		if isZeroValue {
			newSignVal = 1
		}

		err = nStrDtoElectron.setSignValue(
			&sum,
			newSignVal,
			ePrefix.XCtx("sum"))

		return sum, err
	}

	// Sign Values ARE Already Equal!

	sum,
		err = nStrDtoHelper.signValuesAreEqualAddNumStrs(
		numSepsDto,
		&n1DtoSetup,
		&n2DtoSetup,
		ePrefix)

	return sum, err
}

// multiplyInPlace - Receives two NumStrDto input parameters
// labeled 'numStrDto' and 'multiplier'. The numeric value
// for 'numStrDto' is multiplied by the numeric value of
// the result of this multiplication operation, or product,
// is then stored in 'numStrDto'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSepsDto          NumericSeparatorsDto
//     - An instance of NumericSeparatorsDto which will be used to supply
//       the numeric separators for the new NumStrDto instance returned
//       by this method. Numeric separators include the Thousands
//       Separator, Decimal Separator and the Currency Symbol.
//
//       The data fields included in the NumericSeparatorsDto are
//       listed as follows:
//
//          type NumericSeparatorsDto struct {
//
//            DecimalSeparator   rune // Character used to separate
//                                    //  integer and fractional digits ('.')
//
//            ThousandsSeparator rune // Character used to separate thousands
//                                    //  (1,000,000,000
//
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//       If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), they will
//       be reset to USA default values. USA default numeric
//       separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
//
//
//  numStrDto        *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and overwrite the internal member variable data
//       values of 'numStrDto' in order to achieve the method's
//       objectives.
//
//       The numerical value of 'numStrDto' will be multiplied
//       by the numerical value of input parameter 'multiplier'.
//       The result or product of this multiplication will then
//       be stored in this same NumStrDto instance.
//
//       If this NumStrDto instance proves to be invalid, this
//       method will return an error.
//
//
//  multiplier          *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplier' will be multiplied
//       by the numerical value of input parameter 'numStrDto'.
//       The result, or product, will then be stored in 'numStrDto'.
//
//       If this NumStrDto instance proves to be invalid, an error
//       will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
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
func (nStrDtoUtil *numStrDtoUtility) multiplyInPlace(
	numSepsDto NumericSeparatorsDto,
	numStrDto *NumStrDto,
	multiplier *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoUtility.multiplyInPlace()")

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			numStrDto,
			ePrefix.XCtx("Initial validity test for 'numStrDto'"))

	if err != nil {
		return err
	}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			multiplier,
			ePrefix.XCtx("Initial validity test for 'multiplier'"))

	if err != nil {
		return err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	var productNDto NumStrDto

	nStrDtoHelper := numStrDtoHelper{}

	productNDto,
		err = nStrDtoHelper.multiplyNumStrs(
		numSepsDto,
		numStrDto,
		multiplier,
		ePrefix.XCtx("numStrDto x multiplier"))

	if err != nil {
		return err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.copyIn(
		numStrDto,
		&productNDto,
		ePrefix.XCtx("productNDto->numStrDto"))

	return err
}

// setNumStr - Sets the value of of input parameter 'numStrDto' to
// to the numeric value of the number string input parameter,
// 'numStr'.
//
//  This method WILL CHANGE AND OVERWRITE the internal member
//  variable data values in order to achieve the method's
//  objectives.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto        *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and OVERWRITE the internal member variable data
//       values of 'numStrDto' in order to achieve the method's
//       objectives.
//
//       The NumStrDto instance be populated with the numeric
//       value contained in input parameter 'numStr'.
//
//  numSepsDto          NumericSeparatorsDto
//     - An instance of NumericSeparatorsDto which will be used to
//       supply the numeric separators for input parameter
//       'numStrDto' when it is populated with a new numeric value
//       extracted from input parameter, 'numStr'.
//
//       Numeric separators include the Thousands Separator, Decimal
//       Separator and the Currency Symbol.
//
//       The data fields included in the NumericSeparatorsDto are
//       listed as follows:
//
//          type NumericSeparatorsDto struct {
//
//            decimalSeparator   rune // Character used to separate
//                                    //  integer and fractional digits ('.')
//
//            integerDigitsSeparator rune // Character used to separate thousands
//                                        //  (1,000,000,000)
//
//            integerDigitsGroupingSequence  []uint // Usually []uint{3} for thousands
//          }
//
//       If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), they will
//       be reset to USA default values. USA default numeric
//       separators are listed as follows:
//
//         Thousands Separator: ','
//           Decimal Separator: '.'
//
//
//  numStr              string
//     - A valid number string. A valid number string 'may' be
//       prefixed with a numeric sign value of plus ('+') or
//       minus ('-'). The absence of a leading numeric sign
//       character will default the numeric value to plus or a
//       positive numeric value. A valid number string 'may'
//       also include a decimal delimiter such as a decimal
//       point to separate integer and fractional digits
//       within the number string. With these two exceptions,
//       all other characters in a valid number string must be
//       numeric values represented by text characters between
//       '0' and '9'.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
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
func (nStrDtoUtil *numStrDtoUtility) setNumStr(
	numStrDto *NumStrDto,
	numSepsDto NumericSeparatorsDto,
	numStr string,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoUtility.setNumStr()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrDto' is invalid!\n"+
			"'numStrDto' is a 'nil' pointer.\n",
			ePrefix.String())

	}

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			numStrDto,
			ePrefix.XCtx(
				"Initial validity test for 'numStrDto'"))

	if err != nil {
		return err
	}

	if len(numStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is an empty string!\n",
			ePrefix.String())

		return err
	}

	nStrDtoNanobot := numStrDtoNanobot{}
	var newNumStrDto NumStrDto

	newNumStrDto,
		err = nStrDtoNanobot.newNumStr(
		numSepsDto,
		numStr,
		ePrefix.XCtx(
			"numSepsDto, numStr"))

	if err != nil {
		return err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&newNumStrDto,
		numSepsDto,
		ePrefix.XCtx(
			"numSepsDto->newNumStrDto"))

	if err != nil {
		return err
	}

	err =
		nStrDtoElectron.copyIn(
			numStrDto,
			&newNumStrDto,
			ePrefix.XCtx("newNumStrDto->numStrDto"))

	return err
}

// subtractNumStrs - Subtracts the numeric values represented by two
// NumStrDto objects.
//
// Input parameter 'subtrahend' is subtracted from input parameter
// 'minuend'. The computed 'difference' of this subtraction operation
// is returned as a new instance of NumStrDto.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  minuend             *NumStrDto
//     - A pointer to a NumStrDto instance. The numeric value
//       encapsulated by the second input parameter, 'subtrahend',
//       will be subtracted from this NumStrDto instance, 'minuend'.
//
//       This method WILL NOT change the values of internal member
//       variables in order to achieve the method's objectives.
//
//       If this instance of NumStrDto proves to be invalid, an
//       error will be returned.
//
//
//  subtrahend          *NumStrDto
//     -  A pointer to a NumStrDto instance. The numeric value
//        encapsulated by this NumStrDto instance, 'subtrahend',
//        will be subtracted from the first input parameter,
//        'minuend'.
//
//       This method WILL NOT change the values of internal member
//       variables in order to achieve the method's objectives.
//
//       If this instance of NumStrDto proves to be invalid, an error
//       will be returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If error prefix information is NOT needed, set this
//       parameter to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  difference          NumStrDto
//     - If this method completes successfully, this parameter will
//       encapsulate the numeric difference obtained by subtracting
//       the numeric value of input parameter 'subtrahend' from that
//       of input parameter 'minuend'.
//
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
func (nStrDtoUtil *numStrDtoUtility) subtractNumStrs(
	numSepsDto NumericSeparatorsDto,
	minuend *NumStrDto,
	subtrahend *NumStrDto,
	ePrefix *ErrPrefixDto) (
	difference NumStrDto,
	err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoUtility.multiplyInPlace()")

	nStrDtoElectron := numStrDtoElectron{}

	difference = nStrDtoElectron.newBaseZeroNumStrDto(0)
	err = nil

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			minuend,
			ePrefix.XCtx("Initial validity test for 'minuend'"))

	if err != nil {
		return difference, err
	}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			subtrahend,
			ePrefix.XCtx("Initial validity test for 'subtrahend'"))

	if err != nil {
		return difference, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	var n1NumDto, n2NumDto NumStrDto
	var compare int
	var isReversed bool

	nStrDtoMolecule := numStrDtoMolecule{}

	n1NumDto,
		n2NumDto,
		compare,
		isReversed,
		err =
		nStrDtoMolecule.formatForMathOps(
			minuend,
			subtrahend,
			ePrefix.XCtx(
				"minuend, subtrahend"))

	if err != nil {
		return difference, err
	}

	if compare == 0 {
		// Numeric Values are Equal!
		difference =
			nStrDtoElectron.newBaseZeroNumStrDto(n1NumDto.precision)

		err = nStrDtoElectron.setNumericSeparatorsDto(
			&difference,
			numSepsDto,
			ePrefix.XCtx(
				"&difference, numSepsDto,"))

		return difference, err
	}

	newSignVal := n1NumDto.signVal
	precision := n1NumDto.precision

	if n1NumDto.signVal != n2NumDto.signVal {
		// Sign Values ARE NOT Equal!
		nStrDtoElectron := numStrDtoElectron{}

		err = nStrDtoElectron.setSignValue(
			&n1NumDto,
			1,
			ePrefix.XCtx("n1NumDto Sign=1"))

		if err != nil {
			return difference, err
		}

		err = nStrDtoElectron.setSignValue(
			&n2NumDto,
			1,
			ePrefix.XCtx("n2NumDto Sign=1"))

		if err != nil {
			return difference, err
		}

		// Sign Values are NOW Equal!

		nStrDtoHelper := numStrDtoHelper{}

		difference,
			err =
			nStrDtoHelper.signValuesAreEqualAddNumStrs(
				numSepsDto,
				&n1NumDto,
				&n2NumDto,
				ePrefix)

		if err != nil {
			return difference, err
		}

		err = nStrDtoElectron.setSignValue(
			&difference,
			newSignVal,
			ePrefix.XCtx("difference<-newSignVal"))

		return difference, err
	}

	// Sign Values ARE Equal!
	// Change sign for subtraction
	newSignVal = n1NumDto.signVal

	if isReversed {
		newSignVal = newSignVal * -1
	}

	lenN1AllRunes := len(n1NumDto.absAllNumRunes)

	n1IntAry := make([]int, lenN1AllRunes)
	n2IntAry := make([]int, lenN1AllRunes)
	n3IntAry := make([]int, lenN1AllRunes)

	for i := 0; i < lenN1AllRunes; i++ {

		n1IntAry[i] = int(n1NumDto.absAllNumRunes[i]) - 48
		n2IntAry[i] = int(n2NumDto.absAllNumRunes[i]) - 48

	}

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0
	// Main Subtraction Routine
	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = n1IntAry[j]
		n2 = n2IntAry[j]
		n3 = 0

		if n1-carry-n2 < 0 {
			n1 += 10
			n3 = n1 - n2 - carry
			carry = 1
		} else {
			n3 = n1 - n2 - carry
			carry = 0
		}

		n3IntAry[j] = n3

	}

	nStrDtoMech := numStrDtoMechanics{}

	difference,
		err =
		nStrDtoMech.findIntArraySignificantDigitLimits(
			numSepsDto,
			n3IntAry,
			precision,
			newSignVal,
			ePrefix.XCtx("n3IntAry"))

	return difference, err
}
