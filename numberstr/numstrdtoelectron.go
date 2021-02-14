package numberstr

import (
	"fmt"
	"sync"
)

type numStrDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Receives pointers to two NumStrDto objects, 'numStrDto'
// and 'nInDto'. This method then proceeds to copy all member
// variable data fields from 'nInDto' to 'numStrDto'.
//
// If 'nInDto' is an invalid instance of NumStrDto, an error will
// be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE AND OVERWRITE all values of internal member
//       variables to achieve the method's objectives.
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
//       If this NumStrDto instance proves to be invalid, an error
//       will be returned.
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
func (nStrDtoElectron *numStrDtoElectron) copyIn(
	numStrDto *NumStrDto,
	nInDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.copyIn()")

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

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err = nStrDtoQuark.testNumStrDtoValidity(
		nInDto,
		ePrefix.XCtx(
			"Testing validity of 'nInDto'"))

	if err != nil {
		return err
	}

	err = nStrDtoQuark.copyInLowLevel(
		numStrDto,
		nInDto,
		ePrefix)

	_,
		err = nStrDtoQuark.testNumStrDtoValidity(
		numStrDto,
		ePrefix.XCtx("Final Validity Test for 'numStrDto'"))

	return err
}

// copyOut - Creates a deep copy of the data fields contained in
// input parameter 'numStrDto' and returns that data as a new
// instance NumStrDto.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
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
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDt
//     - If this method completes successfully, a deep copy of
//       input parameter 'numStrDto' will be returned.
//
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
func (nStrDtoElectron *numStrDtoElectron) copyOut(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.copyOut()")

	newNumStrDto = NumStrDto{}
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	nStrDtoQuark := numStrDtoQuark{}

	_,
		err = nStrDtoQuark.testNumStrDtoValidity(
		numStrDto,
		ePrefix.XCtx(
			"Testing validity of input parameter 'numStrDto'"))

	if err != nil {
		return newNumStrDto, err
	}

	err = nStrDtoQuark.copyInLowLevel(
		&newNumStrDto,
		numStrDto,
		ePrefix.XCtx(
			"numStrDto-> newNumStrDto"))

	if err != nil {
		return newNumStrDto, err
	}

	_,
		err = nStrDtoQuark.testNumStrDtoValidity(
		&newNumStrDto,
		ePrefix.XCtx(
			"Final validity test on 'newNumStrDto'"))

	return newNumStrDto, err
}

// getAbsFracRunes - Returns all of the fractional digits
// to the right of the decimal place in the NumStrDto input
// parameter, 'numStrDto'. The fractional digits will be returned
// as an array of runes. The rune array is not signed; that is,
// the rune array does not contain a '+' or '-' character in the
// first array position. The rune array is therefore said to
// represent the absolute value of the fractional digits in the
// 'numStrDto' numeric value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Any fractional digits contained in
//       the numeric value encapsulated by this NumStrDto instance
//       will be returned as an array of runes.
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
//  absFracRunes        []rune
//     - If this method completes successfully, an array of runes will
//       be returned. This rune array contains the fractional digits to
//       the right of the decimal point extracted from input parameter
//       numStrDto.
//
//       Note that the rune array is not signed; that is, the rune array
//       will not contain a '+' or '-' character in the first array
//       position. The rune array is therefore said to represent the
//       absolute value of the fractional digits in the input parameter
//       'numStrDto' numeric value.
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
func (nStrDtoElectron *numStrDtoElectron) getAbsFracRunes(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	absFracRunes []rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron.getAbsFracRunes()")

	absFracRunes = make([]rune, 0, 20)

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return absFracRunes, err
	}

	precision := int(numStrDto.precision)

	lenAllNums := len(numStrDto.absAllNumRunes)

	if lenAllNums == 0 ||
		precision < 0 ||
		precision > lenAllNums {
		return absFracRunes, err
	}

	absFracRunes = make([]rune, precision, precision+50)

	lenIntNums := lenAllNums - int(numStrDto.precision)

	for i := lenIntNums; i < lenAllNums; i++ {
		absFracRunes[i-lenIntNums] = numStrDto.absAllNumRunes[i]
	}

	return absFracRunes, err
}

