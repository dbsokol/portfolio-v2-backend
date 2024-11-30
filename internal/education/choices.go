package education

import "portfolio/internal/api/choices"

var DegreeChoices = choices.NewIntChoices([]choices.IntChoice{
	{Public: "BACHELORS", Private: 1},
	{Public: "MASTERS", Private: 2},
})

var EducationOrderingChoices = choices.NewTextChoices([]choices.TextChoice{
	{Public: "START_DATE", Private: "start_date"},
	{Public: "-START_DATE", Private: "-start_date"},
	{Public: "END_DATE", Private: "end_date"},
	{Public: "-END_DATE", Private: "-end_date"},
	{Public: "DEGREE", Private: "degree"},
	{Public: "-DEGREE", Private: "-degree"},
	{Public: "CREATED_AT", Private: "created_at"},
	{Public: "-CREATED_AT", Private: "-created_at"},
})
