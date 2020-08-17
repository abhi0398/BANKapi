package services

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/my/main/entity"
)

type TService interface {
	TInsertTable(tab interface{}) error
	TSelectID(tab interface{}, idt int) (interface{}, error)
	TSelectAll(tab interface{}) (interface{}, error)
	Transactions(tab *entity.Transaction) error
}

type tservice struct {
	db *pg.DB
}

func TNew() TService {
	opts := &pg.Options{
		User:     "postgres",
		Password: "abhi123",
		Addr:     "localhost:5432",
		Database: "project",
	}
	db := pg.Connect(opts)

	if db == nil {
		log.Printf("failed  connection")
		os.Exit(100)
	}
	log.Printf("connected for service")

	return &tservice{
		db: db,
	}
}

func (s *tservice) TInsertTable(tab interface{}) error {

	log.Printf("insert")
	err := s.db.Insert(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (s *tservice) TSelectID(tab interface{}, idt int) (interface{}, error) {
	err := s.db.Model(tab).Where("id =?", idt).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}

	return tab, err
}
func (s *tservice) TSelectAll(tab interface{}) (interface{}, error) {
	err := s.db.Model(tab).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}

func (s *tservice) Transactions(tab *entity.Transaction) error {

	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		var balance float64
		err := tx.Model((*entity.Account)(nil)).
			Column("balance").
			Where("id = ?", tab.FromAccountNo).
			Select(&balance)
		log.Println(balance)
		if err != nil {
			log.Printf("couldnt fetch and error %s", err)
			return err
		}
		if balance < tab.Amount {
			err = fmt.Errorf("amount is greater than the balance:%v of account:%v", balance, tab.FromAccountNo)
			return err
		}
		//log.Println(balance)
		balance = balance - tab.Amount
		fmt.Println(balance, tab.Amount)
		_, err = tx.Model((*entity.Account)(nil)).Set("balance = ?", balance).Where("id= ?", tab.FromAccountNo).Update()
		if err != nil {
			return err
		}
		err = tx.Model((*entity.Account)(nil)).Column("balance").Where("id = ?", tab.ToAccountNo).Select(&balance)
		balance = balance + tab.Amount
		_, err = tx.Model((*entity.Account)(nil)).Set("balance = ?", balance).Where("id= ?", tab.ToAccountNo).Update()
		return err
	})
	//fmt.Println(tab)
	if err == nil {
		s.TInsertTable(tab)
	}
	return err

}
