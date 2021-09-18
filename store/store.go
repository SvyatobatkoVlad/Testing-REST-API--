package store

import (
	"context"
	"github.com/mattermost/mattermost-server/config"
	"log"
)

//Store contains all repositories
type Store struct{
	Pg *pg.DB
}


//Creates new store
func New(ctx context.Context) (*Store, error){
	cfg := config.Get()

	//connect to Postgres
	pgDB, err := pg.Dial()
	if err != nil{
		return nil, erros.Wrap(err, "pgdb.Dial failed")
	}

	//Run Postgres migrations
	if pgDB != nil{
		log.Println("Running PostgreSQL migration... ")
		if err := runPgMigrations(); err != nil{
			return nil, erros.Wrap(err, "runPgMigrations failed")
			}


	}

}

