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

/*

	setupDto := NumStrFmtSpecSetupDto{}
	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 0
	setupDto.IdString = ""
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 0
	setupDto.CountryIdString = ""
	setupDto.CountryDescription = "Country Setup - "
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = ""
	setupDto.CountryAbbreviatedName = ""

	setupDto.CountryAlternateNames =
		[]string{
			"" }

	setupDto.CountryCodeTwoChar = ""
	setupDto.CountryCodeThreeChar = ""
	setupDto.CountryCodeNumber = ""

	setupDto.AbsoluteValFmt = ""
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
			TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = ""
	setupDto.CurrencyNegativeValueFmt = ""
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = ""
	setupDto.CurrencyName = ""
	setupDto.CurrencySymbol = '0'
	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '0'
	setupDto.IntegerDigitsSeparator = '0'
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

	err = nStrFmtSpecDto.SetFromFmtSpecSetupDto(
		&setupDto,
		ePrefix+
			"United States Setup\n")

	return nStrFmtSpecDto, err

*/
