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

// New - Creates and returns a new instance of NumStrFmtSpecDto.
//
func (fmtSpecDto NumStrFmtSpecDto) New(
	idNo uint64,
	idString string,
	description string,
	tag string,
	countryCulture NumStrFmtSpecCountryDto,
	absoluteValue NumStrFmtSpecAbsoluteValueDto,
	currencyValue NumStrFmtSpecCurrencyValueDto,
	signedNumValue NumStrFmtSpecSignedNumValueDto,
	sciNotation NumStrFmtSpecSciNotationDto,
	ePrefix string) (
	NumStrFmtSpecDto,
	error) {

	if fmtSpecDto.lock == nil {
		fmtSpecDto.lock = new(sync.Mutex)
	}

	fmtSpecDto.lock.Lock()

	defer fmtSpecDto.lock.Unlock()

	ePrefix += "NumStrFmtSpecDto.New() "

	newFmtSpecDto := NumStrFmtSpecDto{}

	nStrFmtSpecDtoMech := numStrFmtSpecDtoMechanics{}

	err := nStrFmtSpecDtoMech.setNumStrFmtSpecDto(
		&newFmtSpecDto,
		idNo,
		idString,
		description,
		tag,
		countryCulture,
		absoluteValue,
		currencyValue,
		signedNumValue,
		sciNotation,
		ePrefix)

	return newFmtSpecDto, err
}
