## 第二周作业

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

   ### dao.go

```go
package dao

import (
	"database/sql"
	"homework/week02/model"
	// "log"

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
			//当遇到一个 sql.ErrNoRows 的时候，这不是一个错误， 不应该 Wrap 这个 error，也不能抛给上层
		} else {
			//log.Fatal(err)
			return user, errors.Wrap(err, "get user name failed")
		}
	}
	return user, nil
}
```





### main.go

```go
package main

import (
	"fmt"
	"homework/week02/dao"
	"log"
)

func main() {

	dao := dao.NewDao()
	user, err := dao.GetUserName(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user name: ", user.Username)
}

```