package tx

import (
	"log"
)

func missingClose() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_ = tx
}
