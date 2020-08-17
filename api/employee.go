package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/my/main/controller"
	"github.com/my/main/entity"
	"github.com/my/main/services"
)

var (
	cs  services.CService      = services.CNew()
	cc  controller.CController = controller.CNew(cs)
	as  services.AService      = services.ANew()
	ac  controller.AController = controller.ANew(as)
	ls  services.LService      = services.LNew()
	lc  controller.LController = controller.LNew(ls)
	ts  services.TService      = services.TNew()
	tc  controller.TController = controller.TNew(ts)
	err error
	db  *pg.DB
)

func EmployeeApi(router *gin.Engine) {

	employeeGroup := router.Group("employee")
	{
		//employeeGroup.GET("/account", GetAccount)
		employeeGroup.GET("/account", GetAccountByid)
		employeeGroup.POST("/account/", AccountInsertVal)
		employeeGroup.DELETE("/account/:id", DeleteAccount)
		employeeGroup.PUT("/account/:id", UpdateAccount)

		//employeeGroup.GET("/customer", GetCustomer)
		employeeGroup.GET("/customer", GetCustomerByid)
		employeeGroup.POST("/customer/", CustomerInsertVal)
		employeeGroup.DELETE("/customer/:id", DeleteCustomer)
		employeeGroup.PUT("/customer/:id", UpdateCustomer)

		//employeeGroup.GET("/loan", GetLoan)
		employeeGroup.GET("/loan", GetLoanByid)
		employeeGroup.POST("/loan/", LoanInsertVal)
		employeeGroup.DELETE("/loan/:id", DeleteLoan)
		employeeGroup.PUT("/loan/:id", UpdateLoan)

		//employeeGroup.GET("/transaction", GetTransaction)
		employeeGroup.GET("/transaction", GetTransactionByid)
		employeeGroup.POST("/transaction", Trans)
	}
}
func GetAccountByid(ctx *gin.Context) {
	var tab []entity.Account
	var res interface{}
	res, err = ac.AccountSelectID(&tab, ctx)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			ctx.JSON(200, gin.H{"error": "not found"})
		} else {
			ctx.JSON(200, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(200, res)
	}
}

func AccountInsertVal(ctx *gin.Context) {
	var tab entity.Account
	err = ac.AccountInsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteAccount(ctx *gin.Context) {
	var tab entity.Account
	res, id, er := ac.AccountDeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateAccount(ctx *gin.Context) {
	tab := entity.Account{}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := ac.AccountUpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//-------------------
func GetCustomerByid(ctx *gin.Context) {
	var tab []entity.Customer
	var res interface{}
	res, err = cc.CustomerSelectID(&tab, ctx)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			ctx.JSON(200, gin.H{"error": "not found"})
		} else {
			ctx.JSON(200, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(200, res)
	}
}

func CustomerInsertVal(ctx *gin.Context) {
	var tab entity.Customer
	err = cc.CustomerInsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteCustomer(ctx *gin.Context) {
	var tab entity.Customer
	res, id, er := cc.CustomerDeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateCustomer(ctx *gin.Context) {
	tab := entity.Customer{}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := cc.CustomerUpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//---------------------
func GetLoanByid(ctx *gin.Context) {
	var tab []entity.Loan
	var res interface{}
	res, err = lc.LoanSelectID(&tab, ctx)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			ctx.JSON(200, gin.H{"error": "not found"})
		} else {
			ctx.JSON(200, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(200, res)
	}
}

func LoanInsertVal(ctx *gin.Context) {
	var tab entity.Loan
	err = lc.LoanInsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteLoan(ctx *gin.Context) {
	var tab entity.Loan
	res, id, er := lc.LoanDeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateLoan(ctx *gin.Context) {
	tab := entity.Loan{}
	err = ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := lc.LoanUpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//-------------------------

func GetTransactionByid(ctx *gin.Context) {
	var tab []entity.Transaction
	var res interface{}
	res, err = tc.TransactionSelectID(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}
func Trans(ctx *gin.Context) {

	er := tc.Transactions(ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"value": "transacted"})
}
