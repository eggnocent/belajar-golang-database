package belajargolangdatabase

import (
	"context"
	"fmt"
	"testing"
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
