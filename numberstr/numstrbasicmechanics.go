package numberstr

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
)

type numStrBasicMechanics struct {
	lock *sync.Mutex
}

// convertInt64ToFractionalValue - Converts an int64 value to a float64 with
// all digits to the right of the decimal place.
func (nStrBasicMech *numStrBasicMechanics) convertInt64ToFractionalValue(i64 int64) (float64, error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	ex := 1
	f64 := float64(i64)
	exp := math.Pow10(ex)
	for f64 > exp {
		ex++
		exp = math.Pow10(ex)
	}

	r64 := f64 / math.Pow10(ex)

	return r64, nil

}

// convertRunesToInt64 - Converts a rune array to an int64 value.
//
func (nStrBasicMech *numStrBasicMechanics) convertRunesToInt64(
	rAry []rune, signVal int) (int64, error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	lNumRunes := len(rAry)

	if lNumRunes == 0 {
		return int64(0), errors.New("Incoming rune array is empty!\n")
	}

	if rAry[0] == '+' || rAry[0] == '-' {
		rAry = rAry[1:]
		lNumRunes = len(rAry)
	}

	numVal := int64(0)
	for i := 0; i < lNumRunes; i++ {

		if rAry[i] < '0' || rAry[i] > '9' {
			return int64(0), fmt.Errorf("Number string contained non-numeric characters: %v\n", string(rAry))
		}

		numVal *= 10
		numVal += int64(rAry[i] - 48)

	}

	numVal = numVal * int64(signVal)

	return numVal, nil

}

// ConvertStrToIntNumRunes - Receives an integer string and returns
// a slice of runes. Note, thousands separator (',') and currency signs ('$')
// will be ignored and excluded from the array of runes returned by this
// method. In order to take advantage of this exclusion, the thousands
// separator and the currency symbol must be previously initialized in
// the NumStrBasicUtility data fields.
func (nStrBasicMech *numStrBasicMechanics) convertStrToIntNumRunes(
	str string,
	thousandsSeparator rune,
	currencySymbol rune) []rune {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	baseRunes := []rune(str)
	var numRunes []rune
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false

	for i := 0; i < lBaseRunes; i++ {

		if (baseRunes[i] == '-' || baseRunes[i] == '+') &&
			len(numRunes) == 0 && isEndRunes == false &&
			i < lBaseRunes-2 &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' {

			numRunes = append(numRunes, baseRunes[i])
			isStartRunes = true

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			numRunes = append(numRunes, baseRunes[i])
			isStartRunes = true

		} else if baseRunes[i] == thousandsSeparator || baseRunes[i] == currencySymbol {

			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	return numRunes
}

// delimitDecimalCurrencyStr - Inserts a Currency Symbol and a Thousands
// Separator in a number string containing a decimal point.
func (nStrBasicMech *numStrBasicMechanics) delimitDecimalCurrencyStr(
	rawStr string,
	thousandsSeparator rune,
	decimal rune,
	currency rune) string {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	const maxStr = 256
	outStr := [maxStr]rune{}
	inStr := []rune(rawStr)
	lInStr := len(inStr)
	iCnt := 0
	outIdx := maxStr - 1
	outIdx1 := maxStr - 1
	outIdx2 := maxStr - 1
	r1 := [maxStr]rune{}
	r2 := [maxStr]rune{}
	decimalIsFound := false

	for i := lInStr - 1; i >= 0; i-- {
		if inStr[i] == decimal {
			r1[outIdx1] = decimal
			outIdx1--
			decimalIsFound = true
			continue
		}

		if !decimalIsFound {
			r1[outIdx1] = inStr[i]
			outIdx1--
		} else {
			r2[outIdx2] = inStr[i]
			outIdx2--
		}
	}

	var ptr *[maxStr]rune

	if !decimalIsFound {
		ptr = &r1
	} else {
		ptr = &r2
	}

	lInterPart := len(ptr)

	for i := lInterPart - 1; i >= 0; i-- {

		if ptr[i] >= '0' && ptr[i] <= '9' {

			iCnt++
			outStr[outIdx] = ptr[i]
			outIdx--

			if iCnt == 3 {
				iCnt = 0
				outStr[outIdx] = thousandsSeparator
				outIdx--
			}

			continue
		}

		// Check and allow for decimal
		// separators and sign designators
		if ptr[i] == '-' ||
			ptr[i] == '+' ||
			(ptr[i] == currency && currency != 0) {

			outStr[outIdx] = ptr[i]
			outIdx--

		}

	}

	if !decimalIsFound {
		return string(outStr[outIdx+1:])
	}

	return string(outStr[outIdx+1:]) + string(r1[outIdx1+1:])
}

// delimitThousands - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
// NOTE: This method will not handle number strings containing decimal fractions
// and currency characters. For these options see method ns.DlimDecCurrStr(),
// above.
func (nStrBasicMech *numStrBasicMechanics) delimitThousands(
	pureNumStr string,
	thousandsSeparator byte) string {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	const maxStr = 256
	outStr := [maxStr]byte{}
	lInStr := len(pureNumStr)
	iCnt := 0
	outIdx := maxStr - 1

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumStr[i] >= '0' && pureNumStr[i] <= '9' {

			iCnt++
			outStr[outIdx] = pureNumStr[i]
			outIdx--

			if iCnt == 3 && i != 0 {
				iCnt = 0
				outStr[outIdx] = thousandsSeparator
				outIdx--
			}

			continue
		}

		// Check and allow for sign designators
		if pureNumStr[i] == '-' ||
			pureNumStr[i] == '+' {

			outStr[outIdx] = pureNumStr[i]
			outIdx--

		}

	}

	return string(outStr[outIdx+1:])
}

