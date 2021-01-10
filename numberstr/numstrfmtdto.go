package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFormatDto struct {
	valueDisplaySpec         NumStrValSpec
	positiveValueFmt         string
	negativeValueFmt         string
	currencyFmt              CurrencySymbolDto
	decimalSeparator         rune
	thousandsSeparator       rune
	turnOnThousandsSeparator bool
	numberFieldDto           numberFieldDto
	lock                     *sync.Mutex
}

// copyIn - Receives pointers to an instance of NumStrFormatDto
// labeled as input parameter, 'nStrFmtDtoIn'. This method then
// proceeds to copy all of the data from 'nStrFmtDtoIn' into the
// current NumStrFormatDto instance.  When the copy process is
// completed, the data values in 'nStrFmtDtoIn' and the current
// NumStrFormatDto instance will be identical.
//
// IMPORTANT
//
// Be advised, this method will overwrite the data values contained
// in the current instance of NumStrFormatDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtDtoIn        *NumStrFormatDto
//     - A pointer to an instance of NumStrFormatDto. This method
//       will copy all of the data values contained in
//       'numStrFmtDtoIn' to the current instance of
//       NumStrFormatDto.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. This
//       error prefix, 'ePrefix' will be prefixed to the beginning
//       of the error message.
//
func (nStrFmtDto *NumStrFormatDto) CopyIn(
	nStrFmtDtoIn *NumStrFormatDto,
	ePrefix string) (err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.CopyIn() "

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	err =
		nStrFmtDtoMech.copyIn(
			nStrFmtDto,
			nStrFmtDtoIn,
			ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current NumStrFormatDto
// instance.
//
func (nStrFmtDto *NumStrFormatDto) CopyOut() NumStrFormatDto {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	newNumStrFmtDto,
		_ :=
		nStrFmtDtoMech.copyOut(
			nStrFmtDto,
			"")

	return newNumStrFmtDto
}

// GetValueDisplaySpec - Returns the NumStrValSpec
// value.
//
func (nStrFmtDto *NumStrFormatDto) GetValueDisplaySpec() (
	valueDisplaySpec NumStrValSpec) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	valueDisplaySpec = nStrFmtDto.valueDisplaySpec

	return valueDisplaySpec
}

// GetPositiveValueFormat - Returns the Positive Value Format
// string. This string contains formatting characters and placeholders
// used to format positive numeric values in number strings.
//
func (nStrFmtDto *NumStrFormatDto) GetPositiveValueFormat() string {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	return nStrFmtDto.positiveValueFmt
}

// GetNegativeValueFormat - Returns the Negative Value Format
// string. This string contains formatting characters and placeholders
// used to format negative numeric values in number strings.
//
func (nStrFmtDto *NumStrFormatDto) GetNegativeValueFormat() string {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	return nStrFmtDto.negativeValueFmt
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFormatDto instance to determine whether the current instance
// is a valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If this method completes successfully, the returned boolean
//       value will signal whether the current NumStrFormatDto is
//       valid, or not. If the current NumStrFormatDto contains
//       valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
//
func (nStrFmtDto *NumStrFormatDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	isValid,
		_ =
		nStrFmtDtoMech.testNumStrFormatDtoValidity(
			nStrFmtDto,
			"")

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFormatDto instance to determine whether the current instance
// is a valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. Said text will be prefixed
//       to the beginning of the error message.
//
//       If this instance of NumStrFormatDto contains invalid data, a
//       detailed error message will be returned identifying the invalid
//       data item.
//
func (nStrFmtDto *NumStrFormatDto) IsValidInstanceError(
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.IsValidInstanceError() "

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	_,
		err =
		nStrFmtDtoMech.testNumStrFormatDtoValidity(
			nStrFmtDto,
			ePrefix)

	return err
}

// New - Creates and returns a new instance of
// NumStrFormatDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  valueDisplaySpec    NumStrValSpec
//     - An enumeration value for Number String Value Specification.
//       Used to specify display of positive and and negative values,
//       or absolute values.
//
//
//  positiveValueFmt           string
//     - A string specifying the number string format to be used in
//       formatting positive numeric values in text strings. Valid
//       formats for positive numeric values (NOT Currency) are listed
//       as follows:
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Positive Value Format
//
//       Positive Value Formatting Terminology:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the numeric value of a number
//                     string. This place holder signals that the
//                     actual length of the numeric value including
//                     formatting characters and symbols such  as
//                     Thousands Separators, Decimal Separators and
//                     Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format string,
//                     the plus sign ('+') specifies  where the plus sign will
//                     be placed for positive numeric values.
//
//    Absence of "+" - The absence of a plus sign ('+') means that the positive
//                     numeric value will be displayed in text with out a
//                     plus sign ('+'). This is the default for positive number
//                     formatting.
//
//
//  negativeValueFmt           string
//     - An enumeration value for Number String Negative Value Format
//       Mode. Used to specify formatting for negative numeric values
//       (NOT Currency). Valid formats for negative numeric values are
//       listed as follows:
//
//               -127.54   The Default Negative Value Format String
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -127.54
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -NUMFIELD
//               - NUMFIELD
//               NUMFIELD-
//               NUMFIELD -
//               (-) NUMFIELD
//               (-)NUMFIELD
//               NUMFIELD(-)
//               NUMFIELD (-)
//               (NUMFIELD)
//               ( NUMFIELD )
//               (NUMFIELD)
//               ( NUMFIELD )
//
//       Negative Value Formatting Terminology:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the numeric value of a number
//                     string. This place holder signals that the
//                     actual length of the numeric value including
//                     formatting characters and symbols such  as
//                     Thousands Separators, Decimal Separators and
//                     Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the format string,
//                     the minus sign ('-') specifies where the minus sign will
//                     be positioned in the text string containing the negative
//                     numeric value.
//
//             "(-)" - These three characters are often used in Europe and the
//                     United Kingdom to classify a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are frequently used in
//                     the United States of America to classify a numeric value
//                     as negative.
//
//
//  currencyFmt                CurrencySymbolDto
//     - A valid, populated instance of CurrencySymbolDto. This
//       structure contains all the information necessary to format
//       currency symbols. If this parameter is unpopulated or
//       invalid, an error will be returned.
//
//
//  decimalSeparator           rune
//     - This parameter holds the character used to separate the
//       integer and fractional components of a floating point
//       number string. In the United States, the standard decimal
//       separator is the decimal point ('.').
//          Example:  123.456
//
//
//  thousandsSeparator         rune
//     - This parameter holds the character used to separate thousands
//       in the integer component of a number string. In the United
//       States, the standard thousands separator is the comma.
//         Example:  1,000,000,000
//
//
//  turnOnThousandsSeparator   bool
//     - Simply setting the Thousands Separator character will not
//       ensure that character is actually used in formatting number
//       strings. In addition, it is necessary to activate the use of
//       the designated thousands character in formatting text number
//       strings.
//
//       To turn on the insertion of thousands separators in the
//       formatting of number strings, set input parameter
//       'turnOnThousandsSeparator' to 'true'.
//
//
//  numberFieldLength          int
//     - This is the requested length of the number field in which
//       the number string will be displayed. If this field length
//       is greater than the actual length of the number string,
//       the number string will be right justified within the
//       the number field. If the actual number string length is
//       greater than the number field length, the number field
//       field length will be automatically expanded to display the
//       entire number string.
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newFmtDto                  NumStrFormatDto
//     - If this method completes successfully, a new instance of NumStrFormatDto
//       be returned to calling function.
//
//
//  err                        error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDto NumStrFormatDto) New(
	valueDisplaySpec NumStrValSpec,
	positiveValueFmt string,
	negativeValueFmt string,
	currencyFmt CurrencySymbolDto,
	decimalSeparator rune,
	thousandsSeparator rune,
	turnOnThousandsSeparator bool,
	numberFieldLength int,
	ePrefix string) (
	newFmtDto NumStrFormatDto,
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.New() "

	if !valueDisplaySpec.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'valueDisplaySpec' is invalid!\n"+
			"valueDisplaySpec='%v'\n",
			nStrFmtDto.valueDisplaySpec.XValueInt())
		return newFmtDto, err
	}

	if decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is a zero value rune!\n"+
			"Invalid 'decimalSeparator' value!\n", ePrefix)
		return newFmtDto, err
	}

	if thousandsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparator' is a zero value rune!\n"+
			"Invalid 'thousandsSeparator' value!\n", ePrefix)
		return newFmtDto, err
	}

	nStrFmtNanobot := numStrFormatNanobot{}

	_,
		err = nStrFmtNanobot.testNumStrFormatValidity(
		positiveValueFmt,
		ePrefix+"Testing positiveValueFmt validity. ")

	if err != nil {
		return newFmtDto, err
	}

	_,
		err = nStrFmtNanobot.testNumStrFormatValidity(
		negativeValueFmt,
		ePrefix+"Testing negativeValueFmt validity. ")

	if err != nil {
		return newFmtDto, err
	}

	err = currencyFmt.IsValidInstanceError(
		ePrefix +
			"Testing 'currencyFmt' validity. ")

	if err != nil {
		return newFmtDto, err
	}

	newFmtDto.valueDisplaySpec = valueDisplaySpec
	newFmtDto.negativeValueFmt = negativeValueFmt
	newFmtDto.positiveValueFmt = positiveValueFmt
	newFmtDto.currencyFmt = currencyFmt.CopyOut()
	newFmtDto.turnOnThousandsSeparator = turnOnThousandsSeparator
	newFmtDto.thousandsSeparator = thousandsSeparator
	newFmtDto.decimalSeparator = decimalSeparator
	newFmtDto.numberFieldDto.requestedNumFieldLength = numberFieldLength

	newFmtDto.lock = new(sync.Mutex)

	return newFmtDto, err
}

