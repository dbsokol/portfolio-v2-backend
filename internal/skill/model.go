package skill

import (
	"encoding/json"
	"errors"
	"portfolio/internal"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	internal.Model
	Name      string    `gorm:"type:varchar(35);not null;unique" json:"name"`
	StartDate time.Time `gorm:"type:date" json:"-"`
	Type      int       `gorm:"type:smallint;not null" json:"type"`
}

// MarshalJSON formats StartDate as "YYYY-MM-DD" in JSON output
func (s Skill) MarshalJSON() ([]byte, error) {
	type Alias Skill // Create an alias to avoid infinite recursion
	return json.Marshal(&struct {
		Alias
		StartDate string `json:"startDate"`
		Type      string `json:"type"`
	}{
		Alias:     (Alias)(s),
		StartDate: s.StartDate.Format("2006-01-02"),
		Type:      SkillTypeChoices.GetPublic(s.Type),
	})
}

func (s *Skill) BeforeSave(tx *gorm.DB) (err error) {

	skillType := s.Type

	_, err = SkillTypeChoices.Validate(SkillTypeChoices.GetPublic(skillType))
	if err != nil {
		return errors.New("field must be one of: " + strings.Join(SkillTypeChoices.GetPublicList(), ", "))
	}

	return nil
}
