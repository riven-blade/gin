package contract

const IDKey = "nice:id"

type IDService interface {
	NewID() string
}
