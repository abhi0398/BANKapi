package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/services"
)

type Controller interface {
	BankInsertTable(tab interface{}, ctx *gin.Context) error
	BankSelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	//BankSelectAll(tab interface{}) (interface{}, error)
	BankDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error)
	BankUpdateRow(tab interface{}, id int) (int, error)
}

type controller struct {
	s services.Service
}

func New(ser services.Service) Controller {
	return &controller{
		s: ser,
	}
}

func (c *controller) BankInsertTable(tab interface{}, ctx *gin.Context) error {

	err := ctx.ShouldBindJSON(&tab)
	log.Printf("insert")
	err = c.s.InsertTable(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (c *controller) BankSelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	//id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	id := ctx.Query("id")
	if id != "" {
		idd, _ := strconv.Atoi(id)
		tab, err := c.s.SelectID(tab, idd)
		if err != nil {
			log.Printf("error while inserting :%v", err)
		}
		return tab, err
	} else {
		tab, err := c.s.SelectAll(tab)
		if err != nil {
			log.Printf("error while selecting :%v", err)
		}
		return tab, err
	}

}

/*func (c *controller) BankSelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.SelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}*/
func (c *controller) BankDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error) {
	iid := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(iid)
	res, delerr := c.s.DeleteRow(tab, id)

	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	return res, iid, delerr
}

func (c *controller) BankUpdateRow(tab interface{}, id int) (int, error) {
	//iid := ctx.Params.ByName("id")
	//id, _ := strconv.Atoi(iid)
	res, upderr := c.s.UpdateRow(tab, id)
	if upderr != nil {
		log.Printf("error %s", upderr)
		return 0, upderr
	}
	return res, nil
}
