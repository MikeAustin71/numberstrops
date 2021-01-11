package numberstr

import "sync"

// NumStrFormatCountry - Returns the number string formats used
// by specific countries.
//
// Sources:
// https://gist.github.com/bzerangue/5484121
// http://symbologic.info/currency.htm
// http://www.xe.com/symbols.php
// https://www.countrycode.org/
// https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//  https://www.codeproject.com/articles/78175/international-number-formats
//  https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//  https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://english.stackexchange.com/questions/124797/how-to-write-negative-currency-in-text
//  https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
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
		[]int{3}
	currencyFmt.integerDigitsSeparator = '.'
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = numberFieldDto{
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
		[]int{3}
	absValFmt.integerDigitsSeparator = '.'
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = numberFieldDto{
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
		[]int{3}
	signedNumValFmt.integerDigitsSeparator = '.'
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = numberFieldDto{
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

// Australia - Returns the number string format used in the
// Australia.
func (nStrFmtCountry *NumStrFormatCountry) Australia() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U00000024'
	newNStrFmtDto.currencyFmt.currencyCode = "AUD"
	newNStrFmtDto.currencyFmt.currencyName = "Dollar"
	newNStrFmtDto.currencyFmt.countryName = "Australia"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "Australia"
	newNStrFmtDto.currencyFmt.alternateCountryName = "Commonwealth of Australia"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "AU"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "AUS"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "036"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$ 127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "-$ 127.54"
	newNStrFmtDto.currencyFmt.turnOnIntegerDigitsSeparation = true
	newNStrFmtDto.currencyFmt.currencyDecimalDigits = 2

	newNStrFmtDto.decimalSeparator = '.'
	newNStrFmtDto.integerDigitsSeparator = ','
	newNStrFmtDto.integerDigitSeparation = []int{3}
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// Canada - Returns the number string format used in the
// Canada.
func (nStrFmtCountry *NumStrFormatCountry) Canada() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U00000024'
	newNStrFmtDto.currencyFmt.currencyCode = "CAD"
	newNStrFmtDto.currencyFmt.currencyName = "Dollar"
	newNStrFmtDto.currencyFmt.countryName = "Canada"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "Canada"
	newNStrFmtDto.currencyFmt.alternateCountryName = "Canada"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "CA"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "CAN"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "124"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$ 127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "-$ 127.54"
	newNStrFmtDto.currencyFmt.turnOnIntegerDigitsSeparation = true
	newNStrFmtDto.currencyFmt.currencyDecimalDigits = 2

	newNStrFmtDto.decimalSeparator = '.'
	newNStrFmtDto.integerDigitsSeparator = ','
	newNStrFmtDto.integerDigitSeparation = []int{3}
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// CanadaFrench - Returns the number string format used in
// French Canada.
func (nStrFmtCountry *NumStrFormatCountry) CanadaFrench() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U00000024'
	newNStrFmtDto.currencyFmt.currencyCode = "CAD"
	newNStrFmtDto.currencyFmt.currencyName = "Dollar"
	newNStrFmtDto.currencyFmt.countryName = "Canada"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "Canada French"
	newNStrFmtDto.currencyFmt.alternateCountryName = "French Canadian"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "CA"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "CAN"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "124"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "127.54 $"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "127.54 $-"
	newNStrFmtDto.currencyFmt.turnOnIntegerDigitsSeparation = true
	newNStrFmtDto.currencyFmt.currencyDecimalDigits = 2

	newNStrFmtDto.decimalSeparator = ','
	newNStrFmtDto.integerDigitsSeparator = '.'
	newNStrFmtDto.integerDigitSeparation = []int{3}
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// China - Returns the number string format used in the
// Peoples Republic of China.
func (nStrFmtCountry *NumStrFormatCountry) China() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U000000a5'
	newNStrFmtDto.currencyFmt.currencyCode = "CNY"
	newNStrFmtDto.currencyFmt.currencyName = "Yuan"
	newNStrFmtDto.currencyFmt.countryName = "China"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "CHN"
	newNStrFmtDto.currencyFmt.alternateCountryName = "Peoples Republic of China"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "CN"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "CHN"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "156"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$ 127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "-$ 127.54"
	newNStrFmtDto.currencyFmt.turnOnIntegerDigitsSeparation = true
	newNStrFmtDto.currencyFmt.currencyDecimalDigits = 2

	newNStrFmtDto.decimalSeparator = '.'
	newNStrFmtDto.integerDigitsSeparator = ','
	newNStrFmtDto.integerDigitSeparation = []int{3}
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// France - Returns the number string format used in the
// French Republic.
//
// https://www.ibm.com/support/pages/english-and-french-currency-formats
// https://freeformatter.com/france-standards-code-snippets.html
func (nStrFmtCountry *NumStrFormatCountry) France() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 250
	newNStrFmtDto.currencyFmt.currencySymbol = '\U000020ac'
	newNStrFmtDto.currencyFmt.currencyCode = "EUR"
	newNStrFmtDto.currencyFmt.currencyName = "Euro"
	newNStrFmtDto.currencyFmt.countryName = "France"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "France"
	newNStrFmtDto.currencyFmt.alternateCountryName = "French Republic"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "FR"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "FRA"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "250"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "127.54 $"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "127.54 $ -"
	newNStrFmtDto.currencyFmt.turnOnIntegerDigitsSeparation = true
	newNStrFmtDto.currencyFmt.currencyDecimalDigits = 2
	newNStrFmtDto.integerDigitSeparation = []int{3}

	newNStrFmtDto.decimalSeparator = ','
	newNStrFmtDto.integerDigitsSeparator = ' '
	newNStrFmtDto.integerDigitSeparation = []int{3}
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// Germany - Returns the number string format used in the
// Federal Republic of Germany.
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Germany() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U000020ac'
	newNStrFmtDto.currencyFmt.currencyCode = "EUR"
	newNStrFmtDto.currencyFmt.currencyName = "Euro"
	newNStrFmtDto.currencyFmt.countryName = "Germany"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "Germany"
	newNStrFmtDto.currencyFmt.alternateCountryName = "Federal Republic of Germany"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "DE"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "DEU"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "276"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "127.54$"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "127.54$-"
	newNStrFmtDto.currencyFmt.turnOnIntegerDigitsSeparation = true
	newNStrFmtDto.currencyFmt.currencyDecimalDigits = 2

	newNStrFmtDto.decimalSeparator = ','
	newNStrFmtDto.integerDigitsSeparator = '.'
	newNStrFmtDto.integerDigitSeparation = []int{3}
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// Italy - Returns the number string format used in the
// Italian Republic.
//
//
//
func (nStrFmtCountry *NumStrFormatCountry) Italy() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U000020ac'
	newNStrFmtDto.currencyFmt.currencyCode = "EUR"
	newNStrFmtDto.currencyFmt.currencyName = "Euro"
	newNStrFmtDto.currencyFmt.countryName = "Italy"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "Italy"
	newNStrFmtDto.currencyFmt.alternateCountryName = "Italian Republic"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "IT"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "ITA"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "380"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "- $127.54"

	newNStrFmtDto.decimalSeparator = ','
	newNStrFmtDto.integerDigitsSeparator = '.'
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

func (nStrFmtCountry *NumStrFormatCountry) Israel() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U000020aa'
	newNStrFmtDto.currencyFmt.currencyCode = "ILS"
	newNStrFmtDto.currencyFmt.currencyName = "Shekel"
	newNStrFmtDto.currencyFmt.countryName = "Israel"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "Israel"
	newNStrFmtDto.currencyFmt.alternateCountryName = "State of Israel"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "IL"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "ISR"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "376"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "- $127.54"

	newNStrFmtDto.decimalSeparator = '.'
	newNStrFmtDto.integerDigitsSeparator = ','
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

func (nStrFmtCountry *NumStrFormatCountry) UnitedKingdom() NumStrFormatDto {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	newNStrFmtDto := NumStrFormatDto{}

	newNStrFmtDto.valueDisplaySpec =
		NumStrValSpec(0).SignedNumberValue()

	newNStrFmtDto.positiveValueFmt = "127.54"
	newNStrFmtDto.negativeValueFmt = "-127.54"

	newNStrFmtDto.currencyFmt = CurrencySymbolDto{}

	newNStrFmtDto.currencyFmt.idNo = 1
	newNStrFmtDto.currencyFmt.currencySymbol = '\U000000a3'
	newNStrFmtDto.currencyFmt.currencyCode = "GBP"
	newNStrFmtDto.currencyFmt.currencyName = "Pound"
	newNStrFmtDto.currencyFmt.countryName = "United Kingdom"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "UK"
	newNStrFmtDto.currencyFmt.alternateCountryName = "United Kingdom of Great Britain and Northern Ireland"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "GB"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "GBR"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "826"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "(-) $127.54"

	newNStrFmtDto.decimalSeparator = '.'
	newNStrFmtDto.integerDigitsSeparator = ','
	newNStrFmtDto.turnOnIntegerDigitSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// UnitedStates - Returns the number string format used in the
// United States.
func (nStrFmtCountry *NumStrFormatCountry) UnitedStates() map[NumStrValSpec]NumStrFormatter {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	currencyFmt := NumStrFormatter{}
	currencyFmt.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()
	currencyFmt.idNo = 840
	currencyFmt.idString = "840"
	currencyFmt.description = "Currency Format United States"
	currencyFmt.countryName = "United States"
	currencyFmt.abbreviatedCountryName = "USA"
	currencyFmt.alternateCountryName = "United States of America"
	currencyFmt.countryCodeTwoChar = "US"
	currencyFmt.countryCodeThreeChar = "USA"
	currencyFmt.countryCodeNumber = "840"
	currencyFmt.positiveValueFmt = "$127.54"
	currencyFmt.negativeValueFmt = "-$127.54"
	currencyFmt.decimalSeparator = '.'
	currencyFmt.currencySymbol = '\U00000024'
	currencyFmt.currencyDecimalDigits = 2
	currencyFmt.currencyCode = "USD"
	currencyFmt.currencyName = "Dollar"
	currencyFmt.integerDigitsGroupingSequence =
		[]int{3}
	currencyFmt.integerDigitsSeparator = ','
	currencyFmt.turnOnIntegerDigitsSeparation = true
	currencyFmt.numFieldDto = numberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	currencyFmt.lock = new(sync.Mutex)

	absValFmt := NumStrFormatter{}
	absValFmt.valueDisplaySpec = NumStrValSpec(0).AbsoluteValue()
	absValFmt.idNo = 840
	absValFmt.idString = "840"
	absValFmt.description = "United States"
	absValFmt.countryName = "United States"
	absValFmt.abbreviatedCountryName = "USA"
	absValFmt.alternateCountryName = "United States of America"
	absValFmt.countryCodeTwoChar = "US"
	absValFmt.countryCodeThreeChar = "USA"
	absValFmt.countryCodeNumber = "840"
	absValFmt.positiveValueFmt = "127.54"
	absValFmt.negativeValueFmt = "127.54"
	absValFmt.decimalSeparator = '.'
	absValFmt.currencySymbol = 0
	absValFmt.currencyDecimalDigits = -1
	absValFmt.currencyCode = ""
	absValFmt.currencyName = ""
	absValFmt.integerDigitsGroupingSequence =
		[]int{3}
	absValFmt.integerDigitsSeparator = ','
	absValFmt.turnOnIntegerDigitsSeparation = false
	absValFmt.numFieldDto = numberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}
	absValFmt.lock = new(sync.Mutex)

	signedNumValFmt := NumStrFormatter{}
	signedNumValFmt.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	signedNumValFmt.idNo = 840
	absValFmt.idString = "840"
	signedNumValFmt.description = "Signed Number Format United States"
	signedNumValFmt.countryName = "United States"
	signedNumValFmt.abbreviatedCountryName = "USA"
	signedNumValFmt.alternateCountryName = "United States of America"
	signedNumValFmt.countryCodeTwoChar = "US"
	signedNumValFmt.countryCodeThreeChar = "USA"
	signedNumValFmt.countryCodeNumber = "840"
	signedNumValFmt.positiveValueFmt = "127.54"
	signedNumValFmt.negativeValueFmt = "-127.54"
	signedNumValFmt.decimalSeparator = '.'
	signedNumValFmt.currencySymbol = 0
	signedNumValFmt.currencyDecimalDigits = -1
	signedNumValFmt.currencyCode = ""
	signedNumValFmt.currencyName = ""
	signedNumValFmt.integerDigitsGroupingSequence =
		[]int{3}
	signedNumValFmt.integerDigitsSeparator = ','
	signedNumValFmt.turnOnIntegerDigitsSeparation = false
	signedNumValFmt.numFieldDto = numberFieldDto{
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
