package unittest

import (
	"errors"
	"strconv"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
}

func NewDbConfig(host, user, password, port, dbName string) (*DbConfig, error) {
	if host == "" {
		return nil, errors.New("host is empty")
	}
	if user == "" {
		return nil, errors.New("user is empty")
	}
	if password == "" {
		return nil, errors.New("password is empty")
	}
	if port == "" {
		return nil, errors.New("port is empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("port is not a number")
	}

	if dbName == "" {
		return nil, errors.New("db name is empty")
	}
	return &DbConfig{
		Host:     host,
		User:     user,
		Password: password,
		Port:     port,
		DbName:   dbName,
	}, nil
}
