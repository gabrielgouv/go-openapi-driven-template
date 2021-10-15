package domain

type Echo struct {
	Date         string
	Message      string
	RandomNumber float32
}

type EchoRepository interface {
	FindEchoDate() string
}
