package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Ping_20200310_085612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Ping_20200310_085612{}
	m.Created = "20200310_085612"

	migration.Register("Ping_20200310_085612", m)
}

// Run the migrations
func (m *Ping_20200310_085612) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE ping(`id` int(11) NOT NULL AUTO_INCREMENT,`content` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Ping_20200310_085612) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `ping`")
}
