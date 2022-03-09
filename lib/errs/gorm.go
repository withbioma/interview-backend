package errs

import (
	"errors"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

const (
	// Complete list of postgres error: https://www.postgresql.org/docs/9.3/errcodes-appendix.html
	pgNotNullConstraintViolation = "23502"
	pgForeignKeyViolation        = "23503"
	pgUniqueConstraintViolation  = "23505"
)

// IsGormNotFound returns true if error is related to gorm.ErrRecordNotFound.
func IsGormNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// TryConvertPostgresError best effort conversion to pg.Error.
func TryConvertPostgresError(err error) (*pgconn.PgError, bool) {
	pgError, ok := err.(*pgconn.PgError)
	return pgError, ok
}

// IsPostgresUniqueConstraintViolationError returns true if error is related to unique constraint.
func IsPostgresUniqueConstraintViolationError(err error) bool {
	pqerr, ok := TryConvertPostgresError(err)
	if !ok {
		return false
	}
	return pqerr.Code == pgUniqueConstraintViolation
}

// IsPostgresNotNullConstraintViolationError returns true if error is related to not null constraint.
func IsPostgresNotNullConstraintViolationError(err error) bool {
	pqerr, ok := TryConvertPostgresError(err)
	if !ok {
		return false
	}
	return pqerr.Code == pgNotNullConstraintViolation
}

// IsPostgresForeignKeyViolationError returns true if error is related to foreign key violation.
func IsPostgresForeignKeyViolationError(err error) bool {
	pqerr, ok := TryConvertPostgresError(err)
	if !ok {
		return false
	}
	return pqerr.Code == pgForeignKeyViolation
}