// SetCurrencyFormat - Controls the use of Currency Symbols
// in formatted number strings.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  turnOnCurrencyFormatting   bool
//     - If this input parameter is set to 'true' the currency symbol
//       designated in parameter 'currencyFormat' will be used to insert
//       that currency symbol in formatted number strings. If set to
//       'false', the currency symbol WILL NOT be inserted into formatted
//       number strings.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtDto *NumStrFormatDto) SetCurrencyFormat(
	turnOnCurrencyFormatting bool) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	if turnOnCurrencyFormatting == true {

		nStrFmtDto.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()

		return
	}

	// turnOnCurrencyFormatting must be 'false'

	if nStrFmtDto.valueDisplaySpec == NumStrValSpec(0).CurrencyValue() {

		nStrFmtDto.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()

	}

	return
}

// SetCurrencyFormat - Sets the currency symbol and determines
// whether currency symbols will be formatted and displayed in
// text number strings.
//
// Simply setting the currency symbol will not ensure that the
// currency symbol is actually used in the formatting of text
// number strings.
//
// To ensure the use of the currency symbol in formatting text
// number strings, set the second input parameter to 'true'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  currencyFormat             CurrencySymbolDto
//     - If this instance of CurrencySymbolDto is valid, it will be
//       used to configure the currency symbol and related currency
//      information in the current NumStrFormatDto instance.
//
//
//  turnOnCurrencyFormatting   bool
//     - If this input parameter is set to 'true' the currency symbol
//       designated in parameter 'currencyFormat' will be used to insert
//       that currency symbol in formatted number strings. If set to
//       'false', the currency symbol WILL NOT be inserted into formatted
//       number strings.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. This
//       error prefix, 'ePrefix' will be prefixed to the beginning
//       of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetCurrencySymbol(
	currencyFormat CurrencySymbolDto,
	turnOnCurrencyFormatting bool,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetNegativeValueFormat() "

	err = currencyFormat.IsValidInstanceError(
		ePrefix +
			"Testing validity of input parameter 'currencyFormat'. ")

	if err != nil {
		return err
	}

	nStrFmtDto.currencyFmt = currencyFormat.CopyOut()

	if turnOnCurrencyFormatting == true {

		nStrFmtDto.valueDisplaySpec = NumStrValSpec(0).CurrencyValue()

		return err
	}

	// turnOnCurrencyFormatting must be 'false'

	if nStrFmtDto.valueDisplaySpec == NumStrValSpec(0).CurrencyValue() {
		nStrFmtDto.valueDisplaySpec = NumStrValSpec(0).SignedNumberValue()
	}

	return err
}