// getAbsFracRunesLength - Returns the length of the
// fractional digits in the number string.
//
func (nStrDtoElectron *numStrDtoElectron) getAbsFracRunesLength(
	numStrDto *NumStrDto) (
	absFracRunesLen int) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	absFracRunesLen = 0

	if numStrDto == nil {
		return absFracRunesLen
	}

	return int(numStrDto.precision)
}

// GetAbsIntRunes - Returns all of the integer digits included
// in the current NumStrDto numeric value as an array of runes.
// The returned rune array does not contain a sign value in the
// first position and therefore represents the absolute or positive
// value of all the integer digits. The integer digits of a NumStrDto
// numeric includes all of the digits to the left of the decimal point.
//
// If the current NumStrDto consists of zero integers and fractional
// digits (Example: '0.123456'), this method will return a rune array
// consisting one array element with a '0' value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The integer digits extracted from this instance of NumStrDto
//       will be returned as an array of runes.
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
//  absIntRunes         []rune
//     - If this method completes successfully, this parameter will
//       return an array of runes containing positive integer digits
//       extracted from input parameter, 'numStrDto'.
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
func (nStrDtoElectron *numStrDtoElectron) getAbsIntRunes(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	absIntRunes []rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.getAbsIntRunes()")

	absIntRunes = []rune{}

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return absIntRunes, err
	}

	lenAllNum := len(numStrDto.absAllNumRunes)

	precision := int(numStrDto.precision)

	if lenAllNum == 0 ||
		precision < 0 ||
		precision >= lenAllNum {
		return absIntRunes, err
	}

	lenIntNum := lenAllNum - precision

	absIntRunes = make([]rune, lenIntNum, lenIntNum+50)

	for i := 0; i < lenIntNum; i++ {
		absIntRunes[i] = numStrDto.absAllNumRunes[i]
	}

	return absIntRunes, err
}

// getAbsIntRunesLength - Returns the length of the
// integer portion of the number string.
//
func (nStrDtoElectron *numStrDtoElectron) getAbsIntRunesLength(
	numStrDto *NumStrDto) (
	absIntRunesLen int) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	absIntRunesLen = 0

	if numStrDto == nil {
		return absIntRunesLen
	}

	absIntRunesLen =
		len(numStrDto.absAllNumRunes) -
			int(numStrDto.precision)

	return absIntRunesLen
}

// getAbsoluteValueNumStr - Returns all numeric digits residing in
// the NumStrDto instance passed as input parameter 'numStrDto'. This
// number string is a pure, unsigned number string consisting entirely
// of numeric digits. There is no leading number sign ('+' or '-') and
// this is no decimal delimiter ('.'). If fractional digits exists they
// are included and NOT separated by a decimal separator.
//
//     Examples:
//
//     numStrDto
//      value      result
//    --------     ------
//      123.45      12345
//     -123.45      12345
//
// Effectively, this method will return the member variable 'absAllNumRunes'
// from the input parameter, 'numStrDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The absolute value rune array from this instance will be
//       returned.
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
// absValNumStr         string
//     - A string of numeric digits representing the absolute numeric value
//       encapsulated in the NumStrDto instance passed as input parameter
//       'numStrDto'.
//
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
//
func (nStrDtoElectron *numStrDtoElectron) getAbsoluteValueNumStr(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	absValNumStr string,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.getAbsoluteValueNumStr()")

	absValNumStr = ""
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return absValNumStr, err
	}

	absValNumStr = string(numStrDto.absAllNumRunes)

	return absValNumStr, err
}

// getDecimalSeparator - returns the character designated
// as the decimal separator for the current NumStrDto instance.
//
// In the USA, the decimal separator is the period character ('.').
//
// Example:  123.456
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
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
//  decimalSeparator    rune
//     - If this method completes successfully, this rune value will
//       contain the decimal separator associated with input parameter
//       'numStrDto'.
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
func (nStrDtoElectron *numStrDtoElectron) getDecimalSeparator(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	decimalSeparator rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron.getDecimalSeparator()")

	decimalSeparator = 0
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return decimalSeparator, err
	}

	err = numStrDto.fmtSpec.currencyValue.IsValidInstanceError(
		ePrefix.XCtx(
			"Testing validity of currency value."))

	if err != nil {
		return decimalSeparator, err
	}

	decimalSeparator =
		numStrDto.fmtSpec.currencyValue.numberSeparatorsDto.GetDecimalSeparator()

	return decimalSeparator, err

	//if numStrDto.decimalSeparator == 0 {
	//	numStrDto.decimalSeparator = '.'
	//}
	//
	//decimalSeparator = numStrDto.decimalSeparator
	//
	//return decimalSeparator, err
}

