package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my/main/api"
)

func main() {

	/*opts := &pg.Options{
		User:     "postgres",
		Password: "Pcmwithsc",
		Addr:     "localhost:5432",
		Database: "project",
	}
	db := pg.Connect(opts)
	if db == nil {
		log.Printf("failed  connection")
		os.Exit(100)
	}
	log.Printf("connected")
	entity.CreateTable(db)
	db.Close()

	*/

	router := gin.Default()
	api.AdminApi(router)
	api.EmployeeApi(router)

	router.Run(":8080")

}

/*func GetBank(ctx *gin.Context) {
	var tab []entity.Bank
	var res interface{}
	res, err = c.BankSelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}*/

//--------------------

/*
{
    "title" : "hello1",
    "description" : "testing1" ,
    "url" : "https://www.youtube.com/watch?v=qR0WnWL2o1Q&t=32ss" ,
    "author" : {
                    "name" : "Ayushi2",
                    "age" : 23,
                    "email" : "Ayushi2@gmail.com"
                }
}*/
