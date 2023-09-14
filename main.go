package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

func randomEmailAddress() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	randomChars := make([]byte, 6)
	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(alphabet))
		randomChars[i] = alphabet[randomIndex]
	}
	randomString := string(randomChars)
	return randomString + "@example.com"
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	uuidObj := uuid.New()
	uuidStr := uuidObj.String()

	sqlStatement := fmt.Sprintf("INSERT INTO users (email) VALUES ($1) /* UUID:%s */;", uuidStr)
	_, err = db.Exec(sqlStatement, randomEmailAddress())
	if err != nil {
		panic(err)
	}
}
