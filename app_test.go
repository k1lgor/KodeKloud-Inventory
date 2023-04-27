package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(DbUser, DbPassword, "test")
	if err != nil {
		log.Fatal("Error occured while initialising the database")
	}
	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		quantity INT,
		price FLOAT(10,7),
		PRIMARY KEY (id)
	);`

	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER TABLE products AUTO_INCREMENT=1")
	log.Println("clearTable")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES('%v', %v, %v)", name, quantity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Println(err)
	}
}
func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 100, 500)
	req, _ := http.NewRequest("GET", "/product/1", nil)
	res := sendRequest(req)
	checkStatusCode(t, http.StatusOK, res.Code)
}

func checkStatusCode(t *testing.T, expectedStatusCode, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}
func sendRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	a.Router.ServeHTTP(rec, req)
	return rec
}