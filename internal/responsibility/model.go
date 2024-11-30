package responsibility

type Responsibility struct {
	ID           uint `gorm:"primary_key"`
	ExperienceID uint
	Description  string `gorm:"type:text;"`
}