// getCurrencySymbol - Returns the character currently designated
// as the currency symbol for this number string.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in:
//   datetime/numstrdtoconstants.go
//
// Example: $123.45
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
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
//  currencySymbol      rune
//     - If this method completes successfully, this rune value will
//       contain the currency symbol associated with input parameter
//       'numStrDto'.
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
func (nStrDtoElectron *numStrDtoElectron) getCurrencySymbol(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	currencySymbol rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron.getCurrencySymbol()")

	currencySymbol = 0
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return currencySymbol, err
	}
	err = numStrDto.fmtSpec.currencyValue.IsValidInstanceError(
		ePrefix.XCtx(
			"Testing validity of currency value."))

	currencySymbol =
		numStrDto.fmtSpec.currencyValue.GetCurrencySymbol()

	return currencySymbol, err

	//if numStrDto.currencySymbol == 0 {
	//	numStrDto.currencySymbol = '$'
	//}
	//
	//currencySymbol = numStrDto.currencySymbol
	//
	//return currencySymbol, err
}

// getNumericSignValue - Returns an integer value specifying the
// numeric sign value of the NumStrDto instance submitted as an
// input parameter ('numStrDto').
//
// The returned numeric sign value will always be one of two
// values:
//
//         +1 = a positive number greater than or equal to zero ('0')
//                       - OR -
//         -1 = a negative number less than zero ('0')
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The 'precision' value returned by this method will be extracted
//       from this NumStrDto instance.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  signValue           int
//     - If this method completes successfully, it will return the
//       numeric sign value associated with the numeric value
//       encapsulated by input parameter 'numStrDto'. The numeric
//       sign value associated with this NumStrDto instance will
//       always be one of two values:
//
//         +1 = a positive number greater than or equal to zero ('0')
//                       - OR -
//         -1 = a negative number less than zero ('0')
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
func (nStrDtoElectron *numStrDtoElectron) getNumericSignValue(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	signValue int,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron.getNumericSignValue()")

	signValue = 0
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return signValue, err
	}

	signValue = numStrDto.signVal

	if signValue != 1 &&
		signValue != -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Sign Value is invalid!\n"+
			"numStrDto.signVal= '%v'\n",
			ePrefix.String(),
			numStrDto.signVal)
	}

	return signValue, err
}

// getPrecisionUint - Returns the precision associated with the
// NumStrDto instance passed as an input parameter.
//
// 'precision', as defined here, specifies the number of numeric
// digits to the right of the decimal place. To compute the
// location of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count left
// 'precision' digits.
//
// The value of 'precision' returned by this method will always be
// greater than or equal to zero '0' ( precision >= zero ).
//
// Internally, type NumStrDto stores 'precision' as an unsigned
// integer value.
//
//
// Example:
//   1.234     getPrecisionUint() = 3
//   5         getPrecisionUint() = 0
//   0.12345   getPrecisionUint() = 5
//
//    Number                    Fractional
//    String     precision        Number
//    123456         3            123.456
//
// IMPORTANT
//
// This method does NOT validate the input parameter 'numStrDto'
// before returning the value. To run a validity check on the
// 'numStrDto' instance, first call the following method:
//
//  numStrDtoQuark.testNumStrDtoValidity() (bool, error)
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The 'precision' value returned by this method will be extracted
//       from this NumStrDto instance.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  precision           uint
//     - This method returns the 'precision' value associated with
//       the input parameter 'numStrDto' instance as a type uint.
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
func (nStrDtoElectron *numStrDtoElectron) getPrecisionUint(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	precision uint,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron.getPrecisionUint()")

	precision = 0
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return precision, err
	}

	precision = numStrDto.precision

	return precision, err
}

