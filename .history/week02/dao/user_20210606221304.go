package dao

import (
	"database/sql"
	"homework/week02/model"
	"log"

	"github.com/pkg/errors"
)

func (d *Dao) AddUser() {
	//d.db.insert(.....)
}

func (d *Dao) GetUser(id int) (model.User, error) {
	rows := d.db.QueryRow("select username,address,age,mobile,sex from t_user where id = ?", id)

	var user model.User
	err := rows.Scan(&user.Username, &user.Address, &user.Age, &user.Mobile, &user.Sex)
	if err != nil {
		//log.Fatal(err)
		return user, errors.Wrap(err, "get user failed")
	}
	return user, nil
}

func (d *Dao) GetUserName(id int) (model.User, error) {
	var user model.User
	err := d.db.QueryRow("select username,address,age,mobile,sex from t_user where id = ?", id).Scan(&user.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层
		} else {
			//log.Fatal(err)
			return user, errors.Wrap(err, "get user name failed")
		}
	}
	return user, err
}
