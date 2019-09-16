package roksandb

const (
	tableSeries = "series"
)

// Series contains details of series
type Series struct {
	ID      string `json:"id" bson:"id"`
	BrandID string `json:"brandId" bson:"brandId"`
	Name    string `json:"name" bson:"name"`
	Lang    string `json:"lang" bson:"lang"`
	ImgPath string `json:"imgPath" bson:"imgPath"`
}

// InsertSeries inserts series
func InsertSeries(series Series) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+tableSeries+"(id,brandId,name,lang,imgPath) VALUES(?,?,?,?,?)", series.ID, series.BrandID, series.Name, series.Lang, series.ImgPath)
	return
}

// FindSerieses finds series
func FindSerieses(lang string) (series []Series, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM "+tableSeries+" WHERE lang=?", lang)
	for rows.Next() {
		var (
			id      string
			brandID string
			name    string
			lang    string
			imgPath string
		)
		if err = rows.Scan(&id, &name, &lang, &imgPath); err != nil {
			return
		}
		series = append(series, Series{
			Name:    name,
			Lang:    lang,
			BrandID: brandID,
			ID:      id,
			ImgPath: imgPath,
		})
	}
	return
}

// FindSeries finds series
func FindSeries(name string) (series Series, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM "+tableSeries+" WHERE name=?", name)
	var (
		id      string
		brandID string
		lang    string
		imgPath string
	)
	if err = row.Scan(&id, &brandID, &name, &lang, &imgPath); err != nil {
		return
	}
	series = Series{
		ID:      id,
		BrandID: brandID,
		Name:    name,
		Lang:    lang,
		ImgPath: imgPath,
	}
	return
}

// UpdateSeries update series details depends of the id
func UpdateSeries(id string, series Series) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "UPDATE " + tableSeries + " SET name=?, lang=?, brandId=?, series=?, imgPath=? WHERE id=?"
	_, err = db.Exec(command, series.Name, series.Lang, series.BrandID, series.ID, series.ImgPath, id)
	return
}

// DeleteSeries delete series
func DeleteSeries(id string) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "DELETE FROM " + tableSeries + " WHERE id=?"
	_, err = db.Exec(command, id)
	return
}