// getThousandsSeparator - returns a rune which represents
// the character currently used to separate thousands in
// the display of the current NumStrDto number string
// designated by input parameter 'numStrDto' .
//
// In the USA, the thousands separator is a comma character.
//
// Example: 1,000,000,000
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
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
//  integerDigitsSeparator  rune
//     - If this method completes successfully, this rune value will
//       contain the thousands separator associated with input parameter
//       'numStrDto'.
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
func (nStrDtoElectron *numStrDtoElectron) getThousandsSeparator(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	thousandsSeparator rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.getThousandsSeparator()")

	thousandsSeparator = 0
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return thousandsSeparator, err
	}

	thousandsSeparator =
		numStrDto.fmtSpec.currencyValue.
			numberSeparatorsDto.
			GetIntegerDigitsSeparator()

	return thousandsSeparator, err

	//if numStrDto.thousandsSeparator == 0 {
	//	numStrDto.thousandsSeparator = ','
	//}
	//
	//thousandsSeparator = numStrDto.thousandsSeparator
	//
	//return thousandsSeparator, err
}

// isNumStrZeroValue - Returns 'true' if all the digits in
// the number string for the passed NumStrDto instance are
// zero.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Separator values from this NumStrDto
//       instance will be used to populate the returned new instance of
//       NumStrDto.
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
//  isZeroValue         bool
//     - If this method completes successfully, this returned boolean
//       value will signal whether all the numeric digits in input
//       parameter 'numStrDto' are zero. If 'true', the numeric value
//       is zero and all digits in the number string are are zero.
//       If 'false', one or more digits are non-zero and the numeric
//       value is non-zero.
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
//
func (nStrDtoElectron *numStrDtoElectron) isNumStrZeroValue(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	isZeroValue bool,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.isNumStrZeroValue()")

	isZeroValue = false
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return isZeroValue, err
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {

		isZeroValue = false

	} else {

		isZeroValue = true

		for i := 0; i < lenAbsAllNumRunes; i++ {
			if numStrDto.absAllNumRunes[i] != '0' {
				isZeroValue = false
			}
		}

	}

	return isZeroValue, err
}

// isFractionalValue - Returns 'true' if the numeric value of
// input parameter 'numStrDto' (type NumStrDto) includes a
// fractional value; that is, the number has fractional digits
// to the right of the decimal point.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. The numeric value encapsulated by
//       this instance of NumStrDto will be evaluated to determine if
//       it contains fractional digits to the right of the decimal
//       point.
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
//  isFractionalValue   bool
//     - If this method completes successfully, this returned boolean
//       value will signal whether the numeric value encapsulated by
//       input parameter 'numStrDto' includes fractional digits to the
//       right of the decimal point. If 'true', 'numStrDto' has fractional
//       digits. If 'false', 'numStrDto' is an integer value with no
//       fractional digits to the right of the decimal point.
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
//
func (nStrDtoElectron *numStrDtoElectron) isFractionalValue(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	isFractionalValue bool,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix.SetEPref("numStrDtoElectron.isFractionalValue()")

	isFractionalValue = false
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return isFractionalValue, err
	}

	if numStrDto.precision > 0 {
		isFractionalValue = true
	} else {
		isFractionalValue = false
	}

	return isFractionalValue, err
}

