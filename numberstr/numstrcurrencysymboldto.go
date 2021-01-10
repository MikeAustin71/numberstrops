package numberstr

import "sync"

type CurrencySymbolDto struct {
	idNo                          int    // Arbitrary Id Number used to track this instance
	currencySymbol                rune   // Unicode UTF-8 currency symbol
	currencyDecimalDigits         int    // The number of digits after the decimal separator for currency presentations
	turnOnIntegerDigitsSeparation bool   // Turns on thousands separator for currency presentations
	currencyCode                  string // Wold Currency Code (3-Characters). Reference: http://www.xe.com/symbols.php
	currencyName                  string // The common name of this currency i.e 'Dollar', 'Yuan' etc.
	countryName                   string // Full Name of country associated with this currency
	abbreviatedCountryName        string // Abbreviated Country Name
	alternateCountryName          string // Alternate Country Name Example: United States of America
	countryCodeTwoChar            string // ISO-3166 Two Character Country Code https://en.wikipedia.org/wiki/ISO_3166-2
	countryCodeThreeChar          string // ISO-3166 Three Character Country Code https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	countryCodeNumber             string // ISO-3166 Country Code Number: 3-Numeric Digits https://en.wikipedia.org/wiki/ISO_3166-1_numeric
	positiveValCurrFormat         string
	// Positive Value Currency Format Samples
	//  Positive Value Format Codes
	//      127.54 = Placeholder for the actual un-formatted Number.
	//
	//    NUMFIELD = Placeholder for the number field in which the
	//               number string will be displayed. The number field
	//               is usually wider than the actual formatted number
	//               and is right justified within the number string.
	//
	//           + = Plus Sign Placement if specified
	//
	//           $ = Currency Symbol Placement
	//
	//    $127.54
	//    $ 127.54
	//    127.54$
	//    127.54 $
	//
	//    $127.54+
	//    $127.54 +
	//    $ 127.54+
	//    $ 127.54 +
	//
	//    $+127.54
	//    $ +127.54
	//    $+ 127.54
	//    $ + 127.54
	//
	//    127.54$+
	//    127.54$ +
	//    127.54 $+
	//    127.54$ +
	//    127.54 $ +
	//
	//    127.54+$
	//    127.54+ $
	//    127.54 +$
	//    127.54+ $
	//    127.54 + $
	//
	//    +127.54$
	//    +127.54 $
	//    + 127.54$
	//    + 127.54 $
	//
	//    $NUMFIELD
	//    $ NUMFIELD
	//    NUMFIELD$
	//    NUMFIELD $
	//
	//    $NUMFIELD+
	//    $NUMFIELD +
	//    $ NUMFIELD+
	//    $ NUMFIELD +
	//
	//    $+NUMFIELD
	//    $ +NUMFIELD
	//    $+ NUMFIELD
	//    $ + NUMFIELD
	//
	//    NUMFIELD$+
	//    NUMFIELD$ +
	//    NUMFIELD $+
	//    NUMFIELD$ +
	//    NUMFIELD $ +
	//
	//    NUMFIELD+$
	//    NUMFIELD+ $
	//    NUMFIELD +$
	//    NUMFIELD+ $
	//    NUMFIELD + $
	//
	//    +NUMFIELD$
	//    +NUMFIELD $
	//    + NUMFIELD$
	//    + NUMFIELD $

	negativeValCurrFormat string
	// Negative Value Currency Format Samples
	//      127.54 = The actual un-formatted number
	//    NUMFIELD = The number field in which the
	//               number string will be displayed.
	//               The number field is usually wider
	//               than the actual formatted number
	//               and is right justified within the
	//               number string.
	//           - = Minus Sign Placement
	//           $ = Currency Symbol Placement
	//          () = Parentheses Placement
	//         (-) = Parentheses & Minus Sign Placement
	//
	//   (-) $127.54
	//   (-) $ 127.54
	//   (-)$127.54
	//   (-)$ 127.54
	//   $ (-)127.54
	//   $ (-) 127.54
	//   $(-)127.54
	//   $(-) 127.54
	//   (-) 127.54$
	//   (-) 127.54 $
	//   (-)127.54$
	//   (-)127.54 $
	//   127.54(-) $
	//   127.54 (-) $
	//   127.54(-)$
	//   127.54 (-)$
	//   127.54$(-)
	//   127.54$ (-)
	//   127.54 $ (-)
	//   127.54 $(-)
	//   $127.54(-)
	//   $127.54 (-)
	//   $ 127.54(-)
	//   $ 127.54 (-)
	//    - $127.54
	//    - $ 127.54
	//    -$127.54
	//    -$ 127.54
	//    $ -127.54
	//    $ - 127.54
	//    $-127.54
	//    $- 127.54
	//    - 127.54$
	//    - 127.54 $
	//    -127.54$
	//    -127.54 $
	//    127.54- $
	//    127.54 - $
	//    127.54-$
	//    127.54 -$
	//    127.54$-
	//    127.54$ -
	//    127.54 $ -
	//    127.54 $-
	//    $127.54-
	//    $127.54 -
	//    $ 127.54-
	//    $ 127.54 -
	//   ( $127.54 )
	//   ( $ 127.54 )
	//   ($ 127.54)
	//   ($127.54)
	//    $(127.54)
	//    $ (127.54)
	//    $( 127.54 )
	//    $ ( 127.54 )
	//   ( 127.54$ )
	//   ( 127.54 $ )
	//   ( 127.54 $)
	//   (127.54$)
	//   (127.54)$
	//   (127.54) $
	//   ( 127.54 )$
	//   ( 127.54 ) $
	//   (-) $NUMFIELD
	//   (-) $ NUMFIELD
	//   (-)$NUMFIELD
	//   (-)$ NUMFIELD
	//   $ (-)NUMFIELD
	//   $ (-) NUMFIELD
	//   $(-)NUMFIELD
	//   $(-) NUMFIELD
	//   (-) NUMFIELD$
	//   (-) NUMFIELD $
	//   (-)NUMFIELD$
	//   (-)NUMFIELD $
	//   NUMFIELD(-) $
	//   NUMFIELD (-) $
	//   NUMFIELD(-)$
	//   NUMFIELD (-)$
	//   NUMFIELD$(-)
	//   NUMFIELD$ (-)
	//   NUMFIELD $ (-)
	//   NUMFIELD $(-)
	//   $NUMFIELD(-)
	//   $NUMFIELD (-)
	//   $ NUMFIELD(-)
	//   $ NUMFIELD (-)
	//    - $NUMFIELD
	//    - $ NUMFIELD
	//    -$NUMFIELD
	//    -$ NUMFIELD
	//    $ -NUMFIELD
	//    $ - NUMFIELD
	//    $-NUMFIELD
	//    $- NUMFIELD
	//    - NUMFIELD$
	//    - NUMFIELD $
	//    -NUMFIELD$
	//    -NUMFIELD $
	//    NUMFIELD- $
	//    NUMFIELD - $
	//    NUMFIELD-$
	//    NUMFIELD -$
	//    NUMFIELD$-
	//    NUMFIELD$ -
	//    NUMFIELD $ -
	//    NUMFIELD $-
	//    $NUMFIELD-
	//    $NUMFIELD -
	//    $ NUMFIELD-
	//    $ NUMFIELD -
	//   ( $NUMFIELD )
	//   ( $ NUMFIELD )
	//   ($ NUMFIELD)
	//   ($NUMFIELD)
	//    $(NUMFIELD)
	//    $ (NUMFIELD)
	//    $( NUMFIELD )
	//    $ ( NUMFIELD )
	//   ( NUMFIELD$ )
	//   ( NUMFIELD $ )
	//   ( NUMFIELD $)
	//   (NUMFIELD$)
	//   (NUMFIELD)$
	//   (NUMFIELD) $
	//   ( NUMFIELD )$
	//   ( NUMFIELD ) $

	lock *sync.Mutex // Used to coordinate multi-threaded operations
}

