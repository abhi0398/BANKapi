package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/controller"
	"github.com/my/main/entity"
	"github.com/my/main/services"
)

var (
	s  services.Service       = services.New()
	c  controller.Controller  = controller.New(s)
	bs services.BService      = services.BNew()
	bc controller.BController = controller.BNew(bs)
	er error
)

func AdminApi(router *gin.Engine) {

	adminGroup := router.Group("admin")
	{

		//adminGroup.GET("/bank", GetBank)
		adminGroup.GET("/bank", GetBankByid)
		adminGroup.POST("/bank/", BankInsertVal)
		adminGroup.DELETE("/bank/:id", DeleteBank)
		adminGroup.PUT("/bank/:id", UpdateBank)

		//adminGroup.GET("/branch", GetBranch)
		adminGroup.GET("/branch", GetBranchByid)
		adminGroup.POST("/branch/", BranchInsertVal)
		adminGroup.DELETE("/branch/:id", DeleteBranch)
		adminGroup.PUT("/branch/:id", UpdateBranch)
	}

}

func GetBankByid(ctx *gin.Context) {
	var tab []entity.Bank
	var res interface{}
	res, er = c.BankSelectID(&tab, ctx)
	if er != nil {
		if er.Error() == "pg: no rows in result set" {
			ctx.JSON(200, gin.H{"error": "not found"})
		} else {
			ctx.JSON(200, gin.H{"error": er.Error()})
		}
	} else {
		ctx.JSON(200, res)
	}
}

func BankInsertVal(ctx *gin.Context) {
	var tab entity.Bank
	er = c.BankInsertTable(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": er.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteBank(ctx *gin.Context) {
	var tab entity.Bank
	res, id, err := c.BankDeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateBank(ctx *gin.Context) {

	//updateval := ctx.Query("updval")
	//updateval := ctx.Params.ByName("updval")
	tab := entity.Bank{
	//ID:   id,
	//Name: updateval,
	}
	er = ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, err := c.BankUpdateRow(&tab, tab.ID)

	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//-----------------------
func GetBranchByid(ctx *gin.Context) {
	var tab []entity.Branch
	var res interface{}
	res, er = bc.BranchSelectID(&tab, ctx)
	if er != nil {
		if er.Error() == "pg: no rows in result set" {
			ctx.JSON(200, gin.H{"error": "not found"})
		} else {
			ctx.JSON(200, gin.H{"error": er.Error()})
		}
	} else {
		ctx.JSON(200, res)
	}
}

func BranchInsertVal(ctx *gin.Context) {
	var tab entity.Branch
	er = bc.BranchInsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": er.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteBranch(ctx *gin.Context) {
	var tab entity.Branch
	res, id, err := bc.BranchDeleteRow(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateBranch(ctx *gin.Context) {
	//updateval := ctx.Query("updval")
	//updateval := ctx.Params.ByName("updval")
	tab := entity.Branch{
	//ID:   id,
	//Address: updateval,
	}
	er = ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, err := bc.BranchUpdateRow(&tab, tab.ID)

	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}
