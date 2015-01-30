package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	//loadCategoreis()
	//loadUsers()
	loadCalls()

}

func toJsonFromFile(filename string) []map[string]interface{} {
	data, err := ioutil.ReadFile(filename)
	oldStr := []byte(`\'`)
	newStr := []byte(`'`)
	data = bytes.Replace(data, oldStr, newStr, -1)
	oldStr = []byte("\t")
	newStr = []byte(` `)
	data = bytes.Replace(data, oldStr, newStr, -1)
	categories := make([]map[string]interface{}, 0)
	if err != nil {
		fmt.Println("in file: ", filename, err)
		return categories
	}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		fmt.Println("error unmarshalling ", filename, err)
	}
	return categories
}

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

func getStringOrEmpty(i interface{}) string {
	if v, ok := i.(string); ok {
		return v
	} else {
		return ""
	}
}

func parseOldDateFormat(ti interface{}) time.Time {
	form := "2006-01-02 15:04:05"
	t := ti.(string)
	timeCreated, err := time.Parse(form, t)
	if err != nil {
		fmt.Println(err)
		timeCreated = time.Now()
	}
	return timeCreated
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
