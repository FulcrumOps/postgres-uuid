package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

var log = logrus.New()

func getRandomEmailAddress() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	randomChars := make([]byte, 6)
	for i := 0; i < 6; i++ {
		randomIndex := r.Intn(len(alphabet))
		randomChars[i] = alphabet[randomIndex]
	}
	randomString := string(randomChars)
	return randomString + "@example.com"
}

func getNewUUID() string {
	uuidObj := uuid.New()
	uuidStr := uuidObj.String()
	return uuidStr
}

func getBunDB() *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func insertViaRawSQL() error {
	log.Info("Inserting a new user via Raw SQL")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := fmt.Sprintf("/* UUID: %s */ INSERT INTO users (email) VALUES ($1);", getNewUUID())
	log.Info(sqlStatement)

	_, err = db.Exec(sqlStatement, getRandomEmailAddress())
	if err != nil {
		return err
	}
	return nil
}

func insertViaRawBun() error {
	log.Info("Inserting a new user via Raw Bun")

	db := getBunDB()
	defer db.Close()

	sqlStatement := fmt.Sprintf("/* UUID: %s */ INSERT INTO users (email) VALUES (?);", getNewUUID())
	log.Info(sqlStatement)

	if _, err := db.NewInsert().NewRaw(sqlStatement, getRandomEmailAddress()).Exec(context.Background()); err != nil {
		return err
	}
	return nil
}

func insertViaNativeBun() error {
	log.Info("Inserting a new user via Native Bun")

	db := getBunDB()
	defer db.Close()

	type User struct {
		Email string
	}
	user := &User{Email: getRandomEmailAddress()}

	uuid := getNewUUID()
	log.Infof("UUID: %s	Inserting new user into db", uuid)
	if _, err := db.NewInsert().Model(user).Exec(context.Background()); err != nil {
		log.Errorf("UUID: %s	Failed to insert user: %v", uuid, err)
		return err
	}
	return nil
}

func littleBobbyTables() error {
	fmt.Print("Do you want to run an unprepared statement (y/n): ")
	var input string
	fmt.Scanln(&input)
	if input == "y" {

		log.Info("Inserting a bad user via Raw SQL without prepared statements")
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
		if err != nil {
			return err
		}
		defer db.Close()

		sqlStatement := fmt.Sprintf("/* UUID: %s */ INSERT INTO users (email) VALUES ($1);", getNewUUID())
		log.Info(sqlStatement)

		emailAddress := `Little Bobby Tables <bobby@'); DROP TABLE users; -->`

		// Since this is not a prepared statement, this will drop the users table
		sqlStatement = fmt.Sprintf("/* UUID: %s */ INSERT INTO users (email) VALUES ('%s');", getNewUUID(), emailAddress)
		log.Info(sqlStatement)

		_, err = db.Exec(sqlStatement)
		if err != nil {
			return err
		}
	}

	err := insertViaRawSQL()
	if err != nil {
		return err
	}
	return nil
}

func main() {

	err := insertViaRawSQL()
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	err = insertViaRawBun()
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	err = insertViaNativeBun()
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	err = littleBobbyTables()
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

}
