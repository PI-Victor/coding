// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2016-11-23 14:57:47
//
// run snippet with Ctl-Return

package main

import (
	"fmt"
	"reflect"
)

type dbInstance interface {
	Connect()
	Disconnect()
}

type (
	dbHandler string
	dbDriver string
)

type mongoDBInstance struct {
	handler dbHandler
	driver dbDriver
}

func (m *mongoDBInstance) Connect() {
	m.handler = "test"
}

func (m *mongoDBInstance) Disconnect() {
	m.handler = ""
}

type redisDBInstance struct  {
	handler dbHandler
	URI string
	driver dbDriver
}

func (r *redisDBInstance) Connect() {
	r.handler = "test"
}

func (r *redisDBInstance) Disconnect() {
	r.handler = ""
}

func (r *redisDBInstance) GetDriver() dbDriver{
	return r.driver 
}

func newRedisDBInstance() *redisDBInstance{
	return &redisDBInstance{
		driver: "N/A",
	}
}	

func newMongoDBInstance() *mongoDBInstance{
	return &mongoDBInstance{}
}

func main() {
	// NOTE: when passing a type that implements an certain interface, to a
	// fn that expects said interface, only the methods that the interface
	// defines are available at runtime.
	// e.g.: If the type has more functions, regDBHandler doesn't recognize
	// them because it expects an interface that has only 2 functions.
	newMongo := newMongoDBInstance()
	newRedis := newRedisDBInstance()
	regDBHandler(newRedis)
	regDBHandler(newMongo)
}


func regDBHandler(dbInst dbInstance) {
	instance := reflect.TypeOf(dbInst)
	// Here, dbInst is not available because the interface doesn't contain
	// that function.
	// dbInst.GetDriver()
	fmt.Println(instance)
}
