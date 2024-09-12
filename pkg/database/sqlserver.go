package database

import (
	"context"
	"fmt"
	"foxomni/pkg/config"
	"foxomni/pkg/errs"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type SQL struct {
	DB *sqlx.DB
}

func NewSQL(conf config.SQLServerConfig) (*SQL, error) {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		conf.Server,
		conf.Username,
		conf.Password,
		conf.Port,
		conf.Database,
	)

	db, err := sqlx.Connect("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * time.Duration(conf.ConnMaxLifeTime))

	return &SQL{
		DB: db,
	}, nil
}

func (sql *SQL) InTx(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := sql.DB.Beginx()
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	if err := fn(tx); err != nil {
		if err1 := tx.Rollback(); err1 != nil {
			return errs.New(errs.Internal, err1)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}
