package enums

type UserRole int

const (
	Employee UserRole = iota
	ProjectManager
	Administrator
)
