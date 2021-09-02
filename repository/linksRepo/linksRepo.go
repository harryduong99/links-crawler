package linksRepo

import (
	"links-crawler/config"
	"links-crawler/models"
)

type LinksRepo interface {
	StoreLink(models.Link) error
	StoreLinks([]models.Link) error
	IsLinkExist(string) bool
	All() []models.Link
}

func GetLinksRepo(dbType string) LinksRepo {
	switch dbType {
	case config.MONGODB:
		return &MongoLinksRepo{}
	case config.MYSQL:
		return &MysqlLinksRepo{}
	case config.POSTGRES:
		return &PostgresLinksRepo{}
	case config.XLSX:
		return &XlsxLinksRepo{}
	}

	return nil
}
