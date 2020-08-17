package services

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
)

type AService interface {
	AInsertTable(tab interface{}) error
	ASelectID(tab interface{}, idt int) (interface{}, error)
	ASelectAll(tab interface{}) (interface{}, error)
	ADeleteRow(tab interface{}, id int) (int, error)
	AUpdateRow(tab interface{}, id int) (int, error)
}

type aservice struct {
	db *pg.DB
}

func ANew() AService {
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

	return &aservice{
		db: db,
	}
}

func (s *aservice) AInsertTable(tab interface{}) error {

	log.Printf("insert")
	err := s.db.Insert(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (s *aservice) ASelectID(tab interface{}, idt int) (interface{}, error) {
	err := s.db.Model(tab).Where("id =?", idt).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}

	return tab, err
}
func (s *aservice) ASelectAll(tab interface{}) (interface{}, error) {
	err := s.db.Model(tab).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}
func (s *aservice) ADeleteRow(tab interface{}, id int) (int, error) {
	res, delerr := s.db.Model(tab).Where("id=?", id).Delete()
	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	r := res.RowsAffected()
	fmt.Println(r)
	return r, delerr
}

func (s *aservice) AUpdateRow(tab interface{}, id int) (int, error) {
	res, upderr := s.db.Model(tab).Where("id=?", id).Update()
	if upderr != nil {
		log.Printf("error, reached service %s", upderr)
	}
	r := res.RowsAffected()
	fmt.Println(r)
	return r, upderr
}
