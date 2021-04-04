package numberstr

import (
	"fmt"
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
//  https://www.xe.com/iso4217.php#U
//  https://unicode-table.com/en
//
//  https://en.wikipedia.org/wiki/Indian_numbering_system
//  https://en.wikipedia.org/wiki/Chinese_numerals
//
//
// Countries:
//

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

// Albania - Returns the number string format used in
// The Republic of Albania.
//
//  https://www.xe.com/currency/all-albanian-lek
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Albania(
	fmtCollection *FormatterCollection,
	ePrefix *ErrPrefixDto) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"NumStrFormatCountry." +
			"Albania()")

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		FormatterBinary{}.NewUnitedStatesDefaultsPtr(),
		ePrefix.XCtx(
			"Binary"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		FormatterHexadecimal{}.NewUnitedStatesDefaultsPtr(),
		ePrefix.XCtx(
			"Hexadecimal"))

	if err != nil {
		return err
	}

}
