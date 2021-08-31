package driver

import "song-chord-crawler/config"

type IDbFactory interface {
	ConnectDatabase()
}

func GetDbDriverFactory(dbType string) IDbFactory {

	switch dbType {
	case config.MONGODB:
		return &MongoDB{}
	case config.MYSQL:
		return &Mysql{}
	}

	return nil
}
