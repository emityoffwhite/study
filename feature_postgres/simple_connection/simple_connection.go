package simple_connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func CheckConnection() {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil { // вызываем Ping() с контекстом
		panic(err)
	}
	fmt.Println("Подключение к базе данных произошло успешно!")
}
