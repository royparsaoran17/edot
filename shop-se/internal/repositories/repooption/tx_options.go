package repooption

import "database/sql"

type TxOptions struct {
	Tx              *sql.Tx
	NotCommitInRepo bool
}

type TxOption func(*TxOptions)

func WithTx(tx *sql.Tx) TxOption {
	return func(options *TxOptions) {
		options.Tx = tx
		options.NotCommitInRepo = true
	}
}