// SetDecimalSeparator - Sets the character which will be used to
// separate the integer and fractional components of a number
// string.
//
// In the United States, the standard decimal separator is the the
// decimal point ('.') or period character.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparator    rune
//     - If this character is a non-zero value, it will be used as
//       the decimal separator in formatting floating point number
//       strings. If this parameter is set to zero ('0'), an error
//       will be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. This
//       error prefix, 'ePrefix' will be prefixed to the beginning
//       of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetDecimalSeparator(
	decimalSeparator rune,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetDecimalSeparator() "

	if decimalSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' contains a zero value.\n"+
			"'decimalSeparator' is invalid!\n",
			ePrefix)
		return err
	}

	nStrFmtDto.decimalSeparator = decimalSeparator

	return
}

// SetNegativeValueFormat - Sets the format string used to
// format negative values in number strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  negativeValueFmt           string
//     - An enumeration value for Number String Negative Value Format
//       Mode. Used to specify formatting for negative numeric values
//       (NOT Currency). Valid formats for negative numeric values are
//       listed as follows:
//
//               -127.54   The Default Negative Value Format String
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -127.54
//               - 127.54
//               127.54-
//               127.54 -
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//               (127.54)
//               ( 127.54 )
//               (127.54)
//               ( 127.54 )
//               -NUMFIELD
//               - NUMFIELD
//               NUMFIELD-
//               NUMFIELD -
//               (-) NUMFIELD
//               (-)NUMFIELD
//               NUMFIELD(-)
//               NUMFIELD (-)
//               (NUMFIELD)
//               ( NUMFIELD )
//               (NUMFIELD)
//               ( NUMFIELD )
//
//       Negative Value Formatting Terminology:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the numeric value of a number
//                     string. This place holder signals that the
//                     actual length of the numeric value including
//                     formatting characters and symbols such  as
//                     Thousands Separators, Decimal Separators and
//                     Currency Symbols.
//
//               "-" - The Minus Sign ('-'). If present in the format string,
//                     the minus sign ('-') specifies where the minus sign will
//                     be positioned in the text string containing the negative
//                     numeric value.
//
//             "(-)" - These three characters are often used in Europe and the
//                     United Kingdom to classify a numeric value as negative.
//
//              "()" - Opposing parenthesis characters are frequently used in
//                     the United States of America to classify a numeric value
//                     as negative.
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetNegativeValueFormat(
	negativeValueFmt string,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetNegativeValueFormat() "

	nStrFmtNanobot := numStrFormatNanobot{}

	_,
		err = nStrFmtNanobot.testNumStrFormatValidity(
		negativeValueFmt,
		ePrefix+
			fmt.Sprintf("negativeValueFmt='%v' ",
				negativeValueFmt))

	if err != nil {
		return err
	}

	nStrFmtDto.negativeValueFmt = negativeValueFmt

	return err
}

