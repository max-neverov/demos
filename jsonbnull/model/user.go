package model

// User represents a user.
type User struct {
	Name     string    `json:"name"      db:"name"`
	Age      int       `json:"age"       db:"age"`
	SomeInfo *SomeInfo `json:"some_info" db:"some_info"`
}

// SomeInfo contains some useful info.
type SomeInfo struct {
	Whatever string
}
