package services

import (
	"context"
	"database/sql"
	"log"

	pgdb "github.com/cipta-ageung/simas-db/proto"
	"github.com/micro/go-micro/config"
)

// ServiceConnection : struct
type ServiceConnection struct{}

// SetupDb : method
func (g *ServiceConnection) SetupDb(ctx context.Context, req *pgdb.ServiceApp, rsp *pgdb.ServiceDb) error {

	config.LoadFile("./config/db.json")
	config.Get("services", req.Svc).Scan(&rsp)

	if rsp != nil {

		log.Print("setup db for services : " + req.Svc)
		connectionDb := "postgres://" + rsp.GetUser() + ":" + rsp.GetPassword() + "@" + rsp.GetHost() +
			":" + rsp.GetPort() + "/" + rsp.GetDbname() + "?sslmode=disable"

		db, err := sql.Open("postgres", connectionDb)
		if err != nil {
			log.Fatal(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		} else {
			log.Print(db)
		}
	}

	return nil
}
