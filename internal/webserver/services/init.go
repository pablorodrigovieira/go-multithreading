package services

import (
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/configs"
)

var envConfig *configs.Configuration

func Init(c *configs.Configuration) {
	envConfig = c
}
