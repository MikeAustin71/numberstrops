package numberstr

import (
	"errors"
	"fmt"
	"sync"
)

type numStrDtoQuark struct {
	lock *sync.Mutex
}

// copyIn - Receives an incoming NumStrDto object
// and copies the information to the current NumStrDto
// data fields.
//
// IMPORTANT
//
// This method will not test input parameters 'numStrDto'
// or nInDto for validity. Use numStrDtoElectron.copyIn()
// if validity tests are needed.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       change and overwrite the values of internal member variables to achieve the method's objectives.
//
//       This NumStrDto will receive all the data values contained
//       in input parameter 'nInDto'. When the copy operation is
//       completed, the data values in 'numStrDto' will be identical
//       to those contained in input parameter, 'nInDto'.
//
//
//  nInDto              *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       All data values in this NumStrDto instance will be copied to
//       input parameter 'numStrDto'.
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoQuark *numStrDtoQuark) copyInLowLevel(
	numStrDto *NumStrDto,
	nInDto *NumStrDto,
	ePrefix string) (
	err error) {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	ePrefix += "nStrDtoQuark.copyInLowLevel() "
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return err
	}

	if nInDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'nInDto' is INVALID!\n" +
			"nInDto = nil pointer!\n")
		return err
	}

	numStrDto.signVal = nInDto.signVal
	numStrDto.absAllNumRunes = nInDto.absAllNumRunes
	numStrDto.precision = nInDto.precision
	numStrDto.thousandsSeparator = nInDto.thousandsSeparator
	numStrDto.decimalSeparator = nInDto.decimalSeparator
	numStrDto.currencySymbol = nInDto.currencySymbol

	return err
}

// testNumStrDtoValidity - Evaluates the validity of a NumStrDto
// instance. If the instance is valid this method returns 'true'
// and sets the returned error type to nil.
//
// If the instance is invalid, the method returns 'false' plus an
// error message.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Member variables will be tested for
//       validity.
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
//  isValid             bool
//     - If this method completes successfully, this boolean value will
//       signal 'false' if 'numStrDto' member variables contains an
//       invalid value. Absent processing errors, if the validity tests
//       show that 'numStrDto' member variables values are valid, this
//       'isValid' flag will be set to 'true'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//       An error will also be returned if the method determines that
//       one or more of the 'numStrDto' member variables contain
//       invalid values.
//
func (nStrDtoQuark *numStrDtoQuark) testNumStrDtoValidity(
	numStrDto *NumStrDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	ePrefix += "numStrDtoQuark.testNumStrDtoValidity() "

	isValid = false
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return isValid, err
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if numStrDto.thousandsSeparator == 0 {
		numStrDto.thousandsSeparator = ','
	}

	if numStrDto.decimalSeparator == 0 {
		numStrDto.decimalSeparator = '.'
	}

	if numStrDto.currencySymbol == 0 {
		numStrDto.currencySymbol = '$'
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {
		err = errors.New(ePrefix + "\n" +
			"- Error: 'numStrDto' Number string is a ZERO length string!\n")
		return isValid, err
	}

	if int(numStrDto.precision) >= lenAbsAllNumRunes {
		err = errors.New(ePrefix + "\n" +
			"Error: precision does match number string. Type is Corrupted!")
		return isValid, err
	}

	if numStrDto.signVal != 1 && numStrDto.signVal != -1 {
		err = fmt.Errorf(ePrefix+"\n + "+
			"Sign Value is INVALID. Should be +1 or -1.\n"+
			"This Sign Value (numStrDto.signVal) is %v\n",
			numStrDto.signVal)

		return isValid, err
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if numStrDto.absAllNumRunes[i] < '0' || numStrDto.absAllNumRunes[i] > '9' {

			err = errors.New(ePrefix + "\n" +
				"Error: Non-Numeric character found in " +
				"'numStrDto' number string!\n")
			return isValid, err
		}
	}

	return isValid, err
}
