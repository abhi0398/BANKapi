package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/services"
)

type BController interface {
	BranchInsertTable(tab interface{}, ctx *gin.Context) error
	BranchSelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	//BranchSelectAll(tab interface{}) (interface{}, error)
	BranchDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error)
	BranchUpdateRow(tab interface{}, id int) (int, error)
}

type bcontroller struct {
	s services.BService
}

func BNew(ser services.BService) BController {
	return &bcontroller{
		s: ser,
	}
}

func (c *bcontroller) BranchInsertTable(tab interface{}, ctx *gin.Context) error {

	err := ctx.ShouldBindJSON(&tab)
	log.Printf("insert")
	err = c.s.BInsertTable(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (c *bcontroller) BranchSelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	if id != "" {
		idd, _ := strconv.Atoi(id)
		tab, err := c.s.BSelectID(tab, idd)
		if err != nil {
			log.Printf("error while inserting :%v", err)
		}
		return tab, err
	} else {
		tab, err := c.s.BSelectAll(tab)
		if err != nil {
			log.Printf("error while selecting :%v", err)
		}
		return tab, err
	}
}

/*func (c *bcontroller) BranchSelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.BSelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}*/
func (c *bcontroller) BranchDeleteRow(tab interface{}, ctx *gin.Context) (int, string, error) {
	iid := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(iid)
	res, delerr := c.s.BDeleteRow(tab, id)

	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	return res, iid, delerr
}

func (c *bcontroller) BranchUpdateRow(tab interface{}, id int) (int, error) {
	//iid := ctx.Params.ByName("id")
	//id, _ := strconv.Atoi(iid)
	res, upderr := c.s.BUpdateRow(tab, id)
	if upderr != nil {
		log.Printf("error %s", upderr)
		return 0, upderr
	}
	return res, nil
}
