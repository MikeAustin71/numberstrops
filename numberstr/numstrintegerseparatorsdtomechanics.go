package numberstr

import (
	"fmt"
	"sync"
)

type numStrIntSeparatorsDtoMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numStrIntSeparatorsDtoMechanics.
//
func (intSepsDtoMech numStrIntSeparatorsDtoMechanics) ptr() *numStrIntSeparatorsDtoMechanics {

	if intSepsDtoMech.lock == nil {
		intSepsDtoMech.lock = new(sync.Mutex)
	}

	intSepsDtoMech.lock.Lock()

	defer intSepsDtoMech.lock.Unlock()

	newMech := new(numStrIntSeparatorsDtoMechanics)

	newMech.lock = new(sync.Mutex)

	return newMech
}

// setToUSADefaults - Sets the member variable data
// values of the incoming NumStrIntSeparatorsDto instance to United
// States default values.
//
func (intSepsDtoMech *numStrIntSeparatorsDtoMechanics) setToUSADefaults(
	intSepsDto *NumStrIntSeparatorsDto,
	ePrefix *ErrPrefixDto) (
	err error) {

	if intSepsDtoMech.lock == nil {
		intSepsDtoMech.lock = new(sync.Mutex)
	}

	intSepsDtoMech.lock.Lock()

	defer intSepsDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorsDtoMechanics." +
			"setToUSADefaults()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	intSepsDto.intSeparators = nil

	intSepsDto.intSeparators =
		make([]NumStrIntSeparator, 1)

	err =
		intSepsDto.intSeparators[0].SetToUSADefaults(
			ePrefix.XCtx(
				"intSepsDto.intSeparators[0]"))

	return err
}

// setWithComponents - Sets the member variable data values of the
// incoming NumStrIntSeparatorsDto instance with the array values
// passed by input parameter 'integerSeparatorsDto'.
//
func (intSepsDtoMech *numStrIntSeparatorsDtoMechanics) setWithComponents(
	intSepsDto *NumStrIntSeparatorsDto,
	integerSeparators []NumStrIntSeparator,
	ePrefix *ErrPrefixDto) (
	err error) {

	if intSepsDtoMech.lock == nil {
		intSepsDtoMech.lock = new(sync.Mutex)
	}

	intSepsDtoMech.lock.Lock()

	defer intSepsDtoMech.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	}

	ePrefix.SetEPref(
		"numStrIntSeparatorsDtoMechanics." +
			"setWithComponents()")

	if intSepsDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSepsDto' is invalid!\n"+
			"'intSepsDto' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if intSepsDto.lock == nil {
		intSepsDto.lock = new(sync.Mutex)
	}

	lenIncomingSeps := len(integerSeparators)

	if lenIncomingSeps == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerSeparatorsDto' is invalid!\n"+
			"'integerSeparatorsDto' is a ZERO length array.\n",
			ePrefix.String())

		return err
	}

	intSepsDto.intSeparators =
		make([]NumStrIntSeparator, lenIncomingSeps, lenIncomingSeps+5)

	for i := 0; i < lenIncomingSeps; i++ {
		err =
			intSepsDto.intSeparators[i].CopyIn(
				&integerSeparators[i],
				ePrefix.XCtx(
					fmt.Sprintf(
						"integerSeparatorsDto[%v]",
						i)))

		if err != nil {
			return err
		}
	}

	return err
}