// SetPositiveValueFormatMode - Sets the Positive Value Format
// for the current NumStrFormatDto instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  positiveValueFmt           string
//     - A string specifying the number string format to be used in
//       formatting positive numeric values in text strings. Valid
//       formats for positive numeric values (NOT Currency) are listed
//       as follows:
//               "+NUMFIELD"
//               "+ NUMFIELD"
//               "NUMFIELD+"
//               "NUMFIELD +"
//               "NUMFIELD"
//               "+127.54"
//               "+ 127.54"
//               "127.54+"
//               "127.54 +"
//               "127.54" THE DEFAULT Positive Value Format
//
//       Positive Value Formatting Terminology:
//
//        "NUMFIELD" - Placeholder for a number field. A number field has
//                     a string length which is equal to or greater than
//                     the actual numeric value string length. Actual
//                     numeric values are right justified within number
//                     fields for text displays.
//
//          "127.54" - Place holder for the numeric value of a number
//                     string. This place holder signals that the
//                     actual length of the numeric value including
//                     formatting characters and symbols such  as
//                     Thousands Separators, Decimal Separators and
//                     Currency Symbols.
//
//               "+" - The Plus Sign ('+'). If present in the format string,
//                     the plus sign ('+') specifies  where the plus sign will
//                     be placed for positive numeric values.
//
//    Absence of "+" - The absence of a plus sign ('+') means that the positive
//                     numeric value will be displayed in text with out a
//                     plus sign ('+'). This is the default for positive number
//                     formatting.
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetPositiveValueFormatMode(
	positiveValueFmt string,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetNegativeValueFormat() "

	nStrFmtNanobot := numStrFormatNanobot{}

	_,
		err = nStrFmtNanobot.testNumStrFormatValidity(
		positiveValueFmt,
		ePrefix+
			fmt.Sprintf("positiveValueFmt='%v' ",
				positiveValueFmt))

	if err != nil {
		return err
	}

	nStrFmtDto.positiveValueFmt = positiveValueFmt

	return err
}

