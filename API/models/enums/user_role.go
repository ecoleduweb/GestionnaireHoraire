package enums

type UserRole int

const (
	Employee UserRole = iota
	ProjectManager
	Administrator
)

type ProjectStatus int

const (
	InProgress UserRole = iota
	Cancel
	Finish
	NotStart
)
