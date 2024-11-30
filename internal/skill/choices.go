package skill

import "portfolio/internal/api/choices"

// SkillType choices
var SkillTypeChoices = choices.NewIntChoices([]choices.IntChoice{
	{Public: "LANGUAGE", Private: 1},
	{Public: "FRAMEWORK", Private: 2},
	{Public: "CLOUD", Private: 3},
})

// SkillOrdering choices
var SkillOrderingChoices = choices.NewTextChoices([]choices.TextChoice{
	{Public: "CREATED_AT", Private: "created_at"},
	{Public: "-CREATED_AT", Private: "-created_at"},
	{Public: "NAME", Private: "name"},
	{Public: "-NAME", Private: "-name"},
	{Public: "START_DATE", Private: "start_date"},
	{Public: "-START_DATE", Private: "-start_date"},
})
