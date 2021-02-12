package numberstr

import (
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
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (nStrDtoQuark *numStrDtoQuark) copyInLowLevel(
	numStrDto *NumStrDto,
	nInDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("nStrDtoQuark.copyInLowLevel()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if nInDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'nInDto' is INVALID!\n"+
			"nInDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	err =
		numStrDto.fmtSpec.CopyIn(
			&nInDto.fmtSpec,
			ePrefix)

	if err != nil {
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

// empty - Performs an 'empty' operation on an instance of NumStrDto.
// All member variable data values encapsulated by input parameter
// 'numStrDto' are set to their zero or default values.
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       set ALL the data values of internal member variables
//       to their zero or default values.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
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
func (nStrDtoQuark *numStrDtoQuark) empty(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoQuark.empty()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Perform an 'empty' operation
	numStrDto.signVal = 0
	numStrDto.absAllNumRunes = make([]rune, 0, 50)
	numStrDto.precision = 0
	numStrDto.fmtSpec = NumStrFmtSpecDto{}
	numStrDto.thousandsSeparator = ','
	numStrDto.decimalSeparator = '.'
	numStrDto.currencySymbol = '$'

	return err
}

// ptr - Returns a pointer to a new instance of numStrDtoQuark.
//
func (nStrDtoQuark numStrDtoQuark) ptr() *numStrDtoQuark {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	newNumStrDtoQuark := new(numStrDtoQuark)

	newNumStrDtoQuark.lock = new(sync.Mutex)

	return newNumStrDtoQuark
}

// setDefaultFormatSpec - Sets the format specification for a
// NumStrDto object to the default values. The default numeric
// format specification is United States.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       set the numeric format parameters to default values. Default
//       numeric format parameters are those of the United States.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
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
func (nStrDtoQuark *numStrDtoQuark) setDefaultFormatSpec(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoQuark.setDefaultFormatSpec()")

	if numStrDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrDto' is invalid!\n"+
			"'numStrDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = NumStrFormatCountry{}.Ptr().UnitedStates(
		&numStrDto.fmtSpec,
		ePrefix)

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
//       An error will also be returned if the method determines that
//       one or more of the 'numStrDto' member variables contain
//       invalid values.
//
func (nStrDtoQuark *numStrDtoQuark) testNumStrDtoValidity(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrDtoQuark.lock == nil {
		nStrDtoQuark.lock = new(sync.Mutex)
	}

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoQuark.testNumStrDtoValidity()")

	isValid = false
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	err = numStrDto.fmtSpec.IsValidInstanceError(
		ePrefix.XCtx(
			"Testing validity of " +
				"numStrDto.fmtSpec"))

	if err != nil {
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
		err = fmt.Errorf("%v\n"+
			"- Error: 'numStrDto' Number string is a ZERO length string!\n",
			ePrefix.String())

		return isValid, err
	}

	if int(numStrDto.precision) >= lenAbsAllNumRunes {

		err = fmt.Errorf("%v\n"+
			"Error: precision does match number string.\n"+
			"Thus 'NumStrDto' instance is Corrupted!\n",
			ePrefix.String())

		return isValid, err
	}

	if numStrDto.signVal != 1 && numStrDto.signVal != -1 {

		err = fmt.Errorf("%v\n"+
			"Sign Value is INVALID. Should be +1 or -1.\n"+
			"This Sign Value (numStrDto.signVal) is %v\n",
			ePrefix.String(),
			numStrDto.signVal)

		return isValid, err
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if numStrDto.absAllNumRunes[i] < '0' || numStrDto.absAllNumRunes[i] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: Non-Numeric character found in "+
				"'numStrDto' number string!\n",
				ePrefix.String())

			return isValid, err
		}
	}

	return isValid, err
}
