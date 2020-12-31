package numberstr

import (
	"sync"
)

type numStrDtoUtility struct {
	lock *sync.Mutex
}

func (nStrDtoUtil *numStrDtoUtility) addNumStrs(
	addend1 *NumStrDto,
	addend2 *NumStrDto,
	ePrefix string) (
	sum NumStrDto,
	err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	ePrefix += "numStrDtoUtility.addNumStrs() "

	nStrDtoElectron := numStrDtoElectron{}

	sum = nStrDtoElectron.newBaseZeroNumStrDto(0)
	err = nil

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			addend1,
			ePrefix+"Initial validity test for 'addend1' ")

	if err != nil {
		return sum, err
	}

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			addend2,
			ePrefix+"Initial validity test for 'addend2' ")

	if err != nil {
		return sum, err
	}

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
			ePrefix+"n1DtoSetup Sign=1 ")

		if err != nil {
			return sum, err
		}

		err = nStrDtoElectron.setSignValue(
			&n2DtoSetup,
			1,
			ePrefix+"n1DtoSetup Sign=1 ")

		if err != nil {
			return sum, err
		}

		if compare == 0 {
			// Numeric Values are Equal!
			sum = nStrDtoElectron.newBaseZeroNumStrDto(
				n1DtoSetup.precision)
			return sum, err
		}

		// Sign Values are NOW Equal!

		sum,
			err = nStrDtoHelper.signValuesAreEqualSubtractNumStrs(
			&n1DtoSetup,
			&n2DtoSetup,
			isOrderReversed,
			ePrefix+"n1DtoSetup, n2DtoSetup ")

		var isZeroValue bool

		isZeroValue,
			err = nStrDtoElectron.isNumStrZeroValue(
			&sum,
			ePrefix+"sum ")

		if err != nil {
			return sum, err
		}

		if isZeroValue {
			newSignVal = 1
		}

		err = nStrDtoElectron.setSignValue(
			&sum,
			newSignVal,
			ePrefix+"sum ")

		return sum, err
	}

	// Sign Values ARE Already Equal!

	sum,
		err = nStrDtoHelper.signValuesAreEqualAddNumStrs(
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
//  numStrDto        *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and overwrite the internal member variable data
//       values in order to achieve the method's objectives.
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
//  err                error
//     - If this method completes successfully, the returned error Type is
//       set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. The 'ePrefix' text will be
//       prefixed to the beginning of the returned error message.
//
func (nStrDtoUtil *numStrDtoUtility) multiplyInPlace(
	numStrDto *NumStrDto,
	multiplier *NumStrDto,
	ePrefix string) (
	err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	ePrefix += "numStrDtoUtility.multiplyInPlace() "

	nStrDtoElectron := numStrDtoElectron{}

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			numStrDto,
			ePrefix+"Initial validity test for 'numStrDto' ")

	if err != nil {
		return err
	}

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			multiplier,
			ePrefix+"Initial validity test for 'multiplier' ")

	if err != nil {
		return err
	}

	var productNDto NumStrDto

	nStrDtoHelper := numStrDtoHelper{}

	productNDto,
		err = nStrDtoHelper.multiplyNumStrs(
		numStrDto,
		multiplier,
		ePrefix+"numStrDto x multiplier ")

	if err != nil {
		return err
	}

	err = nStrDtoElectron.copyIn(
		numStrDto,
		&productNDto,
		ePrefix+"productNDto->numStrDto ")

	return err
}

// setNumStr - Sets the value of the current NumStrDto instance
// to the number string received as input.
func (nStrDtoUtil *numStrDtoUtility) setNumStr(
	numStrDto *NumStrDto,
	numSepsDto NumericSeparatorDto,
	numStr string,
	ePrefix string) (err error) {

	if nStrDtoUtil.lock == nil {
		nStrDtoUtil.lock = new(sync.Mutex)
	}

	nStrDtoUtil.lock.Lock()

	defer nStrDtoUtil.lock.Unlock()

	ePrefix += "numStrDtoUtility.setNumStr() "

	var newNumStrDto NumStrDto
	err = nil

	nStrDtoNanobot := numStrDtoNanobot{}

	newNumStrDto,
		err = nStrDtoNanobot.newNumStr(
		numSepsDto,
		numStr,
		ePrefix)

	if err != nil {
		return err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&newNumStrDto,
		numSepsDto,
		ePrefix)

	if err != nil {
		return err
	}

	err =
		nStrDtoElectron.copyIn(
			numStrDto,
			&newNumStrDto,
			ePrefix+"newNumStrDto->numStrDto ")

	if err != nil {
		return err
	}

	_,
		err =
		nStrDtoElectron.testNumStrDtoValidity(
			numStrDto,
			ePrefix+"Final validity test for 'numStrDto' ")

	return err
}
