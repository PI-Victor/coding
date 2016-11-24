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
	GetDriver() dbDriver
}


type dbHandler string

type dbDriver string

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

func (m *mongoDBInstance) GetDriver() dbDriver {
	//	return m.driver
	var drv dbDriver
	drv = "N/A"
	return drv
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
	var drv dbDriver
	drv = "test"
	return &redisDBInstance{
		driver: drv,
	}
}	

func newMongoDBInstance() *mongoDBInstance{
	return &mongoDBInstance{
	}
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
	fmt.Println(newRedis.GetDriver())
	fmt.Println(newMongo.GetDriver())
}


func regDBHandler(dbInst dbInstance) {
	instance := reflect.TypeOf(dbInst)
	dbInst.GetDriver()
	fmt.Println(instance)
}
