package numberstr

import (
	"fmt"
	"sync"
)

type NumStrFormatDto struct {
	valueDisplaySpec NumStrValSpec
	positiveValueFmt string
	negativeValueFmt string
	currencyFmt      CurrencySymbolDto
	lock             *sync.Mutex
}

// New - Creates and returns a new instance of
// NumStrFormatDto.
//
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
//  positiveValueFmt    string
//     - A string specifying the number string format to be used in
//       formatting positive numeric value in text strings. Valid
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
//        "NUMFIELD" - Signifies a number field. A number field has a string
//                     length which is equal to or greater than the actual
//                     numeric value string length. Actual numeric values
//                     are right justified within number fields for text
//                     displays.
//
//          "127.54" - Represents the actual length of the numeric value
//                     including formatting characters and symbols such
//                     as Thousands Separators, Decimal Separators and
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
//  negativeValueFmt    string
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
//        "NUMFIELD" - Signifies a number field. A number field has a string
//                     length which is equal to or greater than the actual
//                     numeric value string length. Actual numeric values
//                     are right justified within number fields for text
//                     displays.
//
//          "127.54" - Represents the actual length of the numeric value
//                     including formatting characters and symbols such
//                     as Thousands Separators, Decimal Separators and
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
//  newFmtDto           NumStrFormatDto
//     - If this method completes successfully, a new instance of NumStrFormatDto
//       be returned to calling function.
//
//
//  err                 error
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

	newFmtDto.valueDisplaySpec = valueDisplaySpec
	newFmtDto.negativeValueFmt = negativeValueFmt
	newFmtDto.positiveValueFmt = positiveValueFmt
	newFmtDto.currencyFmt = currencyFmt.CopyOut()

	return newFmtDto, err
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

// GetPositiveValueFormatMode - Returns the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) GetPositiveValueFormatMode() (
	positiveValueFmtMode NumStrPosValFmtMode) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	positiveValueFmtMode = nStrFmtDto.positiveValueFmt

	return positiveValueFmtMode
}

// GetNegativeValueFormatMode - Returns the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) GetNegativeValueFormatMode() (
	negativeValueFmtMode NumStrNegValFmtMode) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	negativeValueFmtMode = nStrFmtDto.negativeValueFmt

	return negativeValueFmtMode
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
//  isValid             bool
//     - If this instance of NumStrFormatDto is invalid, this boolean
//       value will be set to 'false'. Otherwise, if this instance is
//       valid, the value will be set to 'true'.
//
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
	isValid bool,
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.IsValidInstanceError() "

	isValid = false

	if !nStrFmtDto.valueDisplaySpec.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Member variable 'valueDisplaySpec' is invalid!\n"+
			"valueDisplaySpec='%v'\n",
			nStrFmtDto.valueDisplaySpec.XValueInt())
		return isValid, err
	}

	if !nStrFmtDto.positiveValueFmt.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Member variable 'positiveValueFmt' is invalid!\n"+
			"positiveValueFmt='%v'\n",
			nStrFmtDto.positiveValueFmt.XValueInt())
		return isValid, err
	}

	if !nStrFmtDto.negativeValueFmt.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Member variable 'negativeValueFmt' is invalid!\n"+
			"negativeValueFmt='%v'\n",
			nStrFmtDto.negativeValueFmt.XValueInt())
		return isValid, err
	}

	isValid = true

	return isValid, err
}

// SetDefaults - This method will set all internal
// member variables to their default values.
//
// Default values for Number String formatting are
// listed as follows:
//
//  valueDisplaySpec = NumStrValSpec(0).StdPlusOrMinus()
//  positiveValueFmt = NumStrPosValFmtMode(0).NoLeadingPlusSign()
//  negativeValueFmt = NumStrNegValFmtMode(0).LeadingMinusSign()
//
func (nStrFmtDto *NumStrFormatDto) SetDefaults() {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDto.valueDisplaySpec = NumStrValSpec(0).StdPlusOrMinus()
	nStrFmtDto.positiveValueFmt = NumStrPosValFmtMode(0).NoLeadingPlusSign()
	nStrFmtDto.negativeValueFmt = NumStrNegValFmtMode(0).LeadingMinusSign()
}

// SetNegativeValueFormatMode - Sets the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) SetNegativeValueFormatMode(
	negativeValueFmtMode NumStrNegValFmtMode,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetNegativeValueFormatMode() "

	if !negativeValueFmtMode.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'negativeValueFmtMode' is invalid!\n"+
			"negativeValueFmtMode='%v'\n",
			negativeValueFmtMode.XValueInt())
		return err
	}

	negativeValueFmtMode = nStrFmtDto.negativeValueFmt

	return err
}

// SetPositiveValueFormatMode - Sets the Positive Value Format
// Mode or NumStrPosValFmtMode value.
//
func (nStrFmtDto *NumStrFormatDto) SetPositiveValueFormatMode(
	positiveValueFmtMode NumStrPosValFmtMode,
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetPositiveValueFormatMode() "

	if !positiveValueFmtMode.XIsValid() {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Input parameter 'positiveValueFmtMode' is invalid!\n"+
			"positiveValueFmtMode='%v'\n",
			positiveValueFmtMode.XValueInt())
		return err
	}

	nStrFmtDto.positiveValueFmt = positiveValueFmtMode

	return err
}

// SetValueDisplaySpec - Sets the NumStrValSpec
// value.
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
