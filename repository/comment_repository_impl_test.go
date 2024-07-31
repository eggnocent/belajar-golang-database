package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repositoryyy@test.go",
		Comment: "Test repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 44)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	commentalls, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	// lihat ada berapa dengan perulangan
	for _, commentall := range commentalls {
		fmt.Println(commentall)
	}

	//lihat dengan isi nya
	//fmt.Println(commentalls)

}
