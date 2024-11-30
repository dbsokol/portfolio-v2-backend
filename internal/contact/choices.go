package contact

import "portfolio/internal/api/choices"

var ContactOrderingChoices = choices.NewTextChoices([]choices.TextChoice{
	{Public: "CREATED_AT", Private: "created_at"},
	{Public: "-CREATED_AT", Private: "-created_at"},
	{Public: "NAME", Private: "name"},
	{Public: "-NAME", Private: "-name"},
})
