package simple_connection

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {

	connString := os.Getenv("CONN_STRING")

	return pgx.Connect(ctx, connString)

}
