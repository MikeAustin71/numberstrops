package numberstr

import "sync"

type NumStrFmtSpecSetupDto struct {
	IdNo                                      uint64
	IdString                                  string
	Description                               string
	Tag                                       string
	CountryIdNo                               uint64
	CountryIdString                           string
	CountryDescription                        string
	CountryTag                                string
	CountryCultureName                        string
	CountryAbbreviatedName                    string
	CountryAlternateNames                     []string
	CountryCodeTwoChar                        string
	CountryCodeThreeChar                      string
	CountryCodeNumber                         string
	AbsoluteValFmt                            string
	AbsoluteValTurnOnIntegerDigitsSeparation  bool
	AbsoluteValNumSeps                        NumericSeparators
	AbsoluteValNumField                       NumberFieldDto
	CurrencyPositiveValueFmt                  string
	CurrencyNegativeValueFmt                  string
	CurrencyDecimalDigits                     uint
	CurrencyCode                              string
	CurrencyCodeNo                            string
	CurrencyName                              string
	CurrencySymbols                           []rune
	MinorCurrencyName                         string
	MinorCurrencySymbols                      []rune
	CurrencyTurnOnIntegerDigitsSeparation     bool
	CurrencyNumSeps                           NumericSeparators
	CurrencyNumField                          NumberFieldDto
	SignedNumValPositiveValueFmt              string
	SignedNumValNegativeValueFmt              string
	SignedNumValTurnOnIntegerDigitsSeparation bool
	SignedNumValNumSeps                       NumericSeparators
	SignedNumValNumField                      NumberFieldDto
	SciNotSignificandUsesLeadingPlus          bool
	SciNotMantissaLength                      uint
	SciNotExponentChar                        rune
	SciNotExponentUsesLeadingPlus             bool
	SciNotNumFieldLen                         int
	SciNotNumFieldTextJustify                 TextJustify
	Lock                                      *sync.Mutex
}
