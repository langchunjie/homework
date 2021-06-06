package dao

import (
	"database/sql"
	"fmt"
	"log"
)

func (d *Dao) AddUser() {
	//d.db.insert(.....)
}

func (d *Dao) GetUser(id int) (user model.User, e error {
	var name string
	err := d.db.QueryRow("select * from users where id = ?", id).Scan(&name)
	rows := db.QueryRow("select username,address,age,mobile,sex from t_user where id = ? and username = ?",id,name)

	var user model.User
	err := rows.Scan(&user.Username,&user.Address,&user.Age,&user.Mobile,&user.Sex)	
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}
