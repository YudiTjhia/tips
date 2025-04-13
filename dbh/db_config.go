package dbh

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type DbType int

const (
	DbType_Postgre DbType = 1
)

type DbConfig struct {
	Host        string
	User        string
	Password    string
	Port        string
	DbName      string
	MaxConns    int
	MaxIdle     int
	MaxLifeTime time.Duration

	DbType     DbType
	DriverName string
}

func NewDbConfigForRDMS(driverName, host, user, password, port, dbName string,
	dbType DbType,
	maxConns, maxIdle int,
	maxLifeTime time.Duration,
) (*DbConfig, error) {

	if driverName == "" {
		return nil, errors.New("driver name is empty")
	}
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
	if maxConns <= 0 {
		return nil, errors.New("maxConns <= 0")
	}
	if maxIdle <= 0 {
		return nil, errors.New("maxIdle <= 0")
	}
	if maxLifeTime <= 0 {
		return nil, errors.New("maxLifeTime <= 0")
	}
	if maxLifeTime < 1*time.Minute {
		return nil, errors.New("maxLifeTime < 1 minute")
	}

	return &DbConfig{
		Host:        host,
		User:        user,
		Password:    password,
		Port:        port,
		DbName:      dbName,
		MaxConns:    maxConns,
		MaxIdle:     maxIdle,
		MaxLifeTime: maxLifeTime,

		DbType:     dbType,
		DriverName: driverName,
	}, nil
}

func NewDbConfigForRedis(host, password, port string) (*DbConfig, error) {

	if host == "" {
		return nil, errors.New("host is empty")
	}
	if password == "" {
		return nil, errors.New("password is empty")
	}
	if port == "" {
		return nil, errors.New("port is empty")
	}
	return &DbConfig{
		Host:     host,
		Password: password,
		Port:     port,
	}, nil

}

func (d DbConfig) ConnString() string {
	switch d.DbType {
	case DbType_Postgre:
		dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			d.User,
			d.Password,
			d.Host,
			d.Port,
			d.DbName)
		return dbConnStr
	default:
		panic(fmt.Sprintf("unsupported db type: %d", d.DbType))
	}
}
