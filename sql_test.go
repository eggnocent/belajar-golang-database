package belajargolangdatabase

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('dino', 'hamid')"
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
		var id, name, email string
		var balance int32
		var rating float64
		var birth_date, create_at time.Time
		var married bool

		rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &create_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("=====================")
		fmt.Println("id:", id)
		fmt.Println("nama:", name)
		fmt.Println("email:", email)
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("birth date:", birth_date)
		fmt.Println("married:", married)
		fmt.Println("Dibuat pada:", create_at)
	}
}
