package tx

import (
	"database/sql"
	"log"
)

func returnTx() (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return tx, nil
}

func returnTxShort() (*sql.Tx, error) {
	return nil, nil
}
