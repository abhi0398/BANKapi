package services

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
)

type BService interface {
	BInsertTable(tab interface{}) error
	BSelectID(tab interface{}, idt int) (interface{}, error)
	BSelectAll(tab interface{}) (interface{}, error)
	BDeleteRow(tab interface{}, id int) (int, error)
	BUpdateRow(tab interface{}, id int) (int, error)
}

type bservice struct {
	db *pg.DB
}

func BNew() BService {
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

	return &bservice{
		db: db,
	}
}

func (s *bservice) BInsertTable(tab interface{}) error {

	log.Printf("insert")
	err := s.db.Insert(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (s *bservice) BSelectID(tab interface{}, idt int) (interface{}, error) {
	err := s.db.Model(tab).Where("id =?", idt).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}

	return tab, err
}
func (s *bservice) BSelectAll(tab interface{}) (interface{}, error) {
	err := s.db.Model(tab).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}
func (s *bservice) BDeleteRow(tab interface{}, id int) (int, error) {
	res, delerr := s.db.Model(tab).Where("id=?", id).Delete()
	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	r := res.RowsAffected()
	fmt.Println(r)
	return r, delerr
}

func (s *bservice) BUpdateRow(tab interface{}, id int) (int, error) {
	res, upderr := s.db.Model(tab).Where("id=?", id).Update()
	if upderr != nil {
		log.Printf("error, reached service %s", upderr)
	}
	r := res.RowsAffected()
	fmt.Println(r)
	return r, upderr
}