// newBaseZeroNumStrDto - Returns a new NumStrDto object
// initialized to zero value. If the parameter 'precision'
// is set to a value greater than zero, then an equal number
// of zero digits will be added to the right of the decimal
// point.
//
// This method will return a zero NumStrDto instance set to
// default USA separators (Currency Symbol, Thousands Separator,
// Decimal Separator).
//
// Example Usage:
//
// Input Parameter      newNumStrDto
//    precision            Result
//       0                  "0"
//       2                  "0.00"
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  precision       uint
//     - The number of zeros which will be placed to the right of the
//       decimal point in the returned new instance of NumStrDto.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - This method will return a new instance of NumStrDto populated
//       with zero values. The separator values will be set to USA
//       defaults (Currency Symbol, Thousands Separator,
//       Decimal Separator).
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     Input Parameter   newNumStrDto
//       precision          Result
//          0                 "0"
//          2                 "0.00"
//
func (nStrDtoElectron *numStrDtoElectron) newBaseZeroNumStrDto(
	precision uint) (
	newNumStrDto NumStrDto) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix := ErrPrefixDto{}.Ptr()

	ePrefix.SetEPref(
		"numStrDtoElectron.newBaseZeroNumStrDto()")

	newNumStrDto = NumStrDto{}

	newNumStrDto.absAllNumRunes = make([]rune, 0, 50)

	// Set USA defaults for thousands separators,
	// decimal separators and currency Symbols

	_ = numStrDtoQuark{}.ptr().
		setDefaultFormatSpec(
			&newNumStrDto,
			ePrefix)

	newNumStrDto.signVal = 1
	newNumStrDto.signVal = 1
	newNumStrDto.precision = 0

	newNumStrDto.absAllNumRunes =
		append(newNumStrDto.absAllNumRunes, '0')

	if precision > 0 {

		for i := uint(0); i < precision; i++ {
			newNumStrDto.absAllNumRunes = append(newNumStrDto.absAllNumRunes, '0')
		}

		newNumStrDto.precision = precision
	}

	newNumStrDto.lock = new(sync.Mutex)

	return newNumStrDto
}

// newEmptyNumStrDto - Returns a new instance of NumStrDto. This
// instance is initialized to 'Empty' values. To fully understand
// 'Empty' values see method 'numStrDtoQuark.empty()'.
//
// 'Empty' values are useful because newNumStrDto.absAllNumRunes is
// initialized to make([]rune, 0, 50). This is useful for 'appending'
// runes to this array because one is sure to start out with zero
// pre-configured elements. This ensures that the first element
// appended to newNumStrDto.absAllNumRunes will actually be the first
// element in that array.
//
// If a new NumStrDto instance with zero values is required, see method
// 'numStrDtoElectron.newZeroNumStrDto()'.
//
func (nStrDtoElectron *numStrDtoElectron) newEmptyNumStrDto(
	ePrefix *ErrPrefixDto) (
	NumStrDto,
	error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.newEmptyNumStrDto()")

	nStrDtoQuark := numStrDtoQuark{}

	newNumStrDto := NumStrDto{}

	err := nStrDtoQuark.empty(
		&newNumStrDto,
		ePrefix)

	return newNumStrDto, err
}

// newZeroNumStrDto - returns a new NumStrDto initialized
// to zero value. If the parameter numFracDigits is set
// to a value greater than zero, then an equal number of
// zero characters will be added to the right of the
// decimal point.
//
// Examples:
//
// Input Parameter     newNumStrDto
//  numFracDigits        Result
//     0                  "0"
//     2                  "0.00"
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Separator values from this NumStrDto
//       instance will be used to populate the returned new instance of
//       NumStrDto.
//
//
//  numFracDigits       uint
//     - The number of zeros which will be placed to the right of the
//       decimal point in the returned new instance of NumStrDto.
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
//  newNumStrDto        NumStrDto
//     - If this method completes successfully, this returned new instance
//       of NumStrDto will be populated with zero values and will contain
//       the same separator values as those passed through input parameter
//       'newNumStrDto'.
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
//
func (nStrDtoElectron *numStrDtoElectron) newZeroNumStrDto(
	numStrDto *NumStrDto,
	numFracDigits uint,
	ePrefix *ErrPrefixDto) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.newZeroNumStrDto() ")

	newNumStrDto = NumStrDto{}
	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return newNumStrDto, err
	}

	newNumStrDto.signVal = 1

	_ = numStrDtoQuark{}.ptr().
		setDefaultFormatSpec(
			&newNumStrDto,
			ePrefix)

	newNumStrDto.signVal = 1
	newNumStrDto.precision = 0

	newNumStrDto.absAllNumRunes = make([]rune, 0, 50)

	newNumStrDto.absAllNumRunes =
		append(newNumStrDto.absAllNumRunes, '0')

	if numFracDigits > 0 {

		for i := uint(0); i < numFracDigits; i++ {
			newNumStrDto.absAllNumRunes = append(newNumStrDto.absAllNumRunes, '0')
		}

		newNumStrDto.precision = numFracDigits
	}

	newNumStrDto.lock = new(sync.Mutex)

	return newNumStrDto, err
}

