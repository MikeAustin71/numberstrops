package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatterQuark struct {
	lock *sync.Mutex
}

// copyIn - Copies all data elements from 'nStrFormtrIncoming'
// to 'nStrFormtrTarget'.
func (nStrFormtrQuark *numStrFormatterQuark) copyIn(
	nStrFormtrTarget *NumStrFormatter,
	nStrFormtrIncoming *NumStrFormatter,
	ePrefix string) (
	err error) {

	ePrefix += "numStrFormatterQuark.copyOut() "

	if nStrFormtrTarget == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFormtrTarget' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	if nStrFormtrIncoming == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFormtrIncoming' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFormtrTarget.valueDisplaySpec =
		nStrFormtrIncoming.valueDisplaySpec

	nStrFormtrTarget.idNo =
		nStrFormtrIncoming.idNo

	nStrFormtrTarget.idString =
		nStrFormtrIncoming.idString

	nStrFormtrTarget.description =
		nStrFormtrIncoming.description

	nStrFormtrTarget.tag =
		nStrFormtrIncoming.tag

	nStrFormtrTarget.countryName =
		nStrFormtrIncoming.countryName

	nStrFormtrTarget.abbreviatedCountryName =
		nStrFormtrIncoming.abbreviatedCountryName

	nStrFormtrTarget.alternateCountryName =
		nStrFormtrIncoming.alternateCountryName

	nStrFormtrTarget.countryCodeTwoChar =
		nStrFormtrIncoming.countryCodeTwoChar

	nStrFormtrTarget.countryCodeThreeChar =
		nStrFormtrIncoming.countryCodeThreeChar

	nStrFormtrTarget.countryCodeNumber =
		nStrFormtrIncoming.countryCodeNumber

	nStrFormtrTarget.positiveValueFmt =
		nStrFormtrIncoming.positiveValueFmt

	nStrFormtrTarget.negativeValueFmt =
		nStrFormtrIncoming.negativeValueFmt

	nStrFormtrTarget.decimalSeparator =
		nStrFormtrIncoming.decimalSeparator

	nStrFormtrTarget.currencySymbol =
		nStrFormtrIncoming.currencySymbol

	nStrFormtrTarget.currencyDecimalDigits =
		nStrFormtrIncoming.currencyDecimalDigits

	nStrFormtrTarget.currencyCode =
		nStrFormtrIncoming.currencyCode

	nStrFormtrTarget.currencyName =
		nStrFormtrIncoming.currencyName

	nStrFormtrTarget.integerDigitsGroupingSequence =
		make([]uint, 0)

	nStrFormtrTarget.integerDigitsGroupingSequence =
		make([]uint, len(nStrFormtrIncoming.integerDigitsGroupingSequence))

	for i := 0; i < len(nStrFormtrIncoming.integerDigitsGroupingSequence); i++ {
		nStrFormtrTarget.integerDigitsGroupingSequence[i] =
			nStrFormtrIncoming.integerDigitsGroupingSequence[i]
	}

	nStrFormtrTarget.integerDigitsSeparator = nStrFormtrIncoming.integerDigitsSeparator

	nStrFormtrTarget.turnOnIntegerDigitsSeparation =
		nStrFormtrIncoming.turnOnIntegerDigitsSeparation

	nStrFormtrTarget.numFieldDto.requestedNumFieldLength =
		nStrFormtrIncoming.numFieldDto.requestedNumFieldLength

	nStrFormtrTarget.numFieldDto.actualNumFieldLength =
		nStrFormtrIncoming.numFieldDto.actualNumFieldLength

	nStrFormtrTarget.numFieldDto.minimumNumFieldLength =
		nStrFormtrIncoming.numFieldDto.minimumNumFieldLength

	return err
}