// getCountryAndCurrency - Returns the Country and Currency runes as
// an instance of NumStrBasicUtility.
//
func (nStrBasicMech *numStrBasicMechanics) getCountryAndCurrency(
	country string,
	ePrefix string) (
	NumStrBasicUtility,
	error) {

	if nStrBasicMech.lock == nil {
		nStrBasicMech.lock = new(sync.Mutex)
	}

	nStrBasicMech.lock.Lock()

	defer nStrBasicMech.lock.Unlock()

	ePrefix += "numStrBasicMechanics.getCountryAndCurrency() "
	ns := NumStrBasicUtility{}

	lcStr := strings.ToLower(country)

	if strings.Contains(lcStr, "united states") {
		ns.Nation = "United States"
		ns.CurrencySymbol = NumStrCurrencySymbols[28]
		ns.ThousandsSeparator = ','
		ns.DecimalSeparator = '.'
		return ns, nil
	}

	if strings.Contains(lcStr, "united kingdom") {
		ns.Nation = "United Kingdom"
		ns.CurrencySymbol = NumStrCurrencySymbols[29]
		ns.DecimalSeparator = '.'
		return ns, nil
	}

	if strings.Contains(lcStr, "australia") {
		ns.Nation = "Australia"
		ns.CurrencySymbol = NumStrCurrencySymbols[0]
		return ns, nil
	}

	if strings.Contains(lcStr, "brazil") {
		ns.Nation = "Brazil"
		ns.CurrencySymbol = NumStrCurrencySymbols[1]
		return ns, nil
	}

	if strings.Contains(lcStr, "canada") {
		ns.Nation = "Canada"
		ns.CurrencySymbol = NumStrCurrencySymbols[2]
		ns.ThousandsSeparator = ','
		ns.DecimalSeparator = '.'
		return ns, nil
	}

	if strings.Contains(lcStr, "china") {
		ns.Nation = "China"
		ns.CurrencySymbol = NumStrCurrencySymbols[3]
		return ns, nil
	}

	if strings.Contains(lcStr, "colombia") {
		ns.Nation = "Colombia"
		ns.CurrencySymbol = NumStrCurrencySymbols[4]
		return ns, nil
	}

	if strings.Contains(lcStr, "czech") {
		ns.Nation = "Czechoslovakia"
		ns.CurrencySymbol = NumStrCurrencySymbols[5]
		return ns, nil
	}

	if strings.Contains(lcStr, "egypt") {
		ns.Nation = "Egypt"
		ns.CurrencySymbol = NumStrCurrencySymbols[6]
		return ns, nil
	}

	if strings.Contains(lcStr, "euro") {
		ns.Nation = "Euro"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "germany") {
		ns.Nation = "Germany"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "france") {
		ns.Nation = "France"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "italy") {
		ns.Nation = "Italy"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "spain") {
		ns.Nation = "Spain"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return ns, nil
	}

	if strings.Contains(lcStr, "hungary") {
		ns.Nation = "Hungary"
		ns.CurrencySymbol = NumStrCurrencySymbols[8]
		return ns, nil
	}

	if strings.Contains(lcStr, "iceland") {
		ns.Nation = "Iceland"
		ns.CurrencySymbol = NumStrCurrencySymbols[9]
		return ns, nil
	}

	if strings.Contains(lcStr, "indonesia") {
		ns.Nation = "Indonesia"
		ns.CurrencySymbol = NumStrCurrencySymbols[10]
		return ns, nil
	}

	if strings.Contains(lcStr, "israel") {
		ns.Nation = "Israel"
		ns.CurrencySymbol = NumStrCurrencySymbols[11]
		return ns, nil
	}

	if strings.Contains(lcStr, "japan") {
		ns.Nation = "Japan"
		ns.CurrencySymbol = NumStrCurrencySymbols[12]
		return ns, nil
	}

	if strings.Contains(lcStr, "korea") {
		ns.Nation = "Korea"
		ns.CurrencySymbol = NumStrCurrencySymbols[13]
		return ns, nil
	}

	if strings.Contains(lcStr, "malaysia") {
		ns.Nation = "Malaysia"
		ns.CurrencySymbol = NumStrCurrencySymbols[14]
		return ns, nil
	}

	if strings.Contains(lcStr, "mexico") {
		ns.Nation = "Mexico"
		ns.CurrencySymbol = NumStrCurrencySymbols[15]
		return ns, nil
	}

	if strings.Contains(lcStr, "norway") {
		ns.Nation = "Norway"
		ns.CurrencySymbol = NumStrCurrencySymbols[16]
		return ns, nil
	}

	if strings.Contains(lcStr, "netherlands") {
		ns.Nation = "Netherlands"
		ns.CurrencySymbol = NumStrCurrencySymbols[17]
		return ns, nil
	}

	if strings.Contains(lcStr, "pakistan") {
		ns.Nation = "Pakistan"
		ns.CurrencySymbol = NumStrCurrencySymbols[18]
		return ns, nil
	}

	if strings.Contains(lcStr, "russia") {
		ns.Nation = "Russia"
		ns.CurrencySymbol = NumStrCurrencySymbols[19]
		return ns, nil
	}

	if strings.Contains(lcStr, "saudi") {
		ns.Nation = "Saudi Arabia"
		ns.CurrencySymbol = NumStrCurrencySymbols[20]
		return ns, nil
	}

	if strings.Contains(lcStr, "south africa") {
		ns.Nation = "South Africa"
		ns.CurrencySymbol = NumStrCurrencySymbols[21]
		return ns, nil
	}

	if strings.Contains(lcStr, "sweden") {
		ns.Nation = "Sweden"
		ns.CurrencySymbol = NumStrCurrencySymbols[22]
		return ns, nil
	}

	if strings.Contains(lcStr, "switzerland") {
		ns.Nation = "Switzerland"
		ns.CurrencySymbol = NumStrCurrencySymbols[23]
		return ns, nil
	}

	if strings.Contains(lcStr, "taiwan") {
		ns.Nation = "Taiwan"
		ns.CurrencySymbol = NumStrCurrencySymbols[24]
		return ns, nil
	}

	if strings.Contains(lcStr, "turkey") {
		ns.Nation = "Turkey"
		ns.CurrencySymbol = NumStrCurrencySymbols[25]
		return ns, nil
	}

	if strings.Contains(lcStr, "venezuela") {
		ns.Nation = "Venezuela"
		ns.CurrencySymbol = NumStrCurrencySymbols[26]
		return ns, nil
	}

	if strings.Contains(lcStr, "viet nam") {
		ns.Nation = "Viet Nam"
		ns.CurrencySymbol = NumStrCurrencySymbols[27]
		return ns, nil
	}

	return ns,
		fmt.Errorf(ePrefix+"\n"+
			"Failed to initialize country, %v.\n", country)

}
