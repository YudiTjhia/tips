package dbh

import (
	"testing"
	"time"
)

func Test_Dbh(t *testing.T) {

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

	dbConfig := NewDbConfigForRDMS(driverName, host, user, password, port, dbName,
		dbType, maxConns, maxIdle, maxLifeTime)

	dbh := NewDbh(dbConfig)
	err := dbh.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}

}
