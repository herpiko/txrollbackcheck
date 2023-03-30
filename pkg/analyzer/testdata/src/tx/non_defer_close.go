package tx

import (
	"log"
)

func nonDeferClose() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tx.Rollback()
}
