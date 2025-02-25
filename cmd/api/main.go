package main

import (	
	"log"

	"github.com/nettokrt/golang_api_study/internal/env"
	"github.com/nettokrt/golang_api_study/internal/store"
)


func main()  {

	cfg := config{
		addr: env.GetString("ADDR",":8080"),
	}
	store := store.NewStorage(nil)
	
	app := &application{
		config: cfg,
		store: store,
	}




	mux := app.mount()
	log.Fatal(app.run(mux))
}
