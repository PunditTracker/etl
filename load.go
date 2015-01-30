package main

import (
	"fmt"
)

func loadCategoreis() {
	db, _ := getDB()
	categories := toJsonFromFile("categories.json")
	fmt.Println(categories[0])
	for _, c := range categories {
		fmt.Println(c["parent_id"])
		if c["parent_id"] == 0.0 {
			name := c["name"].(string)
			fmt.Println("add cat")
			db.Save(&PtCategory{
				Name: name,
			})
		} else {
			fmt.Println("add subcat")
			name := c["name"].(string)
			pid := int64(c["parent_id"].(float64))

			db.Save(&PtSubcategory{
				Name:        name,
				ParentCatId: pid,
			})
		}
	}
}

func loadUsers() {
	db, _ := getDB()
	db = db.Debug()
	users := toJsonFromFile("users.json")
	fmt.Println(users[0])
	for _, user := range users {
		uname := getStringOrEmpty(user["slug"])
		fname := getStringOrEmpty(user["first_name"])
		lname := getStringOrEmpty(user["last_name"])
		email := getStringOrEmpty(user["email"])
		av_url := getStringOrEmpty(user["avatar"])

		timeCreated := parseOldDateFormat(user["created"])
		id := int64(user["id"].(float64))

		newUser := PtUser{
			Id:         id,
			Username:   uname,
			FirstName:  fname,
			LastName:   lname,
			Email:      email,
			Password:   "NONE",
			Avatar_URL: av_url,
			Created:    timeCreated,
		}
		db.FirstOrCreate(&newUser)
	}
}

func loadCalls() {
	db, _ := getDB()
	db = db.Debug()
	calls := toJsonFromFile("calls.json")
	fmt.Println(calls[0])
	//	size := len(calls)

	for _, c := range calls {

		id := int64(c["id"].(float64))
		creId := int64(c["user_id"].(float64))
		catId := int64(c["category_id"].(float64))

		newPrediction := PtPrediction{
			Id:         id,
			CreatorId:  creId,
			CategoryId: catId,
			Deadline:   parseOldDateFormat(c["approval_time"]),
			Created:    parseOldDateFormat(c["created"]),
			Title:      getStringOrEmpty(c["prediction"]),
		}
		db.FirstOrCreate(&newPrediction)

	}
}
