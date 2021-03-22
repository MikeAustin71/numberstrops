package numberstr

import "sync"

type FormatterHexadecimal struct {
	numStrFmtType                 NumStrFormatTypeCode
	useUpperCaseLetters           bool
	leftPrefix                    string // '0x'
	turnOnIntegerDigitsSeparation bool
	integerSeparators             NumStrIntSeparatorsDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// FormatterHexadecimal instance  to the data fields
// of the current instance of FormatterHexadecimal
// instance.
//
// If input parameter 'incomingFormatterHex' is judged to be
// invalid, this method will return an error.
//
// IMPORTANT
//
// Be advised, all of the data fields in the current
// FormatterHexadecimal instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFormatterHex    *FormatterHexadecimal
//     - A pointer to an instance of FormatterHexadecimal.
//       The data values in this object will be copied to the
//       current FormatterHexadecimal instance.
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
func (formatterHex *FormatterHexadecimal) CopyIn(
	incomingFormatterHex *FormatterHexadecimal,
	ePrefix *ErrPrefixDto) error {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"CopyIn()")

	return formatterHexadecimalElectron{}.ptr().
		copyIn(
			formatterHex,
			incomingFormatterHex,
			ePrefix.XCtx(
				"formatterHex"))
}

// CopyOut - Creates and returns a deep copy of the current
// FormatterHexadecimal instance.
//
// If the current FormatterHexadecimal instance is judged
// to be invalid, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  FormatterHexadecimal
//     - If this method completes successfully, a new instance of
//       FormatterHexadecimal will be created and returned
//       containing all of the data values copied from the current
//       instance of FormatterHexadecimal.
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
func (formatterHex *FormatterHexadecimal) CopyOut(
	ePrefix *ErrPrefixDto) (
	FormatterHexadecimal,
	error) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"FormatterHexadecimal." +
			"CopyOut()")

	return formatterHexadecimalElectron{}.ptr().
		copyOut(
			formatterHex,
			ePrefix.XCtx(
				"formatterHex"))
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterHexadecimal instance to
// their initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (formatterHex *FormatterHexadecimal) Empty() {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	_ = formatterHexadecimalQuark{}.ptr().
		empty(formatterHex, nil)

	formatterHex.lock.Unlock()

	formatterHex.lock = nil
}

// GetLeftPrefix - Returns the 'Left Prefix' string configured for
// the current FormatterHexadecimal instance. Often, hexadecimal
// digits displayed in a text string are prefixed with the
// characters, "0x".
//   Example:  0x14AF
//
// The returned string controls this prefix.
//
func (formatterHex *FormatterHexadecimal) GetLeftPrefix() string {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.leftPrefix

}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for
// FormatterHexadecimal objects is
// NumStrFormatTypeCode(0).Hexadecimal().
//
// This method is required by interface INumStrFormatter.
//
func (formatterHex *FormatterHexadecimal) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.numStrFmtType
}

// GetUseUpperCaseLetters - Returns the 'Use Upper Case Letters'
// flag. This boolean flag determines whether Hexadecimal Digits
// 'A' through 'F' will be expressed as Upper Case Letters or Lower
// Case Letters ('a' through 'f').
//
func (formatterHex *FormatterHexadecimal) GetUseUpperCaseLetters() bool {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	return formatterHex.useUpperCaseLetters
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterHexadecimal instance to determine whether
// the current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - This returned boolean value will signal whether the current
//       FormatterHexadecimal is valid, or not. If the
//       current FormatterHexadecimal contains valid data,
//       this method returns 'true'. If the data is invalid, this
//       method will return 'false'.
//
func (formatterHex *FormatterHexadecimal) IsValidInstance() (
	isValid bool) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	isValid,
		_ = formatterHexadecimalQuark{}.ptr().
		testValidityOfFormatterHexadecimal(
			formatterHex,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterSignedNumber instance to determine
// whether the current instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//  error
//     - If the current instance of FormatterSignedNumber
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (formatterHex *FormatterHexadecimal) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("FormatterHexadecimal.IsValidInstanceError()",
		"Testing Validity of 'formatterHex'")

	_,
		err :=
		formatterHexadecimalQuark{}.ptr().
			testValidityOfFormatterHexadecimal(
				formatterHex,
				ePrefix)

	return err
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterHexadecimal object. For Signed Number
// formatters, the Number String Format Type Code is set to
// NumStrFormatTypeCode(0).Hexadecimal().
//
// This method is required by interface INumStrFormatter.
//
func (formatterHex *FormatterHexadecimal) SetNumStrFormatTypeCode() {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	formatterHex.numStrFmtType =
		NumStrFormatTypeCode(0).Hexadecimal()

	return
}

// SetUseUpperCaseLetters - Sets the 'Use Upper Case Letters' flag.
// This boolean flag determines whether Hexadecimal Digits 'A'
// through 'F' will be expressed as Upper Case Letters or Lower
// Case Letters ('a' through 'f').
//
// If input parameter 'useUpperCaseLetters' is set to 'true',
// hexadecimal alphabetic digits will be expressed as upper case
// letters. Conversely, if the parameter is 'false', hexadecimal
// alphabetic digits will be expressed as lower case letters.
//
func (formatterHex *FormatterHexadecimal) SetUseUpperCaseLetters(
	useUpperCaseLetters bool) {

	if formatterHex.lock == nil {
		formatterHex.lock = new(sync.Mutex)
	}

	formatterHex.lock.Lock()

	defer formatterHex.lock.Unlock()

	formatterHex.useUpperCaseLetters =
		useUpperCaseLetters

	return
}
