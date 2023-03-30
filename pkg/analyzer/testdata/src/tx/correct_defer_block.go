package tx

import (
	"log"
)

func correctDeferBlock() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Print("problem closing rows")
		}
	}()
}
