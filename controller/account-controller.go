package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/services"
)

type AController interface {
	AccountInsertTable(tab interface{}, ctx *gin.Context) error
	AccountSelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	//AccountSelectAll(tab interface{}) (interface{}, error)
	AccountDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error)
	AccountUpdateRow(tab interface{}, id int) (int, error)
}

type acontroller struct {
	s services.AService
}

func ANew(ser services.AService) AController {
	return &acontroller{
		s: ser,
	}
}

func (c *acontroller) AccountInsertTable(tab interface{}, ctx *gin.Context) error {

	err := ctx.ShouldBindJSON(&tab)
	log.Printf("insert")
	err = c.s.AInsertTable(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (c *acontroller) AccountSelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	if id != "" {
		idd, _ := strconv.Atoi(id)
		tab, err := c.s.ASelectID(tab, idd)
		if err != nil {
			log.Printf("error while inserting :%v", err)
		}
		return tab, err
	} else {
		tab, err := c.s.ASelectAll(tab)
		if err != nil {
			log.Printf("error while selecting :%v", err)
		}
		return tab, err
	}
}

/*func (c *acontroller) AccountSelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.ASelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}*/
func (c *acontroller) AccountDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error) {
	iid := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(iid)
	res, delerr := c.s.ADeleteRow(tab, id)

	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	return res, iid, delerr
}

func (c *acontroller) AccountUpdateRow(tab interface{}, id int) (int, error) {
	//iid := ctx.Params.ByName("id")
	//id, _ := strconv.Atoi(iid)
	res, upderr := c.s.AUpdateRow(tab, id)
	if upderr != nil {
		log.Printf("error %s", upderr)
		return 0, upderr
	}
	return res, nil
}
