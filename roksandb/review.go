package roksandb

const (
	tableReview = "review"
)

// Review contains data structure of review
type Review struct {
	Name    string `json:"name" bson:"name"`
	Lang    string `json:"lang" bson:"lang"`
	Details string `json:"details" bson:"details"`
	URL     string `json:"url" bson:"url"`
}

// InsertReview inserts review
func InsertReview(review Review) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+tableReview+"(name,lang,details,url) VALUES(?,?,?,?)", review.Name, review.Lang, review.Details, review.URL)
	return
}

// FindReviews finds reviews
func FindReviews(lang string) (reviews []Review, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM "+tableReview+" WHERE lang=?", lang)
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
		reviews = append(reviews, Review{
			Name:    name,
			Lang:    lang,
			Details: "",
			URL:     "",
		})
	}
	return
}

// FindReview finds review
func FindReview(name string) (review Review, err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM "+tableReview+" WHERE name=?", name)
	var (
		details string
		url     string
		lang    string
	)
	if err = row.Scan(&name, &details, &url, &lang); err != nil {
		return
	}
	review = Review{
		Name:    name,
		Lang:    lang,
		Details: details,
		URL:     url,
	}
	return
}

// UpdateReview update review details depends of the id
func UpdateReview(id string, review Review) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "UPDATE " + tableReview + " SET name=?, lang=?, details=?, url=? WHERE id=?"
	_, err = db.Exec(command, review.Name, review.Lang, review.Details, review.URL, id)
	return
}

// DeleteReviews delete reviews
func DeleteReviews(id string) (err error) {
	db, err := db()
	if err != nil {
		return
	}
	defer db.Close()
	command := "DELETE FROM " + tableBrand + " WHERE id=?"
	_, err = db.Exec(command, id)
	return
}
