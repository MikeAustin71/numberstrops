package numberstr

import (
	"sync"
)

// NumStrFmtSpecDigitsSeparatorsDto - This type encapsulates the
// number separators used to format numeric values in text strings.
// These separators include the 'Decimal Separator', The 'Integer
// Digits Separator' and the 'Integer Digits Grouping Sequence'.
//
// decimalSeparator rune
//
// The 'Decimal Separator' is used to separate integer and
// fractional digits within a floating point number display.
//
// integerDigitsSeparator rune
//
// This type also encapsulates the integer digits separator, often
// referred to as the 'Thousands Separator'. This is used to
// separate thousands digits within the integer component of a
// number string.
//
// integerDigitsGroupingSequence []uint
//
// Related to the integer digits separator, the integer digits
// grouping sequence is also encapsulated in this type. The integer
// digits grouping sequence is used to identify the digits which
// will be grouped and separated by the integer digits separator.
//
// In most western countries integer digits to the left of the
// decimal separator (a.k.a. decimal point) are separated into
// groups of three digits representing a grouping of 'thousands'
// like this: '1,000,000,000,000'. In this case the parameter
// integer digits grouping sequence would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted as like this:
//               '6,78,90,00,00,00,00,000'
// The right most group has three digits and all the others are
// grouped by two. In this case integer digits grouping sequence
// would be configured as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//

type NumStrFmtSpecDigitsSeparatorsDto struct {
	decimalSeparator              rune
	integerDigitsSeparator        rune
	integerDigitsGroupingSequence []uint
	lock                          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrFmtSpecDigitsSeparatorsDto instance to the data fields
// of the current instance of NumStrFmtSpecDigitsSeparatorsDto
// instance.
//
// If the incoming NumStrFmtSpecDigitsSeparatorsDto is judged to be
// invalid, this method will return an error.
//
// This method will OVERWRITE the data fields of the current
// NumStrFmtSpecDigitsSeparatorsDto instance.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) CopyIn(
	incomingSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	ePrefix.SetEPref("NumStrFmtSpecDigitsSeparatorsDto.CopyIn()")

	nStrFmtSpecDigitsSepsElectron :=
		numStrFmtSpecDigitsSeparatorsDtoElectron{}

	return nStrFmtSpecDigitsSepsElectron.copyIn(
		nStrFmtSpecDigitsSepDto,
		incomingSpecDigitsSepDto,
		ePrefix)
}

// CopyOut - Returns a deep copy of the current
// NumStrFmtSpecDigitsSeparatorsDto instance.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) CopyOut() NumStrFmtSpecDigitsSeparatorsDto {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	nStrFmtSpecDigitsSepsElectron :=
		numStrFmtSpecDigitsSeparatorsDtoElectron{}

	newDigitsSepDto,
		_ := nStrFmtSpecDigitsSepsElectron.copyOut(
		nStrFmtSpecDigitsSepDto,
		new(ErrPrefixDto))

	return newDigitsSepDto
}

// GetDecimalSeparator - Returns the decimal separator character.
//
// The decimal separator separates integer and fractional digits
// in a floating point number string. In the United States, the
// decimal separator is the period ('.') or decimal point.
//
//             Example: 1234.456
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) GetDecimalSeparator() rune {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	return nStrFmtSpecDigitsSepDto.decimalSeparator
}

// GetIntegerDigitsSeparator - Returns the integer digits separator
// character.
//
// The integer digits separator is also known as the 'thousands'
// separator. In the United States the standard integer digits
// separator is the comma character (',').
//
//  Example: 1,000,000,000,000
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) GetIntegerDigitsSeparator() rune {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	return nStrFmtSpecDigitsSepDto.integerDigitsSeparator
}

// GetIntegerDigitsGroupingSequence - Returns a deep copy of the
// integer digits grouping sequence as an array of unsigned
// integers. This refers to grouping of integer digits within a
// string of numeric digits.
//
// In most western countries integer digits to the left of the
// decimal separator (a.k.a. decimal point) are separated into
// groups of three digits representing a grouping of 'thousands'
// like this: '1,000,000,000,000'. In this case the integer digits
// grouping sequence would be configured as:
//        integerDigitsGroupingSequence = []uint{3}
//
// In some countries and cultures other integer groupings are
// used. In India, for example, a number might be formatted as
// like this: '6,78,90,00,00,00,00,000'. The right most group
// has three digits and all the others are grouped by two digits.
// In this case integer digits grouping sequence would be
// configured as:
//        integerDigitsGroupingSequence = []uint{3,2}
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) GetIntegerDigitsGroupingSequence() []uint {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	lenGrpSeq := len(nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence)

	groupingSeq := make([]uint, lenGrpSeq, 10)

	if lenGrpSeq == 0 {
		return groupingSeq
	}

	for i := 0; i < lenGrpSeq; i++ {
		groupingSeq[i] =
			nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence[i]
	}

	return groupingSeq
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrFmtSpecDigitsSeparatorsDto instance to determine whether
// the current instance is valid in all respects.
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
//       value will signal whether the current NumStrFmtSpecDigitsSeparatorsDto
//       is valid, or not. If the current NumStrFmtSpecDigitsSeparatorsDto
//       contains valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	nStrFmtSpecDigitsSepsQuark :=
		numStrFmtSpecDigitsSeparatorsDtoQuark{}

	isValid,
		_ = nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
		nStrFmtSpecDigitsSepDto,
		new(ErrPrefixDto))

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrFmtSpecDigitsSeparatorsDto instance to determine
// whether the current instance is valid in all respects.
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
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
//       If this instance of NumStrFmtSpecDigitsSeparatorsDto contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) IsValidInstanceError(
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	ePrefix.SetEPrefCtx(
		"NumStrFmtSpecDigitsSeparatorsDto.IsValidInstanceError()",
		"Testing Validity of 'nStrFmtSpecDigitsSepDto'")

	nStrFmtSpecDigitsSepsQuark := numStrFmtSpecDigitsSeparatorsDtoQuark{}

	_,
		err := nStrFmtSpecDigitsSepsQuark.testValidityOfNumSepsDto(
		nStrFmtSpecDigitsSepDto,
		ePrefix)

	return err
}