// CopyIn - Receives a pointer to a CurrencySymbolDto object and
// copies the internal data fields to the current CurrencySymbolDto
// instance.
//
// All of the data fields in the current CurrencySymbolDto will be
// overwritten.
//
func (curSymDto *CurrencySymbolDto) CopyIn(
	currSymDtoIn *CurrencySymbolDto,
	ePrefix string) (err error) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	ePrefix += "CurrencySymbolDto.CopyIn() "

	curSymMech := currencySymbolMechanics{}

	err = curSymMech.copyIn(
		curSymDto,
		currSymDtoIn,
		ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current CurrencySymbolDto
// instance.
//
func (curSymDto *CurrencySymbolDto) CopyOut() CurrencySymbolDto {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymMech := currencySymbolMechanics{}

	var newCurrSymDto = CurrencySymbolDto{}

	newCurrSymDto, _ = curSymMech.copyOut(curSymDto, "")

	return newCurrSymDto
}

// *******************************************************
//               GETTER METHODS BEGIN HERE
// Use 'Getter' methods to retrieve information stored in
// 'CurrencySymbolDto' data fields for the current instance.
//
// *******************************************************

// GetAbbreviatedCountryName - Gets the abbreviated country name
// associated with this CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) GetAbbreviatedCountryName() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.abbreviatedCountryName

}

