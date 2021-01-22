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

	newFmtSpecDto.idNo = idNo
	newFmtSpecDto.idString = idString
	newFmtSpecDto.description = description
	newFmtSpecDto.tag = tag

	err := newFmtSpecDto.countryCulture.CopyIn(
		&countryCulture,
		ePrefix)

	if err != nil {
		return NumStrFmtSpecDto{}, err
	}

	err = newFmtSpecDto.absoluteValue.CopyIn(
		&absoluteValue,
		ePrefix)

	if err != nil {
		return NumStrFmtSpecDto{}, err
	}

	err = newFmtSpecDto.currencyValue.CopyIn(
		&currencyValue,
		ePrefix)

	if err != nil {
		return NumStrFmtSpecDto{}, err
	}

	err = newFmtSpecDto.signedNumValue.CopyIn(
		&signedNumValue,
		ePrefix)

	if err != nil {
		return NumStrFmtSpecDto{}, err
	}

	err = newFmtSpecDto.sciNotation.CopyIn(
		&sciNotation,
		ePrefix)

	return newFmtSpecDto, err
}
