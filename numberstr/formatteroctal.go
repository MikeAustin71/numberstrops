package numberstr

import "sync"

// FormatterOctal - This structure stores formatting data
// for octal numeric values displayed in text number strings.
// Member data elements are listed as follows:
//
//
// Reference:
//
//   https://en.wikipedia.org/wiki/Octal
//
type FormatterOctal struct {
	numStrFmtType                 NumStrFormatTypeCode
	leftPrefix                    string // A prefix added to beginning of number string
	rightSuffix                   string // A suffix or postfix added to end of number string.
	turnOnIntegerDigitsSeparation bool
	integerSeparators             NumStrIntSeparatorsDto
	numFieldDto                   NumberFieldDto
	lock                          *sync.Mutex
}

// Empty - Deletes and resets the data values of all member
// variables within the current FormatterOctal instance to
// their initial 'zero' values.
//
// This method is required by interface INumStrFormatter.
//
func (fmtOctal *FormatterOctal) Empty() {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	_ = formatterOctalQuark{}.ptr().
		empty(fmtOctal, nil)

	fmtOctal.lock.Unlock()

	fmtOctal.lock = nil

	return
}

// GetNumStrFormatTypeCode - Returns the Number String Format Type
// Code. The Number String Format Type Code for FormatterOctal
// objects is NumStrFormatTypeCode(0).Octal().
//
// This method is required by interface INumStrFormatter.
//
func (fmtOctal *FormatterOctal) GetNumStrFormatTypeCode() NumStrFormatTypeCode {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	return fmtOctal.numStrFmtType

}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterOctal instance to determine whether that instance is
// valid in all respects.
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
//     - This returned boolean value will signal whether the
//       current FormatterOctal is valid, or not. If the current
//       FormatterOctal instance contains valid data, this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
func (fmtOctal *FormatterOctal) IsValidInstance() (
	isValid bool) {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	isValid,
		_ = formatterOctalQuark{}.ptr().
		testValidityOfFormatterOctal(
			fmtOctal,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current FormatterOctal instance to determine whether that
// instance is valid in all respects.
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
//     - If the current instance of FormatterOctal contains invalid
//       data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (fmtOctal *FormatterOctal) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPrefCtx("FormatterOctal."+
		"IsValidInstanceError()",
		"Testing Validity of 'formatterHex'")

	_,
		err :=
		formatterOctalQuark{}.ptr().
			testValidityOfFormatterOctal(
				fmtOctal,
				ePrefix)

	return err
}

// SetNumStrFormatTypeCode - Sets the Number String Format Type
// coded for this FormatterOctal object. For Octal Number
// formatters, the Number String Format Type Code is set to
// NumStrFormatTypeCode(0).Octal().
//
// This method is required by interface INumStrFormatter.
//
func (fmtOctal *FormatterOctal) SetNumStrFormatTypeCode() {

	if fmtOctal.lock == nil {
		fmtOctal.lock = new(sync.Mutex)
	}

	fmtOctal.lock.Lock()

	defer fmtOctal.lock.Unlock()

	fmtOctal.numStrFmtType =
		NumStrFormatTypeCode(0).Octal()

	return
}
