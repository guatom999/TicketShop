package main

import (
	"context"

	"github.com/guatom999/TicketShop/config"
	"github.com/guatom999/TicketShop/database"
	"github.com/guatom999/TicketShop/modules/users"
	"github.com/guatom999/TicketShop/server"
)

func main() {

	ctx := context.Background()

	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg).GetDb()

	db.AutoMigrate(&users.Users{})

	server.NewEchoServer(db, &cfg).Start(ctx)

}
