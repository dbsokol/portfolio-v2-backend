package contact

import "portfolio/internal"

type Contact struct {
	internal.Model
	Name  string `gorm:"type:varchar(35);not null;unique" json:"name"`
	Value string `gorm:"type:varchar(100);" json:"value"`
	URL   string `gorm:"type:varchar(100);" json:"url"`
}
