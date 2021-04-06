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
	errorPrefix interface{}) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	var ePrefix *ErrPrefixDto

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatCountry."+
			"Albania()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	// Adds Binary, Hexadecimal, Octal and
	// Scientific Notation Formatters
	err = numStrFormatCountryMechanics{}.ptr().
		addBaseFormatters(
			fmtCollection,
			ePrefix.XCtx("fmtCollection"))

	if err != nil {
		return err
	}

	var country FormatterCountry

	country,
		err = FormatterCountry{}.NewWithComponents(
		8,               // idNo
		"008",           // idString
		"Country Setup", // Desc
		"",              // Tag
		"Albania",       // Country Culture
		"Albania",       // Abbrv Country Name
		[]string{
			"The Republic of Albania",
			"Republic of Albania"}, // Alternate Cntry Names
		"AL",  // countryCodeTwoChar
		"ALB", // countryCodeThreeChar
		"008", // countryCodeNumber
		ePrefix.XCtx("country"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&country,
		ePrefix.XCtx(
			"&country"))

	if err != nil {
		return err
	}

	var absVal FormatterAbsoluteValue

	absVal,
		err = FormatterAbsoluteValue{}.NewBasic(
		",",                    // decimalSeparatorChars
		" ",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // absoluteValFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx("absVal"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&absVal,
		ePrefix.XCtx(
			"&absVal"))

	if err != nil {
		return err
	}

	var currency FormatterCurrency

	currency,
		err = FormatterCurrency{}.NewBasic(
		",",         // decimalSeparatorChars
		" ",         // thousandsSeparatorChars
		true,        // turnOnThousandsSeparator
		"$ 127.54",  // positiveValueFmt
		"$ -127.54", // negativeValueFmt
		2,           // decimalDigits
		"ALL",       // currencyCode
		"008",       // currencyCodeNo
		"Lek",       // currencyName
		[]rune{
			'\U0000004c',
			'\U00000065',
			'\U0000006b',
		}, //currencySymbols
		"Qindarkë",             // minorCurrencyName
		make([]rune, 0),        // minorCurrencySymbols
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"currency"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&currency,
		ePrefix.XCtx(
			"&currency"))

	if err != nil {
		return err
	}

	var signedNum FormatterSignedNumber

	signedNum,
		err = FormatterSignedNumber{}.NewBasic(
		",",                    // decimalSeparatorChars
		" ",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // positiveValueFmt
		"-127.54",              // negativeValueFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"signedNum"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&signedNum,
		ePrefix.XCtx(
			"&signedNum"))

	return err
}

// Argentina - Returns the number string format used in the
// Argentina.
//
// https://freeformatter.com/argentina-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Argentina(
	fmtCollection *FormatterCollection,
	errorPrefix interface{}) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	var ePrefix *ErrPrefixDto

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatCountry."+
			"Argentina()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	// Adds Binary, Hexadecimal, Octal and
	// Scientific Notation Formatters
	err = numStrFormatCountryMechanics{}.ptr().
		addBaseFormatters(
			fmtCollection,
			ePrefix.XCtx("fmtCollection"))

	if err != nil {
		return err
	}

	var country FormatterCountry

	country,
		err = FormatterCountry{}.NewWithComponents(
		32,              // idNo
		"032",           // idString
		"Country Setup", // Desc
		"",              // Tag
		"Argentina",     // Country Culture
		"Argentina",     // Abbrv Country Name
		[]string{
			"Argentine Republic",
			"The Argentine Republic"}, // Alternate Cntry Names
		"AR",  // countryCodeTwoChar
		"ARG", // countryCodeThreeChar
		"032", // countryCodeNumber
		ePrefix.XCtx("country"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&country,
		ePrefix.XCtx(
			"&country"))

	if err != nil {
		return err
	}

	var absVal FormatterAbsoluteValue

	absVal,
		err = FormatterAbsoluteValue{}.NewBasic(
		",",                    // decimalSeparatorChars
		".",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // absoluteValFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx("absVal"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&absVal,
		ePrefix.XCtx(
			"&absVal"))

	if err != nil {
		return err
	}

	var currency FormatterCurrency

	currency,
		err = FormatterCurrency{}.NewBasic(
		",",                    // decimalSeparatorChars
		".",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"$ 127.54",             // positiveValueFmt
		"-$ 127.54",            // negativeValueFmt
		2,                      // decimalDigits
		"ARS",                  // currencyCode
		"032",                  // currencyCodeNo
		"Peso",                 // currencyName
		[]rune{'\U00000024'},   //currencySymbols
		"Centavo",              // minorCurrencyName
		make([]rune, 0),        // minorCurrencySymbols
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"currency"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&currency,
		ePrefix.XCtx(
			"&currency"))

	if err != nil {
		return err
	}

	var signedNum FormatterSignedNumber

	signedNum,
		err = FormatterSignedNumber{}.NewBasic(
		",",                    // decimalSeparatorChars
		".",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // positiveValueFmt
		"- 127.54",             // negativeValueFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"signedNum"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&signedNum,
		ePrefix.XCtx(
			"&signedNum"))

	return err
}

