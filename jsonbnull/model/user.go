package model

// User represents a user.
type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	SomeInfo *SomeInfo `json:"some_info"`
}

// SomeInfo contains some useful info.
type SomeInfo struct {
	Whatever string
}
