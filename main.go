package main

import (
	"database/sql"
	"fmt"
	"log"

	// postgres ドライバ
	_ "github.com/lib/pq"
)

// TestUser : テーブルデータ
type TestUser struct {
	UserID   int
	Password string
}

// メイン関数
func main() {
	// Db: データベースに接続するためのハンドラ
	var Db *sql.DB
	// Dbの初期化
	Db, err := sql.Open("postgres", "host=localhost user=postgres password=mysecretpassword dbname=signage sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// userテーブルから全件取得します
	rows, err := Db.Query("SELECT * FROM TEST_USER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 取得したデータを処理します
	for rows.Next() {
		var user_id int
		var user_password string
		//var email string
		// 必要なカラムに対応する変数を定義します

		err := rows.Scan(&user_id, &user_password)
		if err != nil {
			log.Fatal(err)
		}
		// 取得したデータを利用して処理を行います
		fmt.Printf("ID: %d, Username: %s\n", user_id, user_password)
	}

	// エラーチェック
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
