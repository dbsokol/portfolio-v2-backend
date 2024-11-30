package parameters

import (
	"portfolio/internal/api/parameters/charparam"
	"portfolio/internal/api/parameters/choicelistparam"
	"portfolio/internal/api/parameters/dateparam"
	"portfolio/internal/api/parameters/intparam"
)

// Re-export functions for convenience
var (
	CharParam           = charparam.CharParam
	DateParam           = dateparam.DateParam
	IntChoiceListParam  = choicelistparam.IntChoiceListParam
	IntParam            = intparam.IntParam
	TextChoiceListParam = choicelistparam.TextChoiceListParam
)
