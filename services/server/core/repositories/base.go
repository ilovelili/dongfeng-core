// Package repositories is the database access layer
// We use dapper lib to handle sql ORM
package repositories

import (
	"database/sql"

	"github.com/ilovelili/dongfeng/core/services/utils"
	"github.com/olivere/dapper"

	"sync"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbclient struct {
	db      *sql.DB
	session *dapper.Session
}

var (
	instance dbclient
	once     sync.Once
)

// dbclient singleton db client
func client() dbclient {
	once.Do(func() {
		config := utils.GetConfig()
		db, err := sql.Open("mysql", connectionstring())
		if err == nil {
			instance = dbclient{
				db:      db,
				session: dapper.New(db).Dialect(dapper.MySQL).Debug(config.MySQL.AllowDebug),
			}
		}
	})

	return instance
}

func session() *dapper.Session {
	return client().session
}

// insertTx insert with a transaction
func insertTx(entity interface{}) (err error) {
	db, session := client().db, client().session
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// Insert
	err = session.InsertTx(tx, entity)
	if err != nil {
		tx.Rollback()
		return
	}

	// Commit transaction
	err = tx.Commit()
	return
}

// updateTx update with a transaction
func updateTx(entity interface{}) (err error) {
	db, session := client().db, client().session
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// Update
	err = session.UpdateTx(tx, entity)
	if err != nil {
		tx.Rollback()
		return
	}

	// Commit transaction
	err = tx.Commit()
	return
}

// mysql connection string
func connectionstring() string {
	config := utils.GetConfig()
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	return fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", config.MySQL.User, config.MySQL.Password, config.MySQL.Host, config.MySQL.DataBase)
}

// Table dapper.Q Wrapper
func Table(tablename string) *dapper.Query {
	return dapper.Q(dapper.MySQL, tablename)
}

// norows no rows error sometimes can be ignored
func norows(err error) bool {
	return err == sql.ErrNoRows
}
