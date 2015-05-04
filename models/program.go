package models

// Program しょぼかるのProgに相当.
type Program struct {
	ID      int
	AnimeID int `db:"index"`
	Title   string
	Chapter string
}
