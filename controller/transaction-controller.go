package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/entity"
	"github.com/my/main/services"
)

type TController interface {
	TransactionSelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	//TransactionSelectAll(tab interface{}) (interface{}, error)
	Transactions(cts *gin.Context) error
}

type tcontroller struct {
	s services.TService
}

func TNew(ser services.TService) TController {
	return &tcontroller{
		s: ser,
	}
}

func (c *tcontroller) TransactionSelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	if id != "" {
		idd, _ := strconv.Atoi(id)
		tab, err := c.s.TSelectID(tab, idd)
		if err != nil {
			log.Printf("error while inserting :%v", err)
		}
		return tab, err
	} else {
		tab, err := c.s.TSelectAll(tab)
		if err != nil {
			log.Printf("error while selecting :%v", err)
		}
		return tab, err
	}
}

/*func (c *tcontroller) TransactionSelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.TSelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}*/

func (c *tcontroller) Transactions(ctx *gin.Context) error {

	tab := &entity.Transaction{}
	err := ctx.ShouldBindJSON(&tab)
	if err != nil {
		log.Printf("error in json:%v", err)
	}
	log.Println(tab)
	errr := c.s.Transactions(tab)
	if errr != nil {
		log.Printf("error in transactions:%v", errr)
	}
	return errr

}
