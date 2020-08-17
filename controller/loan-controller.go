package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/services"
)

type LController interface {
	LoanInsertTable(tab interface{}, ctx *gin.Context) error
	LoanSelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	//LoanSelectAll(tab interface{}) (interface{}, error)
	LoanDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error)
	LoanUpdateRow(tab interface{}, id int) (int, error)
}

type lcontroller struct {
	s services.LService
}

func LNew(ser services.LService) LController {
	return &lcontroller{
		s: ser,
	}
}

func (c *lcontroller) LoanInsertTable(tab interface{}, ctx *gin.Context) error {

	err := ctx.ShouldBindJSON(&tab)
	log.Printf("insert")
	err = c.s.LInsertTable(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (c *lcontroller) LoanSelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	if id != "" {
		idd, _ := strconv.Atoi(id)
		tab, err := c.s.LSelectID(tab, idd)
		if err != nil {
			log.Printf("error while inserting :%v", err)
		}
		return tab, err
	} else {
		tab, err := c.s.LSelectAll(tab)
		if err != nil {
			log.Printf("error while selecting :%v", err)
		}
		return tab, err
	}
}

/*func (c *lcontroller) LoanSelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.LSelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}*/
func (c *lcontroller) LoanDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error) {
	iid := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(iid)
	res, delerr := c.s.LDeleteRow(tab, id)

	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	return res, iid, delerr
}

func (c *lcontroller) LoanUpdateRow(tab interface{}, id int) (int, error) {
	//iid := ctx.Params.ByName("id")
	//id, _ := strconv.Atoi(iid)
	res, upderr := c.s.LUpdateRow(tab, id)
	if upderr != nil {
		log.Printf("error %s", upderr)
		return 0, upderr
	}
	return res, nil
}
