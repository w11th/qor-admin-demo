package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
)

type Product struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "demo.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	adm := admin.New(&admin.AdminConfig{
		SiteName: "DEMO",
		DB:       db,
	})
	adm.AddResource(&Product{})

	mux := http.NewServeMux()
	adm.MountTo("admin", mux)
	fmt.Println("Server listened. Open: http://127.0.0.1:9000")
	err = http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
