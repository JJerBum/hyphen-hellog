package database

import "hyphen-hellog/database/repository"

func Get() *repository.DBType {
	return repository.Get()
}
