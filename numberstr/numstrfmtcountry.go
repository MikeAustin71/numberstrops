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
//
//
//
// Countries:
//
//  United Kingdom
//  United States

type NumStrFormatCountry struct {
	lock *sync.Mutex
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
	newNStrFmtDto.thousandsSeparator = ','
	newNStrFmtDto.turnOnThousandsSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}

// UnitedStates - Returns the number string format used in the
// United States.
func (nStrFmtCountry *NumStrFormatCountry) UnitedStates() NumStrFormatDto {

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
	newNStrFmtDto.currencyFmt.currencyCode = "USD"
	newNStrFmtDto.currencyFmt.currencyName = "Dollar"
	newNStrFmtDto.currencyFmt.countryName = "United States"
	newNStrFmtDto.currencyFmt.abbreviatedCountryName = "USA"
	newNStrFmtDto.currencyFmt.alternateCountryName = "United States of America"
	newNStrFmtDto.currencyFmt.countryCodeTwoChar = "US"
	newNStrFmtDto.currencyFmt.countryCodeThreeChar = "USA"
	newNStrFmtDto.currencyFmt.countryCodeNumber = "840"
	newNStrFmtDto.currencyFmt.positiveValCurrFormat = "$127.54"
	newNStrFmtDto.currencyFmt.negativeValCurrFormat = "-$127.54"

	newNStrFmtDto.decimalSeparator = '.'
	newNStrFmtDto.thousandsSeparator = ','
	newNStrFmtDto.turnOnThousandsSeparator = false
	newNStrFmtDto.numberFieldDto.requestedNumFieldLength = -1

	newNStrFmtDto.lock = new(sync.Mutex)

	return newNStrFmtDto
}
