package numberstr

import "sync"

type NumStrFmtSpecCurrencyValueDto struct {
	positiveValueFmt              string
	negativeValueFmt              string
	decimalDigits                 int
	currencyCode                  string
	currencyName                  string
	currencySymbol                rune
	turnOnIntegerDigitsSeparation bool
	numberSeparators              NumStrFmtSpecDigitsSeparatorsDto
	numFieldLen                   NumberFieldDto
	lock                          *sync.Mutex
}