// NewWithDefaults - Creates and returns a new instance of
// NumStrFmtSpecDigitsSeparatorsDto. This type encapsulates the
// digit separators used in formatting a number string for text
// display. Digit separators are commonly referred to as the
// decimal point and thousands separator characters. In addition,
// Digit Separators include the integer digits grouping sequence.
//
// This method default the integer digits grouping sequence to that
// used in most of the world.
//
// In most countries, integer digits to the left of the decimal
// separator (a.k.a. decimal point) are separated into groups of
// three (3) digits representing a grouping of 'thousands' like
// this: '1,000,000,000,000'. In this case the 'integer digits
// grouping sequence' would be configured as:
//        integer digits grouping sequence = []uint{3}
// This method applies the 3-digit integer digits grouping sequence
// by default.
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted as like this:
// '6,78,90,00,00,00,00,000'. The right most group has three digits
// and all the others are grouped by two. In this case 'integer
// digits grouping sequence' would be configured as:
//        integerDigitsGroupingSequence = []uint{3,2}
//
// Again, this method will automatically set the 'integer digits
// grouping sequence' to a default of 3-digits or []uint{3}.
//
// If you wish to configure the 'integer digits grouping sequence'
// to some value other than the default, see method:
//     NumStrFmtSpecDigitsSeparatorsDto.NewFromComponents()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparator           rune
//     - A text character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the standard decimal separator is the period
//       ('.') or decimal point.
//
//       If this input parameter is set to zero, an error will be
//       returned.
//
//
//  integerDigitsSeparator     rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//       If this input parameter is set to zero, an error will be
//       returned.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDigitsSeparatorsDto
//     - If this method completes successfully, new instance of
//       NumStrFmtSpecDigitsSeparatorsDto will be created and
//       returned. The 'integer digits grouping sequence' will be
//       automatically set to default of 3-digits or []uint{3}.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecDigitsSepDto NumStrFmtSpecDigitsSeparatorsDto) NewWithDefaults(
	ePrefix *ErrPrefixDto,
	decimalSeparator rune,
	integerDigitsSeparator rune) (
	NumStrFmtSpecDigitsSeparatorsDto,
	error) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	ePrefix.SetEPref(
		"NumStrFmtSpecDigitsSeparatorsDto.NewWithDefaults()")

	newDigitsSepsDto := NumStrFmtSpecDigitsSeparatorsDto{}

	nStrFmtSpecDigitsSepsDtoMech :=
		numStrFmtSpecDigitsSepsDtoMechanics{}

	err := nStrFmtSpecDigitsSepsDtoMech.setDigitsSeps(
		&newDigitsSepsDto,
		decimalSeparator,
		integerDigitsSeparator,
		[]uint{3},
		ePrefix)

	return newDigitsSepsDto, err

}

// NewFromComponents - Creates and returns a new instance of NumStrFmtSpecDigitsSeparatorsDto.
// This type encapsulates the digit separators used in formatting a
// number string for text display.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparator               rune
//     - A text character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the standard decimal separator is the period
//       ('.') or decimal point.
//
//
//  integerDigitsSeparator         rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrFmtSpecDigitsSeparatorsDto
//     - If this method completes successfully, new instance of
//       NumStrFmtSpecDigitsSeparatorsDto will be created and
//       returned.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecDigitsSepDto NumStrFmtSpecDigitsSeparatorsDto) NewFromComponents(
	decimalSeparator rune,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix *ErrPrefixDto) (
	NumStrFmtSpecDigitsSeparatorsDto,
	error) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	ePrefix.SetEPref("NumStrFmtSpecDigitsSeparatorsDto.NewFromComponents()")

	newDigitsSepsDto := NumStrFmtSpecDigitsSeparatorsDto{}

	nStrFmtSpecDigitsSepsDtoMech :=
		numStrFmtSpecDigitsSepsDtoMechanics{}

	err := nStrFmtSpecDigitsSepsDtoMech.setDigitsSeps(
		&newDigitsSepsDto,
		decimalSeparator,
		integerDigitsSeparator,
		integerDigitsGroupingSequence,
		ePrefix)

	return newDigitsSepsDto, err
}

