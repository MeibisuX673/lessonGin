package model

type ErrorInterface interface {
	Error() string
	GetStatus() int
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *Error) Error() string {

	return e.Message

}

func (e *Error) GetStatus() int {

	return e.Status

}
