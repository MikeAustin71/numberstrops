package numberstr

import (
	"sync"
)

// NumStrFmtSpecDto - Encapsulates all of the format specifications
// required to format the numeric values contained in type NumStrDto.
//
type NumStrFmtSpecDto struct {
	idNo           uint64
	idString       string
	description    string
	tag            string
	countryCulture FormatterCountry
	absoluteValue  FormatterAbsoluteValue
	currencyValue  FormatterCurrency
	signedNumValue FormatterSignedNumber
	sciNotation    FormatterSciNotation
	lock           *sync.Mutex
}
