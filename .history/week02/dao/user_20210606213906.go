package dao

import (
	"database/sql"
	"homework/week02/model"
	"log"
)

func (d *Dao) AddUser() {
	//d.db.insert(.....)
}

func (d *Dao) GetUser(id int) (model.User, error) {
	rows := d.db.QueryRow("select username,address,age,mobile,sex from t_user where id = ?", id)

	var user model.User
	err := rows.Scan(&user.Username, &user.Address, &user.Age, &user.Mobile, &user.Sex)
	if err != nil {
			log.Fatal(err)
		}
	}
	return user, err
}

func (d *Dao) GetUserName(id int) (model.User, error) {
	rows := d.db.QueryRow("select username,address,age,mobile,sex from t_user where id = ?", id)

	var user model.User
	err := rows.Scan(&user.Username, &user.Address, &user.Age, &user.Mobile, &user.Sex)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	return user, err
}
