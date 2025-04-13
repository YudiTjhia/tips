package dbh

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_DbConfigNegative(t *testing.T) {

	_, err := NewDbConfigForRDMS("", "", "", "", "", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "driver name is empty")

	_, err = NewDbConfigForRDMS("postgres", "", "", "", "", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "host is empty")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "", "", "", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "user is empty")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "", "", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "password is empty")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "port is empty")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"abc", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "port is not a number")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"9896", "",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "db name is empty")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"9896", "account_db",
		DbType_Postgre, 0, 0, 0)
	assert.Equal(t, err.Error(), "maxConns <= 0")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"9896", "account_db",
		DbType_Postgre, 1, 0, 0)
	assert.Equal(t, err.Error(), "maxIdle <= 0")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"9896", "account_db",
		DbType_Postgre, 1, 1, 0)
	assert.Equal(t, err.Error(), "maxLifeTime <= 0")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"9896", "account_db",
		DbType_Postgre, 1, 1, 0)
	assert.Equal(t, err.Error(), "maxLifeTime <= 0")

	_, err = NewDbConfigForRDMS("postgres", "192.168.101.18", "account_db", "123456",
		"9896", "account_db",
		DbType_Postgre, 1, 1, 2*time.Second)
	assert.Equal(t, err.Error(), "maxLifeTime < 1 minute")
}

func Test_DbConfig(t *testing.T) {

	driverName := "postgres"
	host := "192.168.101.18"
	user := "account_db_user"
	password := "1YWNjb3VudF91c2VyOjk4OTY="
	port := "9896"
	dbName := "account_db"
	dbType := DbType_Postgre
	maxConns := 100
	maxIdle := 10
	maxLifeTime := 24 * time.Hour

	dbConfig, err := NewDbConfigForRDMS(driverName, host, user, password, port, dbName,
		dbType, maxConns, maxIdle, maxLifeTime)
	if err != nil {
		t.Fatalf("Failed to create DB config: %v", err)
	}

	assert.Equal(t, driverName, dbConfig.DriverName)
	assert.Equal(t, host, dbConfig.Host)
	assert.Equal(t, user, dbConfig.User)
	assert.Equal(t, password, dbConfig.Password)
	assert.Equal(t, port, dbConfig.Port)
	assert.Equal(t, dbName, dbConfig.DbName)
	assert.Equal(t, dbType, dbConfig.DbType)
	assert.Equal(t, maxConns, dbConfig.MaxConns)
	assert.Equal(t, maxIdle, dbConfig.MaxIdle)
	assert.Equal(t, maxLifeTime, dbConfig.MaxLifeTime)

}
