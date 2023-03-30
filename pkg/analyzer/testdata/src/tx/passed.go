package tx

import (
	"database/sql"
	"log"
)

func passedAndClosed() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	closedPassed(tx)
}

func closedPassed(tx *sql.Tx) {
	tx.Rollback()
}

func passedAndNotClosed() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	dontClosedPassed(tx)
}

func dontClosedPassed(*sql.Tx) {

}
