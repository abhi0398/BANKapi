package entity

import "time"

//Branch store branch details
type Branch struct {
	tableName struct{}  `sql:"branch_details"`
	ID        int       `json:"id" binding:"required" sql:"id,pk"`
	Address   string    `json:"address" sql:"address,notnull"`
	FkBrankID int       `json:"fkbankid" sql:"fk_bank_id ,type:int REFERENCES bank_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	Time      time.Time `json:"time" sql:"time,notnull,default:now()"`
	account   []*Account
	loan      []*Loan
}

//insert into branch_details values(224, 'Delhi', 122 );
