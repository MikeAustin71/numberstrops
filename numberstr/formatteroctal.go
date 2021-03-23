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
