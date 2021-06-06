package dao

import (
	"database/sql"	
	"homework/week02/database"
)


type Dao struct {
	db *sql.DB
}

func NewDao() *Dao {
	return &Dao{
		db: database.GetDataBase(),
	}
}




