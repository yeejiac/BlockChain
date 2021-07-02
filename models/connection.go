package models

type Connection struct {
	Sequence int  `json:"sequence"`
	Status   bool `json:"status"`
}

type Mode int

const (
	RELEASE Mode = iota
	DEVELOPE
	TEST
)
