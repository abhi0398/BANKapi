package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/services"
)

type CController interface {
	CustomerInsertTable(tab interface{}, ctx *gin.Context) error
	CustomerSelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	//CustomerSelectAll(tab interface{}) (interface{}, error)
	CustomerDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error)
	CustomerUpdateRow(tab interface{}, id int) (int, error)
}

type ccontroller struct {
	s services.CService
}

func CNew(ser services.CService) CController {
	return &ccontroller{
		s: ser,
	}
}

func (c *ccontroller) CustomerInsertTable(tab interface{}, ctx *gin.Context) error {

	err := ctx.ShouldBindJSON(&tab)
	log.Printf("insert")
	err = c.s.CInsertTable(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (c *ccontroller) CustomerSelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	if id != "" {
		idd, _ := strconv.Atoi(id)
		tab, err := c.s.CSelectID(tab, idd)
		if err != nil {
			log.Printf("error while inserting :%v", err)
		}
		return tab, err
	} else {
		tab, err := c.s.CSelectAll(tab)
		if err != nil {
			log.Printf("error while selecting :%v", err)
		}
		return tab, err
	}
}

/*func (c *ccontroller) CustomerSelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.CSelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}*/
func (c *ccontroller) CustomerDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error) {
	iid := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(iid)
	res, delerr := c.s.CDeleteRow(tab, id)

	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	return res, iid, delerr
}

func (c *ccontroller) CustomerUpdateRow(tab interface{}, id int) (int, error) {
	//iid := ctx.Params.ByName("id")
	//id, _ := strconv.Atoi(iid)
	res, upderr := c.s.CUpdateRow(tab, id)
	if upderr != nil {
		log.Printf("error %s", upderr)
		return 0, upderr
	}
	return res, nil
}
