package numberstr

import (
	"fmt"
	"math"
	"sync"
)

type numStrDtoHelper struct {
	lock *sync.Mutex
}

// multiplyNumStrs - Multiplies two NumStrDto instances and returns
// the product as a new NumStrDto instance.
//
// If either of the two input parameters, 'multiplicand' or
// 'multiplier' proves to be invalid, this method will return an
// error.
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
//  multiplicand        *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplicand' will be multiplied
//       by input parameter 'multiplier' to generate and return
//       the product as NumStrDto instance.
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
//       by input parameter 'multiplicand' to generate and return
//       the product as NumStrDto instance.
//
//       If this NumStrDto instance proves to be invalid, this
//       method will return an error.
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
//  product             NumStrDto
//     - If this method completes successfully, the product obtained from
//       multiplying input parameters 'multiplicand' and 'multiplier' will
//       be returned in a new instance of NumStrDto.
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
func (nStrDtoHelper *numStrDtoHelper) multiplyNumStrs(
	numSepsDto NumericSeparatorsDto,
	multiplicand *NumStrDto,
	multiplier *NumStrDto,
	ePrefix *ErrPrefixDto) (
	product NumStrDto,
	err error) {

	if nStrDtoHelper.lock == nil {
		nStrDtoHelper.lock = new(sync.Mutex)
	}

	nStrDtoHelper.lock.Lock()

	defer nStrDtoHelper.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoHelper.multiplyNumStrs()")

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	product = nStrDtoElectron.newBaseZeroNumStrDto(0)

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			multiplicand,
			ePrefix.XCtx("multiplicand"))

	if err != nil {
		return product, err
	}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			multiplier,
			ePrefix.XCtx("multiplier"))

	if err != nil {
		return product, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	lenN1AbsAllRunes := len(multiplicand.absAllNumRunes)
	lenN2AbsAllRunes := len(multiplier.absAllNumRunes)

	var n1Setup NumStrDto
	var n2Setup NumStrDto

	if lenN2AbsAllRunes > lenN1AbsAllRunes {

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix.XCtx("multiplier -> n1Setup"))

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix.XCtx("multiplicand -> n2Setup"))

		if err != nil {
			return product, err
		}

	} else if lenN1AbsAllRunes > lenN2AbsAllRunes {

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix.XCtx("multiplicand -> n1Setup"))

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix.XCtx("multiplier -> n2Setup"))

		if err != nil {
			return product, err
		}

	} else {
		// Must be lenN1AbsAllRunes == lenN2AbsAllRunes

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix.XCtx("multiplicand -> n1Setup"))

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix.XCtx("multiplier -> n2Setup"))

		if err != nil {
			return product, err
		}

	}

	newPrecision := n1Setup.precision + n2Setup.precision
	newSignVal := 1

	if n1Setup.signVal == n2Setup.signVal {
		newSignVal = 1
	} else {
		// Must be n1Setup.signVal != n2Setup.signVal
		newSignVal = -1
	}

	lenN1AbsAllRunes = len(n1Setup.absAllNumRunes)
	lenN2AbsAllRunes = len(n2Setup.absAllNumRunes)
	lenLevels := lenN2AbsAllRunes
	lenNumPlaces := (lenN1AbsAllRunes + lenN2AbsAllRunes) + 1

	intMAry := make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		intMAry[i] = make([]int, lenNumPlaces)
	}

	intFinalAry := make([]int, lenNumPlaces+1)

	carry := 0
	levels := 0
	place := 0
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0
	for i := lenN2AbsAllRunes - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := lenN1AbsAllRunes - 1; j >= 0; j-- {

			n1 = int(n1Setup.absAllNumRunes[j]) - 48
			n2 = int(n2Setup.absAllNumRunes[i]) - 48
			n3 = (n1 * n2) + carry
			n4 = int(math.Mod(float64(n3), 10.00))

			intMAry[levels][place] = n4

			carry = n3 / 10

			place--
		}

		intMAry[levels][place] = carry
		carry = 0
		levels++
	}

	carry = 0
	n1 = 0
	n2 = 0
	n3 = 0
	n4 = 0
	for i := 0; i < lenLevels; i++ {
		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = intFinalAry[j+1]
			n2 = intMAry[i][j]
			n3 = n1 + n2 + carry
			n4 = 0

			if n3 > 9 {
				n4 = int(math.Mod(float64(n3), 10.0))
				carry = n3 / 10

			} else {
				n4 = n3
				carry = 0
			}

			intFinalAry[j+1] = n4
		}

		if carry > 0 {
			intFinalAry[0] = carry
		}

	}

	nStrDtoMech := numStrDtoMechanics{}

	product, err =
		nStrDtoMech.findIntArraySignificantDigitLimits(
			numSepsDto,
			intFinalAry,
			newPrecision,
			newSignVal,
			ePrefix.XCtx("intFinalAry"))

	return product, err
}

