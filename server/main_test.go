package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var app App

func TestHello(t *testing.T) {
	mockResponse := `{"message":"helloworld"}`
    r := gin.Default()
    r.GET("/", HelloWorld)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}


func TestMain(m *testing.M){
    
    app.Initialize(
        os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
    )
    ensureTableExists()
    code := m.Run()
    clearTable()
    os.Exit(code)
}


func ensureTableExists() {
    if _, err := app.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`


func clearTable(){
    app.DB.Exec("DELETE FROM products")
    app.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}