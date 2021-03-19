package numberstr

type INumStrFormatter interface {
	Empty()

	GetNumStrFormatTypeCode() NumStrFormatTypeCode

	SetNumStrFormatTypeCode()
}
