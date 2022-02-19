package database

import (
	"context"
	"entgo.io/ent/dialect"

	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"epitime/ent"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

type Database struct {
	Client *ent.Client
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func NewEntDatabase(dba Database) *ent.Client {
	var err error
	dba.Client = Open("postgresql://root:password@127.0.0.1:5432/my_database")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := dba.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	fmt.Println("successful !")
	return dba.Client
}
