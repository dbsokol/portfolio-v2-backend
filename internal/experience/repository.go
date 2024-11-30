package experience

import (
	"portfolio/config"
)

func ListExperiences() ([]Experience, error) {
	var experiences []Experience
	err := config.DB.Preload("Responsibilities").Find(&experiences).Error
	if err != nil {
		return nil, err
	}
	return experiences, nil
}
