package simple_connection

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {

	conn_String := os.Getenv("conn_string")

	return pgx.Connect(ctx, conn_String)

}
