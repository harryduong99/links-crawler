package linksRepo

import (
	"context"
	"database/sql"
	"links-crawler/driver"
	"links-crawler/models"
	"log"
	"strconv"
	"time"
)

type PostgresLinksRepo struct{}

func (linksRepo *PostgresLinksRepo) StoreLink(link models.Link) error {
	query := "INSERT INTO links(url, crawled, domain) VALUES ($1, $2, $3)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := driver.PostgresDB.Client.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, link.Url, false, link.Domain)
	if err != nil {
		log.Printf("Error %s when inserting row into links table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d links created ", rows)
	return nil
}

func (linksRepo *PostgresLinksRepo) StoreLinks(links []models.Link) error {
	query := "INSERT INTO links(url, crawled, domain) VALUES "
	var params []interface{}
	for i, link := range links {
		params = append(params, link.Url, false, link.Domain)

		numFields := 3 // the number of fields you are inserting
		n := i * numFields

		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1] // remove the trailing comma

	res, err := driver.PostgresDB.Client.Exec(query, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into links table", err)
		return err
	}
	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d links created simulatneously", rows)
	return nil
}

func (linksRepo *PostgresLinksRepo) IsLinkExist(href string) bool {
	var id int
	row := driver.PostgresDB.Client.QueryRow("SELECT id from links WHERE url = $1", href)
	err := row.Scan(&id)

	return err != sql.ErrNoRows
}

func (linksRepo *PostgresLinksRepo) All() []models.Link {
	return nil
}
