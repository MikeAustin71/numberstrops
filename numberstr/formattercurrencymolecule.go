package numberstr

import (
	"fmt"
	"sync"
)

type formatterCurrencyMolecule struct {
	lock *sync.Mutex
}

// emptyFormatterCurrency - Deletes all data in a FormatterCurrency
// object passed as input parameter, 'formatterCurrency'.
//
// All member variable data fields in this object will be returned
// to their 'zero' state.
//
func (fmtCurrMolecule *formatterCurrencyMolecule) emptyFormatterCurrency(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	err error) {

	if fmtCurrMolecule.lock == nil {
		fmtCurrMolecule.lock = new(sync.Mutex)
	}

	fmtCurrMolecule.lock.Lock()

	defer fmtCurrMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(

		"formatterCurrencyMolecule." +
			"emptyFormatterCurrency()")

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	formatterCurrency.numStrFormatterType =
		NumStrFormatTypeCode(0).None()

	formatterCurrency.positiveValueFmt = ""

	formatterCurrency.negativeValueFmt = ""

	formatterCurrency.decimalDigits = 0

	formatterCurrency.currencyCode = ""

	formatterCurrency.currencyCodeNo = ""

	formatterCurrency.currencyName = ""

	formatterCurrency.currencySymbols =
		make([]rune, 0)

	formatterCurrency.minorCurrencyName = ""

	formatterCurrency.minorCurrencySymbols =
		make([]rune, 0)

	formatterCurrency.turnOnIntegerDigitsSeparation =
		false

	formatterCurrency.numericSeparators =
		NumericSeparators{}

	formatterCurrency.numFieldLenDto =
		NumberFieldDto{}

	return err
}

// ptr - Returns a pointer to a new instance of
// formatterCurrencyMolecule.
//
func (fmtCurrMolecule formatterCurrencyMolecule) ptr() *formatterCurrencyMolecule {

	if fmtCurrMolecule.lock == nil {
		fmtCurrMolecule.lock = new(sync.Mutex)
	}

	fmtCurrMolecule.lock.Lock()

	defer fmtCurrMolecule.lock.Unlock()

	newFmtCurrencyMolecule :=
		new(formatterCurrencyMolecule)

	newFmtCurrencyMolecule.lock = new(sync.Mutex)

	return newFmtCurrencyMolecule
}

// testValidityOfFormatterCurrency - Tests the validity of
// FormatterCurrency objects.
//
func (fmtCurrMolecule *formatterCurrencyMolecule) testValidityOfFormatterCurrency(
	formatterCurrency *FormatterCurrency,
	ePrefix *ErrPrefixDto) (
	isValid bool,
	err error) {

	if fmtCurrMolecule.lock == nil {
		fmtCurrMolecule.lock = new(sync.Mutex)
	}

	fmtCurrMolecule.lock.Lock()

	defer fmtCurrMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"formatterCurrencyMolecule." +
			"testValidityOfFormatterCurrency()")

	isValid = false

	if formatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'formatterCurrency' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterCurrency.numStrFormatterType !=
		NumStrFormatTypeCode(0).Currency() {

		err = fmt.Errorf("%v\n"+
			"Error: Number String Formatter Type Code is invalid!\n"+
			"formatterCurrency.numStrFormatterType != NumStrFormatTypeCode(0).Currency().\n"+
			"Integer formatterCurrency.numStrFormatterType ='%v'",
			ePrefix.String(),
			formatterCurrency.numStrFormatterType.XValueInt())

		return isValid, err

	}

	if formatterCurrency.currencySymbols == nil {
		formatterCurrency.currencySymbols =
			make([]rune, 0, 10)
	}

	lenCurrencySymbols := len(formatterCurrency.currencySymbols)

	if lenCurrencySymbols == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The Currency Symbol is missing!\n"+
			"The currency symbols rune array is a zero length array.\n",
			ePrefix.String())

		return isValid, err
	}

	if formatterCurrency.minorCurrencySymbols == nil {
		formatterCurrency.minorCurrencySymbols =
			make([]rune, 0, 10)
	}

	err =
		formatterCurrency.numericSeparators.IsValidInstanceError(
			ePrefix.XCtx("Validating 'formatterCurrency' Number Separators"))

	if err != nil {
		return isValid, err
	}

	nStrCurrencyAtom :=
		formatterCurrencyAtom{}

	_,
		err = nStrCurrencyAtom.testCurrencyPositiveValueFormat(
		formatterCurrency,
		ePrefix.XCtx("Testing formatterCurrency.positiveValueFmt"))

	if err != nil {
		return isValid, err
	}

	_,
		err = nStrCurrencyAtom.testCurrencyNegativeValueFormat(
		formatterCurrency,
		ePrefix.XCtx("\nTesting formatterCurrency.positiveValueFmt"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
