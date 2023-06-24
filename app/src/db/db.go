package db

import (
	"fmt"
	"log"

	"github.com/Shota-Kurahashi/ent-with-echo/ent"
	"github.com/Shota-Kurahashi/ent-with-echo/src/config"
)

func Connect() *ent.Client {
	client, err := ent.Open("mysql", config.GetMySQLUrl())

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	return client
}

func Close(client *ent.Client) {
	err := client.Close()

	if err != nil {
		log.Fatalf("failed closing connection to mysql: %v", err)
	}

	fmt.Println("Connection to mysql closed.")
}
