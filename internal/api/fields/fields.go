package fields

import (
	"portfolio/internal/api/fields/charfield"
	"portfolio/internal/api/fields/choicefield"
	"portfolio/internal/api/fields/datefield"
)

var (
	CharField       = charfield.CharField
	DateField       = datefield.DateField
	IntChoiceField  = choicefield.IntChoiceField
	TextChoiceField = choicefield.TextChoiceField
)
