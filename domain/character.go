package domain

type Character struct {
	Data CharacterData `json:"data"`
}

type CharacterData struct {
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"imageUrl,omitempty"`
}

func (c *Character) IsEmpty() bool {
	return c.Data.Name == "" && c.Data.ImageURL == ""
}
