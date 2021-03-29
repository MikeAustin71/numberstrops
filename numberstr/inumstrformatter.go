package numberstr

type INumStrFormatter interface {
	CopyInINumStrFormatter(
		incomingIFormatter INumStrFormatter,
		ePrefix *ErrPrefixDto) error

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
