package main

import (
	"log"
	"os"
	"time"

	pgdb "github.com/cipta-ageung/simas-db/proto"
	services "github.com/cipta-ageung/simas-db/services"

	_ "github.com/lib/pq"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// test runClient
func runClient(service micro.Service) {

	setup := pgdb.NewServiceConnectionService("go.micro.srv.simasdb", service.Client())

	rsp, err := setup.SetupDb(context.TODO(), &pgdb.ServiceApp{Svc: "go.micro.srv.simaslogin"})

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(rsp)
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.simasdb"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "simas_db",
		}),

		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	// Register handler
	pgdb.RegisterServiceConnectionHandler(service.Server(), new(services.ServiceConnection))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
