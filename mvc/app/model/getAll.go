package getall

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	conn "app/package/dbcon"
)

type Content struct {
	Id int `json:"id"`
	Content string `json:"content"`
}

func GetAllModelList() (Content) {
	var content Content

	if conn.DBConnectAccessAll() {
		defer conn.DBcon.Close()

		selectState, err2 := conn.DBcon.Query("SELECT * FROM model")
		if err2 != nil {
			fmt.Println(err2.Error())
			os.Exit(1)
		}
		defer selectState.Close()

		// Continue to loop if there is a newt row
		for selectState.Next() {
			err := selectState.Scan(&content.Id, &content.Content)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("Failed to GetAllModelList")
	}
	return content
}

func UpdateRowModel(id int, content string) (int64){
	var rowAffected int64
	if conn.DBConnectAccessAll() {
		defer conn.DBcon.Close()
		query := `UPDATE model SET content = ? WHERE id = ?`
		stmt, err := conn.DBcon.Prepare(query)
		if err != nil {
			fmt.Println("Error while updating")
		}
		result, err := stmt.Exec(content, id)
		if err != nil {
			fmt.Println("Error Exec")
		}
		rowAffected, err2 := result.RowsAffected()
		if err2 != nil {
			fmt.Println("Error getting rows affect")
		}
		return rowAffected
	} else {
		fmt.Println("Falied to connect database -> UpdateRowModel")
	}
	return rowAffected
}


func GetSingleModelList() (Content) {
	var content Content

	if conn.DBConnectAccessAll() {
		defer conn.DBcon.Close()

		err := conn.DBcon.QueryRow("SELECT id, content FROM model where id = ?", 1).Scan(&content.Id, &content.Content)
		if err != nil {
			fmt.Println("Error getting a single row")
		}
	}
	return content
}
