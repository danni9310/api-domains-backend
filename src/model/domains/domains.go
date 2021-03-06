package domains

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func Insert(db *sql.DB, id string, requestDomain string, requestHash string, previousGrade string) {
	var timeNow int64 = time.Now().UnixNano() / int64(time.Millisecond)
	sqlStatement := `INSERT INTO domainsV2 (id, requestDomain, requestHash, previousGrade, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, id, requestDomain, requestHash, previousGrade, timeNow, timeNow)
	if err != nil {
		panic(err)
	}
}

func Find(db *sql.DB, id string) (string, string, int64, error) {

	sqlStatement := `SELECT requestHash, previousGrade, updated_at FROM domainsV2 WHERE id = $1;`
	var requestHash string
	var previousGrade string
	var updated_at int64
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&requestHash, &previousGrade, &updated_at); err {
	case sql.ErrNoRows:
		return "", "", time.Now().UnixNano() % 1e6 / 1e3, err
	case nil:
		return requestHash, previousGrade, updated_at, err
	default:
		panic(err)
	}
}

type DomainsRequests struct {
	Item []string
}

func FindIById(db *sql.DB) (DomainsRequests, error) {
	var domainsRequests DomainsRequests
	sqlStatement := `SELECT id FROM domainsV2;`
	rows, err := db.Query(sqlStatement)

	defer func() {
		if err != nil {
			fmt.Printf("Error finding by id: %v", err)
		}
		rows.Close()
	}()

	if err != nil {
		return domainsRequests, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return domainsRequests, err
		}
		domainsRequests.Item = append(domainsRequests.Item, id)
	}
	return domainsRequests, nil

}

func Update(db *sql.DB, id string, requestHash string, update_at int64, grade string) {
	var err error
	stmt, err := db.Prepare("UPDATE domainsV2 SET requestHash = $1, updated_at = $2, previousGrade =$3 WHERE id=$4")
	defer func() {
		if err != nil {
			fmt.Printf("Error in update: %v", err)
		}
		stmt.Close()
	}()

	if err != nil {
		return
	}

	res, err := stmt.Exec(requestHash, update_at, grade, id)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	fmt.Println(affect, "rows changed")
}
