package choices

import "errors"

type TextChoice struct {
	Public  string
	Private string
}

type IntChoice struct {
	Public  string
	Private int
}

type TextChoices struct {
	items []TextChoice
}

type IntChoices struct {
	items []IntChoice
}

func NewTextChoices(choices []TextChoice) TextChoices {
	return TextChoices{items: choices}
}

func NewIntChoices(choices []IntChoice) IntChoices {
	return IntChoices{items: choices}
}

func (c TextChoices) Validate(public string) (TextChoice, error) {
	for _, choice := range c.items {
		if choice.Public == public {
			return choice, nil
		}
	}
	return TextChoice{}, errors.New("choice must be one of: " + public)
}

func (c IntChoices) Validate(public string) (IntChoice, error) {
	for _, choice := range c.items {
		if choice.Public == public {
			return choice, nil
		}
	}
	return IntChoice{}, errors.New("choice must be one of: " + public)
}

func (c TextChoices) GetPrivate(public string) (string, error) {
	choice, err := c.Validate(public)
	if err != nil {
		return "", err
	}
	return choice.Private, nil
}

func (c IntChoices) GetPrivate(public string) (int, error) {
	choice, err := c.Validate(public)
	if err != nil {
		return 0, err
	}
	return choice.Private, nil
}

func (c TextChoices) GetPublicList() []string {
	publicList := make([]string, len(c.items))
	for i, choice := range c.items {
		publicList[i] = choice.Public
	}
	return publicList
}

func (c IntChoices) GetPublicList() []string {
	publicList := make([]string, len(c.items))
	for i, choice := range c.items {
		publicList[i] = choice.Public
	}
	return publicList
}

func (c TextChoices) GetPublic(private string) string {
	for _, choice := range c.items {
		if choice.Private == private {
			return choice.Public
		}
	}
	return ""
}

func (c IntChoices) GetPublic(private int) string {
	for _, choice := range c.items {
		if choice.Private == private {
			return choice.Public
		}
	}
	return ""
}