// SetDecimalSeparator - Sets the value of the decimal separator.
//
// The decimal separator separates integer and fractional digits
// in a floating point number string. In the United States, the
// decimal separator is the period ('.') or decimal point.
//
//             Example: 1234.456
//
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) SetDecimalSeparator(
	decimalSeparator rune) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	nStrFmtSpecDigitsSepDto.decimalSeparator = decimalSeparator

}

// SetDigitsSeps - This method will set all of the member variable
// data values for the current instance of NumStrFmtSpecDigitsSeparatorsDto.
//
// The NumStrFmtSpecDigitsSeparatorsDto type encapsulates the digit separators
// used in formatting a number string for text display.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparator               rune
//     - A text character used to separate integer and fractional
//       digits in a floating point number string. In the United
//       States, the standard decimal separator is the period
//       ('.') or decimal point.
//
//
//  integerDigitsSeparator         rune
//     - This separator is also known as the 'thousands' separator.
//       It is used to separate groups of integer digits to the left
//       of the decimal separator (a.k.a. decimal point). In the
//       United States, the standard integer digits separator is the
//       comma (',').
//
//        Example:  1,000,000,000
//
//
//  integerDigitsGroupingSequence  []uint
//     - In most western countries integer digits to the left of the
//       decimal separator (a.k.a. decimal point) are separated into
//       groups of three digits representing a grouping of 'thousands'
//       like this: '1,000,000,000,000'. In this case the parameter
//       'integerDigitsGroupingSequence' would be configured as:
//              integerDigitsGroupingSequence = []uint{3}
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted as
//       like this: '6,78,90,00,00,00,00,000'. The right most group
//       has three digits and all the others are grouped by two. In
//       this case 'integerDigitsGroupingSequence' would be configured
//       as:
//              integerDigitsGroupingSequence = []uint{3,2}
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) SetDigitsSeps(
	decimalSeparator rune,
	integerDigitsSeparator rune,
	integerDigitsGroupingSequence []uint,
	ePrefix *ErrPrefixDto) error {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	ePrefix.SetEPref("NumStrFmtSpecDigitsSeparatorsDto.SetDigitsSeps()")

	nStrFmtSpecDigitsSepsDtoMech :=
		numStrFmtSpecDigitsSepsDtoMechanics{}

	return nStrFmtSpecDigitsSepsDtoMech.setDigitsSeps(
		nStrFmtSpecDigitsSepDto,
		decimalSeparator,
		integerDigitsSeparator,
		integerDigitsGroupingSequence,
		ePrefix.XCtx("Setting Data Values for current instance 'nStrFmtSpecDigitsSepDto'"))
}

// SetIntegerDigitsGroupingSequence - Sets the value of the integer
// digits grouping sequence. This refers to grouping of integer digits
// within a string of numeric digits.
//
// In most western countries integer digits to the left of the
// decimal separator (a.k.a. decimal point) are separated into
// groups of three digits representing a grouping of 'thousands'
// like this: '1,000,000,000,000'. In this case, the integer digits
// grouping sequence would be configured as:
//        integerDigitsGroupingSequence = []uint{3}
//
// In some countries and cultures, other integer groupings are
// used. In India, for example, a number might be formatted as
// like this: '6,78,90,00,00,00,00,000'. The right most group
// has three digits and all the others are grouped by two digits.
// In this case integer digits grouping sequence would be
// configured as:
//        integerDigitsGroupingSequence = []uint{3,2}
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) SetIntegerDigitsGroupingSequence(
	integerDigitsGroupingSeq []uint) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	lenIntDigsGrpSeq := len(integerDigitsGroupingSeq)

	if lenIntDigsGrpSeq == 0 {

		nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence =
			make([]uint, 0, 10)

	} else {

		nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence =
			make([]uint, lenIntDigsGrpSeq, 10)

		for i := 0; i < lenIntDigsGrpSeq; i++ {
			nStrFmtSpecDigitsSepDto.integerDigitsGroupingSequence[i] =
				integerDigitsGroupingSeq[i]
		}
	}

	return
}

// SetIntegerDigitsSeparator - Sets the value of the integer digits
// separator.
//
// This separator is used in formatting a string of numeric digits.
// It is also known as the 'thousands' separator. The integer
// digits separator is used to separate groups of integer digits to
// the left of the decimal separator (a.k.a. decimal point). In the
// United States, the standard integer digits separator is the comma
// (',').
//
//        Example:  1,000,000,000
//
func (nStrFmtSpecDigitsSepDto *NumStrFmtSpecDigitsSeparatorsDto) SetIntegerDigitsSeparator(
	integerDigitsSeparator rune) {

	if nStrFmtSpecDigitsSepDto.lock == nil {
		nStrFmtSpecDigitsSepDto.lock = new(sync.Mutex)
	}

	nStrFmtSpecDigitsSepDto.lock.Lock()

	defer nStrFmtSpecDigitsSepDto.lock.Unlock()

	nStrFmtSpecDigitsSepDto.integerDigitsSeparator =
		integerDigitsSeparator
}
