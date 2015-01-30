package main

import (
	_ "fmt"
)

func init() {
	db, _ := getDB()
	db.DropTable(&PtCategory{})
	db.DropTable(&PtSubcategory{})
	db.DropTable(&PtUser{})
	db.DropTable(&PtPrediction{})
	SetUpDB(db)
}

func main() {
	loadCategoreis()
	loadUsers()
	loadCalls()
}
