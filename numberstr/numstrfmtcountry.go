package numberstr

import (
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
//  https://en.wikipedia.org/wiki/ISO_4217   Currency Codes
//  https://english.stackexchange.com/questions/124797/how-to-write-negative-currency-in-text
//  https://freeformatter.com/i18n-standards-code-snippets.html
//  https://www.evertype.com/standards/euro/formats.html
//  https://www.unicode.org/charts/PDF/U20A0.pdf
//  https://www.rapidtables.com/code/text/unicode-characters.html
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://www.ip2currency.com/currency-symbol
//
//
// Countries:
//
//  Argentina
//  Australia
//  Canada
//  Chile
//  China
//  Columbia
//  France
//  Germany
//  Israel
//  Italy
//  United Kingdom
//  United States

type NumStrFormatCountry struct {
	lock *sync.Mutex
}

// Ptr - Returns a pointer to a new instance of NumStrFormatCountry.
//
func (nStrFmtCountry NumStrFormatCountry) Ptr() *NumStrFormatCountry {
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
//
// https://freeformatter.com/argentina-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Argentina() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 32
	setupDto.IdString = "032"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 32
	setupDto.CountryIdString = "032"
	setupDto.CountryDescription = "Country Setup - Argentina"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Argentina"
	setupDto.CountryAbbreviatedName = "Argentina"

	setupDto.CountryAlternateNames =
		[]string{
			"Argentine Republic",
			"The Argentine Republic"}

	setupDto.CountryCodeTwoChar = "AR"
	setupDto.CountryCodeThreeChar = "ARG"
	setupDto.CountryCodeNumber = "032"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ARS"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "- 127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Australia - Returns the number string format used in
// Australia.
//
func (nStrFmtCountry *NumStrFormatCountry) Australia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 36
	setupDto.IdString = "36"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 36
	setupDto.CountryIdString = "036"
	setupDto.CountryDescription = "Country Setup - Australia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Australia"
	setupDto.CountryAbbreviatedName = "Australia"

	setupDto.CountryAlternateNames =
		[]string{
			"Commonwealth of Australia",
			"The Commonwealth of Australia"}

	setupDto.CountryCodeTwoChar = "AU"
	setupDto.CountryCodeThreeChar = "AUS"
	setupDto.CountryCodeNumber = "036"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "AUD"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Canada - Returns the number string format used in
// Canada.
//
func (nStrFmtCountry *NumStrFormatCountry) Canada() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 124
	setupDto.IdString = "124"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 124
	setupDto.CountryIdString = "124"
	setupDto.CountryDescription = "Country Setup - Canada"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Canada"
	setupDto.CountryAbbreviatedName = "Canada"

	setupDto.CountryAlternateNames =
		[]string{
			"Canada"}

	setupDto.CountryCodeTwoChar = "CA"
	setupDto.CountryCodeThreeChar = "CAN"
	setupDto.CountryCodeNumber = "124"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CAD"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// CanadaFrench - Returns the number string format used in
// French Canada.
//
func (nStrFmtCountry *NumStrFormatCountry) CanadaFrench() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 124
	setupDto.IdString = "124"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 124
	setupDto.CountryIdString = "124"
	setupDto.CountryDescription = "Country Setup - Canada French"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Canada French"
	setupDto.CountryAbbreviatedName = "Canada French"

	setupDto.CountryAlternateNames =
		[]string{
			"Canada French",
			"French Canadian"}

	setupDto.CountryCodeTwoChar = "CA"
	setupDto.CountryCodeThreeChar = "CAN"
	setupDto.CountryCodeNumber = "124"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CAD"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54 -"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

//  Czechia - Returns the number string format used in the
// the The Czech Republic.
//
func (nStrFmtCountry *NumStrFormatCountry) Czechia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 203
	setupDto.IdString = "203"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 203
	setupDto.CountryIdString = "203"
	setupDto.CountryDescription = "Country Setup - Chile"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Czechia"
	setupDto.CountryAbbreviatedName = "Czechia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Czech Republic",
			"Czech Republic"}

	setupDto.CountryCodeTwoChar = "CZ"
	setupDto.CountryCodeThreeChar = "CZE"
	setupDto.CountryCodeNumber = "203"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "CZK"
	setupDto.CurrencyName = "Koruna"
	setupDto.CurrencySymbols = '\U0004b10d'
	setupDto.MinorCurrencyName = "Haler"
	setupDto.MinorCurrencySymbols = 0

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Chile - Returns the number string format used in the
// the The Republic of Chile.
//
func (nStrFmtCountry *NumStrFormatCountry) Chile() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 152
	setupDto.IdString = "152"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 152
	setupDto.CountryIdString = "152"
	setupDto.CountryDescription = "Country Setup - Chile"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Chile"
	setupDto.CountryAbbreviatedName = "Chile"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Chile"}

	setupDto.CountryCodeTwoChar = "CL"
	setupDto.CountryCodeThreeChar = "CHL"
	setupDto.CountryCodeNumber = "152"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "CLP"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.MinorCurrencyName = "centavo"
	setupDto.MinorCurrencySymbols = 0

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// China - Returns the number string format used in the
// Peoples Republic of China.
//
func (nStrFmtCountry *NumStrFormatCountry) China() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 156
	setupDto.IdString = "156"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 156
	setupDto.CountryIdString = "156"
	setupDto.CountryDescription = "Country Setup - China"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "China"
	setupDto.CountryAbbreviatedName = "CHN"

	setupDto.CountryAlternateNames =
		[]string{
			"Peoples Republic of China",
			"The Peoples Republic of China"}

	setupDto.CountryCodeTwoChar = "CN"
	setupDto.CountryCodeThreeChar = "CHN"
	setupDto.CountryCodeNumber = "156"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CNY"
	setupDto.CurrencyName = "Yuan"
	setupDto.CurrencySymbols = '\U000000a5'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Columbia - Returns the number string format used in the
// Columbia.
//
func (nStrFmtCountry *NumStrFormatCountry) Columbia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 170
	setupDto.IdString = "170"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 170
	setupDto.CountryIdString = "170"
	setupDto.CountryDescription = "Country Setup - Columbia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Columbia"
	setupDto.CountryAbbreviatedName = "Columbia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Colombia",
			"Republic of Colombia"}

	setupDto.CountryCodeTwoChar = "CO"
	setupDto.CountryCodeThreeChar = "COL"
	setupDto.CountryCodeNumber = "170"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "COP"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.MinorCurrencyName = "centavo"
	setupDto.MinorCurrencySymbols = 0
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// France - Returns the number string format used in the
// French Republic.
//
// https://www.ibm.com/support/pages/english-and-french-currency-formats
// https://freeformatter.com/france-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) France() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 250
	setupDto.IdString = "250"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 250
	setupDto.CountryIdString = "250"
	setupDto.CountryDescription = "Country Setup - France"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "France"
	setupDto.CountryAbbreviatedName = "France"

	setupDto.CountryAlternateNames =
		[]string{
			"French Republic",
			"The French Republic"}

	setupDto.CountryCodeTwoChar = "FR"
	setupDto.CountryCodeThreeChar = "FRA"
	setupDto.CountryCodeNumber = "250"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $ -"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EUR"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = '\U000020ac'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' '
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54 -"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Germany - Returns the number string format used in the
// Federal Republic of Germany.
//
// https://freeformatter.com/germany-standards-code-snippets.html
// https://www.evertype.com/standards/euro/formats.html
//
func (nStrFmtCountry *NumStrFormatCountry) Germany() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 276
	setupDto.IdString = "276"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 276
	setupDto.CountryIdString = "276"
	setupDto.CountryDescription = "Country Setup - Germany"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Germany"
	setupDto.CountryAbbreviatedName = "Germany"

	setupDto.CountryAlternateNames =
		[]string{
			"Federal Republic of Germany",
			"The Federal Republic of Germany"}

	setupDto.CountryCodeTwoChar = "DE"
	setupDto.CountryCodeThreeChar = "DEU"
	setupDto.CountryCodeNumber = "276"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EUR"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = '\U000020ac'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Italy - Returns the number string format used in the
