package dao

func(d *Dao) AddUser(){
    //d.db.insert(.....)
}

func(d *Dao) GetUser(int id, string name) error{
	var name string
	err = d.db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}