package models

type Connection struct {
	Sequence int  `json:"sequence"`
	Status   bool `json:"status"`
}
