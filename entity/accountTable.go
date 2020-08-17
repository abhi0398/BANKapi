package entity

import "time"

//Account stores the data of the accounts
type Account struct {
	tableName  struct{}  `sql:"account_details"`
	ID         int       `json:"id" binding:"required" sql:"id,pk"`
	Balance    float64   `json:"balance" binding:"gte=0" sql:"balance,notnull"`
	Type       string    `json:"type" sql:"type,type:my_enum,notnull"`
	FkCID      int       `json:"fkcustid" sql:"fk_cust_id ,type:int REFERENCES customer_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	FkBranchID int       `json:"fkbranchid"  sql:"fk_branch_id ,type:int REFERENCES branch_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	Time       time.Time `json:"time" sql:"time,notnull,default:now()"`
	trans      []*Transaction
}

//insert into account_details values(324, 500, 'saving', 20, 222);
