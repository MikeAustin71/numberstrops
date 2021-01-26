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
	AbsoluteValNumFieldLen                    int
	AbsoluteValNumFieldTextJustify            TextJustify
	CurrencyPositiveValueFmt                  string
	CurrencyNegativeValueFmt                  string
	CurrencyDecimalDigits                     uint
	CurrencyCode                              string
	CurrencyName                              string
	CurrencySymbol                            rune
	CurrencyTurnOnIntegerDigitsSeparation     bool
	CurrencyNumFieldLen                       int
	CurrencyNumFieldTextJustify               TextJustify
	DecimalSeparator                          rune
	IntegerDigitsSeparator                    rune
	IntegerDigitsGroupingSequence             []uint
	SignedNumValPositiveValueFmt              string
	SignedNumValNegativeValueFmt              string
	SignedNumValTurnOnIntegerDigitsSeparation bool
	SignedNumValNumFieldLen                   int
	SignedNumValNumFieldTextJustify           TextJustify
	SciNotSignificandUsesLeadingPlus          bool
	SciNotMantissaLength                      uint
	SciNotExponentChar                        rune
	SciNotExponentUsesLeadingPlus             bool
	SciNotNumFieldLen                         int
	SciNotNumFieldTextJustify                 TextJustify
	Lock                                      *sync.Mutex
}
