package numberstr

import (
	"strings"
	"sync"
)

// NumStrFormatCountry - Returns the number string formats used
// by specific countries.
//
// Sources:
//  https://gist.github.com/bzerangue/5484121
//  http://symbologic.info/currency.htm
//  http://www.xe.com/symbols.php
//  https://www.countrycode.org/
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//  https://www.codeproject.com/articles/78175/international-number-formats
//  https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//  https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://english.stackexchange.com/questions/124797/how-to-write-negative-currency-in-text
//  https://freeformatter.com/i18n-standards-code-snippets.html
//  https://www.evertype.com/standards/euro/formats.html
//
//
// Countries:
//
//  Argentina
//  Australia
//  Canada
//  China
//  France
//  Germany
//  Israel
//  Italy
//  United Kingdom
//  United States

type NumStrFormatCountry struct {
	lock *sync.Mutex
}

func (nStrFmtCountry NumStrFormatCountry) NewPtr() *NumStrFormatCountry {
	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	return &NumStrFormatCountry{
		lock: new(sync.Mutex),
	}
}

// Argentina - Returns the number string format used in the
// Argentina.
func (nStrFmtCountry *NumStrFormatCountry) Argentina() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 32
	currencyFmt.idString = "032"
	currencyFmt.description = "Currency Format Argentina"
	currencyFmt.countryName = "Argentina"
	currencyFmt.abbreviatedCountryName = "Argentina"
	currencyFmt.alternateCountryName = "Argentine Republic"
	currencyFmt.countryCodeTwoChar = "AR"
	currencyFmt.countryCodeThreeChar = "ARG"
	currencyFmt.countryCodeNumber = "032"
	currencyFmt.positiveValueFmt = "$ 127.54"
	currencyFmt.negativeValueFmt = "-$ 127.54"
	currencyFmt.decimalSeparator = ','
	currencyFmt.currencySymbol = '\U00000024'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "ARS"
	currencyFmt.currencyName = "Peso"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = '.'
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 32
	absValFmt.idString = "032"
	absValFmt.description = "Absolute Value Format Argentina"
	absValFmt.countryName = "Argentina"
	absValFmt.abbreviatedCountryName = "Argentina"
	absValFmt.alternateCountryName = "Argentine Republic"
	absValFmt.countryCodeTwoChar = "AR"
	absValFmt.countryCodeThreeChar = "ARG"
	absValFmt.countryCodeNumber = "032"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = ','
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = '.'
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 32
	absValFmt.idString = "032"
	signedNumValFmt.description = "Signed Number Format Argentina"
	signedNumValFmt.countryName = "Argentina"
	signedNumValFmt.abbreviatedCountryName = "Argentina"
	signedNumValFmt.alternateCountryName = "Argentine Republic"
	signedNumValFmt.countryCodeTwoChar = "AR"
	signedNumValFmt.countryCodeThreeChar = "ARG"
	signedNumValFmt.countryCodeNumber = "032"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "- 127.54"
	signedNumValFmt.decimalSeparator = ','
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = '.'
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// Australia - Returns the number string format used in
// Australia.
func (nStrFmtCountry *NumStrFormatCountry) Australia() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 36
	currencyFmt.idString = "036"
	currencyFmt.description = "Currency Format: Australia"
	currencyFmt.countryName = "Australia"
	currencyFmt.abbreviatedCountryName = "Australia"
	currencyFmt.alternateCountryName = "Commonwealth of Australia"
	currencyFmt.countryCodeTwoChar = "AU"
	currencyFmt.countryCodeThreeChar = "AUS"
	currencyFmt.countryCodeNumber = "036"
	currencyFmt.positiveValueFmt = "$ 127.54"
	currencyFmt.negativeValueFmt = "-$ 127.54"
	currencyFmt.decimalSeparator = '.'
	currencyFmt.currencySymbol = '\U00000024'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "AUD"
	currencyFmt.currencyName = "Dollar"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = ','
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 36
	absValFmt.idString = "036"
	absValFmt.description = "Absolute Value Format: Australia"
	absValFmt.countryName = "Australia"
	absValFmt.abbreviatedCountryName = "Australia"
	absValFmt.alternateCountryName = "Commonwealth of Australia"
	absValFmt.countryCodeTwoChar = "AU"
	absValFmt.countryCodeThreeChar = "AUS"
	absValFmt.countryCodeNumber = "036"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = '.'
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = ','
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 36
	absValFmt.idString = "036"
	signedNumValFmt.description = "Signed Number Format: Australia"
	signedNumValFmt.countryName = "Australia"
	signedNumValFmt.abbreviatedCountryName = "Australia"
	signedNumValFmt.alternateCountryName = "Commonwealth of Australia"
	signedNumValFmt.countryCodeTwoChar = "AU"
	signedNumValFmt.countryCodeThreeChar = "AUS"
	signedNumValFmt.countryCodeNumber = "036"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "-127.54"
	signedNumValFmt.decimalSeparator = '.'
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = ','
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// Canada - Returns the number string format used in
// Canada.
func (nStrFmtCountry *NumStrFormatCountry) Canada() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 124
	currencyFmt.idString = "124"
	currencyFmt.description = "Currency Format: Canada"
	currencyFmt.countryName = "Canada"
	currencyFmt.abbreviatedCountryName = "Canada"
	currencyFmt.alternateCountryName = "Canada"
	currencyFmt.countryCodeTwoChar = "CA"
	currencyFmt.countryCodeThreeChar = "CAN"
	currencyFmt.countryCodeNumber = "124"
	currencyFmt.positiveValueFmt = "$ 127.54"
	currencyFmt.negativeValueFmt = "-$ 127.54"
	currencyFmt.decimalSeparator = '.'
	currencyFmt.currencySymbol = '\U00000024'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "CAD"
	currencyFmt.currencyName = "Dollar"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = ','
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 124
	absValFmt.idString = "124"
	absValFmt.description = "Absolute Value Format: Canada"
	absValFmt.countryName = "Canada"
	absValFmt.abbreviatedCountryName = "Canada"
	absValFmt.alternateCountryName = "Canada"
	absValFmt.countryCodeTwoChar = "CA"
	absValFmt.countryCodeThreeChar = "CAN"
	absValFmt.countryCodeNumber = "124"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = '.'
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = ','
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 124
	absValFmt.idString = "124"
	signedNumValFmt.description = "Signed Number Format: Canada"
	signedNumValFmt.countryName = "Canada"
	signedNumValFmt.abbreviatedCountryName = "Canada"
	signedNumValFmt.alternateCountryName = "Canada"
	signedNumValFmt.countryCodeTwoChar = "CA"
	signedNumValFmt.countryCodeThreeChar = "CAN"
	signedNumValFmt.countryCodeNumber = "124"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "-127.54"
	signedNumValFmt.decimalSeparator = '.'
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = ','
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// CanadaFrench - Returns the number string format used in
// French Canada.
func (nStrFmtCountry *NumStrFormatCountry) CanadaFrench() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 124
	currencyFmt.idString = "124"
	currencyFmt.description = "Currency Format: Canada French"
	currencyFmt.countryName = "Canada French"
	currencyFmt.abbreviatedCountryName = "Canada French"
	currencyFmt.alternateCountryName = "French Canadian"
	currencyFmt.countryCodeTwoChar = "CA"
	currencyFmt.countryCodeThreeChar = "CAN"
	currencyFmt.countryCodeNumber = "124"
	currencyFmt.positiveValueFmt = "127.54 $"
	currencyFmt.negativeValueFmt = "127.54 $-"
	currencyFmt.decimalSeparator = ','
	currencyFmt.currencySymbol = '\U00000024'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "CAD"
	currencyFmt.currencyName = "Dollar"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = '.'
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 124
	absValFmt.idString = "124"
	absValFmt.description = "Absolute Value Format: Canada French"
	absValFmt.countryName = "Canada French"
	absValFmt.abbreviatedCountryName = "Canada French"
	absValFmt.alternateCountryName = "French Canadian"
	absValFmt.countryCodeTwoChar = "CA"
	absValFmt.countryCodeThreeChar = "CAN"
	absValFmt.countryCodeNumber = "124"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = ','
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = '.'
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 124
	absValFmt.idString = "124"
	signedNumValFmt.description = "Signed Number Format: Canada French"
	signedNumValFmt.countryName = "Canada French"
	signedNumValFmt.abbreviatedCountryName = "Canada French"
	signedNumValFmt.alternateCountryName = "French Canadian"
	signedNumValFmt.countryCodeTwoChar = "CA"
	signedNumValFmt.countryCodeThreeChar = "CAN"
	signedNumValFmt.countryCodeNumber = "124"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "127.54 -"
	signedNumValFmt.decimalSeparator = ','
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = '.'
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// China - Returns the number string format used in the
// Peoples Republic of China.
func (nStrFmtCountry *NumStrFormatCountry) China() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 156
	currencyFmt.idString = "156"
	currencyFmt.description = "Currency Format: China"
	currencyFmt.countryName = "China"
	currencyFmt.abbreviatedCountryName = "CHN"
	currencyFmt.alternateCountryName = "Peoples Republic of China"
	currencyFmt.countryCodeTwoChar = "CN"
	currencyFmt.countryCodeThreeChar = "CHN"
	currencyFmt.countryCodeNumber = "156"
	currencyFmt.positiveValueFmt = "$ 127.54"
	currencyFmt.negativeValueFmt = "-$ 127.54"
	currencyFmt.decimalSeparator = '.'
	currencyFmt.currencySymbol = '\U000000a5'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "CNY"
	currencyFmt.currencyName = "Yuan"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = ','
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 156
	absValFmt.idString = "156"
	absValFmt.description = "Absolute Value Format: China"
	absValFmt.countryName = "China"
	absValFmt.abbreviatedCountryName = "CHN"
	absValFmt.alternateCountryName = "Peoples Republic of China"
	absValFmt.countryCodeTwoChar = "CN"
	absValFmt.countryCodeThreeChar = "CHN"
	absValFmt.countryCodeNumber = "156"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = '.'
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = ','
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 156
	absValFmt.idString = "156"
	signedNumValFmt.description = "Signed Number Format: China"
	signedNumValFmt.countryName = "China"
	signedNumValFmt.abbreviatedCountryName = "CHN"
	signedNumValFmt.alternateCountryName = "Peoples Republic of China"
	signedNumValFmt.countryCodeTwoChar = "CN"
	signedNumValFmt.countryCodeThreeChar = "CHN"
	signedNumValFmt.countryCodeNumber = "156"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "-127.54"
	signedNumValFmt.decimalSeparator = '.'
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = ','
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// France - Returns the number string format used in the
// French Republic.
//
// https://www.ibm.com/support/pages/english-and-french-currency-formats
// https://freeformatter.com/france-standards-code-snippets.html
func (nStrFmtCountry *NumStrFormatCountry) France() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 250
	currencyFmt.idString = "250"
	currencyFmt.description = "Currency Format: France"
	currencyFmt.countryName = "France"
	currencyFmt.abbreviatedCountryName = "France"
	currencyFmt.alternateCountryName = "French Republic"
	currencyFmt.countryCodeTwoChar = "FR"
	currencyFmt.countryCodeThreeChar = "FRA"
	currencyFmt.countryCodeNumber = "250"
	currencyFmt.positiveValueFmt = "127.54 $"
	currencyFmt.negativeValueFmt = "127.54 $ -"
	currencyFmt.decimalSeparator = ','
	currencyFmt.currencySymbol = '\U000020ac'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "EUR"
	currencyFmt.currencyName = "Euro"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = ' '
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 250
	absValFmt.idString = "250"
	absValFmt.description = "Absolute Value Format: France"
	absValFmt.countryName = "France"
	absValFmt.abbreviatedCountryName = "France"
	absValFmt.alternateCountryName = "French Republic"
	absValFmt.countryCodeTwoChar = "FR"
	absValFmt.countryCodeThreeChar = "FRA"
	absValFmt.countryCodeNumber = "250"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = ','
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = ' '
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 250
	absValFmt.idString = "250"
	signedNumValFmt.description = "Signed Number Format: France"
	signedNumValFmt.countryName = "France"
	signedNumValFmt.abbreviatedCountryName = "France"
	signedNumValFmt.alternateCountryName = "French Republic"
	signedNumValFmt.countryCodeTwoChar = "FR"
	signedNumValFmt.countryCodeThreeChar = "FRA"
	signedNumValFmt.countryCodeNumber = "250"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "127.54 -"
	signedNumValFmt.decimalSeparator = ','
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = ' '
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// Germany - Returns the number string format used in the
// Federal Republic of Germany.
//
// https://freeformatter.com/germany-standards-code-snippets.html
// https://www.evertype.com/standards/euro/formats.html
func (nStrFmtCountry *NumStrFormatCountry) Germany() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 276
	currencyFmt.idString = "276"
	currencyFmt.description = "Currency Format: Germany"
	currencyFmt.countryName = "Germany"
	currencyFmt.abbreviatedCountryName = "Germany"
	currencyFmt.alternateCountryName = "Federal Republic of Germany"
	currencyFmt.countryCodeTwoChar = "DE"
	currencyFmt.countryCodeThreeChar = "DEU"
	currencyFmt.countryCodeNumber = "276"
	currencyFmt.positiveValueFmt = "127.54 $"
	currencyFmt.negativeValueFmt = "127.54 $-"
	currencyFmt.decimalSeparator = ','
	currencyFmt.currencySymbol = '\U000020ac'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "EUR"
	currencyFmt.currencyName = "Euro"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = '.'
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 276
	absValFmt.idString = "276"
	absValFmt.description = "Absolute Value Format: Germany"
	absValFmt.countryName = "Germany"
	absValFmt.abbreviatedCountryName = "Germany"
	absValFmt.alternateCountryName = "Federal Republic of Germany"
	absValFmt.countryCodeTwoChar = "DE"
	absValFmt.countryCodeThreeChar = "DEU"
	absValFmt.countryCodeNumber = "276"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = ','
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = '.'
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 276
	absValFmt.idString = "276"
	signedNumValFmt.description = "Signed Number Format: Germany"
	signedNumValFmt.countryName = "Germany"
	signedNumValFmt.abbreviatedCountryName = "Germany"
	signedNumValFmt.alternateCountryName = "Federal Republic of Germany"
	signedNumValFmt.countryCodeTwoChar = "DE"
	signedNumValFmt.countryCodeThreeChar = "DEU"
	signedNumValFmt.countryCodeNumber = "276"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "127.54-"
	signedNumValFmt.decimalSeparator = ','
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = '.'
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// Italy - Returns the number string format used in the
// Italian Republic.
//
// https://freeformatter.com/italy-standards-code-snippets.html
// https://italian.stackexchange.com/questions/5674/what-is-the-correct-way-to-format-currency-in-italian
//
func (nStrFmtCountry *NumStrFormatCountry) Italy() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 380
	currencyFmt.idString = "380"
	currencyFmt.description = "Currency Format: Italy"
	currencyFmt.countryName = "Italy"
	currencyFmt.abbreviatedCountryName = "Italy"
	currencyFmt.alternateCountryName = "Italian Republic"
	currencyFmt.countryCodeTwoChar = "IT"
	currencyFmt.countryCodeThreeChar = "ITA"
	currencyFmt.countryCodeNumber = "380"
	currencyFmt.positiveValueFmt = "127.54 $"
	currencyFmt.negativeValueFmt = "127.54 $-"
	currencyFmt.decimalSeparator = ','
	currencyFmt.currencySymbol = '\U000020ac'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "EUR"
	currencyFmt.currencyName = "Euro"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = '.'
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 380
	absValFmt.idString = "380"
	absValFmt.description = "Absolute Value Format: Italy"
	absValFmt.countryName = "Italy"
	absValFmt.abbreviatedCountryName = "Italy"
	absValFmt.alternateCountryName = "Italian Republic"
	absValFmt.countryCodeTwoChar = "IT"
	absValFmt.countryCodeThreeChar = "ITA"
	absValFmt.countryCodeNumber = "380"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = ','
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = '.'
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 380
	absValFmt.idString = "380"
	signedNumValFmt.description = "Signed Number Format: Italy"
	signedNumValFmt.countryName = "Italy"
	signedNumValFmt.abbreviatedCountryName = "Italy"
	signedNumValFmt.alternateCountryName = "Italian Republic"
	signedNumValFmt.countryCodeTwoChar = "IT"
	signedNumValFmt.countryCodeThreeChar = "ITA"
	signedNumValFmt.countryCodeNumber = "380"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "127.54-"
	signedNumValFmt.decimalSeparator = ','
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = '.'
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// Israel - Returns the number string format used in the
// State of Israel.
//
//  https://freeformatter.com/israel-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Israel() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 376
	currencyFmt.idString = "376"
	currencyFmt.description = "Currency Format: Israel"
	currencyFmt.countryName = "Israel"
	currencyFmt.abbreviatedCountryName = "Israel"
	currencyFmt.alternateCountryName = "State of Israel"
	currencyFmt.countryCodeTwoChar = "IL"
	currencyFmt.countryCodeThreeChar = "ISR"
	currencyFmt.countryCodeNumber = "376"
	currencyFmt.positiveValueFmt = "$ 127.54"
	currencyFmt.negativeValueFmt = "-$ 127.54"
	currencyFmt.decimalSeparator = '.'
	currencyFmt.currencySymbol = '\U000020aa'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "ILS"
	currencyFmt.currencyName = "Shekel"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = ','
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 376
	absValFmt.idString = "376"
	absValFmt.description = "Absolute Value Format: Israel"
	absValFmt.countryName = "Israel"
	absValFmt.abbreviatedCountryName = "Israel"
	absValFmt.alternateCountryName = "State of Israel"
	absValFmt.countryCodeTwoChar = "IL"
	absValFmt.countryCodeThreeChar = "ISR"
	absValFmt.countryCodeNumber = "376"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = '.'
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = ','
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 376
	absValFmt.idString = "376"
	signedNumValFmt.description = "Signed Number Format: Israel"
	signedNumValFmt.countryName = "Israel"
	signedNumValFmt.abbreviatedCountryName = "Israel"
	signedNumValFmt.alternateCountryName = "State of Israel"
	signedNumValFmt.countryCodeTwoChar = "IL"
	signedNumValFmt.countryCodeThreeChar = "ISR"
	signedNumValFmt.countryCodeNumber = "376"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "- 127.54"
	signedNumValFmt.decimalSeparator = '.'
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = ','
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// UnitedKingdom - Returns the number string format used in the
// United Kingdom of Great Britain and Northern Ireland.
//
func (nStrFmtCountry *NumStrFormatCountry) UnitedKingdom() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 826
	currencyFmt.idString = "826"
	currencyFmt.description = "Currency Format: United Kingdom"
	currencyFmt.countryName = "United Kingdom"
	currencyFmt.abbreviatedCountryName = "UK"
	currencyFmt.alternateCountryName = "United Kingdom of Great Britain and Northern Ireland"
	currencyFmt.countryCodeTwoChar = "GB"
	currencyFmt.countryCodeThreeChar = "GBR"
	currencyFmt.countryCodeNumber = "826"
	currencyFmt.positiveValueFmt = "$127.54"
	currencyFmt.negativeValueFmt = "(-) $127.54"
	currencyFmt.decimalSeparator = '.'
	currencyFmt.currencySymbol = '\U000000a3'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "GBP"
	currencyFmt.currencyName = "Pound"
	currencyFmt.integerDigitsGroupingSequence =
		[]uint{3}
	currencyFmt.integerDigitsSeparator = ','
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 826
	absValFmt.idString = "826"
	absValFmt.description = "Absolute Value Format: United Kingdom"
	absValFmt.countryName = "United Kingdom"
	absValFmt.abbreviatedCountryName = "UK"
	absValFmt.alternateCountryName = "United Kingdom of Great Britain and Northern Ireland"
	absValFmt.countryCodeTwoChar = "GB"
	absValFmt.countryCodeThreeChar = "GBR"
	absValFmt.countryCodeNumber = "826"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = '.'
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	absValFmt.integerDigitsSeparator = ','
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 826
	signedNumValFmt.idString = "826"
	signedNumValFmt.description = "Signed Number Format: United Kingdom"
	signedNumValFmt.countryName = "United Kingdom"
	signedNumValFmt.abbreviatedCountryName = "UK"
	signedNumValFmt.alternateCountryName = "United Kingdom of Great Britain and Northern Ireland"
	signedNumValFmt.countryCodeTwoChar = "GB"
	signedNumValFmt.countryCodeThreeChar = "GBR"
	signedNumValFmt.countryCodeNumber = "826"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "-127.54"
	signedNumValFmt.decimalSeparator = '.'
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]uint{3}
	signedNumValFmt.integerDigitsSeparator = ','
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	signedNumValFmt.lock = new(sync.Mutex)

	var numStrFmtMap = make(map[NumStrValSpec]NumStrFormatter)

	numStrFmtMap[NumStrValSpec(0).CurrencyValue()] = currencyFmt
	numStrFmtMap[NumStrValSpec(0).AbsoluteValue()] = absValFmt
	numStrFmtMap[NumStrValSpec(0).SignedNumberValue()] = signedNumValFmt

	return numStrFmtMap
}

