package mysql

import "gorm.io/gorm"

const (
	InstantDefault Instant = "mysql.default"
	InstantWrite   Instant = "mysql.read"
)

type Instant string

type DBConf struct {
	Host            string
	Port            int64
	User            string
	Pass            string
	Database        string
	Charset         string
	Collation       string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type User struct {
	gorm.Model
	Name string
}