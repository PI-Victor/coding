// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2016-12-05 14:07:22
//
// run snippet with Ctl-Return

package main

import (
	"fmt"
)

func main() {
	dbHandle := validateDB()
	fmt.Printf("%#v", dbHandle)
}

type Handler string

type Database interface {
	GetHandler() Handler
	Disconnect() error
	GetDetails() (string, string, Handler)
}

type DBInstance struct {
	handler Handler
	session string
	driver string
}

func newDBInstance(db Database) *DBInstance {
	session, driver, handler := db.GetDetails()
	return &DBInstance{
		handler: handler,
		session: session,
		driver: driver,
	}
}

func validateDB() *DBInstance {
	// function retry logic should return a handler to a specific db
	// that we then pass to instantiate
	mDB := NewMongo()
	mDB.UpdateHandler("")
	sDB := NewMySQL()

	if mDB.handler != "" {
		return newDBInstance(mDB)
	}

	if sDB.handler != "" {
		return newDBInstance(sDB)
	}
	
	return nil
}

func NewMongo() *MongoDB {
	return &MongoDB {
		handler: "test",
		session: "test1",
		driver: "driver",
	}
}

func NewMySQL() *MySQL {
	return &MySQL {
		handler: "test",
		session: "test1",
		driver: "mongo",
	}
}

type MongoDB struct {
	handler Handler
	session string
	driver string
}

func (m *MongoDB) GetDetails() (string, string, Handler) {
	return m.session, m.driver, m.handler
}

func (m *MongoDB) UpdateHandler(handler Handler) {
	m.handler = handler
}

func (m *MongoDB) GetHandler() Handler {
	return m.handler
}

func (m *MongoDB) Disconnect() error {
	m.session = ""
	return nil
}

type MySQL struct {
	handler Handler
	session string
	driver string
}

func (s *MySQL) GetDetails() (string, string, Handler) {
	return s.session, s.driver, s.handler
}
func (s *MySQL) UpdateHandler(handler Handler) {
	s.handler = handler
}

func (s *MySQL) GetHandler() Handler {
	return s.handler
}

func (s *MySQL) Disconnect() error {
	s.session = ""
	return nil
}
