package numberstr

type INumStrFormatter interface {
	Empty()

	GetNumStrFormatTypeCode() NumStrFormatTypeCode

	GetFmtNumStr(
		absValIntRunes []rune,
		absValFracRunes []rune,
		signVal int,
		baseNumSys int,
		ePrefix *ErrPrefixDto) (
		string,
		error)

	SetNumStrFormatTypeCode()
}
