package main

import (
	"log"

	"github.com/Davematteer/SocialApp/internals/db"
	"github.com/Davematteer/SocialApp/internals/env"
	"github.com/Davematteer/SocialApp/internals/env/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		db: dbConfig{
			addr:         env.GetString("ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disabled"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDEL_TIME", "15min"),
		},
	}

	db, err := db.New(cfg.db.addr,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	store := store.NewStorage(db)
	log.Println("database connection pool established")

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
