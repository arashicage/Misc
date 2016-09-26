package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type parameters struct {
	DSN     map[string]map[string]string
	Mapping map[string]map[string]string
}

func New() *parameters {
	return new(parameters)
}

func (p *parameters) GetDSN() {
	dsn := make(map[string]map[string]string)

	db, err := sql.Open("sqlite3", "./conf.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, dsn_swjg, dsn_raw, dsn_raw2, dsn_raw3, valid, priority from DZDZ_ETL_DSN where valid='Y' order by priority asc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, dsn_swjg, dsn_raw, dsn_raw2, dsn_raw3, valid, priority string
		rows.Scan(&id, &dsn_swjg, &dsn_raw, &dsn_raw2, &dsn_raw3, &valid, &priority)
		rec := make(map[string]string)
		rec["id"] = id
		rec["dsn_swjg"] = dsn_swjg
		rec["dsn_raw"] = dsn_raw
		rec["dsn_raw2"] = dsn_raw2
		rec["dsn_raw3"] = dsn_raw3
		rec["valid"] = valid
		rec["priority"] = priority
		dsn[dsn_swjg] = rec
	}
	rows.Close()
	p.DSN = dsn

}

func (p *parameters) GetMapping() {
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world")
}

func (p *parameters) String() {
	for k, v := range p.DSN {
		fmt.Println(k, v)
	}

	for k, v := range p.Mapping {
		fmt.Println(k, v)
	}
}
