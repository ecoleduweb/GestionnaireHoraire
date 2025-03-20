package DAOs

type UserRole int

const (
	Employee UserRole = iota
	ProjectManager
	Administrator
)

type User struct {
	Id         int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	FirstName  string     `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName   string     `json:"last_name" gorm:"type:varchar(50);not null"`
	Email      string     `json:"email" gorm:"type:varchar(255);not null;unique;index"`
	Role       UserRole   `json:"role" gorm:"type:enum(2,1,0);not null;default:0"`
	Activities []Activity `json:"activities" gorm:"foreignKey:UserId;references:Id"`
}
