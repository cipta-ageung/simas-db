package services

import (
	"context"
	"log"

	pgdb "github.com/cipta-ageung/simas-db/proto"
	"github.com/micro/go-micro/config"
)

// ServiceConnection : struct
type ServiceConnection struct{}

// SetupDb : method
func (g *ServiceConnection) SetupDb(ctx context.Context, req *pgdb.ServiceApp, rsp *pgdb.ServiceDb) error {

	// clear or empty first for new request
	rsp = &pgdb.ServiceDb{}

	// load data from file configuration
	config.LoadFile("./config/db.json")

	// set data configuration into type struct ServiceDb
	config.Get("services", req.Svc).Scan(&rsp)

	// check data struct or response
	if rsp != nil {
		log.Print("setup db for services : " + req.Svc)

		// set data response string connection
		rsp.ConnectionDb = "postgres://" + rsp.GetUser() + ":" + rsp.GetPassword() + "@" + rsp.GetHost() +
			":" + rsp.GetPort() + "/" + rsp.GetDbname() + "?sslmode=disable"
	}

	return nil
}
