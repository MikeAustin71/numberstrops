package numberstr

type INumStrFormatter interface {
	Empty()

	GetNumStrFormatTypeCode() NumStrFormatTypeCode

	GetFmtNumStr(
		absValIntegerRunes []rune,
		absValFractionalRunes []rune,
		signVal int,
		baseNumSys BaseNumberSystemType,
		ePrefix *ErrPrefixDto) (
		fmtNumStr string,
		err error)

	SetNumStrFormatTypeCode()
}
