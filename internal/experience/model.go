package experience

import (
	"time"
	"portfolio/internal/responsibility"
)

type Experience struct {
	ID              uint                       `gorm:"primaryKey"`
	Company         string                     `gorm:"type:varchar(100);"`
	Role            string                     `gorm:"type:text;"`
	Mission         string                     `gorm:"type:text;"`
	StartDate       time.Time
	EndDate         *time.Time
	Responsibilities []responsibility.Responsibility `gorm:"foreignKey:ExperienceID"` 
}
