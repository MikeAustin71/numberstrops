package numberstr

type NumberStrFmtConfigDto struct {
	valueDisplaySpec                NumStrValSpec
	idNo                            int
	description                     string
	tag                             string
	countryName                     string // Full Name of country associated with this currency
	abbreviatedCountryName          string // Abbreviated Country Name
	alternateCountryName            string // Alternate Country Name Example: United States of America
	countryCodeTwoChar              string // ISO-3166 Two Character Country Code https://en.wikipedia.org/wiki/ISO_3166-2
	countryCodeThreeChar            string // ISO-3166 Three Character Country Code https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	countryCodeNumber               string // ISO-3166 Country Code Number: 3-Numeric Digits https://en.wikipedia.org/wiki/ISO_3166-1_numeric
	positiveValueFmt                string
	negativeValueFmt                string
	decimalSeparator                rune
	currencySymbol                  rune
	currencyDecimalDigits           int    // The number of digits after the decimal separator for currency presentations
	currencyCode                    string // Wold Currency Code (3-Characters). Reference: http://www.xe.com/symbols.php
	currencyName                    string // The common name of this currency i.e 'Dollar', 'Yuan' etc.
	integerDigitsSeparationSequence []int  // Integer Digit Separation. Usually = 3 for thousands separation
	integerDigitsSeparator          rune   // a.k.a. Thousands Separator
	turnOnIntegerDigitsSeparation   bool
	numFieldDto                     numberFieldDto
}
