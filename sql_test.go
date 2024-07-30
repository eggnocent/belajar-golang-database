package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('5', 'hamid', NULL, 10000, 5.0, NULL, true)"
	//"UPDATE customer SET name = 'uzan' WHERE id = 'dino'"
	//"DELETE FROM customer WHERE id = 'dino'"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("success import new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id,name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("nama:", name)
	}
}

func TestQuesrySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, create_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var create_at time.Time
		var married bool

		rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &create_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("=====================")
		fmt.Println("id:", id)
		fmt.Println("nama:", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		if birth_date.Valid {
			fmt.Println("birth date:", birth_date.Time)
		}
		fmt.Println("married:", married)
		fmt.Println("Dibuat pada:", create_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "EgiWira"
	password := "aya"

	script := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}
