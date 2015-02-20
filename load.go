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
	db = db
	users := toJsonFromFile("users.json")
	fmt.Println(users[0])
	for _, user := range users {

		timeCreated := parseOldDateFormat(user["created"])
		id := int64(user["id"].(float64))
		g := getNumOrZero(user["calls_graded"])
		c := getNumOrZero(user["calls_correct"])
		fmt.Println(g, c)

		newUser := PtUser{
			Id:                 id,
			FirstName:          getStringOrEmpty(user["first_name"]),
			LastName:           getStringOrEmpty(user["last_name"]),
			Email:              getStringOrEmpty(user["email"]),
			FacebookId:         getStringOrEmpty(user["fb_id"]),
			FacebookAuthToken:  getStringOrEmpty(user["fb_access_token"]),
			PredictionsGraded:  g,
			PredictionsCorrect: c,
			Password:           "NONE",
			AvatarUrl:          getStringOrEmpty(user["avatar"]),
			Created:            timeCreated,
		}
		if newUser.Email == "" {
			newUser.Email = "None@None.com"
		}
		db.Create(&newUser)
	}
}

func loadPundits() {
	db, _ := getDB()
	db = db
	pundits := toJsonFromFile("pundits.json")

	for _, pundit := range pundits {
		g := getNumOrZero(pundit["calls_graded"])
		c := getNumOrZero(pundit["calls_correct"])
		fmt.Println(g, c)
		uid := int64(pundit["user_id"].(float64))
		fmt.Println(uid)
		var u PtUser
		db.First(&u, uid)
		u.PredictionsGraded = g
		u.PredictionsCorrect = c
		db.Save(&u)
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

func loadVotes() {
	db, _ := getDB()
	votes := toJsonFromFile("votes.json")
	for _, v := range votes {
		voterId := int64(v["user_id"].(float64))
		predId := int64(v["call_id"].(float64))
		ptvar, isFloat := v["ptvariable"].(float64)
		var voteVal int
		if isFloat {
			if ptvar == 1 {
				voteVal = 4
			} else {
				voteVal = 3
			}
		} else {
			voteVal = 4
		}
		boldness, isFloat := v["boldness"].(float64)
		var avg float64
		if isFloat {
			avg = 1 - boldness
		} else {
			avg = -1
		}
		NewVote := PtVote{
			VoterId:       voterId,
			VotedOnId:     predId,
			AverageAtTime: avg,
			VoteValue:     voteVal,
			Created:       parseOldDateFormat(v["created"]),
		}
		db.Create(&NewVote)
	}
}
