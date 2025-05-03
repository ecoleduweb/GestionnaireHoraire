package enums

import "fmt"

type UserRole int

const (
	Employee UserRole = iota
	ProjectManager
	Administrator
)

func ParseUserRole(role int) (UserRole, error) {
	switch role {
	case int(Employee): // This is 0
		return Employee, nil
	case int(ProjectManager): // This is 1
		return ProjectManager, nil
	case int(Administrator): // This is 2
		return Administrator, nil
	default:
		return UserRole(-1), fmt.Errorf("rôle invalide: %d", role)
	}
}
