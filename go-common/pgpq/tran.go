package pgpq

import "database/sql"

func CloseTran(tx *sql.Tx, e error) error {
	if e != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
