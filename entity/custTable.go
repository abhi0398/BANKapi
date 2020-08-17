package entity

import "time"

//Customer stores the data of the customer
type Customer struct {
	tableName struct{}  `sql:"customer_details"`
	ID        int       `json:"id" binding:"required" sql:"id,pk"`
	Name      string    `json:"name" sql:"name,notnull"`
	Address   string    `json:"address" sql:"address,notnull"`
	City      string    `json:"city" sql:"city,notnull"`
	PhoneNo   string    `json:"phoneno" sql:"phoneno,unique,notnull"`
	Email     string    `json:"email" binding:"required,email" sql:"email,notnull,unique"`
	Time      time.Time `json:"time" sql:"time,notnull,default:now()"`
	account   []*Account
	loan      []*Loan
}

//insert into customer_details values(20,'Aarna','B5', 'Delhi','8750870026 ','ar@gmail.com', '2016-06-22 5:10:25-07' );
//insert into customer_details values(10,'Ayushi','l5', 'Noida','8750870025 ','a@gmail.com');
