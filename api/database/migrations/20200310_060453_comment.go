package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Comment_20200310_060453 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Comment_20200310_060453{}
	m.Created = "20200310_060453"

	migration.Register("Comment_20200310_060453", m)
}

// Run the migrations
func (m *Comment_20200310_060453) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comment(`id` int(11) NOT NULL AUTO_INCREMENT,`content` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Comment_20200310_060453) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `comment`")
}
