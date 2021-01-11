package numberstr

import (
	"fmt"
	"sync"
)

type numStrFormatterQuark struct {
	lock *sync.Mutex
}

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
		make([]int, len(nStrFormtr.integerDigitsGroupingSequence))

	for i := 0; i < len(nStrFormtr.integerDigitsGroupingSequence); i++ {
		newNStrFormtr.integerDigitsGroupingSequence[i] =
			nStrFormtr.integerDigitsGroupingSequence[i]
	}

	newNStrFormtr.integerDigitsSeparator = nStrFormtr.integerDigitsSeparator

	newNStrFormtr.turnOnIntegerDigitsSeparation =
		nStrFormtr.turnOnIntegerDigitsSeparation

	newNStrFormtr.numFieldDto = nStrFormtr.numFieldDto.CopyOut()

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
