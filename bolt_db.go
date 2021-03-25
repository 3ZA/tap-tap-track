package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

type Entry struct {
	Date    time.Time
	Records map[string]*ActivityRecord
}

type ActivityRecord struct {
	Name string // activity name
	Done bool   // true if activity marked as complete
}

type BoltDB struct {
	*bolt.DB
}

func NewBoltDB() (*BoltDB, error) {
	db, err := setupDB()
	if err != nil {
		return nil, err
	}
	return &BoltDB{db}, nil
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("Habits"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("Entries"))
		if err != nil {
			return fmt.Errorf("could not create entries bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func (db *BoltDB) AddEntry(e *Entry) error {
	entryBytes, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("Habits")).Bucket([]byte("Entries")).Put([]byte(timeToDateString(e.Date)), entryBytes)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}
		return nil
	})
	fmt.Println("Added Entry " + PrettyFormat(e))
	return err
}

func (db *BoltDB) RetrieveEntry(t time.Time) (*Entry, error) {
	entry := &Entry{}
	foundEntry := false
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Habits")).Bucket([]byte("Entries"))
		v := b.Get([]byte(timeToDateString(t)))
		if v == nil {
			return nil
		}
		err := json.Unmarshal(v, entry)
		foundEntry = true
		return err
	})
	if err != nil {
		return nil, err
	}
	if foundEntry == false {
		return nil, nil
	}
	fmt.Println("Retrieved Entry" + PrettyFormat(entry))
	return entry, nil
}

func (db *BoltDB) viewAll() {
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Habits")).Bucket([]byte("Entries"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}

func PrettyFormat(i interface{}) string {
	bytes, _ := json.MarshalIndent(i, "", " ")
	return string(bytes)
}
