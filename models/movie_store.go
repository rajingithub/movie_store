package models

type Movie struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate *string `json:"release_date"`
}

type Actor struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Image *string `json:"image"`
}
