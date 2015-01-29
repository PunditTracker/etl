package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	loadCategoreis()

}

func toJsonFromFile(filename string) []map[string]interface{} {
	data, err := ioutil.ReadFile(filename)
	categories := make([]map[string]interface{}, 0)
	if err != nil {
		fmt.Println(err)
		return categories
	}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		fmt.Println(err)
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
