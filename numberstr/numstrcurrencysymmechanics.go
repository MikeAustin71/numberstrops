package numberstr

import (
	"errors"
	"sync"
)

type currencySymbolMechanics struct {
	lock *sync.Mutex
}

// copyIn - Copies the data values from input parameter
// 'currSymDtoIn' to 'currSymDtoOut'. All the data values
// in 'currSymDtoOut' will be overwritten.
//
func (curSymMech *currencySymbolMechanics) copyIn(
	currSymDtoOut *CurrencySymbolDto,
	currSymDtoIn *CurrencySymbolDto,
	ePrefix string) (err error) {

	if curSymMech.lock == nil {
		curSymMech.lock = new(sync.Mutex)
	}

	curSymMech.lock.Lock()

	defer curSymMech.lock.Unlock()

	ePrefix += "currencySymbolMechanics.copyIn() "

	if currSymDtoOut == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'currSymDtoOut' is a 'nil' pointer!\n")

		return err
	}

	if currSymDtoIn == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'currSymDtoIn' is a 'nil' pointer!\n")

		return err
	}

	currSymDtoOut.idNo = currSymDtoIn.idNo

	currSymDtoOut.currencySymbol =
		currSymDtoIn.currencySymbol

	currSymDtoOut.currencyName =
		currSymDtoIn.currencyName

	currSymDtoOut.countryName =
		currSymDtoIn.countryName

	currSymDtoOut.abbreviatedCountryName =
		currSymDtoIn.abbreviatedCountryName

	currSymDtoOut.alternateCountryName =
		currSymDtoIn.alternateCountryName

	currSymDtoOut.countryCodeTwoChar =
		currSymDtoIn.countryCodeTwoChar

	currSymDtoOut.countryCodeThreeChar =
		currSymDtoIn.countryCodeThreeChar

	currSymDtoOut.countryCodeNumber =
		currSymDtoIn.countryCodeNumber

	currSymDtoOut.currencyCode =
		currSymDtoIn.currencyCode

	currSymDtoOut.positiveValCurrFormat =
		currSymDtoIn.positiveValCurrFormat

	currSymDtoOut.negativeValCurrFormat =
		currSymDtoIn.negativeValCurrFormat

	return err
}

// copyOut - Creates and returns a deep copy of the
// input parameter 'currSymDto'.
//
func (curSymMech *currencySymbolMechanics) copyOut(
	currSymDto *CurrencySymbolDto,
	ePrefix string) (
	newCurrSymDto CurrencySymbolDto,
	err error) {

	if curSymMech.lock == nil {
		curSymMech.lock = new(sync.Mutex)
	}

	curSymMech.lock.Lock()

	defer curSymMech.lock.Unlock()

	ePrefix += "currencySymbolMechanics.copyOut() "

	if currSymDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'currSymDto' is a 'nil' pointer!\n")

		return newCurrSymDto, err
	}

	newCurrSymDto.idNo = currSymDto.idNo

	newCurrSymDto.currencySymbol =
		currSymDto.currencySymbol

	newCurrSymDto.currencyName =
		currSymDto.currencyName

	newCurrSymDto.countryName =
		currSymDto.countryName

	newCurrSymDto.abbreviatedCountryName =
		currSymDto.abbreviatedCountryName

	newCurrSymDto.alternateCountryName =
		currSymDto.alternateCountryName

	newCurrSymDto.countryCodeTwoChar =
		currSymDto.countryCodeTwoChar

	newCurrSymDto.countryCodeThreeChar =
		currSymDto.countryCodeThreeChar

	newCurrSymDto.countryCodeNumber =
		currSymDto.countryCodeNumber

	newCurrSymDto.currencyCode =
		currSymDto.currencyCode

	newCurrSymDto.positiveValCurrFormat =
		currSymDto.positiveValCurrFormat

	newCurrSymDto.negativeValCurrFormat =
		currSymDto.negativeValCurrFormat

	newCurrSymDto.lock = new(sync.Mutex)

	return newCurrSymDto, err
}
