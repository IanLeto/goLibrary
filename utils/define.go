package utils

type Task interface {
	Run() error
	Stop() error
}
