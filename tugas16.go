package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type kontak struct {
	ID    int
	Nama  string
	Nomor string
}

func main() {
	getAllData()
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/kontak")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func getAllData() {
	db, err := koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM telepon")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []kontak

	for rows.Next() {
		var each = kontak{}

		var err = rows.Scan(&each.ID, &each.Nama, &each.Nomor)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Tampil data di console
	for _, i := range result {
		fmt.Println(i.ID, i.Nama, i.Nomor)
	}
}
