package db

import (
	"fmt"
	"log"

	"github.com/Shota-Kurahashi/ent-with-echo/ent"
	"github.com/Shota-Kurahashi/ent-with-echo/src/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *ent.Client {
	client, err := ent.Open("mysql", config.GetMySqlUrl())

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return client
}

func Close(client *ent.Client) {
	err := client.Close()

	if err != nil {
		log.Fatalf("failed closing connection to mysql: %v", err)
	}

	fmt.Println("Connection to mysql closed.")
}
