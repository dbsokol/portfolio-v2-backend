package education

import (
	"portfolio/config"
	"time"
)

func CreateEducation(institution string, degree int, major string, startDate, endDate time.Time) (*Education, error) {
	/* create will "undelete" a deleted object if it already exists, otherwise it will create a new one */

	var education Education

	// on create, check if the object already exists but is deleted
	err := config.DB.Debug().Where(
		"institution = ? AND degress = ? AND major = ? AND is_deleted = ?",
		institution,
		degree,
		major,
		true,
	).First(&education).Error

	if err == nil && education.ID != 0 {
		education.IsDeleted = false
		education.StartDate = startDate
		education.EndDate = endDate

		err = config.DB.Save(&education).Error

		if err != nil {
			return nil, err
		}

		return &education, nil
	}

	// otherwise, create a new object
	education = Education{
		Institution: institution,
		Degree:      degree,
		Major:       major,
		StartDate:   startDate,
		EndDate:     endDate,
	}
	err = config.DB.Debug().Create(&education).Error

	if err != nil {
		return nil, err
	}

	return &education, nil
}

func DeleteEducation(uuid string) error {
	/* delete method is treated as soft delete */

	var education Education

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&education).Error

	if err != nil {
		return err
	}

	education.IsDeleted = true

	err = config.DB.Debug().Save(&education).Error

	if err != nil {
		return err
	}

	return nil
}

func GetEducation(uuid string) (*Education, error) {
	var education Education

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&education).Error

	if err != nil {
		return nil, err
	}

	return &education, nil
}

func ListEducation(limit *int, offset *int, degrees []int, ordering []string) ([]Education, int) {
	var limitVal int = 10
	var offsetVal int = 0

	if limit != nil {
		limitVal = *limit
	}

	if offset != nil {
		offsetVal = *offset
	}

	query := config.DB.Debug().Model(&Education{}).Where("is_deleted = ?", false)

	if len(degrees) > 0 {
		query = query.Where("degree IN (?)", degrees)
	}

	if len(ordering) > 0 {
		for _, order := range ordering {
			query = query.Order(order)
		}
	}

	var educations []Education
	var count int64
	query.Count(&count)

	query = query.Limit(limitVal).Offset(offsetVal)
	query.Find(&educations)

	return educations, int(count)
}

func UpdateEducation(education Education, institution *string, degree *int, major *string, startDate, endDate *time.Time) (*Education, error) {

	if institution != nil {
		education.Institution = *institution
	}

	if degree != nil {
		education.Degree = *degree
	}

	if major != nil {
		education.Major = *major
	}

	if startDate != nil {
		education.StartDate = *startDate
	}

	if endDate != nil {
		education.EndDate = *endDate
	}

	err := config.DB.Save(&education).Error

	if err != nil {
		return nil, err
	}

	return &education, nil
}