// Australia - Returns the number string format used in
// Australia.
//
func (nStrFmtCountry *NumStrFormatCountry) Australia(
	fmtCollection *FormatterCollection,
	errorPrefix interface{}) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	var ePrefix *ErrPrefixDto

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatCountry."+
			"Australia()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	// Adds Binary, Hexadecimal, Octal and
	// Scientific Notation Formatters
	err = numStrFormatCountryMechanics{}.ptr().
		addBaseFormatters(
			fmtCollection,
			ePrefix.XCtx("fmtCollection"))

	if err != nil {
		return err
	}

	var country FormatterCountry

	country,
		err = FormatterCountry{}.NewWithComponents(
		36,              // idNo
		"036",           // idString
		"Country Setup", // Desc
		"",              // Tag
		"Australia",     // Country Culture
		"Australia",     // Abbrv Country Name
		[]string{
			"Commonwealth of Australia",
			"The Commonwealth of Australia"}, // Alternate Cntry Names
		"AU",  // countryCodeTwoChar
		"AUS", // countryCodeThreeChar
		"036", // countryCodeNumber
		ePrefix.XCtx("country"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&country,
		ePrefix.XCtx(
			"&country"))

	if err != nil {
		return err
	}

	var absVal FormatterAbsoluteValue

	absVal,
		err = FormatterAbsoluteValue{}.NewBasic(
		".",                    // decimalSeparatorChars
		",",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // absoluteValFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx("absVal"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&absVal,
		ePrefix.XCtx(
			"&absVal"))

	if err != nil {
		return err
	}

	var currency FormatterCurrency

	currency,
		err = FormatterCurrency{}.NewBasic(
		".",                    // decimalSeparatorChars
		",",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"$ 127.54",             // positiveValueFmt
		"$ -127.54",            // negativeValueFmt
		2,                      // decimalDigits
		"AUD",                  // currencyCode
		"036",                  // currencyCodeNo
		"Dollar",               // currencyName
		[]rune{'\U00000024'},   //currencySymbols
		"Cent",                 // minorCurrencyName
		[]rune{'\U000000a2'},   // minorCurrencySymbols
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"currency"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&currency,
		ePrefix.XCtx(
			"&currency"))

	if err != nil {
		return err
	}

	var signedNum FormatterSignedNumber

	signedNum,
		err = FormatterSignedNumber{}.NewBasic(
		".",                    // decimalSeparatorChars
		",",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // positiveValueFmt
		"-127.54",              // negativeValueFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"signedNum"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&signedNum,
		ePrefix.XCtx(
			"&signedNum"))

	return err
}

