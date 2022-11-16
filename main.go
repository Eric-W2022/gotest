package main

import (
	"fmt"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/router"
)

func main() {

	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)

	//log.Fatal(http.ListenAndServe(":80", nil))

	//r.GET("/api/count", service.CounterHandler)

	r := router.Router()
	startErr := r.Run(":80")
	if startErr != nil {
		return
	}
}
