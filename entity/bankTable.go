package entity

import "time"

//Bank stores bank details
type Bank struct {
	tableName struct{}  `sql:"bank_details"`
	ID        int       `json:"id" binding:"required" sql:"id,pk"`
	Name      string    `json:"name"  sql:"name,notnull,unique"`
	Time      time.Time `json:"time"  sql:"time,notnull,default:now()"`
	branch    []*Branch
}

//insert into bank_details values(124, 'SBI', '2018-06-22 5:10:25-07' );
//update http://localhost:8080/admin/Bank/125?updval=icici