// GetAlternateCountryName - Gets the alternate country name
// associated with this CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) GetAlternateCountryName() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.alternateCountryName

}

// GetCountryCodeNumber - Gets the three character country
// code number for the country associated with this CurrencySymbolDto
// instance. The country Code Number consists of three numeric
// digits. ISO-3166 Country Code Number: 3-Numeric Digits.
// Country Code Number may have leading zero digits.
//
// Reference ISO-3166 country codes at:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) GetCountryCodeNumber() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.countryCodeNumber

}

// GetCountryCodeThreeChar - Gets the three character country
// code for the country associated with this CurrencySymbolDto
// instance.
//
// Reference ISO-3166 country codes at:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) GetCountryCodeThreeChar() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.countryCodeThreeChar

}

// GetCountryCodeTwoChar - Gets the two character country
// code for the country associated with this CurrencySymbolDto
// instance.
//
//
// Reference ISO-3166 country codes at:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) GetCountryCodeTwoChar() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.countryCodeTwoChar

}

// GetCountryName - Gets the Country Name associated with
// this CurrencySymbolDto instance.
//
// Reference:
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) GetCountryName() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.countryName

}

// GetCurrencyCode - Gets the currency code used to identify the
// currency symbol associated with this CurrencySymbolDto instance.
//
// Information on Wold Currency Symbols and Currency codes is
// maintained at:
//    https://www.xe.com/symbols.php
//
func (curSymDto *CurrencySymbolDto) GetCurrencyCode() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.currencyCode

}

// GetCurrencyName - Gets the currency name which is associated
// with this CurrencySymbolDto instance.
//
// Information on Wold Currency Symbols and Currency codes is
// maintained at:
//    https://www.xe.com/symbols.php
//
func (curSymDto *CurrencySymbolDto) GetCurrencyName() string {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.currencyName

}

// GetCurrencySymbol - Gets the currency symbol which is associated
// with this CurrencySymbolDto instance. The currency symbol is stored
// as a type 'rune' or UTF-8 (Unicode) symbol.
//
// Information on Wold Currency Symbols and Currency codes is
// maintained at:
//    https://www.xe.com/symbols.php
//
func (curSymDto *CurrencySymbolDto) GetCurrencySymbol() rune {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.currencySymbol

}

// GetIdNo - Gets the Id Number used to identify this
// CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) GetIdNo() int {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	return curSymDto.idNo

}

// IsValidInstance - If the current instance of CurrencySymbolDto
// is valid, this method returns 'true'. If the current instance
// is judged to be invalid, this method returns 'false'.
//
func (curSymDto *CurrencySymbolDto) IsValidInstance() (
	isValid bool) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymMech := currencySymbolMechanics{}

	isValid,
		_ = curSymMech.testCurrencySymbolDtoValidity(curSymDto,
		"CurrencySymbolDto.IsValidInstance() ")

	return isValid
}

