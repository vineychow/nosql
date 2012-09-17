package main

import (
	"fmt"
	"sync"

	"engine/mgo"
	"engine/mgo/bson"
)

var (
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.Collection
	mux     sync.Mutex
)

func init() {
	mux.Lock()
	defer mux.Unlock()

	if session != nil {
		return
	}

	// new session
	session, err := mgo.Dial("127.0.0.1:27017")
	checkErr("dial fail", err)

	session.SetMode(mgo.Monotonic, true)

	// new db
	if db != nil {
		return
	}
	db = session.DB("t2m")

	// new collection
	if c != nil {
		return
	}
	c = db.C("person")
}

func checkErr(data string, err error) {
	if err != nil {
		fmt.Println(data + ": " + err.Error())
	}
}

type Person struct {
	Name  string
	Email string
}

func New() *Person {
	return &Person{
		Name:  "viney",
		Email: "viney.chow@gmail.com",
	}
}

func main() {

	// insert
	err := c.Insert(New())
	checkErr("insert fail", err)

	// find
	m := bson.M{"name": "viney"}
	query := c.Find(m)

	// one
	person := Person{}
	err = query.One(&person)
	checkErr("query fail", err)

	fmt.Println(person.Email)

	if session != nil {
		session.Close()
	}
}
