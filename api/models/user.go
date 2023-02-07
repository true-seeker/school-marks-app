package models

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) GetByID(id string) (*User, error) {
	return &User{
		ID:        "123",
		Name:      "321",
		BirthDay:  "456",
		Gender:    "654",
		PhotoURL:  "789",
		Time:      0,
		Active:    false,
		UpdatedAt: 0,
	}, nil
}
