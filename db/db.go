package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Config DBConfig `yaml:",inline"`
	Conn   *sqlx.DB
}

type DBConfig struct {
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Name     string `yaml:"name" validate:"required"`
}

func (db *DB) Init() {
	db.newDB()
}
