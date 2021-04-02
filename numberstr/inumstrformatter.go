package numberstr

type INumStrFormatter interface {
	CopyInINumStrFormatter(
		incomingIFormatter INumStrFormatter,
		ePrefix *ErrPrefixDto) error

	CopyOutINumStrFormatter(
		ePrefix *ErrPrefixDto) (
		INumStrFormatter,
		error)

	Empty()

	EqualINumStrFormatter(
		incomingIFormatter INumStrFormatter,
		ePrefix *ErrPrefixDto) (
		isEqual bool,
		err error)

	GetNumStrFormatTypeCode() NumStrFormatTypeCode

	GetFmtNumStr(
		absValIntegerRunes []rune,
		absValFractionalRunes []rune,
		signVal int,
		baseNumSys BaseNumberSystemType,
		ePrefix *ErrPrefixDto) (
		fmtNumStr string,
		err error)

	IsValidInstanceError(
		ePrefix *ErrPrefixDto) error

	SetNumStrFormatTypeCode()
}
