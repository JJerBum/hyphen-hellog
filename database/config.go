package database

import (
	"hyphen-hellog/cerrors/exception"
	"os"

	"github.com/joho/godotenv"
)

type config interface {
	Get(string) string
}

type configImpl struct{}

func (configImpl) Get(key string) string {
	return os.Getenv(key)
}

func newConfig(filenames ...string) config {
	err := godotenv.Load(filenames...)
	exception.Sniff(err)
	return configImpl{}
}
