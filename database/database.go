package database

import (
	"context"
	"entgo.io/ent/dialect"
	"epitime/ent/user"

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

func (dba Database) UpdateUserYear(ctx context.Context, email string, year int) error {
	return dba.Client.User.Update().Where(user.Email(email)).SetYear(year).Exec(ctx)
}

func (dba Database) UpdateUserHideModules(ctx context.Context, email string, hideModules string) error {
	return dba.Client.User.Update().Where(user.Email(email)).SetHideModules(hideModules).Exec(ctx)
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func NewEntDatabase() Database {
	var err error
	var dba = Database{}
	dba.Client = Open("postgresql://api:password@127.0.0.1:5432/epitime")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := dba.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	fmt.Println("successful !")
	return dba
}
