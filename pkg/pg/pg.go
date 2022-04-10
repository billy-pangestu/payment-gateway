package pg

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

// Connection ...
type Connection struct {
	Host    string
	DB      string
	User    string
	Pass    string
	Port    int
	Loc     *time.Location
	SslMode string
}

// Connect ...
func (m Connection) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=UTC",
		m.User, m.Pass, m.Host, m.Port, m.DB, m.SslMode,
	)

	db, err := sql.Open("postgres", connStr)

	return db, err
}
