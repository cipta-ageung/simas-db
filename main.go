package main

import (
	"log"
	"time"

	pgdb "github.com/cipta-ageung/simas-db/proto"
	services "github.com/cipta-ageung/simas-db/services"

	_ "github.com/lib/pq"

	"github.com/micro/go-micro"
)

func main() {

	// Initiate Service
	service := micro.NewService(
		micro.Name("go.micro.srv.simasdb"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "simas_db",
		}),
	)
	service.Init()

	// Register handler
	pgdb.RegisterServiceConnectionHandler(service.Server(), new(services.ServiceConnection))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
