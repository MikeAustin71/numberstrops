package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFormatterCollectionGenerator struct {
	lock *sync.Mutex
}

func (nStrFormtrColGen *NumStrFormatterCollectionGenerator) Generate(
	formatterColSetup NumStrFormatterCollectionSetupDto) (
	numStrFormatterCollection map[NumStrValSpec]NumStrFormatter) {

	if nStrFormtrColGen.lock == nil {
		nStrFormtrColGen.lock = new(sync.Mutex)
	}

	nStrFormtrColGen.lock.Lock()

	defer nStrFormtrColGen.lock.Unlock()

	currencyFmt := NumStrFormatter{}

	currencyFmt.valueDisplaySpec =
		NumStrValSpec(0).CurrencyValue()

	currencyFmt.idNo = formatterColSetup.IdNo

	currencyFmt.idString = formatterColSetup.IdString

	currencyFmt.description =
		fmt.Sprintf("Currency Format: %v",
			formatterColSetup.CountryName)

	currencyFmt.countryName =
		formatterColSetup.CountryName

	currencyFmt.abbreviatedCountryName =
		formatterColSetup.AbbreviatedCountryName

	currencyFmt.alternateCountryName =
		formatterColSetup.AlternateCountryName

	currencyFmt.countryCodeTwoChar =
		formatterColSetup.CountryCodeTwoChar

	currencyFmt.countryCodeThreeChar =
		formatterColSetup.CountryCodeThreeChar

	currencyFmt.countryCodeNumber =
		formatterColSetup.CountryCodeNumber

	currencyFmt.positiveValueFmt =
		formatterColSetup.CurrencyPositiveValueFmt

	currencyFmt.negativeValueFmt =
		formatterColSetup.CurrencyNegativeValueFmt

	currencyFmt.decimalSeparator =
		formatterColSetup.DecimalSeparator

	currencyFmt.currencySymbol =
		formatterColSetup.CurrencySymbol

	currencyFmt.currencyDecimalDigits =
		formatterColSetup.CurrencyDecimalDigits

	currencyFmt.currencyCode =
		formatterColSetup.CurrencyCode

	currencyFmt.currencyName =
		formatterColSetup.CurrencyName

	currencyFmt.integerDigitsGroupingSequence =
		make([]uint, len(formatterColSetup.IntegerDigitsGroupingSequence))

	for i := 0; i < len(formatterColSetup.IntegerDigitsGroupingSequence); i++ {
		currencyFmt.integerDigitsGroupingSequence[i] =
			formatterColSetup.IntegerDigitsGroupingSequence[i]
	}

	currencyFmt.integerDigitsSeparator =
		formatterColSetup.IntegerDigitsSeparator

	currencyFmt.turnOnIntegerDigitsSeparation =
		formatterColSetup.CurrencyTurnOnIntegerDigitsSeparation

	currencyFmt.numFieldDto = numberFieldDto{
		requestedNumFieldLength: formatterColSetup.CurrencyNumFieldLen,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()

	absValFmt.idNo = formatterColSetup.IdNo

	absValFmt.idString = formatterColSetup.IdString

	absValFmt.description =
		fmt.Sprintf("Absolute Value Format: %v",
			formatterColSetup.CountryName)

	absValFmt.countryName =
		formatterColSetup.CountryName

	absValFmt.abbreviatedCountryName =
		formatterColSetup.AbbreviatedCountryName

	absValFmt.alternateCountryName =
		formatterColSetup.AlternateCountryName

	absValFmt.countryCodeTwoChar =
		formatterColSetup.CountryCodeTwoChar

	absValFmt.countryCodeThreeChar =
		formatterColSetup.CountryCodeThreeChar

	absValFmt.countryCodeNumber =
		formatterColSetup.CountryCodeNumber

	absValFmt.positiveValueFmt =
		formatterColSetup.AbsoluteValFmt

	absValFmt.negativeValueFmt =
		formatterColSetup.AbsoluteValFmt

	absValFmt.decimalSeparator =
		formatterColSetup.DecimalSeparator

	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""

	absValFmt.integerDigitsGroupingSequence =
		make([]uint, len(formatterColSetup.IntegerDigitsGroupingSequence))

	for i := 0; i < len(formatterColSetup.IntegerDigitsGroupingSequence); i++ {
		absValFmt.integerDigitsGroupingSequence[i] =
			formatterColSetup.IntegerDigitsGroupingSequence[i]
	}

	absValFmt.integerDigitsSeparator =
		formatterColSetup.IntegerDigitsSeparator

	absValFmt.turnOnIntegerDigitsSeparation =
		formatterColSetup.AbsoluteValTurnOnIntegerDigitsSeparation

	absValFmt.numFieldDto = numberFieldDto{
		requestedNumFieldLength: formatterColSetup.AbsoluteValNumFieldLen,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()

	signedNumValFmt.idNo = formatterColSetup.IdNo

	signedNumValFmt.idString = formatterColSetup.IdString

	signedNumValFmt.description =
		fmt.Sprintf("Signed Number Format: %v",
			formatterColSetup.CountryName)

	signedNumValFmt.countryName =
		formatterColSetup.CountryName

	signedNumValFmt.abbreviatedCountryName =
		formatterColSetup.AbbreviatedCountryName

	signedNumValFmt.alternateCountryName =
		formatterColSetup.AlternateCountryName

	signedNumValFmt.countryCodeTwoChar =
		formatterColSetup.CountryCodeTwoChar

	signedNumValFmt.countryCodeThreeChar =
		formatterColSetup.CountryCodeThreeChar

	signedNumValFmt.countryCodeNumber =
		formatterColSetup.CountryCodeNumber

	signedNumValFmt.positiveValueFmt =
		formatterColSetup.SignedNumValPositiveValueFmt

	signedNumValFmt.negativeValueFmt =
		formatterColSetup.SignedNumValNegativeValueFmt

	signedNumValFmt.decimalSeparator =
		formatterColSetup.DecimalSeparator

	signedNumValFmt.currencySymbol = 0

	signedNumValFmt.currencyDecimalDigits = -1

	signedNumValFmt.currencyCode = ""

	signedNumValFmt.currencyName = ""

	signedNumValFmt.integerDigitsGroupingSequence =
		make([]uint, len(formatterColSetup.IntegerDigitsGroupingSequence))

	for i := 0; i < len(formatterColSetup.IntegerDigitsGroupingSequence); i++ {
		signedNumValFmt.integerDigitsGroupingSequence[i] =
			formatterColSetup.IntegerDigitsGroupingSequence[i]
	}

	signedNumValFmt.integerDigitsSeparator =
		formatterColSetup.IntegerDigitsSeparator

	signedNumValFmt.turnOnIntegerDigitsSeparation =
		formatterColSetup.SignedNumValTurnOnIntegerDigitsSeparation

	signedNumValFmt.numFieldDto = numberFieldDto{
		requestedNumFieldLength: formatterColSetup.SignedNumValNumFieldLen,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	signedNumValFmt.lock = new(sync.Mutex)

	sciNotationFmt := NumStrFormatter{}

	sciNotationFmt.valueDisplaySpec =
		NumStrValSpec(0).ScientificNotation()

	sciNotationFmt.idNo = formatterColSetup.IdNo

	sciNotationFmt.idString = formatterColSetup.IdString

	sciNotationFmt.description =
		fmt.Sprintf("Scientific Notation Format: %v",
			formatterColSetup.CountryName)

	sciNotationFmt.countryName =
		formatterColSetup.CountryName

	sciNotationFmt.abbreviatedCountryName =
		formatterColSetup.AbbreviatedCountryName

	sciNotationFmt.alternateCountryName =
		formatterColSetup.AlternateCountryName

	sciNotationFmt.countryCodeTwoChar =
		formatterColSetup.CountryCodeTwoChar

	sciNotationFmt.countryCodeThreeChar =
		formatterColSetup.CountryCodeThreeChar

	sciNotationFmt.countryCodeNumber =
		formatterColSetup.CountryCodeNumber

	sciNotationFmt.decimalSeparator =
		formatterColSetup.DecimalSeparator

	sciNotationFmt.currencySymbol = 0
	sciNotationFmt.currencyDecimalDigits = -1
	sciNotationFmt.currencyCode = ""
	sciNotationFmt.currencyName = ""

	sciNotationFmt.integerDigitsGroupingSequence =
		make([]uint, 0)

	sciNotationFmt.integerDigitsSeparator = 0

	sciNotationFmt.turnOnIntegerDigitsSeparation = false

	sciNotationFmt.sciNotMantissaLength =
		formatterColSetup.SciNotMantissaLength

	sciNotationFmt.sciNotExponentChar =
		formatterColSetup.SciNotExponentChar

	sciNotationFmt.sciNotExponentUsesLeadingPlus =
		formatterColSetup.SciNotExponentUsesLeadingPlus

	sciNotationFmt.numFieldDto = numberFieldDto{
		requestedNumFieldLength: formatterColSetup.SciNotNumFieldLen,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	sciNotationFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt
	numStrFmtMap[NumStrValSpec(0).ScientificNotation()] = sciNotationFmt

	return numStrFmtMap
}
