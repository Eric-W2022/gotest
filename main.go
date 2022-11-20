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

	r := router.Router()
	startErr := r.Run(":80")
	if startErr != nil {
		return
	}
}