// ptr - Returns a pointer to a new instance of
// numStrDtoElectron
//
func (nStrDtoElectron numStrDtoElectron) ptr() *numStrDtoElectron {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	newNumStrDtoElectron :=
		new(numStrDtoElectron)

	newNumStrDtoElectron.lock = new(sync.Mutex)

	return newNumStrDtoElectron
}

// setCurrencySymbol - assigns the input parameter rune as the
// currency symbol to be used by the NumStrDto instance passed
// as an input parameter.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// Note: If a zero value is submitted as input, Currency Symbol
// will default to the USA dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: datetime/numstrdtoconstants.go
//
// Example: $123.45
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and overwrite the value of internal member variable
//       numStrDto.currencySymbol.
//
//
//  currencySymbol      rune
//     - This rune or text character conveys the currency symbol
//       which will populate the internal member variable
//       'numStrDto.currencySymbol' for the input parameter,
//       'numStrDto'.
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
//
func (nStrDtoElectron *numStrDtoElectron) setCurrencySymbol(
	numStrDto *NumStrDto,
	currencySymbol rune,
	ePrefix *ErrPrefixDto) (err error) {

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.setCurrencySymbol()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	numStrDto.fmtSpec.
		currencyValue.
		currencySymbol = currencySymbol

	return err
}

// setDecimalSeparator - Assigns the input parameter rune as the
// decimal separator to be used by the NumStrDto instance passed
// as an input parameter.
//
// In the USA, the decimal separator is the decimal point or
// 'period' ('.').
//
// If a zero value is submitted as input, the default USA decimal
// separator, the decimal point ('.'), will be assigned.
//
// Example: 123.45
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and overwrite the value of internal member variable
//       'numStrDto.decimalSeparator'.
//
//
//  decimalSeparator    rune
//     - This rune or text character conveys the decimal separator
//       character which will populate the internal member variable
//       'numStrDto.decimalSeparator' for the input parameter,
//       'numStrDto'.
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
//
func (nStrDtoElectron *numStrDtoElectron) setDecimalSeparator(
	numStrDto *NumStrDto,
	decimalSeparator rune,
	ePrefix *ErrPrefixDto) (err error) {

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.setDecimalSeparator()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	numStrDto.fmtSpec.
		absoluteValue.
		numberSeparatorsDto.
		decimalSeparator = decimalSeparator

	numStrDto.fmtSpec.
		currencyValue.
		numberSeparatorsDto.
		decimalSeparator = decimalSeparator

	numStrDto.fmtSpec.
		signedNumValue.
		numberSeparatorsDto.
		decimalSeparator = decimalSeparator

	return err
}

// setThousandsSeparator - Sets the value of the character which
// will be used to separate thousands in the display of NumStrDto
// number strings. In the USA the thousands separator is the
// comma (',').
//
// If a zero value is submitted as input, the Thousands Separator
// will default to the USA Thousands Separator (',').
//
// Example: 1,000,000
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and overwrite the value of internal member variable
//       'numStrDto.decimalSeparator'.
//
//
//  integerDigitsSeparator  rune
//     - This rune or text character conveys the thousands separator
//       character which will populate the internal member variable
//       'numStrDto.integerDigitsSeparator' for the input parameter,
//       'numStrDto'.
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
//
func (nStrDtoElectron *numStrDtoElectron) setThousandsSeparator(
	numStrDto *NumStrDto,
	integerDigitsSeparator rune,
	ePrefix *ErrPrefixDto) (err error) {

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.setThousandsSeparator()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if integerDigitsSeparator == 0 {
		integerDigitsSeparator = ','
	}

	numStrDto.fmtSpec.
		absoluteValue.
		numberSeparatorsDto.
		integerDigitsSeparator = integerDigitsSeparator

	numStrDto.fmtSpec.
		absoluteValue.
		numberSeparatorsDto.
		integerDigitsGroupingSequence =
		[]uint{3}

	numStrDto.fmtSpec.
		currencyValue.
		numberSeparatorsDto.
		decimalSeparator = integerDigitsSeparator

	numStrDto.fmtSpec.
		currencyValue.
		numberSeparatorsDto.
		integerDigitsGroupingSequence =
		[]uint{3}

	numStrDto.fmtSpec.
		signedNumValue.
		numberSeparatorsDto.
		decimalSeparator = integerDigitsSeparator

	numStrDto.fmtSpec.
		signedNumValue.
		numberSeparatorsDto.
		integerDigitsGroupingSequence =
		[]uint{3}

	return err
}

