package prisma

import (
	"context"
	"log"

	"github.com/branogarbo/sunswap_backend/prisma/db"
)

var (
	Client *db.PrismaClient
	Ctx    context.Context = context.Background()
)

func Connect() {
	// init db
	Client = db.NewClient()
	if err := Client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}
}