// Austria - Returns the number string format used in
// Austria.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Austria(
	fmtCollection *FormatterCollection,
	errorPrefix interface{}) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	var ePrefix *ErrPrefixDto

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatCountry."+
			"Austria()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	// Adds Binary, Hexadecimal, Octal and
	// Scientific Notation Formatters
	err = numStrFormatCountryMechanics{}.ptr().
		addBaseFormatters(
			fmtCollection,
			ePrefix.XCtx("fmtCollection"))

	if err != nil {
		return err
	}

	var country FormatterCountry

	country,
		err = FormatterCountry{}.NewWithComponents(
		40,                        // idNo
		"040",                     // idString
		"Country Setup - Austria", // Desc
		"",                        // Tag
		"Austria",                 // Country Culture
		"Austria",                 // Abbrv Country Name
		[]string{
			"The Republic of Austria",
			"Republic of Austria"}, // Alternate Cntry Names
		"AT",  // countryCodeTwoChar
		"AUT", // countryCodeThreeChar
		"040", // countryCodeNumber
		ePrefix.XCtx("country"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&country,
		ePrefix.XCtx(
			"&country"))

	if err != nil {
		return err
	}

	var absVal FormatterAbsoluteValue

	absVal,
		err = FormatterAbsoluteValue{}.NewBasic(
		",",                    // decimalSeparatorChars
		".",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // absoluteValFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx("absVal"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&absVal,
		ePrefix.XCtx(
			"&absVal"))

	if err != nil {
		return err
	}

	var currency FormatterCurrency

	currency,
		err = FormatterCurrency{}.NewBasic(
		",",                    // decimalSeparatorChars
		".",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54 $",             // positiveValueFmt
		"127.54- $",            // negativeValueFmt
		2,                      // decimalDigits
		"ATS",                  // currencyCode
		"978",                  // currencyCodeNo
		"Euro",                 // currencyName
		[]rune{'\U000020ac'},   //currencySymbols
		"Cent",                 // minorCurrencyName
		make([]rune, 0),        // minorCurrencySymbols
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"currency"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&currency,
		ePrefix.XCtx(
			"&currency"))

	if err != nil {
		return err
	}

	var signedNum FormatterSignedNumber

	signedNum,
		err = FormatterSignedNumber{}.NewBasic(
		",",                    // decimalSeparatorChars
		".",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // positiveValueFmt
		"-127.54",              // negativeValueFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"signedNum"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&signedNum,
		ePrefix.XCtx(
			"&signedNum"))

	return err
}