// copyOut - Returns a deep copy of the current NumStrFormatter
// instance.
//
func (nStrFormtrQuark *numStrFormatterQuark) copyOut(
	nStrFormtr *NumStrFormatter,
	ePrefix string) (
	newNStrFormtr NumStrFormatter,
	err error) {

	ePrefix += "numStrFormatterQuark.copyOut() "

	newNStrFormtr = NumStrFormatter{}

	if nStrFormtr == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter nStrFormtr is a 'nil' pointer!\n",
			ePrefix)
		return newNStrFormtr, err
	}

	newNStrFormtr.valueDisplaySpec =
		nStrFormtr.valueDisplaySpec

	newNStrFormtr.idNo =
		nStrFormtr.idNo

	newNStrFormtr.idString =
		nStrFormtr.idString

	newNStrFormtr.description =
		nStrFormtr.description

	newNStrFormtr.tag =
		nStrFormtr.tag

	newNStrFormtr.countryName =
		nStrFormtr.countryName

	newNStrFormtr.abbreviatedCountryName =
		nStrFormtr.abbreviatedCountryName

	newNStrFormtr.alternateCountryName =
		nStrFormtr.alternateCountryName

	newNStrFormtr.countryCodeTwoChar =
		nStrFormtr.countryCodeTwoChar

	newNStrFormtr.countryCodeThreeChar =
		nStrFormtr.countryCodeThreeChar

	newNStrFormtr.countryCodeNumber =
		nStrFormtr.countryCodeNumber

	newNStrFormtr.positiveValueFmt =
		nStrFormtr.positiveValueFmt

	newNStrFormtr.negativeValueFmt =
		nStrFormtr.negativeValueFmt

	newNStrFormtr.decimalSeparator =
		nStrFormtr.decimalSeparator

	newNStrFormtr.currencySymbol =
		nStrFormtr.currencySymbol

	newNStrFormtr.currencyDecimalDigits =
		nStrFormtr.currencyDecimalDigits

	newNStrFormtr.currencyCode =
		nStrFormtr.currencyCode

	newNStrFormtr.currencyName =
		nStrFormtr.currencyName

	newNStrFormtr.integerDigitsGroupingSequence =
		make([]uint, len(nStrFormtr.integerDigitsGroupingSequence))

	for i := 0; i < len(nStrFormtr.integerDigitsGroupingSequence); i++ {
		newNStrFormtr.integerDigitsGroupingSequence[i] =
			nStrFormtr.integerDigitsGroupingSequence[i]
	}

	newNStrFormtr.integerDigitsSeparator = nStrFormtr.integerDigitsSeparator

	newNStrFormtr.turnOnIntegerDigitsSeparation =
		nStrFormtr.turnOnIntegerDigitsSeparation

	newNStrFormtr.numFieldDto.requestedNumFieldLength =
		nStrFormtr.numFieldDto.requestedNumFieldLength

	newNStrFormtr.numFieldDto.actualNumFieldLength =
		nStrFormtr.numFieldDto.actualNumFieldLength

	newNStrFormtr.numFieldDto.minimumNumFieldLength =
		nStrFormtr.numFieldDto.minimumNumFieldLength

	newNStrFormtr.lock = new(sync.Mutex)

	return newNStrFormtr, err
}

func (nStrFormtrQuark *numStrFormatterQuark) getValidCurrencyPositiveValFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune(" +$127.54NUMFIELD")
}

func (nStrFormtrQuark *numStrFormatterQuark) getValidCurrencyNegativeValFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune("(-) $127.54NUMFIELD")
}

func (nStrFormtrQuark *numStrFormatterQuark) getValidAbsolutePositiveValFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune(" +127.54NUMFIELD")
}

func (nStrFormtrQuark *numStrFormatterQuark) getValidAbsoluteNegativeValFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune(" +127.54NUMFIELD")
}

func (nStrFormtrQuark *numStrFormatterQuark) getValidSignedNumPositiveValFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune(" +127.54NUMFIELD")
}

func (nStrFormtrQuark *numStrFormatterQuark) getValidSignedNumNegativeValFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune("(-) 127.54NUMFIELD")
}

func (nStrFormtrQuark *numStrFormatterQuark) getAllValidFmtChars() []rune {

	if nStrFormtrQuark.lock == nil {
		nStrFormtrQuark.lock = new(sync.Mutex)
	}

	nStrFormtrQuark.lock.Lock()

	defer nStrFormtrQuark.lock.Unlock()

	return []rune("(-) +$127.54NUMFIELD")
}
