package numberstr

import (
	"errors"
	"fmt"
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

// testCurrencySymbolDtoValidity - Receives an instance of
// CurrencySymbolDto and proceeds to test the validity of the
// member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (curSymMech *currencySymbolMechanics) testCurrencySymbolDtoValidity(
	currSymDto *CurrencySymbolDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if curSymMech.lock == nil {
		curSymMech.lock = new(sync.Mutex)
	}

	curSymMech.lock.Lock()

	defer curSymMech.lock.Unlock()

	ePrefix += "currencySymbolMechanics.testCurrencySymbolDtoValidity() "

	isValid = false
	err = nil

	if currSymDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'currSymDto' is a 'nil' pointer!\n")

		return isValid, err
	}

	if currSymDto.currencySymbol == 0 {
		isValid = false
		err = errors.New(ePrefix + "\n" +
			"Error: Currency Symbol rune is empty!\n")
		return isValid, err
	}

	if len(currSymDto.positiveValCurrFormat) == 0 {
		isValid = false
		err = errors.New(ePrefix + "\n" +
			"Error: Currency Symbol positive value format is empty!\n" +
			"positiveValCurrFormat==\"\"\n")
		return isValid, err
	}

	if len(currSymDto.negativeValCurrFormat) == 0 {
		isValid = false
		err = errors.New(ePrefix + "\n" +
			"Error: Currency Symbol negative value format is empty!\n" +
			"negativeValCurrFormat==\"\"\n")
		return isValid, err
	}

	nStrFmtQuark := numStrFormatQuark{}

	validFmtChars := nStrFmtQuark.getValidFormatChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(currSymDto.positiveValCurrFormat)
	var lenCurrFmt = len(runesToTest)
	var isRuneValid bool

	for i := 0; i < lenCurrFmt; i++ {

		isRuneValid = false

		for j := 0; j < lenValidFmtChars; j++ {
			if runesToTest[i] != validFmtChars[j] {
				continue
			} else {
				isRuneValid = true
				break
			}
		}

		if !isRuneValid {
			isValid = false
			err = fmt.Errorf(ePrefix+"\n"+
				"Error: Currency Symbol positive value contains an invalid character!\n"+
				"CurrencySymbolDto.positiveValCurrFormat invalid char == '%v' at Index [%v] \n",
				string(runesToTest[i]), i)
			return isValid, err
		}
	}

	runesToTest = []rune(currSymDto.negativeValCurrFormat)
	lenCurrFmt = len(runesToTest)

	for i := 0; i < lenCurrFmt; i++ {

		isRuneValid = false

		for j := 0; j < lenValidFmtChars; j++ {
			if runesToTest[i] != validFmtChars[j] {
				continue
			} else {
				isRuneValid = true
				break
			}
		}

		if !isRuneValid {
			isValid = false
			err = fmt.Errorf(ePrefix+"\n"+
				"Error: Currency Symbol negative value contains an invalid character!\n"+
				"CurrencySymbolDto.negativeValCurrFormat invalid char == '%v' at Index [%v] \n",
				string(runesToTest[i]), i)
			return isValid, err
		}
	}

	isValid = true
	return isValid, err
}