// setNumericSeparatorsToDefaultIfEmpty - Sets the numeric
// separators in the passed instance of NumStrDto to their
// their USA defaults IF any of the separators are set to
// zero.
//
// Effectively, this method ensures that numeric separators
// in a NumStrDto instance are set to valid values.
//
// USA default numeric separators are listed as follows:
//
//       Currency Symbol: '$'
//   Thousands Separator: ','
//     Decimal Separator: '.'
//
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that numeric separators
// are set to valid values.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       ensure that separator values are set to valid values. If
//       any of the current separator values are zero, this method
//       method will set those values to USA defaults.
//
//       USA default numeric separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
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
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparatorsToDefaultIfEmpty(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron.setNumericSeparatorsToDefaultIfEmpty()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if !numStrDto.fmtSpec.IsValidInstance() {

		err = numStrDtoQuark{}.ptr().
			setDefaultFormatSpec(
				numStrDto,
				ePrefix)

	}

	return err
}

// setNumericSeparators - Receives three input parameters of type
// 'rune' which will be used to populate the numeric separators for
// the passed NumStrDto instance.
//
// If any of the runes, or text characters, submitted as input
// parameters have a zero value, that value will be defaulted to
// the USA standard numeric separator.
//
// The default USA Numeric Separators are listed as follows:
//
// Decimal Separator period ('.')     = 123.456
// Thousands Separator comma (',')    = 1,000,000,000
// Currency Symbol dollar sign ('$')  = $123
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE and overwrite the value of internal member variables
//       'numStrDto.decimalSeparator', 'numStrDto.integerDigitsSeparator'
//       and 'numStrDto.currencySymbol'.
//
//
//  decimalSeparator    rune
//     - This rune or text character conveys the decimal separator
//       character which will populate the internal member variable
//       'numStrDto.decimalSeparator' for the input parameter,
//       'numStrDto'.
//
//
//  integerDigitsSeparator  rune
//     - This rune or text character conveys the thousands separator
//       character which will populate the internal member variable
//       'numStrDto.integerDigitsSeparator' for the input parameter,
//       'numStrDto'.
//
//
//  currencySymbol    rune
//     - This rune or text character conveys the currency symbol
//       character which will populate the internal member variable
//       'numStrDto.currencySymbol' for the input parameter,
//       'numStrDto'.
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
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparators(
	numStrDto *NumStrDto,
	decimalSeparator rune,
	integerDigitsSeparator rune,
	currencySymbol rune,
	ePrefix *ErrPrefixDto) (err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron." +
			"setNumericSeparatorsToDefaultIfEmpty()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	if integerDigitsSeparator == 0 {
		integerDigitsSeparator = ','
	}

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	numStrDto.fmtSpec.
		absoluteValue.
		numberSeparatorsDto.
		decimalSeparator = decimalSeparator

	numStrDto.fmtSpec.
		currencyValue.
		numberSeparatorsDto.
		decimalSeparator = decimalSeparator

	numStrDto.fmtSpec.
		signedNumValue.
		numberSeparatorsDto.
		decimalSeparator = decimalSeparator

	numStrDto.fmtSpec.
		absoluteValue.
		numberSeparatorsDto.
		integerDigitsSeparator = integerDigitsSeparator

	numStrDto.fmtSpec.
		absoluteValue.
		numberSeparatorsDto.
		integerDigitsGroupingSequence =
		[]uint{3}

	numStrDto.fmtSpec.
		currencyValue.
		numberSeparatorsDto.
		decimalSeparator = integerDigitsSeparator

	numStrDto.fmtSpec.
		currencyValue.
		numberSeparatorsDto.
		integerDigitsGroupingSequence =
		[]uint{3}

	numStrDto.fmtSpec.
		signedNumValue.
		numberSeparatorsDto.
		decimalSeparator = integerDigitsSeparator

	numStrDto.fmtSpec.
		signedNumValue.
		numberSeparatorsDto.
		integerDigitsGroupingSequence =
		[]uint{3}

	numStrDto.fmtSpec.
		currencyValue.
		currencySymbol = currencySymbol

	return err
}

