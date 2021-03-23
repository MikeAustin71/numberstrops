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