// SetNumberFieldLength - Sets the length of the number field in
// which the formatted number string will be displayed.
//
// If the number field length is wider than the actual length of
// the formatted number string, that number string will be right
// justified in the number string. The left margin of the number
// field will be padded with spaces.
//
// Example:
//   numFieldLen = 9
//   numberString =          "123.456"
//   Output Formatted Text = "  123.456"
//
// If the number field length is equal to or less than the actual
// length of the formatted number string, the number string will
// will be displayed in it's actual length.
//
// Example:
//   numFieldLen = -1
//   numberString =          "123.456"
//   Output Formatted Text = "123.456"
//
func (nStrFmtDto *NumStrFormatDto) SetNumberFieldLength(
	numFieldLen int) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDto.numberFieldDto.requestedNumFieldLength =
		numFieldLen
}

// SetThousandsSeparatorCharacter - Sets the value of the thousands
// separator to the rune passed as an input parameter.
//
// Simply setting the Thousands Separator value in the current
// NumStrFormatDto will not actually display that thousands
// separator in text number strings. In addition, the use of
// a thousands separator must be activated by setting the second
// input parameter to 'true'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  thousandsSeparator         rune
//     - This character will be used to separate thousands in number
//       strings formatted for text display.
//       Example:
//         thousandsSeparator = ','   Output: 1,000,000,000
//
//
//  turnOnThousandsSeparator   bool
//     - Simply setting the Thousands Separator character will not
//       ensure that character is actually used in formatting number
//       strings. In addition, it is necessary to activate the use of
//       the designated thousands character in formatting text number
//       strings.
//
//       To turn on the insertion of thousands separators in the
//       formatting of number strings, set input parameter
//       'turnOnThousandsSeparator' to 'true'.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. This
//       error prefix, 'ePrefix' will be prefixed to the beginning
//       of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetThousandsSeparator(
	thousandsSeparator rune,
	turnOnThousandsSeparator bool,
	ePrefix string) (err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetThousandsSeparatorCharacter() "

	if thousandsSeparator == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparator' is a zero value.\n"+
			" 'thousandsSeparator' is invalid!\n", ePrefix)
		return err
	}

	nStrFmtDto.thousandsSeparator = thousandsSeparator
	nStrFmtDto.turnOnThousandsSeparator = turnOnThousandsSeparator

	return err
}

