package main

import (
	"fmt"
	"study/http_server"
)

/*"postgres://postgres:123@localhost:5432/postgres"*/

func main() {

	fmt.Println("Запуска сервера!")
	fmt.Println("Всё работает исправно")

	err := http_server.StartHTTPServer()
	if err != nil {
		fmt.Println("Во время работы сервера произошла ошибка:", err)
	} else {
		fmt.Println("Сервер завершился успешно")
	}

	/*ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)

	if err != nil {

		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		if task.ID == 1 {
			task.Title = "Процесс"
			task.Description = "Начать что то делать"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			return

		}
	}*/
	/*if err := simple_sql.InsertRow(ctx, conn,
		"Обед",
		"Покушать",
		false,
		time.Now()); err != nil {
		panic(err)
	}*/

	/*if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}*/

	/*if err := simple_sql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}*/

	/*fmt.Println("Успешно!")*/

}
