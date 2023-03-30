package tx

import (
	"log"
)

func correctDefer() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
}