// SetThousandsSeparatorOn - Call this method to control the display
// of thousands separators in text number strings.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  turnOnThousandsSeparators  bool
//     - If this input parameter is set to 'true', thousands separators
//       be inserted to number strings for text displays.
//         Example: 1,000,000,000
//
//       If this input parameter is set to 'false' thousands separators,
//       WILL NOT be included in number strings formatted for text display.
//         Example: 1000000000
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (nStrFmtDto *NumStrFormatDto) SetThousandsSeparatorDisplay(
	turnOnThousandsSeparators bool) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDto.turnOnThousandsSeparator = turnOnThousandsSeparators

}

// SetToDefaults - This method will set all internal
// member variables to their default values.
//
// The Number String Format Defaults represent formatting
// parameters used in the United States.
//
func (nStrFmtDto *NumStrFormatDto) SetToDefaults() {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoUtil := numStrFormatDtoUtility{}

	_ = nStrFmtDtoUtil.setToDefaults(
		nStrFmtDto,
		"")

}

// SetToDefaultsIfEmpty - If the current NumStrFormatDto instance
// is empty or invalid, this method will set all of the data fields
// to the default values.
//
// The Number String Format Defaults represent formatting
// parameters used in the United States.
//
func (nStrFmtDto *NumStrFormatDto) SetToDefaultsIfEmpty() {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	isValid,
		_ :=
		nStrFmtDtoMech.testNumStrFormatDtoValidity(
			nStrFmtDto,
			"")

	if isValid {
		return
	}

	nStrFmtDtoUtil := numStrFormatDtoUtility{}

	_ = nStrFmtDtoUtil.setToDefaults(
		nStrFmtDto,
		"")

}

// SetValueDisplaySpec - Sets the NumStrValSpec value. This value determines
// the type of numeric value which will be displayed in text number strings.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  valueDisplaySpec    NumStrValSpec
//     - This parameter must be set to one of the three values
//       listed below:
//
//       AbsoluteValue      NumStrValSpec(0).AbsoluteValue()
//        - This specification signals that a numeric value will be displayed
//          in text as a positive number regardless of whether the native
//          value is positive or negative. Effectively, this means that
//          both negative values and positive values will be displayed as
//          positive numbers.
//
//          Examples:
//
//               Positive Values          Negative Values
//                +132 = +132              -123 = +123
//
//       CurrencyValue      NumStrValSpec(0).CurrencyValue()
//        - The 'Currency Value' specification signals that all numeric values
//          will be displayed in number strings as currency formatted with
//          appropriate currency characters.
//
//          Currency number strings are always displayed as signed numeric
//          values with currency symbols included in the text string. This
//          means that positive values are displayed in text as positive
//          numbers with currency symbols (like the dollar sign) included
//          in the text string. Likewise, negative values are displayed in
//          text as negative numbers with currency symbols (like the dollar
//          sign) included in the text string.
//
//          Examples:
//               Positive Values          Negative Values
//                +132 = $132               -123 = ($123)
//
//       SignedNumberValue  NumStrValSpec(0).SignedNumberValue()
//        - Signals that the numeric value will be displayed in text as a
//          standard positive or negative value contingent upon the number
//          sign associated with the numeric value. NO CURRENCY Symbols will
//          be display in the resulting text number strings.
//
//          This is the default handling for numeric values.
//
//          'SignedNumberValue' means that positive values will be displayed
//           as positive numbers and negative values will be displayed as
//           negative numbers.
//
//           Examples:
//               Positive Values          Negative Values
//                +132 = 132               -123 = -123
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. This
//       error prefix, 'ePrefix' will be prefixed to the beginning
//       of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetValueDisplaySpec(
	valueDisplaySpec NumStrValSpec,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetValueDisplaySpec() "

	if !valueDisplaySpec.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'valueDisplaySpec' is invalid!\n"+
			"valueDisplaySpec='%v'\n",
			valueDisplaySpec.XValueInt())
		return err
	}

	nStrFmtDto.valueDisplaySpec = valueDisplaySpec

	return err
}
