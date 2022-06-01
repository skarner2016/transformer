package mysql

import (
	"fmt"
	"transformer/library/config"
	"testing"
)



func TestNewMysql(t *testing.T) {
	config.InitConfig()

	db, err := NewMysql(InstantDefault)
	if err != nil {
		panic(err)
	}

	db2, err := NewMysql(InstantDefault)

	fmt.Println(fmt.Sprintf("%p", db))
	fmt.Println(fmt.Sprintf("%p", db2))

	//user := new(User)

	//db.First(&user)


	//fmt.Println(user)
}