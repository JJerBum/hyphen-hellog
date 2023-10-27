package initializer

import (
	"hyphen-hellog/cerrors/exception"

	"github.com/joho/godotenv"
)

func LoadEnv(filename ...string) {
	err := godotenv.Load(filename...)
	exception.Sniff(err)
}
