package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtSpecAbsoluteValueDtoUtility struct {
	lock *sync.Mutex
}

func (nStrFmtSpecAbsValDtoUtil *numStrFmtSpecAbsoluteValueDtoUtility) setAbsValDtoWithDefaults(
	nStrFmtSpecAbsValDto *NumStrFmtSpecAbsoluteValueDto,
	absoluteValFmt string,
	turnOnIntegerDigitsSeparation bool,
	decimalSeparatorChar rune,
	thousandsSeparatorChar rune,
	integerDigitsGroupingSequence []uint,
	requestedNumberFieldLen int,
	ePrefix string) (
	err error) {

	if nStrFmtSpecAbsValDtoUtil.lock == nil {
		nStrFmtSpecAbsValDtoUtil.lock = new(sync.Mutex)
	}

	nStrFmtSpecAbsValDtoUtil.lock.Lock()

	defer nStrFmtSpecAbsValDtoUtil.lock.Unlock()

	ePrefix += "\nnumStrFmtSpecAbsoluteValueDtoUtility.setAbsValDtoWithDefaults() "

	if nStrFmtSpecAbsValDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtSpecAbsValDto' is invalid!\n"+
			"'nStrFmtSpecAbsValDto' is a 'nil' pointer.\n",
			ePrefix)
		return err
	}

	numFieldDto := NumberFieldDto{}.NewWithDefaults(requestedNumberFieldLen)

	var numberSeparatorsDto NumStrFmtSpecDigitsSeparatorsDto

	numberSeparatorsDto,
		err = NumStrFmtSpecDigitsSeparatorsDto{}.NewFromComponents(
		decimalSeparatorChar,
		thousandsSeparatorChar,
		integerDigitsGroupingSequence,
		ePrefix+
			fmt.Sprintf("\n"+
				"decimalSeparatorChar='%v'\n"+
				"thousandsSeparatorChar='%v'\n",
				decimalSeparatorChar,
				thousandsSeparatorChar))

	if err != nil {
		return err
	}

	nStrFmtSpecAbsValDtoMech :=
		numStrFmtSpecAbsoluteValueDtoMechanics{}

	err = nStrFmtSpecAbsValDtoMech.setAbsValDto(
		nStrFmtSpecAbsValDto,
		absoluteValFmt,
		turnOnIntegerDigitsSeparation,
		numberSeparatorsDto,
		numFieldDto,
		ePrefix)

	return err
}
