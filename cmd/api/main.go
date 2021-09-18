package main

import (
	"context"
	"log"
)

func main(){
	if err := run(); err != nil{
		log.Fatal(err)
	}
}

func run() error{
	cxt := context.Background()

	//config
	//cfg := config.Get()

	//Init repository store (with mysql/postgresql inside)
	store, err := store.New(cxt)
	if err != nil{
		return erros.Wrap(err, "store.New failed")
	}


	return nil
}
