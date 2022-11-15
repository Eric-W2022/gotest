package main

import "github.com/gin-gonic/gin"

func main() {
	//err := os.Setenv("MYSQL_USERNAME", "root")
	//if err != nil {
	//	return
	//}
	//err1 := os.Setenv("MYSQL_PASSWORD", "STC89c51")
	//if err1 != nil {
	//	return
	//}
	//err2 := os.Setenv("MYSQL_ADDRESS", "sh-cynosdbmysql-grp-6a8ntbek.sql.tencentcdb.com:22515")
	//if err2 != nil {
	//	return
	//}
	//err3 := os.Setenv("MYSQL_DATABASE", "golang_demo")
	//if err3 != nil {
	//	return
	//}
	//
	//if err := db.Init(); err != nil {
	//	panic(fmt.Sprintf("mysql init failed with %+v", err))
	//}
	//
	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)
	//
	//log.Fatal(http.ListenAndServe(":80", nil))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
