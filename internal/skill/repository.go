package skill

import (
	"errors"
	"portfolio/config"
	"time"

	"gorm.io/gorm"
)

func CreateSkill(name string, startDate time.Time, skillType int) (*Skill, error) {
	/* create will "undelete" a deleted skill if it already exists, otherwise it will create a new one */

	var skill Skill

	// on create, check if the skill already exists but is deleted
	err := config.DB.Debug().Where("name = ? AND is_deleted = ?", name, true).First(&skill).Error

	if err == nil && skill.ID != 0 {
		skill.IsDeleted = false
		skill.StartDate = startDate
		skill.Type = skillType

		err = config.DB.Save(&skill).Error

		if err != nil {
			return nil, err
		}

		return &skill, nil
	}

	// otherwise, create a new skill
	skill = Skill{
		Name:      name,
		StartDate: startDate,
		Type:      skillType,
	}

	err = config.DB.Debug().Create(&skill).Error

	if err != nil {
		return nil, err
	}

	return &skill, nil
}

func DeleteSkill(uuid string) error {
	/* delete method is treated as soft delete */

	var skill Skill

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&skill).Error

	if err != nil {
		return err
	}

	skill.IsDeleted = true

	err = config.DB.Debug().Save(&skill).Error

	if err != nil {
		return err
	}

	return nil
}

func GetSkill(uuid string) (*Skill, error) {
	var skill Skill

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&skill).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("skill not found")
	}

	if err != nil {
		return nil, err
	}

	return &skill, nil
}

func ListSkills(limit *int, offset *int, startDateGTE, startDateLTE *time.Time, types []int, ordering []string) ([]Skill, int) {
	limitVal := 10
	offsetVal := 0

	if limit != nil {
		limitVal = *limit
	}

	if offset != nil {
		offsetVal = *offset
	}

	query := config.DB.Debug().Model(&Skill{}).Where("is_deleted = ?", false)

	if startDateGTE != nil {
		startDateGTEUnix := startDateGTE.Unix()
		query = query.Where("start_date >= ?", startDateGTEUnix)
	}

	if startDateLTE != nil {
		startDateLTEUnix := startDateLTE.Unix()
		query = query.Where("start_date <= ?", startDateLTEUnix)
	}

	if len(types) > 0 {
		query = query.Where("type IN ?", types)
	}

	if len(ordering) > 0 {
		for _, order := range ordering {
			query = query.Order(order)
		}
	}

	var skills []Skill
	var count int64
	query.Count(&count)

	query = query.Limit(limitVal).Offset(offsetVal)
	query.Find(&skills)

	return skills, int(count)
}

func UpdateSkill(skill Skill, name *string, startDate *time.Time, skillType *int) (*Skill, error) {
	/* update is treated as partial update */

	nameVal := skill.Name
	if name != nil {
		nameVal = *name
	}

	startDateVal := skill.StartDate
	if startDate != nil {
		startDateVal = *startDate
	}

	skillTypeVal := skill.Type
	if skillType != nil {
		skillTypeVal = *skillType
	}

	// Update the skill fields
	skill.Name = nameVal
	skill.StartDate = startDateVal
	skill.Type = skillTypeVal

	// Save the updated skill
	err := config.DB.Debug().Save(&skill).Error
	if err != nil {
		return nil, err
	}

	return &skill, nil
}
