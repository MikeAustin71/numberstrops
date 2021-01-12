package numberstr

type NumStrFormatterCollectionSetupDto struct {
	IdNo                                      uint64
	IdString                                  string
	CountryName                               string
	AbbreviatedCountryName                    string
	AlternateCountryName                      string
	CountryCodeTwoChar                        string
	CountryCodeThreeChar                      string
	CountryCodeNumber                         string
	AbsoluteValFmt                            string
	AbsoluteValTurnOnIntegerDigitsSeparation  bool
	AbsoluteValNumFieldLen                    int
	SignedNumValPositiveValueFmt              string
	SignedNumValNegativeValueFmt              string
	SignedNumValTurnOnIntegerDigitsSeparation bool
	SignedNumValNumFieldLen                   int
	CurrencyPositiveValueFmt                  string
	CurrencyNegativeValueFmt                  string
	CurrencyDecimalDigits                     int
	CurrencyCode                              string
	CurrencyName                              string
	CurrencySymbol                            rune
	CurrencyTurnOnIntegerDigitsSeparation     bool
	CurrencyNumFieldLen                       int
	DecimalSeparator                          rune
	IntegerDigitsSeparator                    rune
	IntegerDigitsGroupingSequence             []uint
	SciNotMantissaLength                      uint
	SciNotExponentChar                        rune
	SciNotExponentUsesLeadingPlus             bool
	SciNotNumFieldLen                         int
}
