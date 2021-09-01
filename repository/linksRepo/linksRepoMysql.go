package linksRepo

import (
	"context"
	"database/sql"
	"log"
	"song-chord-crawler/driver"
	"song-chord-crawler/models"
	"strings"
	"time"
)

type MysqlLinksRepo struct{}

func (linksRepo *MysqlLinksRepo) StoreLink(link models.Link) error {
	query := "INSERT INTO links(url, crawled, domain) VALUES (?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := driver.MysqlDB.Client.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, link.Url, false, link.Domain)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func (linksRepo *MysqlLinksRepo) StoreLinks(links []models.Link) error {
	query := "INSERT INTO links(url, crawled, domain) VALUES "
	var inserts []string
	var params []interface{}
	for _, v := range links {
		inserts = append(inserts, "(?, ?, ?)")
		params = append(params, v.Url, false, v.Domain)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := driver.MysqlDB.Client.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created simulatneously", rows)
	return nil
}

func (linksRepo *MysqlLinksRepo) IsLinkExist(href string) bool {
	var id int
	row := driver.MysqlDB.Client.QueryRow("SELECT id from links WHERE url = ?", href)

	err := row.Scan(&id)

	return err != sql.ErrNoRows
}

func (linksRepo *MysqlLinksRepo) All() []models.Link {
	return nil
}
