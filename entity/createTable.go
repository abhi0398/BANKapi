package entity

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

/*func CreateTable(db *pg.DB, tab interface{}) {
	opts := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	}
	err := db.CreateTable(tab, opts)
	if err != nil {
		log.Printf("error while creating table:%v", err)
		os.Exit(100)
	}
}*/
func CreateTable(dbb *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true, //don't create new table if already exists, predefind method
	}
	_, err := dbb.Exec("	CREATE TYPE my_enum AS ENUM ('current', 'saving')")
	if err != nil {
		log.Printf("error creating enum %s", err)
	}
	models := []interface{}{
		(*Bank)(nil),
		(*Branch)(nil),
		(*Customer)(nil),
		(*Account)(nil),
		(*Loan)(nil),
		(*Transaction)(nil),
	}

	for _, model := range models {
		err := dbb.CreateTable(model, opts)
		if err != nil {
			log.Printf("error creating tables %s", err)
			return err
		}
	}
	log.Printf("tables created")
	return nil
}
