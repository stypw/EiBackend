package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"main/common/cf"
	"main/common/orm"
	"main/common/tl"

	"main/common/rt"

	_ "main/auth"
	_ "main/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cf := cf.GetConfig()
	aesKey := cf.GetProperty("aes_key").GetString()
	if aesKey != "" {
		tl.SetAesKey(aesKey)
	}
	mysql := cf.GetProperty("mysql").GetString()
	if mysql == "" {
		fmt.Println("config error")
		return
	}

	db, err := sql.Open("mysql", mysql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	orm.SetDb(db)

	rt.Start()

	listen := cf.GetProperty("listen").GetString()
	if listen == "" {
		fmt.Println("config error")
		return
	}
	http.ListenAndServe(listen, nil)
}