// setNumericSeparatorsDto - Sets the values of numeric separators:
//
//                           decimal point separator
//                           thousands separator
//                           currency symbol
//
// These separators are set based on values transmitted through input
// parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators'
// is set to zero or nil, an error will be returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       populate the separator values for Currency Symbol, Thousands
//       Separator and Decimal Separator with values copied from
//       the second input parameter 'customSeparators'.
//
//
//  customSeparators    NumericSeparatorsDto
//     - If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), this
//        method will return an error.
//
//       The separator values contained in this input parameter
//       will be copied to input parameter 'numStrDto'. The data
//       fields included in the NumericSeparatorsDto
//       are listed as follows:
//
//          type NumericSeparatorsDto struct {
//             decimalSeparator              rune
//             integerDigitsSeparator        rune
//             integerDigitsGroupingSequence []uint
//          }
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
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparatorsDto(
	numStrDto *NumStrDto,
	customSeparators NumericSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.setNumericSeparatorsDto()")

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	err = customSeparators.IsValidInstanceError(
		ePrefix.XCtx("Testing 'customSeparators' "))

	if err != nil {
		return err
	}

	err =
		numStrDto.fmtSpec.
			absoluteValue.
			numberSeparatorsDto.
			CopyIn(
				&customSeparators,
				ePrefix)

	if err != nil {
		return err
	}

	err =
		numStrDto.fmtSpec.
			currencyValue.
			numberSeparatorsDto.
			CopyIn(
				&customSeparators,
				ePrefix)

	err =
		numStrDto.fmtSpec.
			signedNumValue.
			numberSeparatorsDto.
			CopyIn(
				&customSeparators,
				ePrefix)

	return err
}

// setNumericSeparatorsToUSADefault - Sets the numeric
// separators in the passed instance of NumStrDto to the
// USA default values.
//
// USA default numeric separators are listed as follows:
//
//       Currency Symbol: '$'
//   Thousands Separator: ','
//     Decimal Separator: '.'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       set the numeric separator values contained in this instance
//       to USA default values.
//
//       USA default numeric separators are listed as follows:
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
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
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparatorsToUSADefault(
	numStrDto *NumStrDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrDtoElectron." +
			"setNumericSeparatorsToUSADefault()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	_ = numStrDtoQuark{}.ptr().
		setDefaultFormatSpec(
			numStrDto,
			ePrefix)

	return err
}

// SetSignValue - Sets the sign of the numeric value for the input parameter
// 'numStrDto'.
//
// Input parameter 'newSignVal' must be set to one of two possible values:
//
//  +1 for positive numeric values greater than or equal to zero
//
//  -1 for negative numeric values less than zero
//
// If 'newSignVal' is not set to one of the two valid numeric sign values
// listed above, this method will return an error.
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
//  newSignVal          int
//     - This input parameter must be set to one of two possible values:
//
//       +1 for positive numeric values greater than or equal to zero
//
//       -1 for negative numeric values less than zero
//
//       If 'newSignVal' is not set to one of the two valid numeric sign
//       values listed above, this method will return an error.
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
//
func (nStrDtoElectron *numStrDtoElectron) setSignValue(
	numStrDto *NumStrDto,
	newSignVal int,
	ePrefix *ErrPrefixDto) (
	err error) {
	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrDtoElectron.setSignValue()")

	err = nil

	if numStrDto == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'numStrDto' is INVALID!\n"+
			"numStrDto = nil pointer!\n",
			ePrefix.String())

		return err
	}

	if newSignVal != -1 && newSignVal != 1 {

		err = fmt.Errorf("%v\n"+
			"Invalid sign value passed. Sign value must be +1 or -1.\n"+
			"This sign value= %v\n",
			ePrefix.String(),
			newSignVal)

		return err
	}

	numStrDto.signVal = newSignVal

	return err
}
