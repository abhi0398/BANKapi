package entity

import "time"

//Loan stores loan details of customer
type Loan struct {
	tableName  struct{}  `sql:"loan_details"`
	ID         int       `json:"id" binding:"required" sql:"id, pk"`
	FkCID      int       `json:"fkcustid" sql:"fk_cust_id ,type:int REFERENCES customer_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	FkBranchID int       `json:"fkbranchid" sql:"fk_branch_id ,type:int REFERENCES branch_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	Amount     float64   `json:"amount" sql:"amount"`
	Type       string    `json:"type" sql:"type"`
	Time       time.Time `json:"time" sql:"time,notnull,default:now()"`
}
