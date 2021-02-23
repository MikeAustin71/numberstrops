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
//  https://en.wikipedia.org/wiki/List_of_circulating_currencies - Symbols with decoding
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
//  Austria
//  Canada
//  CanadaFrench
//  Chile
//  China
//  Columbia
//  Czechia
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
	setupDto.CurrencyCodeNo = "032"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCodeNo = "036"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

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

// Austria - Returns the number string format used in
// Austria.
//
func (nStrFmtCountry *NumStrFormatCountry) Austria() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 40
	setupDto.IdString = "040"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 40
	setupDto.CountryIdString = "040"
	setupDto.CountryDescription = "Country Setup - Austria"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Austria"
	setupDto.CountryAbbreviatedName = "Austria"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Austria",
			"Republic of Austria"}

	setupDto.CountryCodeTwoChar = "AT"
	setupDto.CountryCodeThreeChar = "AUT"
	setupDto.CountryCodeNumber = "040"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ATS"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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

// Brazil - Returns the number string format used in the
// Brazil.
//
func (nStrFmtCountry *NumStrFormatCountry) Brazil() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 76
	setupDto.IdString = "076"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 76
	setupDto.CountryIdString = "076"
	setupDto.CountryDescription = "Country Setup - Brazil"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Brazil"
	setupDto.CountryAbbreviatedName = "Brazil"

	setupDto.CountryAlternateNames =
		[]string{
			"The Federative Republic of Brazil"}

	setupDto.CountryCodeTwoChar = "BR"
	setupDto.CountryCodeThreeChar = "BRA"
	setupDto.CountryCodeNumber = "076"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BRL"
	setupDto.CurrencyCodeNo = "986"
	setupDto.CurrencyName = "Real"
	setupDto.CurrencySymbols = []rune{'\U00000052', '\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = []rune{}

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
	setupDto.CurrencyCodeNo = "124"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

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
	setupDto.CurrencyCodeNo = "124"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

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
	setupDto.CurrencyCodeNo = "152"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCodeNo = "156"
	setupDto.CurrencyName = "Yuan"
	setupDto.CurrencySymbols = []rune{'\U000000a5'}

	setupDto.MinorCurrencyName = "Jiao"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCodeNo = "170"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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

// Czechia - Returns the number string format used in the
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
	setupDto.CountryDescription = "Country Setup - Czechia"
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
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CZK"
	setupDto.CurrencyCodeNo = "203"
	setupDto.CurrencyName = "Koruna"
	setupDto.CurrencySymbols = []rune{
		'\U0000004b', '\U0000010d'}

	setupDto.MinorCurrencyName = "Haler"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0)

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

// Denmark - Returns the number string format used in the
// Kingdom of Denmark.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Denmark() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 208
	setupDto.IdString = "208"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 208
	setupDto.CountryIdString = "208"
	setupDto.CountryDescription = "Country Setup - Denmark"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Denmark"
	setupDto.CountryAbbreviatedName = "Denmark"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Denmark",
			"Kingdom of Denmark"}

	setupDto.CountryCodeTwoChar = "DK"
	setupDto.CountryCodeThreeChar = "DNK"
	setupDto.CountryCodeNumber = "208"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "DKK"
	setupDto.CurrencyCodeNo = "208"
	setupDto.CurrencyName = "Krone"
	setupDto.CurrencySymbols = []rune{'\U0000006b', '\U00000072'}

	setupDto.MinorCurrencyName = "øre"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCode = "FRF"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCode = "DEM"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCodeNo = "376"
	setupDto.CurrencyName = "Shekel"
	setupDto.CurrencySymbols = []rune{'\U000020aa'}

	setupDto.MinorCurrencyName = "Agorot"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyName = "ITL"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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

// Netherlands - Returns the number string format used in the
// Kingdom of Netherlands.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Netherlands() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 528
	setupDto.IdString = "528"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 528
	setupDto.CountryIdString = "528"
	setupDto.CountryDescription = "Country Setup - Netherlands"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Netherlands"
	setupDto.CountryAbbreviatedName = "Netherlands"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of the Netherlands",
			"Kingdom of the Netherlands"}

	setupDto.CountryCodeTwoChar = "NL"
	setupDto.CountryCodeThreeChar = "NLD"
	setupDto.CountryCodeNumber = "528"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = " $ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NLG"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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

// Norway - Returns the number string format used in the
// Kingdom of Norway.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Norway() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 578
	setupDto.IdString = "578"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 578
	setupDto.CountryIdString = "578"
	setupDto.CountryDescription = "Country Setup - Norway"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Norway"
	setupDto.CountryAbbreviatedName = "Norway"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Norway",
			"Kingdom of Norway"}

	setupDto.CountryCodeTwoChar = "NO"
	setupDto.CountryCodeThreeChar = "NOR"
	setupDto.CountryCodeNumber = "578"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NOK"
	setupDto.CurrencyCodeNo = "578"
	setupDto.CurrencyName = "Krone"
	setupDto.CurrencySymbols = []rune{'\U0000006b', '\U00000072'}

	setupDto.MinorCurrencyName = "øre"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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

// Sweden - Returns the number string format used in the
// Kingdom of Sweden.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Sweden() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 752
	setupDto.IdString = "752"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 752
	setupDto.CountryIdString = "752"
	setupDto.CountryDescription = "Country Setup - Sweden"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Sweden"
	setupDto.CountryAbbreviatedName = "Sweden"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Sweden",
			"Kingdom of Sweden"}

	setupDto.CountryCodeTwoChar = "SE"
	setupDto.CountryCodeThreeChar = "SWE"
	setupDto.CountryCodeNumber = "752"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54 $-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "SEK"
	setupDto.CurrencyCodeNo = "752"
	setupDto.CurrencyName = "Krona"
	setupDto.CurrencySymbols = []rune{'\U0000006b', '\U00000072'}

	setupDto.MinorCurrencyName = "øre"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
	setupDto.CurrencyCodeNo = "826"
	setupDto.CurrencyName = "Pound"
	setupDto.CurrencySymbols = []rune{'\U000000a3'}

	setupDto.MinorCurrencyName = "Pence"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

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
// United States.
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
	setupDto.CurrencyCodeNo = "840"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

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
