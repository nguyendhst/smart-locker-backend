// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createLockerStmt, err = db.PrepareContext(ctx, createLocker); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLocker: %w", err)
	}
	if q.createLockerUserStmt, err = db.PrepareContext(ctx, createLockerUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLockerUser: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteLockerStmt, err = db.PrepareContext(ctx, deleteLocker); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteLocker: %w", err)
	}
	if q.deleteLockerUserStmt, err = db.PrepareContext(ctx, deleteLockerUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteLockerUser: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getLockerStmt, err = db.PrepareContext(ctx, getLocker); err != nil {
		return nil, fmt.Errorf("error preparing query GetLocker: %w", err)
	}
	if q.getLockerByLockerNumberStmt, err = db.PrepareContext(ctx, getLockerByLockerNumber); err != nil {
		return nil, fmt.Errorf("error preparing query GetLockerByLockerNumber: %w", err)
	}
	if q.getLockerByLockerNumberAndLocationStmt, err = db.PrepareContext(ctx, getLockerByLockerNumberAndLocation); err != nil {
		return nil, fmt.Errorf("error preparing query GetLockerByLockerNumberAndLocation: %w", err)
	}
	if q.getLockerByNfcSigStmt, err = db.PrepareContext(ctx, getLockerByNfcSig); err != nil {
		return nil, fmt.Errorf("error preparing query GetLockerByNfcSig: %w", err)
	}
	if q.getLockersOfUserStmt, err = db.PrepareContext(ctx, getLockersOfUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetLockersOfUser: %w", err)
	}
	if q.getSensorByIdStmt, err = db.PrepareContext(ctx, getSensorById); err != nil {
		return nil, fmt.Errorf("error preparing query GetSensorById: %w", err)
	}
	if q.getSensorsByTypeStmt, err = db.PrepareContext(ctx, getSensorsByType); err != nil {
		return nil, fmt.Errorf("error preparing query GetSensorsByType: %w", err)
	}
	if q.getSensorsOfLockerStmt, err = db.PrepareContext(ctx, getSensorsOfLocker); err != nil {
		return nil, fmt.Errorf("error preparing query GetSensorsOfLocker: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.updateLockerStmt, err = db.PrepareContext(ctx, updateLocker); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLocker: %w", err)
	}
	if q.updateLockerNfcSigStmt, err = db.PrepareContext(ctx, updateLockerNfcSig); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLockerNfcSig: %w", err)
	}
	if q.updateLockerStatusStmt, err = db.PrepareContext(ctx, updateLockerStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLockerStatus: %w", err)
	}
	if q.updateLockerUserStmt, err = db.PrepareContext(ctx, updateLockerUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLockerUser: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createLockerStmt != nil {
		if cerr := q.createLockerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLockerStmt: %w", cerr)
		}
	}
	if q.createLockerUserStmt != nil {
		if cerr := q.createLockerUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLockerUserStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteLockerStmt != nil {
		if cerr := q.deleteLockerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteLockerStmt: %w", cerr)
		}
	}
	if q.deleteLockerUserStmt != nil {
		if cerr := q.deleteLockerUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteLockerUserStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getLockerStmt != nil {
		if cerr := q.getLockerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLockerStmt: %w", cerr)
		}
	}
	if q.getLockerByLockerNumberStmt != nil {
		if cerr := q.getLockerByLockerNumberStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLockerByLockerNumberStmt: %w", cerr)
		}
	}
	if q.getLockerByLockerNumberAndLocationStmt != nil {
		if cerr := q.getLockerByLockerNumberAndLocationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLockerByLockerNumberAndLocationStmt: %w", cerr)
		}
	}
	if q.getLockerByNfcSigStmt != nil {
		if cerr := q.getLockerByNfcSigStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLockerByNfcSigStmt: %w", cerr)
		}
	}
	if q.getLockersOfUserStmt != nil {
		if cerr := q.getLockersOfUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLockersOfUserStmt: %w", cerr)
		}
	}
	if q.getSensorByIdStmt != nil {
		if cerr := q.getSensorByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSensorByIdStmt: %w", cerr)
		}
	}
	if q.getSensorsByTypeStmt != nil {
		if cerr := q.getSensorsByTypeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSensorsByTypeStmt: %w", cerr)
		}
	}
	if q.getSensorsOfLockerStmt != nil {
		if cerr := q.getSensorsOfLockerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSensorsOfLockerStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.updateLockerStmt != nil {
		if cerr := q.updateLockerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLockerStmt: %w", cerr)
		}
	}
	if q.updateLockerNfcSigStmt != nil {
		if cerr := q.updateLockerNfcSigStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLockerNfcSigStmt: %w", cerr)
		}
	}
	if q.updateLockerStatusStmt != nil {
		if cerr := q.updateLockerStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLockerStatusStmt: %w", cerr)
		}
	}
	if q.updateLockerUserStmt != nil {
		if cerr := q.updateLockerUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLockerUserStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                     DBTX
	tx                                     *sql.Tx
	createLockerStmt                       *sql.Stmt
	createLockerUserStmt                   *sql.Stmt
	createUserStmt                         *sql.Stmt
	deleteLockerStmt                       *sql.Stmt
	deleteLockerUserStmt                   *sql.Stmt
	deleteUserStmt                         *sql.Stmt
	getLockerStmt                          *sql.Stmt
	getLockerByLockerNumberStmt            *sql.Stmt
	getLockerByLockerNumberAndLocationStmt *sql.Stmt
	getLockerByNfcSigStmt                  *sql.Stmt
	getLockersOfUserStmt                   *sql.Stmt
	getSensorByIdStmt                      *sql.Stmt
	getSensorsByTypeStmt                   *sql.Stmt
	getSensorsOfLockerStmt                 *sql.Stmt
	getUserByEmailStmt                     *sql.Stmt
	updateLockerStmt                       *sql.Stmt
	updateLockerNfcSigStmt                 *sql.Stmt
	updateLockerStatusStmt                 *sql.Stmt
	updateLockerUserStmt                   *sql.Stmt
	updateUserStmt                         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                     tx,
		tx:                                     tx,
		createLockerStmt:                       q.createLockerStmt,
		createLockerUserStmt:                   q.createLockerUserStmt,
		createUserStmt:                         q.createUserStmt,
		deleteLockerStmt:                       q.deleteLockerStmt,
		deleteLockerUserStmt:                   q.deleteLockerUserStmt,
		deleteUserStmt:                         q.deleteUserStmt,
		getLockerStmt:                          q.getLockerStmt,
		getLockerByLockerNumberStmt:            q.getLockerByLockerNumberStmt,
		getLockerByLockerNumberAndLocationStmt: q.getLockerByLockerNumberAndLocationStmt,
		getLockerByNfcSigStmt:                  q.getLockerByNfcSigStmt,
		getLockersOfUserStmt:                   q.getLockersOfUserStmt,
		getSensorByIdStmt:                      q.getSensorByIdStmt,
		getSensorsByTypeStmt:                   q.getSensorsByTypeStmt,
		getSensorsOfLockerStmt:                 q.getSensorsOfLockerStmt,
		getUserByEmailStmt:                     q.getUserByEmailStmt,
		updateLockerStmt:                       q.updateLockerStmt,
		updateLockerNfcSigStmt:                 q.updateLockerNfcSigStmt,
		updateLockerStatusStmt:                 q.updateLockerStatusStmt,
		updateLockerUserStmt:                   q.updateLockerUserStmt,
		updateUserStmt:                         q.updateUserStmt,
	}
}
