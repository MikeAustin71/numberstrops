package numberstr

import "sync"

type NumStrFmtSpecCountryDto struct {
	idNo                   uint64
	idString               string
	description            string
	tag                    string
	countryCultureName     string
	abbreviatedCountryName string
	alternateCountryNames  []string
	countryCodeTwoChar     string
	countryCodeThreeChar   string
	countryCodeNumber      string

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecCountryDto instance  to the data fields
// of the current instance of NumStrFmtSpecCountryDto
// instance.
//
func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) CopyIn(
	inComingSpecCntryDto *NumStrFmtSpecCountryDto,
	ePrefix string) error {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFmtSpecCountryDto.CopyIn() "

	nStrFmtSpecCntryQuark :=
		numStrFmtSpecCountryDtoQuark{}

	return nStrFmtSpecCntryQuark.copyIn(
		nStrFmtSpecCntryDto,
		inComingSpecCntryDto,
		ePrefix)
}

func (nStrFmtSpecCntryDto *NumStrFmtSpecCountryDto) CopyOut() NumStrFmtSpecCountryDto {

	if nStrFmtSpecCntryDto.lock == nil {
		nStrFmtSpecCntryDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecCntryDto.lock.Lock()

	defer nStrFmtSpecCntryDto.lock.Unlock()

	nStrFmtSpecCntryQuark :=
		numStrFmtSpecCountryDtoQuark{}

	newFmtSpecCntryDto,
		_ :=
		nStrFmtSpecCntryQuark.copyOut(
			nStrFmtSpecCntryDto,
			"")

	return newFmtSpecCntryDto
}
