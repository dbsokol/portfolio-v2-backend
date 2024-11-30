package contact

import (
	"log"
	"portfolio/config"
)

func CreateContact(name, value string, url *string) (*Contact, error) {
	var contact Contact

	config.DB.Debug().Where("name = ? AND is_deleted = ?", name, true).First(&contact)

	if contact.ID != 0 {
		contact.IsDeleted = false
		contact.Value = value
		contact.URL = *url

		err := config.DB.Save(&contact).Error

		if err != nil {
			return nil, err
		}

		return &contact, nil
	}

	contact = Contact{
		Name:  name,
		Value: value,
		URL:   *url,
	}

	err := config.DB.Debug().Create(&contact).Error

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func DeleteContact(uuid string) error {
	var contact Contact

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&contact).Error

	if err != nil {
		return err
	}

	err = config.DB.Debug().Delete(&contact).Error

	if err != nil {
		return err
	}

	return nil
}

func GetContact(uuid string) (*Contact, error) {
	var contact Contact

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&contact).Error

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func ListContacts(limit, offset *int, ordering []string, nameIContains *string) ([]Contact, int) {
	var limitVal int = 10
	var offsetVal int = 0
	var nameIContainsVal string

	if limit != nil {
		limitVal = *limit
	}

	if offset != nil {
		offsetVal = *offset
	}

	if nameIContains != nil {
		nameIContainsVal = *nameIContains
	}

	query := config.DB.Debug().Model(&Contact{}).Where("is_deleted = ?", false)

	log.Print(nameIContainsVal)

	if nameIContainsVal != "" {
		query = query.Where("name LIKE ?", "%"+nameIContainsVal+"%")
	}

	if len(ordering) > 0 {
		for _, order := range ordering {
			query = query.Order(order)
		}
	}

	var contacts []Contact
	var count int64

	query.Count(&count)

	err := query.Limit(limitVal).Offset(offsetVal).Find(&contacts).Error

	if err != nil {
		return nil, 0
	}

	return contacts, int(count)
}

func UpdateContact(uuid, name, value, url string) (*Contact, error) {
	var contact Contact

	err := config.DB.Debug().Where("uuid = ?", uuid).First(&contact).Error

	if err != nil {
		return nil, err
	}

	contact.Name = name
	contact.Value = value
	contact.URL = url

	err = config.DB.Debug().Save(&contact).Error

	if err != nil {
		return nil, err
	}

	return &contact, nil
}
