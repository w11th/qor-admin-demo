package main

import (
	"net/http"

	"github.com/qor/admin"
)

func main() {
	adm := admin.New(&admin.Config{
		Name: "DEMO",
	})

	mux := http.NewServeMux()
	adm.MountTo("admin", mux)

	_ = http.ListenAndServe(":9000", mux)
}
