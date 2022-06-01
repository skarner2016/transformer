package mysql

import (
	"fmt"
	"transformer/library/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	instantMap map[Instant]*gorm.DB
)

func NewMysql(instant Instant) (*gorm.DB, error) {
	if db, ok := instantMap[instant]; ok {
		return db, nil
	}

	dbConf := new(DBConf)
	if err := config.VipConfig.UnmarshalKey(string(instant), &dbConf); err != nil {
		return nil, err
	}

	// dsn
	dsn := getDSN(dbConf)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(dbConf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConf.ConnMaxLifetime) * time.Minute)

	instantMap = make(map[Instant]*gorm.DB)
	instantMap[instant] = db

	return db, nil
}

func getDSN(c *DBConf) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.User,
			c.Pass,
			c.Host,
			c.Port,
			c.Database)
}