package numberstr

import "sync"

type NumStrFmtSpecSignedNumValueDto struct {
	positiveValueFmt              string
	negativeValueFmt              string
	turnOnIntegerDigitsSeparation bool
	numberSeparators              NumStrFmtSpecDigitsSeparatorsDto
	numFieldLen                   NumberFieldDto
	lock                          *sync.Mutex
}
