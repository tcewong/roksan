package roksandb

import ()

const (
	tableShowroom = "showroom"
)

// Showroom contains data structure of showroom
type Showroom struct {
	Name  string `json:"name" bson:"name"`
	Addr  string `json:"addr" bson:"addr"`
	Phone string `json:"phone" bson:"phone"`
	Time  string `json:"time" bson:"time"`
	Lang  string `json:"lang" bson:"lang"`
}

// InsertShowroom inserts showroom
func InsertShowroom(showroom Showroom) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+tableShowroom+"(name,addr,phone,time,lang) VALUES(?,?,?,?,?)", showroom.Name, showroom.Addr, showroom.Phone, showroom.Time, showroom.Lang)
	return
}

// FindShowrooms finds showroom
func FindShowrooms(lang string) (showrooms []Showroom, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM "+tableShowroom+" WHERE lang=?", lang)
	for rows.Next() {
		var (
			name  string
			addr  string
			phone string
			time  string
			lang  string
		)
		if err = rows.Scan(&name, &addr, &phone, &time, &lang); err != nil {
			return
		}
		showrooms = append(showrooms, Showroom{
			Name:  name,
			Addr:  addr,
			Phone: phone,
			Time:  time,
			Lang:  lang,
		})
	}
	return
}

// UpdateShowroom update showroom details depends of the name
func UpdateShowroom(name string, showroom Showroom) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "UPDATE " + tableShowroom + " SET name=?, addr=?, phone=?, time=?, lang=? WHERE name=?"
	_, err = db.Exec(command, showroom.Name, showroom.Addr, showroom.Phone, showroom.Time, showroom.Lang, name)
	return
}

// DeleteShowrooms delete showrooms
func DeleteShowrooms(name string) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "DELETE FROM " + tableShowroom + " WHERE name=?"
	_, err = db.Exec(command, name)
	return
}
