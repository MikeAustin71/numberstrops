package numberstr

import "sync"

type NumStrFmtSpecHexadecimalDto struct {
	useLeadingPlus                bool
	useCapitalLetters             bool
	leftPrefix                    string // '0x'
	turnOnIntegerDigitsSeparation bool
	integerSeparators             []NumStrIntSeparator
	numFieldLenDto                NumberFieldDto
	lock                          *sync.Mutex
}