// signValuesAreEqualAddNumStrs - Adds two numbers with equal
// numeric sign values. Numeric sign values of +1 or -1 indicate
// whether a number is is positive ('+') or negative ('-').
//
// This method is designed to be called after method
// numStrDtoMolecule.formatForMathOps() has been called on these two
// numeric values.
//
// The 'signValuesAreEqualSubtractNumStrs' is a low level method
// designed to be called by numStrDtoUtility.addNumStrs().
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
//  n1NumDto            *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//
//  n2NumDto            *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//
//  isReversed          bool
//     - If 'isReversed' is set to true, it signals that the original
//       sequence of 'n1NumDto' and 'n2NumDto' have been reversed in
//       order to facilitate the subtraction operation.
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
//     - If this method completes successfully, this parameter will represent
//       the numerical difference between input parameters, 'n1NumDto' and
//       'n2NumDto'.
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
func (nStrDtoHelper *numStrDtoHelper) signValuesAreEqualAddNumStrs(
	numSepsDto NumericSeparatorsDto,
	n1DtoSetup *NumStrDto,
	n2DtoSetup *NumStrDto,
	ePrefix *ErrPrefixDto) (
	sum NumStrDto,
	err error) {

	if nStrDtoHelper.lock == nil {
		nStrDtoHelper.lock = new(sync.Mutex)
	}

	nStrDtoHelper.lock.Lock()

	defer nStrDtoHelper.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoHelper.signValuesAreEqualAddNumStrs()")

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			n1DtoSetup,
			ePrefix.XCtx("Initial validity test for 'n1DtoSetup'"))

	if err != nil {
		return sum, err
	}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			n2DtoSetup,
			ePrefix.XCtx("Initial validity test for 'n2DtoSetup'"))

	if err != nil {
		return sum, err
	}

	if n1DtoSetup.signVal != n2DtoSetup.signVal {
		err = fmt.Errorf("%v\n"+
			"Error: Numeric sign values for input parameters\n"+
			"'n1DtoSetup' and 'n2DtoSetup' are NOT equal!\n"+
			"n1DtoSetup.signVal ='%v'\n"+
			"n2DtoSetup.signVal ='%v'\n",
			ePrefix.String(),
			n1DtoSetup.signVal,
			n2DtoSetup.signVal)
		return sum, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	// Sign Values ARE Equal!

	newSignVal := n1DtoSetup.signVal

	precision := n1DtoSetup.precision
	lenN1AllRunes := len(n1DtoSetup.absAllNumRunes)

	n3IntAry := make([]int, lenN1AllRunes+1)
	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = int(n1DtoSetup.absAllNumRunes[j]) - 48
		n2 = int(n2DtoSetup.absAllNumRunes[j]) - 48

		n3 = n1 + n2 + carry

		carry = 0

		if n3 > 9 {
			n3 = n3 - 10
			carry = 1
		}

		n3IntAry[j+1] = n3

	}

	if carry > 0 {
		n3IntAry[0] = carry
	}

	nStrDtoMech := numStrDtoMechanics{}

	sum,
		err = nStrDtoMech.findIntArraySignificantDigitLimits(
		numSepsDto,
		n3IntAry,
		precision,
		newSignVal,
		ePrefix.XCtx(
			"numSepsDto, n3IntAry"))

	return sum, err
}

// signValuesAreEqualSubtractNumStrs - Subtracts two numbers with
// equal numeric sign values. Numeric sign values of +1 or -1
// indicate whether a number is is positive ('+') or negative ('-').
//
// This method is designed to be called after method
// numStrDtoMolecule.formatForMathOps() has been called on these two
// numeric values.
//
// The 'signValuesAreEqualSubtractNumStrs' is a low level method
// designed to be called by numStrDtoUtility.addNumStrs().
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
//  n1NumDto            *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//
//  n2NumDto            *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//
//  isReversed          bool
//     - If 'isReversed' is set to true, it signals that the original
//       sequence of 'n1NumDto' and 'n2NumDto' have been reversed in
//       order to facilitate the subtraction operation.
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
//     - If this method completes successfully, this parameter will represent
//       the numerical difference between input parameters, 'n1NumDto' and
//       'n2NumDto'.
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
func (nStrDtoHelper *numStrDtoHelper) signValuesAreEqualSubtractNumStrs(
	numSepsDto NumericSeparatorsDto,
	n1NumDto *NumStrDto,
	n2NumDto *NumStrDto,
	isReversed bool,
	ePrefix *ErrPrefixDto) (
	difference NumStrDto,
	err error) {

	if nStrDtoHelper.lock == nil {
		nStrDtoHelper.lock = new(sync.Mutex)
	}

	nStrDtoHelper.lock.Lock()

	defer nStrDtoHelper.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoHelper.signValuesAreEqualSubtractNumStrs()")

	// Sign Values ARE Equal!
	// Change sign for subtraction
	newSignVal := n1NumDto.signVal
	precision := n1NumDto.precision

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
			ePrefix.XCtx(
				"numSepsDto, n3IntAry"))

	return difference, err
}