// Bahrain - Returns the number string format used in
// The Kingdom of Bahrain.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://www.xe.com/currency/bhd-bahraini-dinar
//
func (nStrFmtCountry *NumStrFormatCountry) Bahrain(
	fmtCollection *FormatterCollection,
	errorPrefix interface{}) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	var ePrefix *ErrPrefixDto

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatCountry."+
			"Bahrain()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	// Adds Binary, Hexadecimal, Octal and
	// Scientific Notation Formatters
	err = numStrFormatCountryMechanics{}.ptr().
		addBaseFormatters(
			fmtCollection,
			ePrefix.XCtx("fmtCollection"))

	if err != nil {
		return err
	}

	var country FormatterCountry

	country,
		err = FormatterCountry{}.NewWithComponents(
		48,                        // idNo
		"048",                     // idString
		"Country Setup - Bahrain", // Desc
		"",                        // Tag
		"Bahrain",                 // Country Culture
		"Bahrain",                 // Abbrv Country Name
		[]string{
			"The Kingdom of Bahrain",
			"Kingdom of Bahrain"}, // Alternate Cntry Names
		"BH",  // countryCodeTwoChar
		"BHR", // countryCodeThreeChar
		"048", // countryCodeNumber
		ePrefix.XCtx("country"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&country,
		ePrefix.XCtx(
			"&country"))

	if err != nil {
		return err
	}

	var absVal FormatterAbsoluteValue

	absVal,
		err = FormatterAbsoluteValue{}.NewBasic(
		".",                    // decimalSeparatorChars
		",",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // absoluteValFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx("absVal"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&absVal,
		ePrefix.XCtx(
			"&absVal"))

	if err != nil {
		return err
	}

	var currency FormatterCurrency

	currency,
		err = FormatterCurrency{}.NewBasic(
		".",         // decimalSeparatorChars
		",",         // thousandsSeparatorChars
		true,        // turnOnThousandsSeparator
		"127.54 $",  // positiveValueFmt
		"127.54- $", // negativeValueFmt
		3,           // decimalDigits
		"BHD",       // currencyCode
		"048",       // currencyCodeNo
		"Dinar",     // currencyName
		[]rune{
			'B',
			'D',
		}, //currencySymbols
		"Fils",                 // minorCurrencyName
		make([]rune, 0),        // minorCurrencySymbols
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"currency"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&currency,
		ePrefix.XCtx(
			"&currency"))

	if err != nil {
		return err
	}

	var signedNum FormatterSignedNumber

	signedNum,
		err = FormatterSignedNumber{}.NewBasic(
		".",                    // decimalSeparatorChars
		",",                    // thousandsSeparatorChars
		true,                   // turnOnThousandsSeparator
		"127.54",               // positiveValueFmt
		"127.54-",              // negativeValueFmt
		-1,                     // requestedNumberFieldLen
		TextJustify(0).Right(), // numberFieldTextJustify
		ePrefix.XCtx(
			"signedNum"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&signedNum,
		ePrefix.XCtx(
			"&signedNum"))

	return err
}

// Bangladesh - Returns the number string format used in
// The People's Republic of Bangladesh.
//
// Reference Indian Numbering System:
//  https://en.wikipedia.org/wiki/Indian_numbering_system
//
func (nStrFmtCountry *NumStrFormatCountry) Bangladesh(
	fmtCollection *FormatterCollection,
	errorPrefix interface{}) (
	err error) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	var ePrefix *ErrPrefixDto

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatCountry."+
			"Bangladesh()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	// Adds Binary, Hexadecimal, Octal and
	// Scientific Notation Formatters
	err = numStrFormatCountryMechanics{}.ptr().
		addBaseFormatters(
			fmtCollection,
			ePrefix.XCtx("fmtCollection"))

	if err != nil {
		return err
	}

	var country FormatterCountry

	country,
		err = FormatterCountry{}.NewWithComponents(
		50,                           // idNo
		"050",                        // idString
		"Country Setup - Bangladesh", // Desc
		"",                           // Tag
		"Bangladesh",                 // Country Culture
		"Bangladesh",                 // Abbrv Country Name
		[]string{
			"The People's Republic of Bangladesh",
			"People's Republic of Bangladesh"}, // Alternate Cntry Names
		"BD",  // countryCodeTwoChar
		"BGD", // countryCodeThreeChar
		"050", // countryCodeNumber
		ePrefix.XCtx("country"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&country,
		ePrefix.XCtx(
			"&country"))

	if err != nil {
		return err
	}

	// Indian Numbering System
	// Example: 10,00,00,00,00,000
	numSeps := NumericSeparators{
		decimalSeparators: []rune{'.'},
		integerSeparatorsDto: NumStrIntSeparatorsDto{
			intSeparators: []NumStrIntSeparator{
				{
					intSeparatorChars:          []rune{','},
					intSeparatorGrouping:       3,
					intSeparatorRepetitions:    1,
					restartIntGroupingSequence: false,
					lock:                       nil,
				},
				{
					intSeparatorChars:          []rune{','},
					intSeparatorGrouping:       2,
					intSeparatorRepetitions:    0,
					restartIntGroupingSequence: false,
					lock:                       nil,
				},
			},
			lock: nil,
		},
	}

	numFieldDto := NumberFieldDto{
		requestedNumFieldLength: -1,
		actualNumFieldLength:    0,
		minimumNumFieldLength:   0,
		textJustifyFormat:       TextJustify(0).Right(),
		lock:                    nil,
	}

	var absVal FormatterAbsoluteValue

	absVal,
		err = FormatterAbsoluteValue{}.
		NewWithComponents(
			"127.54",
			true,
			numSeps,
			numFieldDto,
			ePrefix.XCtx("absVal"))

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&absVal,
		ePrefix.XCtx(
			"&absVal"))

	if err != nil {
		return err
	}

	var currency FormatterCurrency

	currency,
		err = FormatterCurrency{}.
		NewWithComponents(
			"$ 127.54",
			"$ -127.54",
			2,
			"BDT",
			"050",
			"Taka",
			[]rune{'\U000009f3'},
			"Paisa",
			make([]rune, 0),
			true,
			numSeps,
			numFieldDto,
			ePrefix)

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&currency,
		ePrefix.XCtx(
			"&currency"))

	if err != nil {
		return err
	}

	var signedNum FormatterSignedNumber

	signedNum,
		err = FormatterSignedNumber{}.
		NewWithComponents(
			"127.54",
			"-127.54",
			true,
			numSeps,
			numFieldDto,
			ePrefix)

	if err != nil {
		return err
	}

	err = fmtCollection.AddReplaceCollectionElement(
		&signedNum,
		ePrefix.XCtx(
			"&signedNum"))

	return err
}
