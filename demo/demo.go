package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
type demo struct {
	Name string
	A *Animal
}

type Animal struct {
	Animal string
}
func main(){
	logger,_:=zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	d:=demo{Name:"123123",A:&Animal{Animal:"dog"}}
	logger.Info(
		"msg",
		zap.Any("key","value"),
		zap.Any("d",d),
		)
}

func get() demo{
	return demo{}
}