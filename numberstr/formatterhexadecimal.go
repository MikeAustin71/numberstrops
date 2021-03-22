package numberstr

import "sync"

type FormatterHexadecimal struct {
	numStrFmtType                 NumStrFormatTypeCode
	useLeadingPlus                bool
	useCapitalLetters             bool
	leftPrefix                    string // '0x'
	turnOnIntegerDigitsSeparation bool
	integerSeparators             NumStrIntSeparatorsDto
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
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
