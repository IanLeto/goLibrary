package logrusDemo

import "go.uber.org/zap"

func NewLogger() {
	zap.NewProduction()
}

func InitLogger() {

}
