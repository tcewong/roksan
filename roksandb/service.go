package roksandb

const (
	tableService = "service"
)

// Service contains data structure of service
type Service struct {
	Name  string `json:"name" bson:"name"`
	Addr  string `json:"addr" bson:"addr"`
	Phone string `json:"phone" bson:"phone"`
	Time  string `json:"time" bson:"time"`
	Lang  string `json:"lang" bson:"lang"`
}

// InsertService inserts service
func InsertService(service Service) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+tableService+"(name,addr,phone,time,lang) VALUES(?,?,?,?,?)", service.Name, service.Addr, service.Phone, service.Time, service.Lang)
	return
}

// FindServices finds service
func FindServices(lang string) (services []Service, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM "+tableService+" WHERE lang=?", lang)
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
		services = append(services, Service{
			Name:  name,
			Addr:  addr,
			Phone: phone,
			Time:  time,
			Lang:  lang,
		})
	}
	return
}

// UpdateService update service details depends of the name
func UpdateService(name string, service Service) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "UPDATE " + tableService + " SET name=?, addr=?, phone=?, time=?, lang=? WHERE name=?"
	_, err = db.Exec(command, service.Name, service.Addr, service.Phone, service.Time, service.Lang, name)
	return
}

// DeleteServices delete services
func DeleteServices(name string) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "DELETE FROM " + tableService + " WHERE name=?"
	_, err = db.Exec(command, name)
	return
}