// UnitedStates - Returns the number string format used in the
// United States.
func (nStrFmtCountry *NumStrFormatCountry) UnitedStates(
	ePrefix string) (
	nStrFmtSpecDto NumStrFmtSpecDto,
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	if len(ePrefix) > 0 &&
		!strings.HasSuffix(ePrefix, "\n ") &&
		!strings.HasSuffix(ePrefix, "\n") {
		ePrefix += "\n"
	}

	ePrefix += "NumStrFormatCountry.UnitedStates() "

	setupDto := NumStrFmtSpecSetupDto{}
	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 840
	setupDto.IdString = "840"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 840
	setupDto.CountryIdString = "840"
	setupDto.CountryDescription = "Country Setup"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "United States"
	setupDto.CountryAbbreviatedName = "USA"

	setupDto.CountryAlternateNames =
		[]string{
			"United States of America",
			"America"}

	setupDto.CountryCodeTwoChar = "US"
	setupDto.CountryCodeThreeChar = "USA"
	setupDto.CountryCodeNumber = "840"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "USD"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbol = '\U00000024'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1
	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	err = nStrFmtSpecDto.SetFromFmtSpecSetupDto(
		&setupDto,
		ePrefix+
			"United States Setup\n")

	return nStrFmtSpecDto, err
}
