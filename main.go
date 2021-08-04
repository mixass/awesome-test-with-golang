package main

import (	
	"github.com/labstack/echo/v4"
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	webServer()
}

func webServer()  {
	e := echo.New()
	e.GET("/license", func(c echo.Context) error {
		mySql()
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func mySql(){
	fmt.Println("Go MySQL Tutorial")


    db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/mixas")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    insert, err := db.Query("INSERT INTO test(test1) VALUES ('TEST' )")

    if err != nil {
        panic(err.Error())
    }
    defer insert.Close()
}