// Italian Republic.
//
// https://freeformatter.com/italy-standards-code-snippets.html
// https://italian.stackexchange.com/questions/5674/what-is-the-correct-way-to-format-currency-in-italian
//
func (nStrFmtCountry *NumStrFormatCountry) Italy() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 380
	setupDto.IdString = "380"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 380
	setupDto.CountryIdString = "380"
	setupDto.CountryDescription = "Country Setup - Italy"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Italy"
	setupDto.CountryAbbreviatedName = "Italy"

	setupDto.CountryAlternateNames =
		[]string{
			"Italian Republic",
			"The Italian Republic"}

	setupDto.CountryCodeTwoChar = "IT"
	setupDto.CountryCodeThreeChar = "ITA"
	setupDto.CountryCodeNumber = "380"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EUR"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = '\U000020ac'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Israel - Returns the number string format used in the
// State of Israel.
//
//  https://freeformatter.com/israel-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Israel() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 376
	setupDto.IdString = "376"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 376
	setupDto.CountryIdString = "376"
	setupDto.CountryDescription = "Country Setup - Israel"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Israel"
	setupDto.CountryAbbreviatedName = "Israel"

	setupDto.CountryAlternateNames =
		[]string{
			"State of Israel",
			"The State of Israel"}

	setupDto.CountryCodeTwoChar = "IL"
	setupDto.CountryCodeThreeChar = "ISR"
	setupDto.CountryCodeNumber = "376"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ILS"
	setupDto.CurrencyName = "Shekel"
	setupDto.CurrencySymbols = '\U000020aa'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = ""
	setupDto.SignedNumValNegativeValueFmt = ""
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// UnitedKingdom - Returns the number string format used in the
// United Kingdom of Great Britain and Northern Ireland.
//
func (nStrFmtCountry *NumStrFormatCountry) UnitedKingdom() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 826
	setupDto.IdString = "826"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 826
	setupDto.CountryIdString = "826"
	setupDto.CountryDescription = "Country Setup - United Kingdom"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "United Kingdom"
	setupDto.CountryAbbreviatedName = "UK"

	setupDto.CountryAlternateNames =
		[]string{
			"United Kingdom of Great Britain and Northern Ireland",
			"England",
			"Great Britain"}

	setupDto.CountryCodeTwoChar = "GB"
	setupDto.CountryCodeThreeChar = "GBR"
	setupDto.CountryCodeNumber = "826"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "(-) $127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "GBP"
	setupDto.CurrencyName = "Pound"
	setupDto.CurrencySymbols = '\U000000a3'
	setupDto.CurrencyName = "Pence"
	setupDto.CurrencySymbols = 0
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// UnitedStates - Returns the number string format used in the
// the United States.
//
func (nStrFmtCountry *NumStrFormatCountry) UnitedStates() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 840
	setupDto.IdString = "840"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 840
	setupDto.CountryIdString = "840"
	setupDto.CountryDescription = "Country Setup - United States"
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

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "USD"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = '\U00000024'
	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = '\U000000a2'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}
