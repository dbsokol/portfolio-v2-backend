package education

import (
	"encoding/json"
	"time"

	"portfolio/internal"
)

type Education struct {
	internal.Model
	Institution string    `gorm:"type:varchar(100);not null;index:idx_education_unique,unique" json:"institution"`
	Degree      int       `gorm:"type:smallint;not null;index:idx_education_unique,unique" json:"degree"`
	Major       string    `gorm:"type:varchar(100);not null;index:idx_education_unique,unique" json:"major"`
	StartDate   time.Time `gorm:"type:date" json:"-"`
	EndDate     time.Time `gorm:"type:date" json:"-"`
}

// MarshalJSON formats StartDate and EndDate as "YYYY-MM-DD" in JSON output
func (e Education) MarshalJSON() ([]byte, error) {
	type Alias Education // Create an alias to avoid infinite recursion
	return json.Marshal(&struct {
		Alias
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Degree    string `json:"degree"`
	}{
		Alias:     (Alias)(e),
		StartDate: e.StartDate.Format("2006-01-02"),
		EndDate:   e.EndDate.Format("2006-01-02"),
		Degree:    DegreeChoices.GetPublic(e.Degree),
	})
}

func (e *Education) BeforeSave() error {
	_, err := DegreeChoices.Validate(DegreeChoices.GetPublic(e.Degree))
	if err != nil {
		return err
	}
	return nil
}
