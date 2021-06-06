package dao

import (
	"database/sql"
	"fmt"
	"log"
)

func (d *Dao) AddUser() {
	//d.db.insert(.....)
}

func (d *Dao) GetUser(id int) error {
	var name string
	err := d.db.QueryRow("select * from users where id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}
