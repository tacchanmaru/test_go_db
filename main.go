package main

import (
	"database/sql"
	// "dbsample/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 現在の良いね数を取得するクエリの実行
	article_id := 1
	const sqlGetNice = `
		select nice from articles where article_id = ?
	`

	row := tx.QueryRow(sqlGetNice, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// 変数nicenumに現在の良いね数を読み込む
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// 良いね数を1増やす
	const sqlUpdateNice = `
		update articles set nice = ? where article_id = ?
	`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()

}