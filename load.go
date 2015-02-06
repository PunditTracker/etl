package main

import (
	"fmt"
)

var (
	categoryIds = make(map[int64]struct{})
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
			id := int64(c["id"].(float64))
			db.Create(&PtCategory{
				Id:   id,
				Name: name,
			})
			categoryIds[id] = struct{}{}
		} else {
			fmt.Println("add subcat")
			name := c["name"].(string)
			pid := int64(c["parent_id"].(float64))

			db.Create(&PtSubcategory{
				Id:          int64(c["id"].(float64)),
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

		timeCreated := parseOldDateFormat(user["created"])
		id := int64(user["id"].(float64))

		newUser := PtUser{
			Id:                id,
			FirstName:         getStringOrEmpty(user["first_name"]),
			LastName:          getStringOrEmpty(user["last_name"]),
			Email:             getStringOrEmpty(user["email"]),
			FacebookId:        getStringOrEmpty(user["fb_id"]),
			FacebookAuthToken: getStringOrEmpty(user["fb_access_token"]),
			Password:          "NONE",
			Avatar_URL:        getStringOrEmpty(user["avatar"]),
			Created:           timeCreated,
		}
		db.Create(&newUser)
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
		_, exists := categoryIds[catId]
		var newPrediction PtPrediction
		if exists {
			newPrediction = PtPrediction{
				Id:         id,
				CreatorId:  creId,
				CategoryId: catId,
				Deadline:   parseOldDateFormat(c["approval_time"]),
				Created:    parseOldDateFormat(c["created"]),
				Title:      getStringOrEmpty(c["prediction"]),
			}
		} else {
			newPrediction = PtPrediction{
				Id:        id,
				CreatorId: creId,
				SubcatId:  catId,
				Deadline:  parseOldDateFormat(c["approval_time"]),
				Created:   parseOldDateFormat(c["created"]),
				Title:     getStringOrEmpty(c["prediction"]),
			}
		}

		db.Create(&newPrediction)

	}
}
