package models

type Model interface {
	Validate() (bool, error)
}
