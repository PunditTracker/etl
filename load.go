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

		name := c["name"].(string)
		fmt.Println("add cat")
		id := int64(c["id"].(float64))
		db.Create(&PtCategory{
			Id:   id,
			Name: name,
		})
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
			Id:                 id,
			FirstName:          getStringOrEmpty(user["first_name"]),
			LastName:           getStringOrEmpty(user["last_name"]),
			Email:              getStringOrEmpty(user["email"]),
			FacebookId:         getStringOrEmpty(user["fb_id"]),
			FacebookAuthToken:  getStringOrEmpty(user["fb_access_token"]),
			PredictionsGraded:  getNumOrZero(user["calls_graded"]),
			PredictionsCorrect: getNumOrZero(user["calss_correct"]),
			Password:           "NONE",
			AvatarUrl:          getStringOrEmpty(user["avatar"]),
			Created:            timeCreated,
		}
		db.Create(&newUser)
	}
}

func loadCalls() {
	db, _ := getDB()
	calls := toJsonFromFile("calls.json")
	fmt.Println(calls[0])
	//	size := len(calls)

	for _, c := range calls {

		id := int64(c["id"].(float64))
		creId := int64(c["user_id"].(float64))
		catId := int64(c["category_id"].(float64))

		var newPrediction PtPrediction
		newPrediction = PtPrediction{
			Id:         id,
			CreatorId:  creId,
			CategoryId: catId,
			Deadline:   parseOldDateFormat(c["approval_time"]),
			Created:    parseOldDateFormat(c["created"]),
			Title:      getStringOrEmpty(c["prediction"]),
		}
		db.Create(&newPrediction)

	}
}
