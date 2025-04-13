package dbh

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"tips/mjson"
	"tips/neth"
)

const SqlNoRows = "sql: no rows in result set"

type Dbh struct {
	DbConfig DbConfig
	DbClient *sqlx.DB
}

func NewDbh(dbConfig DbConfig) Dbh {
	return Dbh{
		DbConfig: dbConfig,
	}
}

func (d *Dbh) Connect() error {

	connString := d.DbConfig.ConnString()
	dbClient, err := sqlx.Connect(d.DbConfig.DriverName, connString)

	nethelper := neth.Neth{}
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v %s", err, nethelper.MaskDSN(connString))
	}

	dbClient.SetMaxOpenConns(d.DbConfig.MaxConns)
	dbClient.SetMaxIdleConns(d.DbConfig.MaxIdle)
	dbClient.SetConnMaxLifetime(d.DbConfig.MaxLifeTime)

	log.Printf("Connected to database %v %s", d.DbConfig.DbName, nethelper.MaskDSN(connString))
	d.DbClient = dbClient
	return nil

}

func (d Dbh) MakeDbUUID(db *sqlx.DB) {
	sql := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`
	db.MustExec(sql)
}

func (d Dbh) RowsAffectedZero(result sql.Result, err error, action string) error {
	if err != nil {
		return fmt.Errorf("error while %s: %w", action, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected count: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("%s: no rows were affected", action)
	}
	return nil
}

func (d Dbh) DbNamedExec(sql string, arg interface{}) (sql.Result, error) {
	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	result, err := d.DbClient.NamedExec(sql, arg)
	if err != nil {
		log.Println("DbExec error=", err.Error())
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("DbExec error=", err.Error())
		return nil, err
	}
	if rows == 0 {
		return nil, fmt.Errorf(SqlNoRows)
	}
	return result, nil
}

func (d Dbh) RowsToMap(rows *sqlx.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		rowMap := make(map[string]interface{})
		err := rows.MapScan(rowMap)
		if err != nil {
			return nil, err
		}
		result = append(result, rowMap)
	}
	return result, nil
}

func (d Dbh) DbNamedQuery(result []map[string]interface{}, sql string, arg interface{}) error {

	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	rows, err := d.DbClient.NamedQuery(sql, arg)
	if err != nil {
		log.Println("DbNamedQuery error=", err.Error())
		return err
	}
	res, err := d.RowsToMap(rows)
	if err != nil {
		log.Println("DbNamedQuery error=", err.Error())
		return err
	}
	result = res
	return nil
}

func (d Dbh) DbOrTrNamedExec(
	tr *sqlx.Tx,
	sql string,
	arg interface{}) (sql.Result, error) {

	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	if tr != nil {
		return d.TrNamedExec(tr, sql, arg)

	} else {
		return d.DbNamedExec(sql, arg)
	}
}

func (d Dbh) TrNamedExec(tr *sqlx.Tx, sql string, arg interface{}) (sql.Result, error) {

	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	result, err := tr.NamedExec(sql, arg)
	if err != nil {
		log.Println("TrNamedExec error=", err.Error())
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("DbExec error=", err.Error())
		return nil, err
	}
	if rows == 0 {
		return nil, fmt.Errorf(SqlNoRows)
	}
	log.Println("affectedrows=", rows)
	return result, nil
}

func (d Dbh) DbGet(result interface{}, sql string, arg ...interface{}) error {

	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	err := d.DbClient.Get(result, sql, arg...)
	if err != nil {
		log.Println("DbGet error=", err.Error())
		return err
	}
	log.Println("result=", mjson.JsonStringify(result))
	return nil
}

func (d Dbh) DbExec(sql string, arg ...interface{}) (sql.Result, error) {

	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))

	result, err := d.DbClient.Exec(sql, arg...)
	if err != nil {
		log.Println("DbExec error=", err.Error())
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("DbExec error=", err.Error())
		return nil, err
	}
	log.Println("affectedrows=", rows)
	return result, nil
}

func (d Dbh) TrExec(tr *sqlx.Tx, sql string, arg ...interface{}) (sql.Result, error) {

	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	result, err := tr.Exec(sql, arg...)
	if err != nil {
		log.Println("TrExec error=", err.Error())
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("TrExec error=", err.Error())
		return nil, err
	}
	log.Println("affectedrows=", rows)
	return result, nil
}

func (d Dbh) DbSelect(db *sqlx.DB, result interface{}, sql string, arg ...interface{}) error {
	log.Println("sql=", sql)
	log.Println("param=", mjson.JsonStringify(arg))
	err := db.Select(result, sql, arg...)
	if err != nil {
		log.Println("DbSelect error=", err.Error())
		return err
	}
	return nil
}
