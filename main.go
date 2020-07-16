package main

import(
	"database/sql"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DbConnection is a Connection pool which can be accessed from anywhere
var DbConnection *sql.DB

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql") // DBをOpen()
	defer DbConnection.Close() // あとでClose()
	cmd := `CREATE TABLE IF NOT EXISTS person(
						name STRING,
						age INT)`
	_, err := DbConnection.Exec(cmd) // commandを作って、Exec(command)で実行する _ はほとんどの場合実行結果を受け取る必要がないため
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)" // ?で後から渡す
	// _, err = DbConnection.Exec(cmd, "Nancy", 19)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 19, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Multiple Select
	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd) // QUERY
	// defer rows.Close()
	// 	var pp []Person
	// 	for rows.Next() { // Next()でrowが存在すればtrueを返し、そうでなければfalseを返す
	// 		var p Person
	// 		err := rows.Scan(&p.Name, &p.Age) // Scan() がstructにデータを入れてくれる
	// 		if err != nil { // 一回一回エラーチェックするか、☆1のようにまとめてチェックさせるか
	// 			log.Println(err)
	// 		}
	// 		pp = append(pp, p)
	// 	}
	// 	err = rows.Err()
	// 	if err != nil { // ☆1
	// 		log.Fatalln(err)
	// 	}

	// 	for _, p := range pp { // ただまとめて表示
	// 		fmt.Println(p.Name, p.Age)
	// 	}

// 	// Single Select
// 	cmd = "SELECT * FROM person where age = ?"
// 	row := DbConnection.QueryRow(cmd, 19) // QueryRow()になる
// 	var p Person
// 	err = row.Scan(&p.Name, &p.Age) // ここは同じ
// 	if err != nil {
// 		if err == sql.ErrNoRows { // sqlのErrNoRowsと一致
// 			log.Println("No Row")
// 		} else {
// 			log.Println(err)
// 		}
// 	}
// 	fmt.Println(p.Name, p.Age) // 重複した結果でも１つのみが結果として出力される

	// // DELETE
	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// table名は?で入力できないため、tableNameをユーザーが入力できるような場合、person; INSERT INTO person (name, age) VALUES ('Mr.X', 100); のようにsql injectionされる場合があるので注意する
	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
}

// Person is a struct for QUERY
type Person struct {
	Name string
	Age int
}