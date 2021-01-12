package numberstr

import "sync"

type NumStrFmtSpecCountryDto struct {
	CountryCultureName     string
	AbbreviatedCountryName string
	AlternateCountryNames  []string
	CountryCodeTwoChar     string
	CountryCodeThreeChar   string
	CountryCodeNumber      string

	lock *sync.Mutex
}
