package main

import (
	"database/sql"
	"fmt"

	"github.com/dlintw/goconf"
	_ "github.com/mattn/go-oci8"
)

var con *sql.DB

func closeCon() {
	if con != nil {
		con.Close()
	}
}

func getConnString(prop *goconf.ConfigFile) string {

	username, _ := prop.GetString("db", "user")
	pass, _ := prop.GetString("db", "pass")
	host, _ := prop.GetString("db", "host")
	port, _ := prop.GetString("db", "port")
	service, _ := prop.GetString("db", "service-name")

	return fmt.Sprintf("%s/%s@%s:%s/%s", username, pass, host, port, service)

}

func openConn(connString string) {

	fmt.Println("getting db connection")

	if nil != props {
		db, err := sql.Open("oci8", connString)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if err = db.Ping(); err != nil {
			fmt.Printf("Error connecting to the database: %s\n", err)
			panic(err)
		}
		con = db
	}

}

func getDB() (*sql.DB, bool) {

	if nil != props {
		return con, true
	}

	fmt.Println("props arent picking up")
	return nil, false

}
