package main

import (
	"fmt"
	"strings"
)

func dropAndReload() {
	db, _ := getDB()
	db.DropTable(&PtCategory{})
	db.DropTable(&PtUser{})
	db.DropTable(&PtPrediction{})
	db.DropTable(&PtVote{})
	SetUpDB(db)
}

func prompt(p string) bool {
	fmt.Println(p)
	var res string
	fmt.Scanln(&res)
	res = strings.ToLower(res)
	if res == "y" || res == "yes" {
		return true
	}
	return false
}

func main() {
	if prompt("Do you want to drop the old tables and reload") {
		dropAndReload()
	}
	loadCategoreis()
	loadUsers()
	loadPundits()
	loadCalls()
	loadVotes()
}
