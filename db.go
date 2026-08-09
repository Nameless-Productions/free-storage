package main

import (
	"encoding/json"
	"log"
	"os"
)

type File struct {
	Name string `json:"name"`
	MessageIDs []string `json:"messageIDs"`
}

type DB struct {
	Files []File `json:"files"`
}

func addMsgID(fileName, messageID string) {
	db, err := os.ReadFile("db.json")
	if err != nil {
		data, _ := json.Marshal(DB{Files: []File{}})
		os.WriteFile("db.json", []byte(data), 0644)
		db = []byte(data)
	}

	var dbData DB
	err = json.Unmarshal(db, &dbData)
	if err != nil {
		log.Fatal(err)
	}

	var foundIndex *int
	for i, f := range dbData.Files {
		if f.Name == fileName {
			foundIndex = &i
			break
		}
	}

	if foundIndex == nil {
		dbData.Files = append(dbData.Files, File{Name: fileName, MessageIDs: []string{messageID}})
		data, _ := json.Marshal(dbData)
		os.WriteFile("db.json", []byte(data), 0644)
		return
	}

	dbData.Files[*foundIndex].MessageIDs = append(dbData.Files[*foundIndex].MessageIDs, messageID)
	
	data, _ := json.Marshal(dbData)
	os.WriteFile("db.json", []byte(data), 0644)
}

func existsInDB(fileName string) bool {
	db, err := os.ReadFile("db.json")
	if err != nil {
		data, _ := json.Marshal(DB{Files: []File{}})
		os.WriteFile("db.json", []byte(data), 0644)
		db = []byte(data)
	}

	var dbData DB
	err = json.Unmarshal(db, &dbData)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range dbData.Files {
		if f.Name == fileName {
			return true
		}
	}
	return false
}