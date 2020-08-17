package entity

import (
	"time"
)

//Transaction stores the transaction details
type Transaction struct {
	tableName     struct{}  `sql:"transaction_details"`
	ID            int       `json:"id" binding:"required" sql:"id, pk"`
	FromAccountNo int       `json:"fromaccountno" sql:"fk_from_account_no ,type:int REFERENCES account_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	ToAccountNo   int       `json:"toaccount" sql:"fk_to_account_no ,type:int REFERENCES account_details(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL"`
	Transtime     time.Time `json:"time" sql:"time,notnull,default:now()"`
	Amount        float64   `json:"amount" sql:"amount,notnull"`
}
