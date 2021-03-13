package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorsDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingIntSepsDto' to input parameter
// 'targetIntSepsDto'.
//
// Be advised - All data fields in 'targetIntSepsDto'
// will be overwritten.
//
// If input parameter 'incomingIntSepsDto' is judged
// to be invalid, this method will return an error.
//
func (intSepsDtoElectron *numStrIntSeparatorsDtoElectron) copyIn(
	targetIntSepsDto *NumStrIntSeparatorsDto,
	incomingIntSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if intSepsDtoElectron.lock == nil {
		intSepsDtoElectron.lock = new(sync.Mutex)
	}

	intSepsDtoElectron.lock.Lock()

	defer intSepsDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrIntSeparatorsDtoElectron." +
		"copyIn()")

	if targetIntSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetIntSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingIntSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingIntSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		numStrIntSeparatorsDtoQuark{}.ptr().
			testValidityOfNumStrIntSepsDto(
				incomingIntSepsDto,
				ePrefix.XCtx(
					"incomingIntSepsDto"))

	if err != nil {
		return err
	}

	if incomingIntSepsDto.intSeparators == nil {
		incomingIntSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	lenIncomingIntSeps :=
		len(incomingIntSepsDto.intSeparators)

	targetIntSepsDto.intSeparators =
		make([]NumStrIntSeparator,
			lenIncomingIntSeps,
			lenIncomingIntSeps+5)

	for i := 0; i < lenIncomingIntSeps; i++ {
		err = targetIntSepsDto.intSeparators[i].
			CopyIn(
				&incomingIntSepsDto.intSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf(
						"incomingIntSepsDto.intSeparators[%v]",
						i)))

		if err != nil {
			return err
		}
	}

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'intSepsDto' styled as a new instance
// of NumStrIntSeparatorsDto.
//
// If input parameter 'intSepsDto' is judged to be
// invalid, this method will return an error.
//
func (intSepsDtoElectron *numStrIntSeparatorsDtoElectron) copyOut(
	intSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	newIntSepsDto NumStrIntSeparatorsDto,
	err error) {

	if intSepsDtoElectron.lock == nil {
		intSepsDtoElectron.lock = new(sync.Mutex)
	}

	intSepsDtoElectron.lock.Lock()

	defer intSepsDtoElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref("numStrIntSeparatorsDtoElectron." +
		"copyOut()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newIntSepsDto, err
	}

	if intSepsDto.intSeparators == nil {
		intSepsDto.intSeparators =
			make([]NumStrIntSeparator, 0, 5)
	}

	_,
		err =
		numStrIntSeparatorsDtoQuark{}.ptr().
			testValidityOfNumStrIntSepsDto(
				intSepsDto,
				ePrefix.XCtx(
					"intSepsDto"))

	if err != nil {
		return newIntSepsDto, err
	}

	newIntSepsDto.lock = new(sync.Mutex)

	lenIntSeps := len(intSepsDto.intSeparators)

	newIntSepsDto.intSeparators =
		make([]NumStrIntSeparator, lenIntSeps,
			lenIntSeps+5)

	for i := 0; i < lenIntSeps; i++ {
		err = newIntSepsDto.intSeparators[i].
			CopyIn(
				&intSepsDto.intSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf(
						"intSepsDto.intSeparators[%v]",
						i)))

		if err != nil {
			return newIntSepsDto, err
		}
	}

	return newIntSepsDto, err
}

// ptr - Returns a pointer to a new instance of
// numStrIntSeparatorsDtoElectron.
//
func (intSepsDtoElectron numStrIntSeparatorsDtoElectron) ptr() *numStrIntSeparatorsDtoElectron {

	if intSepsDtoElectron.lock == nil {
		intSepsDtoElectron.lock = new(sync.Mutex)
	}

	intSepsDtoElectron.lock.Lock()

	defer intSepsDtoElectron.lock.Unlock()

	newElectron := new(numStrIntSeparatorsDtoElectron)

	newElectron.lock = new(sync.Mutex)

	return newElectron
}
