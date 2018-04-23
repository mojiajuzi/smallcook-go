package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	mojiajuzi "github.com/mojiajuzi/db"
)

var mydb *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/smallfood")

	if err != nil {
		fmt.Print("err")
		return
	}

	mydb = db

	http.HandleFunc("/", category)
	http.HandleFunc("/menu", Menu)
	http.ListenAndServe(":8080", nil)
}

func category(w http.ResponseWriter, r *http.Request) {
	rows, err := mydb.Query("select id, category_name, category_desc from menu_categories")
	if err != nil {
		fmt.Print(err)
		return
	}

	defer rows.Close()

	var categoryList []mojiajuzi.Category

	for rows.Next() {
		var category mojiajuzi.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Desc); err != nil {
			log.Fatal(err)
		}

		categoryList = append(categoryList, category)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(categoryList)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}

func Menu(w http.ResponseWriter, r *http.Request) {

	categoryId := r.FormValue("category_id")

	rows, err := mydb.Query("select id, menu_name,menu_time from menus where category_id=?", categoryId)

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	var menus []mojiajuzi.Menu

	for rows.Next() {
		var menu mojiajuzi.Menu
		if err = rows.Scan(&menu.ID, &menu.Name, &menu.Time); err != nil {
			log.Fatal(err)
		}

		menus = append(menus, menu)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	res, err := json.Marshal(menus)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(res)
}
