package numberstr

import "sync"

type NumStrFmtSpecDto struct {
	idNo           uint64
	idString       string
	description    string
	tag            string
	countryCulture NumStrFmtSpecCountryDto
	absoluteValue  NumStrFmtSpecAbsoluteValueDto
	currencyValue  NumStrFmtSpecCurrencyValueDto
	signedNumValue NumStrFmtSpecSignedNumValueDto
	sciNotation    NumStrFmtSpecSciNotationDto
	lock           *sync.Mutex
}
