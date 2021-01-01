package numberstr

import (
	"errors"
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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
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
//  err                error
//     - If this method completes successfully, the returned error Type is
//       set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. The 'ePrefix' text will be
//       prefixed to the beginning of the returned error message.
//
func (nStrDtoHelper *numStrDtoHelper) multiplyNumStrs(
	multiplicand *NumStrDto,
	multiplier *NumStrDto,
	ePrefix string) (
	product NumStrDto,
	err error) {

	if nStrDtoHelper.lock == nil {
		nStrDtoHelper.lock = new(sync.Mutex)
	}

	nStrDtoHelper.lock.Lock()

	defer nStrDtoHelper.lock.Unlock()

	ePrefix += "numStrDtoHelper.multiplyNumStrs() "

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	product = nStrDtoElectron.newBaseZeroNumStrDto(0)

	if multiplicand == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'multiplicand' is INVALID!\n" +
			"multiplicand has a 'nil' pointer!\n")
		return product, err
	}

	if multiplier == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'multiplier' is INVALID!\n" +
			"multiplier  has a 'nil' pointer!!\n")
		return product, err
	}

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			multiplicand,
			ePrefix+"multiplicand ")

	if err != nil {
		return product, err
	}

	_,
		err =
		nStrDtoQuark.testNumStrDtoValidity(
			multiplier,
			ePrefix+"multiplier ")

	if err != nil {
		return product, err
	}

	lenN1AbsAllRunes := len(multiplicand.absAllNumRunes)
	lenN2AbsAllRunes := len(multiplier.absAllNumRunes)

	var n1Setup NumStrDto
	var n2Setup NumStrDto

	if lenN2AbsAllRunes > lenN1AbsAllRunes {

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix+"multiplier -> n1Setup ")

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix+"multiplicand -> n2Setup ")

		if err != nil {
			return product, err
		}

	} else if lenN1AbsAllRunes > lenN2AbsAllRunes {

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix+"multiplicand -> n1Setup ")

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix+"multiplier -> n2Setup ")

		if err != nil {
			return product, err
		}

	} else {
		// Must be lenN1AbsAllRunes == lenN2AbsAllRunes

		n1Setup,
			err = nStrDtoElectron.copyOut(
			multiplicand,
			ePrefix+"multiplicand -> n1Setup ")

		if err != nil {
			return product, err
		}

		n2Setup,
			err = nStrDtoElectron.copyOut(
			multiplier,
			ePrefix+"multiplier -> n2Setup ")

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
			intFinalAry,
			newPrecision,
			newSignVal,
			ePrefix+"intFinalAry ")

	return product, err
}

func (nStrDtoHelper *numStrDtoHelper) signValuesAreEqualAddNumStrs(
	n1DtoSetup *NumStrDto,
	n2DtoSetup *NumStrDto,
	ePrefix string) (
	sum NumStrDto,
	err error) {

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
		n3IntAry,
		precision,
		newSignVal,
		ePrefix)

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
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
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
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrDtoHelper *numStrDtoHelper) signValuesAreEqualSubtractNumStrs(
	n1NumDto *NumStrDto,
	n2NumDto *NumStrDto,
	isReversed bool,
	ePrefix string) (
	difference NumStrDto,
	err error) {

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
			n3IntAry,
			precision,
			newSignVal,
			ePrefix)

	return difference, err
}
