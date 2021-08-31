package driver

type IDbFactory interface {
	ConnectDatabase()
}

const MONGODB = "mongodb"
const MYSQL = "mysql"

func GetDbDriverFactory(dbType string) IDbFactory {

	switch dbType {
	case MONGODB:
		return &MongoDB{}
	case MYSQL:
		return &Mysql{}
	}

	return nil
}