// IsValidInstanceError - Examines the current CurrencySymbolDto
// instance to determine if any of the member data elements are
// invalid.
//
// If the current CurrencySymbolDto is found to be invalid, a
// detailed error message is returned. If the instance is valid,
// the returned error value is set to 'nil'.
//
func (curSymDto *CurrencySymbolDto) IsValidInstanceError(ePrefix string) error {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	ePrefix += "CurrencySymbolDto.IsValidInstanceError() "
	curSymMech := currencySymbolMechanics{}

	_,
		err := curSymMech.testCurrencySymbolDtoValidity(curSymDto,
		ePrefix)

	return err
}

// *******************************************************
//               SETTER METHODS BEGIN HERE
// Use 'Setter' methods to set and store the data field
// values for the current 'CurrencySymbolDto' instance.
//
// *******************************************************

// SetAbbreviatedCountryName - Sets the abbreviated country name
// associated with this CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) SetAbbreviatedCountryName(abbreviatedCountryName string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.abbreviatedCountryName = abbreviatedCountryName

}

// SetAlternateCountryName - Sets the alternate country name
// associated with this CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) SetAlternateCountryName(alternateCountryName string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.alternateCountryName = alternateCountryName

}

// SetCountryCodeNumber - Sets the three character country
// code number for the country associated with this CurrencySymbolDto
// instance. The country Code Number consists of three numeric
// digits. ISO-3166 Country Code Number: 3-Numeric Digits.
//
// Reference ISO-3166 country codes at:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) SetCountryCodeNumber(
	threeCharacterCountryCodeNo string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.countryCodeNumber = threeCharacterCountryCodeNo

}

// SetCountryCodeThreeChar - Sets the three character country
// code for the country associated with this CurrencySymbolDto
// instance.
//
// Reference ISO-3166 country codes at:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) SetCountryCodeThreeChar(threeCharacterCountryCode string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.countryCodeThreeChar = threeCharacterCountryCode

}

// SetCountryCodeTwoChar - Sets the two character country
// code for the country associated with this CurrencySymbolDto
// instance.
//
//
// Reference ISO-3166 country codes at:
//   https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//
func (curSymDto *CurrencySymbolDto) SetCountryCodeTwoChar(twoCharacterCountryCode string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.countryCodeTwoChar = twoCharacterCountryCode

}

// SetCurrencyCode - Sets the currency code used to identify the
// currency symbol associated with this CurrencySymbolDto instance.
//
// Information on Wold Currency Symbols and Currency codes is
// maintained at:
//    https://www.xe.com/symbols.php
//
func (curSymDto *CurrencySymbolDto) SetCurrencyCode(currencyCode string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.currencyCode = currencyCode

}

// SetCurrencySymbol - Sets the currency symbol which is associated
// with this CurrencySymbolDto instance. The currency symbol is stored
// as a type 'rune' or UTF-8 (Unicode) symbol.
//
// Information on Wold Currency Symbols and Currency codes is
// maintained at:
//    https://www.xe.com/symbols.php
//
func (curSymDto *CurrencySymbolDto) SetCurrencyName(currencyName string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.currencyName = currencyName

}

// SetCurrencyName - Sets the currency symbol which is associated
// with this CurrencySymbolDto instance. The currency symbol is stored
// as a type 'rune' or UTF-8 (Unicode) symbol.
//
// Information on Wold Currency Symbols and Currency codes is
// maintained at:
//    https://www.xe.com/symbols.php
//
func (curSymDto *CurrencySymbolDto) SetCurrencySymbol(currencySymbol rune) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.currencySymbol = currencySymbol

}

// SetCountryName - Sets the Country Name associated with
// this CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) SetCountryName(countryName string) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.countryName = countryName

}

// SetIdNo - Sets the Id Number used to identify this
// CurrencySymbolDto instance.
//
func (curSymDto *CurrencySymbolDto) SetIdNo(idNo int) {

	if curSymDto.lock == nil {
		curSymDto.lock = new(sync.Mutex)
	}

	curSymDto.lock.Lock()

	defer curSymDto.lock.Unlock()

	curSymDto.idNo = idNo

}
