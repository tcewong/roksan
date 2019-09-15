package roksandb

const (
	tableBrand = "brand"
)

// Brand contains data structure of showroom
type Brand struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Lang    string `json:"lang" bson:"lang"`
	ImgPath string `json:"imgPath" bson:"imgPath"`
}

// InsertBrand inserts brand
func InsertBrand(brand Brand) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+tableBrand+"(id,name,lang,imgPath) VALUES(?,?,?,?)", brand.ID, brand.Name, brand.Lang, brand.ImgPath)
	return
}

// FindBrands finds brands
func FindBrands(lang string) (brands []Brand, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM "+tableBrand+" WHERE lang=?", lang)
	for rows.Next() {
		var (
			id      string
			name    string
			lang    string
			imgPath string
		)
		if err = rows.Scan(&id, &name, &lang, &imgPath); err != nil {
			return
		}
		brands = append(brands, Brand{
			ID:      id,
			Name:    name,
			Lang:    lang,
			ImgPath: imgPath,
		})
	}
	return
}

// FindBrand finds brand
func FindBrand(id string) (brand Brand, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM "+tableBrand+" WHERE id=?", id)
	var (
		name    string
		lang    string
		imgPath string
	)
	if err = row.Scan(&id, &name, &lang, &imgPath); err != nil {
		return
	}
	brand = Brand{
		ID:      id,
		Name:    name,
		Lang:    lang,
		ImgPath: imgPath,
	}
	return
}

// UpdateBrand update brand details depends of the id
func UpdateBrand(id string, brand Brand) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "UPDATE " + tableBrand + " SET id=?, name=?, lang=?, imgPath=? WHERE id=?"
	_, err = db.Exec(command, brand.ID, brand.Name, brand.Lang, brand.ImgPath, id)
	return
}

// DeleteBrands delete showrooms
func DeleteBrands(id string) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "DELETE FROM " + tableBrand + " WHERE id=?"
	_, err = db.Exec(command, id)
	return
}
