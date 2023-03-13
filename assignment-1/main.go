package main

import (
	model "Golang/assignment-1/models"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type User struct {
	Name    string
	Surname string
	Login   string
}
type Item struct {
	Name    string
	Country string
	Price   int
}

var Items = []Item{}

func ServeError(err error) {
	if err != nil {
		panic(err)
	}
}

func ClearCash(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/index.html")
	ServeError(err)
	ClearCash(w)
	tmpl.Execute(w, nil)
}
func register(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/register.html")
	ServeError(err)
	ClearCash(w)
	tmpl.Execute(w, nil)
}
func home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/home.html")
	ServeError(err)
	ClearCash(w)

	db, err := model.GetDB()
	ServeError(err)
	defer db.Close()
	rows, err := db.Query("select * from items")
	ServeError(err)
	Items = []Item{}
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.Name, &item.Country, &item.Price)
		ServeError(err)
		Items = append(Items, item)
	}
	tmpl.Execute(w, Items)
}
func homeFilter(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/home.html")
	ServeError(err)
	ClearCash(w)

	db, err := model.GetDB()
	ServeError(err)
	defer db.Close()
	rows, err := db.Query("select * from items order by price")
	ServeError(err)
	Items = []Item{}
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.Name, &item.Country, &item.Price)
		ServeError(err)
		Items = append(Items, item)
	}
	tmpl.Execute(w, Items)
}
func homeName(w http.ResponseWriter, r *http.Request) {
	pat := r.FormValue("pattern")
	tmpl, err := template.ParseFiles("template/home.html")
	ServeError(err)
	ClearCash(w)

	db, err := model.GetDB()
	ServeError(err)
	defer db.Close()
	pat = strings.ToLower(pat)
	rows, err := db.Query(fmt.Sprintf("select * from items where lower(name) like '%s'", pat+"%"))
	ServeError(err)
	Items = []Item{}
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.Name, &item.Country, &item.Price)
		ServeError(err)
		Items = append(Items, item)
	}
	tmpl.Execute(w, Items)
}
func wrongRegister(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/wrong-register.html")
	ServeError(err)
	ClearCash(w)
	tmpl.Execute(w, nil)
}
func addItem(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/add-item.html")
	ServeError(err)
	ClearCash(w)
	tmpl.Execute(w, nil)
}
func wrongLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/wrong-login.html")
	ServeError(err)
	ClearCash(w)
	tmpl.Execute(w, nil)
}
func wrongAddItem(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/wrong-add-item.html")
	ServeError(err)
	ClearCash(w)
	tmpl.Execute(w, nil)
}
func saveData(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	login := r.FormValue("login")
	password := r.FormValue("password")
	db, err := model.GetDB()
	ServeError(err)

	if name == "" || surname == "" || login == "" || password == "" {
		http.Redirect(w, r, "/wrong-register", http.StatusSeeOther)
	} else {
		_, err := db.Exec(fmt.Sprintf("insert into users values ('%s', '%s', '%s', '%s')", name, surname, login, password))
		ServeError(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func saveItem(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	country := r.FormValue("country")
	price := r.FormValue("price")
	db, err := model.GetDB()
	ServeError(err)

	if name == "" || country == "" || price == "" {
		http.Redirect(w, r, "/wrong-add-item", http.StatusSeeOther)
	} else {
		_, err := db.Exec(fmt.Sprintf("insert into items values ('%s', '%s', '%s')", name, country, price))
		ServeError(err)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	ServeError(err)
	itemName := ""
	for key, _ := range r.Form {
		itemName = key
		break
	}
	db, err := model.GetDB()
	ServeError(err)
	_, err = db.Exec(fmt.Sprintf("delete from items where name='%s'", itemName))
	ServeError(err)
	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
func enter(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	db, err := model.GetDB()
	ServeError(err)
	rows, err := db.Query(fmt.Sprintf("select name, surname, login from users where login='%s' and password='%s'", login, password))
	ServeError(err)
	defer rows.Close()
	user := User{}
	cnt := 0
	for rows.Next() {
		cnt++
		err := rows.Scan(&user.Name, &user.Surname, &user.Login)
		ServeError(err)
		break
	}
	if cnt == 0 {
		http.Redirect(w, r, "/wrong-login", http.StatusSeeOther)
	}
	//http.Redirect(w, r, "/home", http.StatusSeeOther)

	http.Redirect(w, r, fmt.Sprintf("/home?name=%s&surname=%s&login=%s", user.Name, user.Surname, user.Login), http.StatusSeeOther)
}

func HandleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/save-data", saveData)
	http.HandleFunc("/save-item", saveItem)
	http.HandleFunc("/delete-item", deleteItem)
	http.HandleFunc("/enter", enter)
	http.HandleFunc("/wrong-register", wrongRegister)
	http.HandleFunc("/wrong-login", wrongLogin)
	http.HandleFunc("/wrong-add-item", wrongAddItem)
	http.HandleFunc("/home", home)
	http.HandleFunc("/home-filter", homeFilter)
	http.HandleFunc("/home-name", homeName)
	http.HandleFunc("/add-item", addItem)

	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleFunc()
}